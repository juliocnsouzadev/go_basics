package main

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		list          []int
		wanted        int
		expectedIndex int
		expectedOk    bool
	}{
		{[]int{}, 1, -1, false},
		{[]int{1}, 1, 0, true},
		{[]int{1}, 2, -1, false},
		{[]int{1, 2, 3, 4, 5}, 3, 2, true},
		{[]int{1, 2, 3, 4, 5}, 6, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearch(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearch(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}

func TestBinarySearchFunc(t *testing.T) {
	testCases := []struct {
		list          []Cat
		wanted        Cat
		expectedIndex int
		expectedOk    bool
	}{
		{[]Cat{}, Cat{Name: "Fluffy", Age: 2}, -1, false},
		{[]Cat{{Name: "Fluffy", Age: 2}}, Cat{Name: "Fluffy", Age: 2}, 0, true},
		{[]Cat{{Name: "Fluffy", Age: 2}}, Cat{Name: "Mittens", Age: 3}, -1, false},
		{[]Cat{{Name: "Fluffy", Age: 2}, {Name: "Mittens", Age: 3}}, Cat{Name: "Mittens", Age: 3}, 1, true},
		{[]Cat{{Name: "Fluffy", Age: 2}, {Name: "Mittens", Age: 3}}, Cat{Name: "Whiskers", Age: 1}, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearchFunc(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearchFunc(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}

func TestBinarySearchUnsorted(t *testing.T) {
	testCases := []struct {
		list          []int
		wanted        int
		expectedIndex int
		expectedOk    bool
	}{
		{[]int{3, 2, 1}, 2, 1, true},
		{[]int{1, 3, 2}, 2, 1, true},
		{[]int{1, 2, 3}, 4, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearch(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearch(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}

func TestBinarySearchFuncUnsorted(t *testing.T) {
	testCases := []struct {
		list          []Cat
		wanted        Cat
		expectedIndex int
		expectedOk    bool
	}{
		{[]Cat{{Name: "Mittens", Age: 2}, {Name: "Fluffy", Age: 3}, {Name: "Whiskers", Age: 1}}, Cat{Name: "Mittens", Age: 2}, 1, true},
		{[]Cat{{Name: "Fluffy", Age: 3}, {Name: "Fluffy", Age: 2}, {Name: "Whiskers", Age: 1}}, Cat{Name: "Fluffy", Age: 2}, 0, true},
		{[]Cat{{Name: "Fluffy", Age: 2}, {Name: "Mittens", Age: 3}, {Name: "Whiskers", Age: 1}}, Cat{Name: "Tiger", Age: 1}, -1, false},
	}

	for _, tc := range testCases {
		actualIndex, ok := BinarySearchFunc(tc.list, tc.wanted)
		if actualIndex != tc.expectedIndex || ok != tc.expectedOk {
			t.Errorf("BinarySearchFunc(%v, %v) = (%v, %v), expected (%v, %v)", tc.list, tc.wanted, actualIndex, ok, tc.expectedIndex, tc.expectedOk)
		}
	}
}
