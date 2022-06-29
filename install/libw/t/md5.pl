use Digest::MD5 qw(md5_hex);

$DATA = "0a\n<html><bod\n28\ny>This is a test response</body></html>\n0\n"
;

print md5_hex($DATA), "\n";

