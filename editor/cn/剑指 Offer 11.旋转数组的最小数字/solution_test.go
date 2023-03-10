package leetcode

import (
	"testing"
)

//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//
// 给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3
//,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。
//
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2],
//..., a[n-2]] 。
//
//
//
// 示例 1：
//
//
//输入：numbers = [3,4,5,1,2]
//输出：1
//
//
// 示例 2：
//
//
//输入：numbers = [2,2,2,0,1]
//输出：0
//
//
//
//
// 提示：
//
//
// n == numbers.length
// 1 <= n <= 5000
// -5000 <= numbers[i] <= 5000
// numbers 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
//
//
// 注意：本题与主站 154 题相同：https://leetcode-cn.com/problems/find-minimum-in-rotated-
//sorted-array-ii/
//
// Related Topics 数组 二分查找 👍 747 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	//var min int
	// 先暴力查询一次最小的
	//min = math.MaxInt
	//for _, number := range numbers {
	//	if min > number {
	//		min = number
	//	}
	//}
	//return min

	// 排序再拿第一个
	//sort.Ints(numbers)
	//min = numbers[0]
	//return min

	// 直接迭代
	//for i := 1; i < len(numbers); i++ {
	//	if numbers[i] < numbers[i-1] {
	//		return numbers[i]
	//	}
	//}
	//return numbers[0]

	// 二分查找
	// 左右两个下标
	var x, y = 0, len(numbers) - 1
	for x < y {
		var m = (x + y) / 2
		index := numbers[m]
		if index > numbers[y] {
			x = m + 1
		} else if index < numbers[y] {
			y = m
		} else {
			y = y - 1
		}
	}

	return numbers[x]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestXuanZhuanShuZuDeZuiXiaoShuZiLcof(t *testing.T) {
	t.Log(minArray([]int{3, 4, 5, 2}))
}
