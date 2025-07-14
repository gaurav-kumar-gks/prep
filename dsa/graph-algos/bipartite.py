from collections import deque

def is_bipartite(graph):
    n = len(graph)
    color = [-1] * n  # -1: unvisited, 0: color1, 1: color2

    for start in range(n):
        if color[start] == -1:
            q = deque([start])
            color[start] = 0
            while q:
                node = q.popleft()
                for neighbor in graph[node]:
                    if color[neighbor] == -1:
                        color[neighbor] = 1 - color[node]
                        q.append(neighbor)
                    elif color[neighbor] == color[node]:
                        return False  # Same color on both ends → Not bipartite
    return True

                    
 