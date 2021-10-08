"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/2 10:25:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 844_比较含退格的字符串.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        slow_pointer, fast_pointer = 0, 0


if __name__ == '__main__':
    solution = Solution()
    print(solution.backspaceCompare("ab#c", "ad#c"))
    print(solution.backspaceCompare("ab##", "c#d#"))
    print(solution.backspaceCompare("a##c", "#a#c"))
    print(solution.backspaceCompare("a#c", "b"))
