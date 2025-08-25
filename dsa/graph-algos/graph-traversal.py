"""
GRAPH
"""

"""
Types of Graphs

- Directed / Undirected
- Weighted / Unweighted
- Cyclic / Acyclic (Tree)
- Complete graph: A graph in which there is an edge between every pair of vertices
- Bipartite graph: A graph whose vertices can be divided into two disjoint sets such that every edge connects two vertices from different sets.
- Sparse graph: A graph in which the number of edges is much less than the number of vertices

Types of Edges
- Tree edge: An edge in a DFS tree
- Back edge: An edge that connects a vertex to an ancestor in a DFS tree
- Forward edge: An edge that connects a vertex to a descendant in a DFS tree
- Cross edge: An edge that connects a vertex to a node that is neither an ancestor nor a descendant in a DFS tree

Types of Representation

1. Adjacency Matrix

e.g. 
    A B C
A   0 1 1
B   1 0 0
C   1 0 0

The adjacency matrix for an undirected graph is symmetric
while the one for a directed graph is not necessarily symmetric.

Cons:
    - Space complexity: O(V^2)
    - Adding a new vertex requires creating a new row and a new column
    - Adding a new edge requires updating two entries
    - Determining whether there is an edge between two vertices requires looking up two entries

class GraphMatrix:
    def __init__(self, num_of_vertices):
        self.matrix = [[0]*num_of_vertices for _ in range(num_of_vertices)]
        self.num_of_vertices = num_of_vertices

    def add_edge(self, u, v):
        self.matrix[u][v] = 1
        self.matrix[v][u] = 1  # For undirected graph


2. Adjacency List

e.g.
A -> B -> C
B -> A
C -> A

Pros:
    - Space complexity: O(V + E)
    - Adding a new vertex requires creating a new entry
    - Adding a new edge requires updating one entry
    - Determining whether there is an edge between two vertices requires looking up one entry

Cons:
    - Determining whether there is an edge between two vertices requires iterating over the list

class GraphDict:
    def __init__(self):
        self.graph = {}

    def add_edge(self, u, v):
        if u not in self.graph:
            self.graph[u] = [v]
        else:
            self.graph[u].append(v)
"""

"""
BFS & DFS
"""
from collections import deque

def bfs(s):
    """
    Time complexity: O(V + E)
    Space complexity: O(V)
    """
    
    visited = set()
    queue = deque([s])
    visited.add(s)
    shortest_distance = {s: 0}
    parent = {s: None}   
    while queue:
        s = queue.popleft()
        print(s, end=" ")
        for i in graph[s]:
            if i not in visited:
                shortest_distance[i] = shortest_distance[s] + 1
                parent[i] = s
                queue.append(i)
                visited.add(i)

    # if we want to find the shortest path from s to a node v
    # we can do the following:
    # path = []
    # while v is not None:
    #     path.append(v)
    #     v = parent[v]
    # path.reverse()
    # return path
    
    # if we want to find the shortest distance from s to a node v
    # we can do the following:
    # return shortest_distance[v]


def dfs(graph, s):
    """
    Time complexity: O(V + E)
    Space complexity: O(V)
    """
    stack = [(s, None)]
    color = {s: "gray"}  # Mark as gray when discovered
    iscycle = False
    time = 0
    times = {s: (time, -1)}
    
    while stack:
        current, par = stack[-1]
        unvisited_neighbor = None
        for neighbor in graph[current]:
            if color.get(neighbor, "white") == "white":
                unvisited_neighbor = neighbor
            elif color.get(neighbor) == "gray" and neighbor != par:
                iscycle = True
        
        if unvisited_neighbor:
            # Process unvisited neighbor
            time += 1
            times[unvisited_neighbor] = (time, -1)
            color[unvisited_neighbor] = "gray"
            # visited.add(unvisited_neighbor)
            stack.append((unvisited_neighbor, current))
        else:
            # All neighbors visited, backtrack
            time += 1
            times[current] = (times[current][0], time)
            color[current] = "black"
            stack.pop()
    
    return color, times
                
def dfs_recursive(s, visited = None):
    """
    Time complexity: O(V + E)
    Space complexity: O(V)
    """
    if not visited:
        visited = set()
    visited.add(s)
    print(s, end=" ")
    for i in graph[s]:
        if i not in visited:
            dfs_recursive(i, visited)

"""
A connected component of an undirected graph is a subgraph in which 
every two vertices are connected to each other by a path(s), 
and which is connected to no other vertices outside the subgraph.
"""

"""
Finding Connected Components: DFS or BFS

Start from any node and do BFS / DFS, maintain a visited set. 
All nodes visited in this iteration form a connected component.    
Then just check for the other unvisited set and run DFS / BFS again. 
And so on, until all the nodes are visited.

O(n + m)
"""