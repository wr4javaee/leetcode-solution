# -*- coding: utf-8 -*-

import unittest
from solution_01 import Solution

class Solution01Test(unittest.TestCase):

    solution = Solution()

    def testCase01(self):
        self.assertEqual(self.solution.myAtoi('42'), '42')

    def testCase02(self):
        self.assertEqual(self.solution.myAtoi('42'), '42')

if __name__ == '__main__':
    unittest.main()