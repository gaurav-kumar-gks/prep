"""
Longest prefix which is also a suffix
- online algo - can process characters as they arrive
"""


def longest_prefix_suffix_naive(s):
    n = len(s)
    lps = [0] * n
    end = 2
    while end <= n:
        a = 0
        for k in range(1, end):
            if s[0:k] == s[end - k : end]:
                a = max(a, k)
        lps[end - 1] = a
        end += 1
    print(lps)
    return lps


"""
E.g. -
aaa
012 <- lps

key idea: prefixes are nested 
& hence we don't need to compare the entire prefix

prev_lps = lps[prev_lps - 1] optimization

       $#$#        $#$#i
ind -> 01234       ....i
arr -> asasd.......asasa
lps ->    2x          4?
"""


def longest_prefix_suffix(s):
    n = len(s)
    lps = [0] * n
    for i in range(1, n):
        prev_lps = lps[i - 1]
        while prev_lps > 0 and s[prev_lps] != s[i]:
            prev_lps = lps[prev_lps - 1]
        if s[prev_lps] == s[i]:
            prev_lps += 1
        lps[i] = prev_lps
    print(lps)
    return lps
