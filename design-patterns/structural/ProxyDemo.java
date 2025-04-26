/**
 * Proxy is a structural design pattern that lets you provide a substitute or
 * placeholder for another object.
 * A proxy controls access to the original object, allowing you to perform
 * something either before or after
 * the request gets through to the original object.
 */

// Subject interface
interface Image {
    void display();
}

// Real subject
class RealImage implements Image {
    private String filename;

    public RealImage(String filename) {
        this.filename = filename;
        loadFromDisk();
    }

    private void loadFromDisk() {
        System.out.println("Loading image: " + filename);
    }

    @Override
    public void display() {
        System.out.println("Displaying image: " + filename);
    }
}

// Proxy
class ProxyImage implements Image {
    private RealImage realImage;
    private String filename;

    public ProxyImage(String filename) {
        this.filename = filename;
    }

    @Override
    public void display() {
        if (realImage == null) {
            realImage = new RealImage(filename);
        }
        realImage.display();
    }
}

public class ProxyDemo {
    public static void main(String[] args) {
        Image image = new ProxyImage("test.jpg");

        // Image will be loaded from disk only when display() is called
        System.out.println("First time display:");
        image.display();

        // Image will not be loaded from disk again
        System.out.println("\nSecond time display:");
        image.display();
    }
}