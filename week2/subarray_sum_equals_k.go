//560. 和为 K 的子数组
//https://leetcode-cn.com/problems/subarray-sum-equals-k/

package week2

func subarraySum(nums []int, k int) int {
	nLen := len(nums)
	preSumArr := make([]int, nLen+1) //前缀和数组
	for i := 1; i <= nLen; i++ {
		preSumArr[i] = preSumArr[i-1] + nums[i-1]
	}

	ans := 0
	countMap := map[int]int{} //统计前缀和出现的次数
	for i := 0; i <= nLen; i++ {
		//countMap中存放的是i之前的前缀和统计数据
		//countMap中若存在元素key=j，满足preSumArr[i] - preSumArr[j] = k，即为答案
		if dVal := preSumArr[i] - k; countMap[dVal] > 0 {
			ans += countMap[dVal]
		}
		countMap[preSumArr[i]]++
	}

	return ans
}
