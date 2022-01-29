//875. 爱吃香蕉的珂珂
//https://leetcode-cn.com/problems/koko-eating-bananas/

package week5

func minEatingSpeed(piles []int, h int) int {
	//判定在k小时内能否吃完
	canEat := func(k int) bool {
		needHours := 0
		for i := range piles {
			//吃第i堆香蕉需要的时间，向上取整
			needHours += (piles[i] + k - 1) / k
		}
		return needHours <= h
	}

	left, right := 1, 0
	for i := range piles {
		right += piles[i]
	}

	//时间复杂度NLog(N)
	for left < right {
		mid := (left + right) / 2
		if canEat(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
