"""
Sieve of Eratosthenes

Finding all primes in [1:n] in O(nloglogn)

- Number is prime if none of the smaller primes divides it
- Number of primes <= n is approx. n / ln(n)
- Kth prime number is approx. k*ln(k)

Algo:

1. Create a list of all numbers from 2 to n
    Optimizations:
    - Only odd numbers
    - Use list comprehension for marking
    - Start marking from i*i

2. For each number, if it is prime, mark all its multiples as not prime
    Optimizations:
    - Only mark multiples of prime numbers
    - Start marking from i*i

3. The numbers that are not marked as not prime are the primes
"""

import random
import math

def sieve_of_eratosthenes(n):
    if n < 2: return []
    primes = [2]
    sieve = [True] * ((n+1) // 2)
    # Can use bitwise as well here
    # sieve = bytearray([1]) * ( (n+1) // 2) // 8 + 1)
    # getbit = lambda arr, pos: (arr[pos // 8] >> (pos % 8)) & 1
    # clearbit = lambda arr, pos: arr[pos // 8] &= ~(1 << (pos % 8))
    for i in range(3, int(n**0.5) + 1, 2):
        if sieve[i // 2]:
            for j in range(i * i, n + 1, 2*i):
                sieve[j // 2] = False
    return [2] + [2*i+1 for i in range(1, (n+1)//2) if sieve[i]]

def sieve_linear(n):
    """
    Linear Sieve - O(n) time complexity
    
    Time Complexity: O(n)
    Space Complexity: O(n)
    
    Uses linear sieve algorithm with smallest prime factor
    Useful when we need to calculate prime as well as smallest / least prime factor
    In practise doesn't offer much performance over optimized sieve
    """
    if n < 2:
        return []
    spf = [0] * (n + 1)
    primes = []
    for i in range(2, n + 1):
        if spf[i] == 0:
            spf[i] = i
            primes.append(i)
        
        for p in primes:
            if p > spf[i] or i * p > n:
                break
            spf[i * p] = p
    return primes


def is_prime_trial_division(n: int) -> bool:
    """
    Trial Division - Simplest primality test
    Time Complexity: O(√n)
    Space Complexity: O(1)
    - Small numbers (n < 10^6)
    """
    if n == 2:
        return True
    if n < 2 or n % 2 == 0:
        return False
    # Check odd divisors up to √n
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        if n % i == 0:
            return False
    return True


def is_prime_trial_division_optimized(n: int) -> bool:
    """
    Optimized Trial Division
    Time Complexity: O(√n)
    Space Complexity: O(1)
    """
    if n == 2:
        return True
    if n < 2 or n % 2 == 0:
        return False
    # Check small primes first
    small_primes = [3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37] # next prime 41
    for p in small_primes:
        if n == p: return True
        if n % p == 0: return False
    
    # wheel of 30 pattern for larger numbers
    wheel = [1, 7, 11, 13, 17, 19, 23, 29]
    sqrt_n = int(math.sqrt(n))
    
    for base in range(41, sqrt_n + 1, 30): 
        for offset in wheel:
            i = base + offset
            if i > sqrt_n:
                break
            if n % i == 0:
                return False
    return True


def is_prime_fermat(n: int, k: int = 5) -> bool:
    """
    Fermat's Little Theorem Test (Probabilistic)
    Fermat's Little Theorem: 
        If p is prime, then a^(p-1) with mod p is 1
    
    Time Complexity: O(k log²n log log n) for k iterations
    Space Complexity: O(1)
    
    Algorithm:
    1. Choose random base a where 1 < a < n
    2. Check if a^(n-1) ≡ 1 (mod n)
    3. Repeat k times
    
    This test has false positives (Carmichael numbers)
    """
    if n == 2:
        return True
    if n < 2 or n % 2 == 0:
        return False
    
    def binexpo(base: int, exp: int, mod: int) -> int:
        result = 1
        base %= mod
        while exp > 0:
            if exp % 2 == 1: result = (result * base) % mod
            base = (base << 1) % mod
            exp >>= 1
        return result
    for _ in range(k):
        a = random.randint(2, n - 1)
        if binexpo(a, n - 1, n) != 1:
            return False
    return True
