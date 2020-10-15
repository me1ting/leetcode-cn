package employee_importance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// https://leetcode-cn.com/problems/employee-importance/
// tree的变种，使用容器记录需要计算的员工即可
func getImportance(employees []*Employee, id int) int {
	//陷阱，没有说明employees的id与数组index有关，需要使用map来进行索引
	employeeMap := map[int]*Employee{}
	for _, e := range employees {
		employeeMap[e.Id] = e
	}
	count := 0
	var list []int
	list = append(list, id)
	for len(list) > 0 {
		last := len(list) - 1
		they := employeeMap[list[last]]
		list = list[:last]
		list = append(list, they.Subordinates...)
		count += they.Importance
	}
	return count
}
