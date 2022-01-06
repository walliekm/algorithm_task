//697. 数组的度
//https://leetcode-cn.com/problems/degree-of-an-array/

package week2

//数组元素统计信息
type numStatis struct {
	cnt  int //元素出现的频次
	fidx int //首次出现的位置
	lidx int //最后出现的位置
}

func findShortestSubArray(nums []int) int {
	numStatisMap := map[int]*numStatis{}
	for idx, num := range nums {
		if stInfo, ok := numStatisMap[num]; !ok {
			numStatisMap[num] = &numStatis{cnt: 1, fidx: idx, lidx: idx}
		} else {
			stInfo.cnt++
			stInfo.lidx = idx
		}
	}

	var ansStatis *numStatis
	for _, stInfo := range numStatisMap {
		//频次更高，直接更新结果
		if ansStatis == nil || stInfo.cnt > ansStatis.cnt {
			ansStatis = stInfo
		}

		//频次相等，比较距离长短
		if stInfo.cnt == ansStatis.cnt {
			if stInfo.lidx-stInfo.fidx < ansStatis.lidx-ansStatis.fidx {
				ansStatis = stInfo
			}
		}
	}

	return ansStatis.lidx - ansStatis.fidx + 1
}
