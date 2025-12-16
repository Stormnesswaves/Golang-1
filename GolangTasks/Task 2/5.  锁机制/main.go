// ✅锁机制
// 题目 1：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 2：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Question 1:

	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)

	var x1 = 0
	for z := 0; z < 10; z++ {

		go func(x1 *int, wg *sync.WaitGroup) {
			defer wg.Done()

			for i := 1; i <= 1000; i++ {
				lock.Lock()
				*x1 += i
				lock.Unlock()
			}
		}(&x1, &wg)
	}
	wg.Wait()
	fmt.Println(x1)
	fmt.Println("Question 1: Finished")
	fmt.Println()

	// Question 2:
	var x2 int64 // 计数器
	var lines int = 10
	wg.Add(lines)

	for i := 0; i < lines; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				// atomic.AddInt64：以“不可被打断”的方式完成加1（线程安全）
				atomic.AddInt64(&x2, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("final counter (atomic) = %d (expected %d)\n", x2, int64(lines*1000))
	fmt.Println("Question 2: Finished")
}
