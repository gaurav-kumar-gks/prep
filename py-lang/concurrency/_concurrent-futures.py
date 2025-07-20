"""
concurrent.futures
high-level interface for asynchronously executing callables using threads or processes

executor:
- ThreadPoolExecutor: creates a pool of threads
- ProcessPoolExecutor: creates a pool of processes

future:
- represents the result of an asynchronous computation
"""

from concurrent.futures import (
    ThreadPoolExecutor, as_completed,ProcessPoolExecutor
)

import math
import random
import time

def fetch_url(url):
    time.sleep(random.uniform(0.1, 0.5))
    return url, random.choice([200, 400])

urls = [
    "https://www.example.com",
    "https://www.python.org",
    "https://www.github.com"
]

with ThreadPoolExecutor(max_workers=3) as executor:
    futures = {executor.submit(fetch_url, url): url for url in urls}
    for future in as_completed(futures):
        url = futures[future]
        try:
            url, status = future.result()
            print(f"{url} returned status {status}")
        except Exception as e:
            print(f"{url} generated an exception: {e}")


def compute_factorial(n):
    return math.factorial(n)

numbers = [5, 10, 15, 20]

with ProcessPoolExecutor(max_workers=4) as executor:
    results = list(executor.map(compute_factorial, numbers))
    print(results)