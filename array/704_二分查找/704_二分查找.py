"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/23 20:05:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 704_二分查找.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def search(self, nums, target: int) -> int:
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

        return -1


if __name__ == '__main__':
    solution = Solution()
    print(solution.search([-1, 0, 3, 5, 9, 12], 9))
    print(solution.search([-1, 0, 3, 5, 9, 12], 2))
