"""
Closures
"""
"""
A closure is a function that remembers values in the enclosing scope even if they are not present in memory. 
It captures variables from the outer function's scope and "closes over" them, preserving their values. 
Closures are created when a nested function references variables from its containing (enclosing) function's scope.
"""


def outer_func():
    m = 3

    def inner_func(c=10):
        print(f"y = mx + c => y = {m}*x + {c}")

    return inner_func


# closure = outer_func()
# closure()

outer_func()(3)

# e.g. memoization


def memoize(f):
    cache = {}

    def wrapper(*args):
        if args not in cache:
            result = f(*args)
            cache[args] = result
        return cache[args]

    return wrapper


@memoize
def fibonnaci(n):
    if n <= 1:
        return n

    return fibonnaci(n - 1) + fibonnaci(n - 2)


# adding memoize decorator is equivalent to
# fibonacci = memoize(fibonacci)
