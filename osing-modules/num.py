import os 
import sys 
import phonenumbers
from phonenumbers import carrier, geocoder, timezone
from tabulate import tabulate

def number_scanner(phone_number):
    number = phonenumbers.parse(phone_number)
    description = geocoder.description_for_number(number, "en" )
    supplier = carrier.name_for_number(number, "en")
    info = [["State or Country", "Supplier",],
            [description, supplier,]]
    data = str(tabulate(info, headers="firstrow", tablefmt="github"))
    return data 

if __name__ == "__main__":
    print(" Example: +1 000-000-000")
    number = str(input("Number >>> "))
    os.system("clear")
    print(number_scanner(number))