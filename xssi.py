import requests
from pprint import pprint
from bs4 import BeautifulSoup as bs
from urllib.parse import urljoin
import colorama 
import time as t 
from colorama import Fore, Back, Style, init 
import datetime
import sys 
from datetime import datetime

init()

def xss():

    def get_all_forms(url):
        soup = bs(requests.get(url).content, "html.parser")
        return soup.find_all("form")


    def get_form_details(form):
        details = {}
        action = form.attrs.get("action").lower()
        method = form.attrs.get("method", "get").lower()
        inputs = []
        for input_tag in form.find_all("input"):
            input_type = input_tag.attrs.get("type", "text")
            input_name = input_tag.attrs.get("name")
            inputs.append({"type": input_type, "name": input_name})
        details["action"] = action
        details["method"] = method
        details["inputs"] = inputs
        return details


    def submit_form(form_details, url, value):
        target_url = urljoin(url, form_details["action"])
        inputs = form_details["inputs"]
        data = {}
        for input in inputs:
            if input["type"] == "text" or input["type"] == "search":
                input["value"] = value
            input_name = input.get("name")
            input_value = input.get("value")
            if input_name and input_value:
                data[input_name] = input_value

        if form_details["method"] == "post":
            return requests.post(target_url, data=data)
        else:
            return requests.get(target_url, params=data)


    def scan_xss(url):
        print("\x1b[H\x1b[2J\x1b[3J")
        sc = "defualtxss.txt"
        print(Fore.RED+"----------------------------------------------")
        print(Fore.RED+f"\033[35m[\033[36m+\033[35m] Utilizing Defualt xss script -> {sc}")
        print(Fore.RED+f"\033[35m[\033[36m+\033[35m] Targeting URL -> {url}")
        print(Fore.RED+f"\033[35m[\033[36m+\033[35m] Time Started  -> " + str(datetime.now()))
        t.sleep(2)
        file2 = open('defualtxss.txt', 'r')
        l1    = file2.readlines()
        count = -0
        for line in l1:
            print(Fore.YELLOW+"[-] Payload {} -> {}".format(count, line.strip()))
        print("=============== Are you sure you want to continue ==================== ")
        y = str(input("\033[31m Y/n >>> "))
        if y == 'Y':
            print("[+] Script started at -> " + str(datetime.now()))
        elif y == 'n':
            print("[~] Exiting....")
            sys.exit()
        else:
            print("[-] Not an option...")
            sys.exit()
        while True:
            file = open('defualtxss.txt', 'r')
            l = file.readlines()
            count =+ 1
            for line in l:
                    print("[~] Testing Payload -> {} : {}".format(count, line.strip()))
                    forms = get_all_forms(url)
                    print(f"[+] Detected {len(forms)} forms on {url}.")
                    js_script = "{}".format(line.strip())
                    is_vulnerable = False
                    for form in forms:
                        form_details = get_form_details(form)
                        content = submit_form(form_details, url, js_script).content.decode()
                        if js_script in content:
                            print(f"[+] XSS Detected on {url}") # <- place where it was detected 
                            print(f"[*] Form details:")
                            pprint(form_details)
                            is_vulnerable = True
            return is_vulnerable


    if __name__ == "__main__":
        print("\x1b[H\x1b[2J\x1b[3J")
        #https://xss-game.appspot.com/level1/frame
        url = str(input(" URL >>> "))
        print(scan_xss(url))

if __name__ == "__main__":
    xss()