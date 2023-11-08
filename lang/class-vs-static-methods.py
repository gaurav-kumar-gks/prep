class MyClass:
    def method(self):
        return 'instance method called', self

    @classmethod
    def classmethod(cls):
        return 'class method called', cls

    @staticmethod
    def staticmethod():
        return 'static method called'
    
    
"""
Instance method:

self: points to an instance of MyClass when the method is called 
can access instance (object) and class state through self

obj.method() is equivalent to MyClass.method(obj)
"""

"""
Class method:

cls: points to the class and not the object instance when the method is called
can modify class state that would apply across all instances of the class

cls.method() is equivalent to MyClass.method(cls)
"""

"""
Static method:

no self or cls parameter
cannot access or modify class state
they’re primarily a way to namespace your methods   

They work like regular functions but belong to the class’s (and every instance’s) namespace
"""