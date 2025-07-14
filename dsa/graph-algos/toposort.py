"""
Topological sort

Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices 
such that for every directed edge u -> v, vertex u comes before v in the ordering. 

Time complexity: O(V + E)

Applications:
- Scheduling tasks
- Data serialization
- Dependency resolution
- Instruction scheduling
- Resolving symbol dependencies
- etc.
"""

from collections import deque, defaultdict

def topological_sort_dfs(graph, n):
    """doesn't have cycle detection"""
    
    res = []
    vis = set()
    
    def dfs(node):
        vis.add(node)
        for ne in graph[node]:
            if ne not in vis:
                dfs(ne) 
        res.append(node)
    
    for i in range(n):
        if i not in vis:
            dfs(i)
    return res[::-1]


def topological_sort_bfs(lis, n):
    """has cycle detection"""
    graph = defaultdict(list)
    ind = [0] * n
    for u, v in lis:
        graph[u].append(v)
        ind[v] += 1
    
    q = deque([i for i in range(n) if ind[i] == 0])
    res = []
    
    while q:
        u = q.popleft()
        res.append(u)
        for v in graph[u]:
            ind[v] -= 1
            if ind[v] == 0:
                q.append(v)
    
    return res[::-1] if len(res) == n else []