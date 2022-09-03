package sort

func QuickSort(list []int, start, end int) {
	if start >= end {
		return
	}
	i := start
	j := end
	num := list[i]
	for i < j {
		for i < j && num <= list[j] {
			j--
		}

		for i < j && num >= list[i] {
			i++
		}
		// 交换
		list[i], list[j] = list[j], list[i]
	}

	list[start], list[i] = list[i], list[start]
	QuickSort(list, start, j-1)
	QuickSort(list, j+1, end)
}
