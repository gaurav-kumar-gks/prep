"""
GCD related algo

"""

"""
GCD: Greatest Common Divisor
also known as Highest Common Factor

gcd(a,b) = max{k>0: k | a and k | b}
the largest number which is a divisor of both a and b

- if a or b == 0 then gcd is the second no.
- if both of them are prime then gcd will always be 1
"""
"""
LCM: Least common multiple

lcm(a,b) = a*b / gcd(a,b)

"""

"""
Euclidean Algo for calculating gcd

gcd(a,b)
= a if b == 0
= gcd(b, a%b) otherwise

Time complexity: O(log(min(a,b)))
"""


def gcd_recursive(x, y):
    if y == 0:
        return x
    return gcd_recursive(y, x % y)


def gcd_iterative(x, y):
    while y:
        x, y = y, x % y
    return x


def lcm(x, y):
    return x / gcd_iterative(x, y) * y
