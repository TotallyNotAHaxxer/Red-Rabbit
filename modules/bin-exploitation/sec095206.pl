#!/usr/bin/perl
# cpanel11 - expcheck.pl          Copyright(c) 2006 cPanel Inc.
#                                 All rights Reserved.
# copyright@cpanel.net            http://www.cpanel.net
# This is the offical SEC code for the latest patch of CPANEL 
#
# This file will be run/GET rom the URL
#  http://layer1.cpanel.net/installer/sec092506.pl
# to update to the latest CPANEL Version
#
# in link to `CPanel exploit checker`
#
#
BEGIN {
      unshift @INC,'/scripts';
}

use Socket;
my $httpClient;
eval {
      require cPScript::HttpRequest;
         $httpClient = cPScript::HttpRequest->new();
};



my $gotSigALRM=0;
my $connecttimeout = 30;

my $system;
my $machine;
chomp($system = `uname -s`);
chomp($machine = `uname -m`);

my %UPCONF      = loadcpupdateconf();

my $HTTPB='RELEASE';
if ($UPCONF{'CPANEL'} =~ /stable/i) {
   $HTTPB='STABLE';
} elsif ($UPCONF{'CPANEL'} =~ /current/i) {
   $HTTPB='CURRENT';
} elsif ($UPCONF{'CPANEL'} =~ /edge/i) {
   $HTTPB='EDGE';
} elsif ($UPCONF{'CPANEL'} =~ /nightly/i) {
   $HTTPB='NIGHTLY';
}

my $arch = '';
if ($system =~ /freebsd/i) {
      $arch = '-FREEBSD';
}
if ($machine =~ /x86_64/i) {
      $arch .= '-x86_64';
}
if ($machine =~ /amd64/i) {
      $arch .= '-amd64';
}


print "cPanel Security Patch (w/new wrapper) (sec092506) v1\n";
checkcpsec();
updatewrapper();
print "Patch Complete\n";


sub updatewrapper {
   my $wrap_path='/usr/local/cpanel/bin/cpwrap';
   if (! -e $wrap_path) {
      print "No wrapper present (ok) ... skipped\n";
      return();
   }

   my $updated=0;

   open(WR,'<',$wrap_path);
   while(<WR>) { if (/REMOTE_USER/) { $updated=1;last;  } }
   close(WR); 
   if ($updated == 1) { print "Wrapper was already updated!\n"; return(); }
   print "Updating cpwrap.....";   
   my $cpwrap;
   if (defined($httpClient)) {
      $cpwrap = $httpClient->request(host => "httpupdate.cpanel.net", url => "/cpanelsync/${HTTPB}${arch}/bin/cpwrap");
   } else {
      $cpwrap = httpreqv2("httpupdate.cpanel.net","/cpanelsync/${HTTPB}${arch}/bin/wrap");
   }
   open(WR,'>',$wrap_path);
   print WR $cpwrap;
   close(WR);
   chmod(oct(4755),$wrap_path);
   print "Done\n";
}



sub checkcpsec() {
   patchps('/usr/local/cpanel/bin/mysqladmin');
   patchps('/usr/local/cpanel/bin/hooksadmin');
}


sub patchps {
    my $script = shift;
    if (! -e $script) { return; }

    my %TA=('unshift( @INC, "/usr/local/cpanel" );'=>0,,
    '@INC=grep(!/(^\.|\.\.|\/\.+)/,@INC);'=>0,
    '@INC=grep(/^(\/usr\/lib\d*\/perl|\/usr\/local\/lib\d*\/perl|\/usr\/local\/cpanel)/,@INC);'=>0);
	my @RM = ('\/usr\/lib\/perl|','push(@INC','push( @INC');

	open(SS,'+<',$script);
	print "Patching $script...";
	flock(SS,LOCK_EX);
	my @SF=<SS>;
	foreach my $l (@SF) {
		foreach my $a (keys %TA) {
			if (index($l,$a) > -1) { delete $TA{$a}; next(); } 
		}
	}
	if (!grep(/BEGIN/, @SF)) {
		my @TSF;
		push(@TSF,shift(@SF));
		push(@TSF,"BEGIN {\n","}\n");
		push(@TSF,@SF);
		@SF=@TSF;
	}

	seek(SS,0,0);
	my $skipline=0;
	foreach (@SF) {
		foreach my $rm (@RM) {
			if (index($_,$rm) > -1) { print "-";$skipline=1; last(); }
		}
		if ($skipline)  { $skipline=0; next(); }
		print SS $_;
		if (/BEGIN/) {
			foreach my $a (keys %TA) {
				print "+\n";
				print SS "\t" . $a . "\n";
			}
		}
	}

	flock(SS,LOCK_UN);
	close(SS);
	print "...Done\n";
}


sub loadcpupdateconf {
   my (%UPCONF);
   open(CPUPDATE,"/etc/cpupdate.conf");
   while(<CPUPDATE>) {
      s/\n//g;
      my($name,$value) = split(/=/, $_);
      $UPCONF{$name} = $value;
   }
   close(CPUPDATE);
   return(%UPCONF);
}

sub httpreqv2 {
   my($page);
   my($host,$url,$protocol,$destfile) = @_;
   if ($#HOST_IPs == -1) {
      @HOST_IPs = getAddressList($host);
   }

   if ($protocol eq "" || $protocol eq "0") { $protocol = 'HTTP/1.0'; }
   if ($protocol eq "1") { $protocol = 'HTTP/1.1'; }


   my($status) = 1;
   my($keepalive) = 0;

   if ($destfile ne "") {
      open(DESTFILE,">$destfile");
   }

   print "Fetching http://${host}${url} ($liveconnect)....";

   my @THIS_HOST_IPs = @HOST_IPs;


   eval {
      $SIG{'PIPE'} = $SIG{'ALRM'} = sub {
         print "...Timeout...";
         $gotSigALRM = 1;
      };
      alarm($connecttimeout);
      if ($liveconnect == 0) {
# Iterate through all of the IPs until either a connection
#  is established, or we run out of IPs..
         foreach my $addr (@THIS_HOST_IPs) {
            close(Socket_Handle);
            $gotSigALRM = 0;

            my $proto = getprotobyname('tcp');
            socket(Socket_Handle, AF_INET, SOCK_STREAM, $proto);

            print '@' . "${addr}...";
            $iaddr = inet_aton("$addr")  || do {
               print "...Unable to translate IP address for host: ${host}...";
               reorder($addr);
               alarm($connecttimeout);
               next;
            };
            $port = getservbyname('http', 'tcp');
            $sin = sockaddr_in($port, $iaddr) || do { print "ERROR: $!\n"; };
            connect(Socket_Handle, $sin) || do {
               print "...Unable to connect to host ${host}...";
               reorder($addr);
               alarm($connecttimeout);
               next;
            };
            if ($gotSigALRM) {
               reorder($addr);
               alarm($connecttimeout);
               next();
            }

            $connectedhost = ${addr};

            $liveconnect = 1;
            last;
         }
         if ($liveconnect == 0) {
         close(DESTFILE);
            close(Socket_Handle);
            print "...Failed...";
            return("",0);
         }
      } else {
         print '@' . "${connectedhost}...";
      }
      print "...connected...";

      send Socket_Handle, "GET $url ${protocol}\r\nHost: $host\r\n\r\n",0;

      my $clength = 0;
      my $headers;

      print "...receiving...";
      while(<Socket_Handle>) {
         alarm(20);

         if (/^Connection: (\S+)/i) {
            my $cstatus = $1;
            if ($cstatus =~ /alive/i) {
               $keepalive=1;
            }
         } elsif (/^HTTP\/\d+\.\d+ (\d+)/) {
            if ($1 eq "404" || $1 eq "500" || $1 eq "301") {
               print "Error $1 while fetching url http://$host/$url\n";
               close(DESTFILE);
               close(Socket_Handle);
               $liveconnect = 0;
               return("",0);
            }
         } elsif (/^Content-length: (\d+)/i) {
            $clength = $1;
         }
         if (/^\n$/ || /^\r\n$/ || /^$/) { last; }
      }

      my $togo = $clength;
      my $buffer;
      my $percent;
      my $lastpercent;

      while($togo > 0) {
         alarm(20);
         if ($togo > 4096) {
            read Socket_Handle, $buffer, 4096;
            $togo -= 4096;
         } else {
            read Socket_Handle, $buffer, ${togo};
            $togo = 0;
         }

         $percent = int((($clength - ${togo}) / ${clength}) * 100);
         if ($percent != $lastpercent) {
            print ${percent} . '%' . '...';
         }
         $lastpercent = $percent;

         if ($destfile eq "") {
            $page .= $buffer;
         } else {
            print DESTFILE $buffer;
         }
      }
      alarm(0);
      $status = 1;

   };

   print "...Done\n";

   close(DESTFILE);

   if ($status == 0 || $keepalive == 0)  {
      $liveconnect = 0;
      close(Socket_Handle);
   } else {
      $liveconnect = 1;
   }

   if ($destfile ne "") {
      return($status);
   } else {
      return($page);
   }
}

sub getAddressList {
   my ($host) = @_;

   my @addresses = gethostbyname($host);
   @addresses = map { inet_ntoa($_); } @addresses[4..$#addresses];

   if ($#addresses == -1)  { die "${host} could not be resolved to an ip adddress, please check your /etc/resolv.conf"; }
   return @addresses;
}