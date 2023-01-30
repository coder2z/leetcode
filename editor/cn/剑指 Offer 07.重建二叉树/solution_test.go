package leetcode

import (
	"fmt"
	"testing"
)

//输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//
// 示例 1:
//
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 5000
//
//
//
// 注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-
//preorder-and-inorder-traversal/
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 968 👎 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder 前序
// preorder = [3,9,20,15,7]
// inorder 中序
// inorder = [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	ints := searchIndex(inorder, preorder[0])

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:ints+1], inorder[:ints]),
		Right: buildTree(preorder[ints+1:], inorder[ints+1:]),
	}
}

func searchIndex(data []int, to int) int {
	for i, v := range data {
		if v == to {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

// 前序遍历
func Handle(t *TreeNode) {
	// 前面遍历
	fmt.Println(t.Val)
	if t.Left != nil {
		Handle(t.Left)
	}
	// 中序遍历
	//fmt.Println(t.Val)
	if t.Right != nil {
		Handle(t.Right)
	}
	// 后序遍历
	//fmt.Println(t.Val)
}

func TestZhongJianErChaShuLcof(t *testing.T) {
	//v4 := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 9,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val: 15,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//		},
	//	},
	//}
	//Handle(v4)
	var preorder = []int{3, 9, 20, 15, 7}
	var inorder = []int{9, 3, 15, 20, 7}
	// 下标
	ints := searchIndex(inorder, preorder[0])
	t.Log(ints)
	t.Log(preorder[1:ints+1], inorder[:ints])
	t.Log(preorder[ints+1:], inorder[ints+1:])
}
