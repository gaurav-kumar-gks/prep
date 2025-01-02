"""
Patterns of Problems based on KMP

Substring Search:
- Find the first occurrence of a pattern in a string.
- Count occurrences of a pattern in a string.

Preprocessing Problems:
- Reuse precomputed LPS arrays for multiple pattern matching queries.

Other Applications:
- Finding periodicity in a string.
- Implementing a string-based automation.
"""

from dsa.strings.prefixsuffix import longest_prefix_suffix


def kmp(s, t):
    n = len(s)
    m = len(t)
    lps = longest_prefix_suffix(s)
    result = []
    j = 0
    for i in range(m):
        while j > 0 and s[j] != t[i]:
            j = lps[j - 1]
        if s[j] == t[i]:
            j += 1
        if j == n:
            result.append(i - j + 1)
            j = lps[j - 1]
    return result


def find_pattern_occurrences_using_lps(pattern, text):
    concat_string = pattern + "#" + text
    lps = longest_prefix_suffix(concat_string)
    n = len(pattern)
    result = []

    for i in range(n + 1, len(concat_string)):
        if lps[i] == n:
            result.append(i - 2 * n)
    return result


def count_prefix_appearances(s):
    """
    str = ababab
    lps = 001234
    res = 032100
    """
    n = len(s)
    lps = longest_prefix_suffix(s)
    result = [1] * (n + 1)
    for i in range(n):
        if lps[i]:
            result[lps[i]] += 1
    for i in range(n - 1, 0, -1):
        if lps[i - 1]:
            result[lps[i - 1]] += result[i]
    return result[1:]


def count_prefix_appearances_in_t(s, t):
    n = len(s)
    m = len(t)
    concat = s + "#" + t
    lps = longest_prefix_suffix(concat)
    result = [0] * (n + 1)

    for i in range(n + 1, n + 1 + m):
        if 0 < lps[i] <= n:
            result[lps[i]] += 1
    for i in range(n, 0, -1):
        result[lps[i - 1]] += result[i]
    return result[1:]


def no_of_distinct_substrings(s):
    pass


def smallest_period(s):
    lps = longest_prefix_suffix(s)
    n = len(s)
    if n % (n - lps[-1]) == 0:
        return s[: n - lps[-1]]
    return s


def min_chars_to_make_palindrome(s):
    r = s[::-1]  # Reverse of the string
    concat = s + "#" + r  # Concatenated string
    lps = longest_prefix_suffix(concat)
    longest_palindromic_prefix = lps[-1]
    return len(s) - longest_palindromic_prefix  # Minimum chars to add
