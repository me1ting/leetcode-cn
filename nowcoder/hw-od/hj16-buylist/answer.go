package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := readInputs()
	result := answer(inputs.N, inputs.M, inputs.items)
	fmt.Println(result)
}

func answer(N int, m int, items []*Item) int {
	N /= 10
	groups := buildGroups(items)

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, len(groups)+1)
	}
	for i := 1; i <= N; i++ {
		for arrIndex, g := range groups {
			j := arrIndex + 1
			dp[i][j] = dp[i][j-1]
			for _, m := range g.members {
				if m.price/10 <= i {
					dp[i][j] = MaxInt(dp[i][j], dp[i-m.price/10][j-1]+m.sat)
				}
			}
		}
	}

	return dp[N][len(groups)]
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func buildGroups(items []*Item) []*Group {
	var groups = map[int]*Group{}
	for _, item := range items {
		if item.belong == 0 {
			g, ok := groups[item.num]
			if !ok {
				g = &Group{}
				groups[item.num] = g
			}
			g.main = item
		} else {
			g, ok := groups[item.belong]
			if !ok {
				g = &Group{}
				groups[item.belong] = g
			}
			g.attachments = append(g.attachments, item)
		}
	}
	list := []*Group{}
	for _, g := range groups {
		g.build()
		list = append(list, g)
	}
	return list
}

type Group struct {
	main        *Item
	attachments []*Item
	members     []*Member //0-1分组背包问题的组员
}

type Member struct {
	price int
	sat   int // satisfaction
}

func (g *Group) build() {
	if len(g.attachments) >= 0 {
		g.members = append(g.members, &Member{
			price: g.main.price,
			sat:   g.main.satisfaction(),
		})
	}
	if len(g.attachments) >= 1 {
		g.members = append(g.members, &Member{
			price: g.attachments[0].price + g.main.price,
			sat:   g.attachments[0].satisfaction() + g.main.satisfaction(),
		})
	}
	if len(g.attachments) >= 2 {
		g.members = append(g.members, &Member{
			price: g.attachments[1].price + g.main.price,
			sat:   g.attachments[1].satisfaction() + g.main.satisfaction(),
		})
		g.members = append(g.members, &Member{
			price: g.attachments[0].price + g.attachments[1].price + g.main.price,
			sat:   g.attachments[0].satisfaction() + g.attachments[1].satisfaction() + g.main.satisfaction(),
		})
	}
}

type Item struct {
	num        int
	price      int
	importance int
	belong     int
}

func (i *Item) satisfaction() int {
	return i.price * i.importance
}

type Inputs struct {
	N     int
	M     int
	items []*Item
}

func readInputs() Inputs {
	var inputs Inputs
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	metaLine := lines[0]
	nums := strings.Split(metaLine, " ")
	inputs.N, _ = strconv.Atoi(nums[0])
	inputs.M, _ = strconv.Atoi(nums[1])

	for i, line := range lines[1:] {
		var item Item
		item.num = i + 1

		nums := strings.Split(line, " ")
		item.price, _ = strconv.Atoi(nums[0])
		item.importance, _ = strconv.Atoi(nums[1])
		item.belong, _ = strconv.Atoi(nums[2])

		inputs.items = append(inputs.items, &item)
	}

	return inputs
}
