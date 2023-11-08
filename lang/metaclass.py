"""
Metaclass is a class of a class.
"""


class ModelMeta(type):
    """
    ModelMeta is a metaclass for Model.
    """

    def __new__(cls, name, bases, attrs):
        print(f"ModelMeta.__new__ called for class '{name}'")
        fields = {
            name: value for name, value in attrs.items() if isinstance(value, Field)
        }
        attrs["_fields"] = fields
        return super().__new__(cls, name, bases, attrs)


class Field:
    def __init__(self, data_type):
        self.data_type = data_type
        print(f"Field.__init__ called for '{self.__class__.__name__}'")

    def __call__(self):
        print(f"Field.__call__ called for '{self.__class__.__name__}'")


class IntegerField(Field):
    def __init__(self):
        super().__init__("int")
        print(f"IntegerField.__init__ called")

    def __call__(self):
        print(f"IntegerField.__call__ called")


class CharField(Field):
    def __init__(self, max_length):
        super().__init__("str")
        self.max_length = max_length
        print(f"CharField.__init__ called with max_length={max_length}")

    def __call__(self):
        print(f"CharField.__call__ called")


class Model(metaclass=ModelMeta):
    def __new__(cls, *args, **kwargs):
        print(f"Model.__new__ called for class '{cls.__name__}'")
        instance = super().__new__(cls)
        return instance

    def __init__(self, **kwargs):
        print(f"Model.__init__ called for class '{self.__class__.__name__}'")
        for name, value in kwargs.items():
            setattr(self, name, value)

    def __call__(self, *args, **kwargs):
        print(f"Model.__call__ called for class '{self.__class__.__name__}'")


class Person(Model):
    age = IntegerField()
    name = CharField(max_length=20)

    def __new__(cls, *args, **kwargs):
        print("Person.__new__ called")
        instance = super().__new__(cls)
        return instance

    def __init__(self, **kwargs):
        print("Person.__init__ called")
        super().__init__(**kwargs)

    def __call__(self, *args, **kwargs):
        print("Person.__call__ called")


# name_field = CharField(max_length=100)
# age_field = IntegerField()

person = Person(name="Alice", age=30)
person()
