package basics;

import java.util.*;

// ---------------------------------------------------------------
// GENERIC CLASSES
// ---------------------------------------------------------------
/*
 * - Generic classes allow you to define a class that can work with different data types.
 * - Syntax: class ClassName<T> where T is a type parameter.
 * - Type parameters are placeholders for actual types that will be specified when creating instances.
 * - Generics provide type safety at compile time, preventing ClassCastException at runtime.
 * - Use descriptive type parameter names (T for type, E for element, K for key, V for value, N for number).
 * - Prefer generic classes over raw types for type safety.
 * - Use wildcards when you need flexibility in type parameters.
 * - Cannot use primitive types as type parameters or create arrays of generic types
 * - Type erasure means runtime type information is lost.


 * - Generic methods allow you to define methods that can work with different types.
 * - Syntax: <T> returnType methodName(T parameter).
 * - Use generic methods when the method logic is independent of the class's type parameters.
 * 
 * - Generic methods can be static or instance methods.
 * - Static methods in generic classes must declare their own type parameters.
 * - Generic methods cannot use the class's type parameters if they're static.
 * - Type inference works best when the method parameters are used consistently.
*/


class Pair<K, V> {
    private K key;
    private V value;
    
    public Pair(K key, V value) {
        this.key = key;
        this.value = value;
    }
    
    public K getKey() { return key; }
    public V getValue() { return value; }

    public static <K,V> Pair<K,V> createPair(K key, V value) {
        return new Pair<>(key, value);
    }
    
    @Override
    public String toString() {
        return "(" + key + ", " + value + ")";
    }
}



// ---------------------------------------------------------------
// BOUNDED TYPES
// ---------------------------------------------------------------
/*
 * - Bounded types restrict the types that can be used as type parameters.
 * - Upper bounds: T extends SomeClass (T must be SomeClass or a subclass).
 * - Multiple bounds: T extends Class1 & Interface1 & Interface2.
 * - Lower bounds use super: <? super T>.

 * - Use bounded types to ensure type parameters have required methods/properties.
 * - Use interfaces for bounds when possible (more flexible than classes).
 * - Use multiple bounds sparingly to avoid complexity.
 * - Multiple bounds must list class first, then interfaces.
 * - Wildcard bounds have different syntax and behavior.
 * - Bounded wildcards provide more flexibility than bounded type parameters.
 */


// ---------------------------------------------------------------
// TYPE ERASURE
// ---------------------------------------------------------------
/*
 * - Type erasure is the process by which generic type information is removed at compile time.
 * - All generic types become their raw types (or Object if unbounded) at runtime.
 * - This ensures backward compatibility with pre-generics Java code.
 * - Type information is lost at runtime, but compile-time type checking is preserved.
 * - Be aware that runtime type checking of generic types is not possible.
 * - Use instanceof with raw types, not generic types.
 * - Use reflection carefully with generics.

 * - Cannot create arrays of generic types: new List<String>[10].
 * - Runtime type information is lost.
 * - Type erasure can be bypassed using reflection, but it's complex and not recommended.
 * - The erasure process is why you can't overload methods based on generic types.
*/

class TypeErasureDemo {
    public static void main(String[] args) {
        List<String> stringList = new ArrayList<>();
        List<Integer> intList = new ArrayList<>();
        
        // At runtime, both are just List (raw type)
        System.out.println(stringList.getClass() == intList.getClass()); // true
        
        // Cannot use instanceof with generic types
        // if (stringList instanceof List<String>) // Compile-time error!
        if (stringList instanceof List) { // Must use raw type
            System.out.println("It's a List");
        }
        
        // Cannot create arrays of generic types
        // List<String>[] array = new List<String>[10]; // Compile-time error!
        List<String>[] array = (List<String>[]) new List[10]; // Unchecked cast warning
    }
}

// ---------------------------------------------------------------
// GENERIC CONSTRAINTS
// ---------------------------------------------------------------
/*
 * - Generic constraints limit what types can be used with generics.
 * - Common constraints: extends (upper bound), super (lower bound), wildcards (?).
 * - Wildcards provide flexibility: ? (unbounded), ? extends T (upper bounded), ? super T (lower bounded).
 * - Use wildcards when you need flexibility in method parameters.
 * - Use extends for reading operations, super for writing operations.
 * - Follow PECS principle: Producer Extends, Consumer Super.
 * - Cannot add elements to a collection with ? extends T (except null).
 * - Cannot read elements from a collection with ? super T (except as Object).
 * - Wildcards are used in method parameters, not in class declarations.
 * - The PECS principle helps remember when to use extends vs super.
*/

class GenericConstraintsDemo {
    // Producer method (reads from collection)
    public static double sum(List<? extends Number> numbers) {
        double sum = 0.0;
        for (Number num : numbers) {
            sum += num.doubleValue();
        }
        return sum;
    }
    
    // Consumer method (writes to collection)
    public static void addNumbers(List<? super Integer> numbers) {
        numbers.add(1);
        numbers.add(2);
        numbers.add(3);
    }
    
    // Unbounded wildcard
    public static void printList(List<?> list) {
        for (Object item : list) {
            System.out.print(item + " ");
        }
        System.out.println();
    }
    
    public static void main(String[] args) {
        List<Integer> intList = Arrays.asList(1, 2, 3, 4, 5);
        List<Double> doubleList = Arrays.asList(1.1, 2.2, 3.3);
        
        System.out.println("Sum of integers: " + sum(intList));
        System.out.println("Sum of doubles: " + sum(doubleList));
        
        List<Number> numberList = new ArrayList<>();
        addNumbers(numberList);
        System.out.println("Added numbers: " + numberList);
        
        printList(intList);
        printList(doubleList);
        
        // Demonstrating constraints
        List<? extends Number> producer = intList;
        // producer.add(10); // Compile-time error!
        Number num = producer.get(0); // OK
        
        List<? super Integer> consumer = new ArrayList<Number>();
        consumer.add(10); // OK
        // Integer i = consumer.get(0); // Compile-time error!
        Object obj = consumer.get(0); // OK
    }
}

// ---------------------------------------------------------------
// GENERIC COLLECTIONS
// ---------------------------------------------------------------
/*
 * - Java collections framework is built around generics.
 * - All collection interfaces and classes are generic: List<T>, Set<T>, Map<K,V>.
 * - Generic collections provide type safety and eliminate casting.
 * - Collections can be nested: List<List<String>>, Map<String, List<Integer>>.
 * - Always use generic collections instead of raw types.
 * - Use diamond operator (<>) for type inference when possible.
 * - Be careful with nested generics for readability.

 * Pitfalls:
 * - Raw types bypass type checking and can cause runtime errors.
 * - Nested generics can become complex and hard to read.
 * - Type erasure affects runtime behavior.

 * Advanced Notes:
 * - Collections can be made type-safe using Collections.checkedList/Set/Map.
 * - Custom collections can be built using generic interfaces.

*/

// ---------------------------------------------------------------
// SUMMARY TABLE: Generics Quick Reference
// ---------------------------------------------------------------
/*
| Concept           | Syntax                    | Use Case                    | Example                    |
|-------------------|---------------------------|----------------------------|----------------------------|
| Generic Class     | class Box<T>              | Type-safe containers        | Box<String> box            |
| Generic Method    | <T> void method(T param)  | Type-independent logic      | <T> T max(T a, T b)        |
| Bounded Type      | <T extends Number>        | Constrain type parameters   | <T extends Comparable<T>>  |
| Wildcard          | List<?>                   | Flexible method parameters  | List<? extends Number>     |
| Type Erasure      | Runtime type loss         | Backward compatibility      | List<String> → List        |
| Generic Collection| List<T>, Map<K,V>         | Type-safe data structures   | List<String> list          |
*/ 