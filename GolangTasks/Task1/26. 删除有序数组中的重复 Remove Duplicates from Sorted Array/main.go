func removeDuplicates(nums []int) int {

	// k指向“下一个要放入的新元素”的位置，同时也代表当前不重复数组的长度
	var k = 1

	for i := 1; i < len(nums); i++ { //从第二个元素开始遍历

		// fmt.Println(nums[i], nums[k-1])
		// fmt.Println("i & k & k-1 =", i, k, k-1)

		if nums[i] != nums[k-1] { // 如果遇到一个新值（与上一个保留下来的值不同）

			// fmt.Println("-------------")
			// fmt.Println(nums[i], nums[k-1])

			nums[k] = nums[i] // 把这个新值放到数组前面
			k++         // 不重复长度+1
			// fmt.Println("----------------------------")
		}
	}
	return k
}
