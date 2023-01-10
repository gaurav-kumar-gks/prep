"""
Itertools
"""
import operator
from itertools import *

# Infinite iterators
count(10)  # 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, ...
cycle('ABCD')  # A, B, C, D, A, B, C, D, A, B, C, D, ...
repeat(10, 3)  # 10, 10, 10

# Finite iterators
# accumulate(iterable[, func, *, initial=None])
accumulate([1, 2, 3, 4, 5])  # 1, 3, 6, 10, 15
accumulate([1, 2, 3, 4, 5], operator.mul)  # 1, 2, 6, 24, 120
accumulate([1, 2, 3, 4, 5], max)  # 1, 2, 3, 4, 5
accumulate([1, 2, 3, 4, 5], min)  # 1, 1, 1, 1, 1
accumulate([1, 2, 3, 4, 5], lambda res, item: res * item)  # 1, 2, 6, 24, 120

# chain(*iterables)
list(chain('abc', 'def'))  # ['a', 'b', 'c', 'd', 'e', 'f']
print(list(chain.from_iterable(['abc', 'def'])))  # ['a', 'b', 'c', 'd', 'e', 'f']

product('ABCD', 'xy')  # Ax Ay Bx By Cx Cy Dx Dy
product(range(2), repeat=3)  # 000 001 010 011 100 101 110 111

# groupby(iterable, key=None)
animals = ['duck', 'eagle', 'rat', 'giraffe', 'bear', 'bat', 'dolphin', 'shark', 'lion']
animals.sort(key=len)
c = [(length, list(group)) for length, group in groupby(animals, len)]  # [(3, ['rat', 'bat']), (4, ['duck', 'bear', 'lion']), (5, ['eagle', 'shark']), (7, ['giraffe', 'dolphin'])]
[k for k, g in groupby('AAAABBBCCDAABBB')]  # ['A', 'B', 'C', 'D', 'A', 'B']
[list(g) for k, g in groupby('AAAABBBCCD')]  # [['A', 'A', 'A', 'A'], ['B', 'B', 'B'], ['C', 'C'], ['D']]


combinations(['a', 'b', 'c'], 2)  # ('a', 'b'), ('a', 'c'), ('b', 'c')
combinations_with_replacement(['a', 'b', 'c'], 2)  # ('a', 'a'), ('a', 'b'), ('a', 'c'), ('b', 'b'), ('b', 'c'), ('c', 'c')
permutations(['a', 'b', 'c'])  # ('a', 'b', 'c'), ('a', 'c', 'b'), ('b', 'a', 'c'), ('b', 'c', 'a'), ('c', 'a', 'b'), ('c', 'b', 'a')


"""
functools

only functool that'll be frequently used is cache

just add @cache decorator on the function that you want to cache [acts as memoization]
"""


"""
Operators
operators may come handy in lambda functions

e.g. ->
lambda x: mod(x, 2) == 0
"""
