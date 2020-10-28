package arrays

import (
	"reflect"
	"testing"
)

func checkSum(t *testing.T, expected, got []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

// TestSum test of sum
func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expected := 15
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 4}, []int{3, 5})
	expected := []int{7, 8}
	checkSum(t, expected, got)
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 4}, []int{3, 5})
	expected := []int{6, 5}
	checkSum(t, expected, got)
}
