package no0038_count_and_say

import (
	"bytes"
	"strconv"
)

// 执行用时: 8 ms, 在Count and Say的Go提交中击败了56.19% 的用户
func countAndSay(n int) string {
	// 定义首次报数数值
	return count("1", n)
}

func count(countAndSayStr string, n int) string {
	if n == 1 {
		return countAndSayStr
	}
	// 定义map数组, 存放每一次报数结果的依据
	// 其中, map的key存放数字, value为该数字的个数
	resMaps := make([]map[string]int, 0)
	lastNum := ""
	numCount := 0
	for _, v := range countAndSayStr {
		s := string(v)
		if lastNum != "" && s != lastNum {
			resMaps = append(resMaps, map[string]int{lastNum: numCount})
			lastNum = s
			numCount = 0
		}
		numCount ++
		lastNum = s
	}
	if numCount > 0 {
		resMaps = append(resMaps, map[string]int{lastNum: numCount})
	}

	var b bytes.Buffer
	for _, v := range resMaps {
		for numStrTmp, numCountTmp := range v {
			b.WriteString(strconv.Itoa(numCountTmp))
			b.WriteString(numStrTmp)
		}
	}
	return count(b.String(), n - 1)
}