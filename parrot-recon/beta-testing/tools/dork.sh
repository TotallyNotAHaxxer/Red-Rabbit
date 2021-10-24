echo -e "[~] Targeting Domain ->  $domain"
FILES=("item_id=15")
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
