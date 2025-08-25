"""
DP PATTERNS
"""


"""
DP IDENTIFICATION

When approaching any problem, ask these questions:

1. Optimal Substructure: Can optimal solution be built from optimal solutions of subproblems?
2. Overlapping Subproblems: Are the same subproblems solved multiple times?
3. Choice/Decision Making: At each step, do you have multiple choices to make?
4. State Dependencies: Does current state depend on previous states?

Key Questions to Ask:
- What are the base cases?
- Can I break this into smaller, similar subproblems?
- What information do I need to remember at each step?
- How does the current state relate to previous states?
"""

"""
================================================================================
1. 1D DYNAMIC PROGRAMMING
================================================================================

Pattern Summary:
State can be represented by a single dimension (usually an index or position). 
The current state depends on a fixed number of previous states. 
This pattern is characterized by linear progression where each step builds upon previous steps

Foundation Problem: Fibonacci Sequence
Problem: Calculate the nth Fibonacci number where F(n) = F(n-1) + F(n-2)

DP Identification Framework Analysis:
- Optimal Substructure: Yes, F(n) = F(n-1) + F(n-2) - optimal solution built from smaller subproblems
- Overlapping Subproblems: Yes, F(3) is calculated multiple times in F(5), F(6), etc.
- Choice/Decision Making: No direct choices, but recursive calls to previous states
- State Dependencies: Current state depends on exactly two previous states

State Dependencies in Problem Terms:
Consider we need to find the nth Fibonacci number:

1) Check if we know the (n-1)th Fibonacci number: F(n-1)
2) Check if we know the (n-2)th Fibonacci number: F(n-2)
3) If both are known, F(n) = F(n-1) + F(n-2)

The current Fibonacci number depends on the two immediately preceding Fibonacci numbers.

================================================================================
STANDARD PROBLEMS
================================================================================

1. Climbing Stairs
   Problem: Count ways to climb n stairs with 1 or 2 steps
   Solution: dp[i] = dp[i-1] + dp[i-2]. Base case: dp[0] = 1, dp[1] = 1. Each step can be reached from previous step or two steps back.

2. House Robber
   Problem: Maximum money from robbing houses without adjacent
   Solution: dp[i] = max(dp[i-1], dp[i-2] + nums[i]). Choose between robbing current house + money from i-2, or skip current house and take money from i-1.

3. Longest Increasing Subsequence (LIS)
   Problem: Find longest strictly increasing subsequence
   Solution: dp[i] = max(dp[j] + 1) for all j < i where nums[j] < nums[i]. For each element, find maximum LIS ending at previous smaller elements.

4. Coin Change
   Problem: Minimum coins needed to make amount
   Solution: dp[amount] = min(dp[amount - coin] + 1) for all coins. For each amount, try using each coin and take minimum.

5. Maximum Subarray Sum (Kadane's)
   Problem: Maximum sum of contiguous subarray
   Solution: dp[i] = max(nums[i], dp[i-1] + nums[i]). Either start new subarray at current element or extend previous subarray.
"""

"""
================================================================================
2. 2D DYNAMIC PROGRAMMING
================================================================================

Pattern Summary:
State requires two dimensions to represent (typically two indices or positions). 
The current state depends on multiple previous states, often from different directions (left, up, diagonal). 
This pattern is common in problems involving two sequences, matrices, or grid-based problems where decisions at each position depend on multiple previous positions.

Foundation Problem: Longest Common Subsequence (LCS)
Problem: Find the length of longest common subsequence between two strings

DP Identification Framework Analysis:
- Optimal Substructure: Yes, LCS of strings[0:i] and strings[0:j] depends on smaller substrings
- Overlapping Subproblems: Yes, same substrings compared multiple times
- Choice/Decision Making: Yes, at each position choose to match or skip
- State Dependencies: Current state depends on three previous states

State Dependencies in Problem Terms:
Consider string a and b and we need to find out the answer for longest common subsequence 
ending with current char of a and current char of b:

1) Check if we remove current char of a and b, does the remaining string a and remaining 
   string b match: a[i-1] matches with b[j-1] 
2) Check if we remove current char of a does the remaining string a matches with b: 
   a[i-1] matches with b[j]
3) Check if we remove current char of b does the remaining string b matches with a: 
   a[i] matches with b[j-1]

================================================================================
2D DYNAMIC PROGRAMMING - STANDARD PROBLEMS
================================================================================

1. Longest Common Subsequence (LCS)
   Problem: Length of longest common subsequence between two strings
   Solution: 
   Try skipping character from either string or include if characters match.
   dp[i][j] = max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1] + 1 if match). 

2. Edit Distance
   Problem: Minimum operations to transform one string to another
   Solution: 
   Try delete, insert, or replace operations.
   dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1] + (0 if match else 1)). 

3. Longest Palindromic Subsequence
   Problem: Longest palindrome that is a subsequence
   Solution: 
   Try including/excluding characters from either end.
   dp[i][j] = max(dp[i+1][j], dp[i][j-1], dp[i+1][j-1] + 2 if match). 

4. Interleaving String
   Problem: Check if string is interleaving of two others
   Solution: 
   Check if current character comes from either string.
   dp[i][j] = (dp[i-1][j] and s1[i-1] == s3[i+j-1]) or (dp[i][j-1] and s2[j-1] == s3[i+j-1]). 

5. Distinct Subsequences
   Problem: Count distinct subsequences of string
   Solution: 
   Count ways to form subsequence with/without current character.
   dp[i][j] = dp[i-1][j] + (dp[i-1][j-1] if s[i-1] == t[j-1]). 

"""

"""
================================================================================
3. KNAPSACK PROBLEMS
================================================================================

Pattern Summary:
Resource allocation with constraints where you must make binary decisions (include/exclude) 
for each item while respecting capacity limits. 
The state typically represents the current capacity and items considered, 
with transitions based on whether to include or exclude the current item. 

Foundation Problem: 0/1 Knapsack
Problem: Given items with weights and values, find maximum value that can be obtained 
with weight limit W, each item can be used at most once

DP Identification Framework Analysis:
- Optimal Substructure: Yes, optimal solution for weight W uses optimal solutions for smaller weights
- Overlapping Subproblems: Yes, same weight capacities solved multiple times
- Choice/Decision Making: Yes, binary choice to include/exclude each item
- State Dependencies: Current state depends on previous items and smaller weights

State Dependencies in Problem Terms:
Consider we need to find maximum value for weight limit W using first i items:

1) Check if we exclude the current item: maximum value using first (i-1) items with weight W
2) Check if we include the current item: maximum value using first (i-1) items with 
   weight (W - weight[i]) plus value of current item
3) Choose the maximum of these two options

The current optimal value depends on the optimal value with one less item and either 
the same weight or reduced weight.

================================================================================
KNAPSACK PROBLEMS - STANDARD PROBLEMS
================================================================================

1. 0/1 Knapsack
   Problem: Maximum value with weight limit, each item once
   Solution: 
   Choose between including or excluding current item.
   dp[i][w] = max(dp[i-1][w], dp[i-1][w-wt[i]] + val[i]). 

2. Unbounded Knapsack
   Problem: Maximum value with weight limit, items unlimited
   Solution: 
   Same as 0/1 but can reuse items.
   dp[w] = max(dp[w-wt[i]] + val[i]) for all items. 

3. Subset Sum
   Problem: Check if subset exists with given sum
   Solution: 
   Check if sum can be achieved with/without current element.
   dp[i][sum] = dp[i-1][sum] or dp[i-1][sum-nums[i]]. 

4. Partition Equal Subset Sum
   Problem: Divide array into two equal parts
   Solution: 
   Find subset with sum = total/2 using subset sum approach. If total is odd, impossible.

5. Target Sum
   Problem: Ways to assign +/- to make target sum
   Solution: 
   Convert to subset sum. Find subsets with sum = (total + target)/2.
"""

"""
================================================================================
4. TREE DYNAMIC PROGRAMMING
================================================================================

Pattern Summary:
State at each node depends on the states of its children. 
post-order traversal to ensure children are processed before parents. 
decisions at each node are influenced by the optimal solutions in subtrees.

Foundation Problem: Binary Tree Maximum Path Sum
Problem: Find the maximum path sum in a binary tree where path can start and end at any node

DP Identification Framework Analysis:
- Optimal Substructure: Yes, max path through node depends on max paths through children
- Overlapping Subproblems: Yes, same subtrees processed multiple times
- Choice/Decision Making: Yes, choose to include/exclude children paths
- State Dependencies: Current node depends on left and right child states

State Dependencies in Problem Terms:
Consider we need to find maximum path sum starting from current node:

1) Check if we only use current node value: node.val
2) Check if we extend path through left child: node.val + maxPath(left)
3) Check if we extend path through right child: node.val + maxPath(right)
4) Check if we create a path through current node connecting both children: 
   node.val + maxPath(left) + maxPath(right)

The current node's maximum path depends on the maximum paths that can be achieved 
through its left and right children.

================================================================================
TREE DYNAMIC PROGRAMMING - STANDARD PROBLEMS
================================================================================

1. Binary Tree Maximum Path Sum
   Problem: Maximum path sum in binary tree
   Solution: 
   For each node, calculate max path through node = node.val + max(0, left) + max(0, right). 
   Return max path starting from node = node.val + max(0, left, right).

2. House Robber III
   Problem: Rob houses in binary tree structure
   Solution: 
   For each node, return [rob_this_node, skip_this_node]. 
   rob = node.val + skip_left + skip_right, skip = max(rob_left, skip_left) + max(rob_right, skip_right).

3. Unique Binary Search Trees
   Problem: Count different BST structures
   Solution: 
   Each value can be root, left subtree has i-1 nodes, right has n-i nodes.
   dp[n] = sum(dp[i-1] * dp[n-i]) for i from 1 to n. 

4. Diameter of Binary Tree
   Problem: Longest path between any two nodes
   Solution: 
   For each node, diameter = max_depth_left + max_depth_right. Update global max diameter.

5. Sum Root to Leaf Numbers
   Problem: Sum of all root-to-leaf paths
   Solution: 
   Pass current number to children. If leaf, add to sum. current = current*10 + node.val.
"""

"""
================================================================================
5. STATE COMPRESSION DP
================================================================================

Pattern Summary:
State compression DP uses bit manipulation to represent complex states efficiently, 
typically when dealing with permutations, combinations, or visited/unvisited states. 
The state is encoded as a bitmask where each bit represents a binary decision or state. 

Foundation Problem: Traveling Salesman Problem (TSP)
Problem: Find minimum cost to visit all cities exactly once and return to starting city

DP Identification Framework Analysis:
- Optimal Substructure: Yes, optimal tour depends on optimal subtours
- Overlapping Subproblems: Yes, same city combinations solved multiple times
- Choice/Decision Making: Yes, choose next city to visit
- State Dependencies: Current state depends on visited cities and current position

State Dependencies in Problem Terms:
Consider we need to find minimum cost to visit remaining cities from current city:

1) Check if we visit city j next: minimum cost to visit remaining cities (excluding j) 
   starting from city j, plus cost from current city to j
2) Try this for all unvisited cities j
3) Choose the minimum cost among all possible next cities

The current minimum cost depends on the minimum costs achievable from all possible 
next cities, given the set of cities already visited.

================================================================================
STATE COMPRESSION DP - STANDARD PROBLEMS
================================================================================

1. Traveling Salesman Problem (TSP)
   Problem: Minimum cost to visit all cities once
   Solution: dp[mask][pos] = min(dp[mask|(1<<next)][next] + cost[pos][next]). Try visiting each unvisited city next.

2. Hamiltonian Path
   Problem: Check if path visits all vertices once
   Solution: dp[mask][pos] = any(dp[mask|(1<<next)][next] for all unvisited next). Check if any unvisited vertex can be next.

3. N-Queens II
   Problem: Count valid N-queens configurations
   Solution: Use bitmasks for columns, diagonals. Try placing queen in each row, update masks, recurse.

4. Sudoku Solver
   Problem: Solve Sudoku puzzle
   Solution: Use bitmasks for rows, columns, boxes. Try each valid number in empty cell, update masks.

5. Graph Coloring
   Problem: Minimum colors to color graph
   Solution: dp[mask][colors] = can color vertices in mask with given colors. Try different color assignments.
"""

"""
================================================================================
6. DIGIT DYNAMIC PROGRAMMING
================================================================================

Pattern Summary:
Number generation, counting, or validation where decisions are made digit by digit. 
The state typically includes the 
current position, 
constraints (like digit usage), 
and whether the number being built is tight (equal to the upper bound). 
This pattern is essential for range queries and number theory problems.

Foundation Problem: Count Numbers with Unique Digits
Problem: Count numbers with unique digits in range [0, n]

DP Identification Framework Analysis:
- Optimal Substructure: Yes, count for position i depends on counts for position i-1
- Overlapping Subproblems: Yes, same digit patterns repeated
- Choice/Decision Making: Yes, choose digit at each position
- State Dependencies: Current state depends on previous digits and constraints

State Dependencies in Problem Terms:
Consider we need to count valid numbers at position i with given constraints:

1) Check if we place digit d at position i: count of valid numbers for position (i+1) 
   with updated digit mask and tight constraint
2) Try this for all valid digits d (not used before and within bounds)
3) Sum up all valid choices

The current count depends on the counts achievable by placing each valid digit at 
the current position and recursively counting the remaining positions.

================================================================================
DIGIT DYNAMIC PROGRAMMING - STANDARD PROBLEMS
================================================================================

1. Count Numbers with Unique Digits
   Problem: Numbers without repeating digits
   Solution: dp[pos][mask][tight] = count of valid numbers. 
   - Try each unused digit (not in mask), update mask, check tight constraint

2. Numbers At Most N Given Digit Set
   Problem: Count valid numbers in range
   Solution: dp[pos][mask][tight] = count.
   - Try each digit from allowed set, update tight constraint based on comparison with n

3. Sum of Digits
   Problem: Sum of digits in range [L, R]
   Solution: dp[pos][sum][tight] = sum of all numbers.
   - For each digit, add digit value to sum, recurse with updated tight constraint

4. Beautiful Arrangement
   Problem: Permutations with position constraints
   Solution: dp[mask][pos] = count of valid arrangements.
   - Try each unused number at current position if it satisfies the beautiful arrangement condition

5. Non-negative Integers without Consecutive Ones
   Problem: Count numbers without consecutive 1s
   Solution: dp[pos][prev] = count.
   - If prev=1, can only use 0 (to avoid consecutive 1s). If prev=0, can use 0 or 1.
"""

"""
================================================================================
7. INTERVAL DYNAMIC PROGRAMMING
================================================================================

Pattern Summary:
Interval DP involves solving problems where the state represents an interval or range, 
and the solution for an interval depends on solutions for smaller subintervals. 
The key insight is to try all possible ways of splitting the interval and choose the optimal one. 
This pattern is common in problems involving optimal ordering, matrix operations, and range-based optimizations.

Foundation Problem: Burst Balloons
Problem: Find maximum coins by bursting balloons in optimal order

DP Identification Framework Analysis:
- Optimal Substructure: Yes, optimal bursting order for interval depends on optimal orders for subintervals
- Overlapping Subproblems: Yes, same intervals solved multiple times
- Choice/Decision Making: Yes, choose which balloon to burst last in interval
- State Dependencies: Current interval depends on all possible splits

State Dependencies in Problem Terms:
Consider we need to find maximum coins from bursting balloons in range [i, j]:

1) Check if we burst balloon k last in this range: maximum coins from bursting 
   balloons [i, k-1] plus maximum coins from bursting balloons [k+1, j] plus 
   coins from bursting k (nums[i-1] * nums[k] * nums[j+1])
2) Try this for all possible k in range [i, j]
3) Choose the maximum among all possible last balloons

The current maximum coins depend on the maximum coins achievable from all possible 
ways of splitting the interval by choosing different last balloons to burst.

================================================================================
INTERVAL DYNAMIC PROGRAMMING - STANDARD PROBLEMS
================================================================================

1. Burst Balloons
   Problem: Maximum coins by bursting balloons in optimal order
   Solution: 
   Try bursting balloon k last.
   dp[i][j] = max(dp[i][k-1] + dp[k+1][j] + nums[i-1]*nums[k]*nums[j+1]). 

2. Stone Game
   Problem: Optimal game strategy for stone piles
   Solution: 
   Choose maximum advantage from either end.
   dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1]). 

3. Minimum Cost to Merge Stones
   Problem: Optimal merging strategy
   Solution: 
   Try all possible split points.
   dp[i][j] = min(dp[i][k] + dp[k+1][j] + sum[i:j+1]). 

4. Palindrome Partitioning
   Problem: Minimum cuts for palindromes
   Solution: 
   Find minimum cuts for each position.
   dp[i] = min(dp[j] + 1) for all j < i where s[j:i] is palindrome. 

5. Matrix Chain Multiplication
   Problem: Optimal parenthesization
   Solution: 
   Try all split points.
   dp[i][j] = min(dp[i][k] + dp[k+1][j] + dims[i-1]*dims[k]*dims[j]). 
"""

"""
================================================================================
8. STRING DYNAMIC PROGRAMMING
================================================================================

Pattern Summary:
String DP involves problems where the state represents positions in one or more strings, 
and decisions are made character by character. 
The state typically includes 
current positions in the strings and any constraints or conditions that need to be maintained. 
This pattern is essential for string matching, parsing, and transformation problems.

Foundation Problem: Regular Expression Matching
Problem: Check if string matches pattern with '.' and '*' wildcards

DP Identification Framework Analysis:
- Optimal Substructure: Yes, matching at position i depends on matching at previous positions
- Overlapping Subproblems: Yes, same string-pattern combinations repeated
- Choice/Decision Making: Yes, multiple ways to match at each position
- State Dependencies: Current state depends on previous string and pattern positions

State Dependencies in Problem Terms:
Consider we need to check if string[0:i] matches pattern[0:j]:

1) Check if current characters match and previous parts match: 
   string[i-1] matches pattern[j-1] and string[0:i-1] matches pattern[0:j-1]
2) Check if pattern has '*' and we can skip current pattern character: 
   string[0:i] matches pattern[0:j-2]
3) Check if pattern has '*' and we can repeat previous character: 
   string[0:i-1] matches pattern[0:j] and string[i-1] matches pattern[j-2]

The current matching state depends on the matching states achievable by different 
interpretations of the current pattern character and previous matching results.

================================================================================
STRING DYNAMIC PROGRAMMING - STANDARD PROBLEMS
================================================================================

1. Regular Expression Matching
   Problem: Match string with pattern containing . and *
   Solution: 
   Handle different pattern cases.
   dp[i][j] = dp[i+1][j+1] if match, or dp[i][j+2] if *, or dp[i+1][j] if * matches. 

2. Wildcard Matching
   Problem: Match string with wildcard pattern
   Solution: 
   Handle * as matching multiple characters.
   dp[i][j] = dp[i+1][j+1] if match or ?, or dp[i+1][j] if *. 

3. Longest Palindromic Substring
   Problem: Longest palindrome substring
   Solution: 
   Expand from center or use DP table.
   dp[i][j] = dp[i+1][j-1] and s[i] == s[j]. 

4. Word Break
   Problem: Check if string can be segmented into dictionary words
   Solution: 
   Check if prefix can be broken and suffix is in dict.
   dp[i] = any(dp[j] and s[j:i] in dict) for all j < i. 

5. Decode Ways
   Problem: Count ways to decode string
   Solution: 
   Check validity of current and previous digits.
   dp[i] = dp[i-1] (if valid single digit) + dp[i-2] (if valid two digits). 
"""

"""
================================================================================
9. OPTIMIZATION PROBLEMS
================================================================================

Pattern Summary:
Optimization DP involves finding the best possible solution (maximum, minimum, or optimal) among all feasible solutions. 
The state typically represents the optimal value achievable up to the current position, 
and transitions involve making local optimal decisions that contribute to the global optimum. 
This pattern is fundamental for all optimization problems.

Foundation Problem: Maximum Subarray Sum (Kadane's)
Problem: Find maximum sum of contiguous subarray

DP Identification Framework Analysis:
- Optimal Substructure: Yes, max sum ending at i depends on max sum ending at i-1
- Overlapping Subproblems: Yes, same ending positions considered multiple times
- Choice/Decision Making: Yes, choose to extend previous sum or start fresh
- State Dependencies: Current state depends on previous state and current element

State Dependencies in Problem Terms:
Consider we need to find maximum sum ending at position i:

1) Check if we extend the previous maximum sum: maximum sum ending at (i-1) plus current element
2) Check if we start a new subarray: just the current element
3) Choose the maximum of these two options

The current maximum sum depends on whether it's better to extend the previous 
maximum sum or start fresh with the current element.

================================================================================
OPTIMIZATION PROBLEMS - STANDARD PROBLEMS
================================================================================

1. Maximum Subarray Sum
   Problem: Maximum sum of contiguous subarray
   Solution: 
   Extend previous sum or start new subarray.
   dp[i] = max(nums[i], dp[i-1] + nums[i]). 

2. Maximum Product Subarray
   Problem: Maximum product of contiguous subarray
   Solution: 
   Keep track of max and min products. 
   max_curr = max(nums[i], max_prev*nums[i], min_prev*nums[i]).

3. Best Time to Buy and Sell Stock
   Problem: Maximum profit from stock trading
   Solution: 
   Keep track of minimum price seen so far.
   dp[i] = max(dp[i-1], prices[i] - min_price). 

4. Jump Game
   Problem: Check if can reach end with jump constraints
   Solution: 
   Check if any previous position can reach current.
   dp[i] = any(dp[j] and j + nums[j] >= i) for all j < i. 

5. Maximum Sum Circular Subarray
   Problem: Maximum sum in circular array
   Solution: 
   Find max subarray sum and min subarray sum. 
   Answer is max(max_sum, total - min_sum).
"""

"""
================================================================================
10. COUNTING PROBLEMS
================================================================================

Pattern Summary:
Counting DP involves problems where the goal is to count the number of ways to achieve a certain outcome 
rather than finding an optimal value. 
The state represents the count of valid solutions up to the current position,
and transitions involve adding up all valid ways to reach the current state. 
This pattern is essential for combinatorial problems.

Foundation Problem: Unique Paths
Problem: Count number of unique paths from top-left to bottom-right in grid

DP Identification Framework Analysis:
- Optimal Substructure: Yes, paths to (i,j) depend on paths to (i-1,j) and (i,j-1)
- Overlapping Subproblems: Yes, same positions reached multiple times
- Choice/Decision Making: Yes, choose to come from left or up
- State Dependencies: Current position depends on two adjacent positions

State Dependencies in Problem Terms:
Consider we need to count paths to reach position (i, j):

1) Check if we come from the cell above: number of paths to reach (i-1, j)
2) Check if we come from the cell to the left: number of paths to reach (i, j-1)
3) Total paths = paths from above + paths from left

The current path count depends on the path counts achievable by reaching the two 
adjacent cells that can lead to the current position.

================================================================================
COUNTING PROBLEMS - STANDARD PROBLEMS
================================================================================

1. Unique Paths
   Problem: Count paths from top-left to bottom-right in grid
   Solution: 
   Can come from above or left.
   dp[i][j] = dp[i-1][j] + dp[i][j-1]. 

2. Unique Paths II
   Problem: Count paths with obstacles
   Solution: 
   Skip obstacles.
   dp[i][j] = 0 if obstacle, else dp[i-1][j] + dp[i][j-1]. 

3. Decode Ways
   Problem: Count ways to decode string
   Solution: 
   Check validity of current and previous digits.
   dp[i] = dp[i-1] (if valid single digit) + dp[i-2] (if valid two digits). 

4. Perfect Squares
   Problem: Minimum perfect squares to sum to n
   Solution: 
   Try each perfect square.
   dp[i] = min(dp[i - j*j] + 1) for all j*j <= i. 

5. Combination Sum IV
   Problem: Count combinations that sum to target
   Solution: 
   Try using each number.
   dp[target] = sum(dp[target - num]) for all nums. 
"""

"""
================================================================================
11. GAME THEORY DP
================================================================================

Pattern Summary:
Game theory DP involves solving problems where two or more players make 
optimal moves in turn, and the goal is to determine winning strategies or outcomes. 
The state represents the current game configuration, and transitions involve 
analyzing all possible moves and their outcomes. 
This pattern is essential for competitive games and strategic decision-making problems.

Foundation Problem: Nim Game
Problem: Determine if current player can win given pile of stones

DP Identification Framework Analysis:
- Optimal Substructure: Yes, winning state depends on winning states after moves
- Overlapping Subproblems: Yes, same pile configurations analyzed multiple times
- Choice/Decision Making: Yes, choose which pile and how many stones to take
- State Dependencies: Current state depends on states after all possible moves

State Dependencies in Problem Terms:
Consider we need to determine if current player can win with given piles:

1) Check if we take k stones from pile i: can the opponent win with the resulting pile configuration?
2) Try this for all possible moves (all piles and all possible stone counts)
3) If any move leaves opponent in losing position, current player can win

The current winning state depends on whether any move can force the opponent into 
a losing state.

================================================================================
GAME THEORY DP - STANDARD PROBLEMS
================================================================================

1. Nim Game
   Problem: Determine if current player can win
   Solution: 
   If any move leaves opponent in losing position, current player wins.
   dp[piles] = any(not dp[piles_after_move]). 

2. Stone Game
   Problem: Two-player optimal stone game
   Solution: 
   Choose maximum advantage.
   dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1]). 

3. Predict the Winner
   Problem: Predict game winner
   Solution: 
   Return if dp[0][n-1] >= 0.
   dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1]). 

4. Can I Win
   Problem: Check if can force win
   Solution: 
   If any move leads to opponent loss, can win.
   dp[mask] = any(not dp[mask|(1<<i)] for unused i). 

5. Cat and Mouse
   Problem: Optimal strategy for cat and mouse game
   Solution: 
   Consider all possible moves.
   dp[mouse][cat][turn] = 1 if mouse wins, 2 if cat wins, 0 if draw. 
"""

"""
================================================================================
12. DP ON GRAPHS
================================================================================

Pattern Summary:
DP on graphs involves solving problems where the 
state represents positions or configurations in a graph, and 
transitions follow graph edges or paths. 
The state typically includes the current node/position and any additional constraints or conditions. 
This pattern is essential for graph traversal, path finding, and network optimization problems.

Foundation Problem: Shortest Path with Constraints
Problem: Find shortest path from source to destination with additional constraints

DP Identification Framework Analysis:
- Optimal Substructure: Yes, shortest path to node depends on shortest paths to adjacent nodes
- Overlapping Subproblems: Yes, same node configurations visited multiple times
- Choice/Decision Making: Yes, choose which adjacent node to visit next
- State Dependencies: Current state depends on states of adjacent nodes

State Dependencies in Problem Terms:
Consider we need to find shortest path to node v with constraint C:

1) Check if we come from adjacent node u: shortest path to u with constraint C' plus edge weight (u,v)
2) Try this for all adjacent nodes u that satisfy the constraint
3) Choose the minimum among all valid paths

The current shortest path depends on the shortest paths achievable through all valid adjacent nodes.

================================================================================
DP ON GRAPHS - STANDARD PROBLEMS
================================================================================

1. Shortest Path with Alternating Colors
   Problem: Shortest path with color constraints
   Solution: dp[node][color] = shortest path to node with last edge of given color. Use BFS with state.

2. Number of Ways to Arrive at Destination
   Problem: Count shortest paths
   Solution: dp[node] = sum(dp[prev]) for all previous nodes in shortest path. Use Dijkstra with counting.

3. Minimum Cost to Reach Destination in Time
   Problem: Shortest path with time limit
   Solution: dp[node][time] = minimum cost to reach node within time. Use BFS with time dimension.

4. Graph Coloring
   Problem: Minimum colors for graph
   Solution: dp[mask][colors] = can color vertices in mask with given colors. Try different color assignments.

5. Hamiltonian Path
   Problem: Path visiting all vertices once
   Solution: dp[mask][pos] = can reach position with visited vertices in mask. Check if any unvisited vertex can be next.
"""

"""
================================================================================
13. CONVEX HULL OPTIMIZATION
================================================================================

Pattern Summary:
transition function can be optimized using geometric properties. 
It involves maintaining a convex hull of lines or functions to achieve O(n) or O(n log n) complexity instead of O(n²). 
This pattern is essential for complex optimization problems with specific mathematical properties.

Foundation Problem: Maximum Subarray Sum with Constraints
Problem: Find maximum sum subarray with additional constraints that allow convex hull optimization

DP Identification Framework Analysis:
- Optimal Substructure: Yes, optimal solution depends on optimal solutions of subproblems
- Overlapping Subproblems: Yes, same subproblems solved multiple times
- Choice/Decision Making: Yes, choose optimal transition from multiple candidates
- State Dependencies: Current state depends on previous states with geometric optimization

State Dependencies in Problem Terms:
Consider we need to find maximum value at position i:

1) Check if we use transition from position j: value at j plus some function of (i-j)
2) This can be optimized using convex hull of lines representing different j values
3) Find the line that gives maximum value at position i

The current maximum value depends on the optimal line from the convex hull that maximizes the transition function.

===============================================================================
CONVEX HULL OPTIMIZATION - STANDARD PROBLEMS
================================================================================

1. Maximum Subarray Sum with Constraints
   Problem: Maximum sum with geometric optimization
   Solution: Use convex hull trick. Maintain lines in convex hull, query for maximum at each position.

2. Optimal Partition of String
   Problem: Partition string optimally
   Solution: dp[i] = max(dp[j] + cost(j,i)) for j < i. Use convex hull to optimize transition.

3. Maximum Profit in Job Scheduling
   Problem: Schedule jobs for maximum profit
   Solution: dp[i] = max(dp[j] + profit[i]) for all j where end[j] <= start[i]. Use convex hull optimization.

4. Longest Arithmetic Subsequence
   Problem: Longest arithmetic sequence
   Solution: dp[i][diff] = dp[j][diff] + 1 where nums[i] - nums[j] = diff. Use convex hull for optimization.

5. Maximum Sum of 3 Non-Overlapping Subarrays
   Problem: Three non-overlapping subarrays
   Solution: Use convex hull to find optimal positions for three subarrays.
"""

"""
================================================================================
14. DIVIDE AND CONQUER OPTIMIZATION
================================================================================

Pattern Summary:
Reduces complexity by exploiting the monotonicity of optimal decision points. 
It involves finding the optimal split point for each state by using the fact that optimal decisions are monotonic. 
This pattern is essential for problems where the optimal choice follows a predictable pattern.

Foundation Problem: Optimal Binary Search Tree
Problem: Construct optimal binary search tree with minimum expected search cost

DP Identification Framework Analysis:
- Optimal Substructure: Yes, optimal tree depends on optimal subtrees
- Overlapping Subproblems: Yes, same subtrees solved multiple times
- Choice/Decision Making: Yes, choose optimal root for each subtree
- State Dependencies: Current state depends on optimal splits with monotonic optimization

State Dependencies in Problem Terms:
Consider we need to find optimal root for subtree [i, j]:

1) Check if we use k as root: optimal cost for left subtree [i, k-1] plus optimal cost for right subtree [k+1, j] plus root cost
2) The optimal k is monotonic, allowing binary search optimization
3) Find the optimal k using divide and conquer approach

The current optimal cost depends on the optimal split point found using monotonicity properties.

================================================================================
DIVIDE AND CONQUER OPTIMIZATION - STANDARD PROBLEMS
================================================================================

1. Optimal Binary Search Tree
   Problem: Construct optimal BST
   Solution: dp[i][j] = min(dp[i][k-1] + dp[k+1][j] + freq_sum[i:j+1]). Use monotonicity to optimize split point.

2. Burst Balloons
   Problem: Optimal balloon bursting order
   Solution: dp[i][j] = max(dp[i][k-1] + dp[k+1][j] + nums[i-1]*nums[k]*nums[j+1]). Use divide and conquer optimization.

3. Stone Game V
   Problem: Optimal stone game with divide and conquer
   Solution: Use monotonicity of optimal split points to reduce complexity.

4. Minimum Cost to Merge Stones
   Problem: Optimal merging with monotonicity
   Solution: Use divide and conquer optimization to find optimal split points efficiently.

5. Palindrome Partitioning III
   Problem: Partition string into k palindromes
   Solution: Use divide and conquer optimization for optimal partition points.
"""

"""
================================================================================
15. BACKTRACKING WITH DP
================================================================================

Pattern Summary:
Backtracking with DP involves memoizing backtracking solutions to avoid recalculating the same states multiple times. 
The state represents the current configuration in the backtracking search, and transitions follow the backtracking choices. This pattern is essential for problems that naturally fit backtracking but have overlapping subproblems.

Foundation Problem: N-Queens with Memoization
Problem: Count valid N-queens configurations using backtracking with DP optimization

DP Identification Framework Analysis:
- Optimal Substructure: Yes, valid configurations depend on valid partial configurations
- Overlapping Subproblems: Yes, same partial configurations explored multiple times
- Choice/Decision Making: Yes, choose valid positions for queens
- State Dependencies: Current state depends on states after placing queens

State Dependencies in Problem Terms:
Consider we need to count valid configurations with queens placed up to row i:

1) Check if we place queen at column j in row i: count valid configurations for remaining rows with updated constraints
2) Try this for all valid columns j that don't conflict with existing queens
3) Sum up all valid configurations

The current count depends on the counts achievable by placing queens at all valid positions in the current row.

================================================================================
BACKTRACKING WITH DP - STANDARD PROBLEMS
================================================================================

1. N-Queens
   Problem: Count valid N-queens configurations
   Solution: Use backtracking with memoization. Track columns, diagonals with bitmasks.

2. Sudoku Solver
   Problem: Solve Sudoku with memoization
   Solution: Use backtracking with memoization. Track rows, columns, boxes with bitmasks.

3. Word Search II
   Problem: Find words in grid with trie
   Solution: Use backtracking with trie. Explore grid, match against trie, memoize states.

4. Palindrome Partitioning
   Problem: Partition into palindromes
   Solution: Use backtracking with memoization. Try all valid palindrome partitions.

5. Combination Sum
   Problem: Find combinations that sum to target
   Solution: Use backtracking with memoization. Try each number, recurse with reduced target.
"""

"""
================================================================================
16. DP WITH BINARY SEARCH
================================================================================

Pattern Summary:
Binary search to optimize problems where the answer lies in a search space that can be efficiently narrowed down. 
The DP state represents the feasibility of achieving a certain target value, 
and binary search is used to find the optimal target. 
This pattern is essential for optimization problems with monotonic properties.

Foundation Problem: Split Array Largest Sum
Problem: Split array into k subarrays to minimize the maximum subarray sum

DP Identification Framework Analysis:
- Optimal Substructure: Yes, optimal split depends on optimal splits of subarrays
- Overlapping Subproblems: Yes, same subarray splits solved multiple times
- Choice/Decision Making: Yes, choose where to split the array
- State Dependencies: Current state depends on feasibility of achieving target sum

State Dependencies in Problem Terms:
Consider we need to check if target sum T is achievable with k splits:

1) Check if we can split at position i: can we achieve target T for remaining elements with (k-1) splits?
2) Try this for all valid split positions i
3) Use binary search to find the minimum achievable target sum

The current feasibility depends on whether any split point allows achieving the target with remaining splits.

================================================================================
DP WITH BINARY SEARCH - STANDARD PROBLEMS
================================================================================

1. Split Array Largest Sum
   Problem: Split array to minimize maximum sum
   Solution: Binary search on answer. For each target, use DP to check if achievable.

2. Capacity To Ship Packages Within D Days
   Problem: Minimum capacity to ship packages
   Solution: Binary search on capacity. For each capacity, check if packages can be shipped in D days.

3. Koko Eating Bananas
   Problem: Minimum eating speed
   Solution: Binary search on speed. For each speed, check if can eat all bananas in H hours.

4. Minimum Number of Days to Make m Bouquets
   Problem: Minimum days for bouquets
   Solution: Binary search on days. For each day, check if can make m bouquets.

5. Egg Drop Problem
   Problem: Minimum attempts to find critical floor
   Solution: Binary search on attempts. For each attempt count, use DP to check if feasible.
"""
