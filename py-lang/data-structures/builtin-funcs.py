"""
Most important built-in functions
"""

"""
Filter Map Zip
"""
filter(lambda x: x > 0, [1, -2, 3, -4])  # <filter object at 0x7f9b7c2c8e60>
list(filter(lambda x: x > 0, [1, -2, 3, -4]))  # [1, 3]

map(lambda x: x * 2, [1, 2, 3])  # <map object at 0x7f9b7c2c8e60>
list(map(lambda x: x ** 2, range(1, 11)))  # [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]

zip([1, 2, 3], [4, 5, 6])  # <zip object at 0x7f9b7c2c8e60>
list(zip([1, 2, 3], [4, 5, 6]))  # [(1, 4), (2, 5), (3, 6)]

"""
Bisect
bisect_left: returns the index where the target should be inserted in the sorted list
bisect_right: returns the index after the last occurrence of the target
"""
import bisect

sorted_list = [1, 3, 5, 7, 9]
index = bisect.bisect_left(sorted_list, 5) # 2
index = bisect.bisect_left(sorted_list, 10) # 5
index = bisect.bisect_left(sorted_list, 2) # 1 
index = bisect.bisect_right(sorted_list, 1) # 1
index = bisect.bisect_right(sorted_list, 10) # 5

"""
Functools
"""
from functools import (
    reduce,
    partial,
    lru_cache,
    wraps,
    singledispatch,
)

result = reduce(lambda x, y: x + y, [1, 2, 3, 4, 5])
print(result)  # Output: 15

double = partial(lambda x, y: x * y, 2)
print(double(5))  # Output: 10

@lru_cache()
def fibonacci(n):
    if n < 2: return n
    return fibonacci(n-1) + fibonacci(n-2)
print(fibonacci(10))  # Output: 55

def my_decorator(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        print("Before function call")
        result = func(*args, **kwargs)
        print("After function call")
        return result
    return wrapper

@my_decorator
def say_hello(name):
    """Greet someone by name."""
    print(f"Hello, {name}!")

say_hello("Alice") # Output: Before function call\n Hello, Alice!\n After function call
print(say_hello.__name__)  # Output: say_hello
print(say_hello.__doc__)   # Output: Greet someone by name.

@singledispatch
def process(value):
    return f"Unsupported type: {type(value)}"

@process.register(int)
def _(value):
    return value * 2

@process.register(str)
def _(value):
    return value.upper()

print(process(10))  # Output: 20
print(process("hello"))  # Output: HELLO
print(process(3.14))  # Output: Unsupported type: <c

"""
Itertools
"""
from itertools import (
    chain, 
    zip_longest, 
    starmap, 
    accumulate,
    product, 
    permutations, 
    combinations, 
    count, 
    cycle, 
    repeat, 
    compress, 
    dropwhile, 
    takewhile, 
    groupby,
)
import operator

l = list(chain([1, 2, 3], [4, 5], [6, 7, 8])) # [1, 2, 3, 4, 5, 6, 7, 8]
l = list(zip_longest([1, 2], ['a', 'b', 'c'], fillvalue=0)) # [(1, 'a'), (2, 'b'), (0, 'c')]
l = list(starmap(lambda x, y: x + y, [(1, 2), (3, 4), (5, 6)])) # [3, 7, 11]
l = list(accumulate([1, 2, 3, 4, 5])) # [1, 3, 6, 10, 15]
l = list(accumulate([1, 2, 3, 4, 5], operator.mul)) # [1, 2, 6, 24, 120]
l = list(product([1, 2], ['a', 'b'])) # [(1, 'a'), (1, 'b'), (2, 'a'), (2, 'b')]
l = list(permutations([1, 2, 3], 2)) # [(1, 2), (1, 3), (2, 1), (2, 3), (3, 1), (3, 2)]
l = list(combinations([1, 2, 3], 2)) # [(1, 2), (1, 3), (2, 3)]

for i in count(10, 2):
    if i > 20:
        break
    print(i)  # Output: 10, 12, 14, 16, 18, 20

cycler = cycle(['red', 'green', 'blue'])
for i in range(6):
    print(next(cycler))  # Output: red, green, blue, red, green, blue
    
l = list(repeat(5, 3)) # [5, 5, 5]
    
data, selectors = ['a', 'b', 'c', 'd'], [True, False, True, False]
l = list(compress(data, selectors)) # ['a', 'c']
    
data = [1, 2, 3, 4, 5]
l = list(dropwhile(lambda x: x < 3, data)) # [3, 4, 5]
l = list(takewhile(lambda x: x < 3, data)) # [1, 2]

data = [('a', 1), ('a', 2), ('b', 3), ('b', 4), ('c', 5)]
groups = groupby(data, lambda x: x[0])
d = {key: list(group) for key, group in groups} # {'a': [('a', 1), ('a', 2)], 'b': [('b', 3), ('b', 4)], 'c': [('c', 5)]}


"""
Commonly used
"""
import sys

# for max and min
sys.maxsize  # 9223372036854775807

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

list(reversed([1,2,3])) # [3, 2, 1]
''.join(reversed('abc'))  # 'cba'

range(10)  # [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
range(1, 10)  # [1, 2, 3, 4, 5, 6, 7, 8, 9]
range(1, 10, 2)  # [1, 3, 5, 7, 9]

slice(1, 10)  # slice(1, 10, None)

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
