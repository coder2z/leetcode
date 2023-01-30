package leetcode

import (
	"testing"
)

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
// 示例 1：
//
// 输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//
// 限制：
//
// 0 <= 链表长度 <= 10000
//
// Related Topics 栈 递归 链表 双指针 👍 368 👎 0

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	return append(reversePrint(head.Next), head.Val)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCongWeiDaoTouDaYinLianBiaoLcof(t *testing.T) {
	v1 := ListNode{
		Val:  2,
		Next: nil,
	}
	v2 := ListNode{
		Val:  3,
		Next: &v1,
	}
	v3 := ListNode{
		Val:  1,
		Next: &v2,
	}
	t.Log(reversePrint(&v3))
}
