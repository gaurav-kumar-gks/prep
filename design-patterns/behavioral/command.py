from abc import ABC, abstractmethod


class Command(ABC):
    @abstractmethod
    def execute(self):
        pass

    @abstractmethod
    def undo(self):
        pass


class TypeTextCommand(Command):
    def __init__(self, editor, text):
        self.editor = editor
        self.text = text

    def execute(self):
        self.editor.type_text(self.text)

    def undo(self):
        self.editor.delete_text(len(self.text))


class TextEditor:
    def __init__(self):
        self.content = ""

    def type_text(self, text):
        self.content += text

    def delete_text(self, length):
        self.content = self.content[:-length]

    def __str__(self):
        return self.content


class TextEditorInvoker:
    def __init__(self):
        self.history = []
        self.redo_stack = []

    def execute_command(self, command):
        command.execute()
        self.history.append(command)
        self.redo_stack.clear()

    def undo(self):
        if self.history:
            command = self.history.pop()
            command.undo()
            self.redo_stack.append(command)

    def redo(self):
        if self.redo_stack:
            command = self.redo_stack.pop()
            command.execute()
            self.history.append(command)


if __name__ == "__main__":
    editor = TextEditor()
    invoker = TextEditorInvoker()

    command1 = TypeTextCommand(editor, "Hello, ")
    command2 = TypeTextCommand(editor, "world!")

    invoker.execute_command(command1)
    invoker.execute_command(command2)
    print(editor)  # Output: Hello, world!

    invoker.undo()
    print(editor)  # Output: Hello,

    invoker.redo()
    print(editor)  # Output: Hello, world!
