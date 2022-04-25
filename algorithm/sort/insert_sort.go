/**
 * @Author: cyj19
 * @Date: 2022/4/25 14:32
 */

package sort

// InsertSort 直接插入排序：每一轮取一个元素插入到该元素左边的有序列表形成新的有序列表
func InsertSort(list []int) {
	length := len(list)

	// 一开始有序的元素只有list[0]，所以待排序下标从1开始
	for i := 1; i < length; i++ {
		// 待排序元素
		deal := list[i]
		// 待排序元素左边的元素下标（有序列表的从右往左开始和待排序元素比较）
		j := i - 1
		// 寻找待排序元素的位置
		for ; j >= 0 && deal < list[j]; j-- {
			// 比待排序元素小的元素往后移
			list[j+1] = list[j]
		}
		// 此时deal > list[j]，那么待排序元素的下标就是j+1
		list[j+1] = deal
	}

}
