import os 
import sys 
import time 
import colorama 
import pyfiglet 
from colorama import Fore 
import json
import subprocess


os.system(' clear ')

def restart_program():
   python = sys.executable
   os.execl(python, python, * sys.argv)
   curdir = os.getcwd()

def load_animation(): 
  
    load_str = "Loading DBROWSER."
    ls_len = len(load_str) 
    time.sleep(1)
    animation = "|/-\\|/-\|/-/"
    anicount = 0
      
    counttime = 0        
       
    i = 0                     
  
    while (counttime != 30): 
          

        time.sleep(0.075)  
                              
        load_str_list = list(load_str)  
 
        x = ord(load_str_list[i]) 
          
        y = 0                             

        if x != 32 and x != 46:              
            if x>90: 
                y = x-32
            else: 
                y = x + 32
            load_str_list[i]= chr(y) 
          
        res =''              
        for j in range(ls_len): 
            res = res + load_str_list[j] 
              
        sys.stdout.write("\r"+res + animation[anicount]) 
        sys.stdout.flush() 
  
        load_str = res 
  
          
        anicount = (anicount + 1)% 4
        i =(i + 1)% ls_len 
        counttime = counttime + 1

    else: 
        os.system("clear") 
  
# Driver program 
if __name__ == '__main__':  
    load_animation() 

def restart_program():
   python = sys.executable
   os.execl(python, python, * sys.argv)
   curdir = os.getcwd()

def screen_clear():
   if name == 'nt':
      _ = system('cls')

def CS(X):          # CS = clear sleep
   time.sleep(X)
   os.system("clear")

print(Fore.RED+"")
banner = pyfiglet.figlet_format("DBROW", font = "isometric1" )

CS(2)
print(banner)
print(Fore.RED+"                                           WelC0me to DBROW the hackers browser   ")
print("                                                                               V2.0")
print(" [1] Just browse the net ")
print(" [2] go to my index for my website ") 
print(" [3] just browse google ")
print(" [4] view supported websites ")
print(" [5] see whats new for vserions 2.0 ")
print("--------------------------------------------------------------------------------------------")
# run.py = index 
# rung = run google 
# run1 = run duckduckgo 

N = str(input(" Options ===> "))


####### BROWSE DUCKDUCKGO ###########


if '5' == N:
    CS(2)
    print(banner)
    print("-"*40)
    time.sleep(1)
    print(" fixxed web bugs and index bugs ")
    time.sleep(1)
    print(" added netsniff-ng as a packte monitor ")
    time.sleep(1)
    print(" added bash scripts for setup ")
    time.sleep(1)
    print(" added automation ")
    time.sleep(1)
    print(" added json files for easier loads and time ")
    time.sleep(1)
    print(" added newer websites that are supported for the browser itself ")
    time.sleep(1)
    print(" updated CSS ")
    time.sleep(1)
    print(" added bash script that checks for required packages ")
    time.sleep(1)
    print(" added new terminal for gnome where it opens a new term for netsniff ")
    time.sleep(1)
    print(" added termcolor ")
    time.sleep(1)
    print(" added more proxies into the browser ")
    time.sleep(1)
    print(" added a few more lines and took away input for if proxychains is installed ")
    time.sleep(1)
    print(" [!] stay tunned for further updates on the browser [!] ")
    restart_program()

if '4' == N:
    F = open('links.json','r+',encoding='utf-8') # open encoding
    data = json.load(F) #load the file
    for x in data['prints']:
        print(x) # print value x in this case the fi;e
        time.sleep(0.1)
        restart_program()

if '1' == N:
    time.sleep(1)
    print(" [!] running Dark browser [!] ")
    time.sleep(3)
    Yn = str(input(" would you like to run proxys along side Netsniff-ng Y/n? "))
    time.sleep(1)
    
    if 'n' in Yn:
        time.sleep(1)
        print(" alright then running browser ")
        CS(2)
        print(banner)
        os.system(' sudo python3 run1.py ')
        print(" [!] Stopping tor service and breaking connections [!] ")
        os.system(' sudo service tor stop && clear ')
        print(" would you like to view the cap file from netsniff? ")
        V = str(input(" Y/n: ==> "))

        if 'y' in V:
            time.sleep(1)
            os.system(' sudo wireshark pack.pcap ')
            time.sleep(1)
            print(" have a nice one :D [!] ")
            sys.exit()
        
        elif 'Y' in V:
            time.sleep(1)
            os.system(' clear' )
            os.system(' sudo wireshark pack.pcap ')
            print(" Have a ncie one ")
            sys.exit()

        if 'n' in V:
            CS(2)
            print(" have a nice one :D [!] ")
            sys.exit()
        
        elif 'N' in V:
            CS(2)
            print(" Have a ncie one ")
            sys.exit()

    if 'Y' == Yn:
        time.sleep(1)
        print(" [=] alright then running browser with proxychains and tor service ")
        CS(2)
        print(banner)
        os.system(" chmod +x ./newterm.sh && ./newterm.sh ")
        os.system(' sudo service tor start && proxychains python3 run1.py') 
        print(" [!] Stopping tor service and breaking connections [!] ")
        os.system(' sudo service tor stop && clear ')
        print(" would you like to view the cap file from netsniff? ")
        V = str(input(" Y/n: ==> "))

        if 'y' in V:
            time.sleep(1)
            os.system(' sudo wireshark pack.pcap ')
            time.sleep(1)
            print(" have a nice one :D [!] ")
            sys.exit()
        
        elif 'Y' in V:
            time.sleep(1)
            os.system(' clear' )
            os.system(' sudo wireshark pack.pcap ')
            print(" Have a ncie one ")
            sys.exit()

        if 'n' in V:
            CS(2)
            print(" have a nice one :D [!] ")
            sys.exit()
        
        elif 'N' in V:
            CS(2)
            print(" Have a ncie one ")
            sys.exit()   

##############################################
###########BROWSE GOOGLE 

elif '3' == N:
    time.sleep(1)
    Yn = str(input(" Would you like to use proxies Y/n? "))
    
    if 'Y' in Yn:
           CS(2)
           time.sleep(1)
           print(banner)
           print(" [=] alright then running browser [=] ")
           os.system(" chmod +x ./newterm.sh && ./newterm.sh ")
           os.system(' sudo service tor start && proxychains python3 rung.py  ')
           print(" [!] Stopping tor service and breaking connections [!] ")
           os.system(' sudo service tor stop && clear ')
           CS(2)
           print(" would you like to view the cap file from netsniff? ")
           V = str(input(" Y/n: ==> "))

           if 'y' in V:
               time.sleep(1)
               os.system(' sudo wireshark pack.pcap ')
               time.sleep(1)
               print(" have a nice one :D [!] ")
               sys.exit()
           
           elif 'Y' in V:
               time.sleep(1)
               os.system(' clear' )
               os.system(' sudo wireshark pack.pcap ')
               print(" Have a ncie one ")
               sys.exit()

           if 'n' in V:
               CS(2)
               print(" have a nice one :D [!] ")
               sys.exit()
           
           elif 'N' in V:
               CS(2)
               print(" Have a ncie one ")
               sys.exit()

    if 'n' in Yn:
           CS(2)
           time.sleep(1)
           print(banner)    
           os.system(' python3 rung.py ')
           os.system(' clear ')
           print(" [!] Stopping tor service and breaking connections [!] ")
           os.system(' sudo service tor stop && clear ')
           sys.exit()

###################BROWSE MTY INDEX################ 
elif '2' == N:
      time.sleep(1)
      Yn = str(input(" Would you like to use proxies Y/n? "))
      print(" [!] running Dark browser [!] ")
      
      if 'Y' in Yn:
          CS(2)
          print(banner)
          print(" [=] alright then running script [=] ")
          time.sleep(1)
          os.system(" chmod +x ./newterm.sh && ./newterm.sh ")
          os.system(' sudo service tor start && proxychains python3 run.py ')
          print(" [!] Stopping tor service and breaking connections [!] ")
          os.system(' sudo service tor stop && clear ')
          CS(2)
          print(" would you like to view the cap file from netsniff? ")
          V = str(input(" You: ==> "))

          if 'y' in V:
              time.sleep(1)
              os.system(' sudo wireshark pack.pcap ')
              time.sleep(1)
              print(" have a nice one :D [!] ")
              sys.exit()
          
          elif 'Y' in V:
              time.sleep(1)
              os.system(' sudo wireshark pack.pcap ')
              time.sleep(1)
              print(" Thanks for stopping by :D [+] ")
              sys.exit()


          elif 'n' in V:
              time.sleep(1)
              print(" have a nice one :D [+] ")
              sys.exit()

      if 'n' == Yn:
          CS(2)
          print(" [=] alright then running script [=] ")
          print(banner)
          time.sleep(1)
          os.system(' python3 run.py ')
          print(" [!] Stopping tor service and breaking connections [!] ") 

else:
    print(" [!] that doesnt seem to be a command ")
    restart_program()          
