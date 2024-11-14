from abc import ABC, abstractmethod
from typing import Optional


class LoggerHandler(ABC):

    @abstractmethod
    def log_message(self, level: int, message: str):
        pass

    @abstractmethod
    def set_next(self, handler: 'LoggerHandler'):
        pass


class Logger(LoggerHandler):
    def __init__(self, level: int):
        self.level = level
        self._next_logger: Optional[Logger] = None

    def set_next(self, next_logger: 'Logger'):
        self._next_logger = next_logger
        return next_logger

    def log_message(self, level: int, message: str):
        if self.level <= level:
            self.write(message)
        if self._next_logger:
            self._next_logger.log_message(level, message)

    @abstractmethod
    def write(self, message: str):
        pass


class ConsoleLogger(Logger):
    def __init__(self):
        super().__init__(1)  # INFO level

    def write(self, message: str):
        print(f"Console Logger: {message}")


class FileLogger(Logger):
    def __init__(self):
        super().__init__(2)  # DEBUG level

    def write(self, message: str):
        print(f"File Logger: {message}")


class ErrorLogger(Logger):
    def __init__(self):
        super().__init__(3)  # ERROR level

    def write(self, message: str):
        print(f"Error Logger: {message}")


if __name__ == "__main__":
    console_logger = ConsoleLogger()
    file_logger = FileLogger()
    error_logger = ErrorLogger()
    console_logger.set_next(file_logger).set_next(error_logger)

    console_logger.log_message(1, "This is an information message.")
    console_logger.log_message(2, "This is a debug message.")
    console_logger.log_message(3, "This is an error message.")
