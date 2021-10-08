"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/8 15:46:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 187_重复的DNA序列.py
# @Software: PyCharm
-------------------------------------
"""

from typing import List
from collections import Counter


class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        return [target for target, times in Counter(s[i:i + 10] for i in range(len(s) - 9)).items() if times > 1]


if __name__ == '__main__':
    solution = Solution()
    print(solution.findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
    print(solution.findRepeatedDnaSequences("AAAAAAAAAAAAA"))
