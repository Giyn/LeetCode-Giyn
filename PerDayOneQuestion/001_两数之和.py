"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/28 11:41:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 001_两数之和.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def twoSum(self, nums, target: int):
        for index1, i in enumerate(nums):
            for index2, j in enumerate(nums):
                if i + j == target and index1 != index2:
                    return [index1, index2]


if __name__ == '__main__':
    solution = Solution()
    print(solution.twoSum([2, 7, 11, 15], 9))
    print(solution.twoSum([3, 2, 4], 6))
    print(solution.twoSum([3, 3], 6))
