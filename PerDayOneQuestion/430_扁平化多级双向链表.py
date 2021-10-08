"""
-------------------------------------
# -*- coding: utf-8 -*-
# @Time    : 2021/9/24 11:05:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 430_扁平化多级双向链表.py
# @Software: PyCharm
-------------------------------------
"""


class Node:
    def __init__(self, val, prev, next, child):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child


class Solution:
    def flatten(self, head):
        def dfs(node):
            if not node:
                return
            if not node.child and not node.next:
                return node
            elif node.child:
                last = dfs(node.child)
                if last:
                    last.next = node.next
                if node.next:
                    node.next.prev = last
                node.next = node.child
                node.child.prev = node
                node.child = None
                return dfs(last)
            else:
                return dfs(node.next)

        dfs(head)

        return head


if __name__ == '__main__':
    solution = Solution()
