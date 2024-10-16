package revisit

import (
	"reflect"
	"testing"
)

func TestRevisit(t *testing.T) {
	want := "hello world"
	if got := Revisit("world"); got != want {
		t.Errorf("revisit() = %q, want %q", got, want)
	}
}

func TestMerge(t *testing.T) {
	want := []int{2, 3, 5, 6, 10, 12}
	nums1 := []int{3, 10, 12, 0, 0, 0}
	m := 3
	nums2 := []int{5, 2, 6}
	n := 3
	Merge(nums1, m, nums2, n)
	if got := nums1; !reflect.DeepEqual(got, want) {
		t.Errorf("Merge() = %q, want %q", got, want)
	}
}
