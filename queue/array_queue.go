/**
 * @Author: cyj19
 * @Date: 2022/4/15 13:44
 */

package queue

import (
	"fmt"
	"sync"
)

/**
队列：先进先出，先进入队列的数据最新出来
下面使用数组实现队列
*/

// ArrayQueue 数组队列
type ArrayQueue struct {
	array []string   // 底层切片
	size  int        // 队列元素数量
	mu    sync.Mutex // 互斥锁保证并发安全
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

// Add 元素入队
func (q *ArrayQueue) Add(v string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.array = append(q.array, v)
	q.size = q.size + 1
}

// Remove 元素出队
func (q *ArrayQueue) Remove() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.size == 0 {
		panic("queue is empty")
	}
	v := q.array[0]
	// 创建新切片，空间占用不会越来越大
	newArray := make([]string, q.size-1, q.size-1)
	// 原切片的数据复制到新切片
	for i := 0; i < q.size-1; i++ {
		newArray[i] = q.array[i+1]
	}
	q.array = newArray
	q.size = q.size - 1
	return v
}

// Print 辅助打印
func (q *ArrayQueue) Print() string {
	result := "["
	for index := range q.array {
		if index == 0 {
			result = fmt.Sprintf("%s%s", result, q.array[index])
			continue
		}
		result = fmt.Sprintf("%s %s", result, q.array[index])
	}
	result += "]"
	return result
}
