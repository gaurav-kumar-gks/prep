"""
Thread: A separate flow of execution, it's a lightweight process.
Process: A program in execution.

Threading -> used for concurrency (not parallelism)
"""

from threading import Condition, Event, Semaphore, Thread, Lock, RLock
import time

def print_numbers():
    for i in range(1, 6):
        print(i)
        
thread = Thread(target=print_numbers)
thread.start()
thread.join() # wait for the thread to finish

class PrintNumbers(Thread):
    def run(self):
        for i in range(1, 6):
            print(i)
            
thread = PrintNumbers()
thread.start()
thread.join() # wait for the thread to finish

"""
Lock: synchronization primitive that is not owned by a particular thread when locked

Rlock: reentrant lock, synchronization primitive that may be acquired multiple times by the same thread
in the locked state, some thread owns the lock; 
in the unlocked state, no thread owns it.
"""
def thread_safe_function_using_lock():
    lock = Lock()
    with lock:
        # some thread safe operation, e.g. writing to a file or database
        pass

    lock = Lock()
    is_lock_acquired = lock.acquire(blocking=True, timeout=-1)
    try:
        # some op that needs to be thread safe here
        is_locked = lock.locked()  # returns True if the lock is acquired
        pass
    finally:
        lock.release()


rlock = RLock()

def recursive_function(n):
    rlock.acquire()
    try:
        if n > 0:
            print(f"Recursion level: {n}")
            recursive_function(n - 1)
    finally:
        rlock.release()

recursive_function(3)


"""
Condition:
A Condition variable allows one or more threads to wait until 
they are notified by another thread. 
It is often used in conjunction with a Lock.

wait: release the lock, and block until another thread awakens it
notify: wake up one of the threads waiting on the condition
notify_all: wake up all threads waiting on the condition
"""

"""
Event: synchronization object that can be used to signal between threads.

set: set the internal flag to true
clear: set the internal flag to false
wait: block until the internal flag is true
"""

import random
from queue import Queue

condition = Condition()
queue = Queue()
shutdown_event = Event()

def consumer(consumer_id):
    print(f"Consumer {consumer_id} started...")
    while not shutdown_event.is_set():
        with condition:
            while queue.empty() and not shutdown_event.is_set():
                print(f"Consumer {consumer_id} is waiting for items...")
                condition.wait()  # Wait until notified or shutdown
            if not queue.empty():
                item = queue.get()
                print(f"Consumer {consumer_id} consumed item: {item}")
            condition.notify_all()  # Notify producers or other consumers

def producer(producer_id):
    print(f"Producer {producer_id} started...")
    while not shutdown_event.is_set():
        with condition:
            item = f"item-{producer_id}-{random.randint(1, 100)}"
            queue.put(item)
            print(f"Producer {producer_id} produced item: {item}")
            condition.notify_all()  # Notify consumers
        time.sleep(random.uniform(0.1, 0.5))

def shutdown():
    print("Waiting for a while before initiating shutdown...")
    time.sleep(1)
    print("Initiating shutdown...")
    shutdown_event.set()
    with condition:
        condition.notify_all()  # Wake up all waiting threads


consumer_threads = [Thread(target=consumer, args=(i,)) for i in range(3)]
for thread in consumer_threads:
    thread.start()

producer_threads = [Thread(target=producer, args=(i,)) for i in range(2)]
for thread in producer_threads:
    thread.start()

shutdown_thread = Thread(target=shutdown)
shutdown_thread.start()

for thread in consumer_threads + producer_threads + [shutdown_thread]:
    thread.join()


"""
Semaphore: synchronization primitive that can be used to 
control access to a shared resource through the use of a counter.
"""

semaphore = Semaphore(2)

def limited_access_function(thread_id):
    print(f"Thread {thread_id} is waiting to access the shared resource")
    with semaphore:
        print(f"Thread {thread_id} is accessing the shared resource")
        time.sleep(1)
        print(f"Thread {thread_id} is releasing the shared resource")

threads = [Thread(target=limited_access_function, args=(i,)) for i in range(5)]

for thread in threads:
    thread.start()

for thread in threads:
    thread.join()
