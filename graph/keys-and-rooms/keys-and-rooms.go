package main

/*
bfs，技巧是用map来模拟队列，每层更换新的map
func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]bool{}
	queue := map[int]bool{0: true}

	for len(queue) > 0 {
		tmpQueue := map[int]bool{}
		for i, _ := range queue {
			visited[i] = true
			for _, r := range rooms[i] {
				if !queue[r] && !visited[r] {
					tmpQueue[r] = true
				}
			}
		}
		queue = tmpQueue
	}

	return len(visited) == len(rooms)
}
 */

//dfs
func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]bool{}

	var dfs func(int)

	dfs = func(i int) {
		visited[i] = true
		for _,r:=range rooms[i]{
			if !visited[r]{
				dfs(r)
			}
		}
	}

	dfs(0)

	return len(visited) == len(rooms)
}

func main() {
	canVisitAllRooms([][]int{{1}, {2}, {3}, {}})
}
