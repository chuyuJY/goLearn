package _3_timer

import (
	"math/rand"
	"time"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func fastSortArray(nums []int) {
	fastSort(nums, 0, len(nums)-1)
}

func fastSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(nums, start, end)
	fastSort(nums, start, pivot-1)
	fastSort(nums, pivot+1, end)
}

func partition(nums []int, start, end int) int {
	random := start + rand.Intn(end-start+1)
	nums[random], nums[end] = nums[end], nums[random]
	index := end
	for start < end {
		for start < end && nums[start] < nums[index] {
			start++
		}
		for start < end && nums[end] >= nums[index] {
			end--
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	nums[end], nums[index] = nums[index], nums[end]
	return end
}

func mergeSortArray(nums []int) {
	dst := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dst[i] = nums[i]
	}
	mergeSort(nums, dst, 0, len(nums)-1)
}

func mergeSort(src, dst []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	mergeSort(dst, src, start, mid)
	mergeSort(dst, src, mid+1, end)
	cur := start
	left, right := start, mid+1
	for left <= mid || right <= end {
		if left > mid || (right <= end && src[left] > src[right]) {
			dst[cur] = src[right]
			right++
		} else {
			dst[cur] = src[left]
			left++
		}
		cur++
	}
}
