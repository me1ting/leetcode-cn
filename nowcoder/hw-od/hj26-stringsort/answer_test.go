package hj26stringsort_test

import (
	hj26stringsort "algorithms/nowcoder/hw-od/hj26-stringsort"
	"testing"
)

func TestAnswer(t *testing.T) {
	testcase := "A Famous Saying: Much Ado About Nothing (2012/8)."
	expected := "A aaAAbc dFgghh: iimM nNn oooos Sttuuuy (2012/8)."

	bs := []byte(testcase)
	hj26stringsort.CustomSort(bs)
	if string(bs) != expected {
		t.Fatalf("bad answer:%v", string(bs))
	}
}
