package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestContantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111

	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
