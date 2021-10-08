"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/30 23:08:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 367_有效的完全平方数.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        if num == 1:
            return True

        left = 1
        right = num

        while left <= right:
            mid = left + ((right - left) >> 1)
            if mid ** 2 == num:
                return True
            elif mid ** 2 > num:
                right = mid - 1
            else:
                left = mid + 1

        return False


if __name__ == '__main__':
    solution = Solution()
    print(solution.isPerfectSquare(16))
    print(solution.isPerfectSquare(14))
