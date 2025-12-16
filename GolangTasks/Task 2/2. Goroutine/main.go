package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 1：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

func odds(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func evens(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// 题目 2：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
type Missions struct {
	Name string
	Fn   func()
	Dur  time.Duration
}

func Records(mi *Missions, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	mi.Fn()

	mi.Dur = time.Since(start)

}

func main() {

	var wg1 sync.WaitGroup
	wg1.Add(2)

	// Question 1:
	fmt.Println("奇数")
	go odds(&wg1)

	fmt.Println("偶数")
	go evens(&wg1)

	wg1.Wait()
	fmt.Println("Question 1: Finished")

	// Question 2:
	tasks := []Missions{
		{
			Name: "task-A (sleep 50ms)",
			Fn: func() {
				time.Sleep(3 * time.Second)
				fmt.Println("睡了3秒")
			},
		},

		{
			Name: "task-B (sleep 120ms)",
			Fn: func() {
				time.Sleep(2 * time.Second)
				fmt.Println("睡了2秒")
			},
		},

		{
			Name: "task-C (compute loop)",
			Fn: func() {
				// 模拟 CPU 任务（简单循环求和）
				sum := 0
				for i := 0; i < 200000; i++ {
					sum += i
				}
				_ = sum // 防止编译器优化掉
			},
		},
	}
	fmt.Println(tasks)

	var wg2 sync.WaitGroup
	wg2.Add(len(tasks))

	for i := 0; i < len(tasks); i++ {

		go Records(&tasks[i], &wg2)
	}

	wg2.Wait()

	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i].Dur)
	}
	fmt.Println("Question 2: Finished")

}
