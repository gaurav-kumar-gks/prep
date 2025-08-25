"""
GRAPH: ARTICULATION POINTS & BRIDGES 
"""

"""
Articulation Points (Cut Vertices) and Bridges (Cut Edges) are critical components in graph theory
that help identify vulnerable points in network connectivity.

1. ARTICULATION POINTS (CUT VERTICES):
   - A vertex whose removal increases the number of connected components
   - Critical points in network infrastructure
   - Single points of failure in communication networks

2. BRIDGES (CUT EDGES):
   - An edge whose removal increases the number of connected components
   - Critical connections in network topology
   - Single connections whose failure disconnects the network

3. PROPERTIES:
   - Articulation points and bridges are only defined for undirected graphs
   - A bridge is always part of a path between two articulation points
   - Removing an articulation point can create multiple bridges
   - A graph without articulation points is called biconnected
"""

def find_articulation_points(graph):
    """
    Find only articulation points using Tarjan's algorithm.
    
    Time Complexity: O(V + E)
    Space Complexity: O(V)
    """
    n = len(graph)
    time = [(-1, -1)] * n
    ap = set()
    t = 0
    
    def dfs(node, parent=None):
        time[node] = (t, t)
        t += 1
        children = 0
        for nei in graph[node]:
            if time[nei][0] == -1:
                children += 1
                dfs(nei, node)
                time[node] = (time[node][0], min(time[node][1], time[nei][1]))
                if parent is not None and time[nei][1] >= time[node][0]:
                    ap.add(node)
            elif nei != parent:
                time[node] = (time[node][0], min(time[node][1], time[nei][0]))
                
        if parent is None and children > 1:
            ap.add(node)
    
    for i in range(n):
        if time[i][0] == -1:
            dfs(i)
    
    return list(articulation_points)


def find_bridges(graph):
    """
    Find only bridges using Tarjan's algorithm
    
    Algo:
    1. Perform DFS and maintain discovery time and low value for each vertex
    2. Low value = minimum of discovery time of vertex and discovery times of all back edges
    3. Bridge: edge (u,v) where low value of v > discovery time of u
    
    Time Complexity: O(V + E)
    Space Complexity: O(V)
    """
    n = len(graph)
    time = [(-1, -1)] * n
    bridges = []
    t = 0
    
    def dfs(node, parent=None):
        time[node] = (t, t)
        for nei in graph[node]:
            if time[nei][0] == -1:
                dfs(nei, node)
                time[node] = (time[node][0], min(time[node][1], time[nei][1]))
                if time[nei][1] > time[node][0]:
                    bridges.append((node, nei))
            elif nei != parent:  # Back edge
                time[node] = (time[node][0], min(time[node][1], time[nei][0]))
    
    for i in range(n):
        if time[i][0] == -1:
            dfs(i)
    return bridges
