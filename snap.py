import logging
from urllib.parse import urljoin
import requests
from bs4 import BeautifulSoup
import colorama 
import socket 
from colorama import Fore
import os 
import sys 
import datetime 
from datetime import datetime
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
target1 = sys.argv[0]
target2 = sys.argv[1]
r = requests.get(f'{target2}')
logging.basicConfig(
    format='%(asctime)s %(levelname)s:%(message)s',
    level=logging.INFO)
if 'info' in ():
    os.systen('cat info.txt')
def save():
    with open('output.txt', 'w') as f:
        sys.stdout = f
        logging.info(f'{url}')
#sudo nmap -F -sC
    print(" --------------- Running URL Crawler -------------- ")
class Crawler:
    def __init__(self, urls=[]):
        self.visited_urls = []
        self.urls_to_visit = urls
    def download_url(self, url):
        return requests.get(url).text
    def get_linked_urls(self, url, html):
        soup = BeautifulSoup(html, 'html.parser')
        for link in soup.find_all('a'):
            path = link.get('href')
            if path and path.startswith('/'):
                path = urljoin(url, path)
            yield path
    def add_url_to_visit(self, url):
        if url not in self.visited_urls and url not in self.urls_to_visit:
            self.urls_to_visit.append(url)
    def crawl(self, url):
        html = self.download_url(url)
        for url in self.get_linked_urls(url, html):
            self.add_url_to_visit(url)
    def run(self):
        while self.urls_to_visit:
            url = self.urls_to_visit.pop(0)
            original_stdout = sys.stdout
            print("\033[35m===============================================================")
            logging.info(f'\033[31m[DATA] CRAWLING ===> {url}')
            os.system(f"ruby ajax-whois.rb {url}")
            print("\033[35m===============================================================")
            try:
                self.crawl(url)
            except KeyboardInterrupt:
                sys.exit()
            except Exception:
                logging.exception(f'[!] WARNING ===> FAILED TO CRAWL {url}')
            finally:
                self.visited_urls.append(url)
if __name__ == '__main__':
    Crawler(urls=[f'{target2}']).run()
    save()