require 'packetgen'


eth = PacketGen.gen('Eth', src: '00:11:22:33:44:55', dst: 'ff:ff:ff:ff:ff:ff')
ip  = eth.add('IP', src: '99.11.99.11', dst: '192.168.100.17')
tcp = ip.add('TCP', sport:90, dport: 80, flag_syn: 1, body: "BlackHatRuby 1")
pkt = tcp.to_w('wlan2')

pp pkt 
pkt.to_w('wls1', number: 1, interval: 1)