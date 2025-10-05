import threading
import multiprocessing
import concurrent.futures
import asyncio
import time
import queue
import os

def thread_worker(n):
    time.sleep(n)
    print(f"Thread {threading.current_thread().name} done after {n}s")

def demo_thread():
    print("Demo Thread:")
    threads = [threading.Thread(target=thread_worker, args=(i,), name=f"T{i}", daemon=True) for i in [1, 2]]
    for t in threads:
        t.start()  # Start thread
    while any(t.is_alive() for t in threads):  # Check if alive
        time.sleep(0.5)
    for t in threads:
        t.join(timeout=1.0)  # Join with timeout
    print(f"Thread IDs: {[t.ident for t in threads]}")  # Get ident
    # native_id not used in this example (Python 3.8+ property)

def process_worker(n, q):
    time.sleep(n)
    q.put(f"Process {os.getpid()} done after {n}s")  # Put result to queue

def demo_process():
    print("\nDemo Process:")
    if __name__ == "__main__":  # Required for processes
        mp_q = multiprocessing.Queue()  # Process-safe queue
        processes = [multiprocessing.Process(target=process_worker, args=(i, mp_q), name=f"P{i}", daemon=True) for i in [1, 2]]
        for p in processes:
            p.start()  # Start process
        while any(p.is_alive() for p in processes):  # Check if alive
            time.sleep(0.5)
        results = [mp_q.get() for _ in processes]  # Get from queue
        print("Results:", results)
        for p in processes:
            p.join(timeout=1.0)  # Join with timeout
        print(f"Process PIDs: {[p.pid for p in processes]}")  # Get pid
        print(f"Exitcodes: {[p.exitcode for p in processes]}")  # Get exitcode
        # terminate/kill not used (for forceful stop)
        # sentinel not used (for polling)

def lock_worker(lock, shared):
    with lock:  # Acquire and release
        shared[0] += 1
        print(f"Lock acquired, shared: {shared[0]}")

def demo_lock():
    print("\nDemo Lock:")
    lock = threading.Lock()
    shared = [0]
    threads = [threading.Thread(target=lock_worker, args=(lock, shared)) for _ in range(3)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()
    print(f"Final shared: {shared[0]}")
    # try_acquire not used (non-blocking acquire, Python 3.2+)

def rlock_worker(rlock, shared):
    with rlock:  # Acquire (reentrant)
        with rlock:  # Reacquire same thread
            shared[0] += 1
            print(f"RLock reacquired, shared: {shared[0]}")

def demo_rlock():
    print("\nDemo RLock:")
    rlock = threading.RLock()
    shared = [0]
    threads = [threading.Thread(target=rlock_worker, args=(rlock, shared)) for _ in range(3)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()
    print(f"Final shared: {shared[0]}")

def semaphore_worker(sem, shared):
    with sem:  # Acquire (decrement)
        shared[0] += 1
        time.sleep(0.5)
        print(f"Semaphore acquired, shared: {shared[0]}")

def demo_semaphore():
    print("\nDemo Semaphore:")
    sem = threading.Semaphore(2)  # Initial value 2
    shared = [0]
    threads = [threading.Thread(target=semaphore_worker, args=(sem, shared)) for _ in range(4)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()
    print(f"Final shared: {shared[0]}")
    # release not used directly (increment, used in with)

def event_worker(event):
    print("Waiting for event")
    event.wait(timeout=2.0)  # Wait for set
    if event.is_set():  # Check if set
        print("Event set, proceeding")

def demo_event():
    print("\nDemo Event:")
    event = threading.Event()
    threads = [threading.Thread(target=event_worker, args=(event,)) for _ in range(2)]
    for t in threads:
        t.start()
    time.sleep(1)
    event.set()  # Signal all
    for t in threads:
        t.join()
    event.clear()  # Reset

def condition_producer(cond, shared):
    with cond:  # Acquire lock
        shared.append("item")
        cond.notify()  # Wake one waiter

def condition_consumer(cond, shared):
    with cond:  # Acquire lock
        while not shared:
            cond.wait(timeout=1.0)  # Wait and release lock temporarily
        print(f"Consumed: {shared.pop()}")

def demo_condition():
    print("\nDemo Condition:")
    cond = threading.Condition()  # With internal lock
    shared = []
    prod = threading.Thread(target=condition_producer, args=(cond, shared))
    cons = threading.Thread(target=condition_consumer, args=(cond, shared))
    cons.start()
    time.sleep(0.5)
    prod.start()
    prod.join()
    cons.join()
    # notify_all not used (wake all waiters)

def barrier_worker(barrier, id):
    print(f"Worker {id} waiting")
    barrier.wait(timeout=2.0)  # Wait for all to arrive
    print(f"Worker {id} passed barrier")

def demo_barrier():
    print("\nDemo Barrier:")
    barrier = threading.Barrier(3)  # For 3 parties
    threads = [threading.Thread(target=barrier_worker, args=(barrier, i)) for i in range(3)]
    for t in threads:
        t.start()
    for t in threads:
        t.join()
    print(f"Broken: {barrier.broken}, Parties: {barrier.parties}")  # Check state
    barrier.reset()  # Reset for reuse
    # abort not used (break barrier)

def pool_task(n):
    time.sleep(n)
    return n * n

def demo_threadpool():
    print("\nDemo ThreadPoolExecutor:")
    with concurrent.futures.ThreadPoolExecutor(max_workers=2) as executor:  # Init with workers
        futures = [executor.submit(pool_task, i) for i in [1, 2]]  # Submit tasks
        mapped = list(executor.map(pool_task, [3, 4], timeout=5.0))  # Map with timeout
        results = [f.result(timeout=3.0) for f in futures]  # Get result
        print("Submitted results:", results)
        print("Mapped:", mapped)
    # shutdown not used directly (handled by with, or manual call)

def demo_processpool():
    print("\nDemo ProcessPoolExecutor:")
    if __name__ == "__main__":
        with concurrent.futures.ProcessPoolExecutor(max_workers=2) as executor:  # Init with workers
            futures = [executor.submit(pool_task, i) for i in [1, 2]]  # Submit tasks
            mapped = list(executor.map(pool_task, [3, 4], timeout=5.0))  # Map with timeout
            results = [f.result(timeout=3.0) for f in futures]  # Get result
            print("Submitted results:", results)
            print("Mapped:", mapped)
        # shutdown not used directly (handled by with)

async def async_worker(n):
    await asyncio.sleep(n)  # Async wait
    print(f"Async task done after {n}s")
    return n * n

async def demo_asyncio():
    print("\nDemo Asyncio:")
    loop = asyncio.get_running_loop()  # Get loop
    tasks = [asyncio.create_task(async_worker(i)) for i in [1, 2]]  # Create tasks
    await asyncio.sleep(0.5)  # Non-blocking sleep
    done, pending = await asyncio.wait(tasks, timeout=2.0)  # Wait with timeout
    results = [t.result() for t in done]  # Get result
    print("Done results:", results)
    for t in tasks:
        if not t.done():  # Check done
            t.cancel()  # Cancel pending
    # gather not used (await multiple)
    # ensure_future not used (schedule coroutine)
    # run_in_executor not used (run sync in thread)

if __name__ == "__main__":
    demo_thread()
    demo_process()
    demo_lock()
    demo_rlock()
    demo_semaphore()
    demo_event()
    demo_condition()
    demo_barrier()
    demo_threadpool()
    demo_processpool()
    asyncio.run(demo_asyncio())  # Run async main