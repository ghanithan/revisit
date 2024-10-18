package revisit

import (
	"fmt"
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

func TestRemoveElement(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	want := []int{0, 1, 3, 0, 4, 0, 0, 0}
	got := RemoveElement(input, 2)
	fmt.Println(got, input)
	if !reflect.DeepEqual(input, want) {
		t.Errorf("Remove() = %q, want %q", input, want)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	want := []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0}
	got := RemoveDuplicates(input)
	fmt.Println(got, input)
	if !reflect.DeepEqual(input, want) {
		t.Errorf("RemoveDuplicate() = %q, want %q", input, want)
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	want := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 0}
	got := RemoveDuplicates2(input)
	fmt.Println(got, input)
	if !reflect.DeepEqual(input, want) {
		t.Errorf("RemoveDuplicate2() = %q, want %q", input, want)
	}
}

func TestMajorityElement(t *testing.T) {
	input := []int{1, 0, 1, 3, 1, 1, 2, 3, 1, 1}
	want := 1
	got := MajorityElement(input)
	fmt.Println("majority", got, input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MajorityElement() = %q, want %q", got, want)
	}
}

func TestRotate(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := []int{6, 7, 8, 1, 2, 3, 4, 5}
	Rotate(input, 3)
	fmt.Println("output", input)
	if !reflect.DeepEqual(input, want) {
		t.Errorf("Rotate() = %q, want %q", input, want)
	}
}

func TestMaximumSwap(t *testing.T) {
	input := 98368
	want := 98863
	got := MaximumSwap(input)
	fmt.Println("output", got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MaximumSwap() = %d, want %d", got, want)
	}
}
