"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/29 10:56:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 517_超级洗衣机.py
# @Software: PyCharm
# @url     : https://leetcode-cn.com/problems/length-of-last-word/
-------------------------------------
"""


class Solution:
    def lengthOfLastWord(self, s: str):
        return len(s.split()[-1])


if __name__ == '__main__':
    solution = Solution()
    print(solution.lengthOfLastWord("Hello World "))
    print(solution.lengthOfLastWord("   fly me   to   the moon  "))
