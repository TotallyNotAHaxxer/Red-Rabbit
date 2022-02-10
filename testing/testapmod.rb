require './mod-fake.rb'


def main_fakeap()
    print "Interface    > "
    fake = gets.chomp 
    print "Fake AP Name > "
    apname = gets.chomp
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

main_fakeap