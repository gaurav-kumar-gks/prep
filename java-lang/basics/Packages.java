package basics;

// -----------------------------------------------------------------------------
// JAVA PACKAGES
// -----------------------------------------------------------------------------
/* 
- Java packages were introduced to solve namespace pollution and enable modular code.
Inspired by C++ namespaces, but mapped directly to directory structure.
Java 9 introduced the module system for even stronger encapsulation.


src/
    com/
        example/
            utils/
                StringUtils.java   --> package com.example.utils;

Compiled: com/example/utils/StringUtils.class

- The first non-comment line in a Java file can be a package declaration. The directory structure must match the package name (enforced by javac).
- No two top-level classes in the same package can have the same name.
- Only one public top-level class per file (file name must match class name)
- Package-private is the default if no modifier is given
- Use package-private for internal helpers/utilities
- Subpackages are independent: java.util and java.util.concurrent are unrelated
- Lowercase only, no underscores or hyphens
- Avoid deep nesting unless necessary

import java.util.List; // explicit import
import java.util.*;    // wildcard import (not recursive)
import static java.lang.Math.*; // static import (Java 5+)

- Wildcard imports only import types directly in the package, not subpackages.
- Static imports bring static members (fields/methods) into scope.

- public: accessible everywhere
- protected: accessible in same package and subclasses (even in other packages)
- (default/package-private): accessible only in the same package
- private: accessible only in the same class
*/

// -----------------------------------------------------------------------------
// JAVA MODULE SYSTEM
// -----------------------------------------------------------------------------

/*


Modules are a collection of related packages and resources that:
- Encapsulates internal implementation details
- Explicitly declares its dependencies on other modules
- Controls which packages are accessible to other modules
- Provides a clear API boundary


- Most enterprise applications don't use modules
- Spring ecosystem (Spring Boot, Spring Framework) doesn't require modules
- Maven/Gradle projects work fine without modules


├── com/
│   └── example/
│       └── mymodule/
│           ├── api/
│           │   └── PublicAPI.java
│           ├── internal/
│           │   └── InternalClass.java
│           └── Main.java
|           |__ module-info.java


module com.example.mymodule {
    // Module dependencies
    requires java.base;
    requires com.example.utils;
    requires transitive com.example.logging;
    
    // Package exports
    exports com.example.mymodule.api;
    exports com.example.mymodule.api to com.example.client;
    
    // Package opens (for reflection)
    opens com.example.mymodule.internal;
    opens com.example.mymodule.internal to com.example.framework;
    
    // Service providers
    provides com.example.mymodule.api.ServiceInterface 
        with com.example.mymodule.internal.ServiceImpl;
    
    // Service consumers
    uses com.example.mymodule.api.ServiceInterface;
}
 */

/*

REFLECTION

- Reflection enables Java code to:
- Examine class structure at runtime
- Create objects dynamically
- Invoke methods by name
- Access private fields and methods
- Modify object behavior at runtime


import java.lang.reflect.*;

Class<?> clazz = String.class;
Method[] methods = clazz.getMethods();
Field[] fields = clazz.getDeclaredFields();
Constructor<?>[] constructors = clazz.getConstructors();

*/