clear
red=`tput setaf 1`
white=`tput setaf 7`
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
red='tput setaf 1'
white='tput setaf 7'
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
   echo "[-] Usage: ./osint-con.sh <domain>"
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
echo "$red [+] Running Google Dorker for $domain$white"
chmod +x ./dork.sh ; ./dork.sh $domain