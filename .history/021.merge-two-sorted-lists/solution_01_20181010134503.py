import unittest

"""

"""
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        # 若L1与L2是否均为空, 返回None
        if not l1 and not l2:
            return None
        # 定义新的空链表L，设置头指针head    
        head = ListNode(-1)
        # 定义指针p，指向当前新链表节点
        p = head
        # 当l1与l2非空时，开启循环：
        while l1 and l2:
            if l1.val <= l2.val:
                p.next = l1
                l1 = l1.next
            else:
                p.next = l2
                l2 = l2.next
            p = p.next
        if l1:
            p.next = l1
        else:
            p.next = l2
        # 通过头指针head返回新链表    
        return head.next

class SolutionTest(unittest.TestCase):
    solution = Solution()
    def testCase(self):
        l1 = ListNode(1)
        l1.next = ListNode(2)
        l1.next.next = ListNode(4)
        self.assertEqual(self.solution.mergeTwoLists(l1, l2), True)

if __name__ == '__main__':
    unittest.main()