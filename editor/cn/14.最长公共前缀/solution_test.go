package leetcode

import (
	"testing"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字典树 字符串 👍 2911 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var rst = strs[0]

	for _, v := range strs {
		rst = lsp(rst, v)
	}

	return rst
}

// 去除两个支付穿的公共部分
func lsp(s1 string, s2 string) string {
	count := min(len(s1), len(s2))
	index := 0
	for index < count && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

// min
func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestCommonPrefix(t *testing.T) {

}
