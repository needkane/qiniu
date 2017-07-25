package bubblesort

import (
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 7, 1}
	BubbleSort(values)
	if values[0] != 1 || values[4] != 7 {
		t.Error("Failed Bubblesort")
	}
}
func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 7, 1}
	BubbleSort(values)
	if values[2] != 5 || values[3] != 5 {
		t.Error("Failed Bubblesort")
	}
}
func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("Failed Bubblesort")
	}
}
