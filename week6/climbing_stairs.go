//70. 爬楼梯
//https://leetcode-cn.com/problems/climbing-stairs/description/

package week6

//以f[i]表示爬到第i级台阶的方法
//要到达第i级台阶，可以在第i-1处走1级台阶，或者在i-2处走2级台阶而来
//由此可得转移方程 f[i] = f[i-1] + f[i-2]
//对于f(0)可视为只有1种方法，f(1)也只能走1级台阶
//f(i)只跟i-1和i-2有关，因此可用三个变量作为滚动数组存储
//时间复杂度为O(n)，需要从1~n遍历一次
//空间复杂度为O(1)，需要三个变量存储数据
func climbStairs(n int) int {
	p, q, r := 0, 0, 1

	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
