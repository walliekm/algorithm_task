//1011. 在 D 天内送达包裹的能力
//https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/

package week5

func shipWithinDays(weights []int, days int) int {
	//判定载重能力为loadCapacity时能否按时送达
	canDeliver := func(loadCapacity int) bool {
		currWeight, needDays := 0, 1
		for _, w := range weights {
			if currWeight+w <= loadCapacity {
				currWeight += w
			} else {
				needDays++
				currWeight = w
			}
		}
		return needDays <= days
	}

	var left, right int
	for _, w := range weights {
		right += w
		//船的最低运载能力必须大于等于最大包裹的重量
		if left < w {
			left = w
		}
	}

	//时间复杂度NLog(N)
	for left < right {
		mid := (left + right) / 2
		if canDeliver(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
