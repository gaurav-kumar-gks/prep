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
mid = (l + r) / 2 -> this may overflow in c / cpp etc. 
mid = (low + high) // 2 -> this is lower mid
mid = (low + high + 1) // 2 -> this is upper mid
mid = (low + high) >> 2

note: choice of mid & shrinking logic has to work together 
s.t. the search space is reduced by half and 1 element is always excluded
also, always check the case of <=3 elements
i.e.
a b c
a b
a
with the while loop condition and the shrinking logic

also always check
if using lower mid -> then what if l == mid
if using higher mid -> then what if r == mid
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

"""
Lower bound: The lower bound algorithm finds the first or the smallest index in a sorted array where the value at that index is greater than or equal to a given key i.e. x
"""
def find_first_pos(nums, t):
    """
    it returns the first position of the target element in non decreasing array
    
    e.g. 
    nums = [1, 2, 2, 3, 4, 5] and t = 2
    then it should return 1
    if t = 6 then it should return -1
    """
    if not nums:
        return -1
    l, r = 0, len(nums) - 1
    while l < r:
        m = (l + r) // 2
        if nums[m] >= t:
            r = m
        else:
            l = m + 1
    return r if nums[r] == t else -1
    

"""
Upper bound: The upper bound algorithm finds the last or the largest index in a sorted array where the value at that index is less than or equal to a given key i.e. x
"""
def find_last_pos(nums, t):
    """
    it returns the last position of the target element in non decreasing array
    
    e.g. 
    nums = [1, 2, 2, 3, 4, 5] and t = 2
    then it should return 2
    if t = 6 then it should return -1
    """
    if not nums:
        return -1
    l, r = 0, len(nums) - 1
    while l < r:
        m = (l + r + 1) // 2
        if nums[m] <= t:
            l = m
        else:
            r = m - 1
    return l if nums[l] == t else -1

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

