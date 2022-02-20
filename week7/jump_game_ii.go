//45. 跳跃游戏 II
//https://leetcode-cn.com/problems/jump-game-ii/

package week7

//以f[i]表示到达i位置需要的最少跳跃次数
//则有当 0 <= j < i时，如果满足j + nums[j] >= i，则有转移方程: f[i] = min(f[j]) + 1
//初值f[0]=0，其它为正无穷
//目标f[n-1]
//时间复杂度O(n*n)，空间复杂度O(n)
func jump(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	for i := range f {
		f[i] = 1e9
	}
	f[0] = 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}

	return f[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
