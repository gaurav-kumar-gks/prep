"""
GRAPH: MINIMUM SPANNING TREE (MST)
"""

"""
Subset of the edges of a connected, undirected graph that connects all the vertices together
, without any cycles, and with minimum possible total edge weight.

1. MST PROPERTIES:
   - Connects all vertices in the graph
   - Contains no cycles (tree structure)
   - Has minimum total edge weight
   - Unique if all edge weights are distinct
   - Number of edges = |V| - 1 (where V is vertices)

2. MST APPLICATIONS:
   - Network design (minimum cost to connect all nodes)
   - Clustering algorithms
   - Image segmentation
   - Circuit design
   - Transportation networks
   - Water supply networks

3. ALGORITHMS:
   - Kruskal's Algorithm: Sort edges, add minimum weight edges
   - Prim's Algorithm: Grow tree from single vertex
   - Boruvka's Algorithm: Parallel algorithm for MST

4. TIME COMPLEXITIES:
   - Kruskal's: O(E log E) = O(E log V) with union-find
   - Prim's: O(E log V) with binary heap

5. MST PROPERTIES:
   - Cut Property: Minimum weight edge crossing any cut is in MST
   - Cycle Property: Maximum weight edge in any cycle is not in MST
   - Uniqueness: MST is unique if all edge weights are distinct

"""

import heapq
from collections import defaultdict

def kruskal_mst(edges, n):
    """
    Kruskal's algorithm for finding Minimum Spanning Tree.
    
    Algorithm:
    1. Sort all edges in non-decreasing order of weight
    2. Initialize disjoint set with all vertices
    3. For each edge in sorted order:
       - If edge doesn't create cycle, add to MST
       - Union the vertices connected by edge
    4. Return MST edges
    
    Time Complexity: O(E log E) = O(E log V)
    Space Complexity: O(V)
    """
    res = []
    ds = DisjointSet()
    for i in range(n):
        ds.make_set(i)
    sorted_edges = sorted(edges, key=lambda x: x[2])
    for u, v, weight in sorted_edges:
        if ds.find(u) != ds.find(v):
            res.append((u, v, weight))
            ds.union(u, v)
            if len(res) == n - 1:
                break
    return res


def prim_mst(graph, start=0):
    """
    Prim's algorithm for finding Minimum Spanning Tree.
    
    Algorithm:
    1. Start with single vertex
    2. Add minimum weight edge that connects tree to new vertex
    3. Repeat until all vertices are included
    4. Use priority queue to find minimum weight edge efficiently
    
    Time Complexity: O(E log V) with binary heap
    Space Complexity: O(V)
    """
    res = []
    vis = [False] * n
    pq = []  # Priority queue: (weight, u, v)
    
    vis[start] = True
    for v, weight in graph[start]:
        heapq.heappush(pq, (weight, start, v))
    
    while pq and len(res) < n - 1:
        weight, u, v = heapq.heappop(pq)
        if vis[u] and vis[v]:
            continue
        res.append((u, v, weight))
        new_vertex = v if not vis[v] else u
        vis[new_vertex] = True
        for neighbor, edge_weight in graph[new_vertex]:
            if not vis[neighbor]:
                heapq.heappush(pq, (edge_weight, new_vertex, neighbor))
    return res

"""

1. KRUSKAL'S ALGORITHM:
   - Advantages: Simple, works with disconnected graphs
   - Disadvantages: Requires sorting all edges
   - Best for: Sparse graphs, when edges are pre-sorted

2. PRIM'S ALGORITHM:
   - Advantages: Works well with dense graphs
   - Disadvantages: More complex implementation
   - Best for: Dense graphs, when starting vertex is known

4. CHOOSING THE RIGHT ALGORITHM:
   - Sparse graphs (E ≈ V): Kruskal's
   - Dense graphs (E ≈ V²): Prim's
   - Parallel processing: Boruvka's
   - Online algorithms: Prim's with priority queue
"""
