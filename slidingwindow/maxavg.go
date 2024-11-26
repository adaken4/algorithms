// Given an array of integers nums and an integer k, find the maximum average value of any contiguous subarray of size k.

// Example:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 12.75

// Constraints:
// - n == nums.length
// - 1 <= k <= n <= 10^5
// - -10^4 <= nums[i] <= 10^4

package algos

func maxAvg(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	result := float64(sum) / float64(k)

	for i := k; i < len(nums); i++ {
		// sliding window technique
		sum = sum + nums[i] - nums[i-k]
		windowAvg := float64(sum) / float64(k)
		if windowAvg > result {
			result = windowAvg
		}
	}

	return result
}

// Given an array of positive integers nums and a positive integer target, 
// return the minimal length of a contiguous subarray whose sum is greater than or equal to target. 
// If there is no such subarray, return 0.

// Example:
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Constraints:
// - 1 <= target <= 10^9
// - 1 <= nums.length <= 10^5
// - 1 <= nums[i] <= 10^4

func minLen(arr []int, t int) int {
	result := len(arr) + 1
	sum := 0
	left := 0
	for right := 0; right < len(arr); right++ {
		sum += arr[right]
		for sum >= t {
			sum -= arr[left]
			result = min(result, right-left+1)
			left++
		}
	}
	if result > len(arr) {
		return 0
	}
	return result
}

// Given a string s and an integer k, find the length of the longest substring 
// that contains at most k distinct characters.

// Example 1:
// Input: s = "eceba", k = 2
// Output: 3
// Explanation: The longest substring with at most 2 distinct characters is "ece"

// Example 2:
// Input: s = "aa", k = 1
// Output: 2
// Explanation: The substring "aa" contains only one distinct character

// Constraints:
// - 1 <= s.length <= 10^5
// - 0 <= k <= 50
// - s consists of uppercase and lowercase English letters

func longestSubstrLen(s string, k int) int {
	if k == 0 {
		return 0
	}
	sRunes := []rune(s)
	left := 0
	maxLength := 0
	charCount := map[rune]int{}
	for right := 0; right < len(sRunes); right++ {
		charCount[sRunes[right]]++

		for len(charCount) > k {
			charCount[sRunes[left]]--
			if charCount[sRunes[left]] == 0 {
				delete(charCount, sRunes[left])
			}
			left++
		}
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
