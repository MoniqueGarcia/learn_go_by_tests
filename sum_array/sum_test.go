package main

import "testing"
import "reflect"


func TestSum(t *testing.T) {

	assertResult := func(test testing.TB, numbers []int, got, want int) {
		test.Helper()
		if got != want {
			test.Errorf("expected %d but got %d. Input: %v", want, got, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)
		expected := 6

		assertResult(t, numbers, result, expected)
	})

}

func TestSumAll(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{3, 2, 1}

	result := SumAll(arr1, arr2)
	expected := []int{6, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v. Input: %v and %v", expected, result, arr1, arr2)
	}
}

func TestSumAllTails(t *testing.T) {

	assertSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, got %v.", got, want)
		}
	}

	t.Run("sum some arrays", func(t *testing.T) {
		arr1 := []int{1, 3, 5, 7}
		arr2 := []int{0, 2, 4, 6}
	
		result := SumAllTails(arr1, arr2)
		expected := []int{15, 12}

		assertSum(t, result, expected)
	})

	t.Run("process empty array with security", func(t *testing.T) {
		arr1 := []int{}
		arr2 := []int{1, 2}

		result := SumAllTails(arr1, arr2)
		expected := []int{0, 2}

		assertSum(t, result, expected)
	})

}