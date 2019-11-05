/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (35.90%)
 * Likes:    1499
 * Dislikes: 0
 * Total Accepted:    88.3K
 * Total Submissions: 245.9K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */
package leetcode

import "sort"

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	midM := (m - 1) / 2
	midN := (n - 1) / 2

	if m == 0 {
		if n%2 == 1 {
			return float64(nums2[midN])
		} else {
			return float64(nums2[midN]+nums2[midN+1]) / 2
		}
	} else if m == 1 || m == 2 {
		// n小于3的情况下，取nums2所有元素和nums1的元素进行排序
		if n < 3 {
			for i := 0; i < n; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else if n%2 == 1 {
			// n大于2且为奇数的情况下，取nums2中间3位和nums1的元素进行排序
			for i := midN - 1; i < midN+2; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else {
			// 其他情况下，取nums2的中间4位和nums1的元素进行排序
			for i := midN - 1; i < midN+3; i++ {
				nums1 = append(nums1, nums2[i])
			}
		}
		sort.Ints(nums1)
		m = len(nums1)
		midM = (m - 1) / 2

		if m%2 == 1 {
			return float64(nums1[midM])
		} else {
			return float64(nums1[midM]+nums1[midM+1]) / 2
		}
	}

	// n为奇数时，midNP==midN。n为偶数时，midNP==midN+1。
	midNP := midN
	if n%2 == 0 {
		midNP++
	}

	if nums1[midM] == nums2[midNP] {
		return float64(nums1[midM])
	}
	if nums1[midM] < nums2[midNP] {
		// 取出nums1数组midM下标及之后的元素，和nums2数组 [0, n-midM) 的元素
		return findMedianSortedArrays(nums1[midM:], nums2[:n-midM])
	} else {
		// 取出nums2数组midM下标及之后的元素，和nums1数组 [0, m-midM) 的元素
		return findMedianSortedArrays(nums2[midM:], nums1[:m-midM])
	}
}

// @lc code=end
