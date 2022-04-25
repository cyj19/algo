/**
 * @Author: cyj19
 * @Date: 2022/4/21 10:21
 */

package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(arr)
	t.Log(arr)
}

func TestSelectSort(t *testing.T) {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(arr)
	t.Log(arr)
}

func TestSelectGoodSort(t *testing.T) {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectGoodSort(arr)
	t.Log(arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(arr)
	t.Log(arr)
}

func TestShellSort(t *testing.T) {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	ShellSort(arr)
	t.Log(arr)
}
