# -*- coding: utf-8 -*-

import re

class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        # 以列表形式返回匹配到的字符串
        matchedArr = re.findall(r"^[\+\-]?\d+", str.strip())
        max_int = 2**31-1
        min_int = -2**31
        if matchedArr !=[]:
            if int(matchedArr[0]) > max_int:
                return max_int
            if int(matchedArr[0]) < min_int:
                return min_int
            return int(matchedArr[0])
        else:
            return 0