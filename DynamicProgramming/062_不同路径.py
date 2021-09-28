"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/28 10:08:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 062_不同路径.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        if m <= 0 or n <= 0:
            return 0

        # 创建 m × n dp数组
        dp = [[0 for _ in range(n)] for _ in range(m)]

        # 初始化
        for i in range(m):
            dp[i][0] = 1
        for i in range(n):
            dp[0][i] = 1

        # 递推式
        for i in range(1, m):
            for j in range(1, n):
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1]

        return dp[m - 1][n - 1]


if __name__ == '__main__':
    solution = Solution()
    print(solution.uniquePaths(3, 2))
