
abstract class Handler {
    private Handler next;
    protected String name;
    protected int threshold;

    public Handler(String name, int threshold) {
        this.name = name;
        this.threshold = threshold;
    }

    public Handler setNext(Handler next) {
        this.next = next;
        return next;
    }

    public void handle(int request) {
        if (request <= threshold) {
            processRequest(request);
        } else if (next != null) {
            System.out.println(name + " cannot handle request " + request + ", passing to " + next.name);
            next.handle(request);
        } else {
            System.out.println("No handler available for request " + request);
        }
    }

    protected abstract void processRequest(int request);
}

class ConcreteHandler1 extends Handler {
    public ConcreteHandler1() {
        super("Handler1", 10);
    }

    @Override
    protected void processRequest(int request) {
        System.out.println("Handler1 processing request " + request);
    }
}

class ConcreteHandler2 extends Handler {
    public ConcreteHandler2() {
        super("Handler2", 20);
    }

    @Override
    protected void processRequest(int request) {
        System.out.println("Handler2 processing request " + request);
    }
}

class ConcreteHandler3 extends Handler {
    public ConcreteHandler3() {
        super("Handler3", 30);
    }

    @Override
    protected void processRequest(int request) {
        System.out.println("Handler3 processing request " + request);
    }
}

public class ChainOfResponsibilityDemo {
    public static void main(String[] args) {
        // Create the chain of handlers
        Handler handler1 = new ConcreteHandler1();
        Handler handler2 = new ConcreteHandler2();
        Handler handler3 = new ConcreteHandler3();

        handler1.setNext(handler2).setNext(handler3);

        // Process various requests
        System.out.println("Processing request 5:");
        handler1.handle(5);

        System.out.println("\nProcessing request 15:");
        handler1.handle(15);

        System.out.println("\nProcessing request 25:");
        handler1.handle(25);

        System.out.println("\nProcessing request 35:");
        handler1.handle(35);
    }
}