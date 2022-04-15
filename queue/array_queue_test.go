/**
 * @Author: cyj19
 * @Date: 2022/4/15 13:59
 */

package queue

import (
	"log"
	"testing"
)

func TestArrayQueue_Add(t *testing.T) {
	queue := new(ArrayQueue)
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")
	log.Println(queue.Size())
	log.Println(queue.Print())
}

func TestArrayQueue_Remove(t *testing.T) {
	queue := new(ArrayQueue)
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")

	v := queue.Remove()
	log.Printf("size=%d element=%s \n", queue.Size(), v)
	v = queue.Remove()
	log.Printf("size=%d element=%s \n", queue.Size(), v)
	v = queue.Remove()
	log.Printf("size=%d element=%s \n", queue.Size(), v)

	queue.Add("4")
	log.Println(queue.Print())
}
