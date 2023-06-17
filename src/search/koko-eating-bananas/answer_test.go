package kokoeatingbananas

import "testing"

func minEatingSpeed(piles []int, h int) int {
	theMax := piles[0]
	for _, pile := range piles {
		if pile > theMax {
			theMax = pile
		}
	}

	i, j := 1, theMax

	for i < j {
		mid := i + (j-i)/2
		if canEatOver(piles, h, mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

func canEatOver(piles []int, h int, k int) bool {
	i, n := 0, 0
	for ; n < h && i < len(piles); i++ {
		pile := piles[i]
		n += pile / k
		if pile%k != 0 {
			n++
		}
	}

	return i == len(piles) && n <= h
}

func TestMinEatingSpeed(t *testing.T) {
	piles := []int{312884470}
	h := 312884469
	rightAns := 2

	ans := minEatingSpeed(piles, h)
	if ans != rightAns {
		t.Errorf("bad ans")
	}
}
