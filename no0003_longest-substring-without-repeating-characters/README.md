# LeetCode.0003 longest-substring-without-repeating-characters 无重复字符的最长子串
---

## 题目
---

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

示例 2:

```
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

示例 3:

```
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```

## 方法一 暴力法
---
#### 解题思路：
参考
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/
#### 复杂度分析
* 时间复杂度 O( n^3 )

由于LeetCode此题存在超长字符串测试用例，该方法在该测试用例下运行超时。

#### 实现（Golang）
```
package no0003_longest_substring_without_repeating_characters

// 暴力法
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	ans := 0
	for i := 0; i < strLen; i++ {
		for j := i + 1; j <= strLen; j++ {
			subStr := s[i : j]
			isAllUnique, subStrLen := allUnique(subStr)
			if isAllUnique {
				// 取出
				ans = max(ans, subStrLen)
			}
			continue
		}
	}
	return ans
}

// 所有字符串唯一返回true, 否则返回false
func allUnique(str string) (isAllUnique bool, strLen int) {
	strLen = len(str)
	valiMaps := make(map[int32]int32, strLen)
	for _, v := range str {
		if _, ok := valiMaps[v]; ok {
			isAllUnique = false
			return
		}
		valiMaps[v] = v
	}
	isAllUnique = true
	return
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
```
## 方法二 滑块法（借助哈希表）
---
#### 解题思路：
参考
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/

注意官方解答思路以及Java版本实现存在bug，无法通过测试用例pwwkew

#### 复杂度分析
* 时间复杂度 O( n )
* 
执行用时: 16 ms, 在Longest Substring Without Repeating Characters的Go提交中击败了63.92% 的用户
#### 实现（Golang）

```
package no0003_longest_substring_without_repeating_characters

// 滑块法
func lengthOfLongestSubstring02(s string) int {
	strLen := len(s)
	// 返回值
	ans := 0
	// 滑块区间，小曲笺你好:)
	i := 0
	j := i
	// 用于判断字符串是否重复
	valiMap := make(map[uint8]uint8)

	// 滑块从左向右滑动
	for i < strLen && j < strLen {
		jStr := s[j]
		iStr := s[i]
		_, isRepeat := valiMap[jStr]
		if isRepeat && iStr != jStr {
			// 发现有重复字符, 且i的索引字符不等于j的索引字符, i向右滑动, 并删除map中保存的i字符
			i ++
			delete(valiMap, iStr)
		} else if isRepeat && iStr == jStr {
			// 发现有重复字符, 且i的索引字符等于j的索引字符, i, j均向右滑动
			i ++
			j ++
		} else {
			// 非重复字符, j向右边滑动, 并将结果缓存至map中
			valiMap[jStr] = jStr
			j ++
		}

		// 记录滑块长度
		if j == strLen {
			ans = Max(ans, strLen - i)
			continue
		}
		ans = Max(ans, j - i)
	}
	return ans
}
```

## 方法三 优化滑块法（不借助哈希表）
---
#### 解题思路：
参考
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/

#### 复杂度分析
* 时间复杂度 O( n )

#### 实现（Golang）

```
```
