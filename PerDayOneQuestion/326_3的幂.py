"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/28 11:36:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 326_3的幂.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        while n > 0 and n % 3 == 0:
            n = n // 3

        return n == 1


if __name__ == '__main__':
    solution = Solution()
    print(solution.isPowerOfThree(27))
    print(solution.isPowerOfThree(0))
    print(solution.isPowerOfThree(9))
