"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/10 9:03:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 441_排列硬币.py
# @Software: PyCharm
-------------------------------------
"""

import math


class Solution:
    def arrangeCoins(self, n: int) -> int:
        return int((math.sqrt(1 + 8 * n) - 1) // 2)


if __name__ == '__main__':
    solution = Solution()
    print(solution.arrangeCoins(5))
    print(solution.arrangeCoins(8))
