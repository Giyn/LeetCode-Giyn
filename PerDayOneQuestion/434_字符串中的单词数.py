"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/8 15:44:16
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 434_字符串中的单词数.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def countSegments(self, s: str) -> int:
        return len(s.split())


if __name__ == '__main__':
    solution = Solution()
    print(solution.countSegments("Hello, my name is John"))
