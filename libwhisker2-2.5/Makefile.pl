#!/usr/bin/perl
#
# Generic perl application Makefile
#
# Copyright (c) 2009, Jeff Forristal (wiretrip.net)
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without 
# modification, are permitted provided that the following conditions 
# are met:
#
# - Redistributions of source code must retain the above copyright 
# notice, this list of conditions and the following disclaimer.
#
# - Redistributions in binary form must reproduce the above copyright 
# notice, this list of conditions and the following disclaimer in the 
# documentation and/or other materials provided with the distribution.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS 
# "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT 
# LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS 
# FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE 
# COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, 
# INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, 
# BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; 
# LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER 
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT 
# LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN 
# ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE 
# POSSIBILITY OF SUCH DAMAGE.

$VERSION	=	'2.5';		# version of the app
$PACKAGE	=	'LW2';		# name of the app
$TARGET		=	'LW2.pm';	# target build filename

$SRCDIR		= 	'src';		# dir containing .pl parts
$MAIN		=	'globals.pl_';	# main app logic/global library logic
$HEADER		=	'header.pod';	# POD/file header
$FOOTER		= 	'footer.pod';	# POD/file footer

$LIBRARY	=	 1;		# is it a library?
$HASPOD		=	 1;		# does it have embedded POD?

$DESTDIR	= '';		# installation directory prefix

#### supported build options #########################################

# general commands supported by this makefile

$COMMANDS{clean}	= \&command_clean;
$COMMANDS{lib}		= \&command_build		if($LIBRARY);
$COMMANDS{build}	= \&command_build		if(!$LIBRARY);
$COMMANDS{install}	= \&command_install;
$COMMANDS{uninstall}	= \&command_uninstall;
$COMMANDS{support}	= \&command_support;
$COMMANDS{sockdiag}	= \&command_socket_diag;
$COMMANDS{nopod} 	= \&command_strip_pod		if($HASPOD);


# commands specific to this app

$COMMANDS{install_lw1} = \&command_install_compat;


#### external modules ################################################

# modules to check for and track if they are installed
#
# Module values:
#	0 = just try to load module, but don't error if not available
#	1 = abort build if module isn't available

%MODULES = (
	'Socket'	=> 0,
	'MIME::Base64'	=> 0,
	'MD5'		=> 0,
	'Net::SSLeay'	=> 0,
	'Net::SSL'	=> 0,
	'POSIX'		=> 0
);

#### end config ######################################################

$|++;

# internal vars
%BUILD		= ();
$CWD		= ();
$COMMAND 	= '';
%DESCRIPTIONS	= ();

# first check arguments
if($ARGV[0] eq ''){
	print STDOUT "$PACKAGE version $VERSION build options:\n\n";

	# load the command descriptions
	while(<DATA>){
		tr/\r\n//d;
		my ($name,$desc)=split(/\t/,$_,2);
		$DESCRIPTIONS{$name}=$desc;
	}

	foreach (keys %COMMANDS){
		print STDOUT "- Makefile.pl $_";
		if(defined $DESCRIPTIONS{$_}){
			print STDOUT "\t",$DESCRIPTIONS{$_};
		}
		print STDOUT "\n";
	}
	print STDOUT "\n";
	exit;
}

# the makefile requires Config, Cwd, and Pod::Man modules
$MODULES{Config}	= 0;
$MODULES{Cwd}		= 0;
$MODULES{'Pod::Man'}	= 0;

# next check for external modules
foreach (keys %MODULES){
	eval "use $_;";
	if(!$@){
		$MODULES{$_}++;
	} else {
		if($MODULES{$_}>0){
			print STDERR "Error: module '$_' required.\n";
			exit;
		}
	}
}

# adjust DESTDIR, if needed
$DESTDIR = $ENV{DESTDIR} if(defined $ENV{DESTDIR});

# parse command line build options
while($COMMAND = shift @ARGV){

	if(defined $COMMANDS{$COMMAND}){
		$COMMANDS{$COMMAND}->();
	} else {
		print STDERR "Error: bad build command '$COMMAND'\n";
		exit;
	}
}

exit;

#########################################################################

sub command_clean {
	unlink $TARGET if(-e $TARGET);
	print STDOUT "Clean.\n";
}

sub command_install {
	command_install_library()	if($LIBRARY);
	command_install_pod()		if($HASPOD);
}

sub command_uninstall {
	command_uninstall_library()	if($LIBRARY);
	command_uninstall_pod()		if($HASPOD);
}

sub command_install_pod {
	return if(!$HASPOD);
	if($MODULES{'Pod::Man'}==0){
		print STDERR "WARNING: Pod::Man not available; man page not installed\n";
		return;
	}
	command_build() if(!-e $TARGET);
	die("Can not install without Config.pm") if($MODULES{Config}==0);
	$CWD=&cwd if($MODULES{Cwd}>0);
	my $where=$DESTDIR . $Config{'man3direxp'};
	my $t = $TARGET;
	if($LIBRARY){
		$t="$PACKAGE.3pm";
	} else {
		$t=~s/\.pl$//i;
		$t.='.3';
	}
	if(!-e $where){
	  print STDOUT "WARNING!\n\n",
		"The local man3 site directory does not exist:\n",
		"$where\n\nPlease create this directory and try again.\n\n";
		exit;
	}

	my $parser = Pod::Man->new (
			release => $VERSION, 
			section => 3,
			name => $PACKAGE
		);

	open(IN,'<'.$TARGET)||puke($TARGET);
	$temp = <IN>;
	if($temp=~m/^# NOPOD NOTICE:/){
		print STDERR "Pod has been stripped; not installing man page\n";
		return;
	}
	
	chdir($where);
	open(OUT,'>'.$t)||die("Can't open $where/$t for write");
	chmod 0644, $t;

	$parser->parse_from_filehandle(\*IN,\*OUT);

	close(IN);
	close(OUT);
	if(-s "$t"){
		print STDOUT "$t installed to $where\n";
	} else {
		print STDOUT "Error installing $t to $where\n";
	}
	exit if($MODULES{Cwd}==0);
	chdir($CWD);
}

sub command_uninstall_pod {
	die("Can not uninstall without Config.pm") if($MODULES{Config}==0);
	$CWD=&cwd if($MODULES{Cwd}>0);
	my $where=$DESTDIR . $Config{'man3direxp'};
	my $t = $TARGET;
	if($LIBRARY){
		$t="$PACKAGE.3pm";
	} else {
		$t=~s/\.pl$//i;
		$t.='.3';
	}
	chdir($where);
	if(-e $t){
		unlink $t;
		print STDOUT "$t uninstalled.\n";
	} else {
		print STDOUT "$t not installed.\n";
	}
	exit if($MODULES{Cwd}==0);
	chdir($CWD);
}

sub command_install_library {
	return if(!$LIBRARY);
	command_build() if(!-e $TARGET);
	die("Can not install without Config.pm") if($MODULES{Config}==0);
	$CWD=&cwd if($MODULES{Cwd}>0);
	my $where=$DESTDIR . $Config{'installsitelib'};
	if(!-e $where){
	  print STDOUT "WARNING!\n\n",
		"The local perl site directory does not exist:\n",
		"$where\n\nPlease create this directory and try again.\n\n";
		exit;	
	}
	open(IN,'<'.$TARGET)||puke($TARGET);
	chdir($where);
	open(OUT,'>'.$TARGET)||die("Can't open $where/$TARGET for write");
	chmod 0755, $TARGET;
	while(<IN>){ 
		print OUT; 
	}
	close(IN); 
	close(OUT);
	if(-s "$TARGET"){
		print STDOUT "$TARGET installed to $where\n";
	} else { 
		print STDOUT "Error installing $TARGET to $where\n"; 
	}
	exit if($MODULES{Cwd}==0);
	chdir($CWD);
}

sub command_uninstall_library {
	die("Can not uninstall without Config.pm") if($MODULES{Config}==0);
	$CWD=&cwd if($MODULES{Cwd}>0);
	my $where=$DESTDIR . $Config{'installsitelib'};
	chdir($where);
	if(-e $TARGET){
		unlink $TARGET;
		print STDOUT "$PACKAGE uninstalled.\n";
	} else {
		print STDOUT "$PACKAGE not installed.\n";
	}
	exit if($MODULES{Cwd}==0);
	chdir($CWD);
}

sub command_build {
	$CWD=&cwd if($MODULES{Cwd}>0);

	# open target file for output
	open(OUT,'>'.$TARGET)||die("Can't open $TARGET for write");
	chmod 0755, $TARGET;

	# print out the shebang line
	print OUT "#!",$^X,"\n";

	# embed the package and version info
	print OUT "# $PACKAGE version $VERSION\n";

	# switch to the src directory
	opendir(DIR,$SRCDIR);
	chdir($SRCDIR);

	# print out the initial header and infoz
	readlib($HEADER,1);

	if($LIBRARY){
		print OUT "package $PACKAGE;\n";
		print OUT "\$",$PACKAGE,"::VERSION=\"$VERSION\";\n";
	} else {
		print OUT "\$VERSION=\"$VERSION\";\n";
	}
	print OUT "\$PACKAGE='",$PACKAGE,"';\n";

	# handle main logic
	print OUT "\n";
	if($LIBRARY){
		print OUT "BEGIN {\n";
		print OUT "package $PACKAGE;\n";
		print OUT "\$PACKAGE='",$PACKAGE,"';\n";
	}
	readlib($MAIN,0);
	if($LIBRARY){
		print OUT "\n} # BEGIN\n\n";
	}

	# handle all the source files
	&readlibs;	

	# and now the footer
	readlib($FOOTER,1);
	print OUT "1;\n" if($LIBRARY);

	# we're all done; print status and cleanup
	print STDOUT "$PACKAGE built.\n";
	close(OUT);
	closedir(DIR);
	exit if($MODULES{Cwd}==0);
	chdir($CWD);
}

sub command_strip_pod {
	return if(!$HASPOD);
	command_build() if(!-e $TARGET);
	open(OUT,'>'."$TARGET.nopod") || die("Couldn't open $TARGET.nopod");
	open(IN,'<'.$TARGET) || puke($TARGET);
	&strip_pod;
	close(OUT); 
	close(IN);
	unlink $TARGET;
	rename "$TARGET.nopod", $TARGET;
	chmod 0755, $TARGET;
	print STDOUT "POD removed.\n";
}

sub command_support {
	print STDOUT "Perl $] on '$^O'\n";
	print STDOUT "Architecture: '$Config{archname}'\n"
		if($MODULES{Config}>0);
	print STDOUT "\n";

	foreach $lib (keys %MODULES){
		print STDOUT $lib, ' ->', ' 'x(20-length($lib));
		if($MODULES{$lib}>0){
			print STDOUT 'yes';
			my $name = $lib.'::VERSION';
			if(defined $$name){
				print STDOUT ' (version '.$$name.')';
			}
			print "\n";
		
		} else {
			print STDOUT "no\n";
		}
	}
}

sub command_socket_diag {
        use Socket;
        use POSIX;
        my @what=qw(
                INADDR_ANY INADDR_BROADCAST INADDR_LOOPBACK INADDR_NONE
                AF_INET PF_INET MSG_OOB
                SOCK_DGRAM SOCK_RAW SOCK_SEQPACKET SOCK_STREAM
                SOL_SOCKET SOMAXCONN F_GETFL F_SETFL O_NONBLOCK
                EINPROGRESS EWOULDBLOCK
                SO_BROADCAST SO_KEEPALIVE SO_LINGER SO_OOBINLINE
                SO_RCVBUF SO_RCVLOWAT SO_RCVTIMEO SO_REUSEADDR SO_SNDBUF
                SO_SNDLOWAT SO_SNDTIMEO SO_TYPE SO_USELOOPBACK
        );
        print STDOUT "\nPerl: $^O\n";
        print STDOUT "Perl version: $]\n";
        print STDOUT "Uname processor: ",`uname -m`;
        print STDOUT "Uname kernel: ",`uname -r`;
	print STDOUT "Socket defines:\n\n";
        map { verify($_) } @what;
}


#########################################################################

sub command_install_compat {
	die("Can not install without Config.pm") if($MODULES{Config}==0);
	$CWD=&cwd if($MODULES{Cwd}>0);
	my $where=$DESTDIR . $Config{'installsitelib'};
	if(!-e $where){
	  print STDOUT "WARNING!\n\n",
		"The local perl site directory does not exist:\n",
		"$where\n\nPlease create this directory and try again.\n\n";
		exit;	
	}
	open(IN,'<'.'compat/LW.pm')||puke('compat/LW.pm');
	chdir($where);
	open(OUT,'>'.'LW.pm')||die("Can't open $where/LW.pm for write");
	chmod 0755, 'LW.pm';
	while(<IN>){ 
		print OUT; 
	}
	close(IN); 
	close(OUT);
	if(-s "LW.pm"){
		print STDOUT "LW.pm bridge installed to $where\n";
	} else { 
		print STDOUT "Error installing LW.pm to $where\n"; 
	}
	exit if($MODULES{Cwd}==0);
	chdir($CWD);
}

#########################################################################

sub puke {
	my $file = shift;
	print STDERR "Build error: missing/corrupted file $file\n";
	eval "close(OUT)";
	exit;
}

sub readlib {
	my $file=shift;
	my $replace_flag=shift||0;
	return if(defined $BUILD{$file});
	puke($file) if(!-e $file);
	$BUILD{$file}++;
	open(IN,'<'.$file)||puke($file);
	while(<IN>){		
		next if(m/^#GPL/ || m/^#LIC/ || m/^#BSD/);
		s/\r\n$/\n/;
		if($replace_flag){
			s/\$VERSION/$VERSION/g;
			s/\$TARGET/$TARGET/g;
			s/\$PACKAGE/$PACKAGE/g;
		}
		print OUT $_; 
	}
	close(IN);	
}

sub readlibs {
	my $file;
	my @FF=();
	while($file=readdir(DIR)){
		next if($file=~/^\./);
		next if($file eq $MAIN);
		next if($file eq $HEADER);
		next if($file eq $FOOTER);
		next if($file =~ /^_/);
		push(@FF,$file);
	}
	my @FE = sort @FF;
	foreach $file (@FE){
		readlib($file,0); }
}

	
sub strip_pod {
	my $inpod=0;
	my $last='';

	# put a small notice in the file to keep people from wondering where
	# all the whitespace went...
	print OUT "# NOPOD NOTICE: the documentation and whitespace have been stripped\n";
	print OUT "# from this file in order to reduce filesize.\n#\n\n";

	my $IN_INITIAL_COMMENTS=1;
	while(<IN>){
		s/^[ \t]+//; # remove leading whitespace
		my $line=$_;
		next if(m/^#/ && !$IN_INITIAL_COMMENTS);
		tr/\r\n//d; # remote CRLF
		if($IN_INITIAL_COMMENTS && !m/^#/){ 
			$IN_INITIAL_COMMENTS=0; 
			next; 
		}
		next if($_ eq '');
		$inpod=1 if($line=~/^=(head1|item|pod|back)/);
		if(!$inpod){
			$line=~tr/\r//d;
			print OUT $line if(!($line eq "\n" && $last eq "\n"));
		}
		$inpod=0 if($line=~/^=cut/);
		$last = $line;
	}
}

sub verify {
        my $temp=-1;
        my $name=shift;
        eval { $temp=sprintf("%lu",&$name) };
        $temp="\t$temp" if(length($name)<7);
        print STDOUT "\t$name:\t$temp\n";
}


__DATA__
nopod	Strip the POD documentation and whitespace
clean	Clean up the build tree
lib	Build the library
build	Build the application
install	Install the components to the Perl site directory
uninstall	Uninstall/remove the components from your system
support	List various external module support information
sockdiag	Diagnostics for troubleshooting Socket.pm problems
install_lw1 Install the LW.pm compatiblity bridge
