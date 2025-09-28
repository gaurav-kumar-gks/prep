package basics;

/*
 * List (ArrayList, LinkedList, Vector, Stack)
 * Set (HashSet, LinkedHashSet, TreeSet, EnumSet)
 * Map (HashMap, LinkedHashMap, TreeMap, Hashtable, EnumMap)
 * Queue & Deque (Queue, Deque, PriorityQueue, ArrayDeque, Stack)
 * Streams & Functional Collections
 * Custom Collections (implementing Collection, List, Set, Map)
 */


 /*
  
                                        iterable(I)
                                            |
                                        collection(I)
                /                           |                         \
            list(I)                       queue(I)                   set(I)
        /      |     \                    /    \                    /      \
arraylist linkedlist vector   priorityqueue  deque(I)           hashset     sortedset(I)
                        |                     /    \               |          \
                      stack             linkedlist arraydeque   linkedhashset  navigableset(I)
                                                                                |
                                                                                treeset        

          map(I)
        /       \
sortedmap(I)     abstractmap
      |           /      \
navigablemap(I) hashmap enummap
      |  
treemap
  */









// --------------------------------------------------------------
// List (ArrayList, LinkedList, Vector, Stack)
// --------------------------------------------------------------
/*
 * - List is an ordered collection (sequence) that allows duplicates.
 * - Elements are indexed (get/set by position).
 * - Implementations: ArrayList (resizable array), LinkedList (doubly-linked list), Vector (legacy, synchronized), Stack (legacy, LIFO)
 */
import java.util.*;

class ListDemo {
    public static void main(String[] args) {
        ArrayList<String> arrayList = new ArrayList<String>();
        ArrayList<String> arrayList2 = new ArrayList<String>(Arrays.asList("a", "b", "c"));
        ArrayList<String> arrayList3 = new ArrayList<String>(arrayList2);
        ArrayList<String> arrayList4 = new ArrayList<String>(10);
        ArrayList<String> arrayList5 = new ArrayList<String>(Collections.nCopies(10, "a"));
        
        arrayList3.forEach(System.out::println);
        System.out.println(arrayList4);
        for (String s : arrayList5) {System.out.println(s);}
        arrayList.add("a");
        arrayList.addAll(Arrays.asList("b", "c"));
        System.out.println(arrayList.get(0));
        arrayList.set(2, "d");
        arrayList.indexOf("b");
        arrayList.contains("c");
        arrayList.remove("b");
        arrayList.size();
        arrayList.isEmpty();
        arrayList.clear();

        LinkedList<String> linkedList = new LinkedList<String>();
        linkedList.add("a");
        linkedList.addFirst("b");
        linkedList.addLast("c");
        linkedList.remove("b");
        linkedList.removeFirst();
        linkedList.removeLast();
        linkedList.remove(1);
        linkedList.removeAll(Arrays.asList("a", "c"));
        linkedList.removeIf(s -> s.equals("d"));
        linkedList.forEach(System.out::println);
        linkedList.size();
        linkedList.clear();
    }
}

// --------------------------------------------------------------
// Set (HashSet, LinkedHashSet, TreeSet, EnumSet)
// --------------------------------------------------------------
/*
 * - Set is a collection that does not allow duplicates.
 * - Implementations: HashSet (unordered, fast), LinkedHashSet (insertion order), TreeSet (sorted), EnumSet (for enums).
 */
class SetDemo {
    public static void main(String[] args) {
        System.out.println("set");
        Set<String> hashset = new HashSet<String>();
        Set<String> hashset2 = new HashSet<String>(Arrays.asList("a", "b", "c"));
        Set<String> hashset3 = new HashSet<String>(hashset2);
        Set<String> set4 = new HashSet<String>(Collections.nCopies(10, "a"));
        Set<String> hashset5 = new HashSet<String>(Set.of("a", "b", "c"));
        hashset3.forEach(System.out::println);
        hashset.add("a");
        hashset.addAll(Arrays.asList("b", "c"));
        hashset.contains("b");
        hashset.remove("b");
        hashset.size();
        hashset.isEmpty();
        hashset.clear();
        hashset2.equals(hashset3);

        Set<String> linkedHashSet = new LinkedHashSet<>(Arrays.asList("X", "Y"));
        System.out.println("LinkedHashSet: " + linkedHashSet); // [X, Y]

        TreeSet<Integer> treeSet = new TreeSet<Integer>(Arrays.asList(5, 2, 1, 4, 5));
        Comparator<Integer> customComparator = new Comparator<Integer>() {
            @Override
            public int compare(Integer a, Integer b) {
                return b - a;
            }
        };
        Comparator<Integer> customComparator2 = (a, b) -> b - a;
        TreeSet<Integer> customTreeSet = new TreeSet<Integer>(customComparator);
        TreeSet<String> reverseSet = new TreeSet<String>(Collections.reverseOrder());

        treeSet.first();
        treeSet.last();
        treeSet.headSet(3);
        treeSet.tailSet(3);
        treeSet.subSet(2, 4);
        treeSet.higher(3);
        treeSet.lower(3);
        treeSet.ceiling(3);
        treeSet.floor(3);
        treeSet.pollFirst();
        treeSet.pollLast();
        treeSet.size();
        treeSet.isEmpty();
        treeSet.clear();
        treeSet.equals(customTreeSet);
    }
}


// --------------------------------------------------------------
// Map (HashMap, LinkedHashMap, TreeMap, Hashtable, EnumMap)
// --------------------------------------------------------------
/*
 * Theory:
 * - Map stores key-value pairs. Keys are unique; values can be duplicated.
 * - Implementations: HashMap (unordered, fast), LinkedHashMap (insertion order), TreeMap (sorted), Hashtable (legacy, synchronized), EnumMap (for enums).
 */
class MapDemo {
    public static void main(String[] args) {
        Map<String, Integer> hashMap = new HashMap<>();
        Map<String, Integer> hashMap2 = new HashMap<>(Map.of("A", 1, "B", 2));
        hashMap.put("A", 1);
        hashMap.putIfAbsent("D", 4);
        hashMap.putAll(Map.of("E", 5, "F", 6));
        hashMap.get("B");
        hashMap.containsKey("A");
        hashMap.getOrDefault("Z", -1);
        hashMap.remove("A");
        hashMap.clear();
        for (String key : hashMap.keySet()) System.out.println(key);
        for (Integer value : hashMap.values()) System.out.println(value);
        for (Map.Entry<String, Integer> entry : hashMap.entrySet())
            System.out.println(entry.getKey() + " -> " + entry.getValue());
        hashMap.forEach((k, v) -> System.out.println(k + ":" + v));

        Map<String, Integer> linkedHashMap = new LinkedHashMap<>(Map.of("X", 10, "Y", 20));
        Map<String, Integer> treeMap = new TreeMap<>();
        treeMap.put("C", 100); treeMap.put("A", 200); treeMap.put("B", 300);
        System.out.println("TreeMap: " + treeMap); // {A=200, B=300, C=100}
    }
}


// --------------------------------------------------------------
// Queue & Deque (Queue, Deque, PriorityQueue, ArrayDeque, Stack)
// --------------------------------------------------------------
/*
 * Theory:
 * - Queue: FIFO (first-in, first-out) data structure. Used for scheduling, buffering, etc.
 * - Deque: Double-ended queue (add/remove from both ends). Can be used as stack or queue.
 * - PriorityQueue: Elements ordered by priority (natural order or Comparator).
 * - Stack: LIFO (last-in, first-out). Legacy; use Deque for stack behavior.
 */
class QueueDemo {
    public static void main(String[] args) {
        Queue<String> queue = new LinkedList<>();
        queue.add("A"); queue.add("B");
        System.out.println("Queue remove: " + queue.remove()); // A
        System.out.println("Queue peek: " + queue.peek()); // B

        Deque<String> deque = new ArrayDeque<>();
        deque.addFirst("X"); deque.addLast("Y");
        System.out.println("Deque removeFirst: " + deque.removeFirst()); // X
        System.out.println("Deque removeLast: " + deque.removeLast()); // Y

        PriorityQueue<Integer> pq = new PriorityQueue<>();
        pq.add(3); pq.add(1); pq.add(2);
        System.out.println("PriorityQueue peek: " + pq.peek()); // 1
        while (!pq.isEmpty()) System.out.println(pq.poll()); // 1 2 3
        // pq with comparator for min heap
        PriorityQueue<String> minHeap = new PriorityQueue<>( (a, b) -> a.compareTo(b) );
        // pq with comparator for max heap
        PriorityQueue<String> maxHeap = new PriorityQueue<>( (a, b) -> b.compareTo(a) );
    }
}
/*
 * - Use Deque for stack/queue behavior (ArrayDeque is fast, no capacity restrictions).
 * - Use PriorityQueue for priority-based scheduling.
 * - PriorityQueue does not guarantee full ordering (only head is min/max).
 * - Custom comparators for PriorityQueue.
 */

/*
 * - Collections.synchronizedList/Set/Map wraps collections for thread safety.
 * - java.util.concurrent provides high-performance concurrent collections (ConcurrentHashMap, CopyOnWriteArrayList, etc.).
 * - Use Collections.unmodifiableList/Set/Map or List.of/Set.of/Map.of (Java 9+) for immutable collections.
 */


class UtilityClasses {
        
    // Collections utility class
    public static void collectionsDemo() {
        System.out.println("\n=== COLLECTIONS UTILITY DEMONSTRATION ===");
        
        List<String> list = new ArrayList<>(Arrays.asList("c", "a", "b", "d"));
        
        // Sorting
        Collections.sort(list);               // O(n log n) - natural order
        Collections.sort(list, Collections.reverseOrder()); // O(n log n) - reverse order
        Collections.sort(list, String::compareToIgnoreCase); // O(n log n) - custom comparator
        
        // Searching
        int index = Collections.binarySearch(list, "b"); // O(log n) - must be sorted
        int index2 = Collections.binarySearch(list, "b", String::compareToIgnoreCase); // O(log n) - with comparator
        
        // Reversing and shuffling
        Collections.reverse(list);            // O(n) - reverse order
        Collections.shuffle(list);            // O(n) - random shuffle
        Collections.shuffle(list, new Random(42)); // O(n) - shuffle with seed
        
        // Filling and copying
        List<String> copy = new ArrayList<>(Collections.nCopies(4, "default")); // O(n) - create list with copies
        Collections.fill(copy, "filled");     // O(n) - fill with value
        Collections.copy(copy, list);         // O(n) - copy elements
        
        // Frequency and disjoint
        int frequency = Collections.frequency(list, "a"); // O(n) - count occurrences
        boolean disjoint = Collections.disjoint(list, Arrays.asList("x", "y", "z")); // O(n) - check if disjoint
        
        // Min and max
        String min = Collections.min(list);   // O(n) - minimum element
        String max = Collections.max(list);   // O(n) - maximum element
        String min2 = Collections.min(list, String::compareToIgnoreCase); // O(n) - with comparator
        
        // Synchronized and unmodifiable wrappers
        List<String> syncList = Collections.synchronizedList(list); // Thread-safe wrapper
        List<String> unmodifiableList = Collections.unmodifiableList(list); // Immutable wrapper
        
        // Empty and singleton collections
        List<String> emptyList = Collections.emptyList(); // Immutable empty list
        List<String> singletonList = Collections.singletonList("single"); // Immutable singleton list
        Set<String> singletonSet = Collections.singleton("single"); // Immutable singleton set
        Map<String, String> singletonMap = Collections.singletonMap("key", "value"); // Immutable singleton map
        
        System.out.println("List: " + list);
        System.out.println("Copy: " + copy);
        System.out.println("Frequency of 'a': " + frequency);
        System.out.println("Disjoint: " + disjoint);
        System.out.println("Min: " + min + ", Max: " + max);
    }
}

// --------------------------------------------------------------
// Summary Table/Quick Reference
// --------------------------------------------------------------
/*
 * Type      | Allows Duplicates | Ordered | Sorted | Nulls | Thread Safe | Typical Use
 * ----------|-------------------|---------|--------|-------|-------------|----------------------
 * ArrayList | Yes               | Yes     | No     | Yes   | No          | General purpose list
 * LinkedList| Yes               | Yes     | No     | Yes   | No          | Frequent inserts/removes
 * HashSet   | No                | No      | No     | Yes   | No          | Unique elements
 * TreeSet   | No                | Yes     | Yes    | No    | No          | Sorted unique elements
 * HashMap   | Keys: No, Vals:Yes| No      | No     | Yes   | No          | Key-value pairs
 * TreeMap   | Keys: No, Vals:Yes| Yes     | Yes    | No    | No          | Sorted key-value pairs
 * ArrayDeque| Yes               | Yes     | No     | Yes   | No          | Stack/queue/deque
 */