"""
shortest path algorithms
"""

from collections import deque

"""
Undirected graph

- BFS
- O(V+E)
"""
def undirected_graph_shortest_path(graph, s, n):
    q = deque([s])
    dist = [-1] * n
    dist[s] = 0
    while q:
        n = q.popleft()
        for ne in graph[n]:
            if dist[ne] == -1:
                dist[ne] = dist[n] + 1
                q.append(ne)
                
    return -1

"""
Using toposort
- DAG
- O(V+E)
"""

def toposort(graph, n):
    res = []
    vis = set()
    def dfs(node):
        vis.add(node)
        for ne, _ in graph[node]:
            if ne not in vis:
                dfs(ne)
        res.append(node)
    
    for i in range(n):
        if i not in vis:
            dfs(i)
    
    topo = res[::-1]
    
    dis = [float('inf')] * n
    start = 0
    dis[start] = 0
    
    for node in topo:
        for ne, w in graph[node]:
            if dis[ne] > dis[node] + w:
                dis[ne] = dis[node] + w
    print(dis)
        

"""
Dijkstra's Algorithm
- doesn't work for negative weight cycle
- O((V+E)logV)
"""

import heapq

def dijkstra(graph, n, start):
    dis = [float('inf')] * n
    pq = [(0, start)]
    pars = {start: None}
    dis[start] = 0
    while pq:
        d, n = heapq.heappop(pq)
        if d > dis[n]: continue
        for w, ne in graph[n]:
            if dis[ne] > d + w:
                dis[ne] = d + w
                pars[ne] = n
                heapq.heappush(pq, (dis[ne], ne))

"""
Bellman ford
- works for negative cycles
- O(VE)
"""

def bellman_ford(edges, n, start):
    dist = [float('inf')] * n
    dist[start] = 0

    for _ in range(n - 1):
        for u, v, w in edges:
            if dist[u] + w < dist[v]:
                dist[v] = dist[u] + w

    # Check for negative-weight cycle
    for u, v, w in edges:
        if dist[u] + w < dist[v]:
            raise ValueError("Graph contains negative weight cycle")

    return dist

def floyd_warshall(graph, n):
    dist = [[float('inf')] * n for _ in range(n)]
    for u in range(n):
        dist[u][u] = 0
    for u in range(n):
        for v, w in graph[u]:
            dist[u][v] = w

    for k in range(n):
        for i in range(n):
            for j in range(n):
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
    return dist