import os, sys, requests, json


token_log = f"config-files/token.conf"
token_red = open(f"{token_log}", 'r')
AUTH = token_red.read()

header = {
    'authorization': f'{AUTH}'
}


def banner(file):
    fp = f"txt/{file}"
    file_r = open(f"{fp}", 'r')
    out = file_r.read()
    print("\033[31m", out)
# fuck i hate that there is no (end) statement to use


def retrieve_messages(channelid, header):

    headers = f"{header}"
    
    r = requests.get(f'https://discord.com/api/v9/channels/{channelid}/messages', headers=headers)
    jsonn = json.loads(r.text)
    for value in jsonn:
        print("\033[31m")
        print("[INFO] ===> " + str(datetime.now()))
        print("\03335m")
        print(value, '\n')
        jsonFile = open("data.json", "w")
        jsonFile.write(f"{r.text}\n")
        jsonFile.close()

def main():
    in_in = str(input("Options >"))
    if in_in == "1":
        channel = str(input("Enter channel ID > "))
        retrieve_messages(f"{channel}", f"{header}")



if __name__ == "__main__":
    banner("banner-discord.txt")