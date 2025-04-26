import java.util.HashMap;
import java.util.Map;

interface Prototype {
    Prototype clone();
}

class ConcretePrototype implements Prototype {
    private String field;

    public ConcretePrototype(String field) {
        this.field = field;
    }

    public String getField() {
        return field;
    }

    public void setField(String field) {
        this.field = field;
    }

    @Override
    public Prototype clone() {
        return new ConcretePrototype(this.field);
    }

    @Override
    public String toString() {
        return "ConcretePrototype{field='" + field + "'}";
    }
}

class PrototypeRegistry {
    private static Map<String, Prototype> items = new HashMap<>();

    public static void addItem(String id, Prototype item) {
        items.put(id, item);
    }

    public static Prototype getById(String id) {
        Prototype item = items.get(id);
        return item != null ? item.clone() : null;
    }

    public static Prototype getByColor(String color) {
        for (Prototype item : items.values()) {
            if (item instanceof ConcretePrototype) {
                ConcretePrototype concreteItem = (ConcretePrototype) item;
                if (concreteItem.getField().equals(color)) {
                    return concreteItem.clone();
                }
            }
        }
        return null;
    }
}

public class PrototypeDemo {
    public static void main(String[] args) {
        PrototypeRegistry.addItem("basic", new ConcretePrototype("blue"));
        PrototypeRegistry.addItem("standard", new ConcretePrototype("red"));
        PrototypeRegistry.addItem("premium", new ConcretePrototype("green"));

        Prototype basic = PrototypeRegistry.getById("basic");
        Prototype standard = PrototypeRegistry.getById("standard");
        Prototype premium = PrototypeRegistry.getById("premium");

        if (basic instanceof ConcretePrototype) {
            ((ConcretePrototype) basic).setField("light blue");
        }

        Prototype redPrototype = PrototypeRegistry.getByColor("red");

        System.out.println("Basic: " + basic);
        System.out.println("Standard: " + standard);
        System.out.println("Premium: " + premium);
        System.out.println("Red Prototype: " + redPrototype);
    }
}