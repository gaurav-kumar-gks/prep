"""
Newton's method for finding roots
"""

"""
Lets say we want solution of f(x) = 0

where f(x) is continuous and differentiable on domain [a, b]

suppose we have Xi
then Xi is calculated as follows
Draw the tangent to the graph of f(x) at x=Xi and find the point of intersection of tangent with X axis
Xi+1 is set to this point


Xi+1 = Xi - f(Xi) / f'(Xi)
"""

"""
Popular use case -> Finding square root

lets f(x) = x^2

then lets say we want sqrt of y

then x^2 = y

or we want soln. of x^2 - y =0

so, 
f(x) = x^2 - y
f'(x) = 2x

"""


def sqrt(n):
    x = n
    while True:
        nx = (x + n / x) // 2
        if abs(n - nx) < 1e-9:
            break
        x = nx
    return x
