/**
 * Decorator is a structural design pattern that lets you attach new behaviors
 * to objects
 * by placing these objects inside special wrapper objects that contain the
 * behaviors.
 */

// Component interface
interface Coffee {
    double getCost();

    String getDescription();
}

// Concrete component
class SimpleCoffee implements Coffee {
    @Override
    public double getCost() {
        return 1.0;
    }

    @Override
    public String getDescription() {
        return "Simple coffee";
    }
}

// Abstract decorator
abstract class CoffeeDecorator implements Coffee {
    protected Coffee decoratedCoffee;

    public CoffeeDecorator(Coffee coffee) {
        this.decoratedCoffee = coffee;
    }

    @Override
    public double getCost() {
        return decoratedCoffee.getCost();
    }

    @Override
    public String getDescription() {
        return decoratedCoffee.getDescription();
    }
}

// Concrete decorators
class MilkDecorator extends CoffeeDecorator {
    public MilkDecorator(Coffee coffee) {
        super(coffee);
    }

    @Override
    public double getCost() {
        return super.getCost() + 0.5;
    }

    @Override
    public String getDescription() {
        return super.getDescription() + ", with milk";
    }
}

class SugarDecorator extends CoffeeDecorator {
    public SugarDecorator(Coffee coffee) {
        super(coffee);
    }

    @Override
    public double getCost() {
        return super.getCost() + 0.2;
    }

    @Override
    public String getDescription() {
        return super.getDescription() + ", with sugar";
    }
}

class WhippedCreamDecorator extends CoffeeDecorator {
    public WhippedCreamDecorator(Coffee coffee) {
        super(coffee);
    }

    @Override
    public double getCost() {
        return super.getCost() + 0.7;
    }

    @Override
    public String getDescription() {
        return super.getDescription() + ", with whipped cream";
    }
}

public class DecoratorDemo {
    public static void main(String[] args) {
        // Create a simple coffee
        Coffee coffee = new SimpleCoffee();
        System.out.println("Cost: $" + coffee.getCost() + "; Description: " + coffee.getDescription());

        // Add milk
        coffee = new MilkDecorator(coffee);
        System.out.println("Cost: $" + coffee.getCost() + "; Description: " + coffee.getDescription());

        // Add sugar
        coffee = new SugarDecorator(coffee);
        System.out.println("Cost: $" + coffee.getCost() + "; Description: " + coffee.getDescription());

        // Add whipped cream
        coffee = new WhippedCreamDecorator(coffee);
        System.out.println("Cost: $" + coffee.getCost() + "; Description: " + coffee.getDescription());
    }
}