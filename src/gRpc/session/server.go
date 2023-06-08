package session

import (
	"fmt"
	"net"
	"reflect"
)

type Server struct {
	addr  string
	funcs map[string]reflect.Value //map维护关系,反射处理
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

func (s *Server) Register(rpcName string, f interface{}) {
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	// 若map中没值，则将映射加入map，用于调用
	fval := reflect.ValueOf(f)
	s.funcs[rpcName] = fval
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("监听 %s err :%v", s.addr, err)
		return
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		serSession := NewSession(conn)
		b, err := serSession.Read()
		if err != nil {
			return
		}
		//数据解析
		rpcData, err := decode(b)
		if err != nil {
			return
		}
		//根据读到的Name,得到调用的函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Printf("函数 %s 不存在", rpcData.Name)
			return
		}
		//遍历解析客户端传来的参数,放在切片
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}

		//反射调用方法
		out := f.Call(inArgs)
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		//数据编码,返回客户端
		respRPCData := RPCData{rpcData.Name, outArgs}
		bytes, err := encode(respRPCData)
		if err != nil {
			return
		}
		//将服务端编码后的数据写到客户端
		err = serSession.Write(bytes)
		if err != nil {
			return
		}

	}
}
