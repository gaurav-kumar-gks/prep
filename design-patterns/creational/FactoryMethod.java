
interface Product {
    String operation();
}

class ConcreteProduct1 implements Product {
    @Override
    public String operation() {
        return "{Result of the ConcreteProduct1}";
    }
}

class ConcreteProduct2 implements Product {
    @Override
    public String operation() {
        return "{Result of the ConcreteProduct2}";
    }
}

abstract class Creator {
    abstract Product factoryMethod();

    public String someOperation() {
        Product product = factoryMethod();
        return "Creator: The same creator's code has just worked with " + product.operation();
    }
}

class ConcreteCreator1 extends Creator {
    @Override
    Product factoryMethod() {
        return new ConcreteProduct1();
    }
}

class ConcreteCreator2 extends Creator {
    @Override
    Product factoryMethod() {
        return new ConcreteProduct2();
    }
}

public class FactoryMethod {
    public static void clientCode(Creator creator) {
        System.out.println("Client: I'm not aware of the creator's class, but it still works.\n" +
                creator.someOperation());
    }

    public static void main(String[] args) {
        System.out.println("App: Launched with the ConcreteCreator1.");
        clientCode(new ConcreteCreator1());
        System.out.println();

        System.out.println("App: Launched with the ConcreteCreator2.");
        clientCode(new ConcreteCreator2());
    }
}