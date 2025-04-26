
import java.util.ArrayList;
import java.util.List;

class Product1 {
    private List<String> parts = new ArrayList<>();

    public void add(String part) {
        parts.add(part);
    }

    public void listParts() {
        System.out.println("Product parts: " + String.join(", ", parts));
    }
}

interface Builder {
    void reset();

    void producePartA();

    void producePartB();

    void producePartC();

    Product1 getProduct();
}

class ConcreteBuilder1 implements Builder {
    private Product1 product;

    public ConcreteBuilder1() {
        this.reset();
    }

    @Override
    public void reset() {
        product = new Product1();
    }

    @Override
    public void producePartA() {
        product.add("PartA1");
    }

    @Override
    public void producePartB() {
        product.add("PartB1");
    }

    @Override
    public void producePartC() {
        product.add("PartC1");
    }

    @Override
    public Product1 getProduct() {
        Product1 result = this.product;
        this.reset();
        return result;
    }
}

class Director {
    private Builder builder;

    public void setBuilder(Builder builder) {
        this.builder = builder;
    }

    public void buildMinimalViableProduct() {
        builder.producePartA();
    }

    public void buildFullFeaturedProduct() {
        builder.producePartA();
        builder.producePartB();
        builder.producePartC();
    }
}

public class BuilderDemo {
    public static void main(String[] args) {
        Director director = new Director();
        ConcreteBuilder1 builder = new ConcreteBuilder1();
        director.setBuilder(builder);

        System.out.println("Standard basic product: ");
        director.buildMinimalViableProduct();
        builder.getProduct().listParts();

        System.out.println("\n");

        System.out.println("Standard full featured product: ");
        director.buildFullFeaturedProduct();
        builder.getProduct().listParts();

        System.out.println("\n");

        // Remember, the Builder pattern can be used without a Director class.
        System.out.println("Custom product: ");
        builder.producePartA();
        builder.producePartB();
        builder.getProduct().listParts();
    }
}