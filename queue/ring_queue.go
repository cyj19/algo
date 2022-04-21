/**
 * @Author: cyj19
 * @Date: 2022/4/15 14:57
 */

package queue

import (
	"fmt"
	"sync"
)

/**
环形队列：收尾相连的队列
下面使用数组实现环形队列
*/

type RingQueue struct {
	queue *ringQueue
}

// ringQueue 环形队列
type ringQueue struct {
	array []interface{} // 底层切片
	count int           // 队列中的实际元素数量
	size  int           // 队列大小
	top   int           // 队列顶部下标
	mu    sync.Mutex    // 互斥锁保证并发安全
}

func NewRingQueue(size int) *RingQueue {
	if size < 1 {
		panic("queue size must large than 0")
	}
	queue := new(ringQueue)
	queue.array = make([]interface{}, size, size)
	queue.size = size
	return &RingQueue{queue: queue}
}

// Size 获取队列大小
func (q *RingQueue) Size() int {
	return q.queue.size
}

// Count 获取队列中的实际元素数量
func (q *RingQueue) Count() int {
	return q.queue.count
}

// IsEmpty 队列判空
func (q *RingQueue) IsEmpty() bool {
	return q.queue.count == 0
}

// IsFull 队列判满
func (q *RingQueue) IsFull() bool {
	return q.queue.count == q.queue.size
}

// Add 元素入队
func (q *RingQueue) Add(v interface{}) {
	q.queue.mu.Lock()
	defer q.queue.mu.Unlock()
	// 判满
	if q.queue.count == q.queue.size {
		panic("ring queue is full")
	}
	top := q.queue.top
	count := q.queue.count
	size := q.queue.size
	// 从顶部下标开始计算队列最后元素的下一个下标
	index := (count + top) % size
	q.queue.array[index] = v
	q.queue.count = count + 1

}

// Remove 元素出队
func (q *RingQueue) Remove() interface{} {
	q.queue.mu.Lock()
	defer q.queue.mu.Unlock()
	// 判空
	if q.queue.count == 0 {
		panic("ring queue is empty")
	}
	top := q.queue.top
	v := q.queue.array[top]
	// 元素数量-1
	q.queue.count -= 1
	// 顶部下标后移
	q.queue.top = (top + 1) % q.queue.size
	return v
}

// Print 辅助打印
func (q *RingQueue) Print() string {
	result := "["
	top := q.queue.top
	for i := 0; i < q.queue.count; i++ {
		j := (i + top) % q.queue.size
		if i == 0 {
			result = fmt.Sprintf("%s%s", result, q.queue.array[j])
			continue
		}
		result = fmt.Sprintf("%s %s", result, q.queue.array[j])
	}
	result += "]"
	return result
}
