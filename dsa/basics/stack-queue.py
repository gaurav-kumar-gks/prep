"""
STACK AND QUEUE
"""

"""
MONOTONIC QUEUE

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



"""
MIN STACK
- single stack to store the elements and the minimum element
"""

class MinStack:
    def __init__(self):
        self.stack = []
    def push(self, val):
        if not self.stack:
            self.stack.append((val, val))
        else:
            self.stack.append((val, min(val, self.stack[-1][1])))
    def pop(self):
        return self.stack.pop()[0]
    def top(self):
        return self.stack[-1][0]
    def getMin(self):
        return self.stack[-1][1]

"""
MIN QUEUE

- Increasing Monotonic Queue
- so front would have min (remove from front, add at back)
"""

class MinQueue:
    def __init__(self):
        self.deque = collections.deque([])
    def push(self, val):
        while self.deque and self.deque[-1] > val:
            self.deque.pop()
        self.deque.append(val)
    def pop(self):
        if self.deque[0] == 0:
            self.deque.popleft()
    def getMin(self):
        return self.deque[0]


