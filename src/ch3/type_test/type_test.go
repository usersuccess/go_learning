package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int  =1
	var b int64
	b = int64(a) //支持显式类型转化
	t.Log(a,b)
	var c MyInt
	c = MyInt(b)
	t.Log(c)
}
