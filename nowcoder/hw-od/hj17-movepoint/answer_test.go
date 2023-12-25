package hj17movepoint_test

import (
	hj17movepoint "algorithms/nowcoder/hw-od/hj17-movepoint"
	"testing"
)

func TestMain(t *testing.T) {
	x, y := hj17movepoint.Main()
	if x != 10 || y != -10 {
		t.Fatalf("bad answer: %d,%d", x, y)
	}
}
