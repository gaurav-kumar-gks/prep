"""
TREE
"""

"""
TREE TERMINOLOGY:
    - Root: Topmost node of the tree
    - Parent: Node that has children
    - Child: Node connected to a parent
    - Leaf: Node with no children
    - Internal Node: Node with at least one child
    - Sibling: Nodes with the same parent
    - Ancestor: All nodes on path from root to current node
    - Descendant: All nodes in subtree rooted at current node
    - Height: Length of longest path from root to leaf
    - Depth: Length of path from root to specific node
    - Degree: Number of children a node has


    BALANCED TREE
        - Height difference between left and right subtrees is limited
        - Examples: AVL Tree, Red-Black Tree, B-Tree
        - Advantages: Guaranteed O(log n) operations
        - Applications: Database indexing, File systems
        
    UNBALANCED TREE
        - No height restrictions
        - Can degenerate to linked list (height = n)
        - Disadvantages: O(n) worst-case operations

    COMPLETE TREE
        - All levels are filled except possibly the last level
        - Last level has all nodes as far left as possible
        - Used in: Heap, Priority Queue

    PERFECT TREE
        - All internal nodes have exactly 2 children
        - All leaves are at the same level
        - Used in: BST, AVL Tree

    BINARY TREE
        - Each node has at most 2 children (left and right)
        - Used in: BST, AVL, Red-Black trees, Heaps
        
    TERNARY TREE
        - Each node has at most 3 children
        - Used in: Some search algorithms, Game trees
        
    N-ARY TREE (GENERAL TREE)
        - Each node can have any number of children
        - Used in: File systems, Organization charts, XML/HTML DOM
"""
class TreeNode:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

    def inorder_traversal(self, result):
        """
        We visit the left subtree first, then the root node, and finally the right subtree.

        Time complexity: O(n), where n is the number of nodes in the tree.
        Space complexity: O(h), where h is the height of the tree.
        
        Example:
        --------
        Given the following binary tree:

                    1
                   / \
                  2   3
                 / \
                4   5
                
        The inorder traversal of the tree is [4, 2, 5, 1, 3].
        """
        if self.left:
            self.left.inorder_traversal(result)
        result.append(self.val)
        if self.right:
            self.right.inorder_traversal(result)

    def inorder_traversal_iterative(self):
        """
        Time complexity: O(n), where n is the number of nodes in the tree.
        Space complexity: O(h), where h is the height of the tree.
        """
        result = []
        stack = []
        curr = self
        while curr or stack:
            while curr:
                stack.append(curr)
                curr = curr.left
            curr = stack.pop()
            result.append(curr.val)
            curr = curr.right
        return result


    def preorder_traversal(self, result):
        """
        This method visits the root node first, then the left subtree, and finally the right subtree.
        It appends the value of each visited node to the given result list.

        Time complexity: O(n), where n is the number of nodes in the tree.
        Space complexity: O(h), where h is the height of the tree.

        Example:
        --------
        Given the following binary tree:

                 1
                / \
               2   3
              / \
             4   5

        The preorder traversal of the tree is [1, 2, 4, 5, 3].
        """
        if self is None:
            return
        result.append(self.val)
        if self.left is not None:
            self.left.preorder_traversal(result)
        if self.right is not None:
            self.right.preorder_traversal(result)

    def preorder_traversal_iterative(self):
        """
        Time complexity: O(n), where n is the number of nodes in the tree.
        Space complexity: O(h), where h is the height of the tree.
        """
        result = []
        stack = [self]
        while stack:
            curr = stack.pop()
            result.append(curr.val)
            if curr.right:
                stack.append(curr.right)
            if curr.left:
                stack.append(curr.left)
        return result
    
    
    def postorder_traversal(self, result):
        """
        Perform a postorder traversal of the tree recursively.

        Time complexity: O(n), where n is the number of nodes in the tree.
        Space complexity: O(h), where h is the height of the tree.
        
        Example:
        --------
        Given the following binary tree:

                1
               / \
              2   3
             / \
            4   5
                
        The postorder traversal of the tree is [4, 5, 2, 3, 1].
        """
        if self.left:
            self.left.postorder_traversal(result)
        if self.right:
            self.right.postorder_traversal(result)
        result.append(self.val)

    def postorder_traversal_iterative(self):
        """
        Perform a postorder traversal of the tree iteratively using two stacks.

        Time complexity: O(n), where n is the number of nodes in the tree.
        Space complexity: O(h), where h is the height of the tree.
        """
        result = []
        stack1 = [self]
        stack2 = []
        while stack1:
            curr = stack1.pop()
            stack2.append(curr)
            if curr.left:
                stack1.append(curr.left)
            if curr.right:
                stack1.append(curr.right)
        while stack2:
            result.append(stack2.pop().val)

        return result
    
    def level_order_traversal_iterative(self, root):
        """
        Perform a level order traversal of the binary tree rooted at `root`.

        Time complexity: O(n), where n is the number of nodes in the tree.
        Space complexity: O(n), where n is the number of nodes in the tree.
        
        Example:
        --------
        Given the following binary tree:

                1
               / \
              2   3
             / \
            4   5
            
        The level order traversal of the tree is [[1], [2, 3], [4, 5]].
        """
        if not root:
            return []
        result = []
        queue = [root]
        while queue:
            level = []
            for i in range(len(queue)):
                node = queue.pop(0)
                level.append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            result.append(level)
        
        # import deque
        # q, result = deque([root]) if root else [], []
        # while q:
        #     node = q.popleft()
        #     result.append((node := q.popleft()).val)
        #     q.extend([kid for kid in (node.left, node.right) if kid])
        return result
    

    def level_order_traversal(self, result=None, level=0):
        if result is None:
            result = []
        if level == len(result):
            result.append([])
        result[level].append(self.val)
        if self.left:
            self.left.level_order_traversal(result, level + 1)
        if self.right:
            self.right.level_order_traversal(result, level + 1)
        return result

    def morris_traversal(self):
        res = []
        while self:
            if not self.left:
                res.append(self.val)
                self = self.right
            else:
                # find the inorder predecessor of the current node
                pre = self.left
                while pre and pre.right != self:
                    pre = pre.right
                if not pre.right:
                    pre.right = self
                    self = self.left
                else:
                    pre.right = None
                    res.append(self.val)
                    self = self.right

"""
Different Trees

    1. BINARY SEARCH TREE (BST)
        - Left subtree < Root < Right subtree
        - Inorder traversal gives sorted order
        - Applications: Symbol tables, Database indexing
        - Time Complexity: O(log n) average, O(n) worst case
        
    2. HEAP
        - Complete binary tree with heap property
        - Max Heap: Parent >= Children
        - Min Heap: Parent <= Children
        - Applications: Priority queues, Heap sort
        - Time Complexity: O(log n) for insert/delete, O(1) for max/min
        
    3. AVL TREE
        - Self-balancing BST
        - Height difference between subtrees ≤ 1
        - Rotations maintain balance
        - Applications: Database systems, Real-time systems
        - Time Complexity: O(log n) guaranteed
        
    4. RED-BLACK TREE
        - Self-balancing BST with color properties
        - Root is black, red nodes have black children
        - Black height is same for all paths
        - Applications: C++ STL map/set, Java TreeMap/TreeSet
        - Time Complexity: O(log n) guaranteed
        
    5. SEGMENT TREE
        - Tree for range queries on arrays
        - Each node represents a segment/interval
        - Supports range sum, min, max, etc.
        - Applications: Range queries, Competitive programming
        - Time Complexity: O(log n) for queries and updates
        
    6. FENWICK TREE (BINARY INDEXED TREE)
        - Efficient prefix sum queries
        - Uses bit manipulation for indexing
        - Applications: Range sum queries, Inversion counting
        - Time Complexity: O(log n) for queries and updates
        
    7. SUFFIX TREE
        - Compressed trie of all suffixes of a string
        - Used for string pattern matching
        - Applications: DNA sequence analysis, Text indexing
        - Time Complexity: O(n) construction, O(m) pattern search
        
    8. TRIE (PREFIX TREE)
        - Tree for storing strings
        - Common prefixes are shared
        - Applications: Auto-complete, Spell checking, IP routing
        - Time Complexity: O(m) for search/insert where m is string length
"""