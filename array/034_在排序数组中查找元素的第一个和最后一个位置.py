"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/23 21:32:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 034_在排序数组中查找元素的第一个和最后一个位置.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def searchRange(self, nums, target: int):
        left = 0
        length = len(nums)
        right = len(nums) - 1
        while left <= right:
            middle = left + (right - left) // 2
            if nums[middle] < target:
                left = middle + 1
            else:
                right = middle - 1

        if left != length and nums[left] == target:
            i = left + 1
            while i < length and nums[i] == target:
                i += 1
            return [left, i - 1]

        return [-1, -1]


if __name__ == '__main__':
    solution = Solution()
    print(solution.searchRange([5, 7, 7, 8, 8, 10], 8))
    print(solution.searchRange([5, 7, 7, 8, 8, 10], 6))
    print(solution.searchRange([], 0))
