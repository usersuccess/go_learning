package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0),cap(s0))
	var s = make([]int,3,4)
	t.Log(len(s),cap(s))

}

/**
自增长的代价,存储空间复制
 */
func TestSliceGrowing(t *testing.T)  {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s,i)
		t.Log(len(s),cap(s))
	}
}

func TestSliceComparing(t *testing.T)  {
	/*s1 := []int{}
	s2 := []int{}*/
	/*if s1==s2{//slice只能nil比较

	}*/
}
