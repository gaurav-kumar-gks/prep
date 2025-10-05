"""
DISJOINT SET (UNION-FIND)
"""

"""
Tracks a set of elements partitioned into a number of disjoint (non-overlapping) subsets. 
It provides near-constant-time operations for merging sets and finding which set an element belongs to.


1. DISJOINT SET PROPERTIES:
   - Each element belongs to exactly one disoint set
   - MAKE_SET(x): Create a new set containing element x
   - UNION(x, y): Merge sets containing x and y
   - FIND(x): Find which set element x belongs to

2. REPRESENTATION:
   - Each set is represented as a tree
   - Each element points to its parent
   - Root element points to itself (self-loop)
   - Root represents the set identifier

3. OPTIMIZATIONS:
   - Union by Rank: Attach smaller tree to root of larger tree
   - Path Compression: Make each node point directly to root during find

4. TIME COMPLEXITIES:
   - Without optimizations: O(n) worst case
   - With Union by Rank: O(log n)
   - With Path Compression: O(log n)
   - With both optimizations: O(α(n)) amortized (inverse Ackermann function)
   - α(n) < 5 for all practical purposes (near-constant time)
"""

class DisjointSet:
    
    def __init__(self):
        self.parent = {}  # Maps element to its parent
        self.rank = {}    # Maps element to its rank (height of subtree)
        self.count = 0    # Number of disjoint sets
    
    def make_set(self, x):
        if x not in self.parent:
            self.parent[x] = x
            self.rank[x] = 0
            self.count += 1
    
    def find(self, x):
        """
        Find the representative (root) of the set containing element x.
        
        1. Follow parent pointers until reaching root
        2. Apply path compression: make all nodes on path point directly to root
        
        Time Complexity: O(α(n)) amortized
        Space Complexity: O(1)
        """
        if x not in self.parent:
            raise ValueError(f"Element {x} not found in any set")
        # Path compression: make all nodes on path point directly to root
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]
    
    def union(self, x, y):
        """
        Merge the sets containing elements x and y.
        
        Algorithm:
        1. Find roots of both sets
        2. If roots are same, sets are already merged
        3. Otherwise, use union by rank:
           - Attach smaller rank tree to root of larger rank tree
           - If ranks are equal, increment rank of new root
        
        Time Complexity: O(α(n)) amortized
        Space Complexity: O(1)
        """
        root_x = self.find(x)
        root_y = self.find(y)
        if root_x == root_y:
            return
        if self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        elif self.rank[root_y] < self.rank[root_x]:
            self.parent[root_y] = root_x
        else:
            self.parent[root_y] = root_x
            self.rank[root_x] += 1
        self.count -= 1  # Decrease number of sets

"""
APPLICATIONS

1. CONNECTED COMPONENTS IN GRAPH
2. CYCLE DETECTION IN UNDIRECTED GRAPH
3. KRUSKAL'S ALGORITHM
4. NUMBER OF ISLANDS
5. ACCOUNTS MERGE
"""