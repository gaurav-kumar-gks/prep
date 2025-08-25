"""
Reservoir sampling
"""

"""

Reservoir Sampling is a randomized algorithm for choosing a sample of k items from a list of n items, 
where n is either a very large or unknown number. 
Typically, it's used when the dataset is too large to fit into memory.

The algorithm ensures that each item in the dataset has an equal probability of being included in the sample. 
It's particularly useful for streaming data.

Here's a step-by-step plan for implementing Reservoir Sampling in Python:
"""

import random


def reservoir_sampling(stream, k):
    # Initialize the reservoir with the first k elements from the stream
    reservoir = stream[:k]
    
    # Iterate over the rest of the stream
    for i in range(k, len(stream)):
        # Pick a random index from 0 to i
        j = random.randint(0, i)
        
        # If the random index is within the range of the reservoir, replace the element
        if j < k:
            reservoir[j] = stream[i]
    
    return reservoir
