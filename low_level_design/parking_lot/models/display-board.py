class DisplayBoard:
    def __init__(self, id, handicapped_spot, compact_spot, large_spot, motor_cycle_spot):
        self.__id = id
        self.__handicapped_spot = handicapped_spot  # Refers to an instance of the Handicapped class
        self.__compact_spot = compact_spot  # Refers to an instance of the Compact class
        self.__large_spot = large_spot  # Refers to an instance of the Large class
        self.__motor_cycle_spot = motor_cycle_spot  # Refers to an instance of the Motorcycle class

    # Member functions
    def show_free_slot(self):
        pass


class ParkinRate:
    def __init__(self, hours, rate):
        self.__hours = hours
        self.__rate = rate

    # Member function
    def calculate(self):
        pass


class Entrance:
    def __init__(self, id, ticket):
        self.__id = id

    # ticket here refers to an instance of the ParkingTicket class
    def get_ticket(self):
        pass

class Exit:
    def __init__(self, id, ticket):
        self.__id = id

    # ticket here refers to an instance of the ParkingTicket class
    def validate_ticket(self, ticket):
        # Perform validation logic for the parking ticket
        # Calculate parking charges, if necessary
        # Handle the exit process
        pass
