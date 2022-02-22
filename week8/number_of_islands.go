//200. 岛屿数量
//https://leetcode-cn.com/problems/number-of-islands/

package week8

//维护一个并查集数组，初始化时对二维数组中所有为1的位置标记为一个集合
//遍历二维数组，将相邻的陆地合并
//最终获取集合的个数，即为答案
//时间复杂度为O(NlogN)，N为二维矩阵的元素个数
//空间复杂度为O(n)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	fa := make([]int, m*n)
	for i := range fa {
		//并查集初始化，二维矩阵中等于1的位置才有效，否则赋值为-1，表示无效
		if grid[i/n][i%n] == '1' {
			fa[i] = i
		} else {
			fa[i] = -1
		}
	}

	//二维矩阵下标转换为一维数组下标
	idx := func(i, j int) int {
		return i*n + j
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	unionSet := func(x, y int) {
		x, y = find(x), find(y)
		if x != y {
			fa[x] = y
		}
	}

	//方向数组，只向右向下探寻，避免重复
	dx := []int{1, 0}
	dy := []int{0, 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			for k := 0; k < 2; k++ {
				ni := i + dx[k]
				nj := j + dy[k]
				if ni >= m || nj >= n || grid[ni][nj] == '0' {
					continue
				}

				//合并相邻的陆地
				unionSet(idx(i, j), idx(ni, nj))
			}
		}
	}

	ans := 0
	for i := range fa {
		//答案为集合数（即根节点的个数）
		if fa[i] == i {
			ans++
		}
	}
	return ans
}
