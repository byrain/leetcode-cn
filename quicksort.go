package main

import "fmt"

func partition(nums []int, left, right int) int {
	p := nums[left]
	i := left
	j := right
	for i < j {
		for nums[j] >= p && j > i {
			j--
		}

		for nums[i] <= p && i < j {
			i++
		}
		fmt.Println(i, j)
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

func main() {
	nums := []int{6, 5, 2, 7, 3, 9, 8, 4, 10, 1}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
