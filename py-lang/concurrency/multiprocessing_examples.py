"""
The multiprocessing is used for parallel execution of tasks.
It allows you to create processes, which are separate instances of the Python interpreter, each with its own memory space. 
Useful for CPU-bound tasks, as it allows for true parallelism by using multi-core processors bypassing the GIL (which limits the Threading)

"""

from multiprocessing import (
    Process, Pool, TimeoutError, Pipe, Queue, Array, Value
)
import os
import time

def process_file(file_path):
    print(f"Processing {file_path} in process {os.getpid()}")


def square(x):
    return x * x

def sender(conn):
    conn.send("Hello from sender")
    response = conn.recv()
    print(f"Sender received: {response}")
    conn.close()

def receiver(conn):
    message = conn.recv()
    print(f"Receiver received: {message}")
    conn.send("Hello from receiver")
    conn.close()


def f(q):
    q.put([42, None, 'hello'])
    
def producer(queue):
    for i in range(5):
        item = f"item-{i}"
        queue.put(item)
        print(f"Produced {item}")
        time.sleep(1)

def consumer(queue):
    while True:
        item = queue.get()
        if item is None:
            break
        print(f"Consumed {item}")
        
def shared_memory(v, arr):
    v.value = 3.141592653589793
    arr[0] = 5
    print(v.value)
    print(arr[:])

if __name__ == "__main__":
    file_paths = ["file1.txt", "file2.txt", "file3.txt"]
    processes = [Process(target=process_file, args=(file_path,)) for file_path in file_paths]

    for process in processes:
        process.start()

    for process in processes:
        process.join()
    
    """
    Exchange data between processes using Queue, Pipe
    """
    # Queue: It is thread and process safe.
    queue = Queue()
    producer_process = Process(target=producer, args=(queue,))
    consumer_process = Process(target=consumer, args=(queue,))

    producer_process.start()
    consumer_process.start()

    producer_process.join()
    queue.put(None)  # Signal consumer to exit
    consumer_process.join()
    
    
    # Pipe: It is used for communication between two processes.
    parent_conn, child_conn = Pipe()
    sender_process = Process(target=sender, args=(child_conn,))
    receiver_process = Process(target=receiver, args=(parent_conn,))

    sender_process.start()
    receiver_process.start()

    sender_process.join()
    receiver_process.join()
        
    """
    Sharing state between processes using Value, Array
    """  
    v = Value('d', 0.0)  # 'd' -> double
    arr = Array('i', range(10))  # 'i' -> int
    print(v.value)
    print(arr[:])
    test_process = Process(target=f, args=(v,))
    test_process.start()
    test_process.join()
    
    
    """
    Pool: A pool of worker processes, useful for parallel execution of tasks.
    """
    with Pool(4) as pool:
        # map:
        # blocking, supports only one iterable, may cause high mem usage for long iterables - use imap
        print(pool.map(square, range(10))) # print "[0, 1, 4,..., 81]"
        
        # map_async:
        # evaluate "square(10)" asynchronously
        multiple_results = [pool.map_async(square, (i,)) for i in range(4)]
        for res in multiple_results:
            try:
                print(res.get(timeout=1)) # print "[0, 1, 4, 9]"
            except TimeoutError as e:
                print("TimeoutError in map_async")
        
        # starmap:     
        # supports multiple iterables
        print(pool.starmap(pow, [(2, 5), (3, 2), (10, 3)])) # print "[2^5, 3^2, 10^3]"
        try:
            print(pool.starmap_async(pow, [(2, 5), (3, 2), (10, 3)]).get()) # print "[2^5, 3^2, 10^3]"
        except TimeoutError as e:
            print("TimeoutError in starmap_async")
        
        
        # imap:
        # evaluate "square(i)" asynchronously
        # results are in order
        for i in pool.imap(square, range(10)):
            print(i)
        
        # imap_unordered:
        # evaluate "square(i)" asynchronously
        # results may be in any order
        for i in pool.imap_unordered(square, range(10)):
            print(i)
        
        # apply:
        # evaluate "square(10)" in this process
        print(pool.apply(square, (10,))) # prints "100"
        
        # apply_async:
        # evaluate "square(20)" asynchronously
        res = pool.apply_async(square, (20,))  # runs in *only* one process
        try:
            print(res.get(timeout=1))  # prints "400"
        except TimeoutError as e:
            print("TimeoutError in apply_async")
        