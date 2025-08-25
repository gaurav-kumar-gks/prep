"""
GRAPH: BIPARTITE GRAPH
"""

"""
A Bipartite Graph is a graph whose vertices can be divided into two disjoint and independent sets
U and V such that every edge connects a vertex in U to one in V. In other words, the graph can be
colored using only two colors such that no two adjacent vertices have the same color.

1. BIPARTITE GRAPH PROPERTIES:
   - Vertices can be partitioned into two independent sets
   - All edges go between the two sets (no edges within same set)
   - A graph is bipartite if and only if it contains no odd-length cycles
   - A graph is bipartite if and only if it is 2-colorable
   - A graph is bipartite if and only if its chromatic number is 2

3. APPLICATIONS:
   - Job assignment problems
   - Marriage problem (stable matching)
   - Network flow problems
   - Resource allocation
   - Scheduling problems
   - Social network analysis

4. ALGORITHMS:
   - BFS/DFS with 2-coloring
   - Maximum matching (Hungarian algorithm)
   - Minimum vertex cover
   - Maximum independent set

5. TIME COMPLEXITIES:
   - Bipartite checking: O(V + E)
   - Maximum matching: O(VE) or O(E√V)
   - Minimum vertex cover: O(VE)

"""

from collections import deque

def is_bipartite(graph):
    """
    Check if a graph is bipartite using BFS with 2-coloring.
    
    Time Complexity: O(V + E)
    Space Complexity: O(V)
    """
    n = len(graph)
    color = [-1] * n  # -1: unvisited, 0: color1, 1: color2

    for start in range(n):
        if color[start] != -1:
            continue
        q = deque([start])
        color[start] = 0
        while q:
            node = q.popleft()
            for nei in graph[node]:
                if color[nei] == -1:
                    color[nei] = 1 - color[node]
                    q.append(nei)
                elif color[nei] == color[node]:
                    return False
    return True


def is_bipartite_dfs(graph):
    """
    Check if a graph is bipartite using DFS with 2-coloring.
    
    Algorithm:
    1. Use DFS to traverse the graph
    2. Assign colors to vertices (0 and 1)
    3. If adjacent vertices have same color, graph is not bipartite
    4. If no conflicts found, graph is bipartite
    
    Time Complexity: O(V + E)
    Space Complexity: O(V)
    """
    n = len(graph)
    color = [-1] * n
    def dfs(node, current_color):
        color[node] = current_color
        for neighbor in graph[node]:
            if color[neighbor] == -1:
                if not dfs(neighbor, 1 - current_color):
                    return False
            elif color[neighbor] == current_color:
                return False
        return True
    for start in range(n):
        if color[start] == -1:
            if not dfs(start, 0):
                return False
    return True


"""
APPLICATIONS

1. CHECK IF GRAPH IS BIPARTITE
   - Use BFS/DFS with 2-coloring
   - Time: O(V + E)

2. MAXIMUM MATCHING IN BIPARTITE GRAPH
   - Use Hungarian algorithm or Ford-Fulkerson
   - Time: O(VE) or O(E√V)

3. MINIMUM VERTEX COVER
   - König's theorem: size of minimum vertex cover equals size of maximum matching
   - Time: O(VE)

4. MAXIMUM INDEPENDENT SET
   - Complement of minimum vertex cover
   - Time: O(VE)

5. JOB ASSIGNMENT PROBLEM
   - Assign jobs to workers optimally
   - Time: O(VE)

6. MARRIAGE PROBLEM
   - Find stable matching between two sets
   - Time: O(V²)

7. NETWORK FLOW PROBLEMS
   - Convert to maximum flow problem
   - Time: O(VE²) with Ford-Fulkerson
"""
