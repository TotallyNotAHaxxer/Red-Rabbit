sudo apt install golang -y 
sudo apt install ruby -y 
sudo apt install python3 -y 
sudo apt install perl
go get github.com/ArkAngeL43/port-scanning/port
go get github.com/briandowns/spinner
go get github.com/google/gopacket
go get github.com/google/gopacket/layers
go get github.com/google/gopacket/pcap
go get github.com/schollz/wifiscan
go get github.com/shirou/gopsutil/cpu
go get github.com/shirou/gopsutil/mem
go get github.com/shirou/gopsutil/host
go get github.com/shirou/gopsutil/disk
go get github.com/sirupsen/logrus
go get github.com/manifoldco/promptui
go get github.com/PuerkitoBio/goquery
go get golang.org/x/net/proxy
go get github.com/bndr/gotabulate
go get github.com/joeljunstrom/go-luhn
go get gopkg.in/vmihailenco/msgpack.v2
go get github.com/alexflint/go-arg
go get github.com/k3a/html2text
go get github.com/steelx/extractlinks
go get golang.org/x/net/html
go get github.com/miekg/dns
go get github.com/theckman/yacspin
go get golang.org/x/crypto/ssh
go get github.com/atotto/clipboard
go mod download golang.org/x/crypto
go mod tidy
echo "[ ! ] DATA: attempted to install 25 Go libs"
sudo gem install colorize
sudo gem install packetgen
sudo gem install open-uri
sudo gem install readline
sudo gem install rubygems
sudo gem install timeout
sudo gem install whois whois-parser
sudo gem install net-ssh
sudo gem install httparty
sudo gem install openssl
sudo gem install net-dns
sudo gem install packetfu
sudo gem install resolv
sudo gem install shodanz
sudo gem install optparse
sudo apt install libpcap-dev 
sudo apt install ssh 
sudo gem install async
echo "[ ! ] DATA: Attempted to install 15 ruby modules"
sudo apt-get install perl 
sudo cpan install 
sudo cpan install Imager::QRCode
sudo cpan install HTTP::Tiny
sudo cpan install Term::ANSIColor
sudo cpan install Image::ExifTool
sudo cpan install Text::Table
sudo cpan install Tk::Clock
sudo cpan install Tk
sudo cpan install Tk::DirTree
sudo cpan install Getopt::Std
sudo cpan install LWP
# barely any python things
sudo pip3 install prettytable
sudo pip3 install scapy
sudo pip3 install colorama
# unzip libw
gunzip libwhisker2-2.5.tar.gz
tar xvf libwhisker2-2.5.tar
cd libwhisker2-2.5
sudo perl Makefile.pl install
