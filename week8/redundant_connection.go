//684. 冗余连接
//https://leetcode-cn.com/problems/redundant-connection/

package week8

//对于一颗树，有n个结点以及n-1条边，本题是n个结点n条边构成的一个无向有环图
//可构造一个并查集，依次遍历所有的边
//检查一条边的两个端点是否在同一个集合内，如果是不同的集合，表示这条边是树的一条有效边，合并两个集合
//反之，如果边的两个端点已经连通，在同一个集合内，则说明这条边是多余的，直接返回
//时间复杂度为O(NlogN)，因为需要遍历所有的边，然后在循环内调用unionSet
//空间复杂度为O(n)，需要一个数组来维护并查集
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	unionSet := func(x, y int) bool {
		x, y = find(x), find(y)
		if x != y {
			fa[x] = y
			return true
		}
		return false
	}

	for _, edge := range edges {
		if !unionSet(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
