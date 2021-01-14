package answer

import "testing"

// https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies/
// 一个用拓扑排序的经典题
// 很灵活的使用了拓扑排序
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// 根据本题的需求，允许传入入度，并对**部分元素进行top排序**
	// 这里对deg进行修改说明需要正确的多次调用这个topSort，这个topSort本质是完整topSort的一部分而已
	topSort := func(graph [][]int, deg []int, items []int) (orders []int) {
		// 初始化存储入度为0的顶点的栈
		// 技巧：如何将slice用作队列
		q := []int{}
		for _, i := range items {
			if deg[i] == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			from := q[0]
			q = q[1:]
			orders = append(orders, from)
			for _, to := range graph[from] {
				deg[to]--
				if deg[to] == 0 {
					q = append(q, to)
				}
			}
		}
		return
	}

	//为item分组
	// demo这里使用m+n是因为无组项目从m组分配起，空闲并不会影响最后的结果
	// 你也可以统计有组的项目和无组的项目，最后汇总
	groupItems := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i // 注意每个无组项目单独一个组，而不是统一分配一个组，后者会影响对拓扑排序的结果
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n)
	groupDegree := make([]int, m+n)

	itemGraph := make([][]int, n)
	itemDegree := make([]int, n)

	for curr, items := range beforeItems {
		currGroup := group[curr]
		for _, before := range items {
			beforeGroup := group[before]
			if currGroup != beforeGroup {
				groupGraph[beforeGroup] = append(groupGraph[beforeGroup], currGroup)
				groupDegree[currGroup]++
			} else {
				itemGraph[before] = append(itemGraph[before], curr)
				itemDegree[curr]++
			}
		}
	}

	groups := make([]int, m+n)
	for i := range groups {
		groups[i] = i
	}

	groupOrders := topSort(groupGraph, groupDegree, groups)
	if len(groupOrders) < m+n {
		return nil
	}

	ans := []int{}
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}

	return ans
}

func TestIt(t *testing.T) {
	n := 8
	m := 2
	group := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems := [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}

	sortItems(n, m, group, beforeItems)
}
