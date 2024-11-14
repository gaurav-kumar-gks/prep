from __future__ import annotations
from abc import ABC


class Mediator(ABC):

    def notify(self, sender, event):
        pass


class ConcreteMediator(Mediator):
    def __init__(self, component1, component2) -> None:
        self._component1 = component1
        self._component1.mediator = self
        self._component2 = component2
        self._component2.mediator = self

    def notify(self, sender, event) -> None:
        print(f"{sender = }  {event = }")
        if event == "C1:A":
            self._component2.do_b()
        elif event == "C2:B":
            self._component1.do_c()


class BaseComponent:

    def __init__(self, mediator = None):
        self._mediator = mediator

    @property
    def mediator(self):
        return self._mediator

    @mediator.setter
    def mediator(self, mediator):
        self._mediator = mediator


class Component1(BaseComponent):
    def do_a(self) -> None:
        self.mediator.notify(self, "C1:A")

    def do_b(self) -> None:
        self.mediator.notify(self, "C1:B")

    def do_c(self) -> None:
        return


class Component2(BaseComponent):
    def do_b(self) -> None:
        self.mediator.notify(self, "C2:B")


if __name__ == "__main__":
    c1 = Component1()
    c2 = Component2()
    mediator = ConcreteMediator(c1, c2)

    print("Event C1:A ")
    c1.do_a()

    print("\n", end="")

    print("Event C1:B")
    c2.do_b()
