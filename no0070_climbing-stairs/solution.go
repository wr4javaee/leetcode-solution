package no0070_climbing_stairs

func climbStairs(n int) int {
	// 边界值校验
	if n <=0 {
		return 0
	}

	// 动态规划
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	// F(N) = F(N-2) + F(N-1)
	// 假设a代指F(N-2), b代指F(N-1)
	a := 1
	b := 2
	res := 0
	for i := 3; i <= n; i++ {
		// F(N) = F(N-2) + F(N-1)
		res = a + b
		// 随着i的遍历, a、b向右移动
		a = b
		b = res
	}

	return res
}
