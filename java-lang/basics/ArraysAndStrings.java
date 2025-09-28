package basics;

/*
STRINGS 

- String is final and immutable (cannot be changed after creation)
- String literals are interned in the string pool (single instance per value)
- String pool is a special area of memory where strings are stored
- String objects created with 'new' are not interned unless .intern() is called
- Immutability enables safe sharing, caching, and use as map keys
- Strings are UTF-16 encoded internally
- For Unicode characters, use char[] instead of String e.g with emoji: 
  char[] emoji = new char[] { '\uD83D', '\uDC31' };
  String emojiString = new String(emoji); // "🐱"

=== STRING CREATION & COMPARISON ===
String a = "hello";
String b = "hello";
String c = new String("hello");
System.out.println(a == b); // true (same pool)
System.out.println(a == c); // false (different object)
System.out.println(a.equals(c)); // true (same value)
System.out.println(a.equalsIgnoreCase("HELLO")); // true
System.out.println(a.compareTo("world")); // negative (lexicographic)
System.out.println(a.compareToIgnoreCase("HELLO")); // 0

=== STRING MANIPULATION ===
String s = "Java,Python,Go";
String[] langs = s.split(","); // ["Java", "Python", "Go"]
String joined = String.join("|", langs); // "Java|Python|Go"
String concat = s.concat("Haskell"); // "Java,Python,GoHaskell"
String replaced = s.replace("Java", "JavaScript"); // "JavaScript,Python,Go"
String replacedAll = s.replaceAll("[aeiou]", "*"); // "J*v*Pyth*n,G*"
String replacedFirst = s.replaceFirst("[aeiou]", "*"); // "J*va,Python,Go"

=== STRING SEARCHING ===
boolean contains = s.contains("Python"); // true
boolean startsWith = s.startsWith("Java"); // true
boolean endsWith = s.endsWith("Go"); // true
int indexOf = s.indexOf("Python"); // 5
int lastIndexOf = s.lastIndexOf("a"); // 3
int indexOfFrom = s.indexOf("a", 2); // 3 (search from index 2)
boolean isEmpty = s.isEmpty(); // false
boolean isBlank = "   ".isBlank(); // true (Java 11+)

=== STRING SUBSTRING & EXTRACTION ===
String substring = s.substring(0, 4); // "Java"
String substringFrom = s.substring(5); // "Python,Go"
char charAt = s.charAt(0); // 'J'
int codePointAt = s.codePointAt(0); // 74 (Unicode value)
int codePointBefore = s.codePointBefore(1); // 74

=== STRING FORMATTING ===
String formatted = String.format("%s-%d", "foo", 42); // "foo-42"
String valueOf = String.valueOf(123); // "123"
String valueOfNull = String.valueOf((Object) null); // "null"

=== STRING REGEX & PATTERN MATCHING ===
String regex = "[A-Za-z]+";
boolean matches = "HelloWorld".matches(regex); // true
String[] splitWithLimit = s.split(",", 2); // ["Java", "Python,Go"]
String[] splitRegex = "a1b2c3".split("\\d"); // ["a", "b", "c"]

=== STRING ENCODING & BYTES ===
byte[] utf8 = s.getBytes(java.nio.charset.StandardCharsets.UTF_8);
String fromBytes = new String(utf8, java.nio.charset.StandardCharsets.UTF_8);
char[] charArray = s.toCharArray(); // ['J','a','v','a',',','P','y','t','h','o','n',',','G','o']

=== STRING WHITESPACE HANDLING ===
String withSpaces = "  Java  ";
String trimmed = withSpaces.trim(); // "Java" (removes leading/trailing whitespace)
String stripped = withSpaces.strip(); // "Java" (Java 11+, handles Unicode whitespace)
String strippedLeading = withSpaces.stripLeading(); // "Java  " (Java 11+)
String strippedTrailing = withSpaces.stripTrailing(); // "  Java" (Java 11+)

=== STRING CASE CONVERSION ===
String toUpperCase = s.toUpperCase(); // "JAVA,PYTHON,GO"
String toLowerCase = s.toLowerCase(); // "java,python,go"

=== STRING LENGTH & VALIDATION ===
int length = s.length(); // 13
boolean isEmpty = s.isEmpty(); // false
boolean isBlank = "   ".isBlank(); // true (Java 11+)

=== STRING INTERNING ===
String interned = s.intern(); // Returns interned version
String newString = new String("hello").intern(); // Interns the string

=== STRING JOINER (Java 8+) ===
StringJoiner joiner = new StringJoiner(", ", "[", "]");
joiner.add("Java").add("Python").add("Go");
String result = joiner.toString(); // "[Java, Python, Go]"

=== STRING LINES (Java 11+) ===
String multiLine = "Line1\nLine2\nLine3";
List<String> lines = multiLine.lines().toList(); // ["Line1", "Line2", "Line3"]

=== STRING REPEAT (Java 11+) ===
String repeated = "Java".repeat(3); // "JavaJavaJava"

=== STRING INDENT (Java 12+) ===
String indented = "Java\nPython".indent(2); // "  Java\n  Python"

=== STRING TRANSFORM (Java 12+) ===
String transformed = s.transform(str -> str.toUpperCase()); // "JAVA,PYTHON,GO"

=== STRING DESCRIBE CONSTABLE (Java 12+) ===
String constant = "hello".describeConstable(); // Optional<String> containing "hello"

=== STRING RESOLVE CONSTANT DESC (Java 12+) ===
String resolved = "hello".resolveConstantDesc(null); // "hello"
*/

/*
ARRAYS

=== ARRAY FUNDAMENTALS ===
- Arrays are fixed-size, zero-based, and covariant (dangerous)
- Covariant: an array of a subclass can be assigned to an array of a superclass
- Arrays store references for objects, primitives for primitive types
- Arrays are not resizable (use ArrayList for dynamic arrays)
- Arrays are objects (arr instanceof Object is true)
- Multidimensional arrays: int[][] matrix = new int[2][3];
- Jagged arrays: int[][] jagged = new int[2][]; jagged[0] = new int[3];

=== ARRAY CREATION ===
int[] arr = new int[3]; // [0, 0, 0]
int[] arr2 = {1, 2, 3}; // [1, 2, 3]
int[] arr3 = new int[]{1, 2, 3}; // [1, 2, 3]
String[] strings = {"Java", "Python", "Go"}; // ["Java", "Python", "Go"]
int[][] matrix = new int[2][3]; // 2x3 matrix
int[][] jagged = new int[2][]; // jagged array
jagged[0] = new int[3]; // first row has 3 elements
jagged[1] = new int[5]; // second row has 5 elements

=== ARRAY ACCESS & MODIFICATION ===
int[] arr = {1, 2, 3};
int length = arr.length; // 3
int first = arr[0]; // 1
arr[1] = 42; // [1, 42, 3]
arr[arr.length - 1] = 99; // [1, 42, 99] (last element)

=== ARRAY COPYING ===
int[] copy = Arrays.copyOf(arr, arr.length); // [1, 42, 99]
int[] copy2 = Arrays.copyOfRange(arr, 1, 3); // [42, 99]
int[] copy3 = Arrays.copyOf(arr, 5); // [1, 42, 99, 0, 0] (padded with 0)
int[] copy4 = Arrays.copyOf(arr, 2); // [1, 42] (truncated)

=== ARRAY TO LIST CONVERSION ===
List<Integer> list = Arrays.asList(arr); // [1, 42, 99] (fixed-size list)
Integer[] boxed = {1, 2, 3};
List<Integer> list2 = Arrays.asList(boxed); // [1, 2, 3]
int[] backToArray = list.stream().mapToInt(Integer::intValue).toArray(); // [1, 42, 99]

=== ARRAY SORTING ===
int[] unsorted = {3, 1, 4, 1, 5};
Arrays.sort(unsorted); // [1, 1, 3, 4, 5] (ascending)
Arrays.sort(unsorted, 1, 4); // sort from index 1 to 3 (exclusive)
Arrays.parallelSort(unsorted); // parallel sorting for large arrays

=== ARRAY SEARCHING ===
int[] sorted = {1, 2, 3, 4, 5};
int index = Arrays.binarySearch(sorted, 3); // 2 (index of 3)
int notFound = Arrays.binarySearch(sorted, 6); // -6 (insertion point)
boolean contains = Arrays.stream(sorted).anyMatch(x -> x == 3); // true

=== ARRAY COMPARISON ===
int[] arr1 = {1, 2, 3};
int[] arr2 = {1, 2, 3};
boolean equals = Arrays.equals(arr1, arr2); // true
boolean deepEquals = Arrays.deepEquals(new int[][]{arr1}, new int[][]{arr2}); // true

=== ARRAY FILLING ===
int[] fillArr = new int[5];
Arrays.fill(fillArr, 42); // [42, 42, 42, 42, 42]
Arrays.fill(fillArr, 1, 4, 99); // [42, 99, 99, 99, 42] (fill from index 1 to 3)

=== ARRAY STREAMING (Java 8+) ===
int[] numbers = {1, 2, 3, 4, 5};
Arrays.stream(numbers).forEach(System.out::println); // 1 2 3 4 5
int sum = Arrays.stream(numbers).sum(); // 15
int max = Arrays.stream(numbers).max().orElse(0); // 5
int min = Arrays.stream(numbers).min().orElse(0); // 1
double average = Arrays.stream(numbers).average().orElse(0.0); // 3.0
int[] evens = Arrays.stream(numbers).filter(x -> x % 2 == 0).toArray(); // [2, 4]
int[] doubled = Arrays.stream(numbers).map(x -> x * 2).toArray(); // [2, 4, 6, 8, 10]

=== ARRAY HASH CODE ===
int hashCode = Arrays.hashCode(arr); // hash code for array
int deepHashCode = Arrays.deepHashCode(new int[][]{arr}); // deep hash code

=== ARRAY STRING REPRESENTATION ===
String arrayString = Arrays.toString(arr); // "[1, 42, 99]"
String deepString = Arrays.deepToString(new int[][]{arr}); // "[[1, 42, 99]]"

=== ARRAY UTILITIES ===
int[] setAll = new int[5];
Arrays.setAll(setAll, i -> i * i); // [0, 1, 4, 9, 16]
Arrays.parallelSetAll(setAll, i -> i * 2); // [0, 2, 4, 6, 8] (parallel)

=== ARRAY MISCELLANEOUS ===
int[] prefix = Arrays.copyOf(arr, 2); // [1, 42] (first 2 elements)
int[] suffix = Arrays.copyOfRange(arr, 1, arr.length); // [42, 99] (from index 1 to end)
int[] reverse = new int[arr.length];
for (int i = 0; i < arr.length; i++) {
    reverse[i] = arr[arr.length - 1 - i]; // [99, 42, 1]
}

=== MULTIDIMENSIONAL ARRAYS ===
int[][] matrix = {{1, 2, 3}, {4, 5, 6}};
int rows = matrix.length; // 2
int cols = matrix[0].length; // 3
int element = matrix[1][2]; // 6
Arrays.sort(matrix[0]); // sort first row: [1, 2, 3]
Arrays.sort(matrix[1]); // sort second row: [4, 5, 6]

=== ARRAY COVARIANCE (DANGEROUS) ===
String[] strings = {"Java", "Python"};
Object[] objects = strings; // Covariant assignment
objects[0] = new Integer(42); // Runtime error! ArrayStoreException
// This is why generics are invariant: List<String> is NOT a List<Object>

=== ARRAY PERFORMANCE TIPS ===
- Use Arrays.copyOf() instead of manual copying
- Use Arrays.sort() for primitive arrays (faster than Collections.sort())
- Use Arrays.parallelSort() for large arrays
- Use Arrays.stream() for functional operations
- Use Arrays.binarySearch() only on sorted arrays
- Use Arrays.fill() for initializing arrays with same value
*/