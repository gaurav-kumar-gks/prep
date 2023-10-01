"""
Monotonic queue
"""

"""
Monotonic - either increasing or decreasing
    
increasing monotonic queue -
    1) if the current element is greater than the last element in the queue, then we can pop the last element in the queue,
    because the last element in the queue will never be the maximum element in the window
    2) if the current element is less than the last element in the queue, then we can append the current element to the queue,
    because the current element may be the maximum element in the window
"""
import collections


def increasing_queue(A):
    queue = collections.deque()
    l = [-1]*len(A)
    r = [-1]*len(A)
    for i, v in enumerate(A):
        while queue and A[queue[-1]] >= v:
            r[queue.pop()] = v
        if queue:
            l[i] = A[queue[-1]]
        queue.append(i)
    return l, r


def decreasing_queue(A):
    queue = collections.deque()
    l = [-1]*len(A)
    r = [-1]*len(A)
    for i, v in enumerate(A):
        while queue and A[queue[-1]] <= v:
            r[queue.pop()] = v
        if queue:
            l[i] = A[queue[-1]]
        queue.append(i)
    return l, r
