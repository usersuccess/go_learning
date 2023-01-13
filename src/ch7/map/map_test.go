package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 :=map[int]int{1:1,2:4,3:9}
	t.Log(m1)
	m2 :=map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d",len(m2))
	m3 := make(map[int]int,10)
	t.Logf("len m3=%d",len(m3))
}

func TestAccessNotExistingKey(t *testing.T)  {
	m1 :=map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v,ok :=m1[3];ok{
		t.Logf("Key 3's value is %d",v )
	}
}

func TestInitArray(t *testing.T)  {
	a :=map[string]int{"hello":1,"nihao":2}
	t.Log(a)
}