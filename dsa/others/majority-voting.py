"""
BOYER-MOORE MAJORITY VOTING ALGORITHM


1. PROBLEM STATEMENT:
   - Given an array of n elements, find the majority element
   - Majority element appears more than n/2 times
   - If no majority element exists, return -1 or None

2. ALGORITHM INTUITION:
   - If a majority element exists, it will "survive" the voting process
   - We maintain a candidate and a count
   - For each element: if it matches candidate, increment count; else decrement
   - If count reaches 0, update candidate to current element
"""

def boyer_moore_majority_vote(arr):
    if not arr:
        return None
    
    candidate = arr[0]
    count = 1
    
    for i in range(1, len(arr)):
        if arr[i] == candidate:
            count += 1
        else:
            count -= 1
            if count == 0:
                candidate, count = arr[i], 1
    
    return candidate if sum(1 for x in arr if x == candidate) > len(arr) // 2 else None





def boyer_moore_majority_vote_generalized(arr, k):
    """
    Generalized Boyer-Moore for finding elements that appear more than n/k times.
    This is a more complex version that can find up to k-1 elements that appear more than n/k times.
    Time Complexity: O(n * k)
    Space Complexity: O(k)
    """
    if not arr or k <= 1:
        return []
    
    candidates = {}
    for element in arr:
        if element in candidates:
            candidates[element] += 1
        elif len(candidates) < k - 1:
            candidates[element] = 1
        else:
            candidates = {key: count - 1 for key, count in candidates.items() if count > 1}
    
    return [candidate for candidate in candidates if sum(1 for x in arr if x == candidate) > len(arr) // k]

