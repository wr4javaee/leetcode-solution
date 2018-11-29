# LeetCode.0001 two-sum 两数之和
---

## 题目
---
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的 两个 整数。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

```
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```

## 方法一 暴力法
---
参考
https://leetcode-cn.com/problems/two-sum/solution/
#### 复杂度分析：
* 时间复杂度 O( n^2 )
* 空间复杂度 O( 1 )

#### 执行效率：
* 执行用时：60 ms，已经战胜 36.27 % 的 golang 提交记录

#### 解题思路：
暴力法很简单。遍历每个元素 xx，并查找是否存在一个值与 target - xtarget−x 相等的目标元素。

#### 实现（Golang）
```
// 方法一：暴力法
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j:= i + 1; j < len(nums); j ++ {
			if v + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```

## 方法二 两遍哈希表
参考
https://leetcode-cn.com/problems/two-sum/solution/
#### 复杂度分析：
* 时间复杂度：O(n)
我们把包含有 n 个元素的列表遍历两次。由于哈希表将查找时间缩短到 O(1) ，所以时间复杂度为 O(n)。

* 空间复杂度：O(n)
所需的额外空间（取决于哈希表中存储的元素数量，该表中存储了 n 个元素

#### 执行效率：
* 执行用时：12 ms，已经战胜 70.38 % 的 golang 提交记录

#### 解题思路：
为了对运行时间复杂度进行优化，我们需要一种更有效的方法来检查数组中是否存在目标元素。如果存在，我们需要找出它的索引。保持数组中的每个元素与其索引相互对应的最好方法是什么？哈希表。

通过以空间换取速度的方式，我们可以将查找时间从 O(n)O(n) 降低到 O(1)O(1)。哈希表正是为此目的而构建的，它支持以 近似 恒定的时间进行快速查找。我用“近似”来描述，是因为一旦出现冲突，查找用时可能会退化到 O(n)O(n)。但只要你仔细地挑选哈希函数，在哈希表中进行查找的用时应当被摊销为 O(1)O(1)。

一个简单的实现使用了两次迭代。在第一次迭代中，我们将每个元素的值和它的索引添加到表中。然后，在第二次迭代中，我们将检查每个元素所对应的目标元素（target - nums[i]target−nums[i]）是否存在于表中。注意，该目标元素不能是 nums[i]nums[i] 本身！

#### 实现（Golang）

```
// 两遍哈希表
func twoSum02(nums []int, target int) []int {
	// 初始化map
	numMap := make(map[int]int)

	// 遍历数组, 将数组中val作为map的key, 数组索引作为map的val
	for i, v := range nums {
		numMap[v] = i
	}

	// 再次遍历数组
	for i, v := range nums {
		expectNum := target - v

		// 按照题目要求不可重复使用同一个索引, 因此这里需要保证数组的索引和map中取出的索引不能相等
		if mapVal, ok := numMap[expectNum]; ok && i != mapVal {
			return []int {i, mapVal}
		}
	}

	// 没有找到结果, 返回空
	return nil
}
```

## 方法三 一遍哈希表（最优解）
---
参考
https://leetcode-cn.com/problems/two-sum/solution/

#### 解题思路：
事实证明，我们可以一次完成。在进行迭代并将元素插入到表中的同时，我们还会回过头来检查表中是否已经存在当前元素所对应的目标元素。如果它存在，那我们已经找到了对应解，并立即将其返回。

#### 复杂度分析：
* 时间复杂度：O(n)
* 空间复杂度：O(n)

#### 执行效率
* 执行用时：8 ms，我的提交执行用时
已经战胜 95.15 % 的 golang 提交记录

#### 实现（Golang）

```
	// 初始化map
	numMap := make(map[int]int)

	// 遍历数组
	for i, v := range nums {
		// 计算期望值
		expectNum := target - v

		if mapVal, ok := numMap[expectNum]; ok {
			return []int {i, mapVal}
		}

		// 将数组中val作为map的key, 数组索引作为map的val
		numMap[v] = i
	}

	// 没有找到结果, 返回空
	return nil
```

### 单元测试
---

```
package no0001_two_sum

import (
	"testing"
)

// 定义单元测试数据结构体
type testCaseModel struct {
	requestModel
	responseModel
}

// 定义被测试方法请求参数结构体
type requestModel struct {
	nums []int
	target int
}

// 定义被测试方法返回参数结构体
type responseModel struct {
	result []int
}

func TestTwoSum(t *testing.T) {
	// 构建测试数据
	testCaseArr := []testCaseModel {
		// 正常数据
		{
			requestModel {
				[]int{2, 7, 11, 15},
				9,
			},
			responseModel {
				[]int{0, 1},
			},
		},
		// 正常数据
		{
			requestModel {
				[]int{3, 2, 4},
				6,
			},
			responseModel {
				[]int{1, 2},
			},
		},
		// 正常数据
		{
			requestModel {
				[]int{3, 3},
				6,
			},
			responseModel {
				[]int{0, 1},
			},
		},
		// 找不到目标值数据
		{
			requestModel {
				[]int{2, 7, 11, 15},
				10,
			},
			responseModel {
				nil,
			},
		},
		// 入参为空的数据
		{
			requestModel {
				[]int{},
				9,
			},
			responseModel {
				nil,
			},
		},
	}

	// 执行测试方法
	// 测试方法一
	for i, v := range testCaseArr {
		actualResult := twoSum01(v.requestModel.nums, v.requestModel.target)
		expectResult := v.responseModel.result
		if resEquals(&actualResult, &expectResult) {
			t.Log("方法一测试通过, 测试用例为", i)
		} else {
			t.Error("方法一测试未通过, 测试用例为", i)
		}
	}

	// 测试方法二
	for i, v := range testCaseArr {
		actualResult := twoSum02(v.requestModel.nums, v.requestModel.target)
		expectResult := v.responseModel.result
		if resEquals(&actualResult, &expectResult) {
			t.Log("方法二测试通过, 测试用例为", i)
		} else {
			t.Error("方法二测试未通过, 测试用例、实际值、期望值分别为", i, actualResult, expectResult)
		}
	}

	// 测试方法三
	for i, v := range testCaseArr {
		actualResult := twoSum03(v.requestModel.nums, v.requestModel.target)
		expectResult := v.responseModel.result
		if resEquals(&actualResult, &expectResult) {
			t.Log("方法三测试通过, 测试用例为", i)
		} else {
			t.Error("方法三测试未通过, 测试用例、实际值、期望值分别为", i, actualResult, expectResult)
		}
	}
}

func resEquals(res1 *[]int, res2 *[]int) bool {
	if *res1 == nil || *res2 == nil {
		return true
	}

	if len(*res1) != len(*res2) {
		return false
	}

	if (*res1)[0] == (*res2)[0] && (*res1)[1] == (*res2)[1] {
		return true
	}

	if (*res1)[0] == (*res2)[1] && (*res1)[1] == (*res2)[0] {
		return true
	}

	return false
}
```

