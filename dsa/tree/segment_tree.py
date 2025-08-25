"""
SEGMENT TREE
"""

"""
1. WHAT IS A SEGMENT TREE?
   - Each node represents a segment/interval of the array
   - Supports efficient range queries and point updates
   - Root represents the entire array [0, n-1]
   - Each internal node has two children representing left and right halves
   - Leaf nodes represent single array elements
   - Usually implemented as an array (1-based indexing)
   - For array of size n, tree size is approximately 4n
   - Node i has children at 2i and 2i+1
   - Parent of node i is at i//2
"""

class SegmentTree:
    """
    Segment Tree implementation for range minimum queries.
    Supports point updates and range queries in O(log n) time.
    """
    
    def __init__(self, arr):
        """
        Initialize segment tree from array.
        Time Complexity: O(n)
        Space Complexity: O(n)
        """
        self.arr = arr
        self.n = len(arr)
        self.tree = [float('inf')] * (4 * self.n)
        if self.n > 0:
            self._build(1, 0, self.n - 1)
    
    def _build(self, node, start, end):
        if start == end:
            self.tree[node] = self.arr[start]
            return
        mid = (start + end) // 2
        self._build(2 * node, start, mid)
        self._build(2 * node + 1, mid + 1, end)
        self.tree[node] = min(self.tree[2 * node], self.tree[2 * node + 1])
    
    def query(self, left, right):
        """
        Query minimum value in range [left, right].
        Time Complexity: O(log n)
        """
        if left < 0 or right >= self.n or left > right:
            raise ValueError(f"Invalid query range: [{left}, {right}]. Valid range is [0, {self.n-1}]")
        return self._query(1, 0, self.n - 1, left, right)
    
    def _query(self, node, start, end, left, right):
        if start > right or end < left:
            return float('inf') # identity element
        if left <= start and end <= right:
            return self.tree[node] # leaf node
        mid = (start + end) // 2
        left_min = self._query(2 * node, start, mid, left, min(mid, right))
        right_min = self._query(2 * node + 1, mid + 1, end, max(mid+1, left), right)
        return min(left_min, right_min) # propogation

    def point_update(self, index, value):
        """
        Update given index with value
        Time Complexity: O(log n)
        """
        if index < 0 or index >= self.n:
            raise ValueError(f"Invalid update index: {index}. Valid range is [0, {self.n-1}]")
        self.arr[index] = value
        self._point_update(1, 0, self.n - 1, index, value)
    
    def _point_update(self, node, start, end, index, value):
        if start == end:
            self.tree[node] = value
            return
        mid = (start + end) // 2
        if index <= mid:
            self._point_update(2 * node, start, mid, index, value)
        else:
            self._point_update(2 * node + 1, mid + 1, end, index, value)
        self.tree[node] = min(self.tree[2 * node], self.tree[2 * node + 1])

    def range_update(self, left, right, value):
        """
        Update range with value (w/o lazy propogation)
        If rebuilding the array: Time Complexity: O(n)
        If doing point update for whole range: Time Complexity: O(nlogn)
        
        """
        if left < 0 or right >= self.n or left > right:
            raise ValueError(f"Invalid update range: [{left}, {right}]. Valid range is [0, {self.n-1}]")
        for i in range(left, right + 1):
            self.arr[i] = value
        self._build(1, 0, self.n - 1)


"""
Examples -
1. Range sum
2. Range Min / Max
3. Range GCD / LCM
4. Range Frequency
5. Range Update
6. Maximum and no. of times it appears: 
    store for a segment (max, n_times) instead of just max
7. Count no. of zeroes, searching for kth zero: 
    store no. of zeroes, for searching kth zero - divide and conquer
8. Finding subarray with max sum: 
    store for a segment (sum, max_presum, max_suffsum, max_subarray_sum)
9. Find smallest number >= x given (l, r, x)
    No modification query:
        store list of elements in sorted way in segment tree (like merge sort algo)
        Memory: O(nlogn) (each element falls in logn segments not all 4n segments)
        Query: O(logn * logn)
    Modification queries:
        store treap in segment tree
"""

# =====================================================================================
# LAZY PROPAGATION SEGMENT TREE (RANGE UPDATES)
# =====================================================================================

class LazySegmentTree:
    """
    Segment Tree with lazy propagation for efficient range updates.
    Supports both point and range updates in O(log n) time.
    """
    
    def __init__(self, arr):
        self.arr = arr
        self.n = len(arr)
        self.tree = [0] * (4 * self.n)
        # lazy array
        # at given index it stores what value needs to be propogated to its children nodes
        # if 0: then nothing needs to be propogated
        self.lazy = [0] * (4 * self.n)
        if self.n > 0:
            self._build(1, 0, self.n - 1)
    
    def _build(self, node, start, end):
        if start == end:
            self.tree[node] = self.arr[start]
            return
        mid = (start + end) // 2
        self._build(2 * node, start, mid)
        self._build(2 * node + 1, mid + 1, end)
        self.tree[node] = self.tree[2 * node] + self.tree[2 * node + 1]
    
    def range_update(self, left, right, value):
        """
        Add value to all elements in range [left, right].
        Args:
            left: Left boundary of update range
            right: Right boundary of update range
            value: Value to add to each element in range
        """
        if left < 0 or right >= self.n or left > right:
            raise ValueError(f"Invalid update range: [{left}, {right}]. Valid range is [0, {self.n-1}]")
        self._range_update(1, 0, self.n - 1, left, right, value)
    
    def _range_update(self, node, start, end, left, right, value):
        # Push down lazy values 
        self._push_down(node, start, end)
        
        # No overlap
        if start > right or end < left:
            return
        
        # Full overlap
        if left <= start and end <= right:
            self.tree[node] += value * (end - start + 1) # update current node
            if start != end: # if not leaf node, then propogate to children
                self.lazy[2 * node] += value
                self.lazy[2 * node + 1] += value
            return
        
        # Partial overlap
        mid = (start + end) // 2
        self._range_update(2 * node, start, mid, left, right, value)
        self._range_update(2 * node + 1, mid + 1, end, left, right, value)
        self.tree[node] = self.tree[2 * node] + self.tree[2 * node + 1]
    
    def _push_down(self, node, start, end):
        """
        Push lazy values down to children.
        """
        if not self.lazy[node]:
            return
        # Update current node
        self.tree[node] += self.lazy[node] * (end - start + 1) 
        # Push to children if not leaf
        if start != end:
            self.lazy[2 * node] += self.lazy[node]
            self.lazy[2 * node + 1] += self.lazy[node]
        self.lazy[node] = 0 # propogated to children, so reset it
    
    def query(self, left, right):
        if left < 0 or right >= self.n or left > right:
            raise ValueError(f"Invalid query range: [{left}, {right}]. Valid range is [0, {self.n-1}]")
        return self._query(1, 0, self.n - 1, left, right)
    
    def _query(self, node, start, end, left, right):
        # Push down lazy values
        self._push_down(node, start, end)
        if start > right or end < left:
            return 0
        if left <= start and end <= right:
            return self.tree[node]
        mid = (start + end) // 2
        left_sum = self._query(2 * node, start, mid, left, right)
        right_sum = self._query(2 * node + 1, mid + 1, end, left, right)
        return left_sum + right_sum
    

"""
COMPLEXITY ANALYSIS
==================

1. TIME COMPLEXITY:
   - Build: O(n)
   - Point Query: O(log n)
   - Range Query: O(log n)
   - Point Update: O(log n)
   - Range Update (with lazy): O(log n)
   - Range Update (without lazy): O(n log n) if doing point updates, O(n) If rebuilding tree

2. SPACE COMPLEXITY:
   - Tree Storage: O(4n)

3. COMPARISON WITH OTHER DATA STRUCTURES:
   
   | Operation      | Array | Sparse Table | Segment Tree |
   |----------------|-------|--------------|--------------|
   | Build          | O(1)  | O(n log n)   | O(n)         |
   | Point Query    | O(1)  | O(1)         | O(log n)     |
   | Range Query    | O(n)  | O(1)         | O(log n)     |
   | Point Update   | O(1)  | O(n log n)   | O(log n)     |
   | Range Update   | O(n)  | O(n log n)   | O(log n)     |

4. WHEN TO USE SEGMENT TREES:
   - Need both range queries and updates
   - Updates are frequent
   - Range queries are common
   - Need flexibility in operations

5. ALTERNATIVES:
   - Sparse Table: For static arrays (no updates)
   - Binary Indexed Tree (Fenwick): For sum queries
   - Sqrt Decomposition: For simple range queries
"""

