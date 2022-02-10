# this file is the only file that is saving me a lifetime 
# for some reason in running the fake ap both module and standalone 
# i keep getting the error
# main.rb:111:in `send': wrong number of arguments (given 2, expected 0) (ArgumentError)
#
#
# i did the following to try and solve it 
# check the module ( worked fine before )
# used it outside of def's
# used it in the main script alone without a main functions 
# used it without calling begin
# used it with calling begin
# used it in so mnay more scenarios despite having the correct information 
# BTW works 100% fine in literally any other file but the main file, im starting 
# to think there might be a bug in main.rb that is causing it to break like that
# in later updated of RR5 this will be fixed or attempted to fix 
#
# however if you find the error for me contact me on github under the issue tab 
# so i can understand wtf just happened lol 
# https://github.com/ArkAngeL43/Red-Rabbit-V5 
#
#
# script starts here 
require './mod-fake.rb'
require 'optiflag'

module Inputs extend OptiFlagSet
    flag "interface" 
    flag "ssidname"
    flag "pkt_tf"
    and_process!
end


interface1 = ARGV.flags.interface
ssidname1  = ARGV.flags.ssidname
packview   = ARGV.flags.pkt_tf

if ARGV.flags.interface.empty?
    puts "[ - ] I need an interface"
    exit!
end
if ARGV.flags.ssidname.empty?
    puts "[ - ] I need an interface"
    exit!
end



def main_fakeap(fake, apname)
    print "Yes | no "
    print "\nWould you like a over view of the packet > "
    yn = gets.chomp
    if yn == "yes" or yn == "y" or yn == "Y"
        FakeAP.clear
        FakeAP.banner
        FakeAP.reciever("#{fake}", "#{apname}", true)
    end
    if yn == "no" or yn == "n" or yn == "N"
        FakeAP.clear
        FakeAP.banner
        FakeAP.reciever("#{fake}", "#{apname}", false)
    end
end

if packview == "true"         
    FakeAP.clear
    FakeAP.banner
    FakeAP.reciever("#{interface1}", "#{ssidname1}", true)
end
if packview == "false"
    FakeAP.clear
    FakeAP.banner
    FakeAP.reciever("#{interface1}", "#{ssidname1}", false)
end