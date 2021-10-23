RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"
RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"
red=`tput setaf 1`
white=`tput setaf 7`
echo "$white [+] A third time, one more time after this :D"
read domain
cd go-serve ; go run go-serve.go
working_dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
results_dir=$working_dir/results
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
#test 1
URL="http://google.com/search?hl=en&safe=off&q="
STRING=`echo $domain | sed 's/ /%20/g'`
URI="$URL%22$domain%22" 
lynx -dump $URI > gone.tmp
sed 's/http/\^http/g' gone.tmp | tr -s "^" "\n" | grep http| sed 's/\ .*//g' > gtwo.tmp
rm gone.tmp
sed '/google.com/d' gtwo.tmp > urls
rm gtwo.tmp
echo "SExtraction -> Extracted `wc -l urls` Saved in ->  '`pwd`/urls' file for drawback."
echo ""
cat urls
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
for i in {16..21} {21..16} ; do echo -en "\e[38;5;${i}m──\e[0m" ; done ; echo 
echo "$RED =========================== TESTING URL AGAIN =================="
URL="http://google.com/search?hl=en&safe=off&q="
STRING=`echo $domain | sed 's/ /%20/g'`
URI="$URL%22$domain%22"
lynx -dump $URI > gone.tmp
sed 's/http/\^http/g' gone.tmp | tr -s "^" "\n" | grep http| sed 's/\ .*//g' > gtwo.tmp
rm gone.tmp
sed '/google.com/d' gtwo.tmp > urls
rm gtwo.tmp
echo "SExtraction -> Extracted `wc -l urls` Saved in ->  '`pwd`/urls' file for drawback."
echo ""
cat urls