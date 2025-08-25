"""
Fibonacci Numbers

Fn = Fn-1 + Fn-2

Properties:
- Cassini Identity: 
    Fn-1 * Fn+1 - Fn^2 = (-1)^n
- Addition:
    Fn+k = Fk * Fn+1 + Fk-1 * Fn
- Gcd:
    gcd(Fm, Fn) = Fgcd(m, n)
- Fibonacci numbers are worst input for euclidean algo
- Fn = (phi^n - psi^n) / sqrt(5)
    phi = (1 + sqrt(5)) / 2
    psi = (1 - sqrt(5)) / 2
"""

def fib(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    return fib(n-1) + fib(n-2)

def fibonacci(n):
    a, b = 0, 1
    for _ in range(n):
        a, b = b, a + b
    return a
