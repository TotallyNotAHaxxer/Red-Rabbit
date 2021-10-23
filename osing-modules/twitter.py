import twint
import os 
import sys 
import time
import colorama 
from colorama import Fore 

def main():
    os.system(' clear ')
    time.sleep(1)
    print(" [+] Loading script [+] ")
    time.sleep(1)
    os.system(' clear ')
    time.sleep(0.1)
    print(Fore.MAGENTA+"             ___________       .__.__  __ ")                                                             
    time.sleep(0.1)
    print(Fore.MAGENTA+"             \__    ___/_  _  _|__|__|/  |_  ___________        ")                   
    time.sleep(0.1)
    print(Fore.RED+"               |    |  \ \/ \/ /  |  \   __\/ __ \_  __ \ ")                         
    time.sleep(0.1)
    print(Fore.MAGENTA+"               |    |   \     /|  |  ||  | \  ___/|  | \/   ")                       
    time.sleep(0.1)
    print(Fore.RED+"               |____|    \/\_/ |__|__||__|  \___  >__|        ")                     
    time.sleep(0.1)
    print(Fore.MAGENTA+"            .___        __         .__  .__  .__\/              ")                   
    time.sleep(0.1)
    print(Fore.RED+"            |   | _____/  |_  ____ |  | |  | |__| ____   ____   ____   ____  ____  ")
    time.sleep(0.1)
    print(Fore.MAGENTA+"            |   |/    \   __\/ __ \|  | |  | |  |/ ___\_/ __ \ /    \_/ ___\/ __ \ ")
    time.sleep(0.1)
    print(Fore.RED+"            |   |   |  \  | \  ___/|  |_|  |_|  / /_/  >  ___/|   |  \  \__\  ___/ ")
    time.sleep(0.1)
    print(Fore.MAGENTA+"            |___|___|  /__|  \___  >____/____/__\___  / \___  >___|  /\___  >___  > ")
    time.sleep(0.1)
    print(Fore.RED+"                     \/          \/            /_____/      \/     \/     \/    \/ ")
    time.sleep(0.1)
    print(Fore.MAGENTA+"       /.\                          ")
    time.sleep(0.1)
    print(Fore.RED+"          \                  ")
    time.sleep(0.1)
    print(Fore.MAGENTA+"      /                    ")
    time.sleep(0.1)
    print(Fore.RED+"     //  /                  ")
    time.sleep(0.1)
    print(Fore.MAGENTA+"     |/ /\_==================")
    time.sleep(0.1)
    print(Fore.RED+"     / /            ")
    time.sleep(0.1)
    print(Fore.MAGENTA+"    / /     ")
    time.sleep(0.1)
    print(Fore.RED+"    \/ ")
    time.sleep(1)
    print(" Usage: input a username and scrape tweets                  | EX: JoeBiden ")
    A = str(input(" @> "))
    os.system(' clear ')
    print("______________________________")
    print("|Would you like the user ID's|")
    print("|To be saved to a output file|")
    print("|____________________________|")
    print(" True or False? ")
    B = str(input(" @> "))
    os.system(' clear ')
    print("|-----------------------|")
    print("|name of file for output|")
    print("|-----------------------|")
    file = str(input(" @> "))
    os.system(' clear ')
    print(Fore.CYAN+" XD ")
    os.system(' clear ')
    print(" ______________________________________________")
    print(" | How much tweets would you like to limit to | ")
    print(" | limit = 1-3200 Tweets Per Search           | ")
    print(" |--------------------------------------------|")
    num = str(input(" @> "))
    os.system(' clear ')
    print(" [=] Scraping tweets [=] ")
    os.system(' clear ')
    #########MAIN CODE############
    c = twint.Config()
    c.Username = f"{A}" #formating for string 
    c.Custom["tweet"] = ["id"]
    c.Custom["user"] = ["bio"] 
    c.Limit = f"{num}"
    c.Store_csv = f"{B}" 
    c.Output = f"{file}"
    twint.run.Search(c)

if __name__ == "__main__":
    main()