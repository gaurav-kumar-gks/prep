"""
*args and **kwargs are mostly used in function definitions. 
*args and **kwargs allow you to pass a variable number of arguments to a function. 
"""

"""
Left side	                     Divider	Right side
Positional-only arguments	        /	    Positional or keyword arguments
Positional or keyword arguments 	*	    Keyword-only arguments
"""


"""
w/o any * or /
"""


def print_three_members(member1, member2, member3):
    print(f"member1 is {member1}")
    print(f"member2 is {member2}")
    print(f"member3 is {member3}")


# print_three_members("a", "b", "c") # OK
# print_three_members(member1="a", member2="b", member3="c") # OK


"""
using * 
"""


def print_three_members1(*, member1, member2, member3):
    print(f"member1 is {member1}")
    print(f"member2 is {member2}")
    print(f"member3 is {member3}")


# print_three_members1("a", "b", "c") # TypeError: print_three_members1() takes 0 positional arguments but 3 were given
# print_three_members1(member1="a", member2="b", member3="c") # OK


"""
using /
"""


def print_three_members2(member1, member2, member3, /):
    print(f"member1 is {member1}")
    print(f"member2 is {member2}")
    print(f"member3 is {member3}")


# print_three_members2("a", "b", "c") # OK
# print_three_members2(member1="a", member2="b", member3="c") # TypeError: print_three_members2() got an unexpected keyword argument 'member1'
def print_varying_members(member1, member2, *args, member3):
    print(f"member1 is {member1}")
    print(f"member2 is {member2}")
    print(f"member3 is {member3}")
    print(f"*args contains {args}")


# print_varying_members("a", "b", "c", "d", "e", member3="f") # OK
# print_varying_members("a", "b", "c", "d", "e", "f") # TypeError: print_varying_members() missing 1 required keyword-only argument: 'member3'
# print_varying_members("a", member2="b", "c", "d", member3="e") # Positional argument after keyword argument
