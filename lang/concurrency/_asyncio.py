"""
asyncio is a library in Python that provides 
a framework for writing concurrent code using the async and await syntax. 

It is designed to handle asynchronous I/O operations, 
allowing you to write code that can perform multiple tasks at once without blocking the execution of your program. 

This is particularly useful for I/O-bound tasks, such as network requests, file operations, or database queries.
"""


import asyncio


async def say_hello():
    await asyncio.sleep(1)
    print("Hello!")


async def fetch_data(x):
    await asyncio.sleep(2)
    return f"Data {x}"


async def main():
    tasks = [asyncio.create_task(say_hello(i)) for i in [2, 1, 3]]
    done, pending = await asyncio.wait(tasks)
    for task in done:
        print(task.result())
    
    results = await asyncio.gather(fetch_data(1), fetch_data(2), fetch_data(3))
    print(results)

    tasks = [asyncio.create_task(say_hello(i)) for i in [2, 1, 3]]
    for completed in asyncio.as_completed(tasks):
        result = await completed
        print(result)


async def waiter(event):
    print("Waiting for the event to be set...")
    await event.wait()  # Wait until the event is set
    print("Event is set! Proceeding...")


async def setter(event):
    print("Setting the event after 3 seconds...")
    await asyncio.sleep(3)  # Simulate some work
    event.set()  # Set the event


async def main():
    event = asyncio.Event()  # Create an Event object
    await asyncio.gather(
        waiter(event), setter(event)
    )  # Run both coroutines concurrently


asyncio.run(main())
