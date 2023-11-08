"""
Dataclass
"""

"""
Dataclass

A dataclass is a class typically containing mainly data, although there aren't really any restrictions.
It is created using the dataclass() decorator.
It is just a regular class created by the decorator and won't be inherited by subclass


dataclass module provides a decorator and functions for automatically adding special methods such as
__init__() and __repr__() to user-defined classes
"""

from dataclasses import dataclass


@dataclass # decorator
class Point:
    y: int
    x: int = 0 # default value can be specified

"""
When you use @dataclass, the following special methods are automatically generated for your class:

__init__(): Initializes the object with values for all attributes.
__repr__(): Provides a string representation of the object, which is useful for debugging.
__eq__(): Implements equality comparison (==) based on attribute values.
__ne__(): Implements inequality comparison (!=) based on attribute values.
__lt__(), __le__(), __gt__(), and __ge__(): Implement comparison operators based on attribute values.
__hash__(): Generates a hash value for the object to be used in data structures like dictionaries and sets.
__str__(): Provides a human-readable string representation of the object.
"""

"""
Some parameters to @dataclass decorator:

order: If true (the default is false), __lt__(), __le__(), __gt__(), and __ge__() methods will be generated.
frozen: If true (the default is false), assigning to fields will generate an exception. This emulates read-only frozen instances.

"""
