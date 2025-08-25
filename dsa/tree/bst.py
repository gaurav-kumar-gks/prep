"""
BINARY SEARCH TREE (BST)
"""

"""
BST PROPERTIES:
    - Each node contains a key (value)
    - Left subtree of a node contains only nodes with keys < node's key
    - Right subtree of a node contains only nodes with keys > node's key
    - Both left and right subtrees must also be binary search trees
    - No duplicate keys (in standard BST)
    - Height can be O(n) in worst case (skewed / unbalanced)
    - Inorder traversal gives sorted order
   

TIME COMPLEXITIES:
    - Search: O(h) where h is height (O(log n) for balanced, O(n) for skewed)
    - Insert: O(h)
    - Delete: O(h)
    - Inorder Traversal: O(n)
    - Find Min/Max: O(h)
"""

class TreeNode:
    def __init__(self, key):
        self.key = key
        self.left = None
        self.right = None
        self.parent = None


class BinarySearchTree:
    """
    Binary Search Tree Implementation
    
    A BST maintains the property that for any node:
    - All nodes in left subtree have keys < node's key
    - All nodes in right subtree have keys > node's key
    """
    
    def __init__(self):
        self.root = None
        self.size = 0
    
    def is_empty(self):
        return self.root is None
    
    def get_size(self):
        return self.size
    
    def insert(self, x, node=None, parent=None):
        """
        Insert a new key into the BST.
        Time Complexity: O(h) where h is height of tree
        Space Complexity: O(h) for recursion stack
        """
        if not node:
            node = TreeNode(x)
            node.parent = parent
            self.size += 1
            return node
        if x < node.val:
            node.left = self.insert(x, node.left, node)
        elif x > node.val: 
            node.right = self.insert(x, node.right, node)
        return node
            
    def search_recursive(self, node, key):
        if node is None or node.key == key:
            return node
        if key < node.key:
            return self.search_recursive(node.left, key)
        return self.search_recursive(node.right, key)
    
    def search(self, key):
        current = self.root
        while current is not None:
            if key == current.key:
                return current
            elif key < current.key:
                current = current.left
            else:
                current = current.right
        return None
    
    def successor(self, key):
        """
        Find the successor (next larger element) of a given key.
         successor
        /
        ancestor
        \
         node
        """
        node = self.search(key)
        if node is None:
            return None
        if node.right:
            return self.find_min(node.right)
        current = node
        parent = current.parent
        while parent is not None and current == parent.right:
            current = parent
            parent = parent.parent
        return parent
    
    def predecessor(self, key):
        """
        Find the predecessor (next smaller element) of a given key.
        predecessor
          \
           ancestor  
           /
        node
        """
        node = self.search(key)
        if node is None:
            return None
        if node.left is not None:
            return self.find_max(node.left)
        current = node
        parent = current.parent
        while parent is not None and current == parent.left:
            current = parent
            parent = parent.parent
        return parent
    
    def delete(self, key):
        """
        Delete a key from the BST.
        
        Algorithm:
        1. Find the node to delete
        2. Case 1: Node is a leaf - simply remove it
        3. Case 2: Node has one child - replace with child
        4. Case 3: Node has two children - replace with successor/predecessor
        
        Time Complexity: O(h)
        Space Complexity: O(h)
        """
        self.root = self._delete_recursive(self.root, key)
    
    def _delete_recursive(self, node, key):
        if not node: return None
        if node.val < key:
            node.right = self._delete_recursive(node.right, key)
            return node
        elif node.val > key:
            node.left = self._delete_recursive(node.left, key)
            return node
        else:
            self.size -= 1
            if node.left is None and node.right is None:
                return None
            elif node.left is None:
                return node.right
            elif node.right is None:
                return node.left
            else:
                successor = self.find_min(node.right)
                node.key = successor.key
                node.right = self._delete_recursive(node.right, successor.key)
        return node
