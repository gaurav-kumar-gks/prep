"""
Descriptors
"""
"""
Descriptors are Python objects that implement a method of the descriptor protocol, 
which gives you the ability to create objects that have special behavior when theyre accessed as attributes of other object

__get__(self, obj, type=None) -> object [if only __get__ is implemented then it's a non-data descriptor]
__set__(self, obj, value) -> None [if __set__ or __delete__ is implemented then it's a data descriptor]
__delete__(self, obj) -> None
__set_name__(self, owner, name)
"""


class Verbose:
    
    def __new__(cls, *args, **kwargs):
        print("creating the Verbose")
        return super().__new__(cls, *args, **kwargs)
    
    def __init__(self) -> None:
        print("initializing the Verbose")
        pass
    
    def __get__(self, obj, type=None):
        print("accessing the Verbose attribute to get the value")
        return 42
    def __set__(self, obj, value):
        print("accessing the Verbose attribute to set the value")
        raise AttributeError("Cannot change the value")

class Foo1:
    attribute1 = Verbose()

    def __init__(self) -> None:
        print("initializing the Foo")
        pass
    
    def __new__(cls, *args, **kwargs):
        print("creating the Foo")
        return super().__new__(cls, *args, **kwargs)

# my_foo_object1 = Foo1()
# x = my_foo_object1.attribute1
# print(x)

"""
property()
property(fget=None, fset=None, fdel=None, doc=None) -> object

property() returns a property object that implements the descriptor protocol
"""

class Foo2:
    @property
    def attribute1(self) -> object:
        print("accessing the attribute to get the value")
        return 42

    @attribute1.setter
    def attribute1(self, value) -> None:
        print("accessing the attribute to set the value")
        raise AttributeError("Cannot change the value")

# my_foo_object = Foo2()
# x = my_foo_object.attribute1
# print(x)

class Foo3:
    def getter(self) -> object:
        print("accessing the attribute to get the value")
        return 42

    def setter(self, value) -> None:
        print("accessing the attribute to set the value")
        raise AttributeError("Cannot change the value")

    attribute1 = property(getter, setter)

# my_foo_object = Foo3()
# x = my_foo_object.attribute1
# print(x)