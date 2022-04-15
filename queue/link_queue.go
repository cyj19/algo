/**
 * @Author: cyj19
 * @Date: 2022/4/15 14:09
 */

package queue

import (
	"fmt"
	"sync"
)

/**
下面使用链表实现队列
*/

// LinkNode 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

// LinkQueue 链表队列
type LinkQueue struct {
	root *LinkNode
	size int
	mu   sync.Mutex
}

func (q *LinkQueue) Size() int {
	return q.size
}

func (q *LinkQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkQueue) Add(v string) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		q.root = new(LinkNode)
		q.root.Value = v
	} else {
		// 创建新节点
		newNode := new(LinkNode)
		newNode.Value = v
		// 遍历找到原链表最后一个节点
		p := q.root
		for p.Next != nil {
			p = p.Next
		}
		p.Next = newNode
	}
	q.size = q.size + 1
}

func (q *LinkQueue) Remove() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.size == 0 {
		panic("queue is empty")
	}
	// 头部节点
	topNode := q.root
	v := topNode.Value
	// 头部节点更改为原头部节点的下一个节点
	q.root = topNode.Next
	q.size = q.size - 1
	return v
}

func (q *LinkQueue) Print() string {
	result := "["
	p := q.root
	for i := 0; i < q.size; i++ {
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
