# Vehicle is an abstract class
from abc import ABC, abstractmethod


class Vehicle(ABC):
    def __init__(self, license_no):
        self.__license_no = license_no

    # ticket here refers to an instance of the ParkingTicket class
    @abstractmethod
    def assign_ticket(self, ticket):
        pass


class Car(Vehicle):
    # ticket here refers to an instance of the ParkingTicket class
    def __init__(self, license_no):
        super().__init__(license_no)

    # ticket here refers to an instance of the ParkingTicket class
    def assign_ticket(self, ticket):
        pass


class Van(Vehicle):
    # ticket here refers to an instance of the ParkingTicket class
    def __init__(self, license_no):
        super().__init__(license_no)

    # ticket here refers to an instance of the ParkingTicket class
    def assign_ticket(self, ticket):
        pass


class Truck(Vehicle):
    # ticket here refers to an instance of the ParkingTicket class
    def __init__(self, license_no, ticket):
        super().__init__(license_no, ticket)

    # ticket here refers to an instance of the ParkingTicket class
    def assign_ticket(self, ticket):
        pass


class MotorCycle(Vehicle):
    # ticket here refers to an instance of the ParkingTicket class
    def __init__(self, license_no, ticket):
        super().__init__(license_no, ticket)

    # ticket here refers to an instance of the ParkingTicket class
    def assign_ticket(self, ticket):
        pass
