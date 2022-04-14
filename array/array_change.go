/**
 * @Author: cyj19
 * @Date: 2022/4/14 16:05
 */

package array

import (
	"fmt"
	"sync"
)

/**
可变数组：golang中内置实现了这种数据结构，slice
下面实现自定义的可变int数组(由于golang语言特性不能使用变量声明数组长度[n]int{}，所以此处使用slice代替[n]int{}，但不使用slice的append函数)
*/

type Array struct {
	array []int      // 固定长度数组
	len   int        // 数组长度
	cap   int        // 数组容量
	lock  sync.Mutex // 互斥锁
}

// Make 创建可变数组
func Make(len, cap int) *Array {
	if len > cap {
		panic("len large than cap")
	}
	array := make([]int, cap, cap)
	s := new(Array)
	s.array = array
	s.len = len
	s.cap = cap
	return s
}

func (a *Array) Len() int {
	return a.len
}

func (a *Array) Cap() int {
	return a.cap
}

// Append 添加元素
func (a *Array) Append(element int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		newCap := a.cap * 2
		if a.cap == 0 {
			newCap = 1
		}
		// 创建新数组
		array := make([]int, newCap, newCap)
		// 旧数组数据复制到新数组
		for k, v := range a.array {
			array[k] = v
		}
		a.array = array
		a.cap = newCap
	}
	// 添加元素
	a.array[a.len] = element
	// 数组长度+1
	a.len += 1
}

// AppendMany 添加多个元素
func (a *Array) AppendMany(elements ...int) {
	for _, v := range elements {
		a.Append(v)
	}
}

// Get 获取元素
func (a *Array) Get(index int) int {
	// 数组越界
	if a.len == 0 || index >= a.len {
		panic("index out of array range")
	}
	return a.array[index]
}

// Println 辅助打印
func (a *Array) Println() string {
	result := "["
	for i := 0; i < a.Len(); i++ {
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, a.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %d", result, a.Get(i))
	}
	result = fmt.Sprintf("%s]", result)
	return result
}
