package session

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  int
}

//查询用户的方法
func queryUser(uid int) (User, error) {
	user := map[int]User{
		0: User{"zs", 20},
		1: User{"ls", 21},
		2: User{"ww", 22},
	}
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("%d err ", uid)

}

func TestRPC(t *testing.T) {
	gob.Register(User{})
	addr := "127.0.0.1:8000"
	//创建服务端
	srv := NewServer(addr)
	//注册方法
	srv.Register("queryUser", queryUser)
	//服务端等待调用
	go srv.Run()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("err")
	}
	//客户端对象
	cli := NewClient(conn)
	var query func(int) (User, error)
	cli.callRPC("queryUser", &query)
	u, err := query(1)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(u)

}
