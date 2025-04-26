import java.util.ArrayList;
import java.util.List;

interface Command {
    void execute();

    void undo();
}

class TypeTextCommand implements Command {
    private TextEditor editor;
    private String text;

    public TypeTextCommand(TextEditor editor, String text) {
        this.editor = editor;
        this.text = text;
    }

    @Override
    public void execute() {
        editor.typeText(text);
    }

    @Override
    public void undo() {
        editor.deleteText(text.length());
    }
}

class TextEditor {
    private String content = "";

    public void typeText(String text) {
        content += text;
    }

    public void deleteText(int length) {
        content = content.substring(0, content.length() - length);
    }

    @Override
    public String toString() {
        return content;
    }
}

class TextEditorInvoker {
    private List<Command> history = new ArrayList<>();
    private List<Command> redoStack = new ArrayList<>();

    public void executeCommand(Command command) {
        command.execute();
        history.add(command);
        redoStack.clear();
    }

    public void undo() {
        if (!history.isEmpty()) {
            Command command = history.remove(history.size() - 1);
            command.undo();
            redoStack.add(command);
        }
    }

    public void redo() {
        if (!redoStack.isEmpty()) {
            Command command = redoStack.remove(redoStack.size() - 1);
            command.execute();
            history.add(command);
        }
    }
}

public class CommandDemo {
    public static void main(String[] args) {
        TextEditor editor = new TextEditor();
        TextEditorInvoker invoker = new TextEditorInvoker();

        Command command1 = new TypeTextCommand(editor, "Hello, ");
        Command command2 = new TypeTextCommand(editor, "world!");

        invoker.executeCommand(command1);
        invoker.executeCommand(command2);
        System.out.println(editor); // Output: Hello, world!

        invoker.undo();
        System.out.println(editor); // Output: Hello,

        invoker.redo();
        System.out.println(editor); // Output: Hello, world!
    }
}