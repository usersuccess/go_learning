package binary_tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func preorderTraverseOld(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	//前序遍历的结果
	res = append(res, root.Val)
	//左子树结果放入
	res = append(res, preorderTraverseOld(root.Left)...)
	//右子树结果放入
	res = append(res, preorderTraverseOld(root.Right)...)
	return res
}

//前序遍历
var res []int

func preorderTraverse(root *TreeNode) []int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	//前序位置
	res = append(res, root.Val)
	traverse(root.Left)
	traverse(root.Right)

}

//打印每一个节点的层数
func traversePrintLevel(root *TreeNode, level int) {
	if root == nil {
		return
	}
	//前序位置
	fmt.Printf("节点 %s 在第 %d 层", root, level)
	traversePrintLevel(root.Left, level+1)
	traversePrintLevel(root.Right, level+1)
}

//最大深度:深度优先算法 DFS
func traverseMaxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//前序位置
	return max(traverseMaxDepthDFS(root.Left), traverseMaxDepthDFS(root.Right)) + 1
}
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

//最大深度:广度优先算法 BFS

//最大深度遍历获取 回溯算法
var depth, resDepth int

func traverseMaxDepth(root *TreeNode) {
	if root == nil {
		return
	}
	//前序位置
	depth++
	if root.Left == nil && root.Right == nil {
		//叶子节点最大深度
		resDepth = max(resDepth, depth)
	}
	traverseMaxDepth(root.Left)
	traverseMaxDepth(root.Right)
	//后序位置
	depth--
}
