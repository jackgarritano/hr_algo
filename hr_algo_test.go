package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	pairs, total := findPairs([]int{-3, 11, -10, -5, 4, 4})
	if total != 5 {
		t.Errorf("findPairs([-3, 11, -10, -5, 4, 4]) got total %d", total)
	}
	if !slicesEqual(pairs, [][]int{{-5, -3}, {11, -10}, {4, 4}}) {
		t.Errorf("findPairs([-3, 11, -10, -5, 4, 4]) got pair %d", pairs)
	}
}

func Test2(t *testing.T) {
	pairs, total := findPairs([]int{2, 10, 11, -9, -3, -4})
	if total != 11 {
		t.Errorf("findPairs([2, 10, 11, -9, -3, -4]) got total %d", total)
	}
	if !slicesEqual(pairs, [][]int{{10, 11}, {-9, -4}, {2, -3}}) {
		t.Errorf("findPairs([2, 10, 11, -9, -3, -4]) got pair %d", pairs)
	}
}

func Test3(t *testing.T) {
	pairs, total := findPairs([]int{2, 2})
	if total != 0 {
		t.Errorf("findPairs([2, 2]) got total %d", total)
	}
	if !slicesEqual(pairs, [][]int{{2, 2}}) {
		t.Errorf("findPairs([2, 2]) got pair %d", pairs)
	}
}

func Test4(t *testing.T) {
	pairs, total := findPairs([]int{2, -2})
	if total != 4 {
		t.Errorf("findPairs([2, -2]) got total %d", total)
	}
	if !slicesEqual(pairs, [][]int{{2, -2}}) {
		t.Errorf("findPairs([2, -2]) got pair %d", pairs)
	}
}

// Function to check if two slices of slices are equal disregarding order
func slicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort slices of slices
	sortSliceOfSlices(a)
	sortSliceOfSlices(b)

	// Compare sorted slices of slices
	return reflect.DeepEqual(a, b)
}

func sortSliceOfSlices(slice [][]int) {
	for _, innerSlice := range slice {
		sort.Ints(innerSlice) // Sort each inner slice
	}
	// Sort the outer slice based on the sorted inner slices
	sort.Slice(slice, func(i, j int) bool {
		for x := 0; x < len(slice[i]) && x < len(slice[j]); x++ {
			if slice[i][x] != slice[j][x] {
				return slice[i][x] < slice[j][x]
			}
		}
		return len(slice[i]) < len(slice[j])
	})
}
