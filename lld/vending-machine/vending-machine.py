from abc import ABC, abstractmethod
from dataclasses import dataclass
from enum import Enum


class VendingMachine:
    def __init__(self):
        self.inventory = Inventory(3, 3, 5)
        self.cash_register = CashRegister(100)
        self.payment_gateway = Razorpay()
        self.state = ProductSelectionState(self)
        self.selected_product: None | Selection = None

    def select_product(self, selection: "Selection"):
        self.state.select_product(selection)

    def insert_money(self, amount: int):
        self.state.insert_money(amount)

    def set_state(self, state: "VendingState"):
        self.state = state

    def get_cash_register(self):
        return self.cash_register

    def get_payment_gateway(self):
        return self.payment_gateway


class VendingState(ABC):
    def __init__(self, vending_machine: "VendingMachine"):
        self.vending_machine = vending_machine

    @abstractmethod
    def select_product(self, selection: "Selection"):
        pass

    @abstractmethod
    def insert_money(self, amount: int):
        pass


class ProductSelectionState(VendingState):
    def select_product(self, selection: "Selection"):
        if not self.vending_machine.inventory._validate_pos(selection):
            print("Invalid selection")
        elif self.vending_machine.inventory._is_pos_empty(selection):
            print("No product at this position")
        else:
            product = self.vending_machine.inventory.get_product(selection)
            if product is None:
                print("No product at this position")
            else:
                self.vending_machine.selected_product = selection
                print(f"Product '{product.name}' selected, please insert money")
                self.vending_machine.set_state(
                    MoneyInsertionState(self.vending_machine)
                )

    def insert_money(self, amount: int):
        print("Please select a product first")


class MoneyInsertionState(VendingState):
    def __init__(self, vending_machine: "VendingMachine"):
        super().__init__(vending_machine)
        self.amount_inserted = 0

    def select_product(self, selection: "Selection"):
        print(
            "Product already selected, please insert money or cancel to select another product"
        )

    def insert_money(self, amount: int):
        self.amount_inserted += amount
        selected_product = self.vending_machine.selected_product
        if selected_product is None:
            print("No product selected")
            self.vending_machine.set_state(ProductSelectionState(self.vending_machine))
        product = self.vending_machine.inventory.get_product(selected_product)
        if product is None:
            print("Error: No product found at the selected position")
            return
        if self.amount_inserted < product.price:
            print("Not enough money inserted")
        else:
            self.vending_machine.set_state(
                DispenseState(self.vending_machine, self.amount_inserted)
            )

    def cancel_selection(self):
        print("Selection cancelled, please select another product")
        self.vending_machine.set_state(ProductSelectionState(self.vending_machine))


class DispenseState(VendingState):
    def __init__(self, vending_machine: "VendingMachine", amount_inserted: int):
        super().__init__(vending_machine)
        self.amount_inserted = amount_inserted

    def select_product(self, selection: "Selection"):
        print("Dispensing in progress, please wait")

    def insert_money(self, amount: int):
        print("Dispensing in progress, please wait")

    def dispense_product(self):
        product = self.vending_machine.inventory.remove_product(
            self.vending_machine.selected_product
        )
        change = self.amount_inserted - product.price
        if change > 0:
            if self.vending_machine.get_cash_register().check_balance() < change:
                print("Not enough change available, transaction cancelled")
                self.vending_machine.set_state(
                    ProductSelectionState(self.vending_machine)
                )
                return
            print(f"Dispensing {product} and returning change: {change}")
        else:
            print(f"Dispensing {product}")
        self.vending_machine.get_cash_register().deduct_money(change)
        self.vending_machine.set_state(ProductSelectionState(self.vending_machine))


class PaymentType(Enum):
    CASH = "cash"
    CARD = "card"


class PaymentRequest:
    def __init__(
        self,
        payment_type: "PaymentType",
        product: "Product",
        quantity: int,
        vending_machine: "VendingMachine",
        amount: int,
        request_id: str,
    ):
        self._payment_type = payment_type
        self._product = product
        self._quantity = quantity
        self._vending_machine = vending_machine
        self._amount = amount
        self._request_id = request_id

    def get_payment_type(self):
        return self._payment_type

    def get_quantity(self):
        return self._quantity

    def get_product(self):
        return self._product

    def get_vending_machine(self):
        return self._vending_machine

    def get_amount(self):
        return self._amount


class PaymentHandler:
    @staticmethod
    def handle(payment_request: "PaymentRequest"):
        payment_type = payment_request.get_payment_type()
        if payment_type == PaymentType.CASH:
            CashPaymentProcessor(payment_request).process_payment()
        elif payment_type == PaymentType.CARD:
            CardPaymentProcessor(payment_request).process_payment()
        else:
            print("Invalid payment type")


class PaymentGateway(ABC):
    @abstractmethod
    def handle_payment(self, amount: int):
        pass


class Razorpay(PaymentGateway):
    def handle_payment(self, amount: int):
        print(f"Received {amount} rupees via Razorpay")


class PaymentProcessor(ABC):
    def __init__(self, payment_request: "PaymentRequest") -> None:
        self.payment_request = payment_request

    @abstractmethod
    def process_payment(self):
        pass


class CashPaymentProcessor(PaymentProcessor):
    def process_payment(self):
        received_amount = self.payment_request.get_amount()
        amount_to_be_deducted = (
            self.payment_request.get_product().price
            * self.payment_request.get_quantity()
        )
        if received_amount < amount_to_be_deducted:
            print("Not enough money inserted")
            return
        amount_to_be_remitted = received_amount - amount_to_be_deducted
        cash_register = self.payment_request.get_vending_machine().get_cash_register()
        current_balance = cash_register.check_balance()
        if current_balance < amount_to_be_remitted:
            print("Not enough balance in cash register")
            return
        cash_register.add_money(received_amount)
        cash_register.deduct_money(amount_to_be_deducted)


class CardPaymentProcessor(PaymentProcessor):
    def process_payment(self):
        received_amount = self.payment_request.get_amount()
        print(f"Received {received_amount} rupees via card")
        payment_gateway = (
            self.payment_request.get_vending_machine()
            .get_payment_gateway()
            .handle_payment(received_amount)
        )


class CashRegister:
    def __init__(self, init_amount: int):
        self._amount = init_amount

    def check_balance(self):
        return self._amount

    def add_money(self, amount):
        self._amount += amount

    def deduct_money(self, amount):
        if self._amount < amount:
            print("Not enough change in cash register")
            return
        self._amount -= amount


@dataclass
class Selection:
    row: int
    col: int


class Inventory:
    def __init__(self, rows=0, cols=0, depth=0) -> None:
        self._rows = rows
        self._cols = cols
        self._depth = depth
        self._grid = [[[] for _ in range(cols)] for _ in range(rows)]

    def _validate_pos(self, selection: "Selection") -> bool:
        return 0 <= selection.row < self._rows and 0 <= selection.col < self._cols

    def _is_pos_empty(self, selection: "Selection") -> bool:
        return (
            self._validate_pos(selection)
            and len(self._grid[selection.row][selection.col]) == 0
        )

    def _is_pos_full(self, selection: "Selection") -> bool:
        return (
            self._validate_pos(selection)
            and len(self._grid[selection.row][selection.col]) == self._depth
        )

    def add_product(self, product: "Product", selection: "Selection") -> None:
        if not self._is_pos_full(selection):
            self._grid[selection.row][selection.col].append(product)

    def remove_product(self, selection: "Selection | None") -> "Product | None":
        if not selection:
            return None
        if not self._is_pos_empty(selection):
            return self._grid[selection.row][selection.col].pop()

    def get_product(self, selection: "Selection | None") -> "Product | None":
        if not selection:
            return None
        if not self._is_pos_empty(selection):
            return self._grid[selection.row][selection.col][-1]


class ProductType(Enum):
    DRINK = "drink"
    SNACK = "snack"
    CANDY = "candy"


class Product:
    def __init__(self, id, name, price, product_type) -> None:
        self._price = price
        self._name = name
        self._id = id
        self._product_type = product_type

    @property
    def price(self):
        return self._price

    @price.setter
    def price(self, price):
        self._price = price

    @property
    def name(self):
        return self._name

    @name.setter
    def name(self, name):
        self._name = name

    def __str__(self):
        return f"{self.name} - {self.price}"

    def __repr__(self):
        return f"Product({self._id}, {self.name}, {self.price})"


if __name__ == "__main__":
    vending_machine = VendingMachine()

    selection = Selection(0, 0)
    product = Product(1, "Soda", 25, ProductType.DRINK)
    vending_machine.inventory.add_product(product, selection)

    user_selection = Selection(3, 3)
    vending_machine.select_product(user_selection)

    user_selection = Selection(1, 1)
    vending_machine.select_product(user_selection)

    user_selection = Selection(0, 0)
    vending_machine.select_product(user_selection)
    vending_machine.insert_money(10)

    vending_machine.insert_money(20)
