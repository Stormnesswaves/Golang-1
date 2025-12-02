func longestCommonPrefix(strs []string) string {
	// 假设字符串切片非空，直接返回空字符串

	var firststr = strs[0] // 取出第一个字符串单词

	for i := 0; i < len(firststr); i++ { // i 循环 一个单词 中的每一个字符

		var char = firststr[i] // 取出 需要比较的单个字符


		for j := 1; j < len(strs); j++ { // j：从第二个字符串单词开始 循环，在同一列上比较

			
			

			if i == len(strs[j]) {
        // 当前单词长度已经到头，说明公共前缀到此为止
				return firststr[:i] // 前缀是 [0, i) 
        
			}else if strs[j][i] != char {
        // 这一列的字符不等于 firststr[i]，说明公共前缀到此为止
        return firststr[:i] // 前缀是 [0, i) 
      }
		}
	}

	// 如果循环没有提前 return，说明第一个字符串本身就是公共前缀
	return firststr
}
