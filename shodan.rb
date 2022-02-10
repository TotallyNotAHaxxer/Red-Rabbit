# Dev => ArkAngeL43
# inspo => Inspired by the OSINT chapter in Black Hat Ruby
#
# Module => Shodan 
# type   => Automate shodan searching 
require 'shodanz'
require 'async'

module Shodan_Main
    # start by reading the config-files for a shodan API key
    # For more information on the SHODAN module please turn to the 12 index of the offical RR5 documentation
    # chap -> shodanz
    # basic search
    def Shodan_Main.Basic_search(host, api_key)
        api = Shodanz.client.new(key: "#{api_key}")
        res = api.host_search("#{host}")
        puts "Current results -> #{res['total']}"
        res['matches'].each do |result|
            puts "IP: #{result['ip_str']}"
            puts result['data'] + "\n"
        end
    end
    # get ports of a host
    def Shodan_Main.host_ports(host, api_key) 
        api = Shodanz.client.new(key: "#{api_key}")
        ms_s = api.host("#{host}")
        puts "
        IPA  => #{ms_s['ip_str']}
        ORG  => #{ms_s['org'] || 'n/a'}
        OS   =>  #{ms_s['os'] || 'n/a'}
        "
        ms_s['data'].each do |item|
            puts "
            Port Scanned =>  #{item['port'] || 'n/a'}
            Host banner  =>  #{item['data'] || 'n/a'}
            "
        end
    end
    # api key information 
    def Shodan_Main.get_api(api_key)
        mk = Shodanz.client.new(key: "#{api_key}")
        puts "You're IPV6 => " + mk.my_ip
        puts "\n\033[39mAPI | Scan Credits     :\033[34m", mk.info['scan_credits']
        puts "\n\033[39mAPI | usage limits     :\033[34m", mk.info['usage_limits'] 
        puts "\n\033[39mAPI | Unlocked>?       :\033[34m", mk.info['unlocked'] 
        puts "\n\033[39mAPI | Query Credits    :\033[34m", mk.info['query_credits'] 
        puts "\n\033[39mAPI | IP's Monitored   :\033[34m", mk.info['monitored_ips']
        puts "\n\033[39mAPI | Plan             :\033[34m", mk.info['plan'] 
        puts "\n\033[39mAPI | Is TelNet        :\033[34m", mk.info['telnet']
        puts "\n\033[39mAPI | Https            :\033[34m", mk.info['https'] 
        puts mk.profile 
    end
    # honeypot test
    def Shodan_Main.Honeypot_score(api_key, host)
        main_api = Shodanz.client.new(key: "#{api_key}")
        puts "Honey pot precentage => " + main_api.honeypot_score("#{host}").to_s
    end
    # banner honeypot
    def Shodan_Main.Async(api_key, ip)
        cn = Shodanz.client.new(key: "#{api_key}")
        cn.streaming_api.banners do |banner|
        if ip = banner["#{ip}"]
            Async do
                score = cn.rest_api.honeypot_score("#{ip}").wait
                puts "[!] #{ip} has a #{score * 100}% chance of being a honeypot"
                rescue Shodanz::Errors::RateLimited
                    sleep rand
                    retry
                    rescue # any other errors
                        puts "[!] An error has occured"
                    next
                end
            end
        end
    end
end

api_ = File.open("config-files/shodan-API-KEY.txt")
key_ = api_.read
puts "Key => #{key_}"


#Shodan_Main.Basic_search("apache", "#{key_}")   | Tested
#Shodan_Main.host_ports("1.1.1.1", "#{key_}")    
#Shodan_Main.get_api("#{key_}")                  | Tested
#hodan_Main.Honeypot_score("#{key_}", "1.1.1.1") | Tested
#Shodan_Main.Async("#{key_}", "1.1.1.1")