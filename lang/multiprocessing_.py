"""
The multiprocessing is used for parallel execution of tasks.
It allows you to create processes, which are separate instances of the Python interpreter, each with its own memory space. 
Useful for CPU-bound tasks, as it allows for true parallelism by using multi-core processors bypassing the GIL (which limits the Threading)

"""

from multiprocessing import Process, Pool
import os

def process_file(file_path):
    print(f"Processing {file_path} in process {os.getpid()}")


def square(x):
    return x * x

if __name__ == "__main__":
    file_paths = ["file1.txt", "file2.txt", "file3.txt"]
    processes = [Process(target=process_file, args=(file_path,)) for file_path in file_paths]

    for process in processes:
        process.start()

    for process in processes:
        process.join()

    with Pool(4) as pool:
        results = pool.map(square, [1, 2, 3, 4, 5])
        print(results)  # Output: [1, 4, 9, 16, 25]