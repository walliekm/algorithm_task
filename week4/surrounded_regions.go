//130. 被围绕的区域
//https://leetcode-cn.com/problems/surrounded-regions/
//对矩阵进行一次遍历，时间复杂度为O(m*n)
//需要额外的visted矩阵和idxList数组以及队列q，每个元素只占一个位置，空间复杂度为O(m*n)

package week4

import "container/list"

func solve(board [][]byte) {
	//m行n列
	m := len(board)
	n := len(board[0])

	//方向数组
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	type idx struct {
		ridx int //行号
		cidx int //列号
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	//广度优先遍历连通字符为O的块
	q := list.New()
	bfs := func(i, j int) {
		isSurrounded := true //是否有效的被包围区域
		idxList := []*idx{}  //包围区域内的元素下标
		visited[i][j] = true
		q.PushBack(&idx{i, j})

		for q.Len() > 0 {
			p := q.Remove(q.Front()).(*idx)

			//元素位于边界上，不是有效的包围区域
			if p.ridx == 0 || p.ridx == m-1 || p.cidx == 0 || p.cidx == n-1 {
				isSurrounded = false
			} else {
				idxList = append(idxList, p)
			}

			for k := 0; k < 4; k++ {
				ni := p.ridx + dx[k]
				nj := p.cidx + dy[k]

				if ni < 0 || ni >= m || nj < 0 || nj >= n {
					continue
				}

				if board[ni][nj] != 'O' || visited[ni][nj] {
					continue
				}

				visited[ni][nj] = true
				q.PushBack(&idx{ni, nj})
			}
		}

		//有效的包围区域，填充X
		if isSurrounded {
			for _, p := range idxList {
				board[p.ridx][p.cidx] = 'X'
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				bfs(i, j)
			}
		}
	}
}
