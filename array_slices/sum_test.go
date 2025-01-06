package main

import (
	"reflect"
	"testing"
)

// func TestSum(t *testing.T) {
// 	numbers := [5]int{1, 2, 3, 4, 5}

// 	got := Sum(numbers)
// 	expected := 15

// 	if got != expected {
// 		t.Errorf("got %d want %d", got, expected)
// 	}
// }

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 5, 6}

		got := Sum(numbers)
		want := 17

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}
func TestSumAllTails(t *testing.T) {
	checkSum := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v but expected %v", got, expected)
		}
	}
	t.Run("Make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		expected := []int{5, 11}
		checkSum(t, got, expected)
	})

	t.Run("safe sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 5, 7})
		expected := []int{0, 12}
		checkSum(t, got, expected)
	})

}
