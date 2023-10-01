"""
Binary Exponentiation
"""

"""
Binary exponentiation
Time complexity: O(log(N))

e.g. calculating a^N in O(log(N)) instead of O(n)

can be used with any operation that has power of associativity
i.e. a*(b*c) = (a*b)*c

core idea:


a^n
= 1 if n==0
= a * (a^(n-1/2))^2 if n is odd
= (a^(n/2))^2 if n is even
"""


# recursive implementation of pow(a, n)
def binary_exponentiation_recursive(a, n):
    if n == 0:
        return 1
    res = binary_exponentiation_recursive(a, (n - 1) // 2)
    if n & 1:
        return a * res * res
    return res * res


def binary_exponentiation_iterative(a, n, mod):
    a %= mod
    res = 1
    while n > 0:
        if n & 1:
            res = (res * a) % mod
        a = (a * a) % mod
        n >>= 1
    return res


"""
Note: if mod is prime

then (a^n) == (a^(n%(m-1))
"""

"""
Multiplying two numbers


a*b = a + a + ...... b times..


a*b
= 0 if b == 0
= a + 2*a*((b-1)/2) if b is odd
= 2*a*(b/2) if b is even
"""
