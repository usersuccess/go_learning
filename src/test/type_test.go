package test

import (
	"go_learning/src/design/create"
	"testing"
)

func TestType(t *testing.T) {
	api := create.NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi,Tom" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := create.NewAPI(2)
	s := api.Say("Tom")
	if s != "hello,Tom" {
		t.Fatal("Type1 test fail")
	}
}
