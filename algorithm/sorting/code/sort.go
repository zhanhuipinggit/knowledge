package code

type quickSort struct{}

func (q *quickSort) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[i] < nums[left] {
			i++
		}
		for i < j && nums[j] > nums[left] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func (q *quickSort) quickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)

}

func siftDown(nums *[]int, n, i int) {
	for true {
		l := 2*i + 1
		r := 2*i + 2
		m := i
		if l < n && (*nums)[l] > (*nums)[m] {
			m = l
		}
		if r < n && (*nums)[r] > (*nums)[m] {
			m = r

		}
		if m == i {
			break
		}

		(*nums)[i], (*nums)[m] = (*nums)[m], (*nums)[i]
		i = m
	}

}

func heapSort(nums *[]int) {
	for i := len(*nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, len(*nums), i)
	}

	for i := len(*nums) - 1; i >= 0; i-- {
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		siftDown(nums, i, 0)
	}

}

func bubbleSort(nums *[]int) {
	for i := len(*nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if (*nums)[j] > (*nums)[j+1] {
				(*nums)[j], (*nums)[j+1] = (*nums)[j+1], (*nums)[j]
			}
		}
	}
}

func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		k := i
		for j := i + 1; j < len(nums); j++ {
			if nums[k] > nums[j] {
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base

	}
}
