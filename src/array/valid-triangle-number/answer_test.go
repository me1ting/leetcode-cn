package validtrianglenumber

import "testing"

func TestTriangleNumber(t *testing.T) {
	nums := []int{2, 2, 3, 4}
	expected := 3

	val := triangleNumber(nums)
	if val != expected {
		t.Errorf("expected %v, return %v", expected, val)
	}
}
