import threading
import logging
from abc import ABC, abstractmethod


class State(ABC):
    def __init__(self, vending_machine):
        self.vending_machine = vending_machine

    @abstractmethod
    def select_item(self):
        pass

    @abstractmethod
    def add_item(self, item: str):
        pass

    @abstractmethod
    def add_money(self, amount: float):
        pass

    @abstractmethod
    def dispense_item(self):
        pass


class HasItemState(State):
    def select_item(self):
        print("Item selected.")
        self.vending_machine.set_state(self.vending_machine.item_requested_state)

    def add_item(self, item: str):
        print("Item already present. Cannot add another item.")

    def add_money(self, amount: float):
        print("Please select an item first.")

    def dispense_item(self):
        print("Please select an item first.")


class ItemRequestedState(State):
    def select_item(self):
        print("Item already selected.")

    def add_item(self, item: str):
        print("Cannot add item. Item already selected.")

    def add_money(self, amount: float):
        print(f"Money added: {amount}")
        self.vending_machine.set_state(self.vending_machine.has_money_state)

    def dispense_item(self):
        print("Please add money first.")


class HasMoneyState(State):
    def select_item(self):
        print("Item already selected.")

    def add_item(self, item: str):
        print("Cannot add item. Item already selected.")

    def add_money(self, amount: float):
        print("Money already added.")

    def dispense_item(self):
        print("Dispensing item...")
        self.vending_machine.set_state(self.vending_machine.no_item_state)


class NoItemState(State):
    def select_item(self):
        print("No item available.")

    def add_item(self, item: str):
        print(f"Item added: {item}")
        self.vending_machine.set_state(self.vending_machine.has_item_state)

    def add_money(self, amount: float):
        print("No item available to add money for.")

    def dispense_item(self):
        print("No item available.")


class VendingMachine:
    def __init__(self):
        self.lock = threading.Lock()
        self.has_item_state = HasItemState(self)
        self.item_requested_state = ItemRequestedState(self)
        self.has_money_state = HasMoneyState(self)
        self.no_item_state = NoItemState(self)
        self.state = self.no_item_state

    def set_state(self, state: State):
        with self.lock:
            self.state = state

    def select_item(self):
        with self.lock:
            self.state.select_item()

    def add_item(self, item: str):
        with self.lock:
            self.state.add_item(item)

    def add_money(self, amount: float):
        with self.lock:
            if amount <= 0:
                logging.error("Invalid amount. Must be greater than zero.")
                return
            self.state.add_money(amount)

    def dispense_item(self):
        with self.lock:
            self.state.dispense_item()


if __name__ == "__main__":
    vending_machine = VendingMachine()

    vending_machine.add_item("Soda")
    vending_machine.select_item()
    vending_machine.add_money(1.0)
    vending_machine.dispense_item()
