package session

import (
	"net"
	"reflect"
)

//客户端
type Client struct {
	conn net.Conn
}

//构建方法
func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

//RPC客户端
func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	//反射初始化fPtr函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	f := func(args []reflect.Value) []reflect.Value {
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		//连接
		cliSession := NewSession(c.conn)
		//编码数据
		reqRPC := RPCData{
			Name: rpcName,
			Args: inArgs,
		}
		b, err := encode(reqRPC)
		if err != nil {
			panic(err)
		}
		//写数据
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}
		//服务端发过来的返回值,读取并解析
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		//解码
		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}

		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			if arg == nil {
				// reflect.Zero()会返回类型的零值的value
				// .out()会返回函数输出的参数类型
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
