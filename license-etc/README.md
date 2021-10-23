# Red-Rabbit-V3.9
` WARNING BEFORE RUNNING THE TOOL PLEASE USE THE FOLLOWING COMMAND ` <br> 
`iw phy phy1 interface add mon0 type monitor && ifconfig mon0 up > THIS NEEDS TO BE IN A ROOT TERMINAL`

<h1> written by </h1>

<img src="programmed.png" alt="homies">


What is Red-Rabbit

Red-Rabbit is a massive framework written from<br>
perl, rust, bash, batch, ruby, and go-lang <br>
this project is a massive one and is written with <br>
alot of effort, this has taken me a long time to write<br>
as i have had issues with me OS being weird lol<br>
not only that but red-rabbit is a really fast penetration testing 
platform and multi tool for new and even experinced hackers 


well you might ask what can it do?<br>

it can<br>

`port scan a host`<br>
`port scan your network`<br>
`port scan a web host`<br>
`spawn Fake Acess points`<br>
`Commit DHCMP attacks`<br>
`test your connection`<br>
`run whois tracers`<br>
`DOS a host`<br>
`test if a url is xss injectable`<br>
`test if a web host is SQL injectable`<br>
`Fuzz a file sharing server`<br>
`Capture all network BSSID's in range`<br>
`Change your interface mode`<br>
`and more `<br>
<br>
LANGUAGES USED<br> 

BATCH, PERL, GO, BASH, RUBY (MAIN), RUST, PYTHON3, PYTHON2

this tool is inspired by one of my first ruby frameworks 
for ethical hacking and web/host discovery
that tool was named Rube-Ster which had alot 
of bugs and a shit ton of work that was thrown 
to the side, so i decided to start this project 
which is well Red-Rabbit, you may ask where 
did the name come from, well initally Rube-Sters 
name cam from ruby hence the Rube- and the ster 
came from a bunny i used to know and it was named mr booster
sadly he passed away, anyway hence the name ster
then red rabbit which is a name derived from the og name 


================
Why this tool? 
===============

this tool has alot of options now for choices such as 

wifi Death, Fake AP, web port scanners, host port scanners 
dns, whois, loggers, banner and title parsers, and is written 
from mainly ruby however utilizes, perl, rust, go, ruby, bash, and batch 
you will notice batch is for win32-64 installs and bash is used for linux installs 



=====================
why so many languages 
=====================
well i wanted speed, and since i am currently learning 
rust, c, perl, ruby, go, and batch i decided to put 
my skills to the very well known test 

i also wanted speed 

to parse the results of a title of the domain
and to grab it faster i used golang 

Go -> go-title.go

i also wanted better exception handeling and easy 
etsting, especially with net/http parsing and result testing 

so i used perl so i can throw the URl's into a list, parse them 
and get faster results for testing a internet or stable connections 
perl is also really good for formatting 

i wanted it to somehwta be cross platform 

so i used bash and batch for the installs 

==========
why rust?
===========

simply for faster current network host identification 

====================
why rust IP sniffing?
====================

rust is really fast and a really good language compared to golang 
sure golang is built from assembly but rust over all is faster 
when it comes to handeling, socks, networking threadings and more 
so i built a small IP sniffer from rust 

============================================================
why make the main file in ruby if other languages are faster
============================================================ 

well currently im reading a few books with ruby, and wanted to put 
my ruby skills to the test to see my limits, and ruby is alot better 
when it comes to offensive security tools with networking and sending
payloads over the network or even making something like a windows 
trojan, so i decided to use it 

if your still confused and want to debate then ask yourself 

why is the biggest exploitation framework and the most powerful (MSF/metasploit)
is 97% built from pure ruby? 

=================================== what can this tool do ======================

spawn fake ap's
deauthenticate clients off a network 
do whois domain tracing 
reverse a dns
launch DHCMP attacks 
Flood networks 
Port Scan Hosts 
Port Scan Web Hosts 
IP Parse 
Find ports on the local network 
check your current connection 
Scan the local area for BSSID's and ranges ( its unorganized )
Fuzz File Sharing Websites ( crash and exploit the servers )
Arp Spoof Clients off the current network 
AP scan 




=============================== REQUIREMENTS ========================
perl
python 
rust 
rustc
crates 
cargo 
cpan 
ruby
bash or batch 
golang 

service/script REQUIREMENTS for modules 

ruby ===

net/http
socket 
time 
awaite
optparse 
iw phy
timeout 
http party 
open uri
uri
whois 
whois-parser
colorize
tty-spinner
ruby-gems 
openssl

Go ====

a fucking os 
a sys 
fmt 
strings 
net/http
net/html 

perl === 
Ansi color 
socket 
Strict 
HTTP Tiny

rust === 
use std::env;
use std::process::Command;
use std::io::{self, Write};
use std::net::{IpAddr, TcpStream};
use std::str::FromStr;
use std::process;
use std::sync::mpsc::{Sender, channel};
use std::thread;


=============================== WARNINGS ====================

ME OR ANYONE WHO CONTRIBUTED OR GAVE IDEAS ARE 
RESPONSIBLE FOR YOUR DUMBASS MAKING DUMBASS DECISIONS 
WE HIGHLY DISCLAIM AGAINST USING THIS TOOL FOR MALICOUS 
ACTIVITY IF YOU HAVE A PROBLEM 






------------------------------------------------------- BUGS/ERRORS IN 3.9 CURRENT STABLE ----------------------------------------------------------------

`in the secondary menu in option 8 the rust scanner is not working properly if you want to use it to scan your local network please run the following command in your prompt `<br>
`cd src ; cargo run -- -j 1000 <HOME NETWORK IP> EX -> 10.0.0.0`

`in the secondary menu where SQLI and XSSI are tested the way links are parsed by BS4 are a bit weird and im quite sorry, i can fix it however this is a seperate file and all ruby does is run that file`<br>

`BLink animation affect in perl test-con.pl` <br>

`sometimes the blinking URL's will cause an issue with display, this wont happen for everyone however it has happened once within the month of testing `<br>

`Length Exception in rust main.rs`<br>

`when the port scanner runs sometimes there will be length exceptions and socket IP binding issues where the length is neaither IPV4-IPV6 or is not supported, note this tool can not scan IP's outside of the private network`<br>
<br>
<br>
`Errcon Refused in URL testiong`<br>
`sometimes perl will throw exceptions like a con was refused however the script will still run`<br>
<br>
                                                                                                 
`BAD CONNECTION test.pl`<br>
`sometimes givin the url, if the indexed HTTP error code is not within the error code list it will throw a warning with testing in connections`<br>
<br>
<br>
`Bug in GO file`<br>
`sometimes when you are during URl/Domain Recon adn at the bottom the go file runs, somtimes it wont properly load, and it wont even print the domain fetched from the URL, i have assumed this depends on the URl and the Domain server your inspecting some sites like google or twitter wont work but smaller connections will`
<br>
<br>

=========================================================== EXAMPLES AND DEMOS OF SCRIPTS ===========================================


NEXT EXAMPLE DATA IS NOT FORMATED LIKE THAT IN THE SCRIPT IT IS ALL FORMATTED IN A LINE ASIDES SERVER SESSIONS

MODULE CHECK 

`[~] Checking Mods Before Run`<br>
`[▇] Checking Module ... Done!`<br>
`[✅️] Module Found`<br>
`[▇] Checking Module ... Done!`<br>
`[✅️] Module Found`<br>
`[▇] Checking Module ... Done!`<br>
`[✅️] Module Found`<br>
`[▇] Checking Module ... Done!`<br>
`[✅️] Module Found`<br>
`[▇] Checking Module ... Done!`<br>
`[✅️] Module Found`<br>
`[▇] Checking Module ... Done!`<br>
`[✅️] Module Found`<br>

  
  ================================================= DNS RECON ===============================

<img src="dns.png" alt="homies">



`EX: twitter.com
Domain Name >>> twitter.com
---------
EX: http://twitter.com
http URL    >>> http://twitter.com
------------------------------------------------------------------------
[+] Created    : 2000-01-21 11:28:17 UTC
[+] Registered : true

#<struct Whois::Parser::Contact 
id=nil, 
type=3, 
name="Tech Admin", 
organization="Twitter, Inc.", 
address="1355 Market Street", 
city="San Francisco", 
zip="94103", 
state="CA", 
country=nil, 
country_code="US", 
phone="+1.4152229670", 
fax="+1.4152220922", 
email="domains-tech@twitter.com", 
url=nil, created_on=nil, 
updated_on=nil>
-------------------------------------------------------------------------
 
     ______     ______     _____     ______     ______     ______     ______     __     ______  
    /\  == \   /\  ___\   /\  __-.  /\  == \   /\  __ \   /\  == \   /\  == \   /\ \   /\__  _\ 
    \ \  __<   \ \  __\   \ \ \/\ \ \ \  __<   \ \  __ \  \ \  __<   \ \  __<   \ \ \  \/_/\ \/ 
     \ \_\ \_\  \ \_____\  \ \____-  \ \_\ \_\  \ \_\ \_\  \ \_____\  \ \_____\  \ \_\    \ \_\ 
      \/_/ /_/   \/_____/   \/____/   \/_/ /_/   \/_/\/_/   \/_____/   \/_____/   \/_/     \/_/ 
         V 3.0                                 Red Rabbit               Scare_Sec Hackers
                                                
                                                ((`\
                                             ___ \\ '--._
                                          .'`   `'    o  )
                                          /    \   '. __.'
                                        _|    /_  \ \_\_
                                       {_\______\-'\__\_\
                                        -----------------

         Date At Start ===> 2021-10-08 20:14:05.674549226 -0400
         Url Target    ===> http://twitter.com
         WWW Target    ===> twitter.com
-------------------------------------------------------
[*] Target is => http://twitter.com
-------------------------------------------------------
[*] Gathering Info on URL => http://twitter.com
[+] good
---------------- BASIC INFORMATION FOR URL -------------- 
"date => Sat, 09 Oct 2021 04:14:01 UTC"
"server => tsa_b"
"location => https://twitter.com/"
"set-cookie => personalization_id=\"v1_GfXLgWue46PteWdF2wB8wA==\"; Max-Age=63072000; Expires=Mon, 09 Oct 2023 04:14:01 GMT; Path=/; Domain=.twitter.com; Secure; SameSite=None, guest_id=v1%3A163375284196230792; Max-Age=63072000; Expires=Mon, 09 Oct 2023 04:14:01 GMT; Path=/; Domain=.twitter.com; Secure; SameSite=None"
"cache-control => no-cache, no-store, max-age=0"
"content-length => 0"
"x-connection-hash => 9b755a57640"
"connection => close"
{"date"=>["Sat, 09 Oct 2021 04:14:01 UTC"], "server"=>["tsa_b"], "location"=>["https://twitter.com/"], "set-cookie"=>["personalization_id=\"v1_GfXLgWue46PteWdF2wB8wA==\"; Max-Age=63072000; Expires=Mon, 09 Oct 2023 04:14:01 GMT; Path=/; Domain=.twitter.com; Secure; SameSite=None", "guest_id=v1%3A163375284196230792; Max-Age=63072000; Expires=Mon, 09 Oct 2023 04:14:01 GMT; Path=/; Domain=.twitter.com; Secure; SameSite=None"], "cache-control"=>["no-cache, no-store, max-age=0"], "content-length"=>["0"], "x-connection-hash"=>["57640"], "connection"=>["close"]}
{"accept-encoding"=>["gzip;q=1.0,deflate;q=0.6,identity;q=0.3"],
 "accept"=>["*/*"],
 "user-agent"=>["Ruby"]}
{"date"=>["Sat, 09 Oct 2021 04:14:02 UTC"],
 "server"=>["tsa_b"],
 "location"=>["https://twitter.com/"],
 "set-cookie"=>
  ["personalization_id=\"v1_9qUMSOtlAzYH42jrPgEwCg==\"; Max-Age=63072000; Expires=Mon, 09 Oct 2023 04:14:02 GMT; Path=/; Domain=.twitter.com; Secure; SameSite=None",
   "guest_id=v1%3A163375284208859520; Max-Age=63072000; Expires=Mon, 09 Oct 2023 04:14:02 GMT; Path=/; Domain=.twitter.com; Secure; SameSite=None"],
 "cache-control"=>["no-cache, no-store, max-age=0"],
 "content-length"=>["0"],
 "x-connection-hash"=>
  ["f6a"],
 "connection"=>["close"]}
-------------------------
[*] Response ~> 301
[*] Checking More Connections..
--------------------------
[*] Gathering Header Info....
[!] Warning, upon further testing of dom-t.rb
[!] Sometimes the server info will go empty
[!] Right now i am planning on fixing this bug 
[!] and massive issue, however this project
[!] was programmed over the course of a month

[*] Query          => 

[*] Scheme         => 
http
[*] Port  Main     => 
80
[*] HOSTNAME       => 
twitter.com
[*] Path           => 

[*] Request URI    => 
/
[*] Server         => tsa_c
[*] Date           => Sat, 09 Oct 2021 04:14:02 GMT
[*] Content        => text/html; charset=utf-8
[*] Response Code  => 301
[*] Last-mod       => 
Sat, 09 Oct 2021 04:14:02 GMT
[*] trans-enc      => 
chunked
`
