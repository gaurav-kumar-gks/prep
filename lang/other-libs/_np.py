import numpy as np
import operator
from itertools import accumulate

# Array Creation
arr = np.array([1, 2, 3, 4]) # [1, 2, 3, 4]
zeros = np.zeros((2, 3)) # [[0., 0., 0.], [0., 0., 0.]]
ones = np.ones((2, 3)) # [[1., 1., 1.], [1., 1., 1.]]
arange = np.arange(0, 10, 2) # [0, 2, 4, 6, 8]
linspace = np.linspace(0, 1, 5) # [0.  , 0.25, 0.5 , 0.75, 1.  ]
identity = np.eye(3) # [[1., 0., 0.], [0., 1., 0.], [0., 0., 1.]]
full = np.full((2, 2), 7) # [[7, 7], [7, 7]]
random_array = np.random.random((2, 2)) # [[0.417022  , 0.72032449], [0.00011437, 0.30233257]]

# Array Manipulation
reshaped = arr.reshape((2, 2)) # [[1, 2], [3, 4]]
flattened = arr.flatten() # [1, 2, 3, 4]
transposed = np.array([[1, 2, 3], [4, 5, 6]]).transpose() # [[1, 4], [2, 5], [3, 6]]
concatenated = np.concatenate((arr, arr)) # [1, 2, 3, 4, 1, 2, 3, 4]
stacked = np.stack((arr, arr)) # [[1, 2, 3, 4], [1, 2, 3, 4]]

# Array Operations
a = np.array([1, 2, 3]) # [1, 2, 3]
b = np.array([4, 5, 6]) # [4, 5, 6]
add_result = np.add(a, b) # [5, 7, 9]
subtract_result = np.subtract(a, b) # [-3, -3, -3]
multiply_result = np.multiply(a, b) # [4, 10, 18]
divide_result = np.divide(a, b) # [0.25, 0.4, 0.5]
dot_result = np.dot(a, b) # 32
power_result = np.power(a, 2) # [1, 4, 9]

# Statistical Operations
arr = np.array([1, 2, 3, 4, 5]) # [1, 2, 3, 4, 5]
mean = np.mean(arr) # 3.0
median = np.median(arr) # 3.0
std_dev = np.std(arr) # 1.4142135623730951
variance = np.var(arr) # 2.0
total = np.sum(arr) # 15
min_value = np.min(arr) # 1
max_value = np.max(arr) # 5

# Logical Operations
arr_bool = np.array([True, True, False])
all_result = np.all(arr_bool) # False
any_result = np.any(arr_bool) # True
logical_and = np.logical_and(arr_bool, [True, False, True]) # [True, False, False]
logical_or = np.logical_or(arr_bool, [True, False, True]) # [True, True, True]

# Indexing and Slicing
element = arr[2] # 3
slice_result = arr[1:4] # [2, 3, 4]
boolean_indexing = arr[arr > 3] # [4, 5]
where_result = np.where(arr > 3) # (array([3, 4]),)

# Linear Algebra
matrix = np.array([[1, 2], [3, 4]]) # [[1, 2], [3, 4]]
determinant = np.linalg.det(matrix) # -2.0000000000000004
inverse = np.linalg.inv(matrix) # [[-2. ,  1. ], [ 1.5, -0.5]]
eigenvalues, eigenvectors = np.linalg.eig(matrix)

# Broadcasting
broadcasted_add = arr + 10 # [11, 12, 13, 14, 15]

# itertools.accumulate for prefix sum
prefix_sum = list(accumulate(arr)) # [1, 3, 6, 10, 15]
cumulative_product = list(accumulate(arr, operator.mul)) # [1, 2, 6, 24, 120]

