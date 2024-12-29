package main

import (
	"fmt"
	"log"
	"sync"
	"syscall"
)

type RawEpoll struct {
	epfd    int                  // epoll 实例的文件描述符
	events  []syscall.EpollEvent // 事件集合
	mu      sync.Mutex           // 并发保护
	handlers map[int]func()      // 回调函数映射
}

// NewRawEpoll 创建一个新的 RawEpoll 实例
func NewRawEpoll() (*RawEpoll, error) {
	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		return nil, fmt.Errorf("failed to create epoll: %w", err)
	}

	return &RawEpoll{
		epfd:    epfd,
		events:  make([]syscall.EpollEvent, 128), // 初始事件容量
		handlers: make(map[int]func()),
	}, nil
}

// Add 注册文件描述符及其回调函数
func (r *RawEpoll) Add(fd int, events uint32, callback func()) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	event := syscall.EpollEvent{
		Events: events,
		Fd:     int32(fd),
	}

	// 将事件添加到 epoll
	if err := syscall.EpollCtl(r.epfd, syscall.EPOLL_CTL_ADD, fd, &event); err != nil {
		return fmt.Errorf("failed to add fd to epoll: %w", err)
	}

	// 保存回调函数
	r.handlers[fd] = callback
	return nil
}

// Remove 移除文件描述符的监听
func (r *RawEpoll) Remove(fd int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 从 epoll 中移除文件描述符
	if err := syscall.EpollCtl(r.epfd, syscall.EPOLL_CTL_DEL, fd, nil); err != nil {
		return fmt.Errorf("failed to remove fd from epoll: %w", err)
	}

	// 删除回调函数
	delete(r.handlers, fd)
	return nil
}

// Wait 开始监听事件
func (r *RawEpoll) Wait() error {
	for {
		// 等待事件触发
		n, err := syscall.EpollWait(r.epfd, r.events, -1)
		if err != nil {
			if err == syscall.EINTR {
				continue // 如果被信号中断，则重试
			}
			return fmt.Errorf("epoll wait error: %w", err)
		}

		// 遍历触发的事件
		for i := 0; i < n; i++ {
			fd := int(r.events[i].Fd)

			r.mu.Lock()
			callback, exists := r.handlers[fd]
			r.mu.Unlock()

			if exists && callback != nil {
				// 执行回调
				callback()
			}
		}
	}
}

// Close 关闭 epoll 实例
func (r *RawEpoll) Close() error {
	return syscall.Close(r.epfd)
}

func main() {
	// 创建 RawEpoll 实例
	epoll, err := NewRawEpoll()
	if err != nil {
		log.Fatalf("Failed to create RawEpoll: %v", err)
	}
	defer epoll.Close()

	// 创建一个监听 socket
	listener, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("Failed to create socket: %v", err)
	}

	// 绑定地址
	addr := syscall.SockaddrInet4{Port: 8080}
	copy(addr.Addr[:], []byte{127, 0, 0, 1})
	if err := syscall.Bind(listener, &addr); err != nil {
		log.Fatalf("Failed to bind: %v", err)
	}

	// 开始监听
	if err := syscall.Listen(listener, 128); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 设置为非阻塞模式
	if err := syscall.SetNonblock(listener, true); err != nil {
		log.Fatalf("Failed to set non-blocking: %v", err)
	}

	// 注册监听 socket 到 epoll
	err = epoll.Add(listener, syscall.EPOLLIN, func() {
		for {
			connFd, _, err := syscall.Accept(listener)
			if err != nil {
				if err == syscall.EAGAIN {
					break // 没有更多连接
				}
				log.Printf("Failed to accept connection: %v", err)
				continue
			}

			// 设置新连接为非阻塞模式
			if err := syscall.SetNonblock(connFd, true); err != nil {
				log.Printf("Failed to set non-blocking: %v", err)
				syscall.Close(connFd)
				continue
			}

			// 注册新连接到 epoll
			epoll.Add(connFd, syscall.EPOLLIN, func() {
				buf := make([]byte, 1024)
				n, err := syscall.Read(connFd, buf)
				if err != nil {
					log.Printf("Failed to read: %v", err)
					epoll.Remove(connFd)
					syscall.Close(connFd)
					return
				}

				if n == 0 {
					log.Printf("Client disconnected")
					epoll.Remove(connFd)
					syscall.Close(connFd)
					return
				}

				log.Printf("Received: %s", string(buf[:n]))
			})
		}
	})
	if err != nil {
		log.Fatalf("Failed to add listener to epoll: %v", err)
	}

	// 启动事件循环
	log.Println("Server is running on :8080")
	if err := epoll.Wait(); err != nil {
		log.Fatalf("Epoll wait error: %v", err)
	}
}