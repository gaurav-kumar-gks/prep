from abc import ABC, abstractmethod


class FileSystemElement(ABC):
    @abstractmethod
    def accept(self, visitor):
        pass


class File(FileSystemElement):
    def __init__(self, name, size):
        self.name = name
        self.size = size

    def accept(self, visitor):
        visitor.visit_file(self)


class Directory(FileSystemElement):
    def __init__(self, name):
        self.name = name
        self.children = []

    def add(self, element):
        self.children.append(element)

    def accept(self, visitor):
        visitor.visit_directory(self)


class Visitor(ABC):
    @abstractmethod
    def visit_file(self, file):
        pass

    @abstractmethod
    def visit_directory(self, directory):
        pass


class SizeCalculatorVisitor(Visitor):
    def __init__(self):
        self.total_size = 0

    def visit_file(self, file):
        self.total_size += file.size

    def visit_directory(self, directory):
        for child in directory.children:
            child.accept(self)


class ListFilesVisitor(Visitor):
    def __init__(self):
        self.files = []

    def visit_file(self, file):
        self.files.append(file.name)

    def visit_directory(self, directory):
        for child in directory.children:
            child.accept(self)


if __name__ == "__main__":
    root = Directory("root")
    file1 = File("file1.txt", 100)
    file2 = File("file2.txt", 200)
    subdir = Directory("subdir")
    subfile = File("subfile.txt", 50)

    root.add(file1)
    root.add(file2)
    root.add(subdir)
    subdir.add(subfile)

    size_calculator = SizeCalculatorVisitor()
    root.accept(size_calculator)
    print(f"Total size: {size_calculator.total_size}")

    list_files = ListFilesVisitor()
    root.accept(list_files)
    print("Files:", list_files.files)
