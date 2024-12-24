# RawEpoll 详解

## 什么是 RawEpoll？

`RawEpoll` 是一种基于 Linux 内核中的 `epoll` 提供的高性能 I/O 多路复用机制的直接实现方式。与常见的高级封装（如 Go 的 `netpoll`、Java 的 `NIO` 等）不同，RawEpoll 通常直接调用 Linux 的 `epoll` 系统调用，绕过额外的抽象层，以获得更高的性能和更细粒度的控制能力。

`epoll` 是 Linux 专门为高并发场景设计的系统调用，它能够高效地监控大量的文件描述符，并只在这些文件描述符状态发生变化时通知应用程序。这使得 `epoll` 非常适合网络服务器、消息队列和其他需要处理大规模并发连接的应用场景。

---

## RawEpoll 的核心概念

### 1. `epoll` 的工作原理

`epoll` 是 Linux 提供的一种 I/O 多路复用机制，它的主要特点是：
- **事件驱动**：只通知那些状态发生变化的文件描述符，而不是轮询所有描述符。
- **高效**：可以监控大规模文件描述符集合，而不会因为文件描述符数量的增加而导致性能下降。

`epoll` 的操作通过以下三个核心函数完成：

- **`epoll_create`**:
  创建一个 epoll 实例，并返回一个文件描述符。

- **`epoll_ctl`**:
  将文件描述符添加、修改或从 epoll 实例中删除。

- **`epoll_wait`**:
  等待事件的发生，并返回已就绪的文件描述符集合。

### 2. `epoll` 的工作模式

- **ET（边沿触发，Edge Triggered）**:
  仅在状态变化时通知一次，需要应用程序处理所有的可用数据。

- **LT（水平触发，Level Triggered）**:
  文件描述符仍处于就绪状态时，每次调用都会通知。默认工作模式。

### 3. 内核实现

`epoll` 使用了一种事件回调机制，基于红黑树存储监控的文件描述符，使用链表存储就绪的事件列表，避免了传统 `select`/`poll` 的线性扫描。

---

## RawEpoll 的主要功能实现

RawEpoll 通常会直接使用 `epoll` 提供的系统调用接口进行开发，以下是 RawEpoll 实现的主要步骤：

### 1. 初始化 epoll 实例
```c
int epfd = epoll_create(1);  // 创建一个 epoll 实例
if (epfd == -1) {
    perror("epoll_create failed");
    return -1;
}
```
`epoll_create` 返回一个 epoll 实例的文件描述符，后续所有操作都通过这个描述符进行。

### 2. 添加文件描述符
```c
struct epoll_event ev;
ev.events = EPOLLIN;  // 监听可读事件
ev.data.fd = sockfd;  // 设置关联的文件描述符
if (epoll_ctl(epfd, EPOLL_CTL_ADD, sockfd, &ev) == -1) {
    perror("epoll_ctl failed");
    return -1;
}
```
通过 `epoll_ctl` 函数可以将文件描述符添加到 epoll 实例中，并指定监听的事件类型（如可读、可写、错误等）。

### 3. 等待事件发生
```c
struct epoll_event events[MAX_EVENTS];  // 存储已就绪的事件
int nfds = epoll_wait(epfd, events, MAX_EVENTS, -1);
if (nfds == -1) {
    perror("epoll_wait failed");
    return -1;
}
for (int i = 0; i < nfds; ++i) {
    int fd = events[i].data.fd;
    // 处理就绪的文件描述符
}
```
`epoll_wait` 阻塞等待事件发生，并返回已就绪的事件数量。

### 4. 修改或移除文件描述符
- 修改文件描述符：
```c
epoll_ctl(epfd, EPOLL_CTL_MOD, sockfd, &ev);
```
- 移除文件描述符：
```c
epoll_ctl(epfd, EPOLL_CTL_DEL, sockfd, NULL);
```

### 5. 关闭 epoll 实例
```c
close(epfd);
```
当不再需要 epoll 实例时，可以关闭文件描述符。

---

## RawEpoll 的特点

1. **高性能**：
    - 支持监控大量文件描述符，性能不会随着文件描述符数量的增加而线性下降。
    - 使用事件回调机制，避免了轮询操作。

2. **灵活性**：
    - 支持多种事件类型（如读、写、异常）。
    - 提供边沿触发（ET）模式，更适合高性能场景。

3. **底层控制**：
    - RawEpoll 直接调用 `epoll` 接口，无额外的封装，可以获得更细粒度的控制权。

---

## RawEpoll 的应用场景

1. **高并发网络服务**：
    - Web 服务器、RPC 框架等高性能网络服务。

2. **消息队列**：
    - 用于高吞吐量的消息生产和消费。

3. **游戏服务器**：
    - 实现大规模连接管理。

4. **日志处理和数据采集**：
    - 监控大量日志文件或数据流的变化。

---

## RawEpoll 的优缺点

### 优点
1. 性能高，特别是在高并发场景下。
2. 灵活性强，可以自定义底层逻辑。
3. 支持大规模文件描述符管理。

### 缺点
1. 实现复杂：需要直接处理底层的事件逻辑，容易出错。
2. 跨平台性差：`epoll` 是 Linux 特有的系统调用，在其他平台（如 Windows、BSD 系统）上无法使用。
3. 开发维护成本高：与高级封装相比，需要更多的开发和维护工作。

---

## RawEpoll 的优化建议

1. **多线程处理**：
    - 使用多个线程处理不同的 epoll 实例，充分利用多核 CPU。

2. **负载均衡**：
    - 将文件描述符分配到不同的 epoll 实例中，避免单个实例过载。

3. **内存管理**：
    - 避免频繁的内存分配和释放操作，可以预分配事件数组或对象池。

4. **ET 模式下的优化**：
    - 在边沿触发模式下，确保一次性读取所有数据，减少事件触发次数。

---

## 代码demo
**rawepoll实现一个简单的ws服务器**
```golang
package main

import (
	"log"
	"net"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"golang.org/x/sys/unix"
)

const (
	MaxEvents = 1024 // 最大事件数
)

// WebSocket 连接管理器
type ConnectionManager struct {
	connections map[int]net.Conn
	lock        sync.Mutex
}

func (cm *ConnectionManager) Add(fd int, conn net.Conn) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.connections[fd] = conn
}

func (cm *ConnectionManager) Remove(fd int) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if conn, ok := cm.connections[fd]; ok {
		conn.Close()
		delete(cm.connections, fd)
	}
}

func (cm *ConnectionManager) Get(fd int) (net.Conn, bool) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	conn, ok := cm.connections[fd]
	return conn, ok
}

func main() {
	// 创建 epoll 实例
	epfd, err := unix.EpollCreate1(0)
	if err != nil {
		log.Fatalf("failed to create epoll instance: %v", err)
	}
	defer unix.Close(epfd)

	// 监听地址
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}
	defer listener.Close()
	log.Println("WebSocket server started on :8080")

	// 初始化连接管理器
	connManager := &ConnectionManager{
		connections: make(map[int]net.Conn),
	}

	// 创建事件数组
	events := make([]unix.EpollEvent, MaxEvents)

	// 接受连接的 Goroutine
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("failed to accept connection: %v", err)
				continue
			}

			// 设置非阻塞模式
			sockfd := int(conn.(*net.TCPConn).File().Fd())
			if err := unix.SetNonblock(sockfd, true); err != nil {
				log.Printf("failed to set non-blocking mode: %v", err)
				conn.Close()
				continue
			}

			// 将连接添加到 epoll
			event := unix.EpollEvent{
				Events: unix.EPOLLIN | unix.EPOLLET,
				Fd:     int32(sockfd),
			}
			if err := unix.EpollCtl(epfd, unix.EPOLL_CTL_ADD, sockfd, &event); err != nil {
				log.Printf("failed to add fd to epoll: %v", err)
				conn.Close()
				continue
			}

			// 添加到连接管理器
			connManager.Add(sockfd, conn)
		}
	}()

	// 主事件循环
	for {
		nfds, err := unix.EpollWait(epfd, events, -1)
		if err != nil {
			if err == unix.EINTR {
				continue
			}
			log.Fatalf("epoll_wait error: %v", err)
		}

		for i := 0; i < nfds; i++ {
			fd := int(events[i].Fd)
			conn, ok := connManager.Get(fd)
			if !ok {
				continue
			}

			// 读取消息
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Printf("error reading message: %v", err)
				connManager.Remove(fd)
				continue
			}

			log.Printf("received message: %s", string(msg))

			// 回显消息
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				log.Printf("error writing message: %v", err)
				connManager.Remove(fd)
			}
		}
	}
}
```


## 总结

`RawEpoll` 是一种直接使用 Linux `epoll` 的高性能 I/O 复用方式，适合需要极致性能和灵活性的场景。尽管实现难度较高，但它为高并发程序提供了强大的工具支持。在现代开发中，很多高级框架（如 Go 的 `netpoll`、Java 的 `NIO`）都基于 `epoll` 实现，但在特殊场景下，直接使用 RawEpoll 能获得更高的性能和更精准的控制。

