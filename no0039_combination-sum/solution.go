package no0039_combination_sum

import (
	"sort"
)

// https://www.cnblogs.com/jimmycheng/p/7198156.html?utm_source=itdadao&utm_medium=referral
// https://studygolang.com/articles/9876
// 执行用时: 4 ms, 在Combination Sum的Go提交中击败了100.00% 的用户
func combinationSum(candidates []int, target int) [][]int {
	// 初始化返回值
	resultArrs := make([][]int, 0, 20)

	// 将candidates从小到大排列
	sort.Ints(candidates)

	// 初始化list, 用于存放一个合法的结果集
	resultArrTmp := make([]int, 0, 20)
	findResult(&resultArrs, &resultArrTmp, &candidates, target, 0)
	return resultArrs
}

func findResult(resultArrs *[][]int, resultArrTmp *[]int, candidates *[]int, target int, candidatesInx int) bool {
	if target == 0 {
		// 找到可解项
		newSlice := make([]int, len(*resultArrTmp))
		copy(newSlice, *resultArrTmp)
		*resultArrs = append(*resultArrs, newSlice)
		// 下一次循环break
		return true
	}

	// 遍历数据源
	for i := candidatesInx; i < len(*candidates); i ++ {
		num := (*candidates)[i]
		*resultArrTmp = append(*resultArrTmp, num)

		if num > target {
			*resultArrTmp = (*resultArrTmp)[:len(*resultArrTmp) - 1]
			break
		}

		// 开启递归调用
		isBreak := findResult(resultArrs, resultArrTmp, candidates, target - num, i)
		*resultArrTmp = (*resultArrTmp)[:len(*resultArrTmp) - 1]

		if isBreak {
			break
		}
	}

	// 下一次循环continue
	return false
}
