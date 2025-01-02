"""
Trie

Applications:
- Auto-completion.
- Spell-checking.
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
