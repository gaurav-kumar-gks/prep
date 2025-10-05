"""
1. PRIME FACTORIZATION:
   - Every integer n > 1 can be uniquely written as:
     n = p1^a1 * p2^a2 * ... * pk^ak
   - where p1, p2, ..., pk are distinct primes
   - and a1, a2, ..., ak are positive integers

2. NUMBER OF DIVISORS:
   - If n = p1^a1 * p2^a2 * ... * pk^ak
   - number of divisors = (a1 + 1) * (a2 + 1) * ... * (ak + 1)

3. SUM OF DIVISORS:
   - If n = p1^a1 * p2^a2 * ... * pk^ak
   - sum_of_divisors = (p1^(a1+1) - 1)/(p1 - 1) * (p2^(a2+1) - 1)/(p2 - 1) * ... * (pk^(ak+1) - 1)/(pk - 1)
   - sum_of_proper_divisors = sum_of_divisors - n

4. EULER'S TOTIENT FUNCTION:
   - φ(n) = number of integers in [1, n] coprime to n
   - φ(n) = n * (1 - 1/p1) * (1 - 1/p2) * ... * (1 - 1/pk)
   - φ(n) is multiplicative: φ(ab) = φ(a) * φ(b) if gcd(a, b) = 1

5. SPECIAL NUMBERS:
   - Perfect: sum of proper divisors = number (6, 28, 496, ...)
   - Abundant: sum of proper divisors > number (12, 18, 20, ...)
   - Deficient: sum of proper divisors < number (most numbers)
   - Amicable: pair (a, b) where σ(a) = b and σ(b) = a
"""

import math
from typing import List, Tuple, Dict
from collections import defaultdict


def get_prime_factors(n):
    """
    Get prime factorization of n
    Returns: Dictionary {prime: exponent}
    Example: 12 = 2^2 * 3^1 -> {2: 2, 3: 1}
    Time Complexity: O(√n)
    Space Complexity: O(log n)
    """
    factors = defaultdict(int)
    while n % 2 == 0:
        factors[2] += 1
        n //= 2
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        while n % i == 0:
            factors[i] += 1
            n //= i
    if n > 1:
        factors[n] += 1
    return factors

def get_all_divisors(n):
    """
    Get all divisors of n (including 1 and n)
    1. Find prime factorization
    2. Generate all combinations of prime factors
    Time Complexity: O(√n + d log d) where d = number of divisors
    Space Complexity: O(d)
    """
    if n <= 1:
        return [n] * n
    
    factors = get_prime_factors(abs(n))
    divisors = [1]
    
    def generate_divisors(prime_factors, current_divisor=1, index=0):
        if index == len(prime_factors):
            if current_divisor != 1:
                divisors.append(current_divisor)
            return
        prime, max_exp = prime_factors[index]
        for exp in range(max_exp + 1):
            generate_divisors(prime_factors, current_divisor * (prime ** exp), index + 1)
    
    prime_list = list(factors.items())
    generate_divisors(prime_list)
    return divisors

def count_divisors_sieve(n: int) -> List[int]:
    """
    Sieve to count divisors for all numbers from 1 to n
    
    Time Complexity: O(n log n)
    Space Complexity: O(n)
    """
    count_divisors = [1] * (n + 1)
    count_divisors[0] = 0
    for i in range(2, n + 1):
        for j in range(i, n + 1, i):
            count_divisors[j] += 1
    return count_divisors

def count_factors_of_prime(n, p):
    """
    Count how many times prime p appears in n!
    Legendre's formula: n/p + n/p**2 + n/p**3 + ...
    """
    count = 0
    power = p
    while power <= n:
        count += n // power
        power *= p
    return count

def euler_totient_sieve(n: int) -> List[int]:
    """
    Sieve to calculate Euler's totient for all numbers from 1 to n
    
    Time Complexity: O(n log log n)
    Space Complexity: O(n)
    
    Returns: List where phi[i] = φ(i)
    """
    phi = list(range(n + 1))
    
    for i in range(2, n + 1):
        if phi[i] == i:  # i is prime
            for j in range(i, n + 1, i):
                phi[j] = phi[j] // i * (i - 1)
    
    return phi
