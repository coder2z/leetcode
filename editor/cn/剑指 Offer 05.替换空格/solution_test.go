package leetcode

import (
	"strings"
	"testing"
)

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
// 示例 1：
//
// 输入：s = "We are happy."
//输出："We%20are%20happy."
//
//
//
// 限制：
//
// 0 <= s 的长度 <= 10000
//
// Related Topics 字符串 👍 402 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	var rsp string
	// 直接操作？
	rsp = strings.ReplaceAll(s, " ", "%20")

	// 自己实现

	// 直接冲
	//sr := []rune(s)
	//var rspsr []rune
	//for _, r := range sr {
	//	if r == ' ' {
	//		rspsr = append(rspsr, []rune("%20")...)
	//	} else {
	//		rspsr = append(rspsr, r)
	//	}
	//}
	//rsp = string(rspsr)

	// 用string build
	sr := []rune(s)
	var rspsr strings.Builder
	for _, r := range sr {
		if r == ' ' {
			rspsr.WriteString("%20")
		} else {
			rspsr.WriteRune(r)
		}
	}
	rsp = rspsr.String()

	return rsp
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTiHuanKongGeLcof(t *testing.T) {
	t.Log(replaceSpace("We are happy."))
}
