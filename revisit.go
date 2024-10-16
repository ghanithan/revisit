package revisit

import (
	"fmt"
)

func Revisit(name string) string {
	greetings := fmt.Sprintf("hello %s", name)
	fmt.Println(greetings)
	return greetings
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	fmt.Println("input", nums1)
	quickSort(nums1)
	fmt.Println("output", nums1)
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
