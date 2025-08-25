"""
GRAPH: TOPOLOGICAL SORT
"""

"""
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
        for nei in graph[node]:
            if nei not in vis:
                vis.add(nei)
                dfs(nei)
        res.append(node)
    
    for node in range(n):
        if node not in vis:
            vis.add(node)
            dfs(node)
    return res[::-1]


def topological_sort_bfs(edges, n):
    """has cycle detection"""
    res = []
    inorder = [0] * n
    graph = defaultdict(list)
    for u, v in edges:
        graph[u].append(v)
        inorder[v] += 1

    q = deque(list(filter(lambda x: inorder[x] == 0, range(n))))
    while q:
        node = q.popleft()
        for nei in graph[node]:
            inorder[nei] -= 1
            res.append(nei)
            if not inorder[nei]:
                q.append(nei)
    return res[::-1] if len(res) == n else []