package no0658_find_k_closest_elements

//给定一个排序好的数组，两个整数 k 和 x，
// 从数组中找到最靠近 x（两数之差最小）的 k 个数。
// 返回的结果必须要是按升序排好的。如果有两个数与 x 的差值一样，优先选择数值较小的那个数。
//
//示例 1:
//
//输入: [1,2,3,4,5], k=4, x=3
//输出: [1,2,3,4]
//
//
//示例 2:
//
//输入: [1,2,3,4,5], k=4, x=-1
//输出: [1,2,3,4]
//
//
//说明:
//
//k 的值为正数，且总是小于给定排序数组的长度。
//数组不为空，且长度不超过 10^4
//数组里的每个元素与 x 的绝对值不超过 10^4
//
//
//更新(2017/9/19):
//这个参数 arr 已经被改变为一个整数数组（而不是整数列表）。 请重新加载代码定义以获取最新更改。
func findClosestElements(arr []int, k int, x int) []int {

	// 题目转义
	// 题目要求从长度为n的数组中找到长度为k的子数组
	// 即从长度为n的数组中排除n-k个数字
	// 因距离x的最远数字一定为以x为中心的子数组的左右边界
	// 因此每次需比较左右边界，并将距离最远的数字排除

	// 解题思路
	// 采用二分法
	// 通过每次比较mid位置与x的距离 与 mid + k位置和x的距离选择二分方向, 最终定位到距离x最近位置

}


//class Solution {
//public:
//vector<int> findClosestElements(vector<int>& arr, int k, int x) {
//int left = 0, right = arr.size() - k;
//while (left < right) {
//int mid = left + (right - left) / 2;
//if (x - arr[mid] > arr[mid + k] - x) left = mid + 1;
//else right = mid;
//}
//return vector<int>(arr.begin() + left, arr.begin() + left + k);
//}
//};
