//538. 把二叉搜索树转换为累加树
//https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
//二叉搜索树的中序遍历会得到一个升序数组
//对数组从后往前遍历，修改每个元素的值为当前值与其后一元素的值之和，即为答案
//时间复杂度为 O(n)，需要遍历每个结点元素
//空间复杂度为 O(n)，需要一个额外数组存储升序结点列表

package week4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	head := root

	//中序遍历得到一个升序数组
	nodeList := []*TreeNode{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		nodeList = append(nodeList, root)
		dfs(root.Right)
	}

	dfs(root)

	//从数组的倒数第二位置往前遍历，修改其值为当前值与其后一元素的值之和
	for i := len(nodeList) - 2; i >= 0; i-- {
		nodeList[i].Val += nodeList[i+1].Val
	}

	return head
}
