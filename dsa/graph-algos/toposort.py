"""
Topological sort

Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge u -> v, vertex u comes before v in the ordering. Topological Sorting for a graph is not possible if the graph is not a DAG.

Algorithm:
1. Run DFS on the graph.
2. Keep track of the finishing time of each node.
3. Sort the nodes in decreasing order of finishing time.
4. Return the sorted nodes.

Time complexity: O(V + E)

Applications:
- Scheduling tasks
- Data serialization
- Dependency resolution
- Instruction scheduling
- Resolving symbol dependencies
- etc.
"""


# code in python using GraphList representation of graph

from collections import deque

class GraphList:
    def __init__(self, num_of_vertices):
        self.adj_list = [[] for _ in range(num_of_vertices)]
        self.num_of_vertices = num_of_vertices

    def add_edge(self, u, v):
        self.adj_list[u].append(v)

    def print_graph(self):
        for i in range(self.num_of_vertices):
            print(f"{i} -> {' -> '.join(map(str, self.adj_list[i]))}")

    def topological_sort(self):
        def dfs(node, visited, stack):
            visited.add(node)
            for neighbor in self.adj_list[node]:
                if neighbor not in visited:
                    dfs(neighbor, visited, stack)
            stack.append(node)

        visited = set()
        stack = deque()
        for node in range(self.num_of_vertices):
            if node not in visited:
                dfs(node, visited, stack)
        return stack
    
    def topological_sort_iterative(self):
        stack = deque()
        visited = set()
        for node in range(self.num_of_vertices):
            if node not in visited:
                stack.append(node)
                visited.add(node)
                while stack:
                    curr = stack[-1]
                    found = False
                    for neighbor in self.adj_list[curr]:
                        if neighbor not in visited:
                            stack.append(neighbor)
                            visited.add(neighbor)
                            found = True
                            break
                    if not found:
                        stack.pop()
        return stack
    
