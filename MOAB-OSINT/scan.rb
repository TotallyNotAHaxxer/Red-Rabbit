require 'net/http'
require 'colorize'
require 'time'

www = ARGV[0] || ''.empty?

if ARGV[0].nil? 
    puts "\n\n\t[ERROR]-[FATAL] NO HOST SPECIFIED OR PARSED\n\n\t".colorize(:red)
    puts "\n\n\t[ERROR]-[FATAL] IN [ [ #{__FILE__} ] ]  TRY SPECIFING HOST\n\n\t".colorize(:red)
    puts "\n\tABORTING........."
    exit!
end



def webscan(www)
    ipa = IPSocket::getaddress("#{www}")
    puts "[+] Scanning Host for ports....."
    ports = 1..65389
    ports.each do |scan|
        begin
            Timeout::timeout(0.1){TCPSocket.new(ipa, scan)}
            rescue
                #puts "[PORT] #{scan} IS [CLOSED]"
            else
                dt = DateTime.now
                puts dt.next_month.strftime("\033[36m[\033[35m%H:%M\033[36m] ") + "\033[35m[\033[36mRED-RABBIT-INF\033[35m] " + "[PORT#{scan}] CAME BACK OPEN"
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

webscan(www)