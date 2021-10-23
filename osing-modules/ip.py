import os
import requests
import json
import colorama 
import urllib
from urllib.request import urlopen as open
import webbrowser
import os 
from colorama import Fore, Back, Style 
import requests 
from termcolor import colored
import time
from requests import get 
import sys 
import pyfiglet
import sys
import socket
from datetime import datetime
os.system('clear')

#colors 
# RED ++++ print(\033[31m)


print(Fore.RED+"                $$$$$$\ $$$$$$$\        $$$$$$$$\ $$$$$$$\   $$$$$$\   $$$$$$\  $$\   $$\ $$$$$$$$\ $$$$$$$\  ")
print(Fore.MAGENTA+"		\_$$  _|$$  __$$\       \__$$  __|$$  __$$\ $$  __$$\ $$  __$$\ $$ | $$  |$$  _____|$$  __$$\  ")
print(Fore.RED+"  		  $$ |  $$ |  $$ |         $$ |   $$ |  $$ |$$ /  $$ |$$ /  \__|$$ |$$  / $$ |      $$ |  $$ | ")
print(Fore.MAGENTA+"  		  $$ |  $$$$$$$  |         $$ |   $$$$$$$  |$$$$$$$$ |$$ |      $$$$$  /  $$$$$\    $$$$$$$  | ")
print(Fore.RED+"  		  $$ |  $$  ____/          $$ |   $$  __$$< $$  __$$ |$$ |      $$  $$<   $$  __|   $$  __$$<  ")
print(Fore.MAGENTA+" 		  $$ |  $$ |               $$ |   $$ |  $$ |$$ |  $$ |$$ |  $$\ $$ |\$$\  $$ |      $$ |  $$ | ") 
print(Fore.RED+"		$$$$$$\ $$ |               $$ |   $$ |  $$ |$$ |  $$ |\$$$$$$  |$$ | \$$\ $$$$$$$$\ $$ |  $$ | ")
print(Fore.MAGENTA+"		\______|\__|               \__|   \__|  \__|\__|  \__| \______/ \__|  \__|\________|\__|  \__| ")
print(Fore.CYAN+"|======================================================================================================|")
print(Fore.BLUE+"|Github: ArkAngeL43 | https://github.com/ArkAngeL43                                                    |")
print(Fore.MAGENTA+"|instagram ark_angel6                                                                                  |")
print(Fore.BLUE+"|NOTES:::::: if you want to exit the script just type exit                                             |")
print(Fore.CYAN+"|======================================================================================================|")
global ip
ip = input("\033[1;36mEnter Your Target IP: \033[1;36m")

if 'Exit' in ip:
          os.system(' clear ')
          time.sleep(1)
          print(" [+] Thanks for stopping by [+] ")
          print(" [=] Have a good one :D     [=]")
          time.sleep(2)
          sys.exit()
elif 'exit' in ip:
            os.system(' clear ')
            time.sleep(1)
            print(" [+] Thanks for stopping by [+] ")
            print(" [=] Have a good one :D     [=]")
            time.sleep(2)
            sys.exit()
url = "http://ip-api.com/json/"
response = open(url + ip)
data = response.read()
values = json.loads(data)
status = values['status']
success = "success"
lat = str(values['lat'])
lon = str(values['lon'])
a = lat + ","
b = lon + "/"
c = b + "data=!3m1!1e3?hl=en"
location = a + c
print("---------------------------------")
time.sleep(0.1)
print(" [=] IP: " + values['query']        )
time.sleep(0.1)
print(" [=] Status: " + values['status']   )
time.sleep(0.1)
print(" [=] city: " + values['city']       )
time.sleep(0.1)
print(" [=] ISP: " + values['isp']         )
time.sleep(0.1)
print(" [=] latitude: " + lat              )
time.sleep(0.1)
print(" [=] longitude: " + lon             )
time.sleep(0.1)
print(" [=] country: " + values['country'] )
time.sleep(0.1)
print(" [=] region: " + values['regionName'])
time.sleep(0.1)
print(" [=] city: " + values['city']       )
time.sleep(0.1)
print(" [=] zip: " + values['zip']         )
time.sleep(0.1)
print(" [=] AS: " + values['as']           )
time.sleep(0.1)
print("---------------------------------")