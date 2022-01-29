//911. 在线选举
//https://leetcode-cn.com/problems/online-election/

package week5

type TopVotedCandidate struct {
	times  []int //时间列表
	ansArr []int //答案数组，其第i个元素为times[i]时刻的答案
}

//遍历persons数组，统计得票情况，将当前最高得票人放入ansArr
//需要额外的ansArr数组记录每阶段的答案，以及votedMap统计得票情况，空间复杂度为O(n)
//需要对persons进行一次数组遍历，时间复杂度为O(n)
func Constructor(persons []int, times []int) TopVotedCandidate {
	ansArr := make([]int, len(persons)) //答案数组，表示对应times位置上得票领先的候选人编号
	votedMap := map[int]int{}           //得票统计，key为候选人编号，val为票数

	for idx, p := range persons {
		votedMap[p]++ //p的得票加1

		//起始位置，得票领先者即当前获票人
		if idx == 0 {
			ansArr[idx] = p
			continue
		}

		//将当前得票人的票数与领先者比较，判定本次计票过后的获胜者，放入ansArr对应的位置
		if votedMap[p] >= votedMap[ansArr[idx-1]] {
			ansArr[idx] = p
		} else {
			ansArr[idx] = ansArr[idx-1]
		}
	}

	return TopVotedCandidate{times, ansArr}
}

func (this *TopVotedCandidate) Q(t int) int {
	//二分查找，查找第一个 <=t 的位置，时间复杂度为: O(LogN)
	left, right := 0, len(this.times)-1
	for left < right {
		mid := (left + right + 1) / 2
		if this.times[mid] <= t {
			left = mid
		} else {
			right = mid - 1
		}
	}

	//ansArr上对应位置的值即为答案
	return this.ansArr[right]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
