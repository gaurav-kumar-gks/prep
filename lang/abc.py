from abc import ABC, ABCMeta, abstractmethod

class MyABC(ABC):
    @abstractmethod
    def my_abstract_method(self, arg1):
        ...
    
    @classmethod
    @abstractmethod
    def my_abstract_classmethod(cls, arg2):
        ...
    
    @staticmethod
    @abstractmethod
    def my_abstract_staticmethod(arg3):
        ...

    @property
    @abstractmethod
    def my_abstract_property(self):
        ...
        
    @my_abstract_property.setter
    @abstractmethod
    def my_abstract_property(self, value):
        ...
        

class MyConcreteClass(MyABC):
    
    def my_abstract_method(self, arg1):
        pass
    
    @classmethod
    def my_abstract_classmethod(cls, arg2):
        pass
    
    @staticmethod
    def my_abstract_staticmethod(arg3):
        pass
    
    @property
    def my_abstract_property(self):
        pass
    
    @my_abstract_property.setter
    def my_abstract_property(self, value):
        pass

# class MyABC(metaclass=ABCMeta):
    # pass