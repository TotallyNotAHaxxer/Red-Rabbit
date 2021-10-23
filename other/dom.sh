RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"
RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"
red=`tput setaf 1`
white=`tput setaf 7`
echo '============== Domain again '
read domain
working_dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
results_dir=$working_dir/results
if [ $# -eq 0 ]
then
    for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo
    cat << "EOF"
 __   __   __  __     __         __   __     ______     __  __     ______     ______    
/\ \ / /  /\ \/\ \   /\ \       /\ "-.\ \   /\  ___\   /\ \_\ \   /\  ___\   /\  == \   
\ \ \'/   \ \ \_\ \  \ \ \____  \ \ \-.  \  \ \ \____  \ \  __ \  \ \  __\   \ \  __<   
 \ \__|    \ \_____\  \ \_____\  \ \_  "\_\  \ \_____\  \ \_\ \_\  \ \_____\  \ \_\ \_\ 
  \/_/      \/_____/   \/_____/   \/_/ \/_/   \/_____/   \/_/\/_/   \/_____/   \/_/ /_/ 
EOF
    for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
    echo -e "$RED [!] I need a fucking domain dumbass "
    echo -e "$RED [!] cool.sh <url> <domain> you fucking dumbass"
    echo -e "$RED [!] clear ; ./cools.sh http://testphp.vulnweb.com/listproducts.php?cat=1 testphp.vulnweb.com"
    exit 1
fi
if [ $# -eq 1 ]
then
    for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo
    cat << "EOF"
 __   __   __  __     __         __   __     ______     __  __     ______     ______    
/\ \ / /  /\ \/\ \   /\ \       /\ "-.\ \   /\  ___\   /\ \_\ \   /\  ___\   /\  == \   
\ \ \'/   \ \ \_\ \  \ \ \____  \ \ \-.  \  \ \ \____  \ \  __ \  \ \  __\   \ \  __<   
 \ \__|    \ \_____\  \ \_____\  \ \_  "\_\  \ \_____\  \ \_\ \_\  \ \_____\  \ \_\ \_\ 
  \/_/      \/_____/   \/_____/   \/_/ \/_/   \/_____/   \/_/\/_/   \/_____/   \/_/ /_/ 
EOF
    for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
    echo -e "$RED [!] I need a fucking domain dumbass "
    echo -e "$RED [!] cool.sh <url> <domain> you fucking dumbass"
    echo -e "$RED [!] clear ; ./cools.sh http://testphp.vulnweb.com/listproducts.php?cat=1 testphp.vulnweb.com"
    exit 1
fi
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo
go run banner.go 
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo
echo -e "[~] Targeting Domain ->  $domain"
FILES="item_id=15"
for ELEMENT in ${FILES[@]}
do
echo -e "[~] Fixed URL -> $domain $FILES"
echo -e "[!] Trying Payload ${FILES}"
URL="http://google.com/search?hl=en&safe=off&q="
STRING=`echo $domain | sed 's/ /%20/g'`
URI="$URL%22$domain%22"
lynx -dump $URI > gone.tmp
sed 's/http/\^http/g' gone.tmp | tr -s "^" "\n" | grep http| sed 's/\ .*//g' > gtwo.tmp
rm gone.tmp
sed '/google.com/d' gtwo.tmp > urls
rm gtwo.tmp
echo "SExtraction -> Extracted `wc -l urls` and listed them in '`pwd`/urls' file for reference."
echo ""
cat urls
done 
FILES1=("cat=1")
for ELEMENT in ${FILES1[@]}
do
echo -e "[~] Fixed URL -> $domain $FILES"
echo -e "[!] Trying Payload ${FILES1}"
URL="http://google.com/search?hl=en&safe=off&q="
STRING=`echo $domain | sed 's/ /%20/g'`
URI="$URL%22$domain%22"
lynx -dump $URI > gone.tmp
sed 's/http/\^http/g' gone.tmp | tr -s "^" "\n" | grep http| sed 's/\ .*//g' > gtwo.tmp
rm gone.tmp
sed '/google.com/d' gtwo.tmp > urls
rm gtwo.tmp
echo "SExtraction -> Extracted `wc -l urls` and listed them in '`pwd`/urls' file for reference."
echo ""
cat urls
done 