package week3

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	//分治递归，将链表数组对半平分为前后两个分组
	//重复分组，直到分组内链表数<=1，到达递归边界
	//合并前后两个分组的链表，返回一个更大的链表，重复此过程，最终得到需要的结果
	var recurMergeFunc func(l, r int) *ListNode
	recurMergeFunc = func(l, r int) *ListNode {
		//递归边界
		if l == r {
			return lists[l]
		}
		if l > r {
			return nil
		}

		mid := (l + r) / 2
		return mergeTwoLists(recurMergeFunc(l, mid), recurMergeFunc(mid+1, r))
	}

	return recurMergeFunc(0, len(lists)-1)
}

//合并两个升序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{} //保护结点
	tailPtr := dummyNode     //结果链表当前的尾节点指针

	//遍历l1和l2，比较两个结点值的大小，选较小者作为新节点加入到结果链表
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tailPtr.Next = l1
			l1 = l1.Next
		} else {
			tailPtr.Next = l2
			l2 = l2.Next
		}
		tailPtr = tailPtr.Next
	}

	//l1比l2长，直接串联l1的剩余结点
	if l1 != nil {
		tailPtr.Next = l1
	}

	//l2比l1长，直接串联l2的剩余结点
	if l2 != nil {
		tailPtr.Next = l2
	}

	return dummyNode.Next
}
