package main

import (
	"sort"
	"testing"
)

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}

	sort.Slice(ids, func(i, j int) bool {
		return values[ids[i]] > values[ids[j]]
	})

	chooseCount := map[int]int{}

	sum, count := 0, 0
	for i := 0; i < n; i++ {
		if count == numWanted {
			break
		}
		value, label := values[ids[i]], labels[ids[i]]
		if chooseCount[label] == useLimit {
			continue
		}
		sum += value
		chooseCount[label]++
		count++
	}

	return sum
}

func TestLargestValsFromLabels(t *testing.T) {
	values := []int{5, 4, 3, 4, 2}
	labels := []int{1, 2, 2, 3, 3}
	numWanted := 3
	useLimit := 1
	expected := 13

	val := largestValsFromLabels(values, labels, numWanted, useLimit)
	if val != expected {
		t.Errorf("expected %d, but got %d", expected, val)
	}
}
