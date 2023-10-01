"""
Sliding window
"""

"""
The sliding window pattern is method aimed at reducing the use of nested loops in an algorithm. 
It’s a variation of the two pointers pattern, where the pointers can be used to set window bounds.

A window is a sublist formed over a part of an iterable data structure. 
It can be used to slide over the data in chunks corresponding to the window size. 
The sliding window pattern allows us to process the data in segments instead of the entire list. 
The segment or window size can be set according to the problem’s requirements


"""

"""
Does my problem match this pattern?

Yes, if both these conditions are fulfilled: ->

1) The problem requires repeated computations on a contiguous set of data elements (a subarray or a substring), 
such that the window moves across the input array from one end to the other. 
The size of the window may be fixed or variable, depending on the requirements of the problem. 
The repeated computations may be a direct part of the final solution, or they may be intermediate steps building up towards the final solution.

2) The computations performed every time the window moves take O(1) time or are a slow-growing function, such as log of a small variable, say k, where k≪n.


No, if either of these conditions are fulfilled: ->
1) The input data structure does not support random access.
2) You have to process the entire data without segmentation.
"""

"""
Idea: 
We have a "window" of 2 pointers, left and right, and we keep increasing the right pointer.

If the element at the right pointer makes the window not valid, we keep moving the left pointer to shrink the window until it becomes valid again.
Then, we update the global min/max with the result from the valid window.
To check if it is valid, we need to store the "state" of the window (ex. frequency of letters, number of distinct integers).


for(right = 0; right < n; right++):
    update window with element at right pointer
    while (condition not valid):
        remove element at left pointer from window, move left pointer to the right
    update global max

"""

"""
exactly K type questions:

can be solved in 2 ways using sliding window
using => exact(k) = atmost(k) - atmast(k-1)
or using prefix sliding window
"""


"""
Concept of adding length of subarray =>

In at most K type of questions
If [i … j] is a valid subarray, ex. [1,2,3,4]

There are these many subarrays - 
4 with length 1, [1], [2], [3], [4]
3 with length 2: [1,2], [2,3], [3,4]
2 with length 3
1 with length 4
So if at every step we added the length:

[1] -> add 1
[1,2] -> add 2
[1,2,3] -> add 3
[1,2,3,4] -> add 4
the sum 1+2+3+4 would be the same as the number of subarrays. So we can add the lengths of the valid subarrays.
"""


"""
Prefixed sliding window

If the subarray [i,j] contains K 1s, and the first p numbers of the subarray are 0, 
then that means all the subarrays from [i,j] to [i+p, j] inclusive are valid, which is p+1 subarrays.

If the subarray [i,j] contains X distinct numbers, and the first p numbers are duplicates (are also in [i + p, j], 
that means all the subarrays from [i,j], [i+1, j], … [i+p, j] have X distinct numbers as well. 
Then if [i,j] is valid, there will be p+1 valid subarrays

"""


"""
When will sliding window not work?

1: Sliding window doesn't work if knowing one element at the edges of the window, does not tell you how to update the state of the window, or whether it becomes valid.
2: Another type of question that doesn't work, is if adding one new element could either increase or decrease the window's state.
3: If given an invalid subarray it is hard to check whether adding or removing from only one end at a time would ever make it valid.

"""
