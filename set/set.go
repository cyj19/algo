/**
 * @Author: cyj19
 * @Date: 2022/4/20 10:28
 */

package set

import (
	"sync"
)

/**
集合Set：存放的数据不可重复，可以没有顺序关系，也可以按值排序，是一种特殊的列表
字典map的键是不可重复的，使用字典实现Set
*/

// Set 集合
type Set struct {
	m            map[interface{}]struct{} // 字典的存放
	len          int                      // 集合大小
	sync.RWMutex                          // 读写锁，保证并发安全 ps:因为后续需要把集合转为列表，需要使用读锁禁止写入，但不禁止并发读
}

// NewSet 创建集合，cap初始容量
func NewSet(cap int) *Set {
	return &Set{
		m: make(map[interface{}]struct{}, cap),
	}
}

func (s *Set) Len() int {
	return s.len
}

// Add 添加元素
func (s *Set) Add(v interface{}) {
	// 加写锁
	s.Lock()
	defer s.Unlock()

	s.m[v] = struct{}{}
	// 重新计算大小
	s.len = len(s.m)
}

// Remove 删除元素
func (s *Set) Remove(v interface{}) {
	// 加写锁
	s.Lock()
	defer s.Unlock()

	// 集合为空，直接返回
	if s.len == 0 {
		return
	}

	delete(s.m, v)
	// 重新计算大小
	s.len = len(s.m)
}

// Has 查找元素是否在集合中
func (s *Set) Has(v interface{}) bool {
	// 加读锁
	s.RLock()
	defer s.RUnlock()

	_, ok := s.m[v]
	return ok
}

// IsEmpty 集合是否为空
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// Clear 清空集合
func (s *Set) Clear() {
	// 加写锁
	s.Lock()
	defer s.Unlock()
	// 字典重新赋值
	s.m = map[interface{}]struct{}{}
	s.len = 0
}

// List 集合转为列表
func (s *Set) List() []interface{} {
	// 加读锁
	s.RLock()
	defer s.RUnlock()

	list := make([]interface{}, 0, s.Len())
	for item := range s.m {
		list = append(list, item)
	}

	return list
}
