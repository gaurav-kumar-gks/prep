"""
Graph: connected-components
"""

"""
Connected Components

A connected component of an undirected graph is a subgraph in which 
every two vertices are connected to each other by a path(s), 
and which is connected to no other vertices outside the subgraph.
"""

"""
Finding Connected Components: DFS or BFS

Start from any node and do BFS / DFS, maintain a visited set. All nodes visited in this iteration form a connected component.    
Then just check for the other unvisited set and run DFS / BFS again. 
And so on, until all the nodes are visited.

O(n + m)
"""

# def connected_components(graph):
#     visited = set()
#     result = []
#     for node in graph:
#         if node not in visited:
#             dfs(node, visited)
#             result.append(connected_component)
#     return result

"""
Finding bridges in a graph

Bridge: An edge in an undirected graph is a bridge if removing it disconnects the graph.

Algorithm:

And edge (u, v) is a bridge if and only if there is no other way to reach u from v or vice versa after removing the edge (u, v).

"""
