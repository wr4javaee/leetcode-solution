package no0040_combination_sum_ii

import (
	"bytes"
	"sort"
	"strconv"
)

// 执行用时: 4 ms, 在Combination Sum II的Go提交中击败了100.00% 的用户
func combinationSum2(candidates []int, target int) [][]int {
	// 初始化返回值
	resultArrs := make([][]int, 0, 20)

	// 将candidates从小到大排列
	sort.Ints(candidates)

	// 初始化list, 用于存放一个合法的结果集
	resultArrTmp := make([]int, 0, 20)
	resultMaps := make(map[string]int, 0)
	findResult(&resultMaps, &resultArrs, &resultArrTmp, &candidates, target, 0)

	return resultArrs
}

func findResult(resultMaps *map[string]int, resultArrs *[][]int, resultArrTmp *[]int, candidates *[]int, target int, candidatesInx int) bool {
	if target == 0 {
		// 找到可解项, 但是要排除重复项
		newSlice := make([]int, len(*resultArrTmp))
		copy(newSlice, *resultArrTmp)

		resultStr := getSliceNumsStr(resultArrTmp)
		if _, ok := (*resultMaps)[resultStr]; ok {
			// 下一次循环break
			return true
		}

		// 如果不存在
		(*resultMaps)[resultStr] = 1
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
		isBreak := findResult(resultMaps, resultArrs, resultArrTmp, candidates, target - num, i + 1)
		*resultArrTmp = (*resultArrTmp)[:len(*resultArrTmp) - 1]

		if isBreak {
			break
		}
	}

	// 下一次循环continue
	return false
}

func getSliceNumsStr(resultArrTmp *[]int) string {
	var b bytes.Buffer
	for _, v := range *resultArrTmp {
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}