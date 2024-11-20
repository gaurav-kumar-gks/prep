"""
Global Interpreter Lock (GIL) 


What:
- Mutex that protects access to Python objects, preventing multiple native threads from executing Python bytecodes simultaneously. 
- impacts how Python handles multi-threading and parallelism.

- GIL ensures that only one thread executes Python bytecode at a time. 
- Even if you have multiple threads, only one can be executing Python code, while others are either waiting or executing non-Python code (e.g., I/O operations).
- GIL is periodically released and reacquired to allow other threads to run. This happens every few milliseconds or when a thread performs a blocking I/O operation.

Why: 
- This lock is necessary because Python's memory management is not thread-safe.
- Python uses reference counting for memory management. 
- Each object has a reference count, and when this count drops to zero, the memory is deallocated. 
- Without the GIL, multiple threads could simultaneously modify the reference count, leading to race conditions and memory corruption.
 

Impact:
1. I/O-Bound Programs: The GIL has minimal impact on I/O-bound programs, where threads spend most of their time waiting for external resources (e.g., network, disk). 
In these cases, threads can release the GIL while waiting, allowing other threads to run.

2. CPU-Bound Programs: The GIL can be a significant bottleneck for CPU-bound programs, where threads perform intensive computations. 
In such cases, the GIL prevents true parallel execution on multi-core processors, as only one thread can execute Python code at a time.
Multi-Core Processors: The GIL limits the ability of Python programs to take full advantage of multi-core processors for CPU-bound tasks. This is because the GIL prevents multiple threads from executing Python code in parallel.

Workarounds:
1. Multiprocessing: The multiprocessing module provides a way to bypass the GIL by using separate processes instead of threads. 
Each process has its own Python interpreter and memory space, allowing true parallel execution.

2. Asyncio: For I/O-bound tasks asyncio can be an effective

3. Using C Extensions: By writing performance-critical code in C or Cython, you can bypass the GIL and achieve parallelism.

4. Jython and IronPython: These implementations of Python do not have a GIL, allowing true parallel execution on multi-core processors.
"""