"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/29 10:56:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 517_超级洗衣机.py
# @Software: PyCharm
# @url     : https://leetcode-cn.com/problems/super-washing-machines/
-------------------------------------
"""


class Solution:
    def findMinMoves(self, machines) -> int:
        sum_, length = sum(machines), len(machines)
        if sum_ % length != 0:
            return -1

        div, res, ans = sum_ // length, 0, 0

        for i in range(length):
            # 前面的洗衣机共有多少衣服
            ans += machines[i]
            # 前面的洗衣机还需要多少衣服
            need = ans - div * (i + 1)
            res = max(res, abs(need), machines[i] - div)

        return res


if __name__ == '__main__':
    solution = Solution()
    print(solution.findMinMoves([0, 2, 0]))
    print(solution.findMinMoves([1, 0, 5]))
    print(solution.findMinMoves([0, 3, 0]))
