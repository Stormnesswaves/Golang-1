package main

import (
	"fmt"
	"math/rand/v2"
)

// 题目 1：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func poadd(z *int) {

	*z = *z + 10
}

// 题目 2：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func poslice(z *[]int) {

  for i := range *z {
		(*z)[i] *= 2

	}
}

func main() {
	// Question 1:
	var i = 5
	poadd(&i)

	fmt.Println(i)

	
	// Question 2:
	var list = make([]int, 8)
	for i := 0; i < len(list); i++ {
    var num = rand.IntN(100) // [0, 100)  	// 生成 0 到 99 之间的随机数
		fmt.Println("0-99:", num)
		list[i] = num
	}

	poslice(&list)
	fmt.Println(list)
}

