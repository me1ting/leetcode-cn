package answer

import "sort"

// https://leetcode-cn.com/problems/accounts-merge/
// 最开始想纯用hashmap，比较繁琐
// 数据的合并，最简便的方法是使用并查集
func accountsMerge(accounts [][]string) [][]string {
	email2ID := map[string]int{}
	id2Email := map[int]string{}
	email2Name := map[string]string{}

	n := 0
	for _, account := range accounts {
		name := account[0]
		emails := account[1:]

		for _, email := range emails {
			if _, ok := email2ID[email]; !ok {
				email2ID[email] = n
				id2Email[n] = email
				n++
				email2Name[email] = name
			}
		}
	}

	uf := NewUF(n)

	for _, account := range accounts {
		emails := account[1:]

		if len(emails) > 0 {
			firstID := email2ID[emails[0]]

			for _, email := range emails[1:] {
				uf.Union(firstID, email2ID[email])
			}
		}
	}

	groups := map[int][]string{}
	for i := 0; i < n; i++ {
		groupID := uf.Find(i)
		if group, ok := groups[groupID]; !ok {
			groups[groupID] = []string{email2Name[id2Email[groupID]], id2Email[i]}
		} else {
			groups[groupID] = append(group, id2Email[i])
		}
	}

	ans := [][]string{}
	for _, group := range groups {
		emails := group[1:]
		sort.Slice(emails, func(i, j int) bool {
			return emails[i] < emails[j]
		})
		ans = append(ans, group)
	}
	return ans
}

type UF struct {
	id   []int
	size []int
}

func NewUF(n int) *UF {
	id := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		rank[i] = 1
	}
	return &UF{
		id:   id,
		size: rank,
	}
}

func (u *UF) Connected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}

func (u *UF) Union(i, j int) {
	m, n := u.Find(i), u.Find(j)
	if m == n {
		return
	}
	if u.size[m] < u.size[n] {
		m, n = n, m
	}
	u.id[n] = m
	u.size[m] += u.size[n]
}

func (u *UF) Find(i int) int {
	if i != u.id[i] {
		u.id[i] = u.Find(u.id[i])
	}
	return u.id[i]
}

func (u *UF) Size(i int) int {
	return u.size[u.Find(i)]
}
