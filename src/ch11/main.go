package main

import (
	"fmt"
)

var testList = make([]int, 10)

func QuickSort(unSortList []int) {
	QuickHelper(unSortList, 0, len(unSortList)-1)
}

func QuickHelper(aList []int, first, last int) {
	if first < last {
		splitPoint := partition(aList, first, last)
		QuickHelper(aList, first, splitPoint-1)
		QuickHelper(aList, splitPoint+1, last)
	}
}

//获取中值
func partition(aList []int, first, last int) int {
	pivotValue := aList[first]
	leftMark := first + 1
	rightMark := last
	done := false
	for !done {
		for leftMark <= rightMark && aList[leftMark] < pivotValue {
			leftMark++
		}
		for rightMark >= leftMark && aList[rightMark] > pivotValue {
			rightMark--
		}
		if rightMark < leftMark {
			done = true
		} else {
			aList[leftMark], aList[rightMark] = aList[rightMark], aList[leftMark]
			/*temp := aList[leftMark]
			aList[leftMark] = aList[rightMark]
			aList[rightMark] = temp*/
		}
	}
	aList[first], aList[rightMark] = aList[rightMark], aList[first]
	/*temp := aList[first]
	aList[first] = aList[rightMark]
	aList[rightMark] = temp*/
	return rightMark
}

/*func main() {
	var v *int
	v = new(int)
	*v = 8
	fmt.Println(*v)
	testList = []int{9, 7, 2, 323, 43, 5, 6, 66, 45}
	QuickSort(testList)
	fmt.Println(testList) //[2 5 6 7 9 43 45 66 323]
}*/

func QuickSelectionSort(aList []int, first, last int) {

}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
	dog{}.say()
}

// Sayer 接口
type Sayer interface {
	say()
}
type dog struct{}

type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

//冒泡排序,O(n^2),稳定
func BubbleSort(unSortList []int) {
	length := len(unSortList)
	for i := 0; i < length-1; i++ {
		for j := 0; j <= length-1-i; j++ {
			if unSortList[j] > unSortList[j+1] {
				unSortList[j], unSortList[j+1] = unSortList[j+1], unSortList[j]
			}
		}
	}

}

//选择排序 O(n^2) 不稳定
func SelectionSort(unSortList []int) {
	length := len(unSortList)
	var minIndex int
	for i := 0; i < length-1; i++ {
		minIndex = i
		for j := i + 1; j < length; j++ {
			if unSortList[minIndex] > unSortList[j] {
				minIndex = j
			}
		}
		unSortList[minIndex], unSortList[i] = unSortList[i], unSortList[minIndex]
	}
}
