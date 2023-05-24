package main

import "testing"

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	adjTable := make([][]int, n+1)
	probabilities := make([]float64, n+1)
	depths := make([]int, n+1)
	parent := make([]int, n+1)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		adjTable[from] = append(adjTable[from], to)
		adjTable[to] = append(adjTable[to], from)
	}

	probabilities[1] = 1
	depths[1] = 0
	parent[1] = 0

	var dfs func(int)
	dfs = func(i int) {
		var p float64
		if i == 1 {
			p = probabilities[i] / float64(len(adjTable[i]))
		} else {
			p = probabilities[i] / float64(len(adjTable[i])-1)
		}

		for _, j := range adjTable[i] {
			if j == parent[i] {
				continue
			}
			probabilities[j] = p
			depths[j] = depths[i] + 1
			parent[j] = i
			dfs(j)
		}
	}

	dfs(1)

	if t < depths[target] {
		return 0
	} else if t == depths[target] {
		return probabilities[target]
	}

	if target == 1 && len(adjTable[target]) == 0 ||
		target > 1 && len(adjTable[target]) == 1 {
		return probabilities[target]
	}

	return 0
}

func TestFrogPosition(test *testing.T) {
	n := 7
	edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
	t := 2
	target := 4
	expected := 0.166666666

	maxDiff := 0.00001
	result := frogPosition(n, edges, t, target)

	if result > expected+maxDiff || result < expected-maxDiff {
		test.Errorf("expected %f, but got %f", expected, result)
	}

	n = 3
	edges = [][]int{{2, 1}, {3, 2}}
	t = 1
	target = 2
	expected = 1

	maxDiff = 0.00001
	result = frogPosition(n, edges, t, target)

	if result > expected+maxDiff || result < expected-maxDiff {
		test.Errorf("expected %f, but got %f", expected, result)
	}

	n = 8
	edges = [][]int{{2, 1}, {3, 2}, {4, 1}, {5, 1}, {6, 4}, {7, 1}, {8, 7}}
	t = 7
	target = 7
	expected = 0

	maxDiff = 0.00001
	result = frogPosition(n, edges, t, target)

	if result > expected+maxDiff || result < expected-maxDiff {
		test.Errorf("expected %f, but got %f", expected, result)
	}
}
