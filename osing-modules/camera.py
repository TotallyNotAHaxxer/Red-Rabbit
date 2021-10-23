import requests
import re 
import os 
import colorama 
from colorama import init, Fore
import pyfiglet
from datetime import datetime 
os.system(' clear ')
print(Fore.RED+" ")
result = pyfiglet.figlet_format("Cam+s", font = "isometric1" )
print(result)
print(Fore.RED+"                               [+] The Camera Stalker")
print("""
|=========================================================================================|
| [1] United States               [31] Mexico                [60] Moldova                 |
| [2] Japan                       [32] Finland               [61] Nicaragua               |
| [3] Italy                       [33] China                 [62] Malta                   |
| [4] Korea                       [34] Chile                 [63] Trinidad And Tobago     |
| [5] France                      [35] South Africa          [64] Soudi Arabia            |
| [6] Germany                     [36] Slovakia              [65] Croatia                 |
| [7] Taiwan                      [37] Hungary               [66] Cyprus                  |
| [8] Russian Federation          [38] Ireland               [67] Pakistan                |
| [9] United Kingdom              [38] Egypt                 [69] United Arab Emirates    |
| [10] Netherlands                [39] Thailand              [70] Kazakhstan              |
| [11] Czech Republic             [40] Ukraine               [71] Kuwait                  |
| [12] Turkey                     [41] Serbia                [71] Venezuela               |
| [13] Austria                    [42] Hong Kong             [73] Georgia                 |
| [14] Switzerland                [43] Greece                [74] Montenegro              |
| [15] Spain                      [44] Portugal              [75] El Salvador             |
| [16] Canada                     [45] Latvia                [76] Luxembourg              |
| [17] Sweden                     [46] Singapore             [77] Curacao                 |
| [18] Israel                     [47] Iceland               [78] Puerto Rico             |
| [19] Iran                       [48] Malaysia              [79] Costa Rica              |
| [20] Poland                     [49] Colombia              [80] Belarus                 |
| [21] India                      [50] Tunisia               [81] Albania                 |
| [22] Norway                     [51] Estonia               [82] Liechtenstein           |
| [23] Romania                    [52] Dominican Republic    [83] Bosnia And Herzegovia   |
| [24] Viet Nam                   [53] Sloveania             [84] Paraguay                |
| [25] Belgium                    [54] Ecuador               [85] Philippines             |
| [26] Brazil                     [55] Lithuania             [86] Faroe Islands           |
| [27] Bulgaria                   [56] Palestinian           [87] Guatemala               |
| [28] Indonesia                  [57] New Zealand           [88] Nepal                   |
| [29] Denmark                    [58] Bangladeh             [89] Peru                    |
| [30] Argentina                  [59] Panama                [90] Uruguay                 |
|==================================== [91] Extra==========================================|
""")
try:
    print()
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
        raise IndexError
    country = countries[num-1]
    res = requests.get(
        f"http://www.insecam.org/en/bycountry/{country}", headers=headers
    )
    last_page = re.findall(r'pagenavigator\("\?page=", (\d+)', res.text)[0]
    for page in range(int(last_page)):
        res = requests.get(
            f"http://www.insecam.org/en/bycountry/{country}/?page={page}",
            headers=headers
        )
        find_ip = re.findall(r"http://\d+.\d+.\d+.\d+:\d+", res.text)
        for ip in find_ip:
            print("\033[1;34m[INFO] ===> " + str(datetime.now()))
            print(".............................................")
            print("\033[1;31m[INFO IP FETCHED] ===> ", ip)
except:
    pass
finally:
    print("\033[1;37m")

