| Languages     | Why the lang                            |
| ------------- | ----------------------------------------|
| Python3       | Easy write for some scripts             |
| Perl          | using the open syntax for Injection     |
| Go-lang       | Speed and Read abilities,also format    |
| Rust          | Speed, mobility, and error handeling    |
| Ruby          | as a main file derivitive from rube-ster|
| Shell         | Grep, wget, godorking, osint, etc       |
| Batch         | Viruses, payloads, bombs,               |
| Assembly      | geting auth cpu and writing data to file|


What is red rabbit? <br>

red rabbit is a exploitation, osint, virus, recon, and hacking script developed with many languages and built for speed to aid hackers in everyday attacks like XSSI, XSSI testing, SQLI testing, PHPI, testing and learning about rouge acess points, file dumping, FTP, running people osint, ajaz spiders and more<br> 

# possible scripts amd usages out of 5-6 menus<br> 
PHPI, SQLI, XSSI, Ajax crawl, ajaz crawl + whois, whois, domain dorker, payload, xss scan, sql scan, ftp scan, port scan, web scan, dhcmp scan, web host scan, whois trace, whois mine, spawn fake ap's, scan local ports, scan the local host, fuzz a file app, web app fuzz, flood a host, scan the local area for bssid's, deauthenticate networks, activate interfaces, grab code notes in html, grab urls in a html file, dive in html, parse html, run MOAB osint, nuke a domain, view my website,  run osint on complex urls 
<br>


# installs and other 
`git clone https://github.com/ArkAngeL43/Red-Rabbit-V4 ; chmod +x ./install.sh ; ./install.sh ; `

# warnings and proper install upon entering<br>
when you use this script when it comes to wireless attacks that use packet gen like <br>
`fake ap`<br>
`Deauthenticaton`<br>
`ftp login captures`<br>
then you will need to run the main file as root like this 

```
Using Current CPU =>  < XAuthenticAMD>
                                           [1] Rouge AP 
                                           [2] Deauth 
                                           [3] Port Scanner 
                                           [A] Web Port Scan 
                                           [4] Flooder 
                                           [5] DHCMP ATK 
                                           [6] Check Connection 
                                           [7] WHOIS Domain 
                                           [8] More 
                                           [9] Start Interface  [aircrack-ng suit]
                                           [01] Start Interface [Iw dev for fake AP]
                                           [02] Stop Interface [iw dev managed mode]
                                           [0] Exit 
                                      >>> 


```

you will need to run 01, then rerun the file MAKE SURE YOU HAVE MON0 as a INTERFACE 802.11

further notes about this secition 
<br>
there will be some errors with the fake AP sometimes it will not spawn this might be because of your network card 
<br>
also when you do option 01 this uses phy dev to erase wlan0 and start a mon0 card, this will be permanant until you chose to stop the interface using 02 ONTOP of that this will also ruin some bluetooth connections if the card is weak, if you do not have a second interface or an extra adapter this will also ruin connections on calls or anything else you may be doing until the script is done, unless its another hacking tool<br>
<br> 

# issues and errors within the script<br>  
<br>
802.11 packetgen utils and wireless cards<br>
<br>
sometimes you wont be able to activate the interface well check the following <br> 
when you started the script and your card were you root?<br>
when you ran the script again after enabling were you root when doing so?<br>
<br>

whois domain <br>
<br>
there is no issue with the script itself it just depends on what domain your targeting, during whois parsing the thing can get stuck and sometimes will not register properly, i prefer using MOAB-OSINT in menu 3 of the script 

SQLI perl module<br>

the PHPI, SQLI, and LW3 module are all perl scripts, this will require some heafty syntaxing. sometimes when you run the PHP script it will error out, i still have not figured out why <br>
<br>
the SQLI script will say "failed to fetch" then the parsed url with query strings, this is probobly because the url isnt being parsed correctly/inputted or the server is not MySQL V5.x or it can be older or even younger, this module ONLY SUPPORTS `MySQL V5.x`<br>
<br>
issues with LW2 and Libwhisker 
<br>
sometimes it will say there is an error in the libwhisker but it shouldnt be an issue, its more of a bug than an issue
<br>
# SCRIPTS USED INSIDE OF RED RABBIT 

`before i go on i want to note that 99% of the scripts are all my scripts asides people like rip, parrot, or some other freinds helped me build `
`scripts like ; LW2, LIBWHISKER, PERL6 and even the php script ( eevn though i fixed 90% of it ) are not mine, i do not claim ownership for these modules or scripts`
injection/exploitation/dne enumeration
-----------
PHPI      -> phpi.pl
SQLI      -> sqli.pl
XSSI      -> xssi.py
port      -> scanner.rs
port web  -> port.rb 
ssh sploit-> ssh-sploit | main.rb 
MOAB OSI  -> moab.sh/go.mod.go.sum.main.go.banner.go.regex.go others 



