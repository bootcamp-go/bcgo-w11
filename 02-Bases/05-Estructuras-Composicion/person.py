def NewPersonJohn():
    return Person('John', 25)

def NewPersonJane():
    return Person('Jane', 22)

class Person:
    name:str
    age:int

    def say_hello(self):
        print(f'Hello, my name is {self.name} and I am {self.age} years old')