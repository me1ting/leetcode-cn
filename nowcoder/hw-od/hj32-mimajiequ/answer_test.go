package hj32mimajiequ_test

import (
	hj32mimajiequ "algorithms/nowcoder/hw-od/hj32-mimajiequ"
	"testing"
)

func TestAnswer(t *testing.T) {
	c := "ABBA"
	r := 4

	if hj32mimajiequ.Answer(c) != r {
		t.Fatal("bad answer")
	}
}
