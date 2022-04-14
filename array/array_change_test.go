/**
 * @Author: cyj19
 * @Date: 2022/4/14 16:24
 */

package main

import (
	"log"
	"testing"
)

func TestArray_Append(t *testing.T) {
	a := Make(0, 0)
	for i := 1; i < 10; i++ {
		a.Append(i)
		log.Printf("len=%d cap=%d \n", a.Len(), a.Cap())
	}

	log.Println(a)
}

func TestArray_AppendMany(t *testing.T) {
	a := Make(0, 0)
	a.AppendMany(1, 2, 3, 4, 5)
	log.Println(a)
}

func TestArray_Println(t *testing.T) {
	a := Make(0, 0)
	a.AppendMany(1, 2, 3, 4, 5)
	log.Println(a.Println())
}
