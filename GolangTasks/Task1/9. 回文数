func isPalindrome(x int) bool {
  if x < 0 {
		return false
	}else if x == 0 {
		return true
	}

	var rev []int

	for x > 0 {
		d := x % 10 // 取最后一位
		rev = append(rev, d)
		x /= 10 // 去掉最后一位
	}

	// 现在 rev 是倒序的，需要反转过来 // 其实无所谓

	xx := len(rev) / 2
	count := 0
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]

		if rev[i]-rev[j] == 0 {
			count++
		}
	}

	if count == xx {
		return true
	} else {
		return false
	}
}
