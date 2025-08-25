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
- Associative: gcd(a, b, c) = gcd(a, gcd(b, c))
"""

"""
LCM: Least common multiple

lcm(a,b) = a*b / gcd(a,b)
"""

"""
Euclidean Algo

gcd(a,b)
= a if b == 0
= gcd(b, a%b) otherwise

Time complexity: O(log(min(a,b)))
"""


def gcd_recursive(a, b):
    if b == 0:
        return a
    return gcd_recursive(b, a % b)


def gcd(a, b):
    while b:
        a, b = b, a % b
    return a

def lcm(a, b):
    return (a * b) // gcd(a, b)

"""
Extended euclidean algo

ax + by = gcd(a, b)
Let's say we find gcd(a,b) = gcd(b, a%b) = gcd
ax + by = gcd(a, b)
bx1 + (a%b)y1 = gcd(b, a%b) = gcd(a, b) = ax + by
bx1 + (a - (a//b * b))y1 = ax + by
ay1 + b(x1 - (a//b * y1)) = ax + by
x = y1
y = x1 - (a//b * y1)


Time complexity: O(log(min(a,b))) 
"""

def extended_gcd_recursive(a, b):
    if b == 0:
        return a, 1, 0
    gcd, x1, y1 = extended_gcd_recursive(b, a % b)
    x = y1
    y = x1 - (a // b) * y1
    return gcd, x, y

def extended_gcd(a, b):
    x, y = 1, 0
    while b > 0:
        q = a // b
        a, b = b, a % b
        x, y = y, x - q * y
    return a, x, y


"""
Modular multiplicative Inverse

- a number x such that (a * x) % m = 1
- Modular inverse exists only if gcd(a, m) = 1
- Used in cryptography, RSA algorithm, solving modular equations

Time Complexity: O(log(min(a,m)))
"""

def mod_inverse(a, m):
    gcd, x, y = extended_gcd(a, m)
    if gcd != 1:
        return None
    return (x % m + m) % m  # Ensure positive result


"""
Linear Diophantine Equation

ax + by = c

- Solve for x, y
- Solution only exists if c is divisible by gcd(a, b)
- Used in number theory problems

Time complexity: O(log(min(a,b)))
"""

def solve_linear_diophantine(a, b, c):
    gcd, x, y = extended_gcd(abs(a), abs(b))
    if c % gcd != 0:
        return None  # No solution exists
    m = c // gcd
    x = x * m * (1 if a > 0 else -1)
    y = y * m * (1 if b > 0 else -1)
    
    solutions = []
    for t in range(-10, 10):
        solutions.append((x + (b // gcd) * t, y - (a // gcd) * t))
    
    return x, y
