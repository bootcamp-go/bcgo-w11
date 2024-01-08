# abstract class Vehicle
class Vehicle:
    def Start():
        pass

    def Stop():
        pass


# implementation of Vehicle
# Car
class Car(Vehicle):
    def Start():
        print('Car started')

    def Stop():
        print('Car stopped')

    
# Truck
class Truck(Vehicle):
    def Start():
        print('Truck started')

    def Stop():
        print('Truck stopped')