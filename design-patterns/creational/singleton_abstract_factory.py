from abc import ABC, ABCMeta, abstractmethod


class SingletonMeta(ABCMeta):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super().__call__(*args, **kwargs)
        return cls._instances[cls]


class Factory(metaclass=SingletonMeta):
    @abstractmethod
    def create_product(self):
        pass


class ConcreteFactory(Factory):
    def create_product(self):
        return ConcreteProduct()


class Product(ABC):
    @abstractmethod
    def operation(self):
        pass


class ConcreteProduct(Product):
    def operation(self):
        print("ConcreteProduct operation")


factory1 = ConcreteFactory()
product = factory1.create_product()
product.operation()
factory2 = ConcreteFactory()
