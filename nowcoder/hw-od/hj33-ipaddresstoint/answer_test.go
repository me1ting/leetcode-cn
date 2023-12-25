package hj33ipaddresstoint_test

import (
	hj33ipaddresstoint "algorithms/nowcoder/hw-od/hj33-ipaddresstoint"
	"testing"
)

func TestIntToIp(t *testing.T) {
	testcase := 167969729
	expected := "10.3.3.193"

	if hj33ipaddresstoint.IntToIp(testcase) != expected {
		t.Fatal("bad answer")
	}
}

func TestIpToInt(t *testing.T) {
	testcase := "10.3.3.193"
	expected := 167969729

	if hj33ipaddresstoint.IpToInt(testcase) != expected {
		t.Fatal("bad answer")
	}
}
