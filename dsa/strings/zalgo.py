def get_z_array_naive(s):
    n = len(s)
    z = [0] * n
    for i in range(1, n):
        l, r = 0, i
        while r < n and s[l] == s[r]:
            l += 1
            r += 1
            z[i] += 1
    print(z)
    return z


def get_z_array(s):
    n = len(s)
    z = [0] * n
    l, r = 0, 0

    for i in range(1, n):
        if i > r:
            # Case 1: i is outside the current window
            l, r = i, i
            while r < n and s[r] == s[r - l]:
                r += 1
            z[i] = r - l
            r -= 1
        else:
            # Case 2: i is inside the current window
            k = i - l  # Corresponding index in the prefix
            if z[k] < r - i + 1:
                z[i] = z[k]  # Use the precomputed Z value
            else:
                # Expand the window beyond R
                l = i
                while r < n and s[r] == s[r - l]:
                    r += 1
                z[i] = r - l
                r -= 1
    return z


def z_algorithm_search(text, pattern):
    concat = pattern + "#" + text
    z = get_z_array(concat)
    m = len(pattern)
    occurrences = []
    for i in range(m + 1, len(z)):
        if z[i] == m:
            occurrences.append(i - m - 1)
    return occurrences


def longest_palindromic_prefix(s):
    z = get_z_array(s + "#" + s[::-1])
    return z[len(s) + 1]


def find_period(s):
    n = len(s)
    z = get_z_array(s)
    for p in range(1, n + 1):
        if n % p == 0 and z[n - p] == p:
            return p
    return n  # If no smaller period exists, the string's period is itself


def count_distinct_substrings(s):
    n = len(s)
    total_distinct = 0
    for i in range(n):
        suffix = s[i:]
        z = get_z_array(suffix)
        total_distinct += len(suffix) - max(z)
    return total_distinct
