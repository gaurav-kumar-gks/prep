"""
Enums
"""
import datetime

from enum import Enum, Flag, auto


class Weekday(Enum):
    MONDAY = 1
    TUESDAY = 2
    WEDNESDAY = 3
    THURSDAY = 4
    FRIDAY = 5
    SATURDAY = 6
    SUNDAY = 7

    @classmethod
    def from_date(cls, date):
        return cls(date.isoweekday())


print(Weekday.MONDAY)  # Weekday.MONDAY
print(Weekday.MONDAY.value)  # 1
print(Weekday(1))  # Weekday.MONDAY
print(Weekday["MONDAY"])  # Weekday.MONDAY
print(Weekday(1) is Weekday.MONDAY)  # True
print(Weekday.from_date(datetime.date(2020, 1, 1)))  # Weekday.WEDNESDAY


class Color(Flag):
    RED = auto()
    GREEN = auto()
    BLUE = auto()


purple = Color.RED | Color.BLUE
white = Color.RED | Color.GREEN | Color.BLUE

print(purple)  # Color.RED|Color.BLUE
print(white)  # Color.RED|Color.GREEN|Color.BLUE
print(purple & white)  # Color.BLUE
print(purple | white)  # Color.RED|Color.GREEN|Color.BLUE
print(purple in white)  # False
print(Color.RED in white)  # True
