"""
Various sorting algorithms
"""


"""
Bubble sort

Time Complexity: O(n^2)
Space Complexity: O(1)
Inplace: Yes
Stable: Yes

What: Sorting an array by repeatedly swapping adjacent elements if they are in wrong order
Identification: If the array is almost sorted
Idea: Bubble up the largest element to the end of the array
Observation: After each pass, the largest element is bubbled up to the end of the array
"""


def bubble_sort(arr):
    """Bubble sort"""
    n = len(arr)
    for i in range(n):
        is_swapped = False
        for j in range(n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                is_swapped = True
        if not is_swapped:
            break
    return arr


"""
Selection sort

Time Complexity: O(n^2)
Inplace: Yes
Stable: No

How: Find the min element and swap it with the first element
Idea: Find the min element and swap it with the first element
Observation: After the first pass, the first element is the smallest element in the array
"""


def selection_sort(arr):
    """Selection sort"""
    n = len(arr)
    for i in range(n):
        min_idx = i
        for j in range(i + 1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]
    return arr


"""
Insertion Sort

Time Complexity: O(n^2)
Inplace: Yes
Stable: Yes

What: Inserting an element in a sorted array
How: Insert the element at the correct position in the sorted array
Idea: To maintain a sorted array and insert the next element at the correct position
Identification: When the array is almost sorted
Unique observation: The array is almost sorted, so the element to be inserted is almost at the correct position
"""


def insertion_sort(arr):
    """Insertion sort"""
    n = len(arr)
    for i in range(1, n):
        key = arr[i]
        j = i - 1
        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key
    return arr


"""
QuickSort

Time Complexity: O(nlogn) in average case, O(n^2) in worst case
Space Complexity: O(logn) in average case, O(n) in worst case
Inplace: Yes
Stable: No

What: Sorting an array
How: Divide and Conquer
Idea: Pick an element as pivot and partition the array around the picked pivot
Partition: Put pivot at its correct position in sorted array and put all smaller elements before pivot and all greater elements after pivot
"""


def quick_sort_lomuto(arr):
    """Quick sort using Lomuto partition scheme"""

    def lomuto_partitioning(arr, l, r):
        pivot = arr[r]
        i = l
        for j in range(l, r):
            if arr[j] < pivot:
                arr[i], arr[j] = arr[j], arr[i]
                i += 1
        arr[i], arr[r] = arr[r], arr[i]
        return i

    def quick_sort_helper(arr, l, r):
        if l < r:
            p = lomuto_partitioning(arr, l, r)
            quick_sort_helper(arr, l, p - 1)
            quick_sort_helper(arr, p + 1, r)

    quick_sort_helper(arr, 0, len(arr) - 1)
    return arr


# quick sort using hoare's partition scheme
def quick_sort_hoare(arr):
    """Quick sort using Hoare's partition scheme"""

    def hoare_partitioning(arr, l, r):
        pivot = arr[l]
        i = l - 1
        j = r + 1
        while True:
            i += 1
            while arr[i] < pivot:
                i += 1
            j -= 1
            while arr[j] > pivot:
                j -= 1
            if i >= j:
                return j
            arr[i], arr[j] = arr[j], arr[i]

    def quick_sort_helper(arr, l, r):
        if l < r:
            p = hoare_partitioning(arr, l, r)
            quick_sort_helper(arr, l, p)
            quick_sort_helper(arr, p + 1, r)

    quick_sort_helper(arr, 0, len(arr) - 1)
    return arr

"""
Merge Sort

Time Complexity: O(nlogn)
Space Complexity: O(n)
Inplace: No
Stable: Yes

How: Divide and Conquer
Idea: Divide the array into two halves, sort the two halves and merge them
"""


def merge_sort(arr):
    """Merge sort"""

    def merge(arr, l, m, r):
        lis = []
        i, j = l, m + 1

        while i <= m and j <= r:
            if arr[i] < arr[j]:
                lis.append(arr[i])
                i += 1
            else:
                lis.append(arr[j])
                j += 1

        if i >= m:
            while j <= r:
                lis.append(arr[j])
                j += 1
        if j >= r:
            while i <= m:
                lis.append(arr[i])
                i += 1
        arr[l:r + 1] = lis

    def merge_sort_helper(arr, l, r):
        if l < r:
            m = (l + (r - 1)) // 2
            merge_sort_helper(arr, l, m)
            merge_sort_helper(arr, m + 1, r)
            merge(arr, l, m, r)

    merge_sort_helper(arr, 0, len(arr) - 1)
    return arr


"""
Dutch National Flag algorithm

Time Complexity: O(n)
Space Complexity: O(1)
Inplace: Yes
Stable: Yes

What: Sorting an array of 0s, 1s and 2s
Idea: Maintain 3 pointers - low, mid and high such that low <= mid <= high
"""


def dutch_national_flag_algorithm(arr):
    """Dutch National Flag Algorithm to sort arr consisting of 0, 1 and 2 """
    n = len(arr)
    low, mid, high = 0, 0, n - 1
    while mid <= high:
        if arr[mid] == 0:
            arr[low], arr[mid] = arr[mid], arr[low]
            low += 1
            mid += 1
        elif arr[mid] == 1:
            mid += 1
        else:
            arr[mid], arr[high] = arr[high], arr[mid]
            high -= 1
    return arr


"""
QuickSelect

Time Complexity: O(n) average case, O(n^2) worst case
Space Complexity: O(1) iterative, O(logn) recursive
Inplace: Yes
Stable: No

What: Find the k-th smallest element in an unsorted array
How: Use partitioning similar to QuickSort but only recurse on one side
Idea: If pivot is at position k, we found the answer. If k < pivot position, search left. If k > pivot position, search right.
Applications: Find median, k-th largest/smallest, top-k elements
"""


def quickselect(arr, k):
    """
    Time Complexity: O(n) average, O(n^2) worst case
    Space Complexity: O(1)
    """
    def partition(arr, left, right):
        pivot = arr[right]
        i = left
        for j in range(left, right):
            if arr[j] < pivot:
                arr[i], arr[j] = arr[j], arr[i]
                i += 1
        arr[i], arr[right] = arr[right], arr[i]
        return i
    
    left, right = 0, len(arr) - 1
    
    while left <= right:
        pivot_index = partition(arr, left, right)
        
        if k == pivot_index:
            return arr[k]
        elif k < pivot_index:
            right = pivot_index - 1
        else:
            left = pivot_index + 1
    
    return arr[k]
