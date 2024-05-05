"""
Graphs
"""
from collections import deque

"""
Representations
"""

"""
Adjacency Matrix

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
"""

class GraphMatrix:
    def __init__(self, num_of_vertices):
        self.matrix = [[0]*num_of_vertices for _ in range(num_of_vertices)]
        self.num_of_vertices = num_of_vertices

    def add_edge(self, u, v):
        self.matrix[u][v] = 1
        self.matrix[v][u] = 1  # For undirected graph

    def print_graph(self):
        for i in range(self.num_of_vertices):
            for j in range(self.num_of_vertices):
                print(self.matrix[i][j], end=" ")
            print()


"""
Adjacency List

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
"""


class GraphList:
    def __init__(self, num_of_vertices):
        self.adj_list = [[] for _ in range(num_of_vertices)]
        self.num_of_vertices = num_of_vertices

    def add_edge(self, u, v):
        self.adj_list[u].append(v)
        self.adj_list[v].append(u)  # For undirected graph

    def print_graph(self):
        for i in range(self.num_of_vertices):
            print(f"{i} -> {' -> '.join(map(str, self.adj_list[i]))}")


class GraphDict:
    def __init__(self):
        self.graph = {}

    def add_edge(self, u, v):
        if u not in self.graph:
            self.graph[u] = [v]
        else:
            self.graph[u].append(v)
        
        # For undirected graph
        if v not in self.graph:
            self.graph[v] = [u]
        else:
            self.graph[v].append(u)

    def print_graph(self):
        for node in self.graph:
            print(f"{node} -> {' -> '.join(map(str, self.graph[node]))}")
            
    def __getitem__(self, node):
        return self.graph[node]

graph = GraphDict()
graph.add_edge(0, 1)
graph.add_edge(0, 2)


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


def dfs_iterative(s):
    """
    Time complexity: O(V + E)
    Space complexity: O(V)
    """
    stack = [s]
    visited = {s}
    color = {s: "white"}
    time = 0
    start_time = {s: time}
    end_time = {}

    while stack:
        s = stack.pop()
        color[s] = "gray"
        time += 1

        all_neighbors_visited = True
        for i in graph[s]:
            if color.get(i) == "white":
                stack.append(i)
                visited.add(i)
                color[i] = "white"
                time += 1
                start_time[i] = time
                all_neighbors_visited = False

        if all_neighbors_visited:
            color[s] = "black"
            end_time[s] = time

    return visited, start_time, end_time
                
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
Types of Graphs

- Directed graph: A graph in which all the edges are directed.
- Undirected graph: A graph in which all the edges are undirected.

- Weighted graph: A graph in which all the edges have weights.
- Unweighted graph: A graph in which all the edges have no weights.

- Cyclic graph: A graph that has at least one cycle.
- Acyclic graph: A graph that has no cycles.

- Connected graph: A graph in which there is a path between every pair of vertices
- Disconnected graph: A graph in which there is no path between at least one pair of vertices

- Complete graph: A graph in which there is an edge between every pair of vertices
- Incomplete graph: A graph in which there is no edge between at least one pair of vertices

- Bipartite graph: A graph whose vertices can be divided into two disjoint sets such that every edge connects two vertices from different sets.
- Non-bipartite graph: A graph that is not bipartite.

- Sparse graph: A graph in which the number of edges is much less than the number of vertices
- Dense graph: A graph in which the number of edges is close to the number of vertices squared

- Tree: A connected, undirected, acyclic graph
- Forest: A collection of trees
"""


"""
Types of Edges

- Tree edge: An edge in a DFS tree
- Back edge: An edge that connects a vertex to an ancestor in a DFS tree
- Forward edge: An edge that connects a vertex to a descendant in a DFS tree
- Cross edge: An edge that connects a vertex to a node that is neither an ancestor nor a descendant in a DFS tree

consider the following graph:


    A -> B -> C
    |    |    |
    v    v    v
    D -> E -> F
    |    |    |
    v    v    v
    G -> H -> I
    
The edges in the graph are as follows:

- Tree edges: A -> B, B -> C, B -> E, D -> E, E -> F, D -> G, G -> H, H -> I
- Back edges: E -> B, G -> D, H -> G, I -> H
- Forward edges: B -> D, C -> F, E -> H, F -> I
- Cross edges: A -> C, A -> F, A -> I, D -> F, D -> I, G -> I
"""

"""

Shortest Path Algorithms: Understand Dijkstra's algorithm, Bellman-Ford algorithm, and Floyd-Warshall algorithm. Know when to use each one and their time complexities.

Minimum Spanning Tree Algorithms: Understand Prim's and Kruskal's algorithms.

Topological Sorting: This is especially important for problems involving scheduling and determining dependencies.

Strongly Connected Components: Understand Tarjan’s algorithm and Kosaraju's algorithm.

Network Flow: Understand the Ford-Fulkerson algorithm and the concept of maximum flow.

Graph Coloring and Bipartite Checking: These are useful for problems involving scheduling and partitioning.

Cycle Detection: Be able to detect cycles in both directed and undirected graphs.

Articulation Points and Bridges: Understand how to find articulation points and bridges in a graph.

Graph Isomorphism and Matching: Understand the concepts, though detailed algorithms are less frequently asked.

Advanced Topics: Depending on the role, you might also need to understand more advanced topics like Eulerian paths and circuits, Hamiltonian cycles, Planar graphs, etc.
"""