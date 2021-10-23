#!/usr/bin/perl

# this script shows how to use the simple (aka 'easy') LW routines to
# fetch and save web pages.  The simple routines were made to make it
# easier to do simple web tasks, without dealing with the internals/
# particulars of libwhisker.

use LW2;

####################################################################

# first let's get the Slashdot homepage and it's HTML data

($response_code, $html_data) = LW2::get_page( "http://slashdot.org/" );

if(!defined $response_code){
	print "There was an error\n";
	exit;
}

if($response_code == 200){ # the page exists

	print "Current Slashdot departments:\n";

	while( $html_data=~m/from the (.+) dept./g ){
		print "\t$1\n";
	}

} else {

	print "Slashdot response was $response_code.\n";
}


####################################################################

# now, let's download the Cipherwar homepage and save it to a file
# named "cipherwar.html"

$file="cipherwar.html";

$response_code = LW2::get_page_to_file( "http://www.cipherwar.com/",
					$file );

if(!defined $response_code){
	print "There was an error retrieving cipherwar.com page\n";
} else {
	print "Cipherwar homepage saved to $file.\n";
}

print "\n";
