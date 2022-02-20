//55. 跳跃游戏
//https://leetcode-cn.com/problems/jump-game/

package week7

//维护一个变量m，表示从当前位置能跳跃的最远距离
//遍历整个数组，动态更新m值
//最后判断m的位置是否到达或超过数组的结尾
//时间复杂度O(n)，空间复杂度O(1)
func canJump(nums []int) bool {
	m := 0
	for i := range nums {
		if i <= m {
			m = max(m, i+nums[i])
		}
	}

	return m >= len(nums)-1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
