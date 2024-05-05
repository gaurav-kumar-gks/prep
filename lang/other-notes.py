"""
About built in types - 
"""

"""
Comparisons

    1. Objects of different types, except different numeric types, never compare equal
    2. Non-identical instances of a class normally compare as non-equal unless the class defines the __eq__() method
    3. Instances of a class cannot be ordered with respect to other instances of the same class, or other types of object, 
    unless the class defines enough of the methods __lt__(), __le__(), __gt__(), and __ge__() (in general, __lt__() and __eq__() are sufficient, if you want the conventional meanings of the comparison operators)
    
"""

"""
Numeric types

    1. // -> result always rounded towards -inf
    2. floor(a/b) == a//b 
    3. ceil(a/b) == -(-a//b)
    4. % -> result has the same sign as the divisor (unlike C)
    5. ** -> has right associativity (unlike C)
    6. int() && float() -> truncates towards 0
    7. round() -> rounds to nearest even value
    8. x >> y -> x // 2**y
    9. x << y -> x * 2**y
    

"""

"""
Boolean types

and, or and != should be preferred over &, | and ^.
"""

"""
Sequence types

    1. in / not in can be used to check if an item is present in a sequence
    2. in can also be used to check if something is a substring of a string e.g. 'ell' in 'Hello'
    3. s*n == n*s equivalent to adding s to itself n times
        note: [x] * n -> creates a list of n references to x
        e.g. lists = [[]] * 3 # [[], [], []] -> creates a list of 3 references to the same list
            lists[0].append(3)
            list # [[3], [3], [3]]
    4. Concatenating immutable sequences always results in a new object. 
        This means that building up a sequence by repeated concatenation will have a quadratic runtime cost in the total sequence length.
        e.g s = ''; for i in range(1000): s += 'a' -> O(n^2)
            alternatively, use ''.join(list) -> O(n)
    5. s.copy() === s[:] -> shallow copy 
        for deepcopy use copy.deepcopy(s)
    6. s.extend(t) === s[len(s):] = t === s += t
    7. s.remove(x) -> removes the first item from the list whose value is x. It is an error if there is no such item.
    8. s.pop(i) -> removes the item at index i and returns it. If no index is specified, a.pop() removes and returns the last item in the list.
    9. s.index(x[, i[, j]]) -> returns the index of the first occurrence of x in s (at or after index i and before index j).
    10. l.sort() -> in place stable sort
    11. tuple can be created as follows: t = 12345, 54321, 'hello!' - so careful of the trailing comma 
    12. str() === type.(object).__str__(object) 
        -> returns a string containing a nicely printable representation of an object
        -> if __str__() is not defined, then __repr__() is used
        For strings, this returns the string itself. 
        The difference with repr(object) is that str(object) does not always attempt to return a string that is acceptable to eval(); its goal is to return a printable string.
"""

"""
Mapping types

Mapping object maps hashable values to arbitrary objects.
Values that are not hashable, that is, values containing lists, dictionaries or other mutable types (that are compared by value rather than by object identity) may not be used as keys.
"""

# For time complexities https://ics.uci.edu/~pattis/ICS-33/lectures/complexitypython.txt