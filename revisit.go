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
	for i := 0; i < m+n; i++ {
		for j := i + 1; j < m+n; j++ {
			if nums1[i] > nums1[j] {
				swap := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = swap
			}
		}
	}

}
