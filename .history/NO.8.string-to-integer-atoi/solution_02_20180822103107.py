class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        # 定义边界
        max_int =  2 ** 31 - 1
        min_int = -2 ** 31
        # 如果不能执行有效的转换，则返回0
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

    def myAtoi2(self, str):
        """
        :type str: str
        :rtype: int
        """
        # 定义边界
        max_int =  2 ** 31 - 1
        min_int = -2 ** 31
        # 如果不能执行有效的转换，则返回0
        result = 0
        # 默认为正数
        sign = 1
        str_len = len(str)

        for i in range(str_len):
            asciiCode = ord(str[i])
            # 这个函数需要丢弃之前的空白字符，直到找到第一个非空白字符。
            if asciiCode == 32:
                continue

            # 之后从这个字符开始，选取一个可选的正号或负号后面跟随尽可能多的数字，并将其解释为数字的值
            # + -> 43 | - -> 45
            if asciiCode == 43:

            print(str[i])
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