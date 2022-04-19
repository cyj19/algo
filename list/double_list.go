/**
 * @Author: cyj19
 * @Date: 2022/4/19 14:11
 */

package list

import "sync"

/**
列表List：可以存放数据的数据结构，数据按顺序排列，依次入队和出队
之前实现的队列和栈就是列表
双端列表（双端队列）：头部和尾部都可以入队出队
下面参照Redis的双端队列实现
*/

// Node 列表节点
type Node struct {
	pre, next *Node
	value     interface{}
}

// GetValue 获取节点值
func (n *Node) GetValue() interface{} {
	return n.value
}

// GetPre 获取前驱节点
func (n *Node) GetPre() *Node {
	return n.pre
}

// GetNext 获取后驱节点
func (n *Node) GetNext() *Node {
	return n.next
}

// HasPre 判断是否存在前驱节点
func (n *Node) HasPre() bool {
	return n.pre != nil
}

// HasNext 判断是否存在后驱节点
func (n *Node) HasNext() bool {
	return n.next != nil
}

// IsNil 是否为空节点
func (n *Node) IsNil() bool {
	return n == nil
}

// DoubleList 双端列表
type DoubleList struct {
	head *Node      // 链表头部节点
	tail *Node      // 链表尾部节点
	len  int        // 列表长度
	mu   sync.Mutex // 互斥锁保证列表的pop弹出并发安全
}

// Len 获取列表长度
func (l *DoubleList) Len() int {
	return l.len
}

// AddNodeFromHead 从头部开始往后，在第n+1节点前插入新节点，比如n=0,表示把新节点添加到第一个节点之前，成为新的头部节点
func (l *DoubleList) AddNodeFromHead(n int, v interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 判断是否越界
	if n != 0 && n >= l.len {
		panic("index out")
	}

	// 获取头部节点
	node := l.head

	// 遍历找出第n+1个节点（因为已经先把头部节点取出来了，所以i=1开始遍历n个，获取到的就是n+1）
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 创建新节点
	newNode := new(Node)
	newNode.value = v

	// 如果定位到的节点为空，说明n <= 0
	if node.IsNil() {
		// 新节点作为头部和尾部
		l.head = newNode
		l.tail = newNode
	} else {
		pre := node.pre
		// 如果定位到的节点的前驱节点为nil，说明该节点是头部节点
		if pre.IsNil() {
			node.pre = newNode
			newNode.next = node

			// 新节点成为头部节点
			l.head = newNode
		} else { // 不是头部节点

			// pre-->newNode-->node
			newNode.next = node
			newNode.pre = pre

			pre.next = newNode
			node.pre = newNode
		}
	}

	// 列表长度+1
	l.len = l.len + 1
}

// AddNodeFromTail 从尾部开始往前，在第n+1节点前插入新节点，如n=0，表示在第一个节点之前插入新节点，成为新的尾部节点
func (l *DoubleList) AddNodeFromTail(n int, v interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 判断是否越界
	if n != 0 && n >= l.len {
		panic("index out")
	}

	// 获取尾部节点
	node := l.tail

	// 定位为第n+1节点
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	// 创建新节点
	newNode := new(Node)
	newNode.value = v

	// 如果定位的节点为nil，说明n<=0，列表为空
	if node.IsNil() {
		l.head = newNode
		l.tail = newNode
	} else {
		// 定位节点的后驱节点
		next := node.next

		// 定位节点的后驱节点为空，说明定位节点是尾部节点，n=0
		if next.IsNil() {
			newNode.pre = node
			node.next = newNode
			l.tail = newNode
		} else { // 不是尾部节点

			// node-->newNode-->next
			newNode.pre = node
			newNode.next = next

			node.next = newNode
			next.pre = newNode
		}
	}

	// 列表长度+1
	l.len = l.len + 1

}

// IndexFromHead 从头部开始往后遍历，获取第n+1个节点，索引从0开始
func (l *DoubleList) IndexFromHead(n int) *Node {
	// 索引等于或大于列表长度，不会有节点
	if n >= l.len {
		return nil
	}

	node := l.head
	// 已取出头部节点，i从1开始
	for i := 1; i <= n; i++ {
		node = node.next
	}

	return node
}

// IndexFromTail 从尾部开始往前遍历，获取第n+1个节点，索引从0开始
func (l *DoubleList) IndexFromTail(n int) *Node {
	// 索引等于或大于列表长度，不会有节点
	if n >= l.len {
		return nil
	}

	node := l.tail
	// 已取出尾部节点，i从1开始
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	return node
}

// 从列表移除节点
func (l *DoubleList) move(node *Node) {
	pre := node.pre
	next := node.next

	// 如果定位节点的前驱后驱都是空，说明定位节点是列表的唯一节点
	if pre.IsNil() && next.IsNil() {
		l.head = nil
		l.tail = nil
	} else if pre.IsNil() { // 前驱节点为空，说明是头部节点
		next.pre = nil
		l.head = next
	} else if next.IsNil() { // 后驱节点为空，说明是尾部节点
		pre.next = nil
		l.tail = pre
	} else { // 移除的是中间节点
		pre.next = next
		next.pre = pre
	}
}

// PopFromHead 从头部开始往后移除第n+1个节点并返回，n从0开始
func (l *DoubleList) PopFromHead(n int) *Node {
	l.mu.Lock()
	defer l.mu.Unlock()

	if n >= l.len {
		return nil
	}

	node := l.head

	// 从列表中定位节点
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 移除节点
	l.move(node)

	// 列表长度-1
	l.len = l.len - 1
	return node
}

// PopFromTail 从尾部开始往前移除第n+1个节点，索引从0开始
func (l *DoubleList) PopFromTail(n int) *Node {
	l.mu.Lock()
	defer l.mu.Unlock()

	if n >= l.len {
		return nil
	}

	node := l.tail

	// 从列表中定位节点
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	// 移除节点
	l.move(node)

	l.len = l.len - 1
	return node
}
