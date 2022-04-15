/**
 * @Author: cyj19
 * @Date: 2022/4/15 15:44
 */

package queue

import (
	"log"
	"testing"
)

func TestRingQueue_Add(t *testing.T) {

	queue := NewRingQueue(5)
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")
	queue.Add("4")
	queue.Add("5")
	log.Println(queue.Print())
	//queue.Add("6")
}

func TestRingQueue_Remove(t *testing.T) {
	queue := NewRingQueue(5)
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")
	queue.Add("4")
	queue.Add("5")
	log.Println(queue.Print(), queue.Count())
	log.Println(queue.queue.array)

	log.Println(queue.Remove())
	log.Println(queue.Remove())
	log.Printf("count=%d \n", queue.Count())

	queue.Add("6")
	queue.Add("7")
	log.Println(queue.Print(), queue.Count())
	log.Println(queue.queue.array)

	log.Println(queue.Remove())
	log.Println(queue.Print(), queue.Count())
	log.Println(queue.queue.array)
}
