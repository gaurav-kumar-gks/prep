"""
Context managers

Used to manage resources and ensure proper setup and cleanup. 
They are typically used with the "with" statement and 
follow the context management protocol by implementing the 
    __enter__() and __exit__() methods
"""

class FileManager:
    def __init__(self, filename, mode):
        self.filename = filename
        self.mode = mode

    def __enter__(self):
        self.file = open(self.filename, self.mode)
        return self.file

    def __exit__(self, exc_type, exc_value, traceback):
        self.file.close()

# Usage:
with FileManager('example.txt', 'w') as file:
    file.write('Hello, world!')