# -*- coding: utf-8 -*-

import unittest
from solution_01 import Solution01

class Solution01Test(unittest.TestCase):

    solution_01 = Solution01()

    def testCase01(self):
        self.assertEqual(self。solution_01.myAtoi(), 'FOO')

    def test_isupper(self):
        self.assertTrue('FOO'.isupper())
        self.assertFalse('Foo'.isupper())

    def test_split(self):
        s = 'hello world'
        self.assertEqual(s.split(), ['hello', 'world'])
        # check that s.split fails when the separator is not a string
        with self.assertRaises(TypeError):
            s.split(2)

if __name__ == '__main__':
    unittest.main()