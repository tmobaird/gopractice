package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d wanted %d given, %v", got, want, numbers)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	fmt.Println(got)
	// Output: 15
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)

	}
}

func ExampleSumAll() {
	got := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(got)
	// Output: [3 9]
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v wanted %v", got, want)
		}
	}

	t.Run("makes the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
	fmt.Println(got)
	// Output: [5 9]
}
