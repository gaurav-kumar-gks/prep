"""
Metaclass are classes that create classes. 
They provide a powerful mechanism for customizing class creation and behavior
"""


class ValidateClass(type):
    def __new__(mcs, name, bases, attrs):
        if not 'required_method' in attrs:
            raise TypeError('Subclass must implement required_method')
        return super().__new__(mcs, name, bases, attrs)

class BaseClass(metaclass=ValidateClass):
    def required_method(self):
        pass

class Subclass(BaseClass):
    def required_method(self):
        pass
    
class AutoProperty(type):
    def __new__(mcs, name, bases, attrs):
        new_attrs = {}
        for key, value in attrs.items():                
            # Check if the key is not a special attribute and the value is a string starting with '_'
            if not key.startswith('__') and isinstance(value, str) and key.startswith('_'):
                new_attrs[key[1:]] = property(fget=lambda self, key=key: getattr(self, key))
            else:
                new_attrs[key] = value
        return super().__new__(mcs, name, bases, new_attrs)

class MyAutoProClass(metaclass=AutoProperty):
    _x = 10
    _y = 20
    _z = "_abc"

class LoggingMeta(type):
    def __new__(mcs, name, bases, attrs):
        cls = super().__new__(mcs, name, bases, attrs)
        print(f"Creating class {name} with bases {bases} and attrs {attrs}")
        for name, value in attrs.items():
            if callable(value):
                def wrapper(*args, **kwargs):
                    print(f"Calling {name}")
                    x = value(*args, **kwargs)
                    print(f"Finished calling {name}")
                    return x
                setattr(cls, name, wrapper)
        return cls

class MyClass(metaclass=LoggingMeta):
    def my_method(self):
        print("Hello from my_method")

if __name__ == "__main__":
    x = MyAutoProClass()