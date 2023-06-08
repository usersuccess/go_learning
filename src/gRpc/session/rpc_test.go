package session

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSessionReadWriter(t *testing.T) {
	addr := "127.0.0.1:7777"
	my_data := "hello"
	wg := sync.WaitGroup{}
	wg.Add(2)
	//写数据协程
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, _ := lis.Accept()
		s := Session{
			conn: conn,
		}
		err = s.Write([]byte(my_data))
		if err != nil {
			t.Fatal(err)
		}
	}()

	//读取数据的协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := Session{
			conn: conn,
		}
		data, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != my_data {
			t.Fatal(err)
		}
		fmt.Printf(string(data))
	}()
	wg.Wait()

}
