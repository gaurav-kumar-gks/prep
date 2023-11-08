import time
import timeit


def fn():
    time.sleep(0.1)


if __name__ == "__main__":
    print(timeit.timeit(fn, number=1))
