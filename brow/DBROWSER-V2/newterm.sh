clear
echo ' running netsniff-ng ' 
sleep 1 
echo ' i need you to run it as root :D '
sleep 1 
gnome-terminal -- sudo netsniff-ng --out pack.pcap
clear 
sleep 1

echo '|===========================|'
echo '|file was saved as pack.pcap|'
echo '|========Have a good one====|'

exit 