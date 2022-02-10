# Author:
#   Sabri Hassanyah | @KINGSABRI
# Description:
#   Create Rogue Wireless Access Point
#   Wireshark filter: wlan.fc.type == 0
# Requirements:
#   $ apt install libpcap-dev
#   $ gem install packetgen
# 
#   Enable monitoring mode on your wireless card as *root*
#   $ iw dev
#   $ iw phy phy1 interface add mon0 type monitor && ifconfig mon0 up
# 
require 'packetgen'
require 'colorize' 

def bnnr()
    puts "\x1b[H\x1b[2J\x1b[3J"
    puts <<-'EOF'.colorize(:red)
     ______     ______     _____     ______     ______     ______     ______     __     ______  
    /\  == \   /\  ___\   /\  __-.  /\  == \   /\  __ \   /\  == \   /\  == \   /\ \   /\__  _\ 
    \ \  __<   \ \  __\   \ \ \/\ \ \ \  __<   \ \  __ \  \ \  __<   \ \  __<   \ \ \  \/_/\ \/ 
     \ \_\ \_\  \ \_____\  \ \____-  \ \_\ \_\  \ \_\ \_\  \ \_____\  \ \_____\  \ \_\    \ \_\ 
      \/_/ /_/   \/_____/   \/____/   \/_/ /_/   \/_/\/_/   \/_____/   \/_____/   \/_/     \/_/ 
    
                                                   ((`\
                                                ___ \\ '--._
                                            .'`   `'    o  )
                                            /    \   '. __.'
                                          _|    /_  \ \_\_
                                         {_\______\-'\__\_\
    -----------------------------------------------------------------------------------
                                        _____ ____ _____
                                       /    /      \    \ with the power of ruby
                                     /____ /_________\____\        i become stronger
                                     \    \          /    /                 and stronger 
                                        \  \        /  /
                                           \ \    / /
                                             \ \/ /
                                               \/
EOF
                     
end


def fakeap(interface, ssidname)
    begin
        iface     = "#{interface}"
        broadcast = "ff:ff:ff:ff:ff:ff"
        bssid     = "aa:aa:aa:aa:aa:aa"
        ssid      = "#{ssidname}"
        pkt = PacketGen.gen('RadioTap').add('Dot11::Management', mac1: broadcast, mac2: bssid, mac3: bssid)
                                    .add('Dot11::Beacon', interval: 0x600, cap: 0x401)
        pkt.dot11_beacon.elements << {type: 'SSID', value: ssid}
        pp pkt
        bnnr
        bt = "Fake Beacon"
        puts "\tBeacon Type       SSID Name            Interface        Beacons sent "
        puts "\t---------------------------------------------------------------------"
        while true
            i = 0
            100000.times do
                i = i + 1
                pkt.to_w(iface)
                puts "\t\033[31m#{bt}        \033[36m#{ssid}          \033[31m#{iface}          \033[36m#{i}" 
            end
        end
    rescue Interrupt
        puts " [ + ] Interupt"
        exit!
    rescue PCAPRUB::PCAPRUBError
        puts " [ - ] YOU MUST BE ROOT TO RUN THIS "
        puts " [ - ] Try: sudo ruby #{__FILE__}"
    end
end

def main()
    begin
        print "Interface > "
        ia = gets.chomp
        print "SSID name > "
        ssname = gets.chomp
        #fakeap("#{ia}", "#{ssname}")
        fakeap("wlan0mon", "name1ssidtypeb")
end
end

main()