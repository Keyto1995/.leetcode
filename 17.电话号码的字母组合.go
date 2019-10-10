/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (50.90%)
 * Likes:    462
 * Dislikes: 0
 * Total Accepted:    49.6K
 * Total Submissions: 97K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var kmap [][]string
	kmap = append(kmap, []string{"a", "b", "c"})
	kmap = append(kmap, []string{"d", "e", "f"})
	kmap = append(kmap, []string{"g", "h", "i"})
	kmap = append(kmap, []string{"j", "k", "l"})
	kmap = append(kmap, []string{"m", "n", "o"})
	kmap = append(kmap, []string{"p", "q", "r", "s"})
	kmap = append(kmap, []string{"t", "u", "v"})
	kmap = append(kmap, []string{"w", "x", "y", "z"})

	result := []string{""}
	for _, value := range digits {
		keys := kmap[value-50]
		var tmpResult []string
		count := len(result)
		for i := 0; i < count; i++ {
			for _, key := range keys {
				tmpResult = append(tmpResult, result[i]+key)
			}
		}
		result = tmpResult
	}
	return result
}

// @lc code=end
