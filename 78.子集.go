/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (74.68%)
 * Likes:    361
 * Dislikes: 0
 * Total Accepted:    39.6K
 * Total Submissions: 52.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */
package leetcode

// @lc code=start
func subsets(nums []int) [][]int {
	count := len(nums)
	var flags []int
	for i := 0; i < count; i++ {
		flags = append(flags, 1<<uint(i))
	}
	var result [][]int
	count = 1 << uint(len(nums))

	for i := 0; i < count; i++ {
		tmp := []int{}
		for j, value := range flags {
			if (i & value) != 0 {
				tmp = append(tmp, nums[j])
			}
		}
		result = append(result, tmp)
	}
	return result
}

// @lc code=end
