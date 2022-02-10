require 'socket'

class Pcap

  def initialize(pcap_file)
    @pcap_file = open(pcap_file, 'wb')
    # Pcap Global https://wiki.wireshark.org/Development/LibpcapFileFormat#Global_Header
    global_header = [
        0xa1b2c3d4,   # magic_number: used to identify pcap files
        2,            # version_major
        4,            # version_minor
        0,            # thiszone
        0,            # sigfigs
        65535,        # snaplen
        1             # network (link-layer), 1 for Ethernet
    ].pack('ISSIIII')
    @pcap_file.write global_header
  end

  def write(data)
    time_stamp  = Time.now.to_f.round(2).to_s.split('.').map(&:to_i)
    data_length = data.length
    # Pcap Record (Packet) Header: https://wiki.wireshark.org/Development/LibpcapFileFormat#Record_.28Packet.29_Header
    packet_header = [
        time_stamp[0],   # ts_sec timestamp seconds
        time_stamp[1],   # ts_usec timestamp microseconds
        data_length,     # incl_len the number of bytes of packet data actually captured
        data_length      # orig_len the length of the packet as it appeared on the network when it was captured
    ].pack('IIII')
    record = "#{packet_header}#{data}"
    @pcap_file.write(record)
  rescue
    @pcap_file.close
  end
end 

pcap   = Pcap.new(ARGV[0])
socket = Socket.new(Socket::PF_PACKET, Socket::SOCK_RAW, 0x03_00)
loop do
  raw_data = socket.recvfrom(65535)[0]
  pcap.write raw_data
end
