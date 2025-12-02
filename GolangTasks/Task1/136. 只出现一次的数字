// 136. 只出现一次的数字
// Brute Force
func singleNumber(nums []int) int {
    for i := 0; i < len(nums); i++ {
        count := 0

        // 内层循环：统计 nums[i] 在整个数组里出现了几次
        for j := 0; j < len(nums); j++ {
            if nums[j] == nums[i] {
                count++
            }
        }

        // 如果只出现了一次，说明它就是题目要求的那个数字
        if count == 1 {
            return nums[i]
        }
    }

    return 0
}
