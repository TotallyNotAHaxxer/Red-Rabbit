import os, sys, time, pandas, colorama, subprocess
from scapy.all import * 
from threading import Thread 
from colorama import Fore, Back, Style
from prettytable import PrettyTable

def scan_networks():
    networks = pandas.DataFrame(columns=["BSSID", "SSID", "dBm_Signal", "Channel", "Crypto"])
    networks.set_index("BSSID", inplace=True)
    def callback(packet):
        if packet.haslayer(Dot11Beacon):
            bssid = packet[Dot11].addr2
            ssid = packet[Dot11Elt].info.decode()
            try:
                dbm_signal = packet.dBm_AntSignal
            except:
                dbm_signal = "N/A"
            stats = packet[Dot11Beacon].network_stats()
            channel = stats.get("channel")
            crypto = stats.get("crypto")
            networks.loc[bssid] = (ssid, dbm_signal, channel, crypto)
    def print_all():
        while True:
            os.system("clear")
            import datetime 
            from datetime import datetime
            import time as t 
            a = str(datetime.now())
            print("\033[35m[\033[36mTime Elapsed > \033[35m]" + a)
            mac = networks
            Mac_table = PrettyTable(["Networks Within Length of Interface"])
            Mac_table.add_row([mac])
            print(Mac_table)
            t.sleep(0.5)
    def change_channel():
        ch = 1
        while True:
            os.system(f"iwconfig {interface} channel {ch}")
            ch = ch % 14 + 1
            time.sleep(0.5)
    if __name__ == "__main__":
        os.system("clear")
        interface = str(input("\033[35m[\033[36m!\033[35m] Interface >>> "))
        print("\033[35m[\033[36m+\033[35m] Starting interface....")
        time.sleep(0.1)
        os.system(f"sudo airmon-ng start {interface}")
        print("\033[35m[\033[36m?\033[35m] Interface Started? ")
        time.sleep(4)
        print("\033[35m[\033[36m+\033[35m] Scanning Networks....")
        time.sleep(3)
        print("\033[35m[\033[36m+\033[35m] CTRL+C When your Done")
        time.sleep(2)
        printer = Thread(target=print_all)
        printer.daemon = True
        printer.start()
        channel_changer = Thread(target=change_channel)
        channel_changer.daemon = True
        channel_changer.start()
        sniff(prn=callback, iface=interface)

if __name__ == "__main__":
    scan_networks()