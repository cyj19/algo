/**
 * @Author: cyj19
 * @Date: 2022/4/15 14:50
 */

package queue

import (
	"log"
	"testing"
)

func TestLinkQueue_Add(t *testing.T) {
	queue := new(LinkQueue)
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")
	log.Printf("size=%d queue=%s \n", queue.Size(), queue.Print())
}

func TestLinkQueue_Remove(t *testing.T) {
	queue := new(LinkQueue)
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")

	v := queue.Remove()
	log.Printf("size=%d element=%s \n", queue.Size(), v)
	v = queue.Remove()
	log.Printf("size=%d element=%s \n", queue.Size(), v)
	v = queue.Remove()
	log.Printf("size=%d element=%s \n", queue.Size(), v)
}
