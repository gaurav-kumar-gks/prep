from __future__ import annotations
from collections.abc import Iterable, Iterator
from typing import Any


"""
Custom Iterators must inherit Iterator and implement the __next__ method, 
which returns the next item from the collection or raises StopIteration when there are no items left.

Iterable objects must inherit Iterable and implement the __iter__ method, 
which returns an Iterator.
"""


class User:

    def __init__(self, name, age):
        self.name = name
        self.age = age


class UsersIterator(Iterator):

    def __init__(self, users_list) -> None:
        self._users = users_list
        self._idx = 0

    def __next__(self) -> Any:
        try:
            value = self._users[self._idx]
            self._idx += 1
        except IndexError:
            raise StopIteration()

        return value


class UsersCollection(Iterable):

    def __init__(self, users_list=None):
        self._collection = users_list or []

    def __getitem__(self, index: int) -> Any:
        return self._collection[index]

    def __iter__(self):
        """
        The __iter__() method returns the iterator object itself, by default we
        return the iterator in ascending order.
        """
        return UsersIterator(self)

    def add_item(self, item: Any) -> None:
        self._collection.append(item)


if __name__ == "__main__":
    users = UsersCollection()
    users.add_item(User("Harry", 12))
    users.add_item(User("Ron", 14))

    print("Users: ")
    print("\n".join(users))
    print("")
