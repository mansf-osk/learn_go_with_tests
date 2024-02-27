package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d and want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("add slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 9})
		want := []int{3, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v and want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v and want %v", got, want)
		}
	}

	t.Run("add tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{}, []int{3, 9})
	}
}
