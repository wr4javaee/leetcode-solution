import unittest

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
        if l1 is None and l2 is None:
            return None
        # 定义新的空链表L，并设置头指针head
        head = ListNode(-1)
        p = head
        while l1 is not None and l2 is not None:
            if l1.val <= l2.val:
                p.next = l1
                l1 = l1.next
            else:
                p.next = l2
                l2 = l2.next
            p = p.next
        if l1 is not None:
            p.next = l1
        else:
            p.next = l2
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