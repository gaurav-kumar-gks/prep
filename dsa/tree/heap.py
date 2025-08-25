"""
HEAP / PRIORITY QUEUE
"""

"""
HEAP PROPERTIES:
    - Complete Binary Tree: All levels are filled except possibly the last level
    - Max Heap: Parent >= Children (root is maximum)
    - Min Heap: Parent <= Children (root is minimum)

ARRAY REPRESENTATION:
    - Parent index: (i-1) // 2
    - Left child: 2*i + 1
    - Right child: 2*i + 2
    - Leaf nodes: from n//2 to n-1

TIME COMPLEXITIES:
    - Get max/min: O(1)
    - Insert: O(log n)
    - Delete (extract max/min): O(log n)
    - Heapify: O(log n)
    - Build heap: O(n)
"""

class MinHeap:
    """
    Min Heap Implementation
    
    A min heap is a complete binary tree where each parent node is smaller than or equal to its children.
    The root node contains the minimum element.
    """
    
    def __init__(self):
        self.heap = []
    
    def parent(self, index):
        return (index - 1) // 2
    
    def left_child(self, index):
        return 2 * index + 1
    
    def right_child(self, index):
        return 2 * index + 2
    
    def has_parent(self, index):
        return self.parent(index) >= 0
    
    def has_left_child(self, index):
        return self.left_child(index) < len(self.heap)
    
    def has_right_child(self, index):
        return self.right_child(index) < len(self.heap)
    
    def swap(self, index1, index2):
        self.heap[index1], self.heap[index2] = self.heap[index2], self.heap[index1]
    
    def peek(self):
        if len(self.heap) == 0:
            raise IndexError("Heap is empty")
        return self.heap[0]
    
    def size(self):
        return len(self.heap)
    
    def insert(self, value):
        self.heap.append(value)
        self._bubble_up(len(self.heap) - 1)
    
    def _bubble_up(self, index):
        """
        Move element up to maintain heap property.
        Compare with parent and swap if necessary until heap property is satisfied.
        """
        while self.has_parent(index) and self.heap[self.parent(index)] > self.heap[index]:
            self.swap(index, self.parent(index))
            index = self.parent(index)
    
    def extract_min(self):
        if len(self.heap) == 0:
            raise IndexError("Heap is empty")
        min_value = self.heap[0]
        self.heap[0] = self.heap[-1]
        self.heap.pop()
        if len(self.heap) > 0:
            self._bubble_down(0)
        return min_value
    
    def _bubble_down(self, index):
        """
        Move element down to maintain heap property.
        Compare with children and swap with smaller child until heap property is satisfied.
        """
        while True:
            small_child = None
            l, r = self.has_left_child(index), self.has_right_child(index):
            if l and r:
                small_child = l if self.heap[self.left_child[index]] < self.heap[self.right_child[index]] else r
            elif l or r:
                small_child = l if l else r
            if self.heap[index] <= self.heap[small_child]: 
                break
            self.swap(index, small_child)
            index = small_child
    
    def delete(self, value):
        """
        Delete a specific value from the heap.
        Time Complexity: O(n) for finding + O(log n) for heapify = O(n)
        Space Complexity: O(1)
        """
        try:
            index = self.heap.index(value)
        except ValueError:
            raise ValueError("Value not found in heap")
        self.heap[index] = self.heap[-1]
        self.heap.pop()
        if len(self.heap) > 0:
            if index == 0 or self.heap[self.parent(index)] <= self.heap[index]:
                self._bubble_down(index)
            else:
                self._bubble_up(index)
    
    def heapify(self, array):
        """
        Convert an array into a min heap.
        Algorithm:
        1. Start from the last non-leaf node (n//2 - 1)
        2. Bubble down each node to its correct position
        3. Work backwards to the root
        Time Complexity: O(n)
        Space Complexity: O(1)
        """
        self.heap = array.copy()
        n = len(self.heap)
        for i in range(n // 2 - 1, -1, -1):
            self._bubble_down(i)
    
    def heap_sort(self):
        """
        Sort the heap in ascending order.
        Algorithm:
        1. Repeatedly extract minimum elements
        2. Store them in a result array
        Time Complexity: O(n log n)
        Space Complexity: O(n)
        """
        sorted_array = []
        while self.size() > 0:
            sorted_array.append(self.extract_min())
        return sorted_array

"""
COMMON APPLICATIONS

1. FIND K-TH LARGEST/SMALLEST ELEMENT
   - Use min heap for k-th largest
   - Use max heap for k-th smallest
   - Time: O(n log k)

2. MERGE K SORTED LISTS
   - Use min heap to get minimum element from all lists
   - Time: O(n log k) where n is total elements, k is number of lists

4. FIND MEDIAN FROM DATA STREAM
   - Use two heaps: max heap for lower half, min heap for upper half
   - Time: O(log n) per operation

6. DIJKSTRA'S ALGORITHM
   - Use priority queue to get minimum distance vertex
   - Time: O((V + E) log V)

7. PRIM'S ALGORITHM
   - Use priority queue to get minimum weight edge
   - Time: O(E log V)

8. HUFFMAN CODING
   - Use priority queue to build Huffman tree
   - Time: O(n log n)
"""