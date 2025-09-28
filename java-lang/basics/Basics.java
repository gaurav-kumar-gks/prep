package basics;



// ============================================================================
// 1. JAVA BASICS & HELLO WORLD
// ============================================================================

/**
 * JAVA PROGRAM STRUCTURE:
 * - Every Java program must have at least one class
 * - The class name must match the filename
 * - The main method is the entry point: public static void main(String[] args)
 * 
 * HELLO WORLD EXAMPLE:
 * public class HelloWorld {
 *     public static void main(String[] args) {
 *         System.out.println("Hello, World!");
 *     }
 * }
 * 
 * COMMAND LINE ARGUMENTS:
 * - Passed as String array to main method
 * - Access via args[0], args[1], etc.
 * - Check length with args.length
 * 
 */

/*

Basic Inputs:

Scanner sc = new Scanner(System.in)
Scanner fs = new Scanner(new File("file.txt"));
while (fs.hasNextLine()) {
    fs.nextLine()
}
sc.nextLine(); 
sc.nextInt();
sc.close();
fs.close();


*/

// ============================================================================
// 2. PRIMITIVE TYPES & RANGES
// ============================================================================

/**
 * PRIMITIVE TYPES OVERVIEW:
 * - Stored directly in memory (stack)
 * - Have fixed size
 * - Default values when not initialized
 * - Passed by value
 * - No methods (unlike objects)
 * 
 * SIZE AND RANGES:
 * - byte: 8 bits (-128 to 127)
 * - short: 16 bits (-32,768 to 32,767)
 * - int: 32 bits (-2^31 to 2^31-1) = -2,147,483,648 to 2,147,483,647
 * - long: 64 bits (-2^63 to 2^63-1) = -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
 * - float: 32 bits (IEEE 754, ~6-7 decimal digits precision)
 * - double: 64 bits (IEEE 754, ~15-16 decimal digits precision)
 * - char: 16 bits (Unicode, 0 to 65,535)
 * - boolean: 1 bit (true/false)
 * 
 * LITERALS:
 * - Integer literals: 42, 0x2A (hex), 052 (octal), 0b101010 (binary)
 * - Long literals: 42L, 42l
 * - Float literals: 3.14f, 3.14F
 * - Double literals: 3.14, 3.14d, 3.14D
 * - Character literals: 'A', '\n', '\t', '\u0041'
 * - String literals: "Hello", "Multi\nline"
 * - Boolean literals: true, false
 *
 * DEFAULT VALUES FOR PRIMITIVE TYPES:
 * - byte, short, int, long: 0
 * - float, double: 0.0
 * - char: '\u0000' (null character)
 * - boolean: false
 * 
 * NOTE: Local variables must be initialized before use!
 * Instance and static variables get default values automatically.
 */

// ============================================================================
// 3. REFERENCE TYPES & OBJECTS
// ============================================================================

/**
 * REFERENCE TYPES OVERVIEW:
 * - Stored as references (pointers to heap memory)
 * - Can be null
 * - Have methods and properties
 * - Passed by reference (reference is passed by value)
 * - Created using 'new' keyword (except for literals)
 * 
 * COMMON REFERENCE TYPES:
 * - String (immutable)
 * - Arrays
 * - Classes
 * - Interfaces
 * - Enums
 * - Wrapper classes (Integer, Double, etc.)
 * 
 * WRAPPER CLASSES:
 * - Byte, Short, Integer, Long, Float, Double, Character, Boolean
 * - Provide object representation of primitives
 * - Enable autoboxing/unboxing (Java 5+)
 * - Have useful utility methods
 */


// ============================================================================
// 4. VARIABLE DECLARATION & SCOPE
// ============================================================================

/**
 * VARIABLE DECLARATION RULES:
 * - Must start with letter, underscore, or dollar sign
 * - Can contain letters, digits, underscores, dollar signs
 * - Case sensitive
 * - Cannot use reserved keywords
 * - Should follow camelCase convention
 * 
 * SCOPE RULES:
 * - Local variables: declared inside method/block
 * - Instance variables: declared inside class, outside methods // e.g. public static String classVar;
 * - Class variables: declared with 'static' keyword // e.g. private String instanceVar;
 * - Parameter variables: method parameters
 * - Shadowing: when a local variable has the same name as a class variable
 * 
 * VARIABLE TYPES:
 * - Local variables: must be initialized before use
 * - Instance variables: get default values 
 * - Static variables: get default values, shared across all instances 
 */

// ============================================================================
// 5. TYPE CONVERSION & CASTING
// ============================================================================

/**
 * TYPE CONVERSION RULES:
 * - Widening conversion: automatic (safe)
 * - Narrowing conversion: requires explicit casting (may lose data)
 * - Reference types: requires explicit casting
 * 
 * WIDENING CONVERSION ORDER:
 * byte → short → int → long → float → double
 * char → int → long → float → double
 * 
 * CASTING:
 * - Explicit casting: (type) expression
 * - Implicit casting: automatic for widening conversions
 * - Check with instanceof before casting reference types
 * 
 * Integer to String:
 * String.valueOf(number)
 * Integer.toString(number)
 * "" + number
 * 
 * String to Integer:
 * Integer.parseInt(string)
 * Double.parseDouble(string)
 * 
 * Character to String:
 * String.valueOf(char)
 * "" + char
 * 
 * String to Character:
 * string.charAt(0)
 * 
 * Character to ascii:
 * (int) char
 * 
 * Ascii to Character:
 * (char) ascii
 */


// ============================================================================
// 6. CONSTANTS & FINAL VARIABLES
// ============================================================================

/**
 * FINAL KEYWORD:
 * - final variables: cannot be reassigned after initialization
 * - final methods: cannot be overridden
 * - final classes: cannot be inherited
 * 
 * CONSTANTS:
 * - Use static final for class constants // public static final String X ="hello"
 * - Use final for method constants // final String X = "Hello";
 * - Convention: UPPER_SNAKE_CASE for constants
 * - Final value must be initialized before use
 */


// ============================================================================
// 7. VAR KEYWORD (JAVA 10+)
// ============================================================================

/**
 * VAR KEYWORD:
 * - Local variable type inference (Java 10+)
 * - Type is inferred from the initializer
 * - Cannot be used for class fields, method parameters, or return types
 * - Cannot be initialized with null (type cannot be inferred)
 * - Cannot be used in lambda parameters
 */


// ============================================================================
// 8. SWITCH
// ============================================================================

/**
 * SUPPORTED TYPES:
 * - byte, short, int, char (traditional)
 * - String, enum (Java 7+)
 * - Any type with pattern matching (Java 17+)
 */

/*
    String season = "SUMMER";
    switch (season.toLowerCase()) {
        case "autumn":
            // no break - fall through
        case "fall":
            System.out.println("Autumn/Fall season");
            break;
        default:
            System.out.println("Unknown season");
    }
*/

/**
 * SWITCH EXPRESSIONS:
 * - Return values instead of just executing statements
 * - Arrow syntax (->) for cleaner code
 * - No fall-through behavior
 * - Must be exhaustive (cover all cases)
 * - Can be used as expressions
 */
/* 
    int month = 6;
    String monthName = switch (month) {
        case 1 -> "January";
        case 2 -> "February";
        default -> {
            System.out.println("Invalid month");
            yield "Invalid month";
        }
    };
*/

/**
 * PATTERN MATCHING OVERVIEW:
 * - instanceof with pattern variables
 * - Switch expressions with patterns
 * - Type-based switching
 * - Null handling in patterns
 */

// Pattern matching with instanceof
// Object obj = "Hello World";
// if (obj instanceof String str) {
//     System.out.println("String length: " + str.length());
// }

// Pattern matching with switch expressions (Java 21+)
// String result = switch (obj) {
//     case String s when s.startsWith("Hello") -> "HelloString: " + s;
//     case Integer i -> "Integer: " + i;
//     case Double d -> "Double: " + d;
//     default -> "Unknown type";
// };

// ============================================================================
// 9. ENUMS
// ============================================================================

/*

    public enum Planet {
        MERCURY("mercury") {
            public Integer getPlanetNumber() { return 1; }
        },
        NEPTUNE("neptune") {
            public Integer getPlanetNumber() { return 8; }
        };
        
        private final String name;

        Planet(String name) {
            this.name = name;
        }
        
        public String getName() { return name; }

        public abstract Integer getPlanetNumber();
        
        public static Planet fromName(String name) {
            return Arrays.stream(Planet.values())
                .filter(planet -> planet.getName().equals(name))
                .findFirst()
                .orElseThrow(() -> new IllegalArgumentException("Invalid planet name: " + name));
        }
    }

*/


// ============================================================================
// 10. MEMORY MANAGEMENT
// ============================================================================

/**
 * MEMORY MANAGEMENT IN JAVA:
 * - Stack: primitive types, method calls, local variables
 * - Heap: objects, arrays, reference types
 * - Garbage Collection: automatic memory management
 * - Memory leaks: possible with improper object references
 * 
 * MEMORY AREAS:
 * - Method Area: class metadata, static variables
 * - Heap: object instances
 * - Stack: method calls, local variables
 * - PC Register: current instruction
 * - Native Method Stack: native method calls
 */




 /**
 * COMPILATION & EXECUTION:
 * 
 * 1. BASIC COMPILATION:
 *    javac HelloWorld.java
 *    java HelloWorld arg1 arg2
 * 
 * 2. WITH PACKAGES:
 *    javac -d . com/example/HelloWorld.java
 *    java com.example.HelloWorld
 * 
 * 3. WITH CLASSPATH:
 *    javac -cp lib/*.jar MyClass.java
 *    java -cp .:lib/* MyClass
 */ 
