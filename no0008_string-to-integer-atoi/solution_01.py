# -*- coding: utf-8 -*-

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
        # 默认为正数
        sign = 1
        # 字符串长度
        str_len = len(str)

        # 规则：字符串为空则不进行转换。
        if str_len == 0:
            return 0

        # 是否需要进行空白符判断
        isNeedFindNonBlank = True
        # 是否需要进行符号判断
        isNeedFindSymbol = True
        # 顺序遍历字符串字符
        for i in range(str_len):
            asciiCode = ord(str[i])
            # 规则：这个函数需要丢弃之前的空白字符，直到找到第一个非空白字符。
            if isNeedFindNonBlank == True and asciiCode == 32:
                continue

            # 规则：之后从这个字符开始，选取一个可选的正号或负号后面跟随尽可能多的数字，并将其解释为数字的值
            # + -> 43 | - -> 45
            # 判断正号
            if isNeedFindSymbol == True and asciiCode == 43:
                isNeedFindNonBlank = False
                isNeedFindSymbol = False
                sign = 1
                continue

            # 判断负号
            if isNeedFindSymbol == True and asciiCode == 45:
                isNeedFindNonBlank = False
                isNeedFindSymbol = False
                sign = -1
                continue

            # 规则：如果字符串中的第一个非空白的字符不是有效的整数，则不进行转换
            # 说明: 如果触发此条件, isNeedFindNonBlank为True意味着字符为非空白字符, isNeedFindSymbol为True意味着字符为非符号字符
            if isNeedFindNonBlank == True and isNeedFindSymbol == True and (asciiCode < 48 or asciiCode > 57):
                return 0

            # 如果第一个数字前没有符号, 则默认符号为正号
            # 0~9 -> 48~57
            if asciiCode >= 48 and asciiCode <= 57:
                # 当发现数字时, 就不再判断空白字符和符号字符了
                isNeedFindNonBlank = False
                isNeedFindSymbol = False
                # 判断是否溢出
                if (result > int(max_int / 10)) or (result == int(max_int / 10) and int(str[i]) > 7):
                    # 溢出了, 返回 INT_MAX（2147483647）或 INT_MIN（-2147483648）
                    if sign == 1:
                        return max_int
                    if sign == -1:
                        return min_int    
                # 开始进行数字转换
                result = result * 10 + int(str[i])
                continue

            # 字符串可以在形成整数的字符后包括多余的字符，将这些字符忽略，这些字符对于函数的行为没有影响。
            break

        # 字符串为空或者只包含空白字符则不进行转换。
        if isNeedFindNonBlank == True:
            return 0

        # 返回结果
        return sign * result