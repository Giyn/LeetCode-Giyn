"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/8 15:37:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 482_密钥格式化.py
# @Software: PyCharm
-------------------------------------
"""


class Solution:
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        res = []
        string = "".join(s.split('-'))
        length = len(string)
        first = length % k
        if first:
            res.append(string[:first].upper())
            string = string[first:]
        for i in range(0, len(string), k):
            res.append(string[i:i + k].upper())

        return "-".join(res)


if __name__ == '__main__':
    solution = Solution()
    print(solution.licenseKeyFormatting("5F3Z-2e-9-w", 4))
    print(solution.licenseKeyFormatting("2-5g-3-J", 2))
