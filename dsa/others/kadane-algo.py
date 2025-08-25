"""
KADANE'S ALGORITHM
TC: O(n)

1. We want to track max sum ending at index i
2. Keep track of current sum and maximum sum
3. If current sum becomes negative, reset to 0
4. Update maximum sum whenever current sum exceeds it
"""

def kadane_algorithm(arr):
    n = len(arr)
    max_sum = float('-inf')
    current_sum = 0
    for i in range(n):
        current_sum += arr[i]
        max_sum = max(current_sum, max_sum)
        current_sum = max(current_sum, 0)
        # to handle all negative numbers
        # if current_sum + arr[i] > arr[i]:
        #     current_sum = current_sum + arr[i]
        # else:
        #     current_sum = arr[i]
    return max_sum

"""
Examples:

1. Maximum subarray sum
    a) Brute Force:
        1) generate all subarrays and then calculate the sum of that subarray
        2) prefix sum may be used for getting subarray sum in O(1)
            Prefix sum: psum[i] = psum[i-1] + arr[i]
    b) Kadane's algorithm
    c) Segment Tree with range query
        each node will have (sum, max_presum, max_suffsum, max_subarray_sum)

2. Maximum submatrix sum
    a) Brute Force:
        1) generate all submatrices and then calculate the sum of that submatrix
        2) prefix sum may be used for getting submatrix sum in O(1)
            Prefix sum: psum[i][j] = psum[i-1][j] + psum[i][j-1] - psum[i-1][j-1] + arr[i][j]
    b) Kadane's algorithm
        1) We need to convert this 2d problem to 1d so that we can apply kadane
        2) Now, if we fix two columns c1 and c2 and for each row calculate the sum b/w col c1 and c2
            we effectively have array of sums for each row
            we can easily apply kadane's algorithm to find the maximum sum subarray
        3) prefix sum should be used to get the sum of the subarray in O(1)

3. Maximum average subarray
    a) Brute Force
    b) Kadane Algo with Binary search
        We would have some bound for the ans e.g. Average could lie b/w m..n
        Binary search for the average
        For each iteration we need to check if we can get a subarray with average >= mid
        e.g. maybe decrease all elements by x then find if subarray with sum > 0 exists or not
"""



