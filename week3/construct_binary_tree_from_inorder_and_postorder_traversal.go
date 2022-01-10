//106. 从中序与后序遍历序列构造二叉树
//https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package week3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	//递归函数定义
	//iStart与iEnd表示中序序列的起始与结束位置
	//pStart与pEnd表示后序序列的起始与结束位置
	//上述起始与结束位置均为闭区间
	var recurBuild func(iStart, iEnd, pStart, pEnd int) *TreeNode
	recurBuild = func(iStart, iEnd, pStart, pEnd int) *TreeNode {
		//递归边界，已没有结点
		if iStart > iEnd {
			return nil
		}

		//后序序列的最后一个元素为根结点
		root := &TreeNode{Val: postorder[pEnd]}

		//在中序序列里查找根结点的位置
		mPos := iStart
		for inorder[mPos] != root.Val {
			mPos++
		}

		//左子树结点数为mPos-iStart
		//后序序列里的左子树结点范围为pStart ~ pStart+(mPos-iStart)-1
		//后序序列里的右子树结点范围为pStart+(mPos-iStart) ~ pEnd-1
		root.Left = recurBuild(iStart, mPos-1, pStart, pStart+mPos-iStart-1)
		root.Right = recurBuild(mPos+1, iEnd, pStart+mPos-iStart, pEnd-1)

		return root
	}

	return recurBuild(0, len(inorder)-1, 0, len(postorder)-1)
}
