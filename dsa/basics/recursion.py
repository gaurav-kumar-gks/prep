"""
RECURSION
"""


"""
Recursion 

A function that calls itself is called a recursive function.

Recursion tree - a tree that shows the recursive calls made by a recursive function.
Branches: Choices made at each level of the recursion tree are called branches.
Branching factor: The number of branches at each level is called the branching factor.
Depth: The number of levels in the recursion tree is called the depth.

Time Complexity: O(b^d) where b: branching factor and d: depth

Things to keep in mind:
1. base condition: check for smallest/largest input to the recursion function
2. hypothesis: assume that the function works for smaller/larger input
3. recursive call: call the recursion function with the new input
4. return value: return the value from the recursion function
"""

"""
Head recursion:

Some computation is being done after the recursive call
The function needs to keep track of the intermediate results for each recursive call, which can lead to higher memory usage


def head_recursive(n):
    if n == 0:
        return
    head_recursive(n-1)
    print(n)
    
head_recursive(5)
    head_recursive(4)
        head_recursive(3)
            head_recursive(2)
                head_recursive(1)
                    head_recursive(0)
                    print(1)
                print(2)
            print(3)
        print(4)
    print(5)

"""

"""
Tail recursion:

Computations are done before the recursive call and the result of the recursive call is returned directly.

The function does not need to keep track of the intermediate results, which can lead to lower memory usage. 

In some programming languages, 
tail recursion can be optimized by the compiler, 
making it more efficient than head recursion.
In Python, tail recursion is not optimized by the compiler (one more python L)


def tail_recursive(n, accumulator=1):
    if n == 0:
        return accumulator
    else:
        return tail_recursive(n-1, n * accumulator)
        

tail_recursive(5, 1)
    tail_recursive(4, 5)
        tail_recursive(3, 20)
            tail_recursive(2, 60)
                tail_recursive(1, 120)
                    tail_recursive(0, 120)
"""

"""
Backtracking is a general algorithm for finding all (or some) solutions to some computational problems,

e.g. finding all the ways to arrange n queens on an n*n chessboard so that no queen can attack any other queen on the chessboard.

Blueprint:

1. identify that we need an algo where we need to exhaust some space, e.g. generate all, all permutations, all combinations, etc.
2. check if the problem can be broken into smaller problems
3. check how many choices we have at each level of the recursion tree
4. with choices there may be some condition to take each chioce
5. now with some constraint take a choice & recurse
6. after that undo that choice


choice -> take / do -> recurse -> untake / undo
"""



"""
Dynamic Programming

How to identify DP problems ?
1. Recursive solution exists 
2. Overlapping Subproblems
3. Optimal Substructure


Overlapping Subproblems: 
A problem has Overlapping Subproblems if it can be broken into subproblems which are reused several times.

Optimal Substructure:
A problem has Optimal Substructure if an optimal solution can be constructed from optimal solutions of its subproblems.


Best way to solve DP problems is to use recursion tree.

always try to follow these
1. try to identify if problem has recursive soln.
2. then try to draw recursive tree
3. check for optimal substructure & optimal subproblems
4. then implement the recursive soln.
5. then memoize the same by storing the results  - dp table - top down approach
6. after that implement the iterative soln. - bottom up approach
"""
