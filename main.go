package main

import (
	"fmt"
)

func main() {
	items := []int{-3, 11, -10, -5, 4, 4}
	fmt.Println(findPairs(items))
}

func findPairs(items []int) (pairs [][]int, total int) {
	timeZoneSort(items)
	sum1 := 0
	for i := 0; i < len(items); i += 2 {
		sum1 += findDistance(items[i], items[i+1])
	}
	sum2 := findDistance(items[0], items[len(items)-1])
	for i := 1; i < len(items)-1; i += 2 {
		sum2 += findDistance(items[i], items[i+1])
	}
	if sum1 <= sum2 {
		for i := 0; i < len(items); i += 2 {
			pairs = append(pairs, []int{items[i], items[i+1]})
		}
		return pairs, sum1
	} else {
		pairs = append(pairs, []int{items[0], items[len(items)-1]})
		for i := 1; i < len(items)-1; i += 2 {
			pairs = append(pairs, []int{items[i], items[i+1]})
		}
		return pairs, sum2
	}
}

func timeZoneSort(items []int) {
	var counts [23]int
	for _, item := range items {
		counts[item+11]++
	}
	currInd := 0
	for ind, count := range counts {
		timeZone := ind - 11
		for i := 0; i < count; i++ {
			items[currInd] = timeZone
			currInd++
		}
	}
}

func findDistance(time1, time2 int) int {
	return min(abs(time1-time2), 24-abs(time1-time2))
}

func abs(i int) int {
	mask := i >> 63 //get sign bit
	return (i ^ mask) - mask
}
