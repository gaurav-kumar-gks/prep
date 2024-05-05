"""
Function overloading

Python does not support function overloading. 
However, you can use the *args and **kwargs arguments to simulate function overloading.

there's a pkg for this - multipledispatch
Code from: https://arpitbhayani.me/blogs/function-overloading/
"""
from inspect import getfullargspec


class Function(object):
    """
    Function is a wrap over standard python function

    An instance of this Function class is also callable
    just like the python function that it wrapped.
    When the instance is "called" like a function it fetches
    the function to be invoked from the virtual namespace and then
    invokes the same.
    """
    def __init__(self, fn):
        self.fn = fn

    def __call__(self, *args, **kwargs):
        """
        Overri̇̇̇̇ding the __call__ function which makes the instance callable.
        """
        # fetching the function to be invoked from the virtual namespace
        # through the arguments.
        namespace = Namespace.get_instance()
        if namespace:
            fn = namespace.get(self.fn, *args)
        if not fn:
            raise Exception("no matching function found.")

        # invoking the wrapped function and returning the value.
        return fn(*args, **kwargs)

    def key(self, args=None):
        """
        Returns the key that will uniquely identifies a function (even when it is overloaded).
        """
        if args is None:
            args = getfullargspec(self.fn).args

        return tuple([
            self.fn.__module__,
            self.fn.__class__,
            self.fn.__name__,
            len(args) or 0,
        ])


class Namespace(object):
    """
    Namespace is the singleton class that is responsible for holding all the functions.
    """
    __instance = None

    def __init__(self):
        if self.__instance is None:
            self.function_map = dict()
            Namespace.__instance = self
        else:
            raise Exception("cannot instantiate Namespace again.")

    @staticmethod
    def get_instance():
        if Namespace.__instance is None:
            Namespace()
        return Namespace.__instance

    def register(self, fn):
        """
        Registers the function in the virtual namespace and returns
        an instance of callable Function that wraps the
        function fn.
        """
        func = Function(fn)
        specs = getfullargspec(fn)
        self.function_map[func.key()] = fn
        return func

    def get(self, fn, *args):
        """get returns the matching function from the virtual namespace.

        return None if it did not fund any matching function.
        """
        func = Function(fn)
        return self.function_map.get(func.key(args=args))


def overload(fn):
    """
    overload is the decorator that wraps the function
    and returns a callable object of type Function.
    """
    namespace = Namespace.get_instance() 
    if namespace:
        return namespace.register(fn)



@overload
def area(length, breadth):
    return length * breadth

@overload
def area(length, breadth):
    # this function will be used when 2 args are given
    # because in namespace function_map dict - the fn. will be overridden as this fn. and fn. above it will have same key
    return length * length


@overload
def area(length):
    return length * length


@overload
def area(length, breadth, height):
    return 2 * (length * breadth + breadth * height + height * length)

# print(locals())
print(f"area(4, 3, 6) is: {area(4, 3, 6)}")
print(f"area(7, 2) is: {area(7, 2)}")
print(f"area(7) is: {area(7)}")

