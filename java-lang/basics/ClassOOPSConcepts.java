package basics;



public class ClassOOPSConcepts {
    
    // ============================================================================
    // METHODS
    // ============================================================================
        
    /**
     * PARAMETER TYPES:
     * - Primitive types: passed by value
     * - Reference types: reference passed by value
     */
    
    /**
     * VARARGS OVERVIEW:
     * - Variable number of arguments
     * - Syntax: Type... parameterName
     * - Must be the last parameter
     * - Treated as an array inside the method
     * - Can be combined with regular parameters
     */
    
    /**
     * METHOD REFERENCES OVERVIEW:
     * - Shorthand for lambda expressions
     * - Syntax: Class::method or object::method
     * - Four types: 
     * 1) static method - e.g. String::length
     * 2) instance method - e.g. String::concat
     * 3) constructor method - e.g. Person::new
     * 4) arbitrary object method - e.g. String::toUpperCase
     *  
     * Used with functional interfaces
     * e.g. java.util.function.Function<String, Integer> lengthFunction = String::length;
     * 
     * Used with streams
     * e.g. numbers.stream().map(Math::abs).forEach(System.out::print);
     */

    /**
     * STATIC VS INSTANCE METHODS:
     * - Static methods: belong to class, can be called without object
     * - Instance methods: belong to object, require object instance
     * - Static methods cannot access instance variables/methods
     * - Instance methods can access static and instance members
     */


    // ============================================================================
    // POLYMORPHISM
    // ============================================================================
    
    /**
     * POLYMORPHISM OVERVIEW:
     * - Runtime Polymorphism: Method overriding
     * - Compile-time Polymorphism: Method overloading
     * - Dynamic method dispatch
     * - Virtual method invocation
     */

    /**
     * METHOD OVERLOADING:
     * - Multiple methods with same name but different parameters
     * - Different number / types / order of parameters
     * - Return type alone is not sufficient for overloading
     * - Compile-time polymorphism
     */

     /**
     * METHOD OVERRIDING:
     * - Method overriding is a runtime polymorphism
     * - Method overriding is a dynamic method dispatch
     * - Method overriding is a virtual method invocation
     */
    

    // ============================================================================
    // INHERITANCE AND COMPOSITION
    // ============================================================================
    
    /**
     * INHERITANCE OVERVIEW:
     * - IS-A relationship
     * - class ChildClass extends BaseClass
     * - Code reusability
     * - @Override
     * - Super keyword
     * - Constructor chaining
     */
    

    /**
     * COMPOSITION OVERVIEW:
     * - HAS-A relationship
     * - class A has class B as instance variable
     * - Object contains other objects
     * - More flexible than inheritance
     * - Better for code reuse
     */
    
    // ============================================================================
    // ENCAPSULATION AND DATA HIDING
    // ============================================================================
    
    /**
     * ENCAPSULATION OVERVIEW:
     * - Data hiding using access modifiers
     * - Controlled access through getters/setters
     * - Validation in setters
     * - Immutable objects
     */
    
    
    // ============================================================================
    // ABSTRACTION (ABSTRACT CLASSES AND INTERFACES)
    // ============================================================================
    
    /**
     * ABSTRACT CLASS:
     * - Abstract classes: Partial implementation
     * - Declaring Abstract class: public abstract class Shape
     * - Declaring Abstract method: public abstract double calculateArea();
     * - Abstract class can have static and final methods / variables
     * - Interfaces: Contract definition
     * - Multiple interface implementation
     * - Default methods in interfaces
     * - Static methods in interfaces
     */

    /**
     * ABSTRACT CLASS:
     * 
     * 1. Can abstract class have constructor? YES
     * 2. Can abstract class have static / final / private methods? YES
     * 3. Can abstract class have instance / static / final variables? YES
     * 4. Can abstract class implement interfaces / extend another abstract class / concrete class? YES
     * 5. Can abstract class have main method? YES
     * 6. Can abstract class be instantiated? NO (but can have anonymous subclasses)
     * 7. Can abstract class have all concrete / all abstract methods? YES (but must be declared abstract)
     * 8. Can abstract method be static / final / private / synchronized / native ? NO
     */
    
    /**
    * INTERFACE:
    * - Interface: Abstract contract
    * - Declaring Interface: public interface Drawable
    * - Declaring Interface method: public void draw();
    * - Interface can have default methods
    * - Interface can have static methods
    * - Interface can have instance variables
    */

    /**
     * INTERFACE:
     * 
     * - All methods are public and abstract by default (Java 8+ can have default/static)
     * - All variables are public, static, and final by default
     * - Cannot have constructors
     * - Cannot have instance variables
     * - Can extend multiple interfaces
     * - Cannot have static blocks
     * - Cannot have instance initialization blocks
     * - Can have default methods (Java 8+)
     * - Can have static methods (Java 8+)
     * - Can have private methods (Java 9+)
     * - Can have private static methods (Java 9+)
     */

    
    /* 
        // Interface example
        public interface Drawable {
            void draw();
            
            // Default method (Java 8+)
            default void drawWithBorder() {
                System.out.println("Drawing border...");
                draw();
                System.out.println("Border drawn.");
            }
            
            // Static method (Java 8+)
            static void drawMultiple(Drawable... drawables) {
                for (Drawable drawable : drawables) {
                    drawable.draw();
                }
            }
        }
    */
    
    // ============================================================================
    // INNER CLASSES
    // ============================================================================
    
    /**
     * INNER CLASS TYPES:
     * 1. Member Inner Class - non-static nested class
     * 2. Static Nested Class - static nested class
     * 3. Local Inner Class - class defined inside method
     * 4. Anonymous Inner Class - unnamed class
     * 5. Lambda Expressions - functional interface implementation
     */
    
    // ============================================================================
    // ASSOCIATION, AGGREGATION, AND COMPOSITION
    // ============================================================================
    
    /**
     * RELATIONSHIP TYPES:
     * - Association: General relationship between objects 
     * - Aggregation: HAS-A relationship (weak ownership) one object does not own the other
     * - Composition: HAS-A relationship (strong ownership) one object owns the other
     */
    
    
    // ============================================================================
    // DEPENDENCY AND COUPLING
    // ============================================================================
    
    /**
     * DEPENDENCY INJECTION:
     * - Constructor injection
     * - Setter injection
     * - Interface-based dependency
     * - Loose coupling
     */

    
    // ============================================================================
    // CONSTRUCTORS AND OBJECT INITIALIZATION
    // ============================================================================
    
    /**
     * CONSTRUCTOR TYPES:
     * - Default constructor
     * - Parameterized constructor
     * - Copy constructor
     * - Constructor chaining
     * - Static initialization blocks
     */
    
    public static class Person {
        private String name;
        private int age;
        private String email;
        private static int personCount = 0;
        
        // Static initialization block
        static {
            System.out.println("Person class is being loaded");
        }
        
        // Instance initialization block
        {
            personCount++;
            System.out.println("Person instance #" + personCount + " is being created");
        }
        
        // Default constructor
        public Person() {
            this("Unknown", 0, "unknown@email.com");
        }
        
        // Parameterized constructor
        public Person(String name, int age, String email) {
            this.name = name;
            this.age = age;
            this.email = email;
        }
        
        // Copy constructor
        public Person(Person other) {
            this(other.name, other.age, other.email);
        }
        
        // Constructor chaining
        public Person(String name) {
            this(name, 0, "unknown@email.com");
        }
        
        public Person(String name, int age) {
            this(name, age, "unknown@email.com");
        }
        
        public static int getPersonCount() {
            return personCount;
        }
    }
    
    // ============================================================================
    // STATIC VS INSTANCE MEMBERS
    // ============================================================================
    
    /**
     * STATIC VS INSTANCE:
     * - Static: Belongs to class, shared among all instances
     * - Instance: Belongs to object, unique per instance
     * - Static methods cannot access instance members
     * - Instance methods can access static members
     */
    
    public static class Counter {
        private static int totalCount = 0; // Static variable
        private int instanceCount = 0;     // Instance variable
        
        public Counter() {
            totalCount++;
            instanceCount++;
        }
        
        // Static method
        public static int getTotalCount() {
            return totalCount;
        }
        
        // Instance method
        public int getInstanceCount() {
            return instanceCount;
        }
        
        // Static method accessing static variable
        public static void resetTotalCount() {
            totalCount = 0;
        }
        
        // Instance method accessing both static and instance variables
        public void displayCounts() {
            System.out.println("Instance count: " + instanceCount);
            System.out.println("Total count: " + totalCount);
        }
    }

    /**
     * STATIC KEYWORD:
     * 
     * 1. Can static method / block access instance variables / method ? NO
     * 2. Can static method / block access static methods / variables ? Yes
     * 3. Can instance method / block access static variables / methods ? YES
     * 4. Can static method be overridden? NO (but can be hidden)
     * 5. Can static method be abstract? NO
     * 6. Can static method be final? YES
     * 7. Can static method be synchronized / native / strictfp ? YES
     * 8. Can static variable be final / volatile ? YES
     * 9. Can static variable be transient? NO (transient is for serialization)
     */

    /**
     * FINAL KEYWORD TRICKY QUESTIONS:
     * 
     * 1. Can final class be extended / have abstract methods? NO
     * 2. Can final class have non-final methods? YES
     * 3. Can final method be overridden / abstract? NO
     * 4. Can final method be static / private? YES
     * 5. Can final variable be reassigned? NO
     * 6. Can final variable be initialized later? YES (blank final)
     * 7. Can final variable be volatile / static / transient ? YES
     * 8. Can final method be synchronized / native / strictfp ? YES
     * 9. Can final class implement interfaces? YES
     */
    
    // ============================================================================
    // ACCESS MODIFIERS AND VISIBILITY
    // ============================================================================
    
    /**
     * ACCESS MODIFIERS:
     * - public: Accessible from anywhere
     * - protected: Accessible within package and subclasses
     * - package-private (default): Accessible within package only
     * - private: Accessible within class only
     */

    
    // ============================================================================
    // MODERN JAVA OOP FEATURES
    // ============================================================================
    
    /**
     * MODERN JAVA FEATURES:
     * - Records (Java 14+): Immutable data carriers
     * - Sealed Classes (Java 17+): Restricted inheritance
     * - Pattern Matching (Java 17+): Type-based switching
     * - Text Blocks (Java 15+): Multi-line strings
     */
    
    /* 
        public record Point(int x, int y) {
            // Compact constructor with validation
            public Point {
                if (x < 0 || y < 0) {
                    throw new IllegalArgumentException("Coordinates must be non-negative");
                }
            }
            
            // Can have methods
        }

        // Sealed Classes (Java 17+)
        public sealed class Shape permits Circle, Rectangle, Triangle { }
    */

    // ============================================================================
    // DESIGN PRINCIPLES
    // ============================================================================
    
    /**
     * SOLID PRINCIPLES:
     * - S: Single Responsibility Principle
     * - O: Open/Closed Principle: Open for extension, closed for modification
     * - L: Liskov Substitution Principle: Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.
     * - I: Interface Segregation Principle: Clients should not be forced to depend on interfaces they do not use.
     * - D: Dependency Inversion Principle: High-level modules should not depend on low-level modules. Both should depend on abstractions.
     */

}