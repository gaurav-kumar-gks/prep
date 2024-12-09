"""
Suffix array

Suffix array is a sorted array of all suffixes of a string.
The suffix array is a data structure that is used to store the lexicographically sorted suffixes of a string.
The suffix array is used in various string processing algorithms like pattern searching, string compression, etc.

e.g. -
s = "banana"
suffixes = ["banana", "anana", "nana", "ana", "na", "a"]
suffixes_sorted = ["a", "ana", "anana", "banana", "na", "nana"]
suffix_index = [5, 3, 1, 0, 4, 2]
"""


def get_suffix_array_naive(s):
    # Time complexity: O(n^2logn)
    n = len(s)
    suffix_array = []
    for i in range(n):
        suffix_array.append((s[i:], i))
    sorted_suffix_array = sorted(suffix_array)
    sorted_suffix_index = [x[1] for x in sorted_suffix_array]
    return sorted_suffix_index


def get_suffix_array(s):
    """Builds the suffix array using an optimized approach (O(n log n))."""
    n = len(s)
    suffix_array = list(range(n))
    rank = list(map(ord, s))
    new_ranks = [0] * n
    k = 1
    while k < n:
        # compare using key: (rank[i], rank[i + k]) because we are sorting index of suffixes with length 2^k
        # at this point rank[i] is the rank of the suffix with length 2^(k-1) starting at i
        suffix_array.sort(key=lambda i: (rank[i], rank[i + k] if i + k < n else -1))
        r = 0
        new_ranks[suffix_array[0]] = 0
        for i in range(1, n):
            prev = suffix_array[i - 1]
            curr = suffix_array[i]
            if (rank[prev], rank[prev + k] if prev + k < n else -1) != (
                rank[curr],
                rank[curr + k] if curr + k < n else -1,
            ):
                r += 1
            new_ranks[curr] = r
        rank = new_ranks[:]
        k *= 2
    return suffix_array
