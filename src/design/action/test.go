package main

import (
	"fmt"
	"go_learning/src/design/create"
	"sync"
	"testing"
	"time"
)

const parCount = 100

func main() {
	// 创建管道
	output1 := make(chan string, 3)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
EXIT:
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
			close(ch)
			break EXIT
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func say(s string, c chan string) {
	c <- s
}
func TestSingleton(t *testing.T) {

	ins1 := create.GetInstance()
	ins2 := create.GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instance := [parCount]*create.Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instance[index] = create.GetInstance()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instance[i] != instance[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
