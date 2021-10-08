"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/8 15:41:44
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 414_第三大的数.py
# @Software: PyCharm
-------------------------------------
"""

from sortedcontainers import SortedList
from typing import List


class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        s = SortedList()
        for num in nums:
            if num not in s:
                s.add(num)
                if len(s) > 3:
                    s.pop(0)
        return s[0] if len(s) == 3 else s[-1]


if __name__ == '__main__':
    solution = Solution()
    print(solution.thirdMax([3, 2, 1]))
    print(solution.thirdMax([1, 2]))
    print(solution.thirdMax([2, 2, 3, 1]))
