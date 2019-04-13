package main

import (
	"runtime"
	"sync"
	"fmt"
	"sync/atomic"
	"time"
)

func exemple() {
	// 1. 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	
	// 2. 设定等待器，类比 Java CountDownLatch
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	fmt.Println("== start ==")
	// 3. 创建一个 goroutine
	go func() {
		defer waitGroup.Done() // CountDownLatch countDown()

		// 打印三遍字母表
		for count := 0; count < 3; count ++ {
			for char := 'a'; char < 'a'+26; char ++ {
				fmt.Printf("%c", char)
			}
		}	
	}()

	// 4. 创建第二个 goroutine 
	go func() {
		defer waitGroup.Done()

		for count := 0; count < 3; count ++ {
			for char := 'A'; char < 'A'+26; char ++ {
				fmt.Printf("%c", char)
			}
		}	
	}()

	// 5. 阻塞 main goroutine
	waitGroup.Wait()
	fmt.Println("== end ==")
}

// 竞争状态

var (
	// 两个 goroutine 同时操作变量
	counter int
	// waitGroup sync.WaitGroup
)

func intCount(s int) {
	defer waitGroup.Done()
	for count := 0; count < 2; count ++ {
		value := counter
		// 当前的 goroutine 主动让出字眼，从线程退出，并放回到队列，
		//让其他 goroutine 继续执行
		runtime.Gosched()
		value ++
		counter = value
		fmt.Println(s,"value", value)
	}
}

var (
	shutdown int64
	waitGroup sync.WaitGroup
)

func doWork(name string) {
	defer waitGroup.Done()
	for {
		time.Sleep(250 * time.Millisecond)
		// 记载关机标志
		fmt.Println(name)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("shutDown, ", name)
			break
		}
	}
}

var (
	// 锁，定义一段临界区
	lock sync.Mutex
)

func intCounts(i int) {
	defer waitGroup.Done()
	for count := 0; count < 2; count ++ {
		lock.Lock()
		// lock 与 unlock 之间的代码都属于临界区， {} 可以省略
		{
			value := counter
			runtime.Gosched()
			value ++
			counter = value
		}
		lock.Unlock()
		// 释放锁
	}
}




func main() {
	// exemple()
	runtime.GOMAXPROCS(1)
	waitGroup.Add(2)

	// go intCount(1)
	// go intCount(2)

	// go doWork("A")
	// go doWork("B")

	go intCounts(1)
	go intCounts(2)

	// time.Sleep(1000 * time.Millisecond)
	// atomic.StoreInt64(&shutdown, 1)
	waitGroup.Wait()
	fmt.Println("counter", counter)
}