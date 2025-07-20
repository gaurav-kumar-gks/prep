"""
NamedTuple:
A named tuple is a tuple with names for each position.
"""
from collections import namedtuple

Point = namedtuple('Point', ['x', 'y'])
p = Point(10, 20)
print(p.x)  # Output: 10
print(p.y)  # Output: 20


"""
Deque:
A deque is a double-ended queue (thread safe)
It can be used to add or remove elements from both ends in O(1)
"""
from collections import deque

d = deque([1, 2, 3])
_ = d[0]  # Output: 1
d.append(4)
d.appendleft(0)
print(d.pop())  # Output: 4
print(d.popleft())  # Output: 0
d.extend([4, 5, 6])
d.extendleft([0, -1, -2])
print(d)  # Output: deque([-2, -1, 0, 1, 2, 3, 4, 5, 6])
d.rotate(1)
print(d)  # Output: deque([6, -2, -1, 0, 1, 2, 3, 4, 5])
d.rotate(-1)
print(d)  # Output: deque([-2, -1, 0, 1, 2, 3, 4, 5, 6])
d.clear()
print(d)  # Output: deque([])


"""
Counter:
A Counter is a dict subclass for counting hashable objects. 
It is an unordered collection where elements are stored as dictionary keys and their counts are stored as dictionary values.
"""
from collections import Counter

c = Counter('gallahad')
print(c)  # Output: Counter({'a': 3, 'l': 2, 'g': 1, 'h': 1, 'd': 1})
c = Counter({'red': 4, 'blue': 2})
c.update({'red': 2, 'blue': 4})
print(c)  # Output: Counter({'red': 6, 'blue': 6})
print(list(c.elements()))  # Output: ['red', 'red', 'red', 'red', 'red', 'red', 'blue', 'blue', 'blue', 'blue', 'blue', 'blue']
print(c.most_common(1))  # Output: [('red', 6)]
c.subtract({'red': 2, 'blue': 4})
print(c)  # Output: Counter({'red': 4, 'blue': 2})
c.clear()
print(c)  # Output: Counter()

"""
OrderedDict:
An OrderedDict is a dict subclass that remembers the order in which its contents are added.
"""
from collections import OrderedDict

od = OrderedDict()
od['a'] = 1
od['b'] = 2

print(od)  # Output: OrderedDict([('a', 1), ('b', 2)])
od.move_to_end('a')
print(od)  # Output: OrderedDict([('b', 2), ('a', 1)])
od.move_to_end('a', last=False)
print(od)  # Output: OrderedDict([('a', 1), ('b', 2)])
od.popitem()
print(od)  # Output: OrderedDict([('a', 1)])
od.popitem(last=False)
print(od)  # Output: OrderedDict([('b', 2)])
od.clear()
print(od)  # Output: OrderedDict()

"""
DefaultDict:
A defaultdict is a dict subclass that calls a factory function to supply missing values.

"""
from collections import defaultdict

dd = defaultdict(int)
dd['a'] = 1
print(dd['a'])  # Output: 1
print(dd['b'])  # Output: 0
dd = defaultdict(lambda: 'empty')
dd['a'] = 'full'
print(dd['a'])  # Output: full
print(dd['b'])  # Output: empty
x = {1: 'a', 2: 'b'}
dd = defaultdict(lambda: 'missing', x)
print(dd[1])  # Output: a
print(dd[3])  # Output: missing


