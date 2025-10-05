"""
BIT MANIPULATIONS
"""

# all bit manipulation tricks from https://leetcode.com/discuss/study-guide/4282051/All-Types-of-Patterns-for-Bits-Manipulations-and-How-to-use-it

def print_binary(n):
    """
    it prints the binary representation of the number
    """
    while n:
        print(n & 1, end='')
        n >>= 1
        

"""

Count the number of set bits
n = 10
count = 0
while n:
    count += n & 1
    n >>= 1

-n = ~n + 1

Find the rightmost set bit
n & -n

Remove the rightmost set bit
n & n-1

Check if the number is even or odd
n & 1

Find the ith bit
n & (1 << i)

Set the ith bit
n | (1 << i)

Unset the ith bit
n & ~(1 << i)

Unset upto ith bit
n & ~((1 << i+1) - 1)

Toggle the ith bit
n ^ (1 << i)

Check if a number is power of 2
n and not(n & (n-1))

Check if a number is power of 4
not (n & (n - 1)) and (n % 3 == 1)

Check if a number is power of 3
n and (pow(3, 19) % n == 0) # 3^19 is the maximum power of 3 that can be stored in 32-bit integer

Multiply a number by 2
n << 1

Divide a number by 2
n >> 1

Convert a lowercase letter to uppercase
c & '_'
c & (1 << 5)

Convert an uppercase letter to lowercase
c | ' '
c | (1 << 5)

swap two numbers
a ^= b
b ^= a
a ^= b

"""

"""
xor tricks

1. x ^ 0 = x
2. x ^ x = 0
3. x ^ y = y ^ x
4. xor of n numbers = xor of xor of n/2 pairs
5. xor of range(0, n) == [n, 1, n+1, 0][n % 4]
6. x ^ ~x = 1

"""
