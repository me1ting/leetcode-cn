package hj29strencode_test

import (
	hj29strencode "algorithms/nowcoder/hw-od/hj29-strencode"
	"log"
	"testing"
)

func TestAnswer(t *testing.T) {
	testcases := []string{"2OA92AptLq5G1lW8564qC4nKMjv8C", "B5WWIj56vu72GzRja7j5"}
	expected := []string{"3pb03bQUmR6h2Mx9675Rd5OlnKW9d", "a4vvhI45UT61fYqIZ6I4"}

	result := hj29strencode.Answer(testcases)

	for i, str := range expected {
		if result[i] != str {
			log.Fatalf("expected %s, return %s", str, result[i])
		}
	}
}
