"""
GRAPH: STRONGLY CONNECTED COMPONENTS (SCC)
"""

"""
A Strongly Connected Component (SCC) is a portion of a directed graph in which there is a path
from each vertex to every other vertex. In other words, it's a maximal strongly connected subgraph.

1. SCC PROPERTIES:
   - Every vertex in SCC is reachable from every other vertex in SCC
   - SCCs are maximal (no larger strongly connected subgraph contains them)
   - Graph can be decomposed into SCCs
   - SCCs form a DAG (Directed Acyclic Graph) when contracted

2. SCC APPLICATIONS:
   - Compiler optimization (finding loops)
   - Social network analysis (finding communities)
   - Circuit design (finding feedback loops)
   - Web page ranking (PageRank algorithm)
   - Dependency resolution

3. ALGORITHMS:
   - Kosaraju's Algorithm: Two DFS passes
   - Tarjan's Algorithm: Single DFS with stack
   - Gabow's Algorithm: Variant of Tarjan's

4. TIME COMPLEXITIES:
   - Kosaraju's: O(V + E)
   - Tarjan's: O(V + E)
   - Gabow's: O(V + E)

5. SCC PROPERTIES:
   - Transitive closure within each SCC
   - No cycles between different SCCs
   - Condensation graph is a DAG

"""

from collections import defaultdict, deque


def kosaraju_scc(graph, n):
    """
    Kosaraju's algorithm for finding Strongly Connected Components.
    
    Algorithm:
    1. Perform DFS on original graph and push vertices to stack in finishing order
    2. Reverse all edges in the graph
    3. Perform DFS on reversed graph in order of stack (top to bottom)
    4. Each DFS tree in step 3 is an SCC
    
    Time Complexity: O(V + E)
    Space Complexity: O(V + E)
    """
    vis = [False] * n
    stack = []
    
    def dfs(node):
        for nei in graph[node]:
            if not vis[nei]:
                vis[nei] = True
                dfs(nei)
        stack.append(node)
    
    for node in range(n):
        if not vis[node]:
            vis[node] = True
            dfs(node)
    
    vis[:] = [False] * n
    sccs = []
    
    def reversed_graph_dfs(node):
        scc.append(node)
        for nei in reversed_graph[node]:
            if not vis[nei]:
                vis[node] = True
                reversed_graph_dfs(nei)
    
    while stack:
        node = stack.pop()
        if not vis[node]:
            scc = []
            vis[node] = True
            reversed_graph_dfs(node)
            sccs.append(scc)
    return sccs
