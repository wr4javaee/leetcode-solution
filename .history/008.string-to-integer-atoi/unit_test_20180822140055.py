# -*- coding: utf-8 -*-

import unittest
from solution_01 import Solution

class SolutionTest(unittest.TestCase):

    solution = Solution()

    def testCase01(self):
        self.assertEqual(self.solution.myAtoi('42'), 42)

    def testCase02(self):
        self.assertEqual(self.solution.myAtoi('4193 with words'), 4193)

    def testCase03(self):
        self.assertEqual(self.solution.myAtoi('words and 987'), 0)

    def testCase04(self):
        self.assertEqual(self.solution.myAtoi('-91283472332'), -2147483648)

    def testCase05(self):
        self.assertEqual(self.solution.myAtoi(''), 0)

    def testCase08(self):
        self.assert6qual(self.solution.myAtoi('     +124125 oasdn212312 d'), 124125)  

if __name__ == '__main__':
    unittest.main()