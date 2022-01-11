//210. 课程表 II
//https://leetcode-cn.com/problems/course-schedule-ii/

package week3

import "container/list"

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDeg := make([]int, numCourses)   //入度数组，标记选修某课程前需要修多少门先修课
	toArr := make([][]int, numCourses) //出边数组，标记完成某课程后，可继续选修其它的课程

	for _, pre := range prerequisites {
		//y为x的先修课
		x, y := pre[0], pre[1]
		toArr[y] = append(toArr[y], x)
		inDeg[x]++
	}

	//选修队列，按先进先出顺序选修课程
	queue := list.New()

	//遍历入度数组，入度数为零的课程进入选修队列
	for i := range inDeg {
		if inDeg[i] == 0 {
			queue.PushBack(i)
		}
	}

	ans := make([]int, 0, numCourses)
	for queue.Len() > 0 {
		x := queue.Remove(queue.Front()).(int)
		ans = append(ans, x)

		//遍历课程x的后续课程，将其入度减一，当入度为零时，进入选修队列
		for _, y := range toArr[x] {
			inDeg[y]--
			if inDeg[y] == 0 {
				queue.PushBack(y)
			}
		}
	}

	//已选长度等于可选数，不存在死循环，即为答案
	if len(ans) == numCourses {
		return ans
	}

	return nil
}
