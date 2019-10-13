/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (23.86%)
 * Likes:    1428
 * Dislikes: 0
 * Total Accepted:    103.3K
 * Total Submissions: 428.9K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */
package main

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	size := len(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		for l, r := i+1, size-1; l < r; {
			if nums[l]+nums[r]+nums[i] < 0 {
				l++
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}
			}
		}
	}
	return result
}

// @lc code=end
