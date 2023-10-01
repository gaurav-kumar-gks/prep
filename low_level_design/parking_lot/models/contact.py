# create contact class considering common use cases for lld design round, note that we already made a address class, follow best oops practices and use type hints
from address import Address


class Contact:
    def __init__(self, name: str, phone_number: str, email: str, address: Address):
        self.name = name
        self.phone_number = phone_number
        self.email = email
        self.address = address
