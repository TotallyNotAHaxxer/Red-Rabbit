#!/usr/bin/perl
#
# crawl_demo v1.0
#
# This file is an example of how to use LW2's new crawl function,
# which is a bit different than the LW (version 1) function.
#

use LW2;

########################################################################
#
# The following code shows how to crawl the site http://www.example.com/
# to a depth of 2 (homepage + all URLs on homepage).
#
########################################################################

$TARGET='www.example.com';

#
# First we'll use the new LW2 new request function.  $REQUEST is
# really a hash reference (aka \%REQUEST).
#
my $REQUEST = LW2::http_new_request( 
		host=>$TARGET,
		method=>'GET',
		timeout=>10
		);


#
# Next, let's change the 'User-Agent' header for our crawler
#
$REQUEST->{'User-Agent'} = 'libwhisker-crawler-demo/1.0';
#
# Note: we can also use $$REQUEST{'User-Agent'} too


#
# Now we'll set the port number, just as an example of setting control
# values using $REQUEST
#
$REQUEST->{whisker}->{port}=80;
#
# Note: again, $$REQUEST{whisker}->{port} would work too


#
# It's always good practice to call http_fixup_request(), to
# ensure everything is protocol-compliant
#
LW2::http_fixup_request($REQUEST);

#
# Great!  Now $REQUEST is all set to be used as the baseline request
# for our crawler.
#

#
# Now we need to make a new crawler
#
my $CRAWLER = LW2::crawl_new(
			"http://$TARGET/",	# start URL
			2,			# depth
			$REQUEST		# premade LW request
		);


#
# $CRAWLER is actually another hash reference (\%CRAWLER), which can
# be directly manipulated to do stuff.  The actual values/structures
# it holds are detailed in the 'crawler.txt' file in the docs/ 
# subdirectory of the libwhisker source tarball.
#


#
# Now, let's tell the crawler that we want it to save all cookies.
# LW v1 used a crawl_set_config() function, which no longer exists in
# LW2.
#
$CRAWLER->{config}->{save_cookies}=1;


#
# Let's also tell the crawler to save all the skipped URLs.  Skipped 
# URLs are URLs which we saw, but did not actually crawl.  We'll
# use a different dereferencing style.  
#
$$CRAWLER{config}->{save_skipped}=1;


#
# Ok, so we configured our crawler.  Now, let's run it!  The {crawl}
# element of the $CRAWLER hash is actually an anonymous subroutine,
# which we call.
#
#
$result=$CRAWLER->{crawl}->();
# or $$CRAWLER{crawl}->()
# or &$CRAWLER->{crawl}()
# or LW2::crawl($CRAWLER)
# they all do the same thing :)

#
# Our crawler returns once it's crawled all available URLs.  
# Let's be good and check for an error...
#
if(!defined $result){
	print "There was an error:\n";
	print $CRAWLER->{errors}->[0];
} else {
	print "We crawled $result URL(s)\n";
}

#
# First, let's print a list of all URLs we found.
#
my ($key, $value);

# $TRACK is a hash ref (\%TRACK)
my $TRACK_HASH = $CRAWLER->{track};

print "\n\nCODE\tURL\n";
while( ($key,$value) = each (%$TRACK_HASH) ){
	print "$value\t$key\n";
}


#
# Next, let's print out any cookies (since we set save_cookies=1)
#
my $COOKIE_HASH = $CRAWLER->{cookies};
print "\n\nCookie name:value\n";
while( ($key,$value) = each (%$COOKIE_HASH) ){
	print "$key: $value\n";
}


#
# That's all there is to it!  Of course, what data is available
# depends on what options you set before you run the crawler.  The
# above is just a general walk-through of basic functionality.
#

#
# If you're curious, you can uncomment the line below and see all the
# data the crawl engine has to offer
#
# print LW2::dumper('crawl',$CRAWLER);

__END__

#
# As a quick recap, below is the minimal amount of code to crawl
# a website, using LW2's default configuration values
#

my $REQUEST = LW2::http_new_request();
my $CRAWLER = LW2::crawl_new( "http://$TARGET/", 2, $REQUEST);
&$CRAWLER->{crawl};
print "List of crawled URLs:\n";
while( my($key,$value)=each(%{$CRAWLER->{track}}){
	print "$value\t$key\n";
}

#
# That's seriously all it takes.
#


#
# Also, if you wish to rerun a crawl(), you can reset the results of
# a previous crawl:
#
$CRAWLER->{reset}->();
$CRAWLER->{crawl}->('http://new.target.com/');

#
# If you don't reset the crawler between crawls, then the old results
# will be used during the next crawl--this can be useful if you want
# to 'resume' a crawl using past results (especially if you dump/restore
# the $CRAWLER object).  In this example, we'll first crawl everything
# in /dir1/ on www.target.com and record it.  Then we'll immediately
# crawl /dir2/ on the same host; the second crawl will reuse the crawl
# results and cache of the first crawl, and thus won't duplicate calls
# to URLs we saw in the first crawl.  The '3' is setting a new depth,
# overriding what we specified in crawl_new().
#
$CRAWLER->{crawl}->('http://www.target.com/dir1/',3);
$CRAWLER->{crawl}->('http://www.target.com/dir2/',3);
