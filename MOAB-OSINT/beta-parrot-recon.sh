#!/bin/bash
printf "\x1b[H\x1b[2J\x1b[3J"
cat << "EOF"
   ___                    __        ___                  
  / _ \___ ____________  / /____   / _ \___ _______  ___ 
 / ___/ _ `/ __/ __/ _ \/ __(_-<  / , _/ -_) __/ _ \/ _ \
/_/   \_,_/_/ /_/  \___/\__/___/ /_/|_|\__/\__/\___/_//_/
By -> ParrotAssasin15 || -> ArkAngeL43 
   /.\                          
   |  \                  
   /   \                 
  //  /                  
  |/ /\_
 / /            
/ /     
\/ 
EOF

# env variables 
red=`tput setaf 1`
white=`tput setaf 7`
domain=$1
url=$2
domain2=$3
working_dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
results_dir=$working_dir/results
tools_dir=$working_dir/tools


echo "[*] Targeting   -> $domain"
echo "[*] IP Address -> $(host $domain | awk '/has address/ { print $4 ; exit }')"


if [ $# -eq 0 ]
then
   echo "[!] No Domain Defined"
   echo "[-] Usage: ./parrot-recon.sh <domain>"
   exit 0
fi

# root check
if [ `whoami` != "root" ]
then
   echo "[!] This Script Needs To Be Run As Root User"
   exit 0
fi

echo "[+] Setting Up Enviornment"
if [ ! -d "$results_dir" ]
then
   mkdir $results_dir
fi


echo "$red [+] Runnign Google Dorker for $domain$white"
chmod +x ./dork.sh ; ./dork.sh $domain
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m─────\e[0m" ; done ; echo 
echo "$red[+] Starting Request Enumeration$white"
go run go-serve.go $url # $results_dir/rep.txt
ruby http.rb $domain 
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m─────\e[0m" ; done ; echo 
ruby scan.rb $domain
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m─────\e[0m" ; done ; echo 
echo "$red[+] Script Done!$white"
echo "$red[+] Check Your Results Directory For The Output!$white"
echo "$red[!] There Are Some Tools That Cannot Save Output$white"
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m─────\e[0m" ; done ; echo 
