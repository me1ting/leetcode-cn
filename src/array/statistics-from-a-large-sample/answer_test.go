package main

import "testing"

func sampleStats(count []int) []float64 {
	min, max := 0, 255

	for count[min] == 0 {
		min++
	}
	for count[max] == 0 {
		max--
	}

	if min == max {
		return []float64{float64(min), float64(min), float64(min), float64(min), float64(min)}
	}

	countAll := 0
	var sum int64 = 0
	mode := min
	for i := min; i <= max; i++ {
		if count[i] > 0 {
			countAll += count[i]
			sum += int64(i) * int64(count[i])
			if count[i] > count[mode] {
				mode = i
			}
		}
	}

	medianLCount, medianRCount := (countAll-1)/2+1, countAll/2+1
	medianL, medianR := -1, -1
	countSum := 0
	for i := min; medianL < 0 || medianR < 0; i++ {
		if medianL < 0 && countSum+count[i] >= medianLCount {
			medianL = i
		}
		if medianR < 0 && countSum+count[i] >= medianRCount {
			medianR = i
		}
		countSum += count[i]
	}

	return []float64{float64(min), float64(max), float64(sum) / float64(countAll), float64(medianL+medianR) / 2, float64(mode)}
}

func TestSampleStats(t *testing.T) {
	countList := [][]int{
		{0, 1, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 3, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	ansLsit := [][]float64{
		{1.00000, 3.00000, 2.37500, 2.50000, 3.00000},
		{1.00000, 4.00000, 2.18182, 2.00000, 1.00000},
	}

	for i, count := range countList {
		ans := ansLsit[i]
		res := sampleStats(count)

		if !sliceEquals(res, ans) {
			t.Errorf("expected %v, but got %v", ans, res)
		}
	}
}

func sliceEquals(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, x := range a {
		y := b[i]
		if !float64Equals(x, y) {
			return false
		}
	}
	return true
}

const MaxDiff = 0.00001

func float64Equals(a, b float64) bool {
	return a <= b+MaxDiff && b <= a+MaxDiff
}
