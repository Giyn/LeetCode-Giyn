"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/28 10:06:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 070_爬楼梯.py
# @Software: PyCharm
# @url     : https://leetcode-cn.com/problems/climbing-stairs/
-------------------------------------
"""


class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 1:
            return 1

        dp = [0 for _ in range(100)]

        # 初始值
        dp[0] = 1
        dp[1] = 1
        dp[2] = 2

        # 通过递推式来计算出dp[n]
        for i in range(2, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]

        return dp[n]


if __name__ == '__main__':
    solution = Solution()
    print(solution.climbStairs(10))
