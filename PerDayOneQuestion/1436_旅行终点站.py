"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/1 7:38:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1436_旅行终点站.py
# @Software: PyCharm
# @url     : https://leetcode-cn.com/problems/destination-city/
-------------------------------------
"""


class Solution:
    def destCity(self, paths) -> str:
        city_A = {path[0] for path in paths}
        return next(path[1] for path in paths if path[1] not in city_A)


if __name__ == '__main__':
    solution = Solution()
    print(solution.destCity([["London", "New York"], ["New York", "Lima"], ["Lima", "Sao Paulo"]]))
    print(solution.destCity([["B", "C"], ["D", "B"], ["C", "A"]]))
    print(solution.destCity([["A", "Z"]]))
