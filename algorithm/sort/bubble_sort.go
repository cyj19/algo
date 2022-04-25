/**
 * @Author: cyj19
 * @Date: 2022/4/21 10:14
 */

package sort

// BubbleSort 冒泡排序（相邻两个元素依次比较大小交换，第一和第二比较，第二和第三比较，依次类推）
func BubbleSort(list []int) {
	length := len(list)
	// 在一轮中有没有交换过
	didSwap := false
	// 进行length-1轮比较
	for i := length - 1; i > 0; i-- {
		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}

		// 如果在一轮迭代中没有交换，说明现在的list已经是有序的
		if !didSwap {
			return
		}
	}

}
