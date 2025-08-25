"""
GRAPH: SHORTEST PATH
"""

from collections import deque

"""
Undirected graph

- BFS
- O(V+E)
"""
def undirected_graph_shortest_path(graph, s, n):
    # For multi source BFS 
    # add all sources in q and make dis[source] = 0
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
Dijkstra's Algorithm
- doesn't work for negative weight cycle

Algo:
- Use priority queue to get the node with the smallest distance
- Relax all the edges coming out of the node
- If the distance to the node is smaller than the distance in the priority queue, update the distance and parent
- Repeat until all nodes are visited

- O((V+E)logV)
"""

import heapq

def dijkstra(graph, n, start):
    # Shortest path with condition 
    # may require changing what we store in pq and changing dis
    pq = [(0, start)]
    dis = [float('inf')] * n
    dis[start] = 0
    par = {start: None}
    while pq:
        d, node = heapq.heappop(pq)
        if d > dis[node]: 
            continue
        for nei, w in graph[node]:
            if dis[nei] > d + w:
                dis[nei] = d + w
                par[nei] = node
                heapq.heappush(pq, (dis[nei], nei))
    return dis, par
        

"""
Bellman ford
- works for negative cycles

Algo:
- Relax all the edges n-1 times
    Do this n-1 times because the shortest path can have at most n-1 edges
    if after n-1 iterations, we still can relax an edge, then there is a negative cycle

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

"""
Floyd Warshall
- works for negative cycles
- O(V³)
"""

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

"""
Using toposort
- DAG
- O(V+E)
"""

def toposort(graph, start, n):
    res = []
    vis = set()
    def dfs(node):
        for nei, _ in graph[node]:
            if nei not in vis:
                vis.add(nei)
                dfs(nei)
        res.append(node)
    
    for node in range(n):
        if node not in vis:
            vis.add(node)
            dfs(node)
    
    topo = res[::-1]
    dis = [float('inf')] * n
    dis[start] = 0
    for node in topo:
        for ne, w in graph[node]:
            if dis[ne] > dis[node] + w:
                dis[ne] = dis[node] + w
    print(dis)