"""
TWO POINTERS
"""

"""
Characteristics of problem that can be solved by two-pointer pattern:

note: the key here is the sub-problem that we are considering, according to that the below two statements may vary
If a wider scope of the window is valid, the narrower scope of that wider scope is valid must hold.
If a narrower scope of the window is invalid, the wider scope of that narrower scope is invalid must hold.
"""


"""
Running from both ends of array

initialize left & right pointers

while left < right:
    using left & right pointers, check something and then
    move left or right pointer, or both according to the problem
"""


"""
Fast and slow pointers

include a fast pointer and a slow pointer.
The fast pointer moves two nodes for every one node that the slow pointer moves.
When the fast pointer reaches the end of the linked list, the slow pointer will be in the middle.
"""


def find_middle(self):
    slow = self.head
    fast = self.head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    return slow.value


def detect_cycle(self):
    slow = self.head
    fast = self.head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return True
    return False


def find_cycle_start(self):
    slow = self.head
    fast = self.head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            current = self.head
            while current != slow:
                current = current.next
                slow = slow.next
            return current
    return None


"""
N ahead

The idea is to have two pointers, one that moves n nodes ahead of the other. 
When the first pointer reaches the end of the linked list, the second pointer will be at the nth node from the end.
"""


def removeNthFromEnd(head, n):
    fast = slow = head
    for _ in range(n):
        fast = fast.next
    if not fast:
        return head.next
    while fast.next:
        fast = fast.next
        slow = slow.next
    slow.next = slow.next.next
    return head
