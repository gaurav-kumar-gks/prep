"""
Linked list
"""


class Node:
    def __init__(self, value, next=None):
        self.value = value
        self.next = next

    def __str__(self):
        return str(self.value)

    def __repr__(self):
        return str(self.value)

    def __eq__(self, other):
        return self.value == other.value

    def __ne__(self, other):
        return self.value != other.value


class SinglyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
        self.size = 0

    def __str__(self):
        return str(self.to_list())

    def __repr__(self):
        return str(self.to_list())

    def __len__(self):
        return self.size

    def __iter__(self):
        current = self.head
        while current:
            yield current
            current = current.next

    def __getitem__(self, index):
        return self.get(index)

    def __setitem__(self, index, value):
        self.set(index, value)

    def __delitem__(self, index):
        self.delete(index)

    def __eq__(self, other):
        if len(self) != len(other):
            return False
        for i in range(len(self)):
            if self[i] != other[i]:
                return False
        return True

    def __ne__(self, other):
        if len(self) != len(other):
            return True
        for i in range(len(self)):
            if self[i] != other[i]:
                return True
        return False

    def to_list(self):
        return [node.value for node in self]

    def get(self, index):
        if index < 0 or index >= self.size:
            raise IndexError('Index out of range')
        current = self.head
        for _ in range(index):
            current = current.next
        return current

    def set(self, index, value):
        if index < 0 or index >= self.size:
            raise IndexError('Index out of range')
        current = self.head
        for _ in range(index):
            current = current.next
        current.value = value

    def insert(self, index, value):
        if index < 0 or index > self.size:
            raise IndexError('Index out of range')
        new_node = Node(value)
        if index == 0:
            new_node.next = self.head
            self.head = new_node
        elif index == self.size:
            self.tail.next = new_node
            self.tail = new_node
        else:
            prev = self.get(index - 1)
            new_node.next = prev.next
            prev.next = new_node
        self.size += 1

    def append(self, value):
        self.insert(self.size, value)

    def prepend(self, value):
        self.insert(0, value)

    def delete(self, index):
        if index < 0 or index >= self.size:
            raise IndexError('Index out of range')
        if index == 0:
            self.head = self.head.next
        else:
            prev = self.get(index - 1)
            prev.next = prev.next.next
        self.size -= 1

    def pop(self):
        self.delete(self.size - 1)

    def remove(self, value):
        current = self.head
        prev = None
        while current:
            if current.value == value:
                if prev:
                    prev.next = current.next
                else:
                    self.head = current.next
                self.size -= 1
                return True
            prev = current
            current = current.next
        return False

    def reverse(self):
        current = self.head
        prev = None
        while current:
            next_node = current.next
            current.next = prev
            prev = current
            current = next_node
        self.head = prev

    def copy(self):
        new_list = SinglyLinkedList()
        for node in self:
            new_list.append(node.value)
        return new_list

    def clear(self):
        self.head = None
        self.tail = None
        self.size = 0

    def count(self, value):
        count = 0
        for node in self:
            if node.value == value:
                count += 1
        return count

    def index(self, value):
        for i, node in enumerate(self):
            if node.value == value:
                return i
        raise ValueError('Value not found')

    def extend(self, other):
        for node in other:
            self.append(node.value)


class DoublyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
        self.size = 0

    def __str__(self):
        return str(self.to_list())

    def __repr__(self):
        return str(self.to_list())

    def __len__(self):
        return self.size

    def __iter__(self):
        current = self.head
        while current:
            yield current
            current = current.next

    def __getitem__(self, index):
        return self.get(index)

    def __setitem__(self, index, value):
        self.set(index, value)

    def __delitem__(self, index):
        self.delete(index)

    def __eq__(self, other):
        if len(self) != len(other):
            return False
        for i in range(len(self)):
            if self[i] != other[i]:
                return False
        return True

    def __ne__(self, other):
        if len(self) != len(other):
            return True
        for i in range(len(self)):
            if self[i] != other[i]:
                return True
        return False

    def to_list(self):
        return [node.value for node in self]

    def get(self, index):
        if index < 0 or index >= self.size:
            raise IndexError('Index out of range')
        if index < self.size // 2:
            current = self.head
            for _ in range(index):
                current = current.next
        else:
            current = self.tail
            for _ in range(self.size - index - 1):
                current = current.prev
        return current

    def set(self, index, value):
        if index < 0 or index >= self.size:
            raise IndexError('Index out of range')
        node = self.get(index)
        node.value = value

    def insert(self, index, value):
        if index < 0 or index > self.size:
            raise IndexError('Index out of range')
        new_node = Node(value)
        if index == 0:
            if self.head:
                self.head.prev = new_node
                new_node.next = self.head
                self.head = new_node
            else:
                self.head = new_node
                self.tail = new_node
        elif index == self.size:
            self.tail.next = new_node
            new_node.prev = self.tail
            self.tail = new_node
        else:
            prev = self.get(index - 1)
            next_node = prev.next
            prev.next = new_node
            new_node.prev = prev
            new_node.next = next_node
            next_node.prev = new_node
        self.size += 1

    def append(self, value):
        self.insert(self.size, value)

    def prepend(self, value):
        self.insert(0, value)

    def delete(self, index):
        if index < 0 or index >= self.size:
            raise IndexError('Index out of range')
        if index == 0:
            self.head = self.head.next
            if self.head:
                self.head.prev = None
            else:
                self.tail = None
        elif index == self.size - 1:
            self.tail = self.tail.prev
            self.tail.next = None
        else:
            node = self.get(index)
            node.prev.next = node.next
            node.next.prev = node.prev
        self.size -= 1

    def pop(self):
        self.delete(self.size - 1)

    def remove(self, value):
        current = self.head
        while current:
            if current.value == value:
                if current.prev:
                    current.prev.next = current.next
                else:
                    self.head = current.next
                if current.next:
                    current.next.prev = current.prev
                else:
                    self.tail = current.prev
                self.size -= 1
                return True
            current = current.next
        return False

    def reverse(self):
        current = self.head
        while current:
            current.prev, current.next = current.next, current.prev
            current = current.prev
        self.head, self.tail = self.tail, self.head

    def copy(self):
        new_list = DoublyLinkedList()
        for node in self:
            new_list.append(node.value)
        return new_list

    def clear(self):
        self.head = None
        self.tail = None
        self.size = 0

    def count(self, value):
        count = 0
        for node in self:
            if node.value == value:
                count += 1
        return count

    def index(self, value):
        for i, node in enumerate(self):
            if node.value == value:
                return i
        raise ValueError('Value not found')

    def extend(self, other):
        for node in other:
            self.append(node.value)




