# -*- coding: utf-8 -*-

import unittest
from solution_01 import Solution

class SolutionTest(unittest.TestCase):

    solution = Solution()

    def testCase01(self):
        self.assertEqual(self.solution.myAtoi('42'), 42)

if __name__ == '__main__':
    unittest.main()