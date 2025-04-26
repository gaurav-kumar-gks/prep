import java.util.ArrayList;
import java.util.List;

/**
 * Composite is a structural design pattern that lets you compose objects into
 * tree structures
 * and then work with these structures as if they were individual objects.
 */

// Component interface
interface FileSystemComponent {
    void display(String indent);

    long getSize();
}

// Leaf class
class File implements FileSystemComponent {
    private String name;
    private long size;

    public File(String name, long size) {
        this.name = name;
        this.size = size;
    }

    @Override
    public void display(String indent) {
        System.out.println(indent + "File: " + name + " (" + size + " bytes)");
    }

    @Override
    public long getSize() {
        return size;
    }
}

// Composite class
class Directory implements FileSystemComponent {
    private String name;
    private List<FileSystemComponent> children;

    public Directory(String name) {
        this.name = name;
        this.children = new ArrayList<>();
    }

    public void add(FileSystemComponent component) {
        children.add(component);
    }

    public void remove(FileSystemComponent component) {
        children.remove(component);
    }

    @Override
    public void display(String indent) {
        System.out.println(indent + "Directory: " + name);
        for (FileSystemComponent component : children) {
            component.display(indent + "  ");
        }
    }

    @Override
    public long getSize() {
        long totalSize = 0;
        for (FileSystemComponent component : children) {
            totalSize += component.getSize();
        }
        return totalSize;
    }
}

public class CompositeDemo {
    public static void main(String[] args) {
        // Create files
        FileSystemComponent file1 = new File("document.txt", 100);
        FileSystemComponent file2 = new File("image.jpg", 200);
        FileSystemComponent file3 = new File("data.csv", 150);

        // Create directories
        Directory dir1 = new Directory("Documents");
        Directory dir2 = new Directory("Pictures");
        Directory root = new Directory("Root");

        // Build the tree structure
        dir1.add(file1);
        dir2.add(file2);
        dir2.add(file3);
        root.add(dir1);
        root.add(dir2);

        // Display the structure
        System.out.println("File System Structure:");
        root.display("");

        // Calculate total size
        System.out.println("\nTotal size: " + root.getSize() + " bytes");
    }
}