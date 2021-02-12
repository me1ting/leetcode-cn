package answer

import "testing"

func getRow(rowIndex int) []int {
	row := []int{1}

	for i := 2; i <= rowIndex+1; i++ {
		nextRow := make([]int, 0, len(row)+1)
		nextRow = append(nextRow, 1)
		for j := 1; j < i-1; j++ {
			nextRow = append(nextRow, row[j-1]+row[j])
		}
		nextRow = append(nextRow, 1)
		row = nextRow
	}

	return row
}

func TestGetRow(t *testing.T) {
	rowIndex := 4
	row := []int{1, 4, 6, 4, 1}
	val := getRow(rowIndex)

	if len(row) < len(val) {
		t.Errorf("return %v, expected %v", val, row)
	}

	for i := 0; i < len(val); i++ {
		if val[i] != row[i] {
			t.Errorf("return %v, expected %v", val, row)
		}
	}
}
