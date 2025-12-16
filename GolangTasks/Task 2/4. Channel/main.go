// 题目 1：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 题目 2：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Question 1:
	var ch11 = make(chan int)
	var exit = make(chan bool)

	fmt.Println(ch11)

	go func(c chan int) {
		for i := 1; i <= 10; i++ {
			c <- i
			fmt.Println("输入数据: ", i)
		}
		close(c)
	}(ch11)

	go func(c chan int, e chan bool) {
		for {
			v, flag := <-c
			if !flag {
				break
			}
			fmt.Println("读到的数据: ", v)
		}
		e <- true
		close(e)
	}(ch11, exit)

	for {
		<-exit
		break
	}
	fmt.Println("Question 1: Finished")

	// Question 2:
	var ch22 = make(chan int, 20)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(c chan<- int) {
		for i := 1; i <= 100; i++ {
			c <- i
			// fmt.Println("输入数据: ", i)
		}
		close(c)
	}(ch22)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		for v := range ch22 {
			fmt.Printf("%d; ", v)
		}
		fmt.Println()

	}(ch22, &wg)

	wg.Wait()
	fmt.Println("Question 2: Finished")

}
