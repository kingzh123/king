package goroutine

import (
	"fmt"
	"time"
)

type User struct {
	Name string
}

func useChannelErrDemo() {
	// 这是一个错误的示例，因为创建无缓存通道，必须选有接收者，否则管道可能会被堵塞。
	c := make(chan int)
	c <- 10
	cc := make(chan []User)
	fmt.Println(cc)
}

func useChanelDemo(channel chan int) {
	// ch <- 10 将一个值发送到管道
	// a := <- ch 从管道接收一个值
	// <-ch 从管道接收一个值，且忽略此值
	n := <-channel
	fmt.Println("接收数据：", n)
}

func cacheChannelClose() {
	c := make(chan int)
	cc := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // 关闭队列 应在发送方完成，接收方不处理关闭
		for i := 0; i < 10; i++ {
			cc <- i
		}
		close(cc)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println("c: ", data)
		} else {
			break
		}
	}
	for data := range cc {
		fmt.Println("cc: ", data)
	}
}

func channelSelectDemo() {
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println("val: ", val)
	}()
	time.Sleep(2 * time.Second)
	select {
	case ch <- 10:
		fmt.Println("send")
	case val := <-ch:
		fmt.Println("Received: ", val)
	default:

		fmt.Println("no activity")
	}

}

func channelSelectDemo2() {
	ch1 := make(chan int)
	ch1Bool := false
	ch2 := make(chan int)
	ch2Bool := false

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 100; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for !ch1Bool || !ch2Bool {
		select {
		case data, ok := <-ch1:
			if ok {
				fmt.Println("ch1: ", data)
			} else {
				ch1Bool = true
			}
		case data, ok := <-ch2:
			if ok {
				fmt.Println("ch2: ", data)
			} else {
				ch2Bool = true
			}
		}
	}
}

// 只读通道
func readOnlyChannel(ch <-chan int) {
	for c := range ch {
		fmt.Println(c)
	}
}

func writeOnlyChannel(ch chan<- int) {
	for i := 0; i < 30; i++ {
		ch <- i
	}
	close(ch)
}

func channelBase() {
	ch := make(chan int)
	fmt.Printf("值：%v, 地址：%p\n", ch, &ch)

	ch2 := make(chan int)
	go func() {
		ch2 <- 10 // 非缓冲channel必须另起一个 goroutine 才能启动
	}()
	fmt.Println(<-ch2)

	ch3 := make(chan int, 1)
	ch3 <- 20 // 缓存go程可以不必另起一个 goroutine
	fmt.Println("缓存go程", <-ch3)
}

// Channel
// 管道是引用类型
// 管道内的数据不限定类型，但是只能同时存在一种类型
// 管道遵循先入先出的原则
// 管道是线程安全的
// 管道本身有锁
// 在遍历管道的时候必须要被关闭，否则不知道在什么时候终止
func Channel() {
	// 声明通道
	var a chan int
	var aa chan bool
	var aaa chan []int
	// 创建通道
	b := make(chan int)
	bb := make(chan int, 1)
	fmt.Println(a, b, aa, aaa, bb)
	//useChannelErrDemo()
	c := make(chan int)
	go useChanelDemo(c)
	c <- 10
	fmt.Println("over")
	// ---管道的类型--
	//channelBase()
	// --缓存通道--
	//cacheChannelClose()
	// --单向通道--
	//ch := make(chan int, 20)
	//go writeOnlyChannel(ch)
	//readOnlyChannel(ch)
	//--channel select--
	//channelSelectDemo()
	channelSelectDemo2()
}
