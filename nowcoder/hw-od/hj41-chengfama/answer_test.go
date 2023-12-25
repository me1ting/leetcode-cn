package hj41chengfama_test

import (
	hj41chengfama "algorithms/nowcoder/hw-od/hj41-chengfama"
	"testing"
)

func TestAnswer(t *testing.T) {
	types := 10
	weights := []int{2000, 1999, 1998, 1997, 1996, 1995, 1994, 1993, 1992, 1991}
	counts := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

	if hj41chengfama.Answer2(types, weights, counts) != 16601 {
		t.Fatal("bad answer")
	}
}
