func plusOne(digits []int) []int {
    var x int = len(digits)

    if digits[x-1] < 9 {
		fmt.Println("digits[x-1] < 9")
		digits[x-1]++
		// fmt.Println(digits)
		return digits
	}

    for i := len(digits) - 1; i >= 0; i-- {
		fmt.Println("i =", i)

		if digits[i] != 9 {
			println("digits[i] =", digits[i])
			digits[i] += 1
			println("digits[i] += 1 =", digits[i])

			for j := i + 1; j < x; j++ {
				digits[j] = 0
			}
			// fmt.Println("digits", digits)
			return digits
		}
	}
    digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits


}
