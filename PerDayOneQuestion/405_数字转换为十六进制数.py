"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/2 9:55:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 405_数字转换为十六进制数.py
# @Software: PyCharm
-------------------------------------
"""

CONV = "0123456789abcdef"


class Solution:
    def toHex(self, num: int) -> str:
        res = []
        for _ in range(8):
            res.append(num % 16)
            num = num // 16
            if not num:
                break

        return "".join(CONV[n] for n in res[::-1])


if __name__ == '__main__':
    solution = Solution()
    print(solution.toHex(26))
    print(solution.toHex(-1))
    print(solution.toHex(-50))
