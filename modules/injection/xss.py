import os, sys, datetime, requests, argparse, threading, socket, re
import time as t
import colorama 
#from imports 
from datetime import datetime 
from colorama import Fore, Back, Style, init 
from bs4 import BeautifulSoup as bs
from urllib.parse import urljoin
from urllib.parse import urlparse
from pprint import pprint
from prettytable import PrettyTable
#color
init()
def reline():
    with open(f"{payload}", 'r') as fp:
        ln = len(fp.readlines())
        print("\033[31mDetected Payloads in file -> ", ln)
def parsesocket():
    rel = f"{url}"
    rels = re.search(r'^http[s]*:\/\/[\w\.]*', rel).group()
    l = f"{url}"
    lu = l.split("//")[-1].split("/")[0].split('?')[0]
    print("\033[35m[*] Socket Form -> " + lu)
    sstr = socket.gethostbyname(f"{lu}")
    print("\033[35m[*] Socket Name -> " + sstr)
    print("\033[35m[*] Regexed URL -> " + rels)
def system(X):
    os.system(X)
def gethttp():
    command=f"go run http.go -u {url}"
    parsesocket()
    print("\033[35m[*] Started At  : " + str(datetime.now()))
    print("\033[35m────────────────────────────────────────────")
    print("\033[49m\033[31m[*] Gathering X-Frame request headers......")
    system(command)
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
    sc = f"{payload}"
    print(Fore.RED+"----------------------------------------------")
    print(Fore.RED+f"\033[35m[\033[36m+\033[35m] Utilizing Defualt xss script -> {sc}")
    print(Fore.RED+f"\033[35m[\033[36m+\033[35m] Targeting URL -> {url}")
    print(Fore.RED+f"\033[35m[\033[36m+\033[35m] Time Started  -> " + str(datetime.now()))
    t.sleep(2)
    file2 = open(f"{payload}", 'r')
    l1    = file2.readlines()
    count = -0
    for line in l1:
        while True:
            file = open(f"{payload}", 'r')
            l = file.readlines()
            count =+ 1
            for line in l:
                    print("[~] Testing Payload -> {} : {}".format(count, line.strip()))
                    forms = get_all_forms(url)
                    js_script = "{}".format(line.strip())
                    is_vulnerable = False
                    for form in forms:
                        form_details = get_form_details(form)
                        content = submit_form(form_details, url, js_script).content.decode()
                        if js_script in content:
                            print(f"\033[31m\033[43m[+] XSSI Returned True At -> \033[49m" + str(datetime.now()))
                            print(f"\033[31m\033[43m[+] XSS Detected -> \033[49m{url}\033[49m") 
                            ptable = PrettyTable(["Content and form details"])
                            ptable.add_row([form_details])
                            print("\033[32m",ptable)
                            is_vulnerable = True
        return is_vulnerable
if __name__ == "__main__":
    url = sys.argv[1]
    payload = sys.argv[2]
    gethttp()
    reline()
    print(scan_xss(f'{url}'))
    print("\033[31m[*] Ended At : " + str(datetime.now()))