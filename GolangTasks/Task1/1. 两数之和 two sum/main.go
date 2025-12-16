func twoSum(nums []int, target int) []int {
	// 记录：值 -> 下标
	var m = make(map[int]int)
	fmt.Println(m)

	for i, v := range nums {
		// complement 是“还差多少能到 target”
		gap := target - v
		fmt.Println("gap =", gap)
		fmt.Println("m[gap] =", m[gap])

		// 看看这个“差值”以前有没有见过
		if j, ok := m[gap]; ok {
			fmt.Println("见过，m[gap] =", m[gap])
			// 如果见过，就说明 nums[j] + nums[i] == target
			return []int{j, i}
		}

		// 把当前这个数记下来，方便后面的元素来查
		m[v] = i
	}

	// 题目保证一定有解，其实这里不会执行
	return nil
}
