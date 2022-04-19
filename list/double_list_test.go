/**
 * @Author: cyj19
 * @Date: 2022/4/19 15:42
 */

package list

import (
	"testing"
)

func TestDoubleList_AddNodeFromHead(t *testing.T) {
	l := new(DoubleList)

	// 在列表头部插入新元素
	l.AddNodeFromHead(0, "a")
	l.AddNodeFromHead(0, "b")
	l.AddNodeFromHead(0, "c")
	// 在列表尾部插入新元素
	l.AddNodeFromTail(0, "d")
	l.AddNodeFromTail(0, "e")

	l.AddNodeFromTail(l.Len()-1, "f")
	l.AddNodeFromHead(l.Len()-1, "g")

	node := l.head
	for !node.IsNil() {
		t.Log(node.value)
		node = node.next
	}
}

func TestDoubleList_IndexFromHead(t *testing.T) {

	l := new(DoubleList)

	// 在列表头部插入新元素
	l.AddNodeFromHead(0, "a")
	l.AddNodeFromHead(0, "b")
	l.AddNodeFromHead(0, "c")
	// 在列表尾部插入新元素
	l.AddNodeFromTail(0, "d")
	l.AddNodeFromTail(0, "e")

	node := l.head
	for !node.IsNil() {
		t.Log(node.value)
		node = node.next
	}

	t.Log(l.IndexFromHead(2))
	t.Log(l.IndexFromTail(3))

}

func TestDoubleList_PopFromHead(t *testing.T) {
	l := new(DoubleList)

	// 在列表头部插入新元素
	l.AddNodeFromHead(0, "a")
	l.AddNodeFromHead(0, "b")
	l.AddNodeFromHead(0, "c")
	// 在列表尾部插入新元素
	l.AddNodeFromTail(0, "d")
	l.AddNodeFromTail(0, "e")

	node := l.head
	for !node.IsNil() {
		t.Log(node.value)
		node = node.next
	}

	t.Log(l.PopFromHead(2))
	node = l.head
	for !node.IsNil() {
		t.Log(node.value)
		node = node.next
	}

	t.Log(l.PopFromTail(3))
	node = l.head
	for !node.IsNil() {
		t.Log(node.value)
		node = node.next
	}
}
