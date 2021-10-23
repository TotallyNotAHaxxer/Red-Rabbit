import os 
import sys
import time as t 
import datetime 
from datetime import datetime 
import socket
import requests
from requests import get 
import colorama 
import json 
import random 
import platform
import scapy.all as scapy
from colorama import init 
from colorama import Fore, Back, Style 
import re
import speedtest

init() 

ip = get('https://api.ipify.org').text

uname = platform.uname()

def CS(X):
    t.sleep(X)
    os.system('clear')


def scan1():
    ip_add_range_pattern = re.compile("^(?:[0-9]{1,3}\.){3}[0-9]{1,3}/[0-9]*$")
    while True:
        print(Fore.RED+"")
        t.sleep(2)
        print("----------------------------------------------------------------")
        t.sleep(0.1)
        print(" MAKE SURE ITS A RANGE (ex 192.168.1.0/24) ")
        t.sleep(0.1)
        print(" MAKE SURE YOU RAN THIS PROGRAM AS ROOT ")
        t.sleep(0.1)
        print(" MAKE SURE YOU ARE ON THE CURRENT NETWORK ")
        t.sleep(0.1)
        print(" MAKE SURE YOU HAVE PREMISSION TO DO THIS ")
        t.sleep(0.1)
        print("----------------------------------------------------------------")
        ip_add_range_entered = input("\nIPA to send ARP to ==> ")
        if ip_add_range_pattern.search(ip_add_range_entered):
            print(f"{ip_add_range_entered} is a valid IP range")
            break
    print(Fore.GREEN+"")
    arp_result = scapy.arping(ip_add_range_entered)

def banner_1():
    print(Fore.RED+"""
▓█████▄  ██▓  ██████  ▄████▄   ▒█████   ██▀███  ▓█████▄      ██████  ███▄    █  ██▓ ██▓███  ▓█████  ██▀███
▒██▀ ██▌▓██▒▒██    ▒ ▒██▀ ▀█  ▒██▒  ██▒▓██ ▒ ██▒▒██▀ ██▌   ▒██    ▒  ██ ▀█   █ ▓██▒▓██░  ██▒▓█   ▀ ▓██ ▒ ██▒
░██   █▌▒██▒░ ▓██▄   ▒▓█    ▄ ▒██░  ██▒▓██ ░▄█ ▒░██   █▌   ░ ▓██▄   ▓██  ▀█ ██▒▒██▒▓██░ ██▓▒▒███   ▓██ ░▄█ ▒
░▓█▄   ▌░██░  ▒   ██▒▒▓▓▄ ▄██▒▒██   ██░▒██▀▀█▄  ░▓█▄   ▌     ▒   ██▒▓██▒  ▐▌██▒░██░▒██▄█▓▒ ▒▒▓█  ▄ ▒██▀▀█▄
░▒████▓ ░██░▒██████▒▒▒ ▓███▀ ░░ ████▓▒░░██▓ ▒██▒░▒████▓    ▒██████▒▒▒██░   ▓██░░██░▒██▒ ░  ░░▒████▒░██▓ ▒██▒
 ▒▒▓  ▒ ░▓  ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░ ▒▒▓  ▒    ▒ ▒▓▒ ▒ ░░ ▒░   ▒ ▒ ░▓  ▒▓▒░ ░  ░░░ ▒░ ░░ ▒▓ ░▒▓░
 ░ ▒  ▒  ▒ ░░ ░▒  ░ ░  ░  ▒     ░ ▒ ▒░   ░▒ ░ ▒░ ░ ▒  ▒    ░ ░▒  ░ ░░ ░░   ░ ▒░ ▒ ░░▒ ░      ░ ░  ░  ░▒ ░ ▒░
 ░ ░  ░  ▒ ░░  ░  ░  ░        ░ ░ ░ ▒    ░░   ░  ░ ░  ░    ░  ░  ░     ░   ░ ░  ▒ ░░░          ░     ░░   ░
   ░     ░        ░  ░ ░          ░ ░     ░        ░             ░           ░  ░              ░  ░   ░
 ░                   ░                           ░
    """)
    print("""                                                                                         Version 1.0
                                                                                              Authors ~~> Scare Sec Hackers""")

def menu1():
    print(Fore.GREEN+"""
▓█████▄  ██▓  ██████  ▄████▄   ▒█████   ██▀███  ▓█████▄      ██████  ███▄    █  ██▓ ██▓███  ▓█████  ██▀███
▒██▀ ██▌▓██▒▒██    ▒ ▒██▀ ▀█  ▒██▒  ██▒▓██ ▒ ██▒▒██▀ ██▌   ▒██    ▒  ██ ▀█   █ ▓██▒▓██░  ██▒▓█   ▀ ▓██ ▒ ██▒
░██   █▌▒██▒░ ▓██▄   ▒▓█    ▄ ▒██░  ██▒▓██ ░▄█ ▒░██   █▌   ░ ▓██▄   ▓██  ▀█ ██▒▒██▒▓██░ ██▓▒▒███   ▓██ ░▄█ ▒
░▓█▄   ▌░██░  ▒   ██▒▒▓▓▄ ▄██▒▒██   ██░▒██▀▀█▄  ░▓█▄   ▌     ▒   ██▒▓██▒  ▐▌██▒░██░▒██▄█▓▒ ▒▒▓█  ▄ ▒██▀▀█▄
░▒████▓ ░██░▒██████▒▒▒ ▓███▀ ░░ ████▓▒░░██▓ ▒██▒░▒████▓    ▒██████▒▒▒██░   ▓██░░██░▒██▒ ░  ░░▒████▒░██▓ ▒██▒
 ▒▒▓  ▒ ░▓  ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░ ▒▒▓  ▒    ▒ ▒▓▒ ▒ ░░ ▒░   ▒ ▒ ░▓  ▒▓▒░ ░  ░░░ ▒░ ░░ ▒▓ ░▒▓░
 ░ ▒  ▒  ▒ ░░ ░▒  ░ ░  ░  ▒     ░ ▒ ▒░   ░▒ ░ ▒░ ░ ▒  ▒    ░ ░▒  ░ ░░ ░░   ░ ▒░ ▒ ░░▒ ░      ░ ░  ░  ░▒ ░ ▒░
 ░ ░  ░  ▒ ░░  ░  ░  ░        ░ ░ ░ ▒    ░░   ░  ░ ░  ░    ░  ░  ░     ░   ░ ░  ▒ ░░░          ░     ░░   ░
   ░     ░        ░  ░ ░          ░ ░     ░        ░             ░           ░  ░              ░  ░   ░
 ░                   ░                           ░
    """)
    print(f"                                                                                     Welcome to SnI4Er {uname.node}")
    print("""                                                                                         Version 1.0
                                                                                              Authors ~~> Scare Sec Hackers""")
    t.sleep(1)
    print("[-] Checking Connections To Discord....")
    t.sleep(1)
    print("[-] Using public IPA -> {}".format(ip))
    r = requests.get('https://discord.com')
    t.sleep(2)
    if r.status_code == 200:
        print ('Looks all good! Sending you to the menu!')
    else:
        print("o(╥﹏╥)o ~ i failed you...")
        print ('[!] Connection refused Try Again later......')
        sys.exit()

######################################################################################################################################################################################################################################################
################################################################################## DEFINITIONS FOR ATTACK MODULES PLACED ABOVE#############################################################################################
# load animation 
def load():
    def loading():
        print(Fore.MAGENTA+"")
        for _ in tqdm(range(100), desc="Sending...", ascii=False, ncols=75):
            time.sleep(0.01)

def read():
    CS(1)
    banner_1()
    print(Fore.MAGENTA+"────────────────────────────────────────────────────────────────────────────────────────────────────────")
    print(f" Running   Node ~~> {uname.node}    ""     Running Public IPA => {}".format(ip),)
    print(f" USERTOKEN AUTH ~~> {AUTH}          ")
    print(Fore.RED+"                                           ┌∩┐(◣ _ ◢)┌∩┐ ")
    print("""
    LIABILITY ~~> ME OR ANY OTHER PERSON WHO OBTAINS A COPY OF THIS FRAMEWORK IS HELD LIABLE FOR YOUR CUMBASS DECISIONS 

    WARNING   ~~> NOTE: THIS SCRIPT USES REQUESTS TO SEND THE POST, GET, DELETE AND OTHER FORMS OF REQUESTS TO DO THE THINGS THIS SCRIPT HAS 
    IT IS SUGGESTED YOU DO NOT OVER USE OR ABUSE THIS SCRIPT AS ONCE THIS SCRIPT IS USED IT CAN SEND UP TO 500-1000 REQUESTS IN UNDER A FEW 
    SECONDS AND CAN RESULT IN A DoS IF YOU ARE NOT CAREFUL

    ACCOUNT   ~~> Due to the ammount of requests this script sends your account may get locked, and you may have to change your password 
    this has nothing to do with the actuall script its just how discords servers respond to the amount of requests, if they see your sending 
    way to much to be a human then they will see your account was hijacked or self botted ( against DISCORD TOS ) this can for real ban you account 
    dont use it to much at once or in one day for the matter, and use it right 

    This Script was written by the Scare-Sec-Hackers Wordlwide, i hope you enjoy this script as it has taken a while to create 
    just dont be dumb   
    """)

def scanner():
    import psutil 
    from tabulate import tabulate

    class InterfaceScanner(object):
        def __init__(self):
            self.instance = psutil.net_if_addrs()
        
        def scanner(self):
            self.interfaces = []
            self.address_ip = []
            self.netmask_ip = []
            self.broadcast_ip = []
            for interface_name, interface_addresses in self.instance.items():
                self.interfaces.append(interface_name)
                for address in interface_addresses:
                    if str(address.family) == 'AddressFamily.AF_INET':
                        self.address_ip.append(address.address)
                        self.netmask_ip.append(address.netmask)
                        self.broadcast_ip.append(address.broadcast)
                data = {"interface" : [*self.interfaces],
                        "IP-Address": [*self.address_ip],
                        "Netmask"   : [*self.netmask_ip],
                        "Broadcast-IP" :[*self.netmask_ip]
                        }
                return tabulate(data, headers="keys", tablefmt="github")
        
        def __str__(self):
            return str(self.scanner())

    if __name__ == "__main__":
        print(InterfaceScanner())

def SPAM_SERVER():
    print(" [!] Loading SPAM Script ")
    t.sleep(1)
    CS(1)
    print(Fore.RED+"")
    banner_1()
    run = str(input(" Message to Spam >>> "))
    print(" -----------------------------------------------------------------------")
    print(" EX ~>  https://discord.com/api/v9/channels/855499460511793196/messages ")
    print("------------------------------------------------------------------------")
    hedscript = str(input("Request URL ID   >>> "))
    header = {
    'authorization': f'{AUTH}'
    }
    payload = {
        'content': f"{run}"  
    }
    CS(1)
    banner_1()
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Sending Message => {run}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Using URL       => {hedscript}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] With AUTH token => {AUTH}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print("\033[35m[\033[36m+\033[35m] [INFO] WITH IPA         => {}".format(ip))
    while True:
        r = requests.post(f"{hedscript}", data=payload, headers=header) # request method 
        t.sleep(1)
        print("\033[37m Shooting MY SHOTSSSS (҂‾ ▵‾)︻デ═一 (˚▽˚’!)/")
        t.sleep(0.1)
        print("\033[31m  Snipe made at => " + str(datetime.now()))
        t.sleep(0.1)
        print(f"\033[31m [DATA] ===> HEADER USED ===> {hedscript} WITH CONNECTION ""{}".format(ip), f"WITH SYSTEM KEY ==> {AUTH} YOU ARE VISIBLE" )


def spam_gifs():
    CS(1)
    banner_1()
    print(" [!] Loading Message Script ")
    CS(1)
    banner_1()
    print("----------------------------------------------")
    print("EXAMPLE => 855499460511793196")
    print("----------------------------------------------")
    ID = str(input("User ID >>> "))
    list = ["https://tenor.com/view/oh-hey-fuck-you-fuck-off-hi-hey-gif-14740754", "https://tenor.com/view/the-cuckshed-dynamis-cr8zgmr-cuck-cuckshed-gif-19781415", "https://tenor.com/view/homo-congratulations-robin-hood-rainbow-gif-8127747", "https://tenor.com/view/todd-chrisley-middle-finger-fuck-you-flip-off-screw-you-gif-10663162", "https://tenor.com/view/flipping-off-flip-off-middle-finger-smile-happy-gif-4746862", "https://tenor.com/view/wtf-haha-flirty-fuck-smile-gif-15931510", "https://tenor.com/view/fuck-you-fuck-up-fucked-off-i-dont-give-a-fuck-idgaf-gif-15388049", "heyhttps://tenor.com/view/rosycheeks-mochi-peach-mochi-cat-cute-kitty-peach-cat-gif-16992614", "https://tenor.com/view/fuck-fuck-you-middle-finger-middle-finger-gif-15294280", "https://tenor.com/view/crazy-die-on-fire-elmo-flames-gif-17266822", "https://tenor.com/view/die-gif-8468069", "https://tenor.com/view/kick-gif-7905607"]
    # make them feel disgusting inside 

    item = random.choice(list)
    header = {
    'authorization': f'{AUTH}'
    }
    payload = {
        'content': "" + item 
    }
    CS(1)
    banner_1()
    c = socket.gethostbyname('www.discord.com')
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Sending Message => {list}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Using USER ID   => {ID}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] With AUTH token => {AUTH}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print("\033[35m[\033[36m+\033[35m] [INFO] WITH IPA         => {}".format(ip))
    while True:
        r = requests.post(f"https://discord.com/api/v9/channels/{ID}/messages", data=payload, headers=header) # throw under a while true to continue the request fetching up to 15 requests a second 
        t.sleep(0.1)
        print("\033[37m Shooting MY SHOTSSSS (҂‾ ▵‾)︻デ═一 (˚▽˚’!)/")                                       #{}".format(ip)
        print(f"\033[31mMessage ~~> {list} \033[36mSent To Server ~~> ", c, "\033[31mUSING PUBLIC IPA ~~> {}".format(ip))



def spam_LINKS():
    CS(1)
    banner_1()
    print(" [!] Loading Message Script ")
    CS(1)
    banner_1()
    print("----------------------------------------------")
    print("EXAMPLE => 855499460511793196")
    print("----------------------------------------------")
    ID = str(input("User ID >>> "))
    list = ["https://tenor.com/view/hug-virtual-hug-hug-sent-gif-5026057", "https://www.pornhub.com/view_video.php?viewkey=ph61112e0fe8f0c",
    "https://www.pornhub.com/view_video.php?viewkey=ph5fa59b5e7f78d", "https://www.pornhub.com/view_video.php?viewkey=ph5f427d2740fe5", 
    "https://www.pornhub.com/view_video.php?viewkey=ph5e33944b200f6", "https://www.pornhub.com/view_video.php?viewkey=ph605e9617a7f1e", 
    "https://www.pornhub.com/view_video.php?viewkey=ph5e18e0d6472d4", "https://www.pornhub.com/view_video.php?viewkey=ph611064c6c573d",
    "https://www.pornhub.com/view_video.php?viewkey=ph60bb6710bc296","https://www.pornhub.com/view_video.php?viewkey=ph5dc61b13a6bd6",
    "https://www.pornhub.com/view_video.php?viewkey=ph5f7ca48d800ba","https://www.pornhub.com/view_video.php?viewkey=ph6047495a560d0",
    "https://www.pornhub.com/view_video.php?viewkey=ph5e25abf7ba478","https://www.pornhub.com/view_video.php?viewkey=ph5a3034097e5f1",
    "https://www.pornhub.com/view_video.php?viewkey=ph603e5caddf123"]
    # make them feel disgusting inside 

    item = random.choice(list)
    header = {
    'authorization': f'{AUTH}'
    }
    payload = {
        'content': "" + item 
    }
    CS(1)
    banner_1()
    c = socket.gethostbyname('www.discord.com')
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Sending Message => {list}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Using USER ID   => {ID}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] With AUTH token => {AUTH}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print("\033[35m[\033[36m+\033[35m] [INFO] WITH IPA         => ")
    while True:
        r = requests.post(f"https://discord.com/api/v9/channels/{ID}/messages", data=payload, headers=header) # throw under a while true to continue the request fetching up to 15 requests a second 
        t.sleep(0.1)
        print("\033[37m Shooting MY SHOTSSSS (҂‾ ▵‾)︻デ═一 (˚▽˚’!)/")                                      #{}".format(ip)
        print(f"\033[31mMessage ~~> {list} \033[36mSent To Server ~~> ", c, "\033[31mUSING PUBLIC IPA ~~> ")

def sniff_1_ID():
    CS(2)
    print(" [!] Osint section #2 ")
    t.sleep(1)
    banner_1()
    osint_stage_1()

def send_message_1(): # just defualt message without spam 
    CS(2)
    print(" [!] Loading Message Script ")
    t.sleep(1)
    print("""
    YOU HAVE 30 SECONDS TO ABORT!!! ME OR MY TEAM ARE NOT HELD LIABLE FOR ANY ILLEGAL ACTIVITIES
    YOU COMMIT WITH THIS TOOL, NOTE THIS TOOL MAKES ALOT OF REQUESTS TO THE DISCORD API AND SERVER
    THIS CAN RESULT IN YOU NEEDING TO RESET YOUR PASSWORD THIS IS BECAUSE OF THE ANTI BOT SERVICE 
    DISCORD USES TO PROTECT THEIR SERVERS AND THEIR CLIENTS !!!!!
    STARTING IN 30 SECONDS PLEASE WAIT !!!!!
    """)
    CS(1)
    print(Fore.RED+"")
    banner_1()
    run = str(input(" Message to send >>> "))
    print(" -------------------------------------------- ")
    print(" EX ~>  https://discord.com/api/v9/channels/855499460511793196/messages ")
    hedscript = str(input("Request URL ID   >>> "))
    header = {
    'authorization': f'{AUTH}'
    }
    payload = {
        'content': f"{run}"  
    }
    CS(1)
    banner_1()
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Sending Message => {run}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Using URL       => {hedscript}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] With AUTH token => {AUTH}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print("\033[35m[\033[36m+\033[35m] [INFO] WITH IPA         => {}".format(ip))
    r = requests.post(f"{hedscript}", data=payload, headers=header) # request method 
    
def spam_message_1():
    print(" [!] Loading Message Script ")
    CS(1)
    banner_1()
    run = str(input(" Message to send >>> "))
    print(" -------------------------------------------- ")
    print(" EX ~>  https://discord.com/api/v9/channels/855499460511793196/messages ")
    hedscript = str(input("Request URL ID   >>> "))
    header = {
    'authorization': f'{AUTH}'
    }
    payload = {
        'content': f"{run}"  
    }
    CS(1)
    banner_1()
    c = socket.gethostbyname('www.discord.com')
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Sending Message => {run}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] Using URL       => {hedscript}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print(f"\033[35m[\033[36m+\033[35m] [INFO] With AUTH token => {AUTH}")
    t.sleep(1)
    print(Fore.RED+"==============================================================================================")
    print("\033[35m[\033[36m+\033[35m] [INFO] WITH IPA         => {}".format(ip))
    while True:
        #r = requests.post(f"{hedscript}", data=payload, headers=header) # throw under a while true to continue the request fetching up to 15 requests a second 
        t.sleep(0.1)
        print(f"\033[31mMessage ~~> {run} \033[36mSent To Server ~~> ", c, "\033[31mUSING PUBLIC IPA ~~> {}".format(ip))


def osint_stage_1():
    CS(2)
    banner_1()
    print("[+] Running Osint Module! ")
    t.sleep(1)
    print(" ------------------------------------------------------------------------------------------- ")
    print("""
    YOU HAVE 30 SECONDS TO ABORT!!! ME OR MY TEAM ARE NOT HELD LIABLE FOR ANY ILLEGAL ACTIVITIES
    YOU COMMIT WITH THIS TOOL, NOTE THIS TOOL MAKES ALOT OF REQUESTS TO THE DISCORD API AND SERVER
    THIS CAN RESULT IN YOU NEEDING TO RESET YOUR PASSWORD THIS IS BECAUSE OF THE ANTI BOT SERVICE 
    DISCORD USES TO PROTECT THEIR SERVERS AND THEIR CLIENTS !!!!!
    STARTING IN 30 SECONDS PLEASE WAIT !!!!!
    """)
    t.sleep(3) #swicth to 30 seconds 
    #855499460511793196
    CHID = str(input(" Channel ID ==> "))
    def retrieve_messages(channelid):
        headers = {
            'authorization': f'{AUTH}'
        }
        
        r = requests.get(f'https://discord.com/api/v9/channels/{channelid}/messages', headers=headers)
        jsonn = json.loads(r.text)
        for value in jsonn:
            t.sleep(1)
            print(Fore.RED+"")
            print(Fore.RED+"[INFO] ===> " + str(datetime.now()))
            t.sleep(1)
            print(Fore.BLUE+"")
            print(value, '\n')
            jsonFile = open("data.json", "w")
            jsonFile.write(f"{r.text}\n")
            jsonFile.close()
    retrieve_messages(f'{CHID}')
    print(" Data Saved to a json file data.json")



#############################################################################################################################################################################################################################

def menu2():
    CS(3)
    import psutil
    A = 'https://www.discord.com'
    uagent = 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36'
    addrs = psutil.net_if_addrs()
    print(Fore.GREEN+"""
▓█████▄  ██▓  ██████  ▄████▄   ▒█████   ██▀███  ▓█████▄      ██████  ███▄    █  ██▓ ██▓███  ▓█████  ██▀███
▒██▀ ██▌▓██▒▒██    ▒ ▒██▀ ▀█  ▒██▒  ██▒▓██ ▒ ██▒▒██▀ ██▌   ▒██    ▒  ██ ▀█   █ ▓██▒▓██░  ██▒▓█   ▀ ▓██ ▒ ██▒
░██   █▌▒██▒░ ▓██▄   ▒▓█    ▄ ▒██░  ██▒▓██ ░▄█ ▒░██   █▌   ░ ▓██▄   ▓██  ▀█ ██▒▒██▒▓██░ ██▓▒▒███   ▓██ ░▄█ ▒
░▓█▄   ▌░██░  ▒   ██▒▒▓▓▄ ▄██▒▒██   ██░▒██▀▀█▄  ░▓█▄   ▌     ▒   ██▒▓██▒  ▐▌██▒░██░▒██▄█▓▒ ▒▒▓█  ▄ ▒██▀▀█▄
░▒████▓ ░██░▒██████▒▒▒ ▓███▀ ░░ ████▓▒░░██▓ ▒██▒░▒████▓    ▒██████▒▒▒██░   ▓██░░██░▒██▒ ░  ░░▒████▒░██▓ ▒██▒
 ▒▒▓  ▒ ░▓  ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░ ▒▒▓  ▒    ▒ ▒▓▒ ▒ ░░ ▒░   ▒ ▒ ░▓  ▒▓▒░ ░  ░░░ ▒░ ░░ ▒▓ ░▒▓░
 ░ ▒  ▒  ▒ ░░ ░▒  ░ ░  ░  ▒     ░ ▒ ▒░   ░▒ ░ ▒░ ░ ▒  ▒    ░ ░▒  ░ ░░ ░░   ░ ▒░ ▒ ░░▒ ░      ░ ░  ░  ░▒ ░ ▒░
 ░ ░  ░  ▒ ░░  ░  ░  ░        ░ ░ ░ ▒    ░░   ░  ░ ░  ░    ░  ░  ░     ░   ░ ░  ▒ ░░░          ░     ░░   ░
   ░     ░        ░  ░ ░          ░ ░     ░        ░             ░           ░  ░              ░  ░   ░
 ░                   ░                           ░
    """)
    print("""                                                                                         Version 1.0
                                                                                              Authors ~~> Scare Sec Hackers""")
    print(Fore.MAGENTA+"────────────────────────────────────────────────────────────────────────────────────────────────────────")
    print(f" Running   Node ~~> {uname.node}    ""     Running Public IPA => {}".format(ip),)
    print(f" USERTOKEN AUTH ~~> {AUTH}          ")
    print(" Running WEBURL ~~> ",A) 
    print(Fore.RED+"                                           ┌∩┐(◣ _ ◢)┌∩┐ ")
    print(Fore.MAGENTA+"──────────────────────────────────────────────────────────────────────────────────────────────────────────|")
    print("| \033[35m[\033[36m1\033[35m] Snipe Codes        | \033[35m[\033[36m2\033[35m] Run OSINT on a Server      | \033[35m[\033[36m3\033[35m] Send a message                            |")
    print("| \033[35m[\033[36m4\033[35m] Sniff ID's         | \033[35m[\033[36m5\033[35m] Spam A Server              | \033[35m[\033[36m6\033[35m] Spam Messages                             |")
    print("| \033[35m[\033[36m7\033[35m] Update SYS!        | \033[35m[\033[36m8\033[35m] Run OSINT on a Person      | \033[35m[\033[36m9\033[35m]                                           |")
    print("| \033[35m[\033[36m10\033[35m]Spam GIFS          | \033[35m[\033[36m11\033[35m] Spam Undescribable links  | \033[35m[\033[36m12\033[35m] Read About ME!                           |")
    print(Fore.MAGENTA+"─────────────────────────────────────────────NETWORKING AND SYS───────────────────────────────────────────|")
    print("| \033[35m[\033[36mA\033[35m] Check Response     | \033[35m[\033[36mC\033[35m] Check IPA                  | \033[35m[\033[36mE\033[35m] Check Internet Speeds                     |")
    print("| \033[35m[\033[36mB\033[35m] Get response code  | \033[35m[\033[36mD\033[35m] Scan Your Network          | \033[35m[\033[36mF\033[35m] Check Interfaces                          |")
    print(Fore.MAGENTA+"──────────────────────────────────────────────────────────────────────────────────────────────────────────|")
    consol()


def consol():
    t.sleep(1)
    SNOOP = str(input(">> "))

    if '1' == SNOOP:
        os.system('python3 dis-snipe.py ')

    elif '2' == SNOOP:
        CS(2)
        banner_1()
        osint_stage_1()
    
    elif '3' == SNOOP:
        CS(2)
        banner_1()
        send_message_1()

    elif '4' == SNOOP:
        CS(2)
        banner_1()
        sniff_1_ID()
    
    elif '5' == SNOOP:
        CS(2)
        SPAM_SERVER()
    
    elif '6' == SNOOP:
        spam_message_1()
    
    elif '7' == SNOOP:
        os.system('sudo apt-get update')
        CS(2)
        AUTH = str(input(" Your Discord Authorization KEY => "))
        menu2()
    
    elif '11' == SNOOP:
        spam_LINKS()
    
    elif '12' == SNOOP:
        read()

    elif '8' == SNOOP:
        OSINT_PERSON_1()

    elif 'A' == SNOOP:
        r = requests.get('https://google.com')
        print(Fore.GREEN+"[+] Checking Connection with => {}".format(ip))
        if r.status_code == 200:
            print(Fore.GREEN+'Connection Seems ʕ •ᴥ•ʔ')
            CS(4)
            AUTH = str(input(" Your Discord Authorization KEY => "))
            menu2()
        else:
            print(Fore.RED+"o(╥﹏╥)o ~ i failed you...")
            print(Fore.RED+'[!] Connection refused Try Again later......')
            sys.exit()
    
    elif 'B' == SNOOP:
        r = requests.get('https://discord.com')
        print("[STATUS CODE] ~~>",r.status_code)
        t.sleep(1)
        CS(3)
        AUTH = str(input(" Your Discord Authorization KEY => "))
        menu2()
    
    elif 'C' == SNOOP:
        print(Fore.GREEN+"[DATA] YOUR PUBLIC IPA IS => {}".format(ip))
        AUTH = str(input(" Your Discord Authorization KEY => "))
        menu2()
    
    elif 'D' == SNOOP:
        print("SCANNING NETWORK")
        CS(3)
        banner_1()
        scan1()
    
    elif 'E' == SNOOP:
        st = speedtest.Speedtest()
        print('                    | Download Speed   => ',end='')
        print(st.download())
        print('                    | UpLoad Speed     => ', end='')
        print(st.upload())
        servernames = []
        st.get_servers(servernames)
    
    elif 'F' == SNOOP:
        scanner()

    else:
        print(" Not a command it seems!")
        CS(2)
        menu2()



if __name__ == "__main__":
    CS(1)
    banner_1()
    AUTH = str(input(" Your Discord Authorization KEY => "))
    CS(1)
    menu1()
    CS(1)
    menu2()
