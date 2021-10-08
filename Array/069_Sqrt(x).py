"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/30 22:51:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 069_Sqrt(x).py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def mySqrt(self, x: int) -> int:
        left = 1
        right = x
        while left <= right:
            mid = left + (right - left) // 2
            if mid * mid == x:
                return mid
            elif mid * mid > x:
                right = mid - 1
            else:
                left = mid + 1

        return right


if __name__ == '__main__':
    solution = Solution()
    print(solution.mySqrt(4))
    print(solution.mySqrt(8))
