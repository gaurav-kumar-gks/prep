"""
Most important built-in functions
"""
from functools import reduce

len('abc')  # 3
len([1, 2, 3])  # 3
len({'a': 1, 'b': 2})  # 2
len({'a', 'b'})  # 2

max(1, 2, 3)  # 3
max([1, 2, 3])  # 3
max('abc')  # 'c'
min(1, 2, 3)  # 1
min([1, 2, 3])  # 1
min('abc')  # 'a'

sum([1, 2, 3])  # 6
sum([1, 2, 3], 10)  # 16
sum({1, 2, 3})  # 6
sum('abc')  # 'abc'

sorted([3, 2, 1])  # [1, 2, 3]
sorted([3, 2, 1], reverse=True)  # [3, 2, 1]
sorted('abc')  # ['a', 'b', 'c']
sorted('cba', reverse=True)  # ['c', 'b', 'a']
sorted([3, 2, 1], key=lambda x: -x)  # [3, 2, 1]
sorted([3, 2, 1], key=lambda x: x % 2)  # [2, 3, 1]
sorted([3, 2, 1], key=lambda x: x % 2, reverse=True)  # [1, 3, 2]
sorted([(1, 2), (3, 1), (2, 3)], key=lambda x: x[0])  # [(1, 2), (2, 3), (3, 1)]
sorted({'a': 1, 'b': 2, 'c': 3}.items(), key=lambda x: x[1])  # [('a', 1), ('b', 2), ('c', 3)]
sorted({'a': 1, 'b': 2, 'c': 3}.items(), key=lambda x: x[1], reverse=True)  # [('c', 3), ('b', 2), ('a', 1)]


range(10)  # [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
range(1, 10)  # [1, 2, 3, 4, 5, 6, 7, 8, 9]
range(1, 10, 2)  # [1, 3, 5, 7, 9]

slice(1, 10)  # slice(1, 10, None)

"""
Filter Map Reduce
"""
filter(lambda x: x > 0, [1, -2, 3, -4])  # <filter object at 0x7f9b7c2c8e60>
list(filter(lambda x: x > 0, [1, -2, 3, -4]))  # [1, 3]

map(lambda x: x * 2, [1, 2, 3])  # <map object at 0x7f9b7c2c8e60>
list(map(lambda x: x ** 2, range(1, 11)))  # [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]

reduce(lambda x, y: x + y, [1, 2, 3])  # 6
reduce(lambda x, y: x + y, [1, 2, 3], 10)  # 16

"""
Zip
"""
zip([1, 2, 3], [4, 5, 6])  # <zip object at 0x7f9b7c2c8e60>
list(zip([1, 2, 3], [4, 5, 6]))  # [(1, 4), (2, 5), (3, 6)]


abs(-10)  # 10
all([1, 2, 3, 0])  # False
all([1, 2, 3, 4])  # True
any([1, 2, 3])  # True
any([0, 0, 0])  # False
bin(10)  # '0b1010'
chr(97)  # 'a'
enumerate([1, 2, 3])  # <enumerate object at 0x7f9b7c2c8e60>
list(enumerate([1, 2, 3]))  # [(0, 1), (1, 2), (2, 3)]
float(10)  # 10.0
hex(10)  # '0xa'
int(10.5)  # 10
int('10')  # 10
int('10', 2)  # 2
int('10', 8)  # 8
int('10', 16)  # 16

"""
getattr(object, name[, default]) 
Return the value of the named attribute of object. name must be a string.

hasattr(object, name)
Return whether the object has an attribute with the given name. (This is done by calling getattr(object, name) and catching AttributeError.)
"""
getattr(10, 'bit_length')  # <built-in method bit_length of int object at 0x7f9b7c2c8e60>
hasattr(10, 'bit_length')  # True


"""
isinstance(object, classinfo)
Return whether an object is an instance of a class or of a subclass thereof.

issubclass(class, classinfo)
Return whether class is a subclass (i.e., a derived class) of classinfo. 
A class is considered a subclass of itself. 

classinfo may be a tuple of class objects, in which case every entry in classinfo will be checked.
"""
isinstance(10, int)  # True
isinstance(10, float)  # False
issubclass(int, int)  # True
issubclass(int, float)  # False
