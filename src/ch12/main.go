package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

/*func main() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	out := gen(2, 3)
	for n := range out {
		fmt.Println(n)
		time.Sleep(5 * time.Second) // done thing
		if true {
			break
		}
	}
}*/

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func gen_optimize(done chan struct{}, num ...int) <-chan int {
	out := make(chan int, 2)
	go func() {
		for _, n := range num {
			fmt.Println("gen:", n)
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}
func main() {
	//无缓冲通道
	ch := make(chan int, 1)
	ch <- 12
	go func() {
		fmt.Println("接收成功", <-ch)
	}()
	time.Sleep(1 * time.Second)

	go func() {
		fmt.Println("接收成功", <-ch)
	}()
	fmt.Println("len:", len(ch))
	ch <- 14
	fmt.Println("le2n:", len(ch))

	//time.Sleep(1000)
	os.Exit(1)
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	done := make(chan struct{})
	defer close(done)
	out := gen_optimize(done, 4, 3, 4, 5, 6, 7)
	for n := range out {
		fmt.Println(n)
		time.Sleep(2 * time.Second) // done thing
		/*if true {
			break
		}*/
	}
}

func main_() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
