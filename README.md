<p align="center">
  <img align="center" width="400" height="91" src="git/red_rabbit_git.jfif">
</p>

# Current version

Version 1.0 | UNSTABLE


Developer Note: It sucks i had to declare this version as the unstable version of red rabbit, it was the most well written form of red rabbit ever however after continuous testing the program just does not seem to work the same as it did during testing. 

# What is red rabbit 

The Red Rabbit project is a cyber security framework designed to touch on the most popular hacking and security topics today, such as digital fornsics, stegonography, network forensics, packet sniffing and forensics, brute forcing, hash cracking, file parsing, filepath pillaging, SSH/FTP/TELNET/SQL recon, web recon, web scanning, vulnerability recon, OSINT and much more among that list. With over 260+ utilities all built into red rabbit locally you should have just enough to help you in say a report, or gathering information on a target such as the hostnames, MX records, A records, even getting information such as mac information or phone number information. Red Rabbit really tries to impliment only the best features into its scripts and programs, adding things most frameworks today dont have, while also using raw pure source code to filter out response bodies, response types and even load configuration files 

# What crack head developed this 

Noone but myself (ArkAngeL) this project was done without development teams, without paid producers, without people and communities contributing, and done alone. Which is why this framework is better than most tool based frameworks out there. Bored of automating tools like aircrack? Bored of constantly having people tell you their framework is legit when the code is hidden? well fucking no more! we dont do that here in fact we are the very opposite of every framework out there, we dont automate tools, we write them, we dont lie to people we speak truthfully to our users.

# Donations 

before i get into this whole word thingy i would like you to consider donating to help 
support this project and keep it going!

Venmo  -> RR6-Development

# Information 

> Supported systems 

Red rabbit supports only linux distro's the following are the systems that red rabbit has been tested on 

| OS Name        | Form      |  Version   |
| -------------- | --------- | ---------- |
| Parrot OS      | Debian    | 5.10+      |
| Archman linux  | Arch      | 2020-11-14 |
| BlackArch      | Arch      | 2021.09.01 |
| Kali Linux     | Debian    | 2022.2     |

the best system that red rabbit has worked on is parrot OS security edition

> Installing Red Rabbit 

Installing red rabbit is not too difficult given the makefile it uses however you will need to do some technical work yourself, red rabbit uses a small amount of libraries a majority of them are popular such as PcapPlusPlus for C++/C/Go wifi commands, pcap parsing, packet analysis and other commands that relate to wifi. PcapPlusPlus can be awkward to install so ill walk you through it, first lets go through the makefile. The makefile for red rabbit is written in Perl its named `Makefile.pl` you will run it like `perl Makefile.pl` when you run it with no options there are a few things that are going to be told to you, and directions such as checking for files, or installing directly here are the options and usage of the makefile.

```
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
MAKEFILE: <RR6> ERROR: No arguments defined, please use the following settings
ARGV: install    | installs all files and directories
ARGV: clean      | will clean and erase the RR6 directory
ARGV: check      | will check and run all modules to see if installed
ARGV: checkS     | Will SUPER check all files and modules and install modules
################################################################################
Usage Install < Arch linux >   : sudo perl Makefile.pl install arch 
Usage Install < Debian Linux > : sudo perl Makefile.pl install debain
Usage Check   < Arch linux >   : sudo perl Makefile.pl check arch 
Usage Check   < Debian linux > : sudo perl Makefile.pl check debian 
Usage Checks  < Arch linux >   : sudo perl Makefile.pl checkS arch
Usage Checks  < Debian linux > : sudo perl Makefile.pl checkS debian
```

as you can see its quite simple, even for begginers just run `sudo perl Makefile.pl <install:check:Checks> <debian:arch>`

sudo is needed in order to run both red rabbit and install the libraries, the install option will install all libraries and other utilities needed for this script, the check option will just check all the files and examples to see if they are there, the Checks will check for all files and install all libs, this is a prefered option.

- Installing PcapPlusPlus

Debian systems command 1 `sudo apt-get install libpcap-dev libpcap-dev libpcap0.8 libpcap0.8-dev`

Debian systems command 2 `git clone https://github.com/seladb/PcapPlusPlus.git`

Debian systems command 3 `cd PcapPlusPlus`

Debian systems command 4 `./configure-linux.sh --default`

Debian systems command 5 `make all`

Debian systems command 5 To build the libraries only (without the unit-tests and examples) run `sudo make install`

[if you get issues use sudo during the process]

# Documentation server 

when or if you get red rabbit to run and you enter its user based input interactive console you have an option in there to start the http server the command for this is `start http server` this server will be started on port 5009, the url for this server is `http://localhost:5509` when you log onto the url or reach it once you start the server you will be prompted for a login like this 

![server login](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Red-Rabbit_server_main_login.png)

the server username is -> `15 15 18 19 22 19 7 13 4 12 23 23 22 16 15 26 4 13 26 14 23 26 8 22 19 7`

the server password is -> `15 15 18 19 22 19 7 13 4 12 23 20 13 18 16 15 26 4 8 18 13 26 14 8 18 19 7 23 26 8 4 12 19`

when you get into the server you will see a few things, tabs, and indexes on the bottom and sides here is an example of the home page 


![server login](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Red-Rabbit_server_main_dashboard.png)

the web server has a few things, a interactive user information dashboard for system info, a errors tab, flags tab, tab for a link to my discord server, my instagram, the red rabbit github, and as well as a home tab. The server works pretty smooth and is well working, however the server is a binary file CLOSED SOURCE, this is due to a algorithm used for encryption that me and my development team do not want open source. It is a simple web innterface and i would like it to be more used than the less because this entire server hosts every bit of documentation on every tool and utility in the red rabbit project, here is an example of the errors tab being used 

![Server errors docs](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Red-Rabbit_server_errors-tab.png)


and the system information tab 

![Server system information tab](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/RR6_DEMO_system_info.png)


# Running red rabbit 


Red rabbit is a pretty easy to use script it comes with many utilities and is written in a variety of languages such as Go, C, and C++. With the use of these languages along with utilitity ideas comes some flags, configuration and even settings for the script to run other files and modules properly without fail which in this section we will be talking about the script and using it. There are a few ways to work with red rabbit, configuration files, flags, and raw input.


> Flags 

I wont be specifying every flag in this script because thats for the interface and documentation server to handel. If you need help with flags you can use --help, --hh to view other flags and options. There are some flags i will go over and one of them helps control input and output for the users screen, say you have a verticle monitor you can use `--reso=""` to specify wether your screen is verticle you would use `--reso="verticle"`the output when you specify a vertical setting will minimalize the terminal banner, the output, and even some really large output. here is an example of the banner when its specified as verticle or landscape 


- verticle

```
 ___ ___ _| |___ ___ ___| |_| |_|_| |_   .               .
|  _| -_| . |___|  _| .'| . | . | |  _| .´  ·  .     .  ·  `.
|_| |___|___|   |_| |__,|___|___|_|_|   :  :  :  (¯)  :  :  :
                                        `.  ·  ` |-| ´  ·  .´
                                         '       |-|       '                                     
                                                /|-|\
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^/ |+| \^^^^^^
RR6>
```

- landscape 

```
┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│    ___o .--.               ____                                          ____                                      __                      ___o .--.    │
│   /___| |--|              /\  _`\                                       /\  _`\                                 __/\ \__                  /___| |OO|    │
│  /'   |_|  |_             \ \,\L\_\    ___     __     _ __    __        \ \,\L\_\     __    ___   __  __  _ __ /\_\ \ ,_\  __  __        /'   |_|  |_   │
│       (_    _)             \/_\__ \   /'___\ /'__`\  /\`'__\/'__`\  __o__\/_\__ \   /'__`\ /'___\/\ \/\ \/\`'__\/\ \ \ \/ /\ \/\ \           (_     _)  │
│       | |   \                /\ \L\ \/\ \__//\ \L\.\_\ \ \//\  __/    |    /\ \L\ \/\  __//\ \__/\ \ \_\ \ \ \/ \ \ \ \ \_\ \ \_\ \           | |   \   │
│       | |___/                \ `\____\ \____\ \__/.\_\\ \_\\ \____\  / \   \ `\____\ \____\ \____\\ \____/\ \_\  \ \_\ \__\\/`____ \          | |___/   │
│                               \/_____/\/____/\/__/\/_/ \/_/ \/____/  _______\/_____/\/____/\/____/ \/___/  \/_/   \/_/\/__/ `/___/> \                   │
│                                                                     /\______\                                                  /\___/                   │
│                                                                     \/______/                                                  \/__/                    │
│                                                                                                                                                         │ 
│           Professional Digital forensics, Network hacking, Stegonography, Recon, OSINT, Bluetooth, CAN and Web Exploitation Expert Secruity Team        │
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

RR6>
```

- shark 

```
                    /""-._
                  .      '-,
                  :         '',
                  ;      *     '.
                  ' *         () '.
                   \               \
                    \      _.---.._ '.  
                     :  .' _.--''-''  \ ,'                        ______           __        ______         __     __     __ __   
       .._            '/.'             . ;                       |   __ \.-----.--|  |______|   __ \.---.-.|  |--.|  |--.|__|  |_ 
        ; `-.          ,     no         \'                       |      <|  -__|  _  |______|      <|  _  ||  _  ||  _  ||  |   _|
         ;   `,         ;    F         ._\                       |___|__||_____|_____|      |___|__||___._||_____||_____||__|____|
          ;    \     _,-'    L           ''--._                                     
           :    \_,-'        A                 '-._              Next generation Offensive Security Framework For Advanced Hackers
            \ ,-'            G          .          '-._                             Powered by Fortran, Go, and Perl
           .'         __.-''; detected  \...,__       '.        
          .'      _,-'       \              \   ''--.,__ '\     
         /   _,--' ;          \             ;           "^.}    
        ;_,-' )     \  )\      )            ;                  
             /       \/  \_.,-'             ;
            /                              ;     _____________________
         ,-'  _,-'''-.    ,-.,            ;      [ Scare Sec Hacker's] Aka: The legacy hackers
      ,-' _.-'        \  /    |/'-._...--'
     :--``             )/
---------------------------------------------------------------------------------------------------------------------------------------------------------
RR6>
```

flags like this will certianly help you, if you need more information on flags be sure to visit the flags segment of the documentation server 

> Configuration files 

In this new version of red rabbit, it was decided to add more use modification ability to the script, both to the code and to how settings and options operate so we added conf/configuration options in the script, when you load red rabbit and say help which will output all utilities and commands, you are quickly going to notice some options have `custom` or `config` at the very end of the command, this means the command is either loaded through config files for settings or uses custom user input say for searching something in a PCAP file.

the first example is the command `search port custom` in the filepath of `RR6/config/scanconf` there is a file labled `settings.yaml` this file will set the standard host to scan, port to start from, and port to end at where the standard is 1-65535, this allows the user to customize and save what target they would like to scan, the file looks like this 

```yml
Data:
  Hostname: "10.0.0.1"
  Port_Start: "1"
  Port_End: "65535"
```

configuration files like these also fall in formats of JSON, TXT, INI, and CONF this just allows certian programs to parse data such as API keys and even targets in a variety of ways, more information on this in the web interface.

> Raw Input 

Most commands in red rabbit ask for input, that is however if you dont use flags. Not all input or tools will be able to use flags, the project does not fully impliment flag features YET however they are still widly used in modules like `[parse, inject, dump, ping, crack, encode, gen, utils, db, ssh, ftp ]` etc. If you do not specify a flag like a filepath flag or interface flag for network gathering then you will most likely be prompted for user input, there is also information in the web interface about flags and how they work. there is not much to say about the raw user input other than it will be asked for if data is not set in config files, json files, settings files, flags, or other types.

> Running the script and looking for certian options 

Out of the entire red rabbit project there is a total of 24 modules which includes search, help, encode, decode, fetch, stalk, inject, dump, parse, run, brute, start, stop, check, utils, db, ssh, api, engine, and atk. These modules have a subset like a tree of modules an example is the following tree 

```
| - engine 
     |> engine ogoogle
     |> engine otwitter
     |> engine ofacebook
     |> engine olinkedin
     |> engine cve
     |> engine tcve
     |> engine shodanh 
```

each module a said above has a subset of tools and that command starts with the module name and ends with its attribute such as `engine cve`, alot of people do not like verbose output or large amounts of output, so to minimalize the script i added help options for each module, its syntax is `help  <module_name>` like this 


```
RR6> help inject 


                    /""-._
                  .      '-,
                  :         '',
                  ;      *     '.
                  ' *         () '.
                   \               \
                    \      _.---.._ '.  
                     :  .' _.--''-''  \ ,'                        ______           __        ______         __     __     __ __   
       .._            '/.'             . ;                       |   __ \.-----.--|  |______|   __ \.---.-.|  |--.|  |--.|__|  |_ 
        ; `-.          ,     no         \'                       |      <|  -__|  _  |______|      <|  _  ||  _  ||  _  ||  |   _|
         ;   `,         ;    F         ._\                       |___|__||_____|_____|      |___|__||___._||_____||_____||__|____|
          ;    \     _,-'    L           ''--._                                     
           :    \_,-'        A                 '-._              Next generation Offensive Security Framework For Advanced Hackers
            \ ,-'            G          .          '-._                             Powered by Fortran, Go, and Perl
           .'         __.-''; detected  \...,__       '.        
          .'      _,-'       \              \   ''--.,__ '\     
         /   _,--' ;          \             ;           "^.}    
        ;_,-' )     \  )\      )            ;                  
             /       \/  \_.,-'             ;
            /                              ;     _____________________
         ,-'  _,-'''-.    ,-.,            ;      [ Scare Sec Hacker's] Aka: The legacy hackers
      ,-' _.-'        \  /    |/'-._...--'
     :--``             )/
---------------------------------------------------------------------------------------------------------------------------------------------------------
   ┌────┳────────────────────┳─────────────────────────────────────────────────────────────────────────┐
   │ C  ┃ inject zip         ┃ Inject a ZIP file / Archive file into a JPG/JPEG image                  │
   │ C  ┃ inject gif         ┃ Inject a given payload into a GIF image                                 │
   │ C  ┃ inject jpg         ┃ Inject a given payload into a JPG/JPEG image                            │
   │ C  ┃ inject webp        ┃ Inject a given payload into a WEBP image                                │
   │ C  ┃ inject png         ┃ Inject a given payload into a PNG image                                 │
   │ C  ┃ inject bmp         ┃ Inject a given payload into a BMP image                                 │
   │ C  ┃ inject payload     ┃ Inject a given payload into a URL to parse it into the headers          │
   │ C  ┃ inject payloadl    ┃ Inject a given payload into a URL from a list of URLS to parse payloads │
   │ C  ┃ inject payloadul   ┃ Inject a given payload list into a URL to parse it into the headers     │
   └───────────────────────────────────────────────────────────────────────────────────────────────────┘
```

this is an example of printing out ALL commands that relate to the branch or root of the module `inject` notice how inject never changes but its attribute changes? thats how navigation works around red rabbit, you can do this with most modules in red rabbit however some extra modules like `atk` or `stalk` are not big enough to fit in a help menu so that was decided to not be added.

# Why should i use red rabbit 

> Open source 

Most frameworks out there even smaller ones that only have 1-5 utilities are very closed source, ones that use API's will claim they use raw databases but they dont they infact will encode it and keep you from seeing wether they are lying or not but when you go to a graphical network mapper you can see clearly your network is being sent to a API, sure Red Rabbit uses few API's but it is very honest with you in terms of where your connection is going and how the responses are being formatted 

> Raw offline databases 

Most scripts that require tracing EG phone numbers, port numbers, XML parsers, XML writers, XML databases, or anything that is offline is all a locally hosted database and what i mean by that is simply a very large amount of information in things such ad the proxy JSON files, the phone number json files, the port number json files and the other files that all are counted as offline. Not all commands are offline but when they are online and are OSINT commands they are quite obviously api by telling the user how true these API's are and letting the module be a API in itself 

> Truthful 

Most frameworks today that peoople make on github say `raw free source code wifi framework from scratch` or `raw password brute forcing framework` when in reality these frameworks are just a bunch of OS commands that capture the output from tools like aircrack-ng, or hydra which is good if you just need a second hand for output but is bad when you lie to people who are using your script when you say raw code and its not really raw code its just automation of another persons code.

> Fast and reliable 

Red rabbitsa previous versions were NOT reliable but for version 1.0 it now comes with organized modules for things like recon as well as threaded chains for heavy processes for resource saving. Modules like Brute, search, utils, db, ssh are all threaded and well written with saving code that does not take as long to make responses or find data unless a target is malformed 

> Easy to use 

Red rabbit unlike most of the frameworks out there come easy to use with a out of the box working documentation server and a heavy support behind its tools, it comes easy to use even coming down to flags, is highly descriptive of its utilities, and is written with pride.

> Professional and respected 

Red Rabbit is a highly professional script and may not seem like the average better framework for people who use something like metasploit (since for now it cant even come close to MSF for payload delivery and other things) however it manages to keep it professional with much options and a diverse scheme of output for the user, it also managed to keep processes light, on the low, and stealthy by allowing the user to control how settings or functions may work in the script such as using configuration files.

> Raw code no automation

All code in the red rabbit project is written pure, or auto generated by template based AI generators which make sure that code is efficient enough to keep going without breaking or having issues in the script such as panics, EOF errors, flag errors, parsing errors etc. Red rabbit holds high respect for its competitors and even inspirations as well when running this code and its functions. Note that no tool, function, option, or program in red rabbit is automating tools like aircrack, sqlmap, sstimap, NoSQLmap, hydra, or any other high end program, it keeps things real and simple for a whole new legacy enviroment.

> Modernized systems 

Code used in red rabbit is modern despite having a legacy style to it. Red Rabbit itself likes to give the user most control over the options and functions inside of the program, control over targets, infromation to send out, requests that are made, workers and routines that are made, configuration and API keys, as well as filtering certian data when parsing through packet captures or raw text files to find certian information or even find info on a file they shall specfically request.


# Configuration for I/O console

as said previously there are ways to work around red rabbit with things such as flags, commands, input, or config files. However there are some other little things to customize output, by defualt every time you enter a command the screen will clear and run the function. however if you would like to keep your history you can simple go to the path `config/io` and click on the file named `IO-Settings.yaml` which has a structure like the following 

```yaml
IO_Setting:
  Clear_On_Command: true
```

When you install red rabbit this will be set to true, if you do not want a large output keep it at true. if you want to not clear the information from other commands, say from something like the engine and search modules and want to use the information from other modules as input for other tools set this to false.

*false -> Do not clear on command, keep history up until terminal output buffer*

*true  -> Clear after entering a command, falls for all 270+ commands*

# Demos

> pcap image parsing and downloading 

**NOTE WHEN RUNNING THIS TOOL THE OUTPUT FILEPATH MUST HAVE A / AT THE END OF THE PATH**

![pcap parsing images](https://github.com/ArkAngeL43/Red-Rabbit-V4/blob/main/git/demo_pcap_Image_downloading.png)

> International phone number checking `Raw and offline database`

![information](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_check_number_info.png)

> Digital forensics with DLL/PE formats 

![PE Digital forensics](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/dump_DLL.png)

> Stegonography with formats such as GIF, JPG, JPEG, JFIF, PNG, etc 

![GIF Stegonography](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_inject_gif.png)

> Credential dumping / pcap parsing and dumping

- SIP AUTH dumping 

![SIP AUTH](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_pcap_SIP_Authentication_message_dumping.png)

- FTP Auth dumping 

![FTP AUTH](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_pcap_FTP_Dumping.png)

> Database utilities 

![DB Util Ping](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_DB_utils.png)

> IP Information 

![IP Information](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_IP_check.png)

> Offline port information database 

![Port information](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Port_Info_Database.png)

> Offline CVE information database 

![CVE information](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_CVE_Info.png)

> Google search engine for information 

![Engine](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_Engine.png)

> Username finder

![Username OSINT](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_User_Search.png)

> Phone Number Information 

![Offline phone number and hash information US](https://github.com/ArkAngeL43/Readme_RR/blob/main/git/Demo_US_Number_Information.png)

# Pre Generation 

Red rabbit is a big project, it has files that are thousands of lines of code of just types, colors, or module settings. A good thing to note would be that those files like some color modules like in this repo  `https://github.com/ArkAngeL43/pre-generated-colors.git` are all pre generated, a majority of the red rabbit project like types and constants and even some files like the SQL modules are all auto generated code, this code should not be touched unless you are a contributing developer of this project. This generator is based on templates created by the offical developer of the red rabbit project, meaning that those files are property of the red rabbit project and the developer. Touching these files, editing them, correcting warnings or errors that may be alerted in a text editor can be false positives and cause massive errors inside of the project, it is asked of the developer that you leave these files alone unless you are a developer or contributing member of the project.


# Legal disclaimer 

Red rabbit is a project that is capeable of knocking networks offline, taking down or sending malicous requests to hosts and networks, and even scanning web applications with a bad intent to do damage. The developers ask that you please use this framework responsibly and not use it to cause damage unless you have the right to do so, or you are using it for educational purposes. Developers are NOT RESPONSIBLE FOR ANY DAMAGES YOU MAKE WITH THIS FRAMEWORK OR LAWS YOU MAY BREAK. 

# command list 

```
  ┌────┳────────────────────┳─────────────────────────────────────────────────────────────────────────┐
  │ C  ┃ help verified      ┃ View all verified and working commands                                  │
  │ C  ┃ commands/h         ┃ View this help menu, sjowing general info                               │
  │ C  ┃ help flags         ┃ View all descriptions of flags                                          │
  │ C  ┃ help advanced      ┃ View all commands with advanced usages and examples                     │
  │ C  ┃ help search        ┃ View all commands for the search module                                 │
  │ C  ┃ help inject        ┃ View all commands for the stegonography module                          │
  │ C  ┃ help parse         ┃ View all commands for the file parsing module                           │
  │ C  ┃ help encode        ┃ View all commands for the text encoding module                          │
  │ C  ┃ help utils         ┃ View all commands for the utilities and etc modules                     │
  │ C  ┃ help brute         ┃ View all commands for the brute forcing modules                         │
  │ C  ┃ help dump          ┃ View all commands for the file dumping modules                          │
  │ C  ┃ help ping          ┃ View all commands for the network discovery modules                     │
  │ C  ┃ help sniff         ┃ View all commands for the network / 802.11 sniffing module              │
  │ C  ┃ help check         ┃ View all commands for the OSINT / verification modules                  │
  │ C  ┃ help commands      ┃ View all general commands and descriptions                              │
  │ C  ┃ help all           ┃ View all general commands and descriptions from config files            │
  │ C  ┃ script settings    ┃ View all settings/flag values that was used                             │
  │ C  ┃ clear              ┃ Clear terminal output                                                   │
  │ C  ┃ cls                ┃ Clear terminal output                                                   │
  │ C  ┃ exit               ┃ Exit the script                                                         │
  │ C  ┃ time               ┃ Get script time                                                         │
  │ C  ┃ stalk mac          ┃ Trace the OUI of a single mac, only one mac can be checked at a time    │
  │ C  ┃ stalk oui          ┃ Trace the OUI of every single mac in a list of macs                     │
  │ C  ┃ stalk hosts        ┃ Trace all hosts on a network given the first 3 octets (10.0.0)          │
  │ C  ┃ stalk hosts2       ┃ Trace all hosts on a network given the first 2 octets (10.0.)           │
  │ C  ┃ make GET           ┃ make a GET request to a given URL and                                   │
  │ C  ┃ make POST          ┃ make a POST request with custom data to a given URL and output body     │
  │ C  ┃ make PUT           ┃ make a PUT request with custom data to a given URL and  output body     │
  │ C  ┃ stalk hosts2       ┃ Trace all hosts on a network given the first 2 octets (10.0.)           │
  │ C  ┃ stalk hosts2       ┃ Trace all hosts on a network given the first 2 octets (10.0.)           │
  │ C  ┃ search scan        ┃ Get all resolved hosts based on the first three octets of an IP (0.0.0) │
  │ C  ┃ search head        ┃ Get the headers of a URL                                                │
  │ C  ┃ fetch titles       ┃ Get the direct title of a website                                       │
  │ C  ┃ fetch documents    ┃ Get all documents on the page (paths only)                              │
  │ C  ┃ fetch links        ┃ Get all links found in the HTML page                                    │
  │ C  ┃ search redirect    ┃ Check if a target URL will redirect you via status code                 │
  │ C  ┃ search subdomains  ┃ Get all subdomains based on a subdomain list from a given domain        │
  │ C  ┃ search ns          ┃ Get the name servers of a domain name                                   │
  │ C  ┃ search mx          ┃ Get all the MX records from a domain name                               │
  │ C  ┃ search waf         ┃ Get all blocked services based on a payload ex (<script>alert()<script>)│
  │ C  ┃ search envf        ┃ Check if targets in a given host list has a .env endpoint to it         │
  │ C  ┃ search env         ┃ Check if a single target has a .env file located in the header          │
  │ C  ┃ search port        ┃ Get all protocals, types, services, etc on a single port number offline │
  │ C  ┃ search port l      ┃ Get all protocals, types, services, etc on a list of ports      offline │
  │ C  ┃ search port live   ┃ Get all ports on a specified host (EX: 10.0.0.1)                        │
  │ C  ┃ search phpbb       ┃ Scan the URL for the version of PHPBB                                   │
  │ C  ┃ search port llive  ┃ Get all active ports on a LIST of hosts (EX: Hostwordlist.txt)          │
  │ C  ┃ search port custom ┃ Get all ports on a host using custom YAML configuration files           │
  │ C  ┃ search endpoint    ┃ Get the endpoint of a URL                                               │
  │ C  ┃ search ip          ┃ Get all hostnames of an IP address                                      │
  │ C  ┃ search txt         ┃ Get all TXT records of a domain name                                    │
  │ C  ┃ search hostname    ┃ Get the hostname of an IP                                               │
  │ C  ┃ search cname       ┃ Get the CNAME records of a domain name                                  │
  │ C  ┃ search srv         ┃ Get the SRV of a domain name                                            │
  │ C  ┃ search robots      ┃ Get the robots.txt from a given URL                                     │
  │ C  ┃ search urls        ┃ Get all URLS, IPAS, Domain names, and servers of a URL that is crawled  │
  │ C  ┃ search banner      ┃ Get the banner of a hostname by making a TCP dial                       │
  │ C  ┃ search server      ┃ Get the server of a domain name or URL                                  │
  │ C  ┃ search ipa         ┃ Get the IP Address of a URL                                             │ 
  │ C  ┃ search route       ┃ Get the route or run a traceroute scan on a domain/url/host             │
  │ C  ┃ search admin       ┃ Scan the URL for admin panels or admin access                           │
  │ C  ┃ search sqli        ┃ Scan the URL for SQL injection vulnerabilities                          │
  │ C  ┃ search lfi         ┃ Scan the URL for LFI vulnerabilities                                    │
  │ C  ┃ search sig         ┃ Scan a unknown file to try and find its filetype                        │
  │ C  ┃ search filepath    ┃ Scan a filepath for a certian file                                      │
  │ C  ┃ search archive     ┃ Scan a JPG image for a archive file and try to unzip it                 │
  │ C  ┃ search geo         ┃ Scan a JPG/JPEG file for a GEO location                                 │
  │ C  ┃ search methods *   ┃ Scan through multiple request methods to see what a URL will accept     │ 
  │ C  ┃ search httpmethod *┃ Print out a list of every HTTP method used in search methods *          │
  │ C  ┃ search knoxss custo┃ `search knoxss custom` scan for xss using the custom api config file    │ 
  │ C  ┃ search xss knoxl   ┃ Test all urls in a list for xss vulnerabilities using the KNOXSS API    │ 
  │ C  ┃ search xss knoxss  ┃ Test a url for xss vulnerabilities using the KOXSS API                  │ 
  │ C  ┃ inject zip         ┃ Inject a ZIP file / Archive file into a JPG/JPEG image                  │
  │ C  ┃ inject gif         ┃ Inject a given payload into a GIF image                                 │
  │ C  ┃ inject jpg         ┃ Inject a given payload into a JPG/JPEG image                            │
  │ C  ┃ inject webp        ┃ Inject a given payload into a WEBP image                                │
  │ C  ┃ inject png         ┃ Inject a given payload into a PNG image                                 │
  │ C  ┃ inject bmp         ┃ Inject a given payload into a BMP image                                 │
  │ C  ┃ inject payload     ┃ Inject a given payload into a URL to parse it into the headers          │
  │ C  ┃ inject payloadl    ┃ Inject a given payload into a URL from a list of URLS to parse payloads │
  │ C  ┃ inject payloadul   ┃ Inject a given payload list into a URL to parse it into the headers     │
  │ C  ┃ dump bootsec info  ┃ Dump your bootsec filename info                                         │ 
  │ C  ┃ dump file info     ┃ Dump normal file inforomation on a file                                 │
  │ C  ┃ dump file info     ┃ Dump normal file inforomation on a file                                 │
  │ C  ┃ dump file binary   ┃ Dump a given file into a hex/binary like output                         │
  │ C  ┃ dump pe info       ┃ Dump a PE/DLL file and get all of its general information and sections  │
  │ C  ┃ dump image metadata┃ Dump an image's metadata                                                │
  │ C  ┃ dump pcap dot11    ┃ Dump all BSSID's, SSID's, and macs in a DOT11 pcap file                 │
  │ C  ┃ dump pcap ftp      ┃ Dump all found FTP username and passwords / credentials in a PCAP file  │
  │ C  ┃ dump pcap ospf     ┃ Dump all found OSPF authentication bytes in a PCAP/CAP file             │
  │ C  ┃ dump pcap smtp     ┃ Dump all authentication/users/passwords in a normal SMTP PCAP/CAP file  │
  │ C  ┃ dump pcap smtppa   ┃ Dump all AUTH PLAIN logins made in a pcap file                          │
  │ C  ┃ dump pcap smtpe    ┃ Dump all email chats, users, and bodies in a PCAP file                  │
  │ C  ┃ dump pcap sipa     ┃ Dump all users and passwords for SIP authentication in a PCAP file      │
  │ C  ┃ dump pcap sipok    ┃ Dump all +OK or OK responses in a SIP PCAP/CAP                          │
  │ C  ┃ dump pcap sipin    ┃ Dump all INVITES in a SIP PCAP/CAP file                                 │
  │ C  ┃ dump pcap sipreg   ┃ Dump all Registers made in a SIP PCAP/CAP file                          │
  │ C  ┃ dump pcap sippg    ┃ Dump all POST and GET requests made in a SIP PCAP/CAP file              │
  │ C  ┃ dump pcap images   ┃ Dump all and download all images that have been found inside of a pcap  │
  │ C  ┃ dump pcap imaplogn ┃ Dump all IMAP logins made in a IMAP PCAP/CAP file                       │
  │ C  ┃ dump pcap custom   ┃ Dump your own custom filters in any PCAP file self output               │
  │ C  ┃ dump pcap          ┃ Dump all packets in a PCAP/CAP file UNFORMATTED                         │
  │ C  ┃ ping udp           ┃ Send out UDP pakcets to ping all living hosts                           │
  │ C  ┃ ping tcp           ┃ Send out TCP packets to ping all living hosts                           │
  │ C  ┃ ping syn           ┃ Send out SYN packets to ping all living hosts                           │
  │ C  ┃ ping icmp          ┃ Send out ICMP packets to ping all living hosts                          │
  │ C  ┃ ping arp           ┃ Send out ARP packets to identify living hosts by MAC and IP             │
  │ C  ┃ crack sha1 list    ┃ Brute force all SHA1 hashes in a list                                   │
  │ C  ┃ crack sha256 list  ┃ Brute force all SHA256 hashes in a list                                 │
  │ C  ┃ crack md5 list     ┃ Brute force all MD5 hashes in a list                                    │
  │ C  ┃ crack sha1 single  ┃ Brute force a single SHA1 hash                                          │
  │ C  ┃ crack sha256 single┃ Brute force a single SHA256 hash                                        │
  │ C  ┃ crack md5 single   ┃ Brute force a single MD5 hashes                                         │
  │ C  ┃ encode md5         ┃ Encode a string in a MD5 hash                                           │
  │ C  ┃ encode sha1        ┃ Encode a string in a SHA1 hash                                          │
  │ C  ┃ encode sha256      ┃ Encode a string in a SHA256 hash                                        │
  │ C  ┃ encode sha512      ┃ Encode a string in a SHA512 hash                                        │
  │ C  ┃ encode base64      ┃ Encode a string in a BASE64 string                                      │
  │ C  ┃ encode base32      ┃ Encode a string in a BASE32 string                                      │
  │ C  ┃ encode md5  list   ┃ Convert all passwords in a wordlist MD5 hashes                          │
  │ C  ┃ encode sha1 list   ┃ Convert all passwords in a wordlist into SHA1 hashes                    │
  │ C  ┃ encode sha256 list ┃ Convert all passwords in a wordlist into SHA256 hashes                  │
  │ C  ┃ encode sha512 list ┃ Convert all passwords in a wordlist into SHA512 hashes                  │
  │ C  ┃ encode base64 list ┃ Convert all passwords in a wordlist into base64 strings                 │
  │ C  ┃ encode base32 list ┃ Convert all passwords in a wordlist into base32 strings                 │
  │ C  ┃ encode rot13       ┃ Encode a string in ROT13                                                │
  │ C  ┃ encode HMAC        ┃ Encode a string with a key in HMAC                                      │
  │ C  ┃ run RR6 GUI        ┃ Run the RR6 GUI for image injection, OUI tracing and more               │
  │ C  ┃ run RR6 scan gui   ┃ Run the RR6 scan GUI for network recon                                  │
  │ C  ┃ Brute SMTP         ┃ Brute force SMTP services and emails                                    │
  │ C  ┃ Brute HTTP auth    ┃ Brute force HTTP authentication threaded option                         │
  │ C  ┃ Brute SSH          ┃ Brute force SSH Servers and usernames                                   │
  │ C  ┃ Brute FTP          ┃ Brute force FTP servers and usernames                                   │
  │ C  ┃ Brute HTTPA        ┃ Brute force HTTP plain authentication                                   │
  │ C  ┃ Brute Telnet       ┃ Brute force TELNET servers using std auth                               │ 
  │ C  ┃ Brute Cpan         ┃ Brute force Cpanel servers using std auth                               │ 
  │ C  ┃ Brute psql         ┃ Brute force PostGreSQL servers using a wordlist                         │ 
  │ C  ┃ Brute mysql        ┃ Brute force MySQL servers using a wordlist                              │ 
  │ C  ┃ Brute mongo        ┃ Brute force Mongo servers using a wordlist                              │ 
  │ C  ┃ Brute ZIP          ┃ Brute force ZIP files with a given thread count, wordlist, and ZIP file │ 
  │ C  ┃ sniff interfaces   ┃ Scan for interfaces on the current machine                              │ 
  │ C  ┃ sniff iface info * ┃ Scan for interfaces on the current machine and grab and output card info│ 
  │ C  ┃ sniff iface info   ┃ Scan for all information on a certian interface on the current system   │ 
  │ C  ┃ sniff interface ip ┃ Scan for all information on a certian interface based on a given IP     │  
  │ C  ┃ sniff tcp          ┃ Sniff all incoming TCP packets that are picked up                       │ 
  │ C  ┃ sniff ip           ┃ Sniff all IP packets that are picked up                                 │ 
  │ C  ┃ sniff dhcp         ┃ Sniff all DHCP packets that come through                                │ 
  │ C  ┃ sniff application  ┃ Sniff all application like packets like Multicast Query Response        │ 
  │ C  ┃ sniff ethernet     ┃ Sniff all Ethernet packets that are picked up on a interface            │
  │ C  ┃ sniff ssid         ┃ Sniff all networks nearby by parsing dot11 beacon requests              │
  │ C  ┃ sniff probe        ┃ Sniff all data of a network by parsing all incoming dot11 probe requests│
  │ C  ┃ sniff 802raw       ┃ Sniff all dot11 packets, and all layers without formatting              │
  │ C  ┃ sniff discovery    ┃ Sniff all networks nearby and get info like RSSI, BSSID, OUI, MAC, SSID │
  │ C  ┃ check cpanel vuln  ┃ Check if the version of cpanel you are using is vulnerable              │
  │ C  ┃ check proton ip    ┃ Check if a IP address belongs to a proton mail IP server                │
  │ C  ┃ check IP           ┃ Check to look for general information of an IP, EG: Country, code       │
  │ C  ┃ check github user  ┃ Check to look for information on a github user, EG: Bio, Location etc   │
  │ C  ┃ check github star  ┃ Check to find all stargazers on a given users repo and info of gazers   │
  │ C  ┃ check github foll  ┃ Check to find all information on followers on a certian github user     │
  │ C  ┃ check github orgr  ┃ Check to find all organization repositories that are public             │
  │ C  ┃ check github usero ┃ Check to find all organizations that the user is tied to by a git user  │
  │ C  ┃ check github usere ┃ Check to find all user events based on a given github username          │
  │ C  ┃ check github orge  ┃ Check to find all organization events based on an organization name     │
  │ C  ┃ check github repos ┃ Check to find all repos and repo information of a given github username │
  │ C  ┃ check proton email ┃ Check if a email address belongs to a proton mail account               │
  │ C  ┃ check cloudflare ip┃ Check if a given IPv6 or IPv4 address belongs to the cloudflare CIDRs   │
  │ C  ┃ check cloudfront ip┃ Check if a given IPv6 or IPv4 address belongs to the cloudfront CIDRs   │
  │ C  ┃ check mcafe ip     ┃ Check if a given IPv6 or IPv4 address belongs to the Mcafe CIDRs        │
  │ C  ┃ check aws ip       ┃ Check if a given IPv6 or IPv4 address belongs to the AWS CIDRs          │
  │ C  ┃ check azure ip     ┃ Check if a IP address belongs to the azure cloud addr prexix list       │
  │ C  ┃ check myip         ┃ Check for your public IP address                                        │
  │ C  ┃ check number       ┃ Check a number hash EX(+381-##-###-####) and try to get information     │
  │ C  ┃ check number us    ┃ Check a number code EX(320) and get city, state, currency etc           │
  │ C  ┃ check number be    ┃ Check a number hash EX(+32(2)###-##-##) and get city, country etc info  │
  │ C  ┃ trace number us    ┃ Check a US number for information and data                              │
  │ C  ┃ parse burp targets ┃ Parse a list of urls and targets to produce a burp suite config file    │
  │ C  ┃ parse config       ┃ Parse the api database to see what API keys you have set in the config  │
  │ C  ┃ parse xmln host    ┃ Parse a NMAP XML file to find the scanned host                          │
  │ C  ┃ parse xmln service ┃ Parse a NMAP XML file to find info on all ports and their services      │
  │ C  ┃ parse xmln ports   ┃ Parse a NMAP XML file to find all scanned ports and their state         │
  │ C  ┃ parse xmln hostn   ┃ Parse a NMAP XML file to find all hostnames in the scan                 │
  │ C  ┃ parse xmln porti   ┃ Parse a NMAP XML file to find all ports and their information           │
  │ C  ┃ parse xmln runs    ┃ Parse a NMAP XML file to find the runtime and exit status of nmap       │
  │ C  ┃ parse xmln hosts   ┃ Parse a NMAP XML file to find all the scanned number of hosts           │
  │ C  ┃ parse xmln hosth   ┃ Parse a NMAP XML file to find all info in the host hint section         │
  │ C  ┃ parse xmln verbose ┃ Parse a NMAP XML file to find settings for the verbose tag              │
  │ C  ┃ parse xmln debug   ┃ Parse a NMAP XML file to find settings for the debugging tag            │
  │ C  ┃ parse xmln times   ┃ Parse a NMAP XML file to find times of the script                       │
  │ C  ┃ parse xmln scaninf ┃ Parse a NMAP XML file to find all scan information                      │
  │ C  ┃ parse xmln *       ┃ Parse a NMAP XML file and run all xmln options for output               │
  │ C  ┃ parse scope        ┃ Parse a burp scope json file to match with hosts in a url like urls.txt │
  │ C  ┃ parse data IP4     ┃ Parse through an entire file to find all IPv4 addresses used            │
  │ C  ┃ parse data IP6     ┃ Parse through an entire file to find all IPv6 addresses used            │
  │ C  ┃ parse data email   ┃ Parse through an entire file to find all email addresses used           │
  │ C  ┃ parse data mac     ┃ Parse through an entire file to find all MAC addresses used             │
  │ C  ┃ parse data custom  ┃ Parse through an entire file with a custom regex string to pillage data │
  │ C  ┃ utils ghosted      ┃ Run the ghosted OS shell for resetting, deleting drives                 │
  │ C  ┃ utils blocker l    ┃ Run the DNS blocker and block a list of hosts from making a connection  │
  │ C  ┃ utils blocker s    ┃ Run the DNS blocker and block a single host from making a connection    │
  │ C  ┃ utils blocker fix  ┃ Run the OS module to repair/reset the hosts file after writing to file  │
  │ C  ┃ utils verify mac   ┃ Verify a MAC address with regex to test the macs quality                │
  │ C  ┃ utils verify email ┃ Verify a Email address with regex to test the emails quality            │
  │ C  ┃ utils verify IP4   ┃ Verify a IPv4 address with regex to test the IP Addresses quality       │
  │ C  ┃ utils verify IP6   ┃ Verify a IPv6 address with regex to test the IP6 Addresses quality      │
  │ C  ┃ utils verify h:p   ┃ Verify a host and port with regex to test the host:port quality         │
  │ C  ┃ utils verify host  ┃ Verify a hostname address with regex to test the hosts quality          │
  │ C  ┃ utils gen emailmail┃ Scrape a website for emails and parse it into a maltego like body report│
  │ C  ┃ utils gen linkmail ┃ Scrape a website all links / hrefs and parse it into a maltego body     │
  │ C  ┃ utils verify host  ┃ Verify a hostname address with regex to test the hosts quality          │
  │ C  ┃ utils gen JPG      ┃ Generate a map / GPS image based on the geo location of a JPG image     │
  │ C  ┃ utils gen map      ┃ Generate a map / GPS image based on the LAT and LON of a input          │
  │ C  ┃ utils gz uncompress┃ Uncompress a GZIP file or folder, this is experimental so watch out     │
  │ C  ┃ utils unzip        ┃ Unzip every file in a zipped file and output to a proper directory      │
  │ C  ┃ utils build gmapl  ┃ Build a google maps link with a lat (lattitude) and lon (longitude)     │
  │ C  ┃ utils gen AES      ┃ Encrypt a string with a assigned or given AES key length must be 32     │
  │ C  ┃ utils dec AES      ┃ Decrypt a string or message in a filename with a given filename         │
  │ C  ┃ utils encf AES     ┃ Encrypt a given text file with a AES key or string of length 32         │
  │ C  ┃ utils decf AES     ┃ Decrypt a given text file with a AES key or string of length 32         │
  │ C  ┃ utils gen rjpg     ┃ Generate a random jpg image with random settings and configurations     │
  │ C  ┃ utils gen xxe php  ┃ Generate a XXE Entity template based on a PHP command                   │
  │ C  ┃ utils gen xxe soap ┃ Generate a XXE Entity template based on a SOAP XXE template and command │
  │ C  ┃ utils gen xxe base ┃ Generate a XXE Entity template based on BASE64 Encoded data             │
  │ C  ┃ utils gen xxe inj  ┃ Generate a normal XXE template with a normal entity                     │
  │ C  ┃ utils gen xxe xinc ┃ Generate a XXE Entity template based on XINC                            │
  │ C  ┃ utils deface htm   ┃ Deface all html files and htm files in a given directory or location    │
  │ C  ┃ utils test binary  ┃ Check if a binary file is vulnerable to any form of ENV vulnerabilities │
  │ C  ┃ utils msfsessions  ┃ Attempt to gather all of your MSF sessions by local and OS variables    │
  │ C  ┃ utils tor stat     ┃ Attempt to check if SOCKS5 is up and running or ready for use           │
  │ C  ┃ utils tor get      ┃ Attempt to download a file from a url with the SOCKS5 proxy             │
  │ C  ┃ ssh check          ┃ Check to see if a SSH server on a host is alive                         │
  │ C  ┃ ssh transfer       ┃ Attempt transfer a file to the ssh server                               │
  │ C  ┃ ssh auth           ┃ Check to see if a username and password will be authenticated           │
  │ C  ┃ ssh auth key       ┃ Check to see if a ssh key will be authenticated with a user and pass    │
  │ C  ┃ ssh exec           ┃ Execute a command over an ssh server with a simple login, user and pass │
  │ C  ┃ ssh auth config    ┃ Attempt to login to a server with the settings in the SSH yaml file     │
  │ C  ┃ ssh dial config    ┃ Attempt to ping a server with the current settings in the ssh yaml file │
  │ C  ┃ ssh check config   ┃ Attempt to check the connection on a server with the config YAML file   │
  │ C  ┃ ssh start shell    ┃ Get a SSH shell on a remote server using a private key and username     │
  │ C  ┃ db ping mongo      ┃ Attempt to ping a MongoDB database running on a remote server           │
  │ C  ┃ db ping mysql      ┃ Attempt to ping a MySQL   database running on a remote server           │
  │ C  ┃ db ping postgres   ┃ Attempt to ping a PostGreSQL database running on a remote server        │
  │ C  ┃ db ping mssql      ┃ Attempt to ping a MS-SQL database running on a remote server            │
  │ C  ┃ db auth mysql      ┃ Attempt to auth a password and user on a local mysql sql server         │
  │ C  ┃ db auth postgres   ┃ Attempt to auth a password and user on a local postgre sql server       │
  │ C  ┃ start http server  ┃ Start the local red rabbit user dashboard and coumentation server :5501 │
  │ C  ┃ start ssh tunel    ┃ Start a SSH tunnel on a remote host, given local addr and port          │
  │ C  ┃ start listen bind  ┃ Start a bind shell on a given address and port to listen for connection │
  │ C  ┃ start listen tcp   ┃ Start a tcp shell on a given address and port to listen for connection  │
  │ C  ┃ start def interface┃ Start the defualt interface based on OS detection as monitor mode       │
  │ C  ┃ start cus interface┃ Start a interface into monitor mode with a given interface name or addr │
  │ C  ┃ stop def interface ┃ Stop a interface that may be in monitor mode using ip link              │
  │ C  ┃ api layerus config ┃ Use the apilayer API to lookup US numbers with custom config            │
  │ C  ┃ api layerbe config ┃ Use the apilayer API to lookup BE/Belgium numbers with custom config    │
  │ C  ┃ api whoisx config  ┃ Use the WHOIS XML API to lookup a website, with custom config           │
  │ C  ┃ api layerus        ┃ Use the apilayer API to lookup US numbers                               │
  │ C  ┃ api layerbe        ┃ Use the apilayer API to lookup BE/Belgium numbers                       │
  │ C  ┃ api whoisxml       ┃ Use the WHOIS XML API to lookup a website                               │
  │ C  ┃ api grab config    ┃ grab all custom placed username, passwords, and api keys                │
  │ C  ┃ engine ogoogle     ┃ Scrape google for infromation given a search query                      │
  │ C  ┃ engine otwitter    ┃ Scrape google for information on a given twitter search query           │
  │ C  ┃ engine ofacebook   ┃ Scrape google for infromation on a given facebook search query          │
  │ C  ┃ engine olinkedin   ┃ Scrape google for infromation on a given linkedin search query          │
  │ C  ┃ engine tcve        ┃ Scrape a govenrment cyber security databases for info on a cve with tor │
  │ C  ┃ engine cve         ┃ Scrape a govenrment cyber security databases for info on a given cve    │
  │ C  ┃ engine shodanh     ┃ Scrape shodan for information on a given hostname or query              │
  │ C  ┃ atk fuzz           ┃ Attempt to fuzz a network service or application to test its buf limit  │
  │ C  ┃ atk timestamp      ┃ Attempt to change the permissions of a given file or folder by ID       │
  │ C  ┃ atk permission     ┃ Attempt to change the timestamp on a given file or folder               │
  │ C  ┃ atk dnsspoof       ┃ Attempt to spoof and poison a given targets dns through ARP spoofing    │
  └───────────────────────────────────────────────────────────────────────────────────────────────────┘
```


# Flags page

[ You used option EHELP, the following below is how to use ]
[ Red Rabbits flags, and how they work in terms of console ]
```
  ┌────┳────────────────────┳─────────────────────────────────────────────────────────────────────────┐
  │ F  ┃ --he  bool         ┃ General help on flags and commands                                      │
  │ F  ┃ --hh  bool         ┃ General help on help commands                                           │
  │ F  ┃ --help bool        ┃ General help on flags                                                   │
  │ F  ┃ --ehelp bool       ┃ Advanced help on flags, commands, help commands, and some flag examples │
  │ F  ┃ --ph int           ┃ Specify image / pixel height when injecting and reconstructing images   │
  │ F  ┃ --pw int           ┃ Specify image / pixel width when injection and reconstructing images    │
  │ F  ┃ --jpgF string      ┃ Set a chunk to inject a JPG image with COM is the main chunk            │
  │ F  ┃ --reso string      ┃ Set a screen resolution <Verticle|Landscape|> for output format         │
  │ F  ┃ --input string     ┃ Set a input file to be injected for stegonography / image injection     │
  │ F  ┃ --output string    ┃ Set a output file or output filename for image manipulation             │
  │ F  ┃ --offset string    ┃ Set a offset to be injected at for image injection / stegonography      │
  │ F  ┃ --payload string   ┃ Set the payload to be used for image injection and other tests          │
  │ F  ┃ --type string      ┃ Set the type of chunk to inject images at such as IEND                  │
  │ F  ┃ --key string       ┃ Set the encryption key for payloads (--ehelp for examples )             │
  │ F  ┃ --filepath string  ┃ Set the general filepath for any file that will be read or used         │
  │ F  ┃ --hashl string     ┃ Set the file of hashes to be used for hash cracking                     │
  │ F  ┃ --wordl string     ┃ Set the file of passwords to be used for brute forcing                  │
  │ F  ┃ --userl string     ┃ Set the file of usernames to be used for brute forcing                  │
  │ F  ┃ --workers int      ┃ Set the amount of go workers for brute force attacks                    │
  │ F  ┃ --interface string ┃ Set the interface to use for network attacks and packet capture         │
  │ F  ┃ --targetm string   ┃ Set the target's mac address for arp poisoning                          │
  │ F  ┃ --targetip string  ┃ Set the target's ip address for arp poisoning                           │
  │ F  ┃ --gatemac string   ┃ Set the target's gateway mac address for arp poisoning                  │
  │ F  ┃ --CIDR / -z string ┃ Set the CIDR to be used for host pinging, scanning, and more            │
  │ F  ┃ --passlen int      ┃ Set the length of a password string to be generated when making lists   │
  │ F  ┃ --target string    ┃ Set the target URL to be set for web attacks such as SQLI, XSS, etc     │
  │ F  ┃ --payloadl string  ┃ Set the file of payloads to be used for XSS, SQLI, Admin panel's etc..  │
  │ F  ┃ --XMLF string      ┃ Set the XML file for NMAP parsing, or other XML file parsing commands   │
  │ F  ┃ --JSONF string     ┃ Set the JSON file for commands that need to parse certian JSON files    │
  │ F  ┃ --PCAP  string     ┃ Set the PCAP file for parsing, this is not used as much as --filepath   │  
  │ F  ┃ --SQ  / S  string  ┃ Set the Set the google search query or thing you would like to search   │ 
  │ F  ┃ --RPP / R  string  ┃ Set the amount of results per google page scraped for the engine module │ 
  │ F  ┃ --PTC / T  string  ┃ Set the amount of google pages to crawl for the engine module           │ 
```


--he, --hh, --help, --ehelp 

are all the same just in different formats and as you are about to figure 
out, --ehelp trys to touch on every flag and its usage properly 



[Flag 1]
    - --PTC / -T

    This flag is used for the engine module, a module that uses google 
    to search for information. if you want to specify the amount of google
    pages the engine module will scrape, then you use this flag with a number 
    say `--PTC=10` means it will scrape 10 pages and return all data found on 
    those 10 pages, -T is another way of specifying --PTC


[Flag 2]
    -> --RPP / -R 

    This flag is also used for the engine module, it allows you to specify 
    the amounts of ResultsPerPage (RPP), specify this with a number only 
    strings or words are not allowed, an example of this usage would be 

    `sudo go run . --RPP=1`

[Flag 3]
    -> --SQ / -S

    This flag like flags 1 and 2 are used for the engine module as well, 
    this flag allows you to specify a certian seach query and have the 
    engine search for all information and links relevating twards that 
    search query. an example usage would be 

    `sudo go run . --SQ="inurl:index.php?id=1"`

[Flag 4]
    -> --PCAP 

    Flag broken, out of order 

[Flag 5]
    -> --JSONF

    Flag broken, out of order

[Flag 6]
    -> --XMLF

    This flag allows you to specify a XML file that needs to be parsed
    the really only option that uses this will be the nmap xmln parsers 
    which will use this flag alot, it will help you in the long run.

    `sudo go run . --XMLF="examples/xml/nmap2.xml"`

[Flag 7]
    -> --payloadl 

    This flag interacts with alot of list / payload list 
    based commands, things like SQL fuzzing, HTTP fuzzing,
    Admin scanning, file scanning, etc. Any option that asks 
    for a payload list will use this like the XMLF option this 
    will be used ALOT throughout red rabbit.

    Example usage 

    `sudo go run . --payloadl="payloads/admin.txt"`

[Flag 8]
    -> --Target 

    This flag is quite unused however if you are planning to use 
    options and modules like the search module or fetch module this 
    flag will help you, example usage is 

    `sudo go run . --target="someurl.com"`


[Flag 9]
    -> --passlen 

    This flag allows you to set a password length or better yet 
    string length / char length for generating password lists for 
    brute forcing attacks

[Flag 10]
    -> --CIDR / z

    This flag allows you to set a given CIDR or network range to scan for 
    using the ping and scan modules, if you want to scan a host range from 

    `10.0.0.1/24` 

    you will specify that in the flag like so 

    `sudo go run -z="10.0.0.1/24"`
    `sudo go run --CIDR="10.0.0.1/24"`

[Flag 11, 12, and 13]
    -> --gatemac, --targetip, --targetm

    All three of these flags must be used together since they are 
    used in only one option and that is the atk dnsspoof, the first option 
    --gatemac needs a gateway mac address as an argument, the --targetip is 
    the ip of the target you want to run the attack on, the targetm is the MAC Address 
    of the target you want to run the attack on.

    example usage 

    `sudo go run . --gatemac="ff:ff:ff:ff:ff:ff", --targetip="0.0.0.0", --targetm="ff:ff:ff:ff:ff:ff"`

[Flag 14]
    -> --interface 

    This flag allows you to specify a interface to sniff devices / packets from 
    a given interface and listen in for certian information using the sniffing 
    module. 

    its example usage is `sudo go run . --interface="wlp5s0"`

[Flag 15]
    -> --workers

    This flag allows you to specify the amount of workers on brute 
    forcing routine or other commands and modules that may start heavy 
    processes, and require workers to run. 

    example usage -> `sudo go run . --workers=200`

[Flag 16, 17, 18]
    -> 
        | --hashl
        | --wordl
        | --userl
    
    All these flags operate the same but are used for different modules.
    for flag `--hashl` you will need to specify a file of hashes something 
    like hashes.txt which is used in the hash cracking module of red rabbit.
    for flag `--wordl` you will need to specify a file of passes something 
    like passes.txt which is used in the brute forcing module of red rabbit.
    for flag `--userl` you will need to specify a file of userns something 
    like useres.txt which is used in the USER search module of red rabbit

[Flag 19]
    -> --filepath 

    --filepath is a flag used alot, this flag allows you to specify a 
    directory to search with, directory to test, or directory to a filename 
    for something like pcap parsing, digital forensics, filepath dumping, file 
    forensics, file stalkers and much more among that list.

    example usage -> `sudo go run . --filepath="some/dir/to/a/file/file.txt"`

[Flag 20, 21, 22, 23, 24, 25, 26, 27, 28]
    -> Flags 
        => --key
        => --type
        => --payload 
        => --offset 
        => --output / o
        => --input / i
        => --decode 
        => --encode

    These flags can all be combined to do certian things, all flags asides --input,
    and --output can only be used in the inject PNG options, this is what these flags 
    can be used for.

    --output / -o   ~> Specify an output of the injected image 
    --input / -i    ~> Specify an input image you want to inject data into

    Advanced side 

    {Decoding a payload}

    if you want to decode a payload that you will inject into the image you will 
    need to use the `--decode` flag with the --key flag. the --key flag is required 
    for this option to run, simply due to the fact that the flag --key will specify 
    the key to the payload you are trying to decode its example usage is 

    `sudo go run main.go --input="image.png" --output="infected.png" --decode --key="someXORKey"`

    {Encoding a payload}

    if you want to encode a payload that you will inject into the image 
    you will need to use the `--encode` flag along with the `--key` flag.
    if you dont know now the `--key` flag is a NEEDED flag for the image 
    injection to start with the encode option, the data must NOT be NIL 
    or empty, if the data is empty it will exit. the key flag specifies
    the XOR key you will use to inject and encode the payload that you 
    will inject into the image. Example usage --

    `sudo go run main.go --input="image.png" --output="infected.png" --encode --key="someXORKey"`

    {Specifying a offset, data chunk, and payload}

    In order to properly inject the png image with the program and module in red rabbit 
    you will need to specify a few main things, as we discussed above a input and output image, 
    but we will also need to specify a data chunk and offset of that data chunk to inject at. 
    if you do not know the offset you want to inject at you can always use red rabbit to dump all 
    offsets and chunk locations using option `dump png meta`. if you have the information you can use 
    the flag `--offset` to specify an offset location, and `--type` to specify a type / chunk type to 
    inject at. in the path examples/stego there is a file named `injected.png` there is a chunk i want to 
    inject which is the IDAT chunk, it is at offset 0x21, if we want to inject this image we can use 
    the flags  we just talked about to specify where we want to inject it but what about data to inject. 
    well if you want  to specify data to inject just use the `--payload` flag to do so as is 
    the following example down below.

    `sudo go run . 
        --input="examples/stego/injected.png" 
        --output="infected_again.png" 
        --type="IDAT" 
        --payload="hello world" 
        --offset="0x21"`


[Flag 29]
    -> --reso

    Reso stands for resolution, for user customization if you have a specific liking for small or 
    large output you can use the --reso flag to specify your monitors display rotation 
    (Landscape/Verticle) are the only formats allowed right now. For a more slimmer build / 
    output use `--reso="Verticle"` for a larger more space wide output use `--reso="Landscape"` or 
    for a certian banner design like the shark use  `--reso="shark"`, finally if you want no banner or 
    large space taken up specify `--reso="none"` which will output and clear the screen without an option.


<h3>Visitors :</h3>
<br>
<img src="https://profile-counter.glitch.me/Red-Rabbit/count.svg" alt="Visitors">
