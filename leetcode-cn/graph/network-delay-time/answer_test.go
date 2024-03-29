package networkdelaytime

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	timess := [][][]int{
		{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
		{{1, 2, 1}},
		{{1, 2, 1}},
		{{3, 5, 78}, {2, 1, 1}, {1, 3, 0}, {4, 3, 59}, {5, 3, 85}, {5, 2, 22}, {2, 4, 23}, {1, 4, 43}, {4, 5, 75}, {5, 1, 15}, {1, 5, 91}, {4, 1, 16}, {3, 2, 98}, {3, 4, 22}, {5, 4, 31}, {1, 2, 0}, {2, 5, 4}, {4, 2, 51}, {3, 1, 36}, {2, 3, 59}},
	}
	ns := []int{
		4,
		2,
		2,
		5,
	}
	ks := []int{
		2,
		1,
		2,
		5,
	}
	expecteds := []int{
		2,
		1,
		-1,
		31,
	}

	for i := 3; i < len(timess); i++ {
		times := timess[i]
		n := ns[i]
		k := ks[i]
		expected := expecteds[i]

		val := networkDelayTime(times, n, k)
		if val != expected {
			t.Errorf("expected %v, return %v", expected, val)
		}
	}
}
