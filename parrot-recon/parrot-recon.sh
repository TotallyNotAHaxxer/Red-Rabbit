#!/bin/bash

#banner
cat << "EOF"
   ___                    __        ___                  
  / _ \___ ____________  / /____   / _ \___ _______  ___ 
 / ___/ _ `/ __/ __/ _ \/ __(_-<  / , _/ -_) __/ _ \/ _ \
/_/   \_,_/_/ /_/  \___/\__/___/ /_/|_|\__/\__/\___/_//_/
                                                         
EOF

# setting up color codes
red=`tput setaf 1`
white=`tput setaf 7`

# enviornment variables for recon tool
domain=$1
working_dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
results_dir=$working_dir/results

# checks to make sure you entered a domain 
if [ $# -eq 0 ]
then
   echo "[!] No Domain Defined"
   echo "[-] Usage: ./parrot-recon.sh <domain>"
   exit 0
fi

# setting up directories for recon tool
echo "[+] Setting Up Enviornment"
if [ ! -d "$results_dir" ]
then
   mkdir $results_dir
fi

# enumeration proccess using tools from install script
echo "$red[+] Starting Nmap Scan$white"
nmap -sV -sC $domain -oA $results_dir/$domain-tcp-scan --open
nmap -sV -sS $domain -oA $results_dir/$domain-udp-scan --open 

echo "$red[+] Starting Sublist3r$white"
sublist3r -d $domain -o $results_dir/subdomains.txt 

echo "$red[+] Starting Subzy Takeover$white"
subzy -targets $results_dir/subdomains.txt > $results_dir/takeover.txt

echo "$red[+] Starting Nikto Scan$white"
nikto -h $domain -o $results_dir/nikto.txt

echo "$red[+] Starting Domain Brute Force$white"
gobuster dir -u $domain -w /usr/share/wordlists/dirbuster/directory-list-1.0.txt -o $results_dir/server-dirs.txt

echo "$red[+] Starting Amass$white"
amass enum -d $domain -o $results_dir/amass.txt -r 8.8.8.8

echo "$red[+] Starting SSL Scans$white"
sslyze --regular $domain > $results_dir/sslyze-regular.txt
sslyze --heartbleed $domain > $results_dir/sslyze-heartbleed.txt
sslyze --robot $domain > $results_dir/sslyze-robot.txt

echo "$red[+] Starting Nuclei Scans$white"
nuclei -u $domain -o $results_dir/nuclei.txt

echo "$red[+] Running Wordpress Scans$white"
wpscan --url $domain -o $results_dir/wordpress.txt --no-banner

echo "$red[+] Script Done!"
