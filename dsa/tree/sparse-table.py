"""
SPARSE TABLE
"""

"""
1. WHAT IS A SPARSE TABLE?
   - A data structure for efficient range queries on static arrays
   - Precomputes answers for all possible ranges of power-of-2 lengths
   - Supports O(1) range queries for idempotent operations
   - Cannot handle updates efficiently (static data structure)

2. STRUCTURE:
   - 2D array where st[i][j] represents the result for range [j, j + 2^i - 1]
   - i: power of 2 representing the length (2^i)
   - j: starting index of the range (0 to n-1)

3. OPERATIONS:
   - Build: O(n log n) - Precompute all power-of-2 ranges
   - Query: O(1) - For idempotent operations (min, max, gcd, lcm)
   - Query: O(log n) - For non-idempotent operations (sum, product)
   - Update: O(n log n) - Rebuild entire table (inefficient)

4. APPLICABLE OPERATIONS:
   - Idempotent: min, max, gcd, lcm (can overlap ranges)
   - Non-idempotent: sum, product (cannot overlap ranges)
   - Must be associative: f(a, f(b, c)) = f(f(a, b), c)

5. WHEN TO USE:
   - Static arrays with frequent range queries
   - Idempotent operations (min, max, gcd, lcm)
   - No or very few updates (updates require rebuilding the entire table in O(n log n) time)
"""

import math
from typing import List, Callable, Any


class SparseTable:
    
    def __init__(self, arr, operation):
        self.arr = arr
        self.n = len(arr)
        self.operation = operation
        self.op_func, self.identity = self._get_operation(operation)
        self.is_idempotent = operation in ["min", "max", "gcd", "lcm"]
        self._build()
    
    def _get_operation(self, operation):
        operations = {
            "min": (min, float('inf')),
            "max": (max, float('-inf')),
            "gcd": (math.gcd, 0),
            "lcm": (lambda x, y: (x * y) // math.gcd(x, y) if x and y else 0, 1),
            "sum": (lambda x, y: x + y, 0),
            "product": (lambda x, y: x * y, 1)
        }
        return operations.get(operation, (min, float('inf')))
    
    def _build(self):
        n = len(self.arr)
        k = math.log2(n) + 1
        st = [[self.identity] * n for _ in range(k + 1)]
        
        for j in range(n):
            st[0][j] = self.arr[j]
        
        for i in range(1, k + 1):
            length = 1 << i
            prev_length = length >> 1
            for j in range(self.n - length + 1):
                st[i][j] = self.op_func(st[i - 1][j], st[i - 1][j + prev_length])
        self.st = st
    
    def query_idempotent(self, l, r):
        # we can overlap the two ranges
        i = math.log2(r - l + 1)
        return self.op_func(self.st[i][l], self.st[i][r - (1 << i) + 1])
    
    def query_non_idempotent(self, l, r):
        # find subranges of length 2^a, 2^b, 2^c untill we reach end
        result = self.identity
        while l <= r:
            i = math.log2(r - l + 1)
            result = self.op_func(result, self.st[i][l])
            l += (1 << i)
        return result
    
    def update(self, index, value):
        self.arr[index] = value
        self._build()  # Rebuild entire table

"""
PERFORMANCE CHARACTERISTICS:
===========================

| Operation                 | Time Complexity | Space Complexity | Best For        |
|---------------------------|-----------------|------------------|-----------------|
| Build                     | O(n log n)      | O(n log n)       | Static arrays   |
| RangeQuery(Idempotent)    | O(1)            | O(n log n)       | Fast queries    |
| RangeQuery(Non-idempotent)| O(log n)        | O(n log n)       | General ops     |
| Point Update              | O(n log n)      | O(n log n)       | Avoid updates   |
| Range Update              | Not supported   | N/A              | Use segment tree|

"""
