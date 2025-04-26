/**
 * Singleton Pattern in Java
 * 
 * Intent:
 * Singleton is a creational design pattern that lets you ensure that a class
 * has only one instance,
 * while providing a global access point to this instance.
 */

// Naive Singleton (not thread-safe)
class SingletonNaive {
    private static SingletonNaive instance;
    private String value;

    private SingletonNaive(String value) {
        this.value = value;
    }

    public static SingletonNaive getInstance(String value) {
        if (instance == null) {
            instance = new SingletonNaive(value);
        }
        return instance;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }
}

// Thread-safe Singleton with double-checked locking
class SingletonThreadSafe {
    private static volatile SingletonThreadSafe instance;
    private String value;

    private SingletonThreadSafe(String value) {
        this.value = value;
    }

    public static SingletonThreadSafe getInstance(String value) {
        if (instance == null) {
            synchronized (SingletonThreadSafe.class) {
                if (instance == null) {
                    instance = new SingletonThreadSafe(value);
                }
            }
        }
        return instance;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }
}

// Enum Singleton (thread-safe by default)
enum SingletonEnum {
    INSTANCE("Initial value");

    private String value;

    SingletonEnum(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }
}

// Client code
public class SingletonDemo {
    public static void main(String[] args) {
        // Test naive singleton
        System.out.println("Testing Naive Singleton:");
        SingletonNaive naive1 = SingletonNaive.getInstance("FOO");
        SingletonNaive naive2 = SingletonNaive.getInstance("BAR");
        System.out.println("Value 1: " + naive1.getValue());
        System.out.println("Value 2: " + naive2.getValue());
        System.out.println("Same instance: " + (naive1 == naive2));

        System.out.println("\n");

        // Test thread-safe singleton
        System.out.println("Testing Thread-Safe Singleton:");
        SingletonThreadSafe safe1 = SingletonThreadSafe.getInstance("FOO");
        SingletonThreadSafe safe2 = SingletonThreadSafe.getInstance("BAR");
        System.out.println("Value 1: " + safe1.getValue());
        System.out.println("Value 2: " + safe2.getValue());
        System.out.println("Same instance: " + (safe1 == safe2));

        System.out.println("\n");

        // Test enum singleton
        System.out.println("Testing Enum Singleton:");
        SingletonEnum enum1 = SingletonEnum.INSTANCE;
        enum1.setValue("FOO");
        SingletonEnum enum2 = SingletonEnum.INSTANCE;
        enum2.setValue("BAR");
        System.out.println("Value: " + enum1.getValue());
        System.out.println("Same instance: " + (enum1 == enum2));
    }
}