package basics;

import java.util.*;
import java.util.function.*;
import java.util.stream.*;

// ---------------------------------------------------------------
// LAMBDA EXPRESSIONS
// ---------------------------------------------------------------
/*
 * - Lambda expressions are anonymous functions, introduced in Java 8.
 * - Syntax: (parameters) -> expression or (parameters) -> { statements }
 * - Used to provide implementations for functional interfaces (interfaces with a single abstract method).
 * - Enable concise, functional-style code, especially for collections and event handling.
 * - Can only access final or effectively final variables from enclosing scope.
 * - Avoid side effects inside lambdas.
 *
 * Advanced Notes:
 * - Lambdas can capture variables from the enclosing context (closures).
 * - Lambdas are compiled to private static or synthetic methods (invokedynamic).
 */

class LambdaDemo {
    public static void main(String[] args) {
        // Single parameter, single statement
        Consumer<String> printer = s -> System.out.println(s);
        printer.accept("Hello Lambda!");

        // Multiple parameters, block body
        BiFunction<Integer, Integer, Integer> sum = (a, b) -> {
            int result = a + b;
            return result;
        };
        System.out.println("Sum: " + sum.apply(2, 3));

        // No parameter
        Runnable r = () -> System.out.println("No params");
        r.run();
    }
}

// ---------------------------------------------------------------
// BUILT-IN FUNCTIONAL INTERFACES
// ---------------------------------------------------------------
/*
 * - Functional Interfaces as nothing but a interface with a single abstract
 * method
 * - Java provides many functional interfaces in java.util.function:
 * - Predicate<T>: boolean test(T t)
 * - Function<T, R>: R apply(T t)
 * - Consumer<T>: void accept(T t)
 * - Supplier<T>: T get()
 * - UnaryOperator<T>: T apply(T t)
 * - BinaryOperator<T>: T apply(T t1, T t2)
 * - BiFunction<T, U, R>, BiConsumer<T, U>, BiPredicate<T, U>
 * 
 * - Checked exceptions are not supported directly.
 * - Compose functions: f.andThen(g), f.compose(g)
 * - Use primitive specializations - IntPredicate, IntConsumer, IntFunction,
 * IntSupplier etc
 */

class BuiltinFunctionalInterfacesDemo {
    public static void main(String[] args) {
        Predicate<String> isEmpty = String::isEmpty;
        // lambda is equivalent to anonymous class
        //
        // Predicate<String> isEmpty = s -> s.isEmpty();
        //
        // Predicate<String> isEmpty = new Predicate<String>() {
        // @Override
        // public boolean test(String s) {
        // return s.isEmpty();
        // }
        // };
        System.out.println(isEmpty.test(""));

        Supplier<Double> random = Math::random;
        System.out.println(random.get());

        // Basic composition
        Function<String, String> trim = String::trim;
        Function<String, String> upper = String::toUpperCase;
        Function<String, String> addPrefix = s -> "PREFIX: " + s;
        Function<String, String> composed = trim.andThen(upper).andThen(addPrefix);
        System.out.println("Composed: " + composed.apply("  hello world  "));

        // Compose vs andThen
        Function<Integer, Integer> f = x -> x + 1;
        Function<Integer, Integer> g = x -> x * 2;
        Function<Integer, Integer> andThen = f.andThen(g); // f then g
        Function<Integer, Integer> compose = f.compose(g); // g then f
        System.out.println("AndThen: " + andThen.apply(2));
        System.out.println("Compose: " + compose.apply(2));

        // Currying with different types
        Function<String, Function<Integer, String>> repeat = s -> n -> s.repeat(n);
        System.out.println("Repeat: " + repeat.apply("Hello").apply(3));

        // Strategy selection
        Map<String, Function<Double, Double>> strategies = Map.of(
                "tax", amount -> amount * 1.1,
                "discount", amount -> amount * 0.9,
                "free", amount -> 0.0);
        Function<Double, Double> strategy = strategies.get("tax");
        System.out.println("Strategy: " + strategy.apply(100.0));

        // Observer as functions
        List<Consumer<String>> observers = Arrays.asList(
                System.out::println,
                s -> System.out.println("LOG: " + s),
                s -> System.out.println("ALERT: " + s.toUpperCase()));
        String event = "Something happened";
        observers.forEach(observer -> observer.accept(event));

        // Commands as functions
        List<Runnable> commands = Arrays.asList(
                () -> System.out.println("Command 1 executed"),
                () -> System.out.println("Command 2 executed"),
                () -> System.out.println("Command 3 executed"));
        commands.forEach(Runnable::run);

        // Functional builder
        Function<String, Function<Integer, Person>> personBuilder = name -> age -> new Person(name, age);
        Person person = personBuilder.apply("John").apply(30);

        // Chain as function composition
        Function<String, String> validate = s -> s == null ? "Invalid" : s;
        Function<String, String> sanitize = s -> s.trim().toLowerCase();
        Function<String, String> format = s -> s.substring(0, 1).toUpperCase() + s.substring(1);
        Function<String, String> process = validate
                .andThen(sanitize)
                .andThen(format);

    }

    static class Person {
        private String name;
        private Integer age;

        public Person(String name, Integer age) {
            this.name = name;
            this.age = age;
        }

        public String getName() {
            return name;
        }

        public Integer getAge() {
            return age;
        }
    }

}

// ---------------------------------------------------------------
// STREAM API
// ---------------------------------------------------------------
/*
 * - Streams provide a declarative, functional way to process collections.
 * - Operations:
 * creation (stream, parallelStream)
 * intermediate (map, filter, sorted, distinct, limit, skip)
 * terminal (forEach, collect, reduce, count, anyMatch, allMatch, noneMatch)
 * - Streams are lazy; operations are only performed on terminal operation.
 * - Streams do not store data; they operate on source collections.
 * - Use streams for complex data processing pipelines.
 * - Prefer method references and lambdas for clarity.
 * - Use parallel streams only when appropriate.
 * - Streams cannot be reused once consumed.
 * - Mutating state inside streams can lead to bugs.
 * - Custom collectors for advanced aggregation.
 */

// ---------------------------------------------------------------
// Collectors
// ---------------------------------------------------------------
/*
 * - Collectors are used to aggregate data from streams.
 * - Collectors are used to group data, sum data, count data etc.
 * - Collectors.toList(), Collectors.toSet(), Collectors.toMap(),
 * Collectors.toCollection() etc
 */

class StreamApiDemo {
    public static void main(String[] args) {
        List<Integer> nums = Arrays.asList(3, 2, 1, 0, 4);
        Map<Integer, List<String>> byLength = Arrays.asList("a", "bb", "ccc", "dd").stream()
                .collect(Collectors.groupingBy(String::length));
        System.out.println(byLength);
        int functionalSum = nums.stream()
                .filter(n -> n % 2 == 0) // Filter: keep only even numbers
                .map(n -> n * n) // Map: square each number
                .reduce(0, Integer::sum); // Reduce: sum all numbers
        System.out.println(functionalSum);
    }

}

// ============================================================================
// OPTIONAL
// ============================================================================

/*
 * Optional is a container object that may or may not contain a non-null value.
 * 
 * KEY CONCEPTS:
 * 
 * Optional.of() - creates an optional with a non-null value
 * Optional.ofNullable() - creates an optional with a nullable value
 * Optional.empty() - creates an empty optional
 * Optional.isPresent() - checks if the optional is not null
 * Optional.isEmpty() - checks if the optional is null
 * Optional.get() - gets the value of the optional, always checks if the
 * optional is present
 * Optional.orElse() - gets the value of the optional or a default value (always
 * evaluates)
 * Optional.orElseGet() - gets the value of the optional or a default value
 * (lazy evaluation)
 * Optional.ifPresentOrElse() - gets the value of the optional or a default
 * value
 * 
 * BEST PRACTICES:
 * - Use Optional for return values, not parameters or fields
 * - Prefer orElseGet() over orElse() for expensive computations
 * - Use map() and flatMap() for chaining operations
 * - Avoid Optional.get() without checking isPresent()
 * - Use Optional.empty() instead of null
 * - Don't use Optional in collections
 * - Don't use for performance-critical code (has overhead)
 * - Don't nest optionals or use it in collections
 */

// ============================================================================
// CUSTOM FUNCTIONAL INTERFACES
// ============================================================================

/*
 * CUSTOM FUNCTIONAL INTERFACES THEORY:
 * 
 * While Java provides many built-in functional interfaces, sometimes you need
 * custom ones for specific use cases. Custom functional interfaces should:
 * 
 * 1. Be annotated with @FunctionalInterface
 * 2. Have exactly one abstract method
 * 3. Can have default and static methods
 * 4. Should be clearly named to indicate their purpose
 * 5. Should follow naming conventions (verb-based names)
 * 
 * COMMON PATTERNS:
 * - Transformer: T -> R
 * - Validator: T -> boolean
 * - Processor: T -> void
 * - Generator: () -> T
 * - Comparator: (T, T) -> int
 * - Mapper: T -> Optional<R>
 * - Handler: T -> Optional<T>
 */
