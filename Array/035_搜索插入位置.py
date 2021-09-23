"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/23 21:12:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 035_搜索插入位置.py
# @Software: PyCharm
# @url     : https://leetcode-cn.com/problems/search-insert-position/
-------------------------------------
"""


class Solution:
    def searchInsert(self, nums, target: int) -> int:
        left = 0
        right = len(nums) - 1
        while left <= right:
            middle = left + (right - left) // 2
            if nums[middle] > target:
                right = middle - 1
            elif nums[middle] < target:
                left = middle + 1
            else:
                return middle

        return right + 1


if __name__ == '__main__':
    solution = Solution()
    print(solution.searchInsert([1, 3, 5, 6], 5))
    print(solution.searchInsert([1, 3, 5, 6], 2))
    print(solution.searchInsert([1, 3, 5, 6], 7))
    print(solution.searchInsert([1, 3, 5, 6], 0))
    print(solution.searchInsert([1], 0))
