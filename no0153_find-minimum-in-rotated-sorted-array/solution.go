package no0153_find_minimum_in_rotated_sorted_array
//
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//请找出其中最小的元素。
//
//你可以假设数组中不存在重复元素。
//
//示例 1:
//
//输入: [3,4,5,1,2]
//输出: 1
//示例 2:
//
//输入: [4,5,6,7,0,1,2]
//输出: 0
func findMin(nums []int) int {
	// 因数组升序、无重复数据
	// 题目要求即找到旋转点
	// 采用变形二分法，思路：
	// 0、若数组维发生旋转、或数组元素个数为1, 直接返回数组首元素
	// 1、若数组发生旋转, 则数组尾元素一定小于首元素
	// 2、若数组发生旋转, mid值左边数组为升序数组条件为：mid值比mid左邻节点值大且大于数组首元素值（12340）
	// 3、若数组发生旋转, mid值右边数组为升序数组条件为: mid值比mid左邻节点大且小于数组首元素值（51234）
	// 4、若数组发生旋转, 数组长度大于1，则二分过程中不会发生mid索引为数组首索引或最后一个索引
	// 4、若数组发生旋转, 则最小值一定处于旋转的非升序部分中
	// 5、若数组发生旋转, mid值小于左邻节点值, 则mid值为最小值

	// 异常校验
	// 数组长度为0
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}

	// 数组长度为1
	if numsLen == 1 {
		return nums[0]
	}

	// 采用变形二分法, 确保mid左邻节点能够取到值
	lowInx := 0
	highInx := numsLen
	firstVal := nums[0]
	// 判断数组是否为旋转数组
	// 对于未旋转数组, 首元素为最小值
	if nums[highInx - 1] > nums[lowInx] {
		return nums[lowInx]
	}

	// 数组肯定为旋转数组, 设置循环条件
	for lowInx < highInx {
		// 向下整除得到midInx
		midInx := (lowInx + highInx) / 2
		midLeftInx := midInx - 1
		midVal := nums[midInx]
		midLeftVal := nums[midLeftInx]

		// 若数组发生旋转, mid值左边数组为升序数组条件为：mid值比mid左邻节点值大且大于lowInx值（12340）
		if midVal > midLeftVal && midVal > firstVal {
			// 1、2、3、4、0, 当mid值为3时, 在右区间找最小值
			lowInx = midInx + 1
			continue
		}

		// 若数组发生旋转, mid值右边数组为升序数组条件为: mid值比mid左邻节点大且小于lowInx值（51234）
		if midVal > midLeftVal && midVal < firstVal {
			// 5、1、2、3、4, 当mid值为2时, 在左区间找最小值
			highInx = midInx
			continue
		}

		// 不会发生mid索引为数组首索引
		// mid值若比左邻节点值小, 则mid值为最小值
		if midVal < midLeftVal {
			return midVal
		}
	}

	// 此题不用考虑lowInx与highInx指向同一个节点的情况,
	// 在发生这一步之前已经判定出来结果了
	return -1
}
