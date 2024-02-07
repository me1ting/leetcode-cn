package maximumprofitofoperatingacentennialwheel_test

import "testing"

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	if boardingCost*4 <= runningCost {
		return -1
	}
	persons := 0
	income := 0
	maxIncome := 0
	maxIncomeI := -1
	for i := 0; i < len(customers) || persons > 0; i++ {
		if (i) < len(customers) {
			persons += customers[i]
		}
		boardingCount := min(persons, 4)
		persons -= boardingCount

		income += boardingCount*boardingCost - runningCost
		if income > maxIncome {
			maxIncome = income
			maxIncomeI = i + 1
		}
	}

	return maxIncomeI
}

func TestIt(t *testing.T) {
	customersList := [][]int{
		{8, 3},
		{10, 9, 6},
		{3, 4, 0, 5, 1},
	}
	boardingCostList := []int{
		5, 6, 1,
	}
	runningCostList := []int{
		6, 4, 92,
	}
	expectedList := []int{
		3, 7, -1,
	}

	for i := 0; i < len(customersList); i++ {
		customers := customersList[i]
		boardingCost := boardingCostList[i]
		runningCost := runningCostList[i]
		expected := expectedList[i]

		result := minOperationsMaxProfit(customers, boardingCost, runningCost)
		if result != expected {
			t.Fatalf("expected %v,result %v", expected, result)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
