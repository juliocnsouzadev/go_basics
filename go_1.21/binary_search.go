package main

import (
	"cmp"
	"slices"
	"sort"
)

type Cat struct {
	Name string
	Age  int
}

func CompareCatFunc() func(someCat, otherCat Cat) int {
	return func(someCat, otherCat Cat) int {
		if cmpName := cmp.Compare(someCat.Name, otherCat.Name); cmpName != 0 {
			return cmpName
		}
		return cmp.Compare(someCat.Age, otherCat.Age)
	}
}

func BinarySearchFunc(list []Cat, wanted Cat) (int, bool) {
	sort.Slice(list, func(i, j int) bool {
		return CompareCatFunc()(list[i], list[j]) < 0
	})
	if index, ok := slices.BinarySearchFunc(list, wanted, CompareCatFunc()); ok {
		return index, true
	}
	return -1, false
}

func BinarySearch(list []int, wanted int) (int, bool) {
	sort.Ints(list)
	if index, ok := slices.BinarySearch(list, wanted); ok {
		return index, true
	}
	return -1, false
}
