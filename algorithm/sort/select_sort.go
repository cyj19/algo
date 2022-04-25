/**
 * @Author: cyj19
 * @Date: 2022/4/21 10:57
 */

package sort

// SelectSort 选择排序（遍历数组，每次找到最小值放到前面）
func SelectSort(list []int) {
	length := len(list)

	for i := 0; i < length-1; i++ {
		min := list[i]
		minIndex := i
		// 每轮迭代找出最小值
		for j := i + 1; j < length; j++ {
			// 更新最小值和最小值下标
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		// 这一轮找到的最小数的下标不等于最开始的下标，交换元素
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

// SelectGoodSort 优化后的选择排序，每轮同时找最大值和最小值
func SelectGoodSort(list []int) {
	length := len(list)

	// 每次都会找最大值和最小值，所以迭代轮数减半
	for i := 0; i < length/2; i++ {
		minIndex := i
		maxIndex := i
		for j := i + 1; j < length-i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
				// 一开始最大最小值是同一个，如果j位置的值大于最大值，意味着j位置的值也大于最小值
				continue
			}

			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		if maxIndex == i && minIndex != length-i-1 {
			// 如果最大值是开头的元素，最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			list[maxIndex], list[length-i-1] = list[length-i-1], list[maxIndex]
			// 再将开头元素和最小值交换
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == length-i-1 {
			// 如果最大值是开头的元素，最小值是最尾的元素
			// 头尾元素直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 如果最大值不是开头元素，把最大值放在尾部，最小值放在头部
			list[i], list[minIndex] = list[minIndex], list[i]
			list[length-i-1], list[maxIndex] = list[maxIndex], list[length-i-1]
		}
	}
}
