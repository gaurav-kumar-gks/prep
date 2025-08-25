"""
RECURSION
"""


"""
Recursion 

Definition: A function that calls itself is called a recursive function.
Can a problem be broken into smaller problems & ans of smaller problems can be used to solve the bigger problem ?


Recursion tree - a tree that shows the recursive calls made by a recursive function.
Branches: Choices made at each level of the recursion tree are called branches.
Branching factor: The number of branches at each level is called the branching factor.
Depth: The number of levels in the recursion tree is called the depth.

Time Complexity: O(b^d) where b is tDefinition: he branching factor and d is the depth of the recursion tree.


How to build recursion tree ?
1. Draw the base case
2. Draw the recursive case
3. Draw the recursive calls
4. Draw the return values
"""

"""
1. base condition: check for smallest/largest input to the recursion function
this is very important, if we don't have a base condition, the recursion will go on forever
2. Hypothesis: assume that the function works for smaller/larger input
3. recursive call: call the recursion function with the new input
4. return value: return the value from the recursion function
"""

"""
Head recursion:

Head recursion is a type of recursion where the recursive call is the first statement in the function. 
In other words, the function performs some operations before making the recursive call.

In head recursion, the recursive call is made first and the computation is done after the recursive call. 
This means that the function needs to keep track of the intermediate results for each recursive call, 
which can lead to higher memory usage


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

Tail recursion is a type of recursion where the recursive call is the last operation in the function. 
This means that all the computations are done before the recursive call and the result of the recursive call is returned directly.

In tail recursion, the computation is done before the recursive call and the result of the recursive call is returned directly. 
This means that the function does not need to keep track of the intermediate results, 
which can lead to lower memory usage. 

In some programming languages, tail recursion can be optimized by the compiler, making it more efficient than head recursion.
In Python, tail recursion is not optimized by the compiler.


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
                    return 120
                return 120
            return 120
        return 120
    return 120
"""


"""
Factorial of a number is the product of all the integers from 1 to that number.

e.g. 
5! = 5 * 4 * 3 * 2 * 1 = 120
4! = 4 * 3 * 2 * 1 = 24

factorial(5) = 5*factorial(4) = 5*4*factorial(3) = 5*4*3*factorial(2) = 5*4*3*2*factorial(1) = 5*4*3*2*1*factorial(0) = 5*4*3*2*1*1 = 120

Recursion tree would look like -> 
factorial(5)
    factorial(4)
        factorial(3)
            factorial(2)
                factorial(1)
                    factorial(0)


Base case: factorial(0) = 1
Hypothesis: factorial(n) gives me the factorial of n
Recursion step: ans = n * factorial(n-1)
Return value: return ans
"""


def factorial(n):
    """
    Time Complexity: O(n)
    Space Complexity: O(n)
    """
    if n <= 1:
        # base condition
        return n
    return n * factorial(n - 1)  # recursive step


def tail_recursion_factorial(n, ans=1):
    """
    Time Complexity: O(n)
    Space Complexity: O(n)
    """
    if n <= 1:
        return ans
    return tail_recursion_factorial(n - 1, ans * n)


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

remember to create a choice diagram, how many choices do we have at each level of the recursion tree
"""



"""
Dynamic Programming

Definition: Dynamic Programming is an algorithmic paradigm that solves a given complex problem by 
a) breaking it into subproblems 
b) stores the results of subproblems to avoid computing the same results again.

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
