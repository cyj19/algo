/**
 * @Author: cyj19
 * @Date: 2022/4/15 10:45
 */

package stack

import (
	"fmt"
	"sync"
)

/**
下面使用单链表实现栈
*/

// LinkNode 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

// LinkStack 链表栈
type LinkStack struct {
	root *LinkNode  // 底层链表根节点
	size int        // 栈元素数量
	mu   sync.Mutex // 互斥锁保证并发安全
}

func (s *LinkStack) Size() int {
	return s.size
}

func (s *LinkStack) IsEmpty() bool {
	return s.size == 0
}

// Push 元素入栈
func (s *LinkStack) Push(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 栈顶为空，新增节点
	if s.root == nil {
		s.root = new(LinkNode)
		s.root.Value = v
	} else { // 新元素插入到链表头部
		// 原链表头部节点
		preNode := s.root
		// 创建新节点
		newNode := new(LinkNode)
		newNode.Value = v
		// 原来链表链接到新节点后面
		newNode.Next = preNode
		// 新节点放在头部
		s.root = newNode
	}
	s.size = s.size + 1

}

// Pop 元素出栈
func (s *LinkStack) Pop() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 栈判空
	if s.IsEmpty() {
		panic("stack is empty")
	}
	// 头部元素
	topNode := s.root
	v := topNode.Value
	// 取出头部元素
	s.root = topNode.Next
	// 元素数量-1
	s.size = s.size - 1

	return v
}

// Peek 获取栈顶元素，但不出栈
func (s *LinkStack) Peek() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.IsEmpty() {
		panic("stack is empty")
	}

	return s.root.Value
}

func (s *LinkStack) Print() string {
	result := "["
	p := s.root
	for i := 0; i < s.Size(); i++ {
		if i == 0 {
			result = fmt.Sprintf("%s%s", result, p.Value)
			p = p.Next
			continue
		}

		result = fmt.Sprintf("%s %s", result, p.Value)
		p = p.Next

	}
	result += "]"
	return result
}
