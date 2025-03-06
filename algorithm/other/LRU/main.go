package main

import "fmt"

// Node 结构，表示双向链表节点
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache 结构
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // 头部是最近使用的
	tail     *Node // 尾部是最久未使用的
}

// Constructor 初始化 LRUCache
func Constructor(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{}, // 哨兵头
		tail:     &Node{}, // 哨兵尾
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get 获取值，并将节点移动到头部
func (l *LRUCache) Get(key int) int {
	if node, exists := l.cache[key]; exists {
		l.moveToHead(node)
		return node.value
	}
	return -1
}

// Put 插入新键值对，如果容量满了，则淘汰最近最少使用的节点
func (l *LRUCache) Put(key int, value int) {
	if node, exists := l.cache[key]; exists {
		// 更新值并移动到头部
		node.value = value
		l.moveToHead(node)
		return
	}

	// 创建新节点
	newNode := &Node{key: key, value: value}
	l.cache[key] = newNode
	l.addToHead(newNode)

	// 超过容量，移除尾部节点
	if len(l.cache) > l.capacity {
		tail := l.removeTail()
		delete(l.cache, tail.key)
	}
}

// moveToHead 将节点移动到链表头部
func (l *LRUCache) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

// addToHead 将节点插入到头部
func (l *LRUCache) addToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

// removeNode 移除双向链表中的某个节点
func (l *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeTail 移除链表尾部节点，并返回它
func (l *LRUCache) removeTail() *Node {
	tail := l.tail.prev
	l.removeNode(tail)
	return tail
}

// 测试代码
func main() {
	lru := Constructor(2)
	lru.Put(1, 10)
	lru.Put(2, 20)
	fmt.Println(lru.Get(1)) // 输出: 10
	lru.Put(3, 30)          // 淘汰 key=2
	fmt.Println(lru.Get(2)) // 输出: -1
	lru.Put(4, 40)          // 淘汰 key=1
	fmt.Println(lru.Get(1)) // 输出: -1
	fmt.Println(lru.Get(3)) // 输出: 30
	fmt.Println(lru.Get(4)) // 输出: 40
}
