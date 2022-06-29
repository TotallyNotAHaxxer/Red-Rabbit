#!/bin/bash

red="\e[0;31m"
green="\e[0;32m"
off="\e[0m"


function linux() {
    if [ -d "/usr/share/RR" ]; then
        echo -e "$red [$green+$red]$off A Directory /usr/share/RR Was Found! Do You Want To Replace It? [Y/n]:" ;
        read replace
        if [ "$replace" = "y" ]; then
          sudo rm -r "/usr/share/RR"
          sudo rm "/usr/share/icons/RR.jpg"
          sudo rm "/usr/share/applications/RR.desktop"
          sudo rm "/usr/local/bin/RR"
          echo -e "$red [$green+$red]$off Installing ...";
          echo -e "$red [$green+$red]$off Creating Symbolic Link ...";
          echo -e "#!/bin/bash go run /usr/share/RR/main.go" '${1+"$@"}' > "RR";
          chmod +x "RR";
          sudo mkdir "/usr/share/RR"
          sudo cp "/home/xea43p3x/Desktop/RR6/main.go" "/usr/share/RR6"
          sudo cp "RR.jpg" "/usr/share/icons"
          sudo cp "RR.desktop" "/usr/share/applications"
          sudo cp "RR" "/usr/local/bin/"
          rm "RR";

    if [ -d "/usr/share/RR" ] ;
    then
    echo -e "$red [$green+$red]$off Tool Successfully Installed And Will Start In 5s!";
    echo -e "$red [$green+$red]$off You can execute tool by typing RR"
    sudo RR;
    else
    echo -e "$red [$green✘$red]$off Tool Cannot Be Installed On Your System! Use It As Portable !";
        exit
    fi 
}

linux