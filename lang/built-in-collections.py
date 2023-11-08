"""
list

list is a dynamic, mutable, heterogeneous collection of elements
"""

l1 = list()
l2 = list([1, 2, 3])
l3 = list((1, 2, 3))
l4 = list(range(1, 4))
l5 = list("abc")
l6 = list()

# important methods of list
l1.append(1)  # append an element to the end of the list
l2.extend([4, 5, 6])  # l2 += [4, 5, 6]
l3.insert(0, 2)  # insert 2 at index 0
l4.remove(2)  # removes the first occurrence of 2
l4.pop(0)  # removes the element at index 0
l5.clear()  # removes all the elements from the list
l6.index(2)  # returns the index of the first occurrence of 2
l6.count(2)  # returns the number of occurrences of 2
l6.sort()  # sorts the list in ascending order
l6.reverse()  # reverses the list
l6.copy()  # returns a shallow copy of the list
l6.sort(reverse=True)  # sorts the list in descending order
l6.sort(key=len)  # sorts the list based on the length of the elements
l6.sort(key=len, reverse=True)  # sorts the list based on the length of the elements in descending order
l6.sort(key=lambda x: x[1])  # sorts the list based on the second element of the elements

"""
array

array is a dynamic, mutable, homogeneous collection of elements
"""
import array

arr1 = array.array("i", [1, 2, 3])  # array of integers

"""
list vs array

list is a dynamic, mutable, heterogeneous collection of elements
array is a dynamic, mutable, homogeneous collection of elements

list is implemented using a dynamic array
array is implemented using a static array

list is slower than array
array is faster than list

list is more flexible than array
array is less flexible than list

list is less memory efficient than array - Each list contains pointers to a block of pointers, each of which in turn points to a full Python object
The differences between the two largely exist because of the aforementioned backend implementation. Arrays in Python are implemented just like C arrays, with a pointer pointing to the first element of the array with the rest existing contiguously in the memory
"""

"""
Set

set is an unordered, mutable, heterogeneous collection of unique elements
"""
s1 = set()  # empty set
s2 = set([1, 2, 3])  # {1, 2, 3}
s3 = set((1, 2, 3))  # {1, 2, 3}
s4 = set(range(1, 4))  # {1, 2, 3}
s5 = set(range(1, 5))  # {1, 2, 3, 4}
s6 = set("abc")  # {'a', 'b', 'c'}

# important methods of set
s1.add(1)  # add an element to the set
s2.update([4, 5, 6])  # s2 |= {4, 5, 6}
s3.remove(2)  # removes the element 2 from the set
s4.discard(2)  # removes the element 2 from the set if it is present
s5.pop()  # removes an arbitrary element from the set
s6.clear()  # removes all the elements from the set
s6.copy()  # returns a shallow copy of the set
s6.union([4, 5, 6])  # s6 | {4, 5, 6}
s6.intersection([1, 2, 3])  # s6 & {1, 2, 3}
s6.difference([1, 2, 3])  # s6 - {1, 2, 3}
# s6.symmetric_difference([1, 2, 3])  # s6 ^ {1, 2, 3}
s6.issubset([1, 2, 3])  # s6 <= {1, 2, 3}
s6.issuperset([1, 2, 3])  # s6 >= {1, 2, 3}
s6.isdisjoint([1, 2, 3])  # s6.isdisjoint({1, 2, 3})

"""
Dictionary

dictionary is an unordered, mutable, heterogeneous collection of key-value pairs
"""
d1 = dict()  # empty dictionary
d2 = dict([(1, 2), (3, 4)])  # dictionary with key-value pairs
d3 = dict.fromkeys([1, 2, 3], 0)  # dictionary with keys and default value
d4 = dict.fromkeys([1, 2, 3])  # dictionary with keys and default value None
d5 = dict.fromkeys([1, 2, 3], [0, 0, 0])  # dictionary with keys and default value [0, 0, 0]
d6 = {1: 2, 3: 4}  # dictionary with key-value pairs

# important methods of dictionary
d1.get(1)  # returns the value associated with the key 1
d1.get(1, 0)  # returns the value associated with the key 1, if the key is not present, returns 0
d2.pop(1)  # removes the key 1 from the dictionary and returns its value
d3.popitem()  # removes an arbitrary key from the dictionary and returns its value
d4.clear()  # removes all the elements from the dictionary
d5.copy()  # returns a shallow copy of the dictionary
d5.keys()  # returns a view of the keys of the dictionary
d5.values()  # returns a view of the values of the dictionary
d5.items()  # returns a view of the items of the dictionary
d6.update({1: 2})  # updates the dictionary with the key-value pair {1: 2}
d6[1] = 2  # updates the dictionary with the key-value pair {1: 2}
_ = d6[1]  # returns the value associated with the key 1


"""
tuple
"""
t1 = tuple()
t2 = tuple([1, 2, 3])
t3 = tuple((1, 2, 3))
t4 = tuple(range(1, 4))
t5 = tuple("abc")

# important methods of tuple
t1.count(1)  # returns the number of occurrences of 1
t2.index(2)  # returns the index of the first occurrence of 2

"""
Dequeue

dequeue is a double-ended queue
"""
from collections import deque

d1 = deque()  # empty dequeue
d2 = deque([1, 2, 3])  # dequeue with elements
d3 = deque((1, 2, 3))  # dequeue with elements
d4 = deque(range(1, 4))  # dequeue with elements

# important methods of dequeue
d1.append(1)  # adds 1 to the right of the dequeue
d2.appendleft(1)  # adds 1 to the left of the dequeue
d3.pop()  # removes the element at the right of the dequeue
d4.popleft()  # removes the element at the left of the dequeue
d1.extend([4, 5, 6])  # adds the elements [4, 5, 6] to the right of the dequeue
d2.extendleft([4, 5, 6])  # adds the elements [4, 5, 6] to the left of the dequeue
d3.rotate(1)  # rotates the dequeue by 1
d4.rotate(-1)  # rotates the dequeue by -1
_ = d1[0]  # returns the element at the left of the dequeue
_ = d2[-1]  # returns the element at the right of the dequeue
len(d3)  # returns the length of the dequeue


"""
Heap

heap is a binary tree where the value of each node is greater than or equal to the value of its children nodes

heapq is a module that provides functions to work with heaps

Python has a min-heap implementation
To implement a max-heap, we can use the negative of the elements and then return the negative of the elements when we pop the elements from the heap

"""
import heapq

h1 = []  # empty heap
h2 = [1, 2, 3]  # heap with elements

# important methods of heap
heapq.heappush(h1, 1)  # adds 1 to the heap
heapq.heappop(h2)  # removes the smallest element from the heap
heapq.heapify(h2)  # converts the list into a heap
_ = h1[0]  # returns the smallest element of the heap
len(h2)  # returns the length of the heap
heapq.nlargest(2, h2)  # returns the 2 largest elements of the heap
heapq.nsmallest(2, h2)  # returns the 2 smallest elements of the heap
