"""
Decorators

Decorators are a powerful and versatile feature in Python used to modify or enhance the behavior of functions or methods. 
They allow you to wrap another function, adding functionality before or after the wrapped function's execution
"""

import time


def timer(func):
    def wrapper(*args, **kwargs):
        start_time = time.time()
        result = func(*args, **kwargs)
        end_time = time.time()
        print(f"Elapsed time: {end_time - start_time}")
        return result

    return wrapper


@timer
def my_function():
    time.sleep(2)


# my_function() # Elapsed time: ...
# adding @timer to the function is equivalent to
# my_function = timer(my_function)


def repeat(n):
    def decorator(func):
        def wrapper(*args, **kwargs):
            results = []
            for _ in range(n):
                result = func(*args, **kwargs)
                results.append(result)
            return results

        return wrapper

    return decorator


@repeat(3)
def roll_dice():
    import random

    return random.randint(1, 6)


# results = roll_dice()
