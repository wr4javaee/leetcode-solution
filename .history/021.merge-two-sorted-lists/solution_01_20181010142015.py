import unittest

"""定义链表结构"""
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

"""比较大小法"""
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
    while l3:
        print(l3.val, end='')
        if l3.next:
            print("->", end='')
        if l3.val != result[i]:
            return False
        l3 = l3.next
        i = i + 1
    return True    

"""测试"""
class SolutionTest(unittest.TestCase):
    solution = Solution()

    def testCase(self):
        l1_data = [1, 2, 4]
        l2_data = [1, 3, 4]
        result = [1, 1, 2, 3, 4, 4]

        L1 = LinkList()
        L1.initList(l1_data)            
        L1.traverse()
        
        L2 = LinkList()
        L2.initList(l2_data)            
        L2.traverse()

        l1 = L1.head
        l2 = L2.head
        l3 = self.solution.mergeTwoLists(l1, l2)
        
        print(isEquals(l3, result))
        # self.assertEqual(self.solution.mergeTwoLists(l1, l2), True)

if __name__ == '__main__':
    unittest.main()






