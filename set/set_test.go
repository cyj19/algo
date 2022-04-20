/**
 * @Author: cyj19
 * @Date: 2022/4/20 11:08
 */

package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	// 初始化一个容量为5的不可重复集合
	s := NewSet(5)

	s.Add(1)
	s.Add(1)
	s.Add(2)
	t.Log("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		t.Log("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		t.Log("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	t.Log("list of all items", s.List())
}
