"""
Josephus Problem

We are given the natural numbers n and k. 
All natural numbers from 1 to n are written in a circle. 
First, count the k-th number starting from the first one and delete it. 
Then k numbers are counted starting from the next one and the k-th one is removed again, 
and so on. The process stops when one number remains. 
It is required to find the last number.


Brute Force -> O(n^2)

Algo:
n = 3 k = 3
    1 2 3 
    s   ^
    

Assume N elements, need to delete Kth element
start at 1, 
delete the K, 
N-1 elements left
if we would've started at 1 ans would've been K
but we start at some start so ans would be (start + K) % N


Optimized -> O(n)
"""

def josephus(n, k):
    res = 0
    for i in range(2, n + 1):
        res = (res + k) % i
    return res + 1