"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/11 22:24:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 273_整数转英文表示.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def __init__(self):
        self.nt = ["Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
                   "Ten", "Eleven", "Twelve",
                   "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen",
                   "Nineteen"]
        self.tens = ["", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty",
                     "Ninety"]
        self.t = ["Thousand", "Million", "Billion"]

    def numberToWords(self, num: int) -> str:
        def helper(num) -> list[str]:
            if num < 20:
                return [self.nt[num]]
            elif num < 100:
                res = [self.tens[num // 10]]
                if num % 10:
                    res += helper(num % 10)
                return res
            elif num < 1000:
                res = [self.nt[num // 100], "Hundred"]
                if num % 100:
                    res += helper(num % 100)
                return res
            for p, w in enumerate(self.t, 1):
                if num < 1000 ** (p + 1):
                    return helper(num // 1000 ** p) + [w] + helper(
                        num % 1000 ** p) if num % 1000 ** p else helper(num // 1000 ** p) + [w]

        return " ".join(helper(num))


if __name__ == '__main__':
    solution = Solution()
    print(solution.numberToWords(12345))
    print(solution.numberToWords(1234567))
    print(solution.numberToWords(1234567891))
