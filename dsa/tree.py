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
