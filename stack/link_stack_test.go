/**
 * @Author: cyj19
 * @Date: 2022/4/15 11:16
 */

package stack

import (
	"log"
	"testing"
)

func TestLinkStack_Push(t *testing.T) {
	stack := new(LinkStack)
	stack.Push("1")
	log.Printf("size=%d \n", stack.Size())
	stack.Push("2")
	log.Printf("size=%d \n", stack.Size())
	stack.Push("3")
	log.Printf("size=%d \n", stack.Size())
}

func TestLinkStack_Pop(t *testing.T) {
	stack := new(LinkStack)
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	v := stack.Pop()
	log.Printf("size=%d element=%s \n", stack.Size(), v)
	v = stack.Pop()
	log.Printf("size=%d element=%s \n", stack.Size(), v)
	v = stack.Pop()
	log.Printf("size=%d element=%s \n", stack.Size(), v)
}

func TestLinkStack_Peek(t *testing.T) {
	stack := new(LinkStack)
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	v := stack.Peek()
	log.Printf("size=%d element=%s \n", stack.Size(), v)
}

func TestLinkStack_Print(t *testing.T) {
	stack := new(LinkStack)
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	log.Println(stack.Print())
}
