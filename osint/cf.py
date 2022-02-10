import requests, re, os, datetime
from datetime import datetime 


with open('cu.txt', 'r') as f:
    print(f.read())
try:
    countries = ["US", "JP", "IT", "KR", "FR", "DE", "TW", "RU", "GB", "NL",
                 "CZ", "TR", "AT", "CH", "ES", "CA", "SE", "IL", "PL", "IR",
                 "NO", "RO", "IN", "VN", "BE", "BR", "BG", "ID", "DK", "AR",
                 "MX", "FI", "CN", "CL", "ZA", "SK", "HU", "IE", "EG", "TH",
                 "UA", "RS", "HK", "GR", "PT", "LV", "SG", "IS", "MY", "CO",
                 "TN", "EE", "DO", "SI", "EC", "LT", "PS", "NZ", "BD", "PA",
                 "MD", "NI", "MT", "IT", "SA", "HR", "CY", "PK", "AE", "KZ",
                 "KW", "VE", "GE", "ME", "SV", "LU", "CW", "PR", "CR", "BY",
                 "AL", "LI", "BA", "PY", "PH", "FO", "GT", "NP", "PE", "UY",
                 "-"]
    headers = {"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"}
    headers = {"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"}
    num = int(input("@>>> "))
    if num not in range(1, 91+1):
        print("Sorry that isnt a option! ")
    #
    country = countries[num-1]
    res = requests.get(f"http://www.insecam.org/en/bycountry/{country}", headers=headers)
    last_page = re.findall(r'pagenavigator\("\?page=", (\d+)', res.text)[0]
    for page in range(int(last_page)):
        res = requests.get(f"http://www.insecam.org/en/bycountry/{country}/?page={page}",headers=headers)
        find_ip = re.findall(r"http://\d+.\d+.\d+.\d+:\d+", res.text)
        with open("out.txt", 'a' ) as out:
            for ip in find_ip:
                print("\033[1;34m[INFO] ===> " + str(datetime.now()))
                print(".............................................")
                print("\033[1;31m[INFO IP FETCHED] ===> ", ip)
                print("\033[32m WARNING: IP's ARE LOGGED TO A FILE -> RR5/modules/osint/out.txt")
                out.write(ip + '\n')
except:
    pass
finally:
    print("\033[1;37m")