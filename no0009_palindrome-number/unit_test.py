# -*- coding: utf-8 -*-

import unittest
from solution_01 import Solution

class SolutionTest(unittest.TestCase):

    solution = Solution()

    def testCase01(self):
        self.assertEqual(self.solution.isPalindrome(121), True)

    def testCase02(self):
        self.assertEqual(self.solution.isPalindrome(-121), False)

if __name__ == '__main__':
    unittest.main()
