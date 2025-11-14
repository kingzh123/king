package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter  int
	counter2 int
	sm       sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	wg.Done()
	sm.Lock() // 进入临界区
	counter++
	sm.Unlock() // 离开临界区
}

// 非安全协程
func increment2() {
	counter2++
}

// RunSyncMutex 同步互斥锁
func RunSyncMutex() {
	start := time.Now()
	var wg sync.WaitGroup //等待一组goroutine完成
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	end := time.Since(start) // 计算时间差。或者使用 time.sub()
	fmt.Println(end)
	fmt.Println("counter is ", counter)
}

// RunNotSyncMutex 非安全的协程，实际开发不能这样使用，会导致数据错乱
func RunNotSyncMutex() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		go increment2()
	}
	end := time.Since(start)
	fmt.Println(end)
	fmt.Printf("counter: %d", end) // 最后生成的数据是错误的
}
