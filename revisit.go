package revisit

import (
	"fmt"
)

func Revisit(name string) string {
	greetings := fmt.Sprintf("hello %s", name)
	fmt.Println(greetings)
	return greetings
}

// https://leetcode.com/problems/rotate-array/submissions/1425243957

func Rotate(nums []int, k int) {
	temp := make([]int, k)
	if len(nums) > k {
		copy(temp, nums[len(nums)-k:])
		for i := len(nums) - 1 - k; i >= 0; i-- {
			nums[i+k] = nums[i]
		}
		for i := 0; i < len(temp); i++ {
			nums[i] = temp[i]
		}
	} else {
		Rotate1(nums, k)
	}

}

func Rotate1(nums []int, k int) {
	if k <= 0 {
		return
	}
	temp := nums[len(nums)-1]
	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = temp
	Rotate1(nums, k-1)

}

func MajorityElement(nums []int) int {
	majority := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		majority[nums[i]]++
	}
	for k, v := range majority {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

func RemoveDuplicates2(nums []int) int {
	count := 0
	threshold := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			threshold++
			if threshold > 1 {
				nums = append(append(nums[:i], nums[i+1:]...), 0)
				count++
			}
		} else {
			threshold = 0
		}
	}
	return len(nums) - count
}

func RemoveDuplicates(nums []int) int {
	count := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(append(nums[:i], nums[i+1:]...), 0)
			count++
		}
	}
	return len(nums) - count
}

func RemoveElement(nums []int, val int) int {
	count := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			count++
			nums = append(append(nums[:i], nums[i+1:]...), 0)
		}
	}
	return len(nums) - count
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	//fmt.Println("input", nums1)
	quickSort(nums1)
	//fmt.Println("output", nums1)
}

func swap(array []int, first, second int) {
	temp := array[first]
	array[first] = array[second]
	array[second] = temp

}

func quickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	pivot := len(array) - 1
	i := -1
	j := 0
	for ; j <= pivot; j++ {
		if array[j] <= array[pivot] {
			i++
			swap(array, i, j)
			if j == pivot {
				pivot = i
			}
		}
	}
	if i < 0 {
		quickSort(array[:len(array)-2])
	} else {
		quickSort(array[0:pivot])
		quickSort(array[pivot:])
	}
}
