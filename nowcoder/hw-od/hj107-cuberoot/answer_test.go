package hj107cuberoot_test

import "fmt"

func main() {
	var n float64
	fmt.Scanf("%v", &n)
	fmt.Printf("%.1f", cuberRoot(n))
}

const diff = 0.000001

func cuberRoot(n float64) float64 {
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	min, max := float64(1), n
	if n < 1 {
		min, max = n, 1
	}

	var mid float64
	for {
		mid = (max-min)/2 + min
		if mid*mid*mid > n+diff {
			max = mid
		} else if mid*mid*mid+diff < n {
			min = mid
		} else {
			break
		}
	}

	if negative {
		return -mid
	}
	return mid
}
