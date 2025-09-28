package basics;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

// -----------------------------------------------------------------------------
//  1. CHECKED VS UNCHECKED EXCEPTIONS
// -----------------------------------------------------------------------------
// - Checked: 
//    1. must be declared or caught (IOException, SQLException)
//    2. subclass of Exception
//    3. checked exceptions are checked at compile time
//    4. generally used for recoverable conditions
// - Unchecked: 
//    1. subclass of RuntimeException, not required to be caught (NPE, IAE), but should be caught 
//    2. subclass of Error
//    3. unchecked exceptions are checked at runtime
//    4. generally used for programming errors hence not required to be caught
// - Errors: subclass of Error, should not be caught (OutOfMemoryError, StackOverflowError)
// -----------------------------------------------------------------------------

/*

Throwable
├── Error (Unchecked)
│   ├── OutOfMemoryError
│   ├── StackOverflowError
│   ├── VirtualMachineError
│   └── LinkageError
└── Exception
    ├── RuntimeException (Unchecked)
    │   ├── NullPointerException
    │   ├── IllegalArgumentException
    │   ├── ArrayIndexOutOfBoundsException
    │   ├── ArithmeticException
    │   └── ClassCastException
    └── Checked Exceptions
        ├── IOException
        ├── SQLException
        ├── ClassNotFoundException
        └── ParseException



 */
// -----------------------------------------------------------------------------
// 5. CUSTOM EXCEPTIONS & CHAINING
// -----------------------------------------------------------------------------
// - Extend Exception (checked) or RuntimeException (unchecked)
// - Always provide constructors for message and cause
// - Use exception chaining to preserve root cause
//
// Example:
//   public class MyException extends Exception {
//       public MyException(String msg, Throwable cause) { super(msg, cause); }
//   }
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 6. BEST PRACTICES
// -----------------------------------------------------------------------------
// - Never catch Throwable or Error (only catch Exception)
// - Always clean up resources (try-with-resources preferred) to prevent resource leaks
// - Log exceptions with stack trace to help debug
// - Avoid empty catch blocks and using exception for control flow
// - Document thrown exceptions in Javadoc to help others understand the code
// - Use exception chaining to preserve root cause (e.g. throw new MyException("Error", e))
// - Don't use broad Exception or Throwable catch blocks
// - Don't throw exceptions from finally blocks (suppresses original exception) 
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 9. CODE EXAMPLES
// -----------------------------------------------------------------------------
public class Exceptions {

    /*
     * Errors
     * - Serious system-level problems
     * - Should not be caught by applications
     * - Examples: OutOfMemoryError, StackOverflowError
     */

    /*
     * 
     * Checked Exceptions (Compile-time)
     * - Must be declared in method signature with throws
     * - Must be handled by caller or propagated
     * - Examples: IOException, SQLException, ClassNotFoundException
     * 
     */

    // Must declare checked exceptions
    public void readFile(String filename) throws IOException {
        FileReader reader = new FileReader(filename);
        // ... file operations
    }

    // Must handle or propagate
    public void processFile(String filename) {
        try {
            readFile(filename);
        } catch (IOException e) {
            System.err.println("File error: " + e.getMessage());
        }
    }

    /*
     * Unchecked Exceptions (Runtime)
     * - RuntimeException subclasses
     * - Not required to be declared or handled
     * - Examples: NullPointerException, IllegalArgumentException
     */

    // Not required to be declared or handled
    public void throwUnchecked() {
        throw new RuntimeException("Unchecked exception");
    }

    /*
     * Try - Catch - Finally
     *
     * - Try: code that may throw an exception
     * - Catch: code that handles the exception
     * - can catch multiple exceptions and can have multi-catch e.g.
     * catch(IOException | SQLException)
     * - Finally: code that is always executed
     * - exception in finally block will suppress the original exception
     */
    public void processUnchecked() {
        try {
            throwUnchecked();
        } catch (RuntimeException e) {
            System.err.println("Caught exception message: " + e.getMessage() + " cause: " + e.getCause());
            System.err.println("Caught exception stack trace: " + e.getStackTrace());
        } catch (Exception e) {
            System.err.println("Caught exception: " + e.getMessage());
        } finally {
            System.out.println("Finally block is always executed");
        }
    }

    /*
     * Try with resources
     */

    // Automatic resource management
    public void processFiles(String filename) throws IOException {
        try (FileReader reader = new FileReader(filename);
                BufferedReader bufferedReader = new BufferedReader(reader)) {

            String line;
            while ((line = bufferedReader.readLine()) != null) {
                System.out.println(line);
            }
        } // Resources automatically closed
    }

    // Custom AutoCloseable resource
    public class DatabaseConnection implements AutoCloseable {
        public void connect() {
            System.out.println("Connecting to database...");
        }

        public void query(String sql) {
            System.out.println("Executing: " + sql);
        }

        @Override
        public void close() {
            System.out.println("Closing database connection...");
        }
    }

    // Usage
    public void databaseOperation() {
        try (DatabaseConnection conn = new DatabaseConnection()) {
            conn.connect();
            conn.query("SELECT * FROM users");
        } // Connection automatically closed
    }

    /*
     * Custom exceptions
     * - Extend Exception (checked) or RuntimeException (unchecked)
     * - Always provide constructors for message and cause
     * - Use exception chaining to preserve root cause
     */

    public class MyCheckedException extends Exception {
        public MyCheckedException(String msg, Throwable cause) {
            super(msg, cause);
        }
    }

    // Wrapping the exception
    public void throwCustomException() throws MyCheckedException {
        throw new MyCheckedException("Wrapped Custom exception", new IOException("IO error"));
    }

    public void processCustomException() {
        try {
            throwCustomException();
        } catch (MyCheckedException e) {
            System.err.println("Caught custom exception: " + e.getMessage());
        }
    }

    public class MyUncheckedException extends RuntimeException {
        public MyUncheckedException(String msg, Throwable cause) {
            super(msg, cause);
        }
    }

    public void throwCustomUncheckedException() {
        throw new MyUncheckedException("Custom unchecked exception", new IOException("IO error"));
    }

    public void processCustomUncheckedException() {
        try {
            throwCustomUncheckedException();
        } catch (MyUncheckedException e) {
            System.err.println("Caught custom unchecked exception: " + e.getMessage());
        }
    }

}