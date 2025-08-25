"""
TRIE (PREFIX TREE)
"""

"""
1. TRIE PROPERTIES:
   - Each node represents a character
   - Common prefixes are shared among strings
   - Leaf nodes or marked nodes indicate end of words
   - Efficient for prefix-based operations
   - Can use more space than hash tables
   - Cache-unfriendly due to pointer chasing

6. TIME COMPLEXITIES:
   - Insert: O(m) where m is word length
   - Search: O(m) where m is word length
   - Prefix search: O(m) where m is prefix length
   - Delete: O(m) where m is word length
   - All Words with a prefix: O(k + m) where k is prefix length, m is total length of all matching words
"""

from collections import defaultdict


class TrieNode:
    def __init__(self):
        # here if ordering of children is not important
        # but if it is -> ordereddict can be used
        # if lexicographical ordering is required -> use []*26 array
        self.children = defaultdict(TrieNode)
        self.is_end_of_word = False


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        node = self.root
        for char in word:
            node = node.children[char]
        node.is_end_of_word = True

    def search(self, word):
        node = self.root
        for char in word:
            node = node.children.get(char)
            if not node:
                return False
        return node.is_end_of_word

    def starts_with(self, prefix):
        node = self.root
        for char in prefix:
            node = node.children.get(char)
            if not node:
                return False
        return True

    def delete(self, word):
        def delete_helper(node, word, index):
            if index == len(word):
                if not node.is_end_of_word:
                    return False
                node.is_end_of_word = False
                return not len(node.children)
            char = word[index]
            if char not in node.children:
                return False
            should_delete_child = delete_helper(node.children[char], word, index + 1)
            if should_delete_child:
                del node.children[char]
                return len(node.children) == 0 and not node.is_end_of_word
            return False
        delete_helper(self.root, word, 0)
    
    def get_all_words_with_prefix(self, prefix):
        """
        Time Complexity: O(k + m) where k is prefix length, m is total length of all matching words
        Space Complexity: O(m)
        """
        def collect_words(node, current_word, words):
            if not node: 
                return
            if node.is_end_of_word:
                words.append(current_word)
            for char, child in node.children.items():
                collect_words(child, current_word + char, words)
        
        node = self.root
        words = []
        for char in prefix:
            if char not in node.children:
                break
            node = node.children[char]
        collect_words(node, prefix, words)
        return words
    
    def longest_common_prefix(self):
        if not self.root.children:
            return ""
        prefix = ""
        node = self.root
        while len(node.children) == 1 and not node.is_end_of_word:
            char = list(node.children.keys())[0]
            prefix += char
            node = node.children[char]
        return prefix
