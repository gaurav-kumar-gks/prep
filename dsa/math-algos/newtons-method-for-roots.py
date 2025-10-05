"""
Newton's method for finding roots
"""

"""
If we want solution of f(x) = 0
where f(x) is continuous and differentiable on domain [a, b]

suppose we have initial guess of Xi
Draw the tangent to the graph of f(x) at x=Xi 
and find the point of intersection of tangent with X axis
Xi+1 is set to this point

Xi+1 = Xi - f(Xi) / f'(Xi)
"""

def sqrt(n):
    """
    f(x) = x*x - n = 0
    f'(x) = 2*x
    f(x)/f'(x) = (x - n/x)/2
    nx = x - f(x)/f'(x) = (x + n/x) / 2
    """
    x = n
    while True:
        nx = (x + n / x) // 2
        if abs(n - nx) < 1e-9:
            # converged
            break
        x = nx
    return x
