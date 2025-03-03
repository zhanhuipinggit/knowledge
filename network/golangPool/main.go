package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Task 定义任务结构体
type Task struct {
	ID   int          // 任务ID
	Data interface{}  // 任务数据
	Proc func() error // 任务处理函数
}

// Pool 协程池结构体
type Pool struct {
	taskChan  chan *Task     // 任务通道
	workerNum int            // 工作协程数量
	wg        sync.WaitGroup // 用于等待任务完成
	quitChan  chan struct{}  // 退出信号
	isClosed  bool           // 池是否已关闭
	mutex     sync.Mutex     // 保护关闭状态
}

// NewPool 创建协程池
func NewPool(workerNum, taskBuffer int) *Pool {
	p := &Pool{
		taskChan:  make(chan *Task, taskBuffer),
		workerNum: workerNum,
		quitChan:  make(chan struct{}),
	}

	p.wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go p.worker()
	}

	return p
}

// worker 工作协程
func (p *Pool) worker() {
	defer p.wg.Done()

	for {
		select {
		case task, ok := <-p.taskChan:
			if !ok { // 通道已关闭
				return
			}
			// 执行任务
			if err := task.Proc(); err != nil {
				fmt.Printf("Task %d failed: %v\n", task.ID, err)
			} else {
				fmt.Printf("Task %d completed\n", task.ID)
			}
		case <-p.quitChan:
			return
		}
	}
}

// Submit 提交任务到池
func (p *Pool) Submit(task *Task) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.isClosed {
		return errors.New("pool is closed")
	}

	select {
	case p.taskChan <- task:
		return nil
	default:
		return errors.New("task queue is full")
	}
}

// Close 优雅关闭协程池
func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.isClosed {
		return
	}

	p.isClosed = true
	close(p.quitChan) // 通知worker停止
	close(p.taskChan) // 关闭任务通道
	p.wg.Wait()       // 等待所有worker退出
}

// 示例使用
func main() {
	// 创建协程池（3个工作协程，任务队列缓冲10）
	pool := NewPool(3, 10)

	// 提交20个任务
	for i := 1; i <= 20; i++ {
		taskID := i
		err := pool.Submit(&Task{
			ID: taskID,
			Proc: func() error {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Processing task %d\n", taskID)
				return nil
			},
		})

		if err != nil {
			fmt.Printf("Submit task %d failed: %v\n", taskID, err)
		}
	}

	// 等待任务处理完成
	time.Sleep(2 * time.Second)

	// 关闭协程池
	pool.Close()
}
