package constant_test

import (
	"math"
	"sort"
	"testing"
)

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

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	//return strings.Join(word1, "") == strings.Join(word2, "")
	//指针比较
	var p1, p2, i, j int
	for p1 < len(word1) && p2 < len(word2) {
		if word1[p1][i] != word2[p2][j] {
			return false
		}
		i++
		if i == len(word1[p1]) {
			p1++
			i = 0
		}

		j++
		if j == len(word2[p2]) {
			p2++
			j = 0
		}

	}

	return p1 == len(word1) && p2 == len(word2)

}

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{i, p}
		}
		hashTable[x] = i
	}
	return nil
}

func threeSum(nums []int) [][]int {
	//1⃣️先考虑特殊情况
	if nums == nil || len(nums) < 3 {
		return nil
	}
	//2⃣️排序处理
	three := make([][]int, 0)
	sort.Ints(nums)
	//3⃣️循环首个值,并去重,另外2个值考虑用双指针
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				three = append(three, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return three
}

func threeSumClosest(nums []int, target int) int {
	//特殊判断
	if len(nums) < 3 {
		return 0
	}
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	best := math.MaxInt32
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	//排序后得到该数值,可以左右指针移动,确定最接近的
	for i := 0; i < len(nums); i++ {
		//重复数据不需要处理
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			update(sum)
			if sum == target {
				return target
				//数据处理对比

			} else if sum > target {
				k--
			} else {
				j++
			}
		}

	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func fourSum(nums []int, target int) [][]int {
	var resp [][]int
	n := len(nums)
	if n < 4 || nums == nil {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		//确定了第一个数字+最后三个数字小于target
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			//确定了前两个数字➕最后两个数字求和小于target
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					resp = append(resp, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return resp

}

func findMaxConsecutiveOnes(nums []int) int {
	var cur, best int

	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			cur++
		} else {
			if cur > best {
				best = cur
			}
			cur = 0
		}
	}
	if cur > best {
		best = cur
	}
	return best
}
func maximumProduct(nums []int) int {

	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[n]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])

}

//如果全是非正数，则最大的三个数相乘同样也为最大乘积。
//
//如果数组中有正数有负数，则最大乘积既可能是三个最大正数的乘积，也可能是两个最小负数（即绝对值最大）与最大正数的乘积。
//
//综上，我们在给数组排序后，分别求出三个最大正数的乘积，以及两个最小负数与最大正数的乘积，二者之间的最大值即为所求答案。
//
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//[2,3,3,4,6,5]
//纯数学解题：
//
//sum(nums) - sum(set(nums)) = 重复的数字
//(1 + len(nums)) * len(nums) // 2 - sum(set(nums)) = 丢失的数字
func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	sort.Ints(nums)
	pre := 0
	for _, v := range nums {
		if v == pre {
			ans[0] = v
		} else if v-pre > 1 {
			ans[1] = pre + 1
		}
		pre = v
	}
	n := len(nums)
	if nums[n-1] != n {
		ans[1] = n
	}
	return ans

}
