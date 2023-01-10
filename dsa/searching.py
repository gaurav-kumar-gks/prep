"""
Various searching algorithms
"""

"""
Binary Search
Time Complexity: O(log n)

What: Searching for an answer maybe a fixed value or some value over a monotonous range
Identification: Given any monotonous function f(x) such that we know it's strictly increaing or decreasing
Idea: To split the search space in half

Finding l & r: 
Check the bounds to find l & r 
e.g. in binary search in ans first try to find what could be the min & max ans -> that'll give the l and r

Finding mid:
mid = (l + r) / 2 -> this may overflow
mid = l + (r - l) // 2 -> this is lower mid
mid = low + (high - low + 1) // 2 -> this is upper mid
mid = (low + high) >> 2

note: choice of mid & shrinking logic has to work together 
s.t. the search space is reduced by half and 1 element is always excluded
also, always check the case of <=3 elements
"""


def check(idx):
    """function to check if the given index / mid satisfies some condition in binary search"""
    pass


def binary_search(l, r):
    """
    Finding the first l s.t. cond(l) == True
    search space kinda like -> F F F F F T T T T
    """
    while l < r:
        mid = l + (r - l) // 2  # lower mid
        if check(mid):
            r = mid
        else:
            l = mid + 1
    # here return l or l-1 depending on the question
    return l


def binary_search_on_floats(l, r):
    """
    Sometimes when we want to do binary search on ans, ans may not be an integer,
    we may be searching on real number line where they are floating point numbers
    so in that case we use some precision to check if the ans is correct or not
    """
    eps = 1e-9
    while abs(r - l) > eps:
        mid = l + (r - l) / 2
        if check(mid):
            r = mid
        else:
            l = mid
    return l


"""
Ternary Search

Time Complexity: O(log n)

What: Searching for an answer maybe a fixed value or some peak value over a unimodal range
Identification: Given any unimodal function f(x) in range [l, r]
unimodal: strictly increasing/decreasing first and then strictly decreasing/increasing e.g. sin(x) in [0, pi]

Idea: To split the search space in 3 halves s.t. 1/3rd of the search space is always excluded

mid1 = l + (r - l) // 3
mid2 = r - (r-l) // 3
"""


def ternary_search(f, l, r):
    eps = 1e-9
    while r - l > eps:
        m1 = l + (r - l) // 3  # here also check whether we need int or float
        m2 = r - (r - l) // 3

        if f(m1) < f(m2):
            l = m1
        else:
            r = m2
    # return what's needed - maybe an index or maybe the value of the function at the given index
    return l
