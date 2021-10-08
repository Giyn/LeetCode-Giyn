"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/30 14:19:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 223_矩形面积.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def computeArea(self, ax1: int, ay1: int, ax2: int, ay2: int, bx1: int, by1: int, bx2: int, by2: int) -> int:
        # 求两个矩形x轴方向相交的长度
        x = max(0, min(ax2, bx2) - max(ax1, bx1))
        # 求两个矩形y轴方向相交的长度
        y = max(0, min(ay2, by2) - max(ay1, by1))
        # 两个矩形的面积和 - 相交的面积
        return (ay2 - ay1) * (ax2 - ax1) + (by2 - by1) * (bx2 - bx1) - (x * y)


if __name__ == '__main__':
    solution = Solution()
    # print(solution.computeArea(ax1=-3, ay1=0, ax2=3, ay2=4, bx1=0, by1=-1, bx2=9, by2=2))
    # print(solution.computeArea(ax1=-2, ay1=-2, ax2=2, ay2=2, bx1=-2, by1=-2, bx2=2, by2=2))
    print(solution.computeArea(-2, -2, 2, 2, 3, 3, 4, 4))
    print(solution.computeArea(-2, -2, 2, 2, -3, 1, -1, 3))
