# -*- coding: utf-8 -*-

import re

class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        # 以列表形式返回匹配到的字符串
        # r原生字符, 将在python里有特殊意义的字符如\b，转换成原生字符（就是去除它在python的特殊意义），
        # 不然会给正则表达式有冲突，为了避免这种冲突可以在规则前加原始字符r
        matchedArr = re.findall(r"^[\+\-]?\d+", str)
        max_int = 2 ** 31 - 1
        min_int = -2 ** 31
        if matchedArr !=[]:
            if int(matchedArr[0]) > max_int:
                return max_int
            if int(matchedArr[0]) < min_int:
                return min_int
            return int(matchedArr[0])
        else:
            return 0