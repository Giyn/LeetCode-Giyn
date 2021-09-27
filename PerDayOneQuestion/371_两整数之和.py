"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/26 15:30:13
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 371_两整数之和.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def getSum(self, a: int, b: int) -> int:
        # if a == 0:
        #     return b
        # if b == 0:
        #     return a
        # return self.getSum((a & b) << 1, a ^ b)

        while b:
            a, b = (a ^ b) & 0xFFFFFFFF, ((a & b) << 1) & 0xFFFFFFFF
        return ~(a ^ 0xFFFFFFFF) if a >= 0x7FFFFFFF else a


if __name__ == '__main__':
    solution = Solution()
    print(solution.getSum(4, 12))
    print(solution.getSum(4, 17))
