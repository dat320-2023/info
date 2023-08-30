class Species:
    def Describe(self):
        print("I belong to species")

class Cat(Species):
    def __init__(self, name):
        self.name = name

class Human(Species):
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def set_name(self, name):
        self.name = name

    def set_age(self, age):
        self.age = age

    def __str__(self):
        return f"Name: {self.name}, Age: {self.age}"
    def Describe(self):
        print( "I am a human ", self)
    



if __name__ == "__main__":
    cat=Cat("Minouch")
    cat.Describe()
    
    human = Human("Geir", 30)
    #print(human)

    #human.set_name("Geir Arne")
    #human.set_age(31)

    #print(human)

    human.Describe()