import pprint

p = set()
q = {}
p is q # True if p and q are the same object (identical)

"""
List comprehension 
"""
[lambda x: x ** 2 for x in range(10) if x % 2 == 0]  # [0, 4, 16, 36, 64]
[lambda x, y: x + y for x in range(10) for y in range(10) if x % 2 == 0 and y % 2 == 0]
list(map(lambda x: x ** 2, range(10)))  # [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]
list(map(lambda x, y: x + y, range(10), range(10)))  # [0, 2, 4, 6, 8, 10, 12, 14, 16, 18]
list(filter(lambda x: x % 2 == 0, range(10)))  # [0, 2, 4, 6, 8]
[[0 for _ in range(3)] for _ in range(3)]  # [[0, 0, 0], [0, 0, 0], [0, 0, 0]]
[[0] * 3 for _ in range(3)]  # [[0, 0, 0], [0, 0, 0], [0, 0, 0]]
list(map(lambda x: [0 for _ in range(3)], range(3)))  # [[0, 0, 0], [0, 0, 0], [0, 0, 0]]

"""
Dict comprehension
"""
d = {x: x ** 2 for x in range(10) if x % 2 == 0}  # {0: 0, 2: 4, 4: 16, 6: 36, 8: 64}

"""
Set comprehension
"""
s = {x ** 2 for x in range(10) if x % 2 == 0}  # {0, 64, 4, 36, 16}

"""
Unpacking using * and ** operator
"""
a, *b, c = [1, 2, 3, 4, 5]  # a = 1, b = [2, 3, 4], c = 5
d = {"A": 10, "B": 20, "C": 30}
print(*d)  # A B C
l = [1, 2, 3]
print(*l)  # 1 2 3
t = (1, 2, 3)
print(*t)  # 1 2 3
d = {'a': 1, 'b': 2, 'c': 3}
{**d, 'd': 4}  # {'a': 1, 'b': 2, 'c': 3, 'd': 4}

# Transposing a matrix
l = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
print(list(zip(*l)))  # [(1, 4, 7), (2, 5, 8), (3, 6, 9)]

# Rotating a matrix by 90 degrees
l = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
print(list(zip(*l[::-1])))  # [(7, 4, 1), (8, 5, 2), (9, 6, 3)]

# Flattening a matrix
l = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
print([x for y in l for x in y])  # [1, 2, 3, 4, 5, 6, 7, 8, 9]

# Flattening a matrix using itertools
import itertools
l = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
list(itertools.chain(*l))  # [1, 2, 3, 4, 5, 6, 7, 8, 9]
list(itertools.chain.from_iterable(l))  # [1, 2, 3, 4, 5, 6, 7, 8, 9]

# Most common element in a list
from collections import Counter
l = [1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9]
Counter(l).most_common(1)  # [(1, 3)]
max(set(l), key=l.count)  # 1

# Ascii conversions
print(ord('a'))  # 97
print(chr(97))  # a

# Random numbers
import random
random.randint(1, 10)  # 1 <= x <= 10
random.randrange(0, 50, 8) # 0 <= x < 50, step = 8
random.random()  # 0 <= x < 1
random.choice([1, 2, 3, 4, 5])  # 1, 2, 3, 4, 5

# Swapping
a = 10
b = 5
a, b = b, a

# Negative indexing
l = [1, 2, 3, 4, 5]
print(l[-1])  # 5
print(l[-2])  # 4

# Reversing a list
l = [1, 2, 3, 4, 5]
print(l[::-1])  # [5, 4, 3, 2, 1]
print(list(reversed(l)))  # [5, 4, 3, 2, 1]

# Multiplying strings
print("abc" * 3)  # abcabcabc

# Replacing the list w/o creating a new object
nums = [1, 2, 3, 4, 5]
k = 2
nums[:] = nums[-k:] + nums[:-k]

# Big number readability
# x = 1_000_000  # 1000000

# Chain operator
a = 1
b = 0 < a < 10  # True

# Pretty print
x = {"A": 10, "B": 20, "C": 30}
print(x)  # {'A': 10, 'B': 20, 'C': 30}
pprint.pprint(x)  # {'A': 10, 'B': 20, 'C': 30} but in nicely formatted way

# advanced methods of printing
a = "abc"
b = "def"
print(f"{a} {b}")  # abc def
print("{} {}".format(a, b))  # abc def
print("%s %s" % (a, b))  # abc def
print(a, b)  # abc def
print(a, b, sep='@')  # abc@def
print(a, b, sep='@', end=',')  # abc@def,
print(f"{a = } {b = }")  # a = 'abc' b = 'def'
