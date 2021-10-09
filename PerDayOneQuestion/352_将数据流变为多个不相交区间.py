"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/10/9 22:02:13
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 352_将数据流变为多个不相交区间.py
# @Software: PyCharm
-------------------------------------
"""

from typing import List


class SummaryRanges:
    def __init__(self):
        self.f = [i for i in range(10001)]
        self.ans = {}

    def find(self, i):
        if self.f[i] != i:
            self.f[i] = self.find(self.f[i])
        return self.f[i]

    def merge(self, i, j):
        fi, fj = self.find(i), self.find(j)
        if fi != fj:
            self.f[fj] = fi
            self.ans[fi] += self.ans[fj]
            del self.ans[fj]

    def addNum(self, val: int) -> None:
        if self.find(val) in self.ans:
            return
        self.ans[val] = 1
        if val > 0 and self.find(val - 1) in self.ans:
            self.merge(val - 1, val)
        if val < 10000 and self.find(val + 1) in self.ans:
            self.merge(val, val + 1)

    def getIntervals(self) -> List[List[int]]:
        return sorted([[i, i + j - 1] for i, j in self.ans.items()])

# Your SummaryRanges object will be instantiated and called as such:
# obj = SummaryRanges()
# obj.addNum(val)
# param_2 = obj.getIntervals()
