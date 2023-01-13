package action

import (
	"go_learning/src/design/create"
	"sync"
	"testing"
)

const parCount = 100

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
