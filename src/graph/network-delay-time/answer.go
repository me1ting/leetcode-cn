package networkdelaytime

import (
	"container/heap"
)

// https://leetcode-cn.com/problems/network-delay-time/

// 思考：所有节点收到信号的时间为单源最短路径中的极大值，如果某个节点无法到达，则不存在，则不存在
// 单源最短路径使用优先队列来实现
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]int, n+1)
	d := make([]int, n+1)
	visited := make([]bool, n+1)

	for i := 1; i < len(graph); i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j < len(graph); j++ {
			graph[i][j] = -1 //使用负值表示边不存在
		}
	}

	for _, time := range times {
		graph[time[0]][time[1]] = time[2]
	}

	h := &RecordHeap{}
	heap.Push(h, &Record{k, 0})

	for h.Len() > 0 {
		r := heap.Pop(h).(*Record)
		to := r.To
		distance := r.Distance

		if !visited[to] {
			visited[to] = true
			d[to] = distance

			for j := 1; j < n+1; j++ {
				if !visited[j] && graph[to][j] >= 0 {
					estimated := distance + graph[to][j]
					if d[j] == 0 || d[j] > estimated { //距离为0有两种情况：确实为0，但已经被标记为visited；距离未初始化，这里为后者。
						d[j] = estimated
						heap.Push(h, &Record{j, estimated})
					}
				}
			}
		}
	}

	for i := 1; i < n+1; i++ {
		if !visited[i] {
			return -1
		}
	}

	max := d[1]
	for i := 2; i < n+1; i++ {
		if d[i] > max {
			max = d[i]
		}
	}
	return max
}

type Record struct {
	To       int
	Distance int
}

//see https://pkg.go.dev/container/heap#example-package-IntHeap
type RecordHeap []*Record

func (h RecordHeap) Len() int           { return len(h) }
func (h RecordHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h RecordHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RecordHeap) Push(x interface{}) {
	*h = append(*h, x.(*Record))
}

func (h *RecordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
