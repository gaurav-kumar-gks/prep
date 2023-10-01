from abc import ABC, abstractmethod
from random import randrange
from typing import List


class Publisher(ABC):
    """
    The Publisher interface declares a set of methods for managing subscribers.
    """

    @abstractmethod
    def subscribe(self, subscriber: Subscriber) -> None:
        pass

    @abstractmethod
    def unsubscribe(self, subscriber: Subscriber) -> None:
        pass

    @abstractmethod
    def notify(self) -> None:
        pass


class ConcretePublisher(Publisher):
    """
    The Publisher owns some important state and notifies observers when the state
    changes.
    """

    _state: int = 0
    """
    For the sake of simplicity, the Subject's state, essential to all
    subscribers, is stored in this variable.
    """

    _subscribers: List[Subscriber] = []
    """
    List of subscribers. In real life, the list of subscribers can be stored
    more comprehensively (categorized by event type, etc.).
    """

    @property
    def state(self):
        return self._state

    @state.setter
    def state(self, value):
        self._state = value

    def subscribe(self, subscriber: Subscriber) -> None:
        print("Publisher: Added an subscriber.")
        self._subscribers.append(subscriber)

    def unsubscribe(self, subscriber: Subscriber) -> None:
        self._subscribers.remove(subscriber)

    """
    The subscription management methods.
    """

    def notify(self) -> None:
        """
        Trigger an update in each subscriber.
        """

        print("Subject: Notifying _subscribers...")
        for subscriber in self._subscribers:
            subscriber.update(self)

    def some_business_logic(self) -> None:
        """
        Usually, the subscription logic is only a fraction of what a Publisher can
        really do. Publishers commonly hold some important business logic, that
        triggers a notification method whenever something important is about to
        happen (or after it).
        """

        print("Publisher: I'm doing something important.")
        self._state = randrange(0, 10)

        print(f"Publisher: My state has just changed to: {self._state}")
        self.notify()


class Subscriber(ABC):
    """
    The Observer interface declares the update method, used by subjects.
    """

    @abstractmethod
    def update(self, publisher: Publisher) -> None:
        """
        Receive update from subject. Can pass
        """
        pass


"""
Concrete Subscribers react to the updates issued by the Publishers they have subbed to.
"""


class ConcreteSubscriberA(Subscriber):
    def update(self, publisher: Publisher) -> None:
        print("ConcreteSubscriberA updated")


class ConcreteSubscriberB(Subscriber):
    def update(self, publisher: Publisher) -> None:
        print("ConcreteSubscriberB updated")


if __name__ == "__main__":
    # The client code.

    publisher = ConcretePublisher()

    subscriber_a = ConcreteSubscriberA()
    publisher.subscribe(subscriber_a)

    subscriber_b = ConcreteSubscriberB()
    publisher.subscribe(subscriber_b)

    publisher.some_business_logic()
    publisher.some_business_logic()

    publisher.unsubscribe(subscriber_a)

    publisher.some_business_logic()
