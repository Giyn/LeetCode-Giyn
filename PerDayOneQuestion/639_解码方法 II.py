"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/27 22:29:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 639_解码方法 II.py
# @Software: PyCharm
# @url     : https://leetcode-cn.com/problems/decode-ways-ii/
-------------------------------------
"""

ONE = {'1': 1, '2': 1, '3': 1, '4': 1, '5': 1, '6': 1, '7': 1, '8': 1, '9': 1, '*': 9}
TWO = {'10': 1, '11': 1, '12': 1, '13': 1, '14': 1, '15': 1, '16': 1, '17': 1, '18': 1, '19': 1,
       '20': 1, '21': 1, '22': 1, '23': 1, '24': 1, '25': 1, '26': 1, '*0': 2, '*1': 2, '*2': 2,
       '*3': 2, '*4': 2, '*5': 2, '*6': 2, '*7': 1, '*8': 1, '*9': 1, '1*': 9, '2*': 6, '**': 15}


class Solution:
    def numDecodings(self, s: str) -> int:
        dp = 1, ONE.get(s[:1], 0)
        for i in range(1, len(s)):
            dp = dp[1], (ONE.get(s[i], 0) * dp[1] + TWO.get(s[i - 1: i + 1], 0) * dp[
                0]) % 1000000007
        return dp[-1]
