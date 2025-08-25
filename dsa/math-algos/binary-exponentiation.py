"""
Binary Exponentiation
"""

"""
Binary exponentiation
Time complexity: O(log(N))

e.g. calculating a^N in O(log(N)) instead of O(n)

can be used with any operation that has power of associativity
i.e. a*(b*c) = (a*b)*c


a^n
= 1 if n==0
= a * (a^(n-1/2))^2 if n is odd
= (a^(n/2))^2 if n is even

Note: if mod is prime
then (a^n) == (a^(n%(m-1))
"""

def binary_exponentiation_recursive(a, n):
    if n == 0: 
        return 1
    if a & 1:
        return a * binary_exponentiation_recursive(a*a, n//2)
    return binary_exponentiation_recursive(a*a, n//2)

def binary_exponentiation(a, n, mod):
    a %= mod
    res = 1
    while n > 0:
        if n & 1:
            res = (res * a) % mod
        a = (a * a) % mod
        n >>= 1
    return res


"""
Multiplying two numbers

a*b = a + a + ...... b times..

a*b
= 0 if b == 0
= a + 2*a*((b-1)/2) if b is odd
= 2*a*(b/2) if b is even
""" 

"""
Applying permutations k times

Sequence of length n, apply permutation k times

- Consider applying permutation as commutative operation
- We can use binary exponentiation to apply permutation k times
- if k is odd, apply permutation to sequence
- if k is even, apply permutation to permutation
- repeat this process until k becomes 0

- Time complexity: O(nlog(k))
"""

"""
Fast application of set of geometric operations to a set of points

Consider coordinate: [x, y, z]
Each transformation on coordinate is a linear operation on coordinate
Matrix Multiplication of nCoordinates+1 x nCoordinates+1 matrix

Shift: add a constant to each coordinate
e.g. shift x by 5, y by 7, z by 9
[x, y, z, 1][[1, 0, 0, 0],[0, 1, 0, 0],[0, 0, 1, 0],[5, 7, 9, 1]]

Scale: multiply a constant to each coordinate
e.g. multiply each x, y, z by 2
[x, y, z, 1][[2, 0, 0, 0],[0, 2, 0, 0],[0, 0, 2, 0],[0, 0, 0, 1]]

Rotate: rotate by a constant angle
e.g. rotate x, y, z by 90 degrees around z axis

For (x, y)
Counterclockwise by theta:
[x, y, 1][[cos(theta), sin(theta), 0], [-sin(theta), cos(theta), 0], [0, 0, 1]]
For (x, y, z)
Roll: Rotate around x axis
Pitch: Rotate around y axis
Yaw: Rotate around z axis

once every transformation is described as a matrix, 
the sequence of transformations can be described as a product of these matrices, 
and a "loop" of k repetitions can be described as the matrix raised to the power of k 
(which can be calculated using binary exponentiation in O(log(k)) like permutation problem). 
This way, the matrix which represents all transformations can be calculated first in O(m*log(k)), 
and then it can be applied to each of the n points in O(n) for a total complexity of O(n + m*log(k)).
"""
