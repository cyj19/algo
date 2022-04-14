/**
 * @Author: cyj19
 * @Date: 2022/4/14 9:41
 */

package link

import (
	"log"
	"testing"
)

const n = 5

var r *Ring

func TestRing(t *testing.T) {
	r = New(n)
	p := r
	i := 0
	for {
		p.Value = i
		p = p.next
		i++
		if p == r {
			break
		}
	}

	t.Run("TestRing_New", TestRing_New)
	t.Run("TestRing_PrevAndNext", TestRing_PrevAndNext)
	t.Run("TestRing_Move", TestRing_Move)
	t.Run("TestRing_Link", TestRing_Link)
	t.Run("TestRing_Unlink", TestRing_Unlink)
	t.Run("TestRing_Len", TestRing_Len)
}

func TestRing_New(t *testing.T) {
	p := r
	for {
		log.Println(p.Value)
		p = p.next
		if p == r {
			break
		}
	}
}

func TestRing_PrevAndNext(t *testing.T) {
	log.Println(r.Prev())
	log.Println(r.Next())
}

func TestRing_Move(t *testing.T) {
	log.Println(r.Move(2))
	log.Println(r.Move(-2))
}

func TestRing_Link(t *testing.T) {
	p := &Ring{Value: 5}
	s := &Ring{Value: 6}
	k := &Ring{Value: 7}

	p.next = s
	p.prev = k
	s.prev = p
	s.next = k
	k.prev = s
	k.next = p

	r.Link(s)
	node := r
	for {
		log.Println(node.Value)
		node = node.next
		if node == r {
			break
		}
	}
}

func TestRing_Unlink(t *testing.T) {
	temp := r.Unlink(2)
	p := r
	for {
		log.Println(p.Value)
		p = p.next
		if p == r {
			break
		}
	}

	log.Println("-------------")
	node := temp
	for {
		log.Println(node.Value)
		node = node.next
		if node == temp {
			break
		}
	}
}

func TestRing_Len(t *testing.T) {
	log.Println(r.Len())
}
