/**
 * @Author: cyj19
 * @Date: 2022/4/15 9:07
 */

package stack

import (
	"fmt"
	"sync"
)

/**
栈：先进后出，先进队的数据最后才出来
下面使用数组实现栈
*/

// ArrayStack 数组栈
type ArrayStack struct {
	array []interface{} // 底层切片
	size  int           // 栈的元素数量
	mu    sync.Mutex    // 互斥锁保证并发安全
}

// Size 获取栈元素数量
func (s *ArrayStack) Size() int {
	return s.size
}

// IsEmpty 判断栈空
func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}

// Push 元素入栈
func (s *ArrayStack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.array = append(s.array, v)
	s.size += 1
}

// Pop 元素出栈
func (s *ArrayStack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.size == 0 {
		panic("stack is empty")
	}

	v := s.array[s.size-1]
	// 创建新切片，空间占用不会越来越大，但移动次数可能会更多
	newArray := make([]interface{}, s.size-1, s.size-1)
	for i := 0; i < s.size-1; i++ {
		newArray[i] = s.array[i]
	}
	s.array = newArray
	s.size = s.size - 1
	return v
}

// Peek 获取栈顶元素，但是不出栈
func (s *ArrayStack) Peek() interface{} {
	if s.size == 0 {
		panic("stack is empty")
	}
	v := s.array[s.size-1]
	return v
}

func (s *ArrayStack) Print() string {
	result := "["
	for i := s.size - 1; i >= 0; i-- {
		if i == s.size-1 {
			result = fmt.Sprintf("%s%s", result, s.array[i])
			continue
		}
		result = fmt.Sprintf("%s %s", result, s.array[i])
	}
	result += "]"
	return result
}
