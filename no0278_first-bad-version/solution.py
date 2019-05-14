# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):
#
# 你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，
# 你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
#
# 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
#
# 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。
# 实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
#
# 示例:
#
# 给定 n = 5，并且 version = 4 是第一个错误的版本。
#
# 调用 isBadVersion(3) -> false
# 调用 isBadVersion(5) -> true
# 调用 isBadVersion(4) -> true
#
# 所以，4 是第一个错误的版本。
import unittest

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """

        # 安全校验
        arr_len = n
        if arr_len == 0:
            return -1

        if arr_len == 1:
            return n

        # 变形二分法查找
        # 当前节点为坏版本且无左临界点, 直接返回当前节点
        # 若mid为错误版本, 则往前半部二分
        # 若mid为正常版本, 则往后半部二分

        nums = range(1, n + 1)
        low_inx = 0
        high_inx = arr_len - 1

        while low_inx <= high_inx:
            mid_inx = (low_inx + high_inx) // 2
            mid_val = nums[mid_inx]
            is_mid_val_bad = isBadVersion(mid_val)
            if is_mid_val_bad:
                mix_left_inx = mid_inx - 1

                # 当前节点为坏版本且无左临界点, 直接返回
                if mix_left_inx < 0:
                    return mid_val

                mix_left_val = nums[mix_left_inx]
                if isBadVersion(mix_left_val):
                    # mid为错误版本, 则往前半部二分
                    high_inx = mid_inx - 1
                else:
                    # 找到第一个错误版本, 返回值
                    return mid_val
            else:
                # mid为正常版本, 则往后半部二分
                low_inx = mid_inx + 1

        # 没找到错误版本
        return -1

def isBadVersion(n):
    if n < 1:
        return False
    else:
        return True


class SolutionTest(unittest.TestCase):

    solution = Solution()

    def testCase01(self):
        res = self.solution.firstBadVersion(2)
        print(res)
        self.assertEqual(res, 1)

if __name__ == '__main__':
    unittest.main()
