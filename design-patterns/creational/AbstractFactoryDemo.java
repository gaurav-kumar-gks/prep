
interface AbstractProductA {
    String usefulFunctionA();
}

class ConcreteProductA1 implements AbstractProductA {
    @Override
    public String usefulFunctionA() {
        return "The result of the product A1.";
    }
}

class ConcreteProductA2 implements AbstractProductA {
    @Override
    public String usefulFunctionA() {
        return "The result of the product A2.";
    }
}

interface AbstractProductB {
    String usefulFunctionB();

    String anotherUsefulFunctionB(AbstractProductA collaborator);
}

class ConcreteProductB1 implements AbstractProductB {
    @Override
    public String usefulFunctionB() {
        return "The result of the product B1.";
    }

    @Override
    public String anotherUsefulFunctionB(AbstractProductA collaborator) {
        String result = collaborator.usefulFunctionA();
        return "The result of the B1 collaborating with the (" + result + ")";
    }
}

class ConcreteProductB2 implements AbstractProductB {
    @Override
    public String usefulFunctionB() {
        return "The result of the product B2.";
    }

    @Override
    public String anotherUsefulFunctionB(AbstractProductA collaborator) {
        String result = collaborator.usefulFunctionA();
        return "The result of the B2 collaborating with the (" + result + ")";
    }
}

interface AbstractFactory {
    AbstractProductA createProductA();

    AbstractProductB createProductB();
}

class ConcreteFactory1 implements AbstractFactory {
    @Override
    public AbstractProductA createProductA() {
        return new ConcreteProductA1();
    }

    @Override
    public AbstractProductB createProductB() {
        return new ConcreteProductB1();
    }
}

class ConcreteFactory2 implements AbstractFactory {
    @Override
    public AbstractProductA createProductA() {
        return new ConcreteProductA2();
    }

    @Override
    public AbstractProductB createProductB() {
        return new ConcreteProductB2();
    }
}

public class AbstractFactoryDemo {
    public static void clientCode(AbstractFactory factory) {
        AbstractProductA productA = factory.createProductA();
        AbstractProductB productB = factory.createProductB();

        System.out.println(productB.usefulFunctionB());
        System.out.println(productB.anotherUsefulFunctionB(productA));
    }

    public static void main(String[] args) {
        System.out.println("Client: Testing client code with the first factory type:");
        clientCode(new ConcreteFactory1());

        System.out.println("\n");

        System.out.println("Client: Testing the same client code with the second factory type:");
        clientCode(new ConcreteFactory2());
    }
}
