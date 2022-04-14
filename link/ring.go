/**
 * @Author: cyj19
 * @Date: 2022/4/14 8:55
 */

package main

/**
链表：一个数据连接另一个数据即为链表
链表类型：单链表、双链表、循环链表
下面实现最为复杂的循环双链表
*/

type Ring struct {
	prev, next *Ring
	Value      interface{}
}

// 初始化循环双链表
func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

// New 创建n个节点的循环双链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	// 已创建节点r，所以只需要再创建n-1个节点
	for ; n > 1; n-- {
		// p的next指向新节点，新节点的prev指向p
		p.next = &Ring{prev: p}
		// 更换当前节点为新节点
		p = p.next
	}
	// 当前节点的next指向r，形成后驱闭环
	p.next = r
	// r的prev指向当前节点，形成前驱闭环
	r.prev = p
	return r
}

// Next 获取r的后驱节点
func (r *Ring) Next() *Ring {
	// 不是双链表，进行初始化
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev 获取r的前驱节点
func (r *Ring) Prev() *Ring {
	// 不是双链表，进行初始化
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

// Move 获取第n个节点，n < 0 往前遍历，n > 0 往后遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	p := r
	switch {
	case n < 0:
		for ; n < 0; n++ {
			p = p.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			p = p.next
		}
	}
	return p
}

// Link 往r后链接节点s，返回r的原后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n

	}
	return n
}

// Unlink 往r后面删除n个节点，返回删除的新循环链表
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	// r链接第n+1个节点，相当于删除r到n+1之间的n个节点
	p := r.Move(n + 1)
	return r.Link(p)
}

func (r *Ring) Len() int {
	count := 0
	if r != nil {
		count = 1
		for p := r.Next(); p != r; p = p.next {
			count++
		}
	}
	return count
}

func main() {

}
