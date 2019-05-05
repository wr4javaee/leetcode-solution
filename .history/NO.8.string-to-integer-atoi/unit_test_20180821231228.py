# -*- coding: utf-8 -*-

import unittest
from solution_01 import Solution

class SolutionTest(unittest.TestCase):

    solution = Solution()

    def testCase01(self):
        print(self.solution.myAtoi('42'))
        # self.assertEqual(self.solution.myAtoi('42'), '42')

    def testCase02(self):
        self.assertEqual(self.solution.myAtoi('4193 with words'), '4193')

if __name__ == '__main__':
    unittest.main()