package hj20passwodchecker_test

import (
	hj20passwodchecker "algorithms/nowcoder/hw-od/hj20-passwodchecker"
	"testing"
)

func TestCheck(t *testing.T) {
	testcase := "021Abc9000"
	expected := true

	if hj20passwodchecker.Check(testcase) != expected {
		t.Fatal("bad answer")
	}
}
