/**
 * @Author: cyj19
 * @Date: 2022/4/25 15:09
 */

package sort

// ShellSort 希尔排序：直接插入排序的改进版，通过多次分组进行插入排序
func ShellSort(list []int) {
	n := len(list)

	// 增量每次减半
	for step := n / 2; step >= 1; step /= 2 {
		// 将step的整倍数分为一组进行插入排序
		for i := step; i < n; i += step {
			// 待排序的数
			deal := list[i]
			// 待排序的数的左边最近一个数的下标
			j := i - step
			// 在有序列表中寻找待排序数的插入位置
			for ; j >= 0 && deal < list[j]; j -= step {
				// 比待排序数大的数往后移
				list[j+step] = list[j]
			}
			// 此时deal > list[j]，那么待排序数的下标就是j+step
			list[j+step] = deal
		}
	}
}
