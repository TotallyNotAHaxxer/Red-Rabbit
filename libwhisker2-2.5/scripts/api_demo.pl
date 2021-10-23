#!/usr/bin/perl

# this script is an example on how to make a HTTP request using libwhisker

$|++;
use LW2;

$TARGET=shift;

# check to make sure target is appropriate
if($TARGET!~m#^https*://#){
	print "Usage: $0 http://target-host/url\n";
	exit;
}

# $request contains the request values passed to libwhisker/server
# $response contains the response values received from libwhisker/server
my $request = LW2::http_new_request();
my $response = LW2::http_new_response();

# %jar will contain all our cookies
my $jar = LW2::cookie_new_jar();

# this is no longer needed, since http_new_request() calls it automatically
#LW2::http_init_request($request);


# set the target host and URI; uri_split() will handle all the SSL
# junk if the URL starts with https://, and also handle abnormal ports,
# setting the URI and parameters, etc.
LW2::uri_split($TARGET, $request);


# we can manually set/override values
#$request->{'whisker'}->{'host'}='www.example.com'; # 'localhost' is the default
#$request->{'whisker'}->{'uri'}="/index.html"; # '/' is the default
#$request->{'whisker'}->{'port'}=80;  # port 80 is default


# example on how to enable SSL support
#
# Note: uri_split() will set these values automatically if it is passed
#       a URL which starts with "https://"
#
#$request->{'whisker'}->{'port'}=443; # default SSL port
#$request->{'whisker'}->{'ssl'}=1; # tell libwhisker to use SSL


# Save the SSL server certificate information for viewing
$request->{'whisker'}->{'ssl_save_info'}=1;


# proxy support
#$request->{'whisker'}->{'proxy_host'}='localhost';
#$request->{'whisker'}->{'proxy_port'}=8080;


# anti-IDS
#$request->{'whisker'}->{'anti_ids'}='12345'; # values are the modes


# basic auth
#LW2::auth_set('basic',$request,'username','password');
#LW2::auth_set('basic-proxy',$request,'username','password');
#LW2::auth_set('ntlm',$request,'username','password');
#LW2::auth_set('ntlm',$request,'username','password','domain');


# special function to tweak the request to make sure it's valid HTTP 
# (not required if you set it all manually, but it never hurts)
LW2::http_fixup_request($request);


# if there's any cookies in the %jar, then set them as appropriate
#LW2::cookie_write($jar,$request);

# actually make the request (and get response)
if(LW2::http_do_request($request,$response)){

	print 'ERROR: ', $response->{'whisker'}->{'error'}, "\n";
	print $response->{'whisker'}->{'data'}, "\n";

} else {

	# save any cookies sent to us into our $jar
	#LW2::cookie_read($jar,$response);

	print $response->{'whisker'}->{'code'}, "\n"; # HTTP return code
	print $response->{'Server'}, "\n";		# Server banner

	# Uncomment following line if you want to see resulting HTML data
	#print $response->{'whisker'}->{'data'}, "\n";


	# If we wanted to save the SSL info
	if($request->{whisker}->{ssl}>0 && 
			$request->{whisker}->{ssl_save_info}>0){
		print 'SSL cipher: ',
			$response->{whisker}->{ssl_cipher},"\n";
		print "Server cert:\n\t",
			$response->{whisker}->{ssl_cert_subject},"\n";
		print "Cert issuer:\n\t",
			$response->{whisker}->{ssl_cert_issuer},"\n";
	}

	# Uncomment following line to dump out the entire response hash
	#print LW2::dump('response', $response);

}


# good practice to clean up our mess
LW2::http_reset();
