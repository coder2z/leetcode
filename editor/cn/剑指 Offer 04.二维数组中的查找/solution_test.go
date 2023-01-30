package leetcode

import (
	"testing"
)

//在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。请完成一个高效的函数，输入这样的一个二
//维数组和一个整数，判断数组中是否含有该整数。
//
//
//
// 示例:
//
// 现有矩阵 matrix 如下：
//
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//
//
// 给定 target = 5，返回 true。
//
// 给定 target = 20，返回 false。
//
//
//
// 限制：
//
// 0 <= n <= 1000
//
// 0 <= m <= 1000
//
//
//
// 注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
//
// Related Topics 数组 二分查找 分治 矩阵 👍 862 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 先暴力来一次
	//for _, ints := range matrix {
	//	for _, i := range ints {
	//		if i == target {
	//			return true
	//		}
	//	}
	//}

	// 二分查找
	//for _, ints := range matrix {
	//	// 看是是否在当前数组里面
	//	searchInts := sort.SearchInts(ints, target)
	//	if searchInts >= 0 && searchInts < len(ints) && target == ints[searchInts] {
	//		return true
	//	}
	//	continue
	//}

	// 树查找
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	x, y := 0, len(matrix[0])-1
	for y >= 0 && x <= len(matrix)-1 && x >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
			continue
		}
		if matrix[x][y] < target {
			x++
			continue
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestErWeiShuZuZhongDeChaZhaoLcof(t *testing.T) {
	t.Log(findNumberIn2DArray([][]int{
		{-5},
	}, -10))
	t.Log(findNumberIn2DArray([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 333))
}
