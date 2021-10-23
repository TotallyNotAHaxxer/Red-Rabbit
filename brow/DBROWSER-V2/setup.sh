
# colors 

g="\033[1;32m"
r="\033[1;31m"
b="\033[1;34m"
w="\033[0m"

clear



echo ' [!] Checking if scripts are installed [!] '
sleep 1
clear 
sleep 1


echo 'checking for if python3 is installed'
sleep 2 
pkg=python3 
status="$(dpkg-query -W --showformat='${db:Status-Status}' "$pkg" 2>&1)"
if [ ! $? = 0 ] || [ ! "$status" = installed ]; then
  sudo apt install $pkg
fi
sleep 1 
echo ' [+] :D yay it is installed [+] '


echo 'checking for if netsniff-ng is installed'
sleep 2 
pkg=netsniff-ng 
status="$(dpkg-query -W --showformat='${db:Status-Status}' "$pkg" 2>&1)"
if [ ! $? = 0 ] || [ ! "$status" = installed ]; then
  sudo apt install $pkg
fi
sleep 1 
echo ' [+] :D yay it is installed [+] '


echo 'checking for if wireshark is installed'
sleep 2 
pkg=wireshark
status="$(dpkg-query -W --showformat='${db:Status-Status}' "$pkg" 2>&1)"
if [ ! $? = 0 ] || [ ! "$status" = installed ]; then
  sudo apt install $pkg
fi
sleep 1 
echo ' [+] :D yay it is installed [+] '

echo 'checking for if proxychains is installed'
sleep 2 
pkg=proxychains
status="$(dpkg-query -W --showformat='${db:Status-Status}' "$pkg" 2>&1)"
if [ ! $? = 0 ] || [ ! "$status" = installed ]; then
  sudo apt install $pkg
fi
sleep 1 
echo ' [+] :D yay it is installed [+] '


echo 'checking for if tor is installed'
sleep 2 
pkg=tor
status="$(dpkg-query -W --showformat='${db:Status-Status}' "$pkg" 2>&1)"
if [ ! $? = 0 ] || [ ! "$status" = installed ]; then
  sudo apt install $pkg
fi
sleep 1 
echo ' [+] :D yay it is installed [+] '

if python -c "import colorama" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh'
fi

sleep 0.5 # time sleep for print and checking 

if python -c "import os" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh'
fi

sleep 0.5

if python -c "import sys" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh installing script'

fi

sleep 0.5

if python -c "import time" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh- this is not right '
fi

sleep 0.5

if python -c "import pyfiglet" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh installing script [+] '
    pip install pyfiglet 
fi

sleep 0.5

if python -c "import PyQt5" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh something went wrong attemtping install [+] '
    pip install PyQt5
fi

sleep 0.5

if python -c "import PyQt5.QtWebEngineWidgets" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh something isnt right installing module'
    pip install PyQt5.QtWebEngineWidgets
fi

sleep 0.5

if python -c "import PyQt5.QtPrintSupport" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'hm i do not see a module, installing !'
    pip install PyQt5.QtPrintSupport
fi

sleep 0.5

if python -c "import PyQt5.QtWidgets" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'doesnt seem to be here installing module [+]'
    pip install PyQt5.QtWidgets
fi

sleep 0.5

if python -c "import PyQt5.QtGui" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh looks like i should install this'
    pip install PyQt5.QtGui
fi


if python -c "import proxychains" &> /dev/null; then
    echo ' [+] all good module is installed [+] '
else
    echo 'uh oh looks like i should install this'
    sudo apt-get install proxychains 
fi

sleep 0.5
clear 
echo ' everything is finished, have fun! :D'
sleep 2

sleep 0.1

echo ' the date is ' 
date 
sleep 1
echo '------------------------------------- '
sleep 0.1
echo ' press enter to continue to browser '
sleep 0.1
echo '------------------------------------- '

sleep 1

read varname 

python3 main.py

exit  # exit the script without any input
