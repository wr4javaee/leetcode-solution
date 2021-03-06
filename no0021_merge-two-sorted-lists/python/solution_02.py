import unittest

"""定义链表结构"""
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

"""递归法"""
class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        if not l1:
            return l2  
        elif not l2:
            return l1
        else:
            if l1.val <= l2.val:  
                l1.next = self.mergeTwoLists(l1.next, l2)
                return l1
            else:
                l2.next = self.mergeTwoLists(l1, l2.next)
                return l2

"""链表工具"""
class LinkList(object):
    def __init__(self):
        self.head = None

    def initList(self, data):
        self.head = ListNode(data[0])
        p = self.head
        for i in data[1:]:
            node = ListNode(i)
            p.next = node
            p = p.next

    def isEmpty(self):
        if not self.head.next:
            return True
        else:
            return False

    def traverse(self):
        if self.isEmpty():
            return
        p = self.head
        while p:  
            print(p.val, end='')
            if p.next:
                print("->", end='')
            p = p.next

def isEquals(list_node, result_arr):
    i = 0
    while list_node:
        print(list_node.val, end='')
        if list_node.next:
            print("->", end='')
        if list_node.val != result_arr[i]:
            return False
        list_node = list_node.next
        i = i + 1
    print("\n")    
    return True    

"""测试"""
class SolutionTest(unittest.TestCase):
    solution = Solution()

    def testCase01(self):
        l1_data = [1, 2, 4]
        l2_data = [1, 3, 4]
        result = [1, 1, 2, 3, 4, 4]

        L1 = LinkList()
        L1.initList(l1_data)            
        
        L2 = LinkList()
        L2.initList(l2_data)            

        l1 = L1.head
        l2 = L2.head
        l3 = self.solution.mergeTwoLists(l1, l2)
        
        self.assertEqual(isEquals(l3, result), True)

    def testCase02(self):
        l1_data = [1, 2, 4]
        result = [1, 2, 4]

        L1 = LinkList()
        L1.initList(l1_data)            

        l1 = L1.head
        l3 = self.solution.mergeTwoLists(l1, None)
        
        self.assertEqual(isEquals(l3, result), True)

    def testCase03(self):
        self.assertEqual(self.solution.mergeTwoLists(None, None), None)

if __name__ == '__main__':
    unittest.main()






