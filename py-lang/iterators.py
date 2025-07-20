"""
Iterators
"""

"""
An iterator is an object that represents a stream of data, allowing you to iterate over its elements one at a time

A Python object is considered an iterator if it implements two special methods:
    __iter__ and __next__
    
__iter__ returns the iterator object itself. typically returns self

__next__ returns the next value from the iterator. If there is no more items to return, it raises StopIteration
"""

class MyIterator:
    def __init__(self, start, end):
        self.current = start
        self.end = end

    def __iter__(self):
        return self

    def __next__(self):
        if self.current >= self.end:
            raise StopIteration
        value = self.current
        self.current += 1
        return value
    
my_iter1 = MyIterator(1, 5)

print("iterating over my_iter1 using for loop:")
for num in my_iter1:
    print(num)  # Outputs: 1 2 3 4
    
my_iter2 = MyIterator(1, 5)

print("iterating over my_iter2 using next():")
while True:
    try:
        print(next(my_iter2))  # Outputs: 1 2 3 4
    except StopIteration:
        break


"""
from collections.abc import Iterator

We can create a iterator by inheriting Iterator
this way we won't have to write the .__iter__() method
"""

