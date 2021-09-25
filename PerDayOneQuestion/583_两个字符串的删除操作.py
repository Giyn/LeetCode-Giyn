"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/25 15:19:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 583_两个字符串的删除操作.py
# @Software: PyCharm
# @url     : https://leetcode-cn.com/problems/delete-operation-for-two-strings/
-------------------------------------
"""


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        dp = [[0] * (len(word2) + 1) for _ in range(len(word1) + 1)]
        for i in range(len(word1) + 1):
            dp[i][0] = i
        for j in range(len(word2) + 1):
            dp[0][j] = j
        for i in range(1, len(word1) + 1):
            for j in range(1, len(word2) + 1):
                if word1[i - 1] == word2[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    dp[i][j] = min(dp[i - 1][j - 1] + 2, dp[i - 1][j] + 1, dp[i][j - 1] + 1)
        return dp[-1][-1]


if __name__ == '__main__':
    solution = Solution()
    print(solution.minDistance("sea", "ate"))
