package dota2_senate

// https://leetcode-cn.com/problems/dota2-senate/
// 个人认为此题的一个判断点在于：
// - ban本轮已经使用了ban权的人，还是ban本轮还未使用ban权的人？
// - 如果ban未使用ban权的人，ban位于前面和后面有什么区别？
// 2. 考虑1100110011001100，如果头两个11去ban最末尾的00，则他们将输，而如果依次ban则它们会赢
// 1. 等同于2。将议会想象成一个环，ban人就是不断缩圈的情况，直到一方获胜。本轮已经使用ban权的人
// 等同于放在当前轮的末尾，因此应当ban本轮未使用ban权的人

// 本题属于贪心算法的习题，这里的贪心就是ban最前面的人是局部最优解（当前轮），也是全局最优解（结果）。

// 实现技巧：算法有了，但如何精巧的实现还需要花一番心思，参考题解：
// https://leetcode-cn.com/problems/dota2-senate/solution/dota2-can-yi-yuan-by-leetcode/
// 此题解的存在许多精巧的地方，如：
// 1. 使用0,1表示阵营；
// 2. 统计人数，人数用于判断结束点
// 3. 统计ban权，在遍历到对方人选时使用，也就是题解中所说的
// “参议员不需要在轮到自己的时候就立即使用禁令，可以等待另一个阵营的参议员投票时再使用。”
// 这是一个很重要的技巧，避免了记录或是查找下一个对方人选的位置
func predictPartyVictory(senate string) string {
	ll := &LinkedList{}
	people, bans := [2]int{0, 0}, [2]int{0, 0}
	const Radiant, Dire byte = 0, 1

	for i := 0; i < len(senate); i++ {
		camp := Radiant
		if senate[i] == 'D' {
			camp = Dire
		}
		people[camp]++

		ll.appendLast(camp)
	}

	for people[Radiant] > 0 && people[Dire] > 0 {
		camp := ll.removeFirst()
		enemy := camp ^ 1
		if bans[enemy] > 0 {
			bans[enemy]--
			people[camp]--
		} else {
			bans[camp]++
			ll.appendLast(camp)
		}
	}

	if people[Radiant]>0 {
		return "Radiant"
	}else{
		return "Dire"
	}
}

type LinkedList struct {
	head, tail *Node
}

type Node struct {
	Next *Node
	Val  byte
}

func (ll *LinkedList) appendLast(val byte) {
	node := &Node{
		Next: nil,
		Val:  val,
	}
	if ll.tail != nil {
		ll.tail.Next = node
	} else {
		ll.head = node
	}

	ll.tail = node
}

func (ll *LinkedList) removeFirst() byte {
	node := ll.head
	if node.Next == nil {
		ll.tail = nil
	}
	ll.head = node.Next
	return node.Val
}
