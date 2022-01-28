//154. 寻找旋转排序数组中的最小值 II
//https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

package week5

func findMin(nums []int) int {
	lastIdx := len(nums) - 1
	left, right := 0, lastIdx

	//排除首尾元素重复的情况
	//最坏情况下，数组内所有元素相同，相当于遍历了一次数组，时间复杂度为O(n)
	for left < lastIdx && nums[left] == nums[lastIdx] {
		left++
	}

	//二分查找，查找第一个小于等于尾部元素的位置，时间复杂度为O(LogN)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= nums[lastIdx] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[right]
}
