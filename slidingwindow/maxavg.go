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