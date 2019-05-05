class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        max_int =  2 ** 31 - 1
        min_int = -2 ** 31
        result = 0

        if not str:
            return result

        i = 0
        while i < len(str) and str[i].isspace():
            i += 1

        if len(str) == i:
            return result

        sign = 1
        if str[i] == "+":
            i += 1
        elif str[i] == "-":
            sign = -1
            i += 1

        while i < len(str) and '0' <= str[i] <= '9':
            if result > (max_int - int(str[i])) / 10:
                return max_int if sign > 0 else min_int
            result = result * 10 + int(str[i])
            i += 1

        return sign * result