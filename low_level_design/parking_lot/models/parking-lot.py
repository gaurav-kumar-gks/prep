# The __ParkingLot is a singleton class that ensures it will have only one active instance at a time
# Both the Entrance and Exit classes use this class to create and close parking tickets
class __ParkingLot(object):
    __instances = None

    def __new__(cls):
        if cls.__instances is None:
            cls.__instances = super(__ParkingLot, cls).__new__(cls)
        return cls.__instances


class ParkingLot(metaclass=__ParkingLot):
    def __init__(self, id, name, address, parking_rate):
        # Call the name, address, parking_rate elements of the customer in the parking lot from the database
        self.__id = id
        self.__name = name
        self.__address = address
        self.__parking_rate = parking_rate

        # Create initial entrance and exit hashmaps respectively
        self.__entrance = {}
        self.__exit = {}

        # Create a hashmap that identifies all currently generated tickets using their ticket number
        self.__tickets = {}

    # entrance here refers to an instance of the Entrance class
    def add_entrance(self, entrance):
        pass

    # exit here refers to an instance of the Exit class
    def add_exit(self, exit):
        pass

    # This function allows parking tickets to be available at multiple entrances
    # vehicle here refers to an Vehicle of the Exit class
    # Returns a ParkingTicket instance
    def get_parking_ticket(self, vehicle):
        pass

    # type here refers to an instance of the ParkingSpot class
    def is_full(self, type):
        pass
