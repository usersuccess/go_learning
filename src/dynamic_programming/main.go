package dynamic_programming

import "math"

//动态规划 dynamic programming
//最优化方法,求最值，核心是穷举【掌握递归】，状态转移方程去穷举
//算法的最优子结构，通过子问题的最值找到原问题的最值
//存在重叠子问题，穷举效率低,加上【备忘录】or【dp table】
//三要素：重叠子问题 最优子结构 状态转移方程
//思维：明确base case ->明确【状态】->明确【选择】->定义dp数组/函数

/*func fib(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}*/

//备忘录
/*func fib(N int) int {
	memo := make([]int, N+1)
	return dpMemo(memo, N)
}
func dpMemo(memo []int, n int) int {
	if n == 1 || n == 0 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = dpMemo(memo, n-1) + dpMemo(memo, n-2)
	return memo[n]
}
*/

//dp数组迭代递推
/*func fib(N int) int {
	if N == 0 {
		return 0
	}
	dp := make([]int, N+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}*/

/*func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	//滚动更新,空间复杂度1,时间复杂度O(n)
	dp_i_1, dp_i_2 := 1, 0
	for i := 2; i <= n; i++ {
		dp_i := dp_i_1 + dp_i_2
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i_1
}*/

//找硬币问题
func coinChange(coins []int, amount int) int {
	//base
	return dp(coins, amount)

}
func dp(coins []int, amount int) int {
	//base
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := dp(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}