/**
 * Adapter is a structural design pattern that allows objects with incompatible
 * interfaces to collaborate.
 * It lets you wrap an otherwise incompatible object in an adapter to make it
 * compatible with another class.
 */

// Target interface that the client expects
interface Target {
    void request();
}

// Adaptee - the class that needs to be adapted
class Adaptee {
    public void specificRequest() {
        System.out.println("Adaptee: specificRequest");
    }
}

// Adapter - adapts the Adaptee to the Target interface
class Adapter implements Target {
    private Adaptee adaptee;

    public Adapter(Adaptee adaptee) {
        this.adaptee = adaptee;
    }

    @Override
    public void request() {
        System.out.println("Adapter: Converting Target request to Adaptee specificRequest");
        adaptee.specificRequest();
    }
}

// Client - uses the Target interface
class Client {
    public void clientCode(Target target) {
        System.out.println("Client: I can work with Target interface");
        target.request();
    }
}

public class AdapterDemo {
    public static void main(String[] args) {
        System.out.println("Client: I can work with Target interface");
        Target target = new Target() {
            @Override
            public void request() {
                System.out.println("Target: request");
            }
        };
        Client client = new Client();
        client.clientCode(target);

        System.out.println("\nClient: I can't work with Adaptee directly");
        Adaptee adaptee = new Adaptee();
        System.out.println("Adaptee: " + adaptee.getClass().getSimpleName());

        System.out.println("\nClient: But I can work with it via the Adapter");
        Adapter adapter = new Adapter(adaptee);
        client.clientCode(adapter);
    }
}