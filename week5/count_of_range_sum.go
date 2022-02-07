//327. 区间和的个数
//https://leetcode-cn.com/problems/count-of-range-sum/

package week5

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)

	//前缀和数组
	preSumArr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSumArr[i] = preSumArr[i-1] + nums[i-1]
	}

	//查找前缀和数组中第一个>=target的位置
	findFirstMinVal := func(target, left, right int) int {
		for left < right {
			mid := (left + right) / 2
			if preSumArr[mid] >= target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return right
	}

	//查找前缀和数组中最后一个<=target的位置
	findLastMaxVal := func(target, left, right int) int {
		for left < right {
			mid := (left + right + 1) / 2
			if preSumArr[mid] <= target {
				left = mid
			} else {
				right = mid - 1
			}
		}
		return right
	}

	merge := func(left, mid, right int) {
		i, j := left, mid+1
		tmpArr := make([]int, right-left+1)
		for k := range tmpArr {
			if j > right || (i <= mid && preSumArr[i] <= preSumArr[j]) {
				tmpArr[k] = preSumArr[i]
				i++
			} else {
				tmpArr[k] = preSumArr[j]
				j++
			}
		}

		for k := range tmpArr {
			preSumArr[left+k] = tmpArr[k]
		}
	}

	ans := 0
	calculate := func(left, mid, right int) {
		//数组的区间和S(i,j) = preSumArr[j] - preSumArr[i]
		//对前缀和数组左侧进行枚举，查找其右侧满足区间和在lower~upper之间的范围
		for i := left; i <= mid; i++ {
			start := findFirstMinVal(lower+preSumArr[i], mid+1, right+1)
			end := findLastMaxVal(upper+preSumArr[i], mid, right)
			if end >= start {
				ans += (end - start + 1)
			}
		}
	}

	var mergeSort func(left, right int)
	mergeSort = func(left, right int) {
		if left >= right {
			return
		}
		mid := (left + right) / 2
		mergeSort(left, mid)
		mergeSort(mid+1, right)
		calculate(left, mid, right)
		merge(left, mid, right)
	}

	mergeSort(0, n)
	return ans
}
