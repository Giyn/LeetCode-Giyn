"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/8 14:46:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 166_分数到小数.py
# @Software: PyCharm
-------------------------------------
"""

from collections import defaultdict


class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        signs = "" if numerator * denominator >= 0 else "-"
        numerator, denominator = abs(numerator), abs(denominator)
        inter_part = numerator // denominator
        remainder, remainder_map, idx = numerator % denominator, defaultdict(int), 3

        if not remainder:
            return signs + str(inter_part)
        else:
            ans = [signs, str(inter_part), "."]

        while remainder:
            if remainder not in remainder_map:
                remainder_map[remainder] = idx
            else:
                ans.insert(remainder_map[remainder], "(")
                ans.append(")")
                break

            remainder *= 10
            ans.append(str(remainder // denominator))
            remainder %= denominator
            idx += 1

        return "".join(ans)


if __name__ == '__main__':
    solution = Solution()
    print(solution.fractionToDecimal(1, 2))
    print(solution.fractionToDecimal(2, 1))
    print(solution.fractionToDecimal(2, 3))
    print(solution.fractionToDecimal(4, 333))
    print(solution.fractionToDecimal(1, 5))
