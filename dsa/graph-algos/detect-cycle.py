"""
GRAPH: DETECTING CYCLE
"""

"""
Undirected graph
BFS - if we visit a node which is already visited but isn't a parent
DFS - if we detect back edge to a visited ancestor
"""

from collections import deque

def detect_cycle_undirected_dfs(graph, n):
    vis = set()
    for node in range(n):
        if node not in vis:
            vis.add(node)
            if has_cycle_dfs_undirected(graph, node, vis, -1):
                return True
            # if has_cycle_bfs_undirected(graph, node, vis, -1):
                # return True
    return False

def has_cycle_dfs_undirected(graph, node, vis, par):
    for ne in graph[node]:
        if ne not in vis:
            vis.add(ne)
            if has_cycle_dfs_undirected(graph, ne, vis, node):
                return True
        elif par != ne:
            return True
    return False

def has_cycle_bfs_undirected(graph, node, vis, par):
    q = deque([(node, par)])
    while q:
        n, par = q.popleft()
        for ne in graph[n]:
            if ne not in vis:
                q.append((ne, n))
                vis.add(ne)
            elif ne != par:
                return True
    return False


    
"""
Directed graph
BFS - if topo sort doesn't visit all the nodes
DFS - detect back edge using the recursion stack
"""

def detect_cycle_directed_dfs(graph, n):
    recStack = set()
    vis = set()
    for i in range(n):
        if i not in vis:
            if has_cycle_dfs_directed(i, graph, vis, recStack):
                return True
    return False

def has_cycle_dfs_directed(node, graph, vis, recStack):
    vis.add(node)
    recStack.add(node)
    for ne in graph[node]:
        if ne not in vis:
            if has_cycle_dfs_directed(ne, graph, vis, recStack):
                return True
        elif ne in recStack:
                return True
    recStack.remove(node)
    return False
