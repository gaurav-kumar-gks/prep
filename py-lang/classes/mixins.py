"""
A mixin is a class that provides methods to other classes
through inheritance but is not intended to stand on its own.

Mixins are used to add reusable functionality to classes
without using multiple inheritance in a way that can lead to
complex and hard-to-maintain code.

"""


class TimestampMixin:
    def get_timestamp(self):
        from datetime import datetime

        return datetime.now().isoformat()


class LoggingMixin:
    def log(self, message):
        print(f"[LOG]: {message}")


class User(LoggingMixin, TimestampMixin):
    def __init__(self, username):
        self.username = username

    def display_username(self):
        self.log(f"Displaying username: {self.username} at {self.get_timestamp()}")
        print(self.username)
