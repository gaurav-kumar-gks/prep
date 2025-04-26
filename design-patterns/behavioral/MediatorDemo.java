import java.util.List;
import java.util.ArrayList;

interface Mediator {
    void notify(Object sender, String event);
}

class ChatMediator implements Mediator {
    private List<Component> components = new ArrayList<>();

    public void addComponent(Component component) {
        components.add(component);
    }

    @Override
    public void notify(Object sender, String event) {
        for (Component component : components) {
            if (component != sender) {
                component.receive(event);
            }
        }
    }
}

abstract class Component {
    protected Mediator mediator;
    protected String name;

    public Component(String name) {
        this.name = name;
    }

    public void setMediator(Mediator mediator) {
        this.mediator = mediator;
    }

    public abstract void send(String message);

    public abstract void receive(String message);
}

class ChatUser extends Component {
    public ChatUser(String name) {
        super(name);
    }

    @Override
    public void send(String message) {
        System.out.println(name + " sends: " + message);
        mediator.notify(this, message);
    }

    @Override
    public void receive(String message) {
        System.out.println(name + " receives: " + message);
    }
}

public class MediatorDemo {
    public static void main(String[] args) {
        ChatMediator mediator = new ChatMediator();

        ChatUser alice = new ChatUser("Alice");
        ChatUser bob = new ChatUser("Bob");
        ChatUser charlie = new ChatUser("Charlie");

        alice.setMediator(mediator);
        bob.setMediator(mediator);
        charlie.setMediator(mediator);

        mediator.addComponent(alice);
        mediator.addComponent(bob);
        mediator.addComponent(charlie);

        System.out.println("Chat session started:");
        alice.send("Hello everyone!");
        System.out.println();
        bob.send("Hi Alice!");
        System.out.println();
        charlie.send("Hey team!");
    }
}