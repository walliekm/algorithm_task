package week3

import (
	"container/list"
)

type chkHelper struct {
	size  int     //结点个数
	inDeg []int   //入度数组
	toArr [][]int //有向图的出边数组
}

//检查给定的数据是否一棵有根树
//根结点没有父结点，入度数为零
//子结点有且仅有一个父结点，即子结点的入度数只能为一
func (this *chkHelper) isTree(excludeEdge []int) bool {
	//根结点
	root := 0

	//遍历入度数组，查找根结点(入度数=0)
	for i := 1; i <= this.size; i++ {
		deg := this.inDeg[i]

		//排除掉一条边，将其入度减一
		if i == excludeEdge[1] {
			deg -= 1
		}

		//当前结点入度大于1，肯定不是合法树
		if deg > 1 {
			return false
		}

		//不止一个结点入度为零，也不是合法树
		if deg == 0 {
			if root != 0 {
				return false
			}

			root = i
		}
	}

	//找不到根结点，不是合法树
	if root == 0 {
		return false
	}

	//标记结点是否已走过
	visited := make([]bool, this.size+1)

	//树遍历走过的节点数量
	visCnt := 0

	//广度优先遍历，将根结点入队
	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		k := q.Remove(q.Front()).(int)
		visCnt++
		visited[k] = true

		for _, v := range this.toArr[k] {
			//排除路径
			if k == excludeEdge[0] && v == excludeEdge[1] {
				continue
			}

			//结点已走过
			if visited[v] {
				continue
			}

			//未走过的节点入队
			q.PushBack(v)
		}
	}

	//遍历结束，如果已走结点数等于总结点数，则说明是有效的树
	return visCnt == this.size
}

func findRedundantDirectedConnection(edges [][]int) []int {
	//结点数，源数据为n个节点n条边
	n := len(edges)

	inDeg := make([]int, n+1)   //入度数组
	toArr := make([][]int, n+1) //有向图的出边数组
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		toArr[from] = append(toArr[from], to)
		inDeg[to]++
	}

	chk := &chkHelper{n, inDeg, toArr}

	//从后往前，依次排除一条边，然后判断剩余路径是否一棵树
	for i := n - 1; i >= 0; i-- {
		if chk.isTree(edges[i]) {
			return edges[i]
		}
	}

	return nil
}
