"""
Generators
"""

"""
Generators functions
returns a generator iterator that yields data one at a time
"""
def countdown(start, end):
    current = start
    while current >= end:
        yield current
        current -= 1

for num in countdown(5, 1):
    print(num)  # Outputs: 5 4 3 2 1
    
"""
Generator expressions
"""

l = [1, 2, 3, 4]
# list comprehension
l1 = [i for i in l]
# generator expression
l2 = (i for i in l)




"""
Data pipeline
ETL: Extract Transform Load
"""

def read_data(filename):
    with open(filename, 'r') as file:
        for line in file:
            print("reading line: ", line)
            yield line.strip()

def filter_data(data, keyword):
    for line in data:
        print("filtering line: ", line)
        if keyword in line:
            yield line

def process_data(data):
    for line in data:
        print("processing line: ", line)
        yield line.upper()

def write_data(data, output_filename):
    with open(output_filename, 'w') as file:
        for line in data:
            print("writing line: ", line)
            file.write(line + '\n')

# data = read_data('./input.txt')
# filtered_data = filter_data(data, 'important')
# processed_data = process_data(filtered_data)
# write_data(processed_data, './output.txt')