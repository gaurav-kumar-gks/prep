"""
Single source shortest path algorithms
"""

"""
Dijkstra's Algorithm

Algorithm:
1. Initialize distances of all vertices as infinite.
2. Create an empty set.  We will use this set to keep track of vertices included in shortest path tree.
3. Assign distance value to the source vertex as 0 so that it is picked first.
4. While the set doesn't include all vertices:
    a. Pick a vertex u which is not in the set and has minimum distance value.
    b. Include u to the set.
    c. Update distance value of all adjacent vertices of u.  
    To update the distance values, iterate through all adjacent vertices. 
    For every adjacent vertex v, 
    if the sum of distance value of u (from source) and weight of edge u-v, 
    is less than the distance value of v, then update the distance value of v.

in general T.C. = O(E + V log V) with min heap
and T.C. = O(E + V^2) with adjacency matrix

"""

import heapq

def dijkstra_sparse_graphs(graph, start):
    distances = {node: float('infinity') for node in graph}
    previous_nodes = {node: None for node in graph}
    distances[start] = 0
    queue = [(0, start)]

    while queue:
        current_distance, current_node = heapq.heappop(queue)

        if current_distance > distances[current_node]:
            continue

        for neighbor, weight in graph[current_node].items():
            distance = current_distance + weight

            if distance < distances[neighbor]:
                distances[neighbor] = distance
                previous_nodes[neighbor] = current_node
                heapq.heappush(queue, (distance, neighbor))

    return distances, previous_nodes

from typing import List, Optional

def dijkstra(graph: List[List[int]], src: int) -> List:
    V = len(graph)
    dist = [float('inf')] * V
    dist[src] = 0
    visited = [False] * V

    for _ in range(V):
        min_dist = float('inf')
        min_index = -1
        for i in range(V):
            if not visited[i] and dist[i] < min_dist:
                min_dist = dist[i]
                min_index = i
        
        if min_index == -1:
            break
        
        visited[min_index] = True

        for v, weight in enumerate(graph[min_index]):
            if weight > 0 and not visited[v] and dist[v] > dist[min_index] + weight:
                dist[v] = dist[min_index] + weight

    return dist

