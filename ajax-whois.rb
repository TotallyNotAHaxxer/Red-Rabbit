require 'socket'
require 'net/http'
require 'colorize'
require 'tty-spinner'
require 'optparse'
require 'httparty'
require 'timeout'
require 'uri'

url     = ARGV[0] || "".empty?

URI.parse("#{url}").port # => 80
uri = URI.parse("#{url}")
date = Time.new 
puts  '         Date At Start ===> '.colorize(:yellow) + date.inspect 
puts  '         Url Target    ===> '.colorize(:yellow) + url    
puts '-------------------------------------------------------'
puts "[*] Target is => ".colorize(:blue) + url
puts '-------------------------------------------------------'  
puts '[*] Gathering Info on URL => '.colorize(:blue) + url
resur = Net::HTTP.get_response(URI.parse(url.to_s))
if resur.code == '200'
    puts '[+] Connection OK' 
elsif resur.code == '301'
    puts '[+] good'
elsif resur.code == '302'
    puts '[+] Domain not found'
elsif resur.code == '202'
    puts '[+] Domain Accepted IPA'
elsif resur.code == '201'
    puts '[+] Domain Created'
elsif resur.code == '204'
    puts '[-] hmmm not much content here'
elsif resur.code == '206'
    puts '[-] little content, but why?.....'
elsif resur.code == '303'
    puts '[-] See another page'
elsif resur.code == '304'
    puts '[-] Domain Isnt modified yet'
elsif resur.code == '305'
    puts '[-] try using proxies'
elsif resur.code == '308'
    puts '[-] perma redirect'
elsif resur.code == '400'
    puts '[-] Bad Request'
elsif resur.code == '403'
    puts '[-] your ip is not wanted here'
elsif resur.code == '405'
    puts '[-] unwanted domain'
elsif resur.code == '404'
    puts '[-] Domain Not Foud'
elsif resur.code == '423'
    puts '[-] locked domain'
elsif resur.code == '425'
    puts '[-] too eraly'
elsif resur.code == '429'
    puts '[-] way to much requests'
elsif resur.code == '413'
    puts '[-] Payload to large'
elsif resur.code == '407'
    puts '[-] hmmmm proxy auth is needed'
elsif resur.code == '410'
    puts '[-] Domain Gone '
elsif resur.code == '500'
    puts '[-] Server Side Error'
elsif resur.code == '503'
    puts '[-] Server Unavalible/Offline'
else
    puts '[-] Server May be offline '
    puts '[+] Trying a new request '
    sleep 1 
    resur = Net:HTTP.get_response(URI.parse(url.to))
    if resur.code == '200'
        puts '[+] Connection OK'.colorize(:blue)
    elsif resur.code == '301'
        puts '[+] good'
    elsif resur.code == '302'
        puts '[+] Domain not found'
    elsif resur.code == '202'
        puts '[+] Domain Accepted IPA'
    elsif resur.code == '201'
        puts '[+] Domain Created'
    elsif resur.code == '204'
        puts '[-] hmmm not much content here'
    elsif resur.code == '206'
        puts '[-] little content, but why?.....'
    elsif resur.code == '303'
        puts '[-] See another page'
    elsif resur.code == '304'
        puts '[-] Domain Isnt modified yet'
    elsif resur.code == '305'
        puts '[-] try using proxies'
    elsif resur.code == '308'
        puts '[-] perma redirect'
    elsif resur.code == '400'
        puts '[-] Bad Request'
    elsif resur.code == '403'
        puts '[-] your ip is not wanted here'
    elsif resur.code == '405'
        puts '[-] unwanted domain'
    elsif resur.code == '404'
        puts '[-] Domain Not Foud'
    elsif resur.code == '423'
        puts '[-] locked domain'
    elsif resur.code == '425'
        puts '[-] too eraly'
    elsif resur.code == '429'
        puts '[-] way to much requests'
    elsif resur.code == '413'
        puts '[-] Payload to large'
    elsif resur.code == '407'
        puts '[-] hmmmm proxy auth is needed'
    elsif resur.code == '410'
        puts '[-] Domain Gone '
    elsif resur.code == '500'
        puts '[-] Server Side Error'
    elsif resur.code == '503'
        puts '[-] Server Unavalible/Offline'
    else
        puts '[-] Second Test Failed '
    end
end

puts '---------------- BASIC INFORMATION FOR URL -------------- '
uri = URI.parse("#{url}")
http = Net::HTTP.new(uri.host, uri.port)
request = Net::HTTP::Get.new(uri.request_uri)
request["User-Agent"] = "My Ruby Script"
request["Accept"] = "*/*"
response = http.request(request)
response["content-type"]
response.each_header do |key, value|
  p "#{key} => #{value}"
end
p response.instance_variable_get("@header")
userag     = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.54 Safari/537.36"
proxy_user = nil
proxy_pass = nil 
uri  = URI.parse("#{url}")
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl        = true if uri.scheme == 'https'
http.verify_mode    = OpenSSL::SSL::VERIFY_NONE
req  = Net::HTTP::Get.new(uri.request_uri)
pp req.to_hash
req["User-Agent"] = "#{userag}"
res   = http.request(req)
res.code
res.message 
res.code_type
res.content_type
pp res.to_hash
puts '-------------------------'
puts '[*] Response ~> '.colorize(:blue) + resur.code.colorize(:red)
sleep 0.1
puts '[*] Checking More Connections..'.colorize(:yellow)
puts '--------------------------'
puts '[*] Gathering Header Info....'.colorize(:yellow)
puts "[!] Warning, upon further testing of #{__FILE__}
[!] Sometimes the server info will go empty
[!] Right now i am planning on fixing this bug 
[!] and massive issue, however this project
[!] was programmed over the course of a month
".colorize(:yellow)
a = HTTParty.get(url).headers


puts '[*] Query          => '.colorize(:blue) , uri.query  
puts '[*] Scheme         => '.colorize(:blue) , uri.scheme
puts '[*] Port  Main     => '.colorize(:blue) , uri.port
puts '[*] HOSTNAME       => '.colorize(:blue) , uri.host
puts '[*] Path           => '.colorize(:blue) , uri.path
puts '[*] Request URI    => '.colorize(:blue) , uri.request_uri 
puts '[*] Server         => '.colorize(:blue) ,  a["server"].colorize(:red)
puts '[*] Date           => '.colorize(:blue) ,  a["date"].colorize(:red)
puts '[*] Content        => '.colorize(:blue) ,  a["content-type"].colorize(:red)
puts '[*] Response Code  => '.colorize(:blue) , resur.code
puts '[*] Last-mod       => '.colorize(:blue) 
puts a["last-modified"]
puts '[*] trans-enc      => '.colorize(:blue) 
puts a["transfer-encoding"]
puts '[*] Connection     => '.colorize(:blue) + a["connection"].colorize(:red)
puts '[*] Access-control => '.colorize(:blue) 
puts a["access-control-allow-origin"]
puts '[*] Cache-control  => '.colorize(:blue) 
puts resur.response["Cache-Control"]
puts '-----------------------SERVER INF--------------------'  
puts '[*] X-Frame-OP      | '.colorize(:blue) 
puts resur.response["x-frame-options"]
puts '[*] X-XSS-Protect   | '.colorize(:blue) 
puts  resur.response["x-xss-protection"]
puts '[*] X-Content-type  | '.colorize(:blue) 
puts resur.response["x-content-type-options"]
puts '[*] Max-Age         |'.colorize(:blue) 
puts resur.response["max-age"]
puts '[*] X-Request-ID    |'.colorize(:blue)
puts resur.response["x-request-id"]
puts '[*] X-Runtime       |'.colorize(:blue)
puts resur.response["x-runtime"]
puts '[*] Access Control  |'.colorize(:blue)
puts resur.response["access_control_max_age"]
puts '[*] Access Allow    |'.colorize(:blue)
puts resur.response["access_control_allow_methods"]
puts '[*] Content Length  |'.colorize(:blue)
puts resur.response["content-length"]
puts '[*] Connection      |'.colorize(:blue)
puts resur.response["connection"]
puts '[*] Content_Dispo   |'.colorize(:blue)
puts resur.response["content_disposition"]
puts '[*] Expires         |'.colorize(:blue)
puts resur.response["expires"]
puts '[*] set-cookie      |'.colorize(:blue)
puts resur.response["set-cookie"]
puts '[*] user-Agent      |'.colorize(:blue)
puts resur.response["user-agent"]
puts '[*] bfcache-opt-in  |'.colorize(:blue)
puts resur.response["bfcache-opt-in"]
puts '[*] Content encode  | '.colorize(:blue)
puts resur.response["content-encoding"]
puts '[*] content-sec     | '.colorize(:blue)
puts resur.response["content-security-policy"]
puts '[*] Session Cookie  |'.colorize(:blue)
puts resur.response["set-cookie"]
puts '[*] strict-trans    |'.colorize(:blue)
puts resur.response["strict-transport-security"]
puts '[*] method          |'.colorize(:blue)
puts resur.response["method"]
puts '----------------------------------------------------------'
puts '----------------------------------------------------------'