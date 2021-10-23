#!/usr/bin/perl

#
# Test::Simple is nice, especially since we can print out additional
# notes (namely, the name of the function we're testing). However,
# older perls don't have Test::Simple, so we're stuck with the less
# flexible Test module.  So we check to see which modules are
# available, and use our own mok() (my_ok) to figure out the proper
# way to invoke ok() (i.e. with our without notes).
#

BEGIN {
	$NUMTESTS = 140;
	eval { require "../LW2.pm" };
	die("Libwhisker2 not installed!?!?") if($@);
	@VERSION = split(/\./, $LW2::VERSION);
	if($VERSION[0] != 2 || $VERSION[1] < 4){
		die("This harness expects LW2 2.4 or later");
	}

	$TESTSIMPLE=0;
	eval "use Test::Simple tests=>$NUMTESTS";
	if($@){
		# no Test::Simple, try to use Test instead
		eval "use Test";
		if($@){
			die("Test or Test::Simple need to be installed");
		}
		plan(tests => $NUMTESTS);
		
	} else {
		$TESTSIMPLE=1;
	}
}


sub mok { # my ok(), for figuring out which ok() to use
	my ($result, $note) = @_;
	if($TESTSIMPLE){
		ok($result, $note);
	} else {
		ok($result);
	}
}


##########################################################################
# Encode functions

# disable MIME::Base64 and use built in version instead
LW2::encode_base64('test');
undef $MIME::Base64::VERSION;

$INITIAL_DATA = "Libwhisker2 default test value";
$EXPECTED = "TGlid2hpc2tlcjIgZGVmYXVsdCB0ZXN0IHZhbHVl";
$temp = LW2::encode_base64($INITIAL_DATA, '');
mok( $temp eq $EXPECTED, 'encode_base64' );
$temp = LW2::decode_base64($temp);
mok($temp eq $INITIAL_DATA, 'decode_base64');

$INITIAL_DATA = "/012/abc/ \\&/=;/";
$EXPECTED = "/%30%31%32/%61%62%63/%20%5c%26/%3d%3b/";
$temp = LW2::encode_uri_hex($INITIAL_DATA);
mok($temp eq $EXPECTED, 'encode_uri_hex');

# encode_uri_randomhex is random, so can't be reliably tested
# encode_uri_randomcase is random, so can't be reliably tested

$INITIAL_DATA = "test";
$EXPECTED = "t\x00e\x00s\x00t\x00";
$temp = LW2::encode_unicode($INITIAL_DATA);
mok($temp eq $EXPECTED, 'encode_unicode');

$INITIAL_DATA = "abc\xc0\xafdef";
$EXPECTED = "abc/def";
$temp = LW2::decode_unicode($INITIAL_DATA);
mok($temp eq $EXPECTED, 'decode_unicode');

# TODO: encode_anti_ids?




##########################################################################
# mdx functions

# disable MD5 and use built in version instead
LW2::md5('test');
undef $MD5::VERSION;

$INITIAL_DATA = "Libwhisker2 default test value";
$EXPECTED = "44c9975b203df2c1de2b0bda1b5f515b";
$temp = LW2::md5($INITIAL_DATA);
mok($temp eq $EXPECTED, 'md5');

$INITIAL_DATA = "Libwhisker2 default test value";
$EXPECTED = "0b1e75f19334a77d2b18dce72311d030";
$temp = LW2::md4($INITIAL_DATA);
mok($temp eq $EXPECTED, 'md4');




##########################################################################
# uri functions

$INITIAL_DATA = 'https://user1:pass1@server:81/dir1/dir2/page?param1#frag';
@EXPECTED = ('/dir1/dir2/page', 'https', 'server', 81, 'param1', 'frag',
	'user1', 'pass1');

@temp = LW2::uri_split($INITIAL_DATA);
$BAD = 0;
for($x=0; $x<8; $x++){
	$BAD++ if($EXPECTED[$x] ne $temp[$x]);
}
mok(!$BAD, 'uri_split');
$temp = LW2::uri_join(@temp);
mok($temp eq $INITIAL_DATA, 'uri_join');


%EXPECTED = ( uri => '/dir1/dir2/page', port => 81, host=>'server',
	ssl=>1, uri_user=>'user1', uri_password=>'pass1',
	parameters=>'param1' );

$BAD = 0;
$REQUEST = LW2::http_new_request();
if(ref($REQUEST) ne 'HASH' || !defined $REQUEST->{whisker}
	|| ref($REQUEST->{whisker}) ne 'HASH'){
	mok(0, 'http_new_request pre-check');
} else {
	mok(1, 'http_new_request pre-check');
	LW2::uri_split($INITIAL_DATA,$REQUEST);
	foreach (keys %EXPECTED){
		if(!defined $REQUEST->{whisker}->{$_} ||
			$REQUEST->{whisker}->{$_} ne $EXPECTED{$_}){
			$BAD++;
		}
	}
}
mok(!$BAD, 'uri_split (hash)');


$INITIAL_DATA = "http://server2/page2";
%EXPECTED = ( uri => '/page2', port => 80, host=>'server2',
	ssl=>0, uri_user=>undef, uri_password=>undef,
	parameters=>undef );

@temp = LW2::uri_split($INITIAL_DATA,$REQUEST);
$BAD=0;
foreach (keys %EXPECTED){
	if(defined $EXPECTED{$_}){
		if(!defined $REQUEST->{whisker}->{$_} ||
			$REQUEST->{whisker}->{$_} ne $EXPECTED{$_}){
			$BAD++;
		}
	} else {
		if(defined $REQUEST->{whisker}->{$_}){
			$BAD++;
		}
	}
}
mok(!$BAD, 'uri_split no.2');

$temp = LW2::uri_join(@temp);
mok($temp eq $INITIAL_DATA, 'uri_join');

$BAD=0;
$BADSTR='';
@INITIAL_DATA = (
	['http://server/dir/page', 'http://baseserver/basedir/basepage', 'http://server/dir/page'],
	['/otherdir/page2', 'http://baseserver/basedir/basepage', 'http://baseserver/otherdir/page2'],
	['page3','http://baseserver/basedir/basepage','http://baseserver/basedir/page3'],
	['/dir/page','/otherdir/otherpage','/dir/page'],
	['/dir/page','http://server','http://server/dir/page'],
	['page','http://server','http://server/page'],
	['page','http://server/stuff','http://server/page'],
	['/dir1/page1','http://server/dir1/page2','http://server/dir1/page1'],
	['dir/dir2/page','http://server/page','http://server/dir/dir2/page'],
	['dir/dir2/page','http://server/dir/','http://server/dir/dir/dir2/page'],
	['dir/page','http://server:81','http://server:81/dir/page'],
	['dir/page','http://server:81/x/y','http://server:81/x/dir/page'],
	['/dir/page','https://server/x/y','https://server/dir/page'],
	['/dir/page','http://server:81/x/y','http://server:81/dir/page'],
	['/dir/page','http://server:81','http://server:81/dir/page'],
	['/dir1/page1?param1','http://server/','http://server/dir1/page1?param1'],
	['/dir1/page1?param1','http://server/dir2/page2?param2','http://server/dir1/page1?param1'],
	['dir1/page1?param1','http://server/dir2/page2?param2','http://server/dir2/dir1/page1?param1'],
	['page1?param1','http://server/dir/?param2','http://server/dir/page1?param1'],
	['page1?param1','http://server/dir/#param2','http://server/dir/page1?param1']
);
for($x=0; $x<~~@INITIAL_DATA; $x++){
	$temp = LW2::uri_absolute($INITIAL_DATA[$x]->[0], $INITIAL_DATA[$x]->[1]);
	if($temp ne $INITIAL_DATA[$x]->[2]){
		$BAD++;
		$BADSTR.=",$x";
	}
}
mok(!$BAD, 'uri_join no.2 '.$BADSTR);

$BAD=0;
$BADSTR='';
@INITIAL_DATA = (
	['http://server/page','http://server/page'],
	['http://server\\page','http://server/page'],
	['http://server/dir1/../dir2/page','http://server/dir2/page'],
	['http://server/dir1/../','http://server/'],
	['http://server/dir1/dir2/..','http://server/dir1/'],
	['http://server/././././././dir1/page','http://server/dir1/page'],
	['http://server/..\\../../x/..\\..\\../a/b','http://server/a/b'],
	['/..\\..///../x//..\\/..\\../a/b','/a/b'],
	['http://server','http://server/'],
	['/dir/.test/page','/dir/.test/page'],
	['/dir/././././','/dir/'],
	['/dir/././././p','/dir/p'],
	['/dir/./././../p','/p'],
	['/d/./p?x','/d/p?x'],
	['/d/./?x','/d/?x'],
	['/d/../?x','/?x'],
	['/d?x/../','/d?x/../'],
	['/d/..#x','/#x']
);
for($x=0; $x<~~@INITIAL_DATA; $x++){
	$temp = LW2::uri_normalize($INITIAL_DATA[$x]->[0],1);
	if($temp ne $INITIAL_DATA[$x]->[1]){
		$BAD++;
		$BADSTR.=", $x";
	}
}
mok(!$BAD, 'uri_normalize '.$BADSTR);

$BAD=0;
$BADSTR='';
@INITIAL_DATA = (
	['page',''],
	['d/','d/'],
	['/d/','/d/'],
	['/d/?p','/d/'],
	['/d/#f','/d/'],
	['/a/b/c','/a/b/'],
	['/a','/']
);
for($x=0; $x<~~@INITIAL_DATA; $x++){
	$temp = LW2::uri_get_dir($INITIAL_DATA[$x]->[0]);
	if($temp ne $INITIAL_DATA[$x]->[1]){
		$BAD++;
		$BADSTR.=", $x";
	}
}
mok(!$BAD, 'uri_get_dir '.$BADSTR);

$BAD=0;
$BADSTR='';
@INITIAL_DATA = (
	['/a/b','/a/b'],
	['/a/b/','/a/b/'],
	['/;foo','/'],
	[';foo/','/'],
	[';foo',''],
	['/a/b;foo','/a/b'],
	['/a;foo/b','/a/b'],
	[';foo/a/','/a/'],
	['/;foo/;bar/b;baz/','///b/'],
	['/b///','/b///']
);
for($x=0; $x<~~@INITIAL_DATA; $x++){
	$temp = LW2::uri_strip_path_parameters($INITIAL_DATA[$x]->[0]);
	if($temp ne $INITIAL_DATA[$x]->[1]){
		$BAD++;
		$BADSTR.=", $x";
	}
}
mok(!$BAD, 'uri_strip_path_parameters '.$BADSTR);

$BAD=0;
$INITIAL_DATA = '/a;A=1/b;B=2/c;C=3/;D=4';
%EXPECTED = (a=>'A=1',b=>'B=2',c=>'C=3',''=>'D=4');
$temp = {};
LW2::uri_strip_path_parameters($INITIAL_DATA, $temp);
foreach( keys %EXPECTED ){
	if(!defined $temp->{$_} || $temp->{$_} ne $EXPECTED{$_}){
		$BAD++;
	}
	delete $temp->{$_};
}
if(keys %$temp){
	$BAD++;
}
mok(!$BAD,'uri_strip_path_parameters');

$BAD=0;
$BADSTR='';
@INITIAL_DATA = (
	['a=1',{a=>'1'}],
	['a',{a=>undef}],
	['&',{}],
	['=',{''=>''}],
	['=1',{''=>'1'}],
	['a=1&b=2&c=3',{a=>'1',b=>'2',c=>'3'}],
	['a=1&b&c=3',{a=>'1',b=>undef,c=>'3'}],
	['a=1&a=2&a=3',{a=>'3'}]
);
for($x=0; $x<~~@INITIAL_DATA; $x++){
	$temp = LW2::uri_parse_parameters( $INITIAL_DATA[$x]->[0] );
	foreach (keys %{ $INITIAL_DATA[$x]->[1] }){
		if(!exists $temp->{$_} || 
				$temp->{$_} ne $INITIAL_DATA[$x]->[1]->{$_}){
			$BAD++;
			$BADSTR.=", $x";
		}
		delete $temp->{$_};
	}
	if(keys %$temp){
		$BAD++;
	}
}
mok(!$BAD,'uri_parse_parameters '.$BADSTR);

$BAD=0;
$INITIAL_DATA = 'a=a1&b=b1&a=a2';
$temp = LW2::uri_parse_parameters( $INITIAL_DATA, 0, 1 );
if(!defined $temp->{a} || ref($temp->{a}) ne 'ARRAY' ||
	$temp->{a}->[0] ne 'a1' || $temp->{a}->[1] ne 'a2' ||
	!defined $temp->{b} || $temp->{b} ne 'b1'){
	$BAD++;
}
mok(!$BAD, 'uri_parse_parameters no.2');

$BADe=0;
$BADSTRe='';
$BADu=0;
$BADSTRu='';
@INITIAL_DATA = (
	['abc','abc'],
	['a%b','a%25b'],
	['a b','a+b'],
	['a    b','a++++b'],
	['a%%%b','a%25%25%25b'],
	['a%25b','a%2525b'],
	['a+b','a%2bb'],
	['+=?&#@;\\/','%2b%3d%3f%26%23%40%3b%5c%2f'],
	["ab\x00cd",'ab%00cd'],
	["ab%\xff\xfe\x01cd",'ab%25%ff%fe%01cd']
);
for($x=0; $x<~~@INITIAL_DATA; $x++){
	$temp = LW2::uri_escape($INITIAL_DATA[$x]->[0]);
	if($temp ne $INITIAL_DATA[$x]->[1]){
		$BADe++;
		$BADSTRe .= ", $x";
	}
	$temp = LW2::uri_unescape($INITIAL_DATA[$x]->[1]);
	if($temp ne $INITIAL_DATA[$x]->[0]){
		$BADu++;
		$BADSTRu .= ", $x";
	}
}
mok(!$BADe, 'uri_escape '.$BADSTRe);
mok(!$BADu, 'uri_unescape '.$BADSTRu);


##########################################################################
# utils functions

# TODO: utils_recperm

# utils_array_shuffle, assumed good, nothing to test

%temp = (); # this test can fail, since it could randomly come up with
						# the same string, but the odds of that should be low
$BAD=0;
for(0..5){
	$temp = LW2::utils_randstr();
	if(length($temp)<10 || length($temp)>20 || 
			$temp=~tr/A-Za-z0-9//c || defined $temp{$temp}){
		$BAD++;
	}
	$temp{$temp}++;
}
mok(!$BAD,'utils_randstr');

$temp = LW2::utils_randstr(11,'ABC');
if(length($temp) != 11 || $temp=~tr/ABC//c){
	print "FAIL: utils_randstr (w/ params)\n";
}

# utils_port_open...


$BAD=0;
%INITIAL_DATA = ( AAA=>'Aaa', bBb=>'Bbb', ccc=>'Ccc', 
	'123'=>'Ddd', ''=>'Eee', "\x80\xc0\xd0\xff"=>'Fff' );
%EXPECTED = ( aaa=>'Aaa', bbb=>'Bbb', ccc=>'Ccc', 
	'123'=>'Ddd', ''=>'Eee', "\x80\xc0\xd0\xff"=>'Fff' );
$temp = LW2::utils_lowercase_keys(\%INITIAL_DATA);
if($temp != 2){
	$BAD++;
}
foreach (keys %EXPECTED){
	if(!defined $INITIAL_DATA{$_} || $INITIAL_DATA{$_} ne $EXPECTED{$_}){
		$BAD++;
	}
	delete $INITIAL_DATA{$_};
}
if(keys %INITIAL_DATA){
	$BAD++;
}
mok(!$BAD,'utils_lowercase_keys');


%INITIAL_DATA = ( AAA=>'1', Aaa=>'2', ''=>'3', BBB=>'4', M=>['X','Y'], m=>'Z' );
$temp = LW2::utils_find_key( \%INITIAL_DATA, 'foo' );
mok(!defined $temp, 'utils_find_key no.1');

$temp = LW2::utils_find_key( \%INITIAL_DATA, '' );
mok(defined $temp && !ref($temp) && $temp eq '3', 'utils_find_key no.2');

@temp = LW2::utils_find_key( \%INITIAL_DATA, '' );
mok(defined $temp[0] && $temp[0] eq '3' && ~~@temp == 1, 'utils_find_key no.3');

$temp = LW2::utils_find_key( \%INITIAL_DATA, 'M' );
mok(!(!defined $temp || ref($temp) ne 'ARRAY' || 
	$temp->[0] ne 'X' || $temp->[1] ne 'Y' || ~~@$temp != 2),'utils_find_key no.4');

@temp = LW2::utils_find_key( \%INITIAL_DATA, 'M' );
mok(!(!defined $temp[0] || !defined $temp[1] || 
	$temp[0] ne 'X' || $temp[1] ne 'Y' || ~~@temp != 2),'utils_find_key no.5');

$temp = LW2::utils_find_lowercase_key( \%INITIAL_DATA, 'bbb' );
mok(!(!defined $temp || $temp ne '4'),'utils_find_lowercase_key no.1');

$temp = LW2::utils_find_lowercase_key( \%INITIAL_DATA, 'BBB' );
mok(!(!defined $temp || $temp ne '4'), 'utils_find_lowercase_key no.2');

@temp = LW2::utils_find_lowercase_key( \%INITIAL_DATA, 'BbB' );
mok(!(!defined $temp[0] || $temp[0] ne '4' || ~~@temp != 1),
	'utils_find_lowercase_key no.3');

$BAD=0;
$temp = LW2::utils_find_lowercase_key( \%INITIAL_DATA, 'aAa' );
if(!defined $temp || ref($temp) ne 'ARRAY'){
	$BAD++;
} else {
	@temp = sort @$temp;
	if($temp[0] ne '1' || $temp[1] ne '2' || ~~@$temp != 2){
		$BAD++;
	}
}
mok(!$BAD,'utils_find_lowercase_key no.4');

$BAD=0;
$temp = LW2::utils_find_lowercase_key( \%INITIAL_DATA, 'm' );
if(!defined $temp || ref($temp) ne 'ARRAY'){
	$BAD++;
} else {
	@temp = sort @$temp;
	if($temp[0] ne 'X' || $temp[1] ne 'Y' || $temp[2] ne 'Z' || ~~@$temp != 3){
		$BAD++;
	}
}
mok(!$BAD,'utils_find_lowercase_key no.5');

%INITIAL_DATA = ( AAA=>'1', Aaa=>'2', ''=>'3', BBB=>'4', M=>['X','Y'], m=>'Z' );
$temp = LW2::utils_delete_lowercase_key( \%INITIAL_DATA, 'm' );
mok(!($temp != 2 || defined $INITIAL_DATA{'M'} || defined $INITIAL_DATA{'m'} ),
	'utils_delete_lowercase_key no.1');

$temp = LW2::utils_delete_lowercase_key( \%INITIAL_DATA, 'AaA' );
mok(!($temp != 2 || defined $INITIAL_DATA{'AAA'} || 
	defined $INITIAL_DATA{'Aaa'} || $INITIAL_DATA{''} ne '3' ||
	$INITIAL_DATA{'BBB'} ne '4'), 'utils_delete_lowercase_key no.2');


$BAD=0;
$INITIAL_DATA = "A\nB\nC";
$temp = LW2::utils_getline(\$INITIAL_DATA);
$BAD++ if($temp ne 'A');
$temp = LW2::utils_getline(\$INITIAL_DATA);
$BAD++ if($temp ne 'B');
$temp = LW2::utils_getline(\$INITIAL_DATA);
$BAD++ if(defined $temp);
$temp = LW2::utils_getline(\$INITIAL_DATA);
$BAD++ if(defined $temp);
$temp = LW2::utils_getline(\$INITIAL_DATA, 0);
$BAD++ if($temp ne 'A');
mok(!$BAD,'utils_getline');

$BAD=0;
$INITIAL_DATA = "A\r\nB\r\nC\rD\r\nE\nF\r\nG";
$temp = LW2::utils_getline_crlf(\$INITIAL_DATA);
$BAD++ if($temp ne 'A');
$temp = LW2::utils_getline_crlf(\$INITIAL_DATA);
$BAD++ if($temp ne 'B');
$temp = LW2::utils_getline_crlf(\$INITIAL_DATA);
$BAD++ if($temp ne "C\rD");
$temp = LW2::utils_getline_crlf(\$INITIAL_DATA);
$BAD++ if($temp ne "E\nF");
$temp = LW2::utils_getline_crlf(\$INITIAL_DATA);
$BAD++ if(defined $temp);
$temp = LW2::utils_getline_crlf(\$INITIAL_DATA);
$BAD++ if(defined $temp);
$temp = LW2::utils_getline_crlf(\$INITIAL_DATA, 0);
$BAD++ if($temp ne 'A');
mok(!$BAD,'utils_getline_crlf');


# TODO: utils_save_page


@ARGV = ('-a','-b','bval','c','-d','','e');
%temp = ();
LW2::utils_getopts('ab:d:', \%temp );
mok(!(!defined $temp{a} || $temp{a} ne '1' ||
	!defined $temp{b} || $temp{b} ne 'bval' ||
	!defined $temp{d} || $temp{d} ne '' ||
	keys %temp != 3 || $ARGV[0] ne 'c' || $ARGV[1] ne 'e'),
	'utils_getopts');
@ARGV = ();

$INITIAL_DATA = "abc def ghijkl";
$temp = LW2::utils_text_wrapper($INITIAL_DATA, "\n", 4);
mok($temp eq "abc\ndef\nghij\nkl\n", 'utils_text_wrapper');

# TODO: brute URL

$BAD=0;
$temp = LW2::utils_join_tag('TAG', { 'a'=>'b' });
$BAD++ if($temp ne '<TAG a="b">');
$temp = LW2::utils_join_tag('TAG', { 'a'=>undef });
$BAD++ if($temp ne '<TAG a>');
$temp = LW2::utils_join_tag('TAG', { 'a'=>'1', 'b'=>'2' });
$BAD++ if($temp ne '<TAG a="1" b="2">' && $temp ne '<TAG b="2" a="1">');
$temp = LW2::utils_join_tag('TAG', {} );
$BAD++ if($temp ne '<TAG>');
mok(!$BAD, 'utils_join_tag');


%INITIAL_DATA = ( 'A'=>'aa', 'B'=>'bb', 
	'whisker'=>{ 'MAGIC'=>31339, 'C'=>'cc' } );
%temp = ();
LW2::utils_request_clone( \%INITIAL_DATA, \%temp );
mok(!(!defined $temp{A} || $temp{A} ne 'aa' ||
	!defined $temp{B} || $temp{B} ne 'bb' ||
	!defined $temp{whisker} || ref($temp{whisker}) ne 'HASH' ||
	!defined $temp{whisker}->{MAGIC} || $temp{whisker}->{MAGIC} != 31339 ||
	!defined $temp{whisker}->{C} || $temp{whisker}->{C} ne 'cc' ||
	keys %temp != 3 || keys %{$temp{whisker}} != 2), 'utils_request_clone no.1');

%INITIAL_DATA = ( 'A'=>['a1','a2'], 
	'whisker'=>{ 'MAGIC'=>31339, 'C'=>['c1','c2'] } );
%temp = ();
LW2::utils_request_clone( \%INITIAL_DATA, \%temp );
mok(!(!defined $temp{A} || ref($temp{A}) ne 'ARRAY' ||
	$temp{A}->[0] ne 'a1' || $temp{A}->[1] ne 'a2' ||
	!defined $temp{whisker} || ref($temp{whisker}) ne 'HASH' ||
	!defined $temp{whisker}->{MAGIC} || $temp{whisker}->{MAGIC} != 31339 ||
	!defined $temp{whisker}->{C} || ref($temp{whisker}->{C}) ne 'ARRAY' ||
	$temp{whisker}->{C}->[0] ne 'c1' || $temp{whisker}->{C}->[1] ne 'c2' ||
	keys %temp != 2 || keys %{$temp{whisker}} != 2), 'utils_request_clone no.2');


# TODO: utils_request_fingerprint

# utils_flatten_lwhash
# - test http_req2line/http_resp2line and http_construct_headers first


##########################################################################
# cookie functions

#
# NOTE! we cheat and use internal knowledge that the cookie jar is just
# a hash reference, in order to verify it's contents for these tests.
# Ideally, we wouldn't be making that assumption, as the structure and
# contents of the cookie jar are subject to change.
#

$JAR = LW2::cookie_new_jar();
mok( defined $JAR && ref($JAR) eq 'HASH', 'cookie_new_jar');
die("Need working cookie jar") if(!defined $JAR || ref($JAR) ne 'HASH');

@INITIAL_DATA = (
	['cook1','v1','d1.com','/u','2',0,'v1','d1.com','/u',undef,0],
	['cook2','v2','d2.com',undef,undef,0,'v2','d2.com','/',undef,0],
	['cook3','v3',undef,undef,undef,1,'v3',undef,'/',undef,1],
	['$cook4','v4',undef,undef,undef,undef,'v4',undef,'/',undef,0],
);
$BAD=0;
foreach (@INITIAL_DATA){
	LW2::cookie_set( $JAR, $_->[0],
		$_->[1], $_->[2], $_->[3], $_->[4], $_->[5]);
	if(defined $JAR->{$_->[0]} && ref($JAR->{$_->[0]}) eq 'ARRAY'){
		for($x=0;$x<5;$x++){
			if( defined $JAR->{$_->[0]}->[$x] && defined $_->[6+$x]){
				$BAD++ if($JAR->{$_->[0]}->[$x] ne $_->[6+$x]);
			} else {
				$BAD++ if(defined $JAR->{$_->[0]}->[$x] && !defined $_->[6+$x]);
				$BAD++ if(!defined $JAR->{$_->[0]}->[$x] && defined $_->[6+$x]);
			}
		}
	} else {
		$BAD++;
	}
	delete $JAR->{$_->[0]};
	$BAD++ if(keys %$JAR);
}
mok(!$BAD,'cookie_set no.1');

LW2::cookie_set( $JAR, '', 'val', undef, undef, undef, undef );
mok( !keys %$JAR, 'cookie_set no.2');

$JAR = LW2::cookie_new_jar();
LW2::cookie_set($JAR,'cook','val');
$BAD=0;
if(defined $JAR->{cook}){
	LW2::cookie_set($JAR,'cook','');
	$BAD++ if(defined $JAR->{cook});
} else {
	$BAD++;
}
mok(!$BAD,'cookie_set no.3');

$JAR = LW2::cookie_new_jar();
LW2::cookie_set($JAR,'cook','val');
$BAD=0;
if(defined $JAR->{cook}){
	LW2::cookie_set($JAR,'cook',undef);
	$BAD++ if(defined $JAR->{cook});
} else {
	$BAD++;
}
mok(!$BAD,'cookie_set no.4');

$JAR = LW2::cookie_new_jar();
LW2::cookie_set($JAR,'A','a');
LW2::cookie_set($JAR,'B','b');
LW2::cookie_set($JAR,'C','c');
@temp = LW2::cookie_get_names($JAR);
@temp = sort @temp;
mok( ($temp[0] eq 'A' && $temp[1] eq 'B' && $temp[2] eq 'C'),
	'cookie_get_names');


$JAR = LW2::cookie_new_jar();
@INITIAL_DATA = (
	['cook1','v1','d1.com','/u','2',0,'v1','d1.com','/u',undef,0],
	['cook2','v2','d2.com',undef,undef,0,'v2','d2.com','/',undef,0],
	['cook3','v3',undef,undef,undef,1,'v3',undef,'/',undef,1],
	['$cook4','v4',undef,undef,undef,undef,'v4',undef,'/',undef,0],
);
$BAD=0;
foreach (@INITIAL_DATA){
	LW2::cookie_set( $JAR, $_->[0],
		$_->[1], $_->[2], $_->[3], $_->[4], $_->[5]);
	@temp = LW2::cookie_get($JAR, $_->[0]);
	if( !defined $temp[0] ){
		$BAD++;
		next;
	}
	for($x=0;$x<5;$x++){
		if( defined $temp[$x] && defined $_->[6+$x]){
			$BAD++ if($temp[$x] ne $_->[6+$x]);
		} else {
			$BAD++ if(defined $temp[$x] && !defined $_->[6+$x]);
			$BAD++ if(!defined $temp[$x] && defined $_->[6+$x]);
		}
	}
}
mok(!$BAD,'cookie_get');



@INITIAL_DATA = (
	['c=v',undef,undef,
			'c','v',undef,'/',undef,0],
	['c=v;',undef,undef,
			'c','v',undef,'/',undef,0],
	['$c=v',undef,undef,
			'$c','v',undef,'/',undef,0],
	['c = v',undef,undef,
			'c','v',undef,'/',undef,0],
	['     c   =  v    ',undef,undef,
			'c','v',undef,'/',undef,0],
	[' c =  v   ;  ',undef,undef,
			'c','v',undef,'/',undef,0],
	['c="v"',undef,undef,
			'c','v',undef,'/',undef,0],
	[' c = "v" ',undef,undef,
			'c','v',undef,'/',undef,0],

	['c=v; path=/',undef,undef,
			'c','v',undef,'/',undef,0],
	['c=v; path=',undef,undef,
			'c','v',undef,'/',undef,0],
	['c=v; path=""',undef,undef,
			'c','v',undef,'/',undef,0],
	['c=v; path=/a/',undef,undef,
			'c','v',undef,'/a',undef,0],
	['c=v; path=/a; path=/b',undef,undef,
			'c','v',undef,'/a',undef,0],
	['c=v; path=a',undef,undef,
			'c','v',undef,'/',undef,0],
	['c=v',undef,'/foo',
			'c','v',undef,'/foo',undef,0],
	['c=v',undef,'',
			'c','v',undef,'/',undef,0],
	['c=v',undef,'/foo/',
			'c','v',undef,'/foo',undef,0],
	['c=v',undef,'foo/',
			'c','v',undef,'/',undef,0],
	['c=v; path=a',undef,'b',
			'c','v',undef,'/',undef,0],
	['c=v; path=/a',undef,'/b',
			'c','v',undef,'/a',undef,0],
	['c=v; path=a',undef,'/b',
			'c','v',undef,'/',undef,0],

	['c=v; domain=.foo.com',undef,undef,
			'c','v','.foo.com','/',undef,0],
	['c=v; domain=.foo.com.',undef,undef,
			'c','v','.foo.com','/',undef,0],
	['c=v; domain=foo.com',undef,undef,
			'c','v','.foo.com','/',undef,0],
	['c=v; domain=a.com; domain=b.com',undef,undef,
			'c','v','.a.com','/',undef,0],
	['c=v; domain=1.2.3.4',undef,undef,
			'c','v','1.2.3.4','/',undef,0],
	['c=v','foo.com',undef,
			'c','v','foo.com','/',undef,0],
	['c=v','.foo.com',undef,
			'c','v','.foo.com','/',undef,0],
	['c=v','1.2.3.4',undef,
			'c','v','1.2.3.4','/',undef,0],
	['c=v; domain=','foo.com',undef,
			'c','v','foo.com','/',undef,0],
	['c=v; domain=""','foo.com',undef,
			'c','v','foo.com','/',undef,0],
	['c=v; domain=a.com','b.com',undef,
			'c','v','.a.com','/',undef,0],

	['c=v; max-age=5',undef,undef,
			'c','v',undef,'/',undef,0],

	['c=v; misc=foobar',undef,undef,
			'c','v',undef,'/',undef,0],

	['c=v; secure',undef,undef,
			'c','v',undef,'/',undef,1],
	['c=v; secure; secure',undef,undef,
			'c','v',undef,'/',undef,1],

	[' c = v ; domain = a.com ; path = /a/ ; ', undef, undef,
			'c','v','.a.com','/a',undef,0],
	[' c = "v" ; domain = "a.com" ; path = "/a/" ; ', undef, undef,
			'c','v','.a.com','/a',undef,0],
	['c=v; domain="a.com" b.com', undef, undef,
			'c','v','.a.com','/',undef,0],
	['c=v; path="/a" secure;', undef, undef,
			'c','v',undef,'/a',undef,0],
	['c=v; path=/a secure; ', undef, undef,
			'c','v',undef,'/a',undef,0]



);
$BAD=0;
foreach (@INITIAL_DATA){
	$JAR = LW2::cookie_new_jar();
	LW2::cookie_parse( $JAR, $_->[0], $_->[1], $_->[2]);
	@temp = LW2::cookie_get($JAR, $_->[3]);
#print "D: ", LW2::dump($_->[0],\@temp), "\n\n";
	if( !defined $temp[0] ){
		$BAD++;
		next;
	}
	for($x=0;$x<5;$x++){
		if( defined $temp[$x] && defined $_->[4+$x]){
			$BAD++ if($temp[$x] ne $_->[4+$x]);
		} else {
			$BAD++ if(defined $temp[$x] && !defined $_->[4+$x]);
			$BAD++ if(!defined $temp[$x] && defined $_->[4+$x]);
		}
	}
}
mok(!$BAD,'cookie_parse no.1');




# cookie_read
# cookie_parse
# cookie_write
# cookie_get_valid_names




##########################################################################
# stream functions

# test buffer stream, to be able to test http stream handling functions 
# later; don't use the stream helper function until after this

$REQUEST = LW2::http_new_request();
$REQUEST->{whisker}->{buffer_stream}++;
$STREAM = LW2::stream_new($REQUEST);
$BAD=0;
if(defined $STREAM){
	$STREAM->{open}->();
	$STREAM->{queue}->('abcd');
	$BAD++ if !$STREAM->{write}->('efg');
	$BAD++ if !$STREAM->{read}->();
	$BAD++ if($STREAM->{bufin} ne 'abcdefg');
	$STREAM->{clearall}->();
	$STREAM->{queue}->('X');
	$STREAM->{queue}->('Y');
	$STREAM->{queue}->('Z');
	$BAD++ if !$STREAM->{write}->();
	$BAD++ if !$STREAM->{read}->();
	$BAD++ if($STREAM->{bufin} ne 'XYZ');
#	$BAD++ if(!(!$STREAM->{read}->() && $STREAM->{eof}==1));
	$STREAM->{close}->();
} else {
	$BAD++;
}
mok(!$BAD,'basic buffer stream testing');
die("Correctly functioning buffer streams are necessary for the rest of the tests")
	if($BAD);


##########################################################################
# html functions

# try to find evil.htm file
$SKIPHTML = 0;
$EVIL = 'evil.htm';
if(!-e $EVIL){
	$EVIL = '../docs/evil.htm' if(-e '../docs' && -d '../docs');
	$EVIL = 'docs/evil.htm' if(-e 'docs' && -d 'docs');
}
if(!-e $EVIL){
	print STDERR "# Unable to find 'evil.htm' file; HTML testing skipped\n";
	$SKIPHTML++;
}

$LITTLEHTML = '<tag param="one">';
$EVILHTML = '';
if(!$SKIPHTML){
	if(!open(IN,"<$EVIL")){
		print STDERR "# Unable to open '$EVIL'\n";
		$SKIPHTML++;
	} else {
		binmode(IN);
		while(<IN>){
			$EVILHTML .= $_; }
		close(IN);
	}
}

sub test_html_minimum {
	my $tag = shift;
	my $hr = shift;
	$HTML_BAD++;
	$HTML_BAD++ if($tag ne 'tag');
	$HTML_BAD++ if(!defined $hr->{param});
	$HTML_BAD++ if(defined $hr->{param} &&
		$hr->{param} ne 'one');
	return;
}

sub test_html_minimum2 {
	my ($tag, $hr, $data, $start, $len, $fr) = @_;
	test_html_minimum($tag, $hr);
	$HTML_BAD++ if($$data ne $LITTLEHTML);
	$HTML_BAD++ if($start != 0);
	$HTML_BAD++ if($len != length($LITTLEHTML));
	$HTML_BAD++ if(!ref($fr) || ref($fr) ne 'ARRAY' ||
		$fr->[0] != 12345);
}

sub test_html_taglist {
	my ($tag, $hr, $data, $start, $len, $fr) = @_;

	return if($HTML_BAD);

	my $T = shift @$fr;

	if(!defined $T){
		$HTML_BAD++;
		return;
	}
	
	if($tag ne $T->[0]){
#print STDERR "HD: current tag '$tag' != expected '", $T->[0], "'\n";
#print STDERR "HD2: next expected '", $fr->[0]->[0], "'\n";
#print STDERR "HD3: next expected '", $fr->[1]->[0], "'\n";
		$HTML_BAD++;
		return;
	}

	if($tag eq '!--'){
		# comments are special
		return if($T->[1]->{'='} eq '*');
	}
	
	foreach(keys %$hr){
		if(!exists $T->[1]->{$_} ||
			$hr->{$_} ne $T->[1]->{$_}){
#print STDERR "HD: bad value for $_\n";
#print STDERR LW2::dump('d', \$T->[1]->{$_}), "\n";
#print STDERR LW2::dump('a', \$hr->{$_}), "\n";
			$HTML_BAD++;
			return;
		}
		delete $T->[1]->{$_};
	}

	if(~~ keys %{ $T->[1] } > 0){
#print STDERR "HD: left over values\n";
		$HTML_BAD++;
		return;
	}
	
}


# really basic test to make sure parser parses a tag
$HTML_BAD = 0;
LW2::html_find_tags(\$LITTLEHTML, \&test_html_minimum);
mok($HTML_BAD==1,'html_find_tags no.1');

# test to make sure all callback parts work
$HTML_BAD = 0;
$FUNCREF = [ 12345 ];
LW2::html_find_tags(\$LITTLEHTML, \&test_html_minimum2, 0, $FUNCREF);
mok($HTML_BAD==1,'html_find_tags no.2');

# verify taglist test function works as expected
$TAGLIST = [ ['tag', { param=>'one' } ] ];
$HTML_BAD = 0;
LW2::html_find_tags(\$LITTLEHTML, \&test_html_taglist, 0, $TAGLIST);
mok($HTML_BAD==0,'html_find_tags no.3');

# basic taglist #1
$TAGLIST = [ ['a',{}],['b',{}],['c',{}],['/c',{}],['/b',{}],['/a',{}] ];
$HTML_BAD = 0;
$TAGHTML = '<a><b><c></c></b></a>';
LW2::html_find_tags(\$TAGHTML, \&test_html_taglist, 0, $TAGLIST);
mok($HTML_BAD==0,'html_find_tags no.4');

# basic taglist #2
$TAGLIST = [ ['!--',{'='=>'-x '}],['script',{'='=>
	'foo<b>bar<i>baz'}],['foo',{}] ];
$HTML_BAD = 0;
$TAGHTML = '<!---x --><script>foo<b>bar<i>baz</script><foo>';
LW2::html_find_tags(\$TAGHTML, \&test_html_taglist, 0, $TAGLIST);
mok($HTML_BAD==0,'html_find_tags no.5');


# now do evil.htm test
if(!$SKIPHTML){

  $TAGLIST = [
	['html',{}], 
	['body',{}], 
	['p',{}],
	['script',{language=>'javascript','='=>
	"\r\ndocument.writeln(\"Do not parse <this> or <that> tag!\");\r\ndocument.writeln(\"<!-- no comments! Do not stop! /script>\");\r\n"}],
	['!--',{'='=>' overzealous '}],
	['<a',{href=>'http://localhost/"'}],
	['/a',{}],
	['/a',{}],
	['form',{method=>'post',action=>'/script'}],
	['input',{type=>'text',value=>'work?'}],
	['input',{type=>'text',value=>"Don't use <blink> \r\n\tanywhere"}],
	['input',{value=>'Continue -->',type=>'submit'}],
	['!--',{'='=>"\r\nDon't stop - -> and -- and -> and <!--\r\nstopped?\r\n"}],
	['!--',{'='=>' nope '}],
	['/a',{}],
	['i',{}], 
	['b',{}], 
	['/i</b',{}], 
	['a',{href=>'/blah>bold?</ a><b/>'}],
	['/a',{}],
	['a',{href=>'/blah"'}], 
	['/a',{}], 
	['/a',{}],
	['/body',{}],
	['/html',{}],
	['!--',{'='=>'*'}]
  ];

  $HTML_BAD = 0;
  LW2::html_find_tags(\$EVILHTML, \&test_html_taglist, 0, $TAGLIST);
  mok($HTML_BAD==0,'html_find_tags no.6');

} else {
	# we need to skip the tests so Test[::Simple] don't complain
	# note: keep the number below in sync with the number of tests
	# above
	for(1..1){
		mok(0, 'html_find_tags placeholder for skipped tests');
	}
}

# tagmap support
$TAGLIST = [ ['b',{}],['c',{}],['/b',{}] ];
$TAGMAP = { 'b'=>1, 'c'=>1, '/b'=>1 };
$HTML_BAD = 0;
$TAGHTML = '<a><b><c></c></b></a>';
LW2::html_find_tags(\$TAGHTML, \&test_html_taglist, 0, $TAGLIST, $TAGMAP);
mok($HTML_BAD==0,'html_find_tags no.7');


##########################################################################
# form functions

# try to find forms.htm file
$SKIPFORM = 0;
$EVIL = 'forms.htm';
if(!-e $EVIL){
	print STDERR "# Unable to find 'forms.htm' file; form testing skipped\n";
	$SKIPFORM++;
}

$EVILFORM = '';
if(!$SKIPFORM){
	if(!open(IN,"<$EVIL")){
		print STDERR "# Unable to open '$EVIL'\n";
		$SKIPFORM++;
	} else {
		binmode(IN);
		while(<IN>){
			$EVILFORM .= $_; }
		close(IN);
	}
}

if(!$SKIPFORM){
	$res = LW2::forms_read(\$EVILFORM);
	#print LW2::dump("res", $res), "\n";
} else {
	for(1..1){
		mok(0, 'form placeholder for skipped tests');
	}
}


##########################################################################
# http functions

@EXPECTED = qw( http_space1 http_space2 version method protocol port
	uri uri_prefix uri_postfix uri_param_sep host timeout
	include_host_in_uri ignore_duplicate_headers normalize_incoming_headers
	lowercase_incoming_headers require_newline_after_headers
	invalid_protocol_return_value ssl ssl_save_info http_eol
	force_close force_open retry trailing_slurp force_bodysnatch
	max_size);

$BAD=0;
$REQUEST = LW2::http_new_request();
$BAD++ if(ref($REQUEST) ne 'HASH' || 
	!defined $REQUEST->{whisker} ||
	ref($REQUEST->{whisker}) ne 'HASH' || 
	!defined $REQUEST->{whisker}->{MAGIC} ||
	$REQUEST->{whisker}->{MAGIC} != 31339);
foreach (@EXPECTED){
	$BAD++ if(!defined $REQUEST->{whisker}->{$_});
}
mok(!$BAD,'http_new_request no.1');

$RESPONSE = LW2::http_new_response();
mok(!(ref($RESPONSE) ne 'HASH' || 
	!defined $RESPONSE->{whisker} ||
	ref($RESPONSE->{whisker}) ne 'HASH' || 
	!defined $RESPONSE->{whisker}->{MAGIC} ||
	$RESPONSE->{whisker}->{MAGIC} != 31340),
	'http_new_response');

$BAD=0;
($REQUEST,$RESPONSE) = LW2::http_new_request();
$BAD++ if(ref($REQUEST) ne 'HASH' || 
	!defined $REQUEST->{whisker} ||
	ref($REQUEST->{whisker}) ne 'HASH' || 
	!defined $REQUEST->{whisker}->{MAGIC} ||
	$REQUEST->{whisker}->{MAGIC} != 31339);
foreach (@EXPECTED){
	$BAD++ if(!defined $REQUEST->{whisker}->{$_});
}
$BAD++ if(ref($RESPONSE) ne 'HASH' || 
	!defined $RESPONSE->{whisker} ||
	ref($RESPONSE->{whisker}) ne 'HASH' || 
	!defined $RESPONSE->{whisker}->{MAGIC} ||
	$RESPONSE->{whisker}->{MAGIC} != 31340);
mok(!$BAD,'http_new_request no.2');

$REQUEST = LW2::http_new_request('A'=>1, 'b'=>2, 'Cc'=>3);
mok(!(!defined $REQUEST->{whisker}->{A} ||
	$REQUEST->{whisker}->{A} ne 1 ||
	!defined $REQUEST->{whisker}->{b} ||
	$REQUEST->{whisker}->{b} ne 2 ||
	!defined $REQUEST->{whisker}->{Cc} ||
	$REQUEST->{whisker}->{Cc} ne 3 ), 'http_new_request no.3');

$TEMP = {};
$BAD=0;
LW2::http_init_request($TEMP);
if(!defined $TEMP->{whisker} || ref($TEMP->{whisker}) ne 'HASH'){
	$BAD++;
} else {
	foreach (@EXPECTED){
		$BAD++ if(!defined $TEMP->{whisker}->{$_});
	}
	$BAD++ if(!defined $TEMP->{whisker}->{MAGIC} ||
		$TEMP->{whisker}->{MAGIC} != 31339);
}
mok(!$BAD, 'http_init_request');


$REQUEST = LW2::http_new_request();
$REQUEST->{whisker}->{method}='A';
$REQUEST->{whisker}->{http_space1}='B';
$REQUEST->{whisker}->{uri_prefix}='C';
$REQUEST->{whisker}->{uri}='D';
$REQUEST->{whisker}->{uri_postfix}='E';
$REQUEST->{whisker}->{http_space2}='F';
$REQUEST->{whisker}->{protocol}='G';
$REQUEST->{whisker}->{version}='H';
$REQUEST->{whisker}->{http_eol}='I';
$temp = LW2::http_req2line($REQUEST);
mok($temp eq 'ABCDEFG/HI', 'http_req2line no.1');

$REQUEST->{whisker}->{uri_param_sep}='J';
$REQUEST->{whisker}->{parameters}='K';
$temp = LW2::http_req2line($REQUEST);
mok($temp eq 'ABCDEJKFG/HI', 'http_req2line no.2');

$REQUEST->{whisker}->{include_host_in_uri}=1;
$REQUEST->{whisker}->{uri_user}='L';
$REQUEST->{whisker}->{uri_password}='M';
$REQUEST->{whisker}->{host}='N';
$REQUEST->{whisker}->{port}=0;
$temp = LW2::http_req2line($REQUEST);
mok($temp eq 'ABhttp://L:M@N:0CDEJKFG/HI', 'http_req2line no.3');

$temp = LW2::http_req2line($REQUEST,1);
mok($temp eq 'CDEJK', 'http_req2line no.4');

$REQUEST->{whisker}->{full_request_override}='123';
$temp = LW2::http_req2line($REQUEST);
mok($temp eq '123', 'http_req2line no.5');

$RESPONSE = LW2::http_new_response();
$RESPONSE->{whisker}->{protocol}='A';
$RESPONSE->{whisker}->{version}='B';
$RESPONSE->{whisker}->{http_space1}='C';
$RESPONSE->{whisker}->{code}='D';
$RESPONSE->{whisker}->{http_space2}='E';
$RESPONSE->{whisker}->{message}='F';
$RESPONSE->{whisker}->{http_eol}='G';
$temp = LW2::http_resp2line($RESPONSE);
mok($temp eq 'A/BCDEFG', 'http_resp2line');

# _http_getline
# _http_get
# _http_getall

$REQUEST = LW2::http_new_request();
$REQUEST->{whisker}->{uri}='';
$REQUEST->{whisker}->{version}='1.1';
$REQUEST->{whisker}->{port}=81;
$REQUEST->{whisker}->{host}='sys';
LW2::utils_delete_lowercase_key($REQUEST,'connection');
LW2::utils_delete_lowercase_key($REQUEST,'host');
LW2::utils_delete_lowercase_key($REQUEST,'content-length');
LW2::utils_delete_lowercase_key($REQUEST,'content-type');
$REQUEST->{whisker}->{data}='abcde';
LW2::http_fixup_request($REQUEST);
mok(( defined $REQUEST->{Connection} && $REQUEST->{Connection} ne '' &&
	defined $REQUEST->{Host} && $REQUEST->{Host} eq 'sys:81' &&
	defined $REQUEST->{'Content-Length'} && $REQUEST->{'Content-Length'}==5 &&
	defined $REQUEST->{'Content-Type'} && $REQUEST->{'Content-Type'} ne '' &&
	$REQUEST->{whisker}->{uri} eq '/'),
	'http_fixup_request no.1');

$REQUEST = LW2::http_new_request();
$REQUEST->{whisker}->{version}='1.0';
$REQUEST->{whisker}->{port}=443;
$REQUEST->{whisker}->{ssl}=1;
$REQUEST->{whisker}->{host}='sys';
$REQUEST->{'content-Length'}=2;
$REQUEST->{'content-Type'}='foo';
$REQUEST->{whisker}->{data}='abcde';
$REQUEST->{whisker}->{proxy_host}='prx';
$REQUEST->{whisker}->{include_host_in_uri}=0;
LW2::utils_delete_lowercase_key($REQUEST,'host');
LW2::utils_delete_lowercase_key($REQUEST,'connection');
LW2::http_fixup_request($REQUEST);
mok(( defined $REQUEST->{Connection} && $REQUEST->{Connection} ne '' &&
	!defined $REQUEST->{Host} &&
	defined $REQUEST->{'Content-Length'} && $REQUEST->{'Content-Length'}==5 &&
	defined $REQUEST->{'content-Type'} && $REQUEST->{'content-Type'} eq 'foo' &&
	$REQUEST->{whisker}->{include_host_in_uri}==1),
	'http_fixup_request no.2');

# we call these in order to traverse the code path and flush out warnings
LW2::http_reset();
mok(1,'http_reset');

LW2::ssl_is_available();
mok(1,'ssl_is_available');

###################################

$INITIAL_DATA = <<EOT;
Content-Length: 5
Connection: Keep-alive
Transfer-Encoding: Identity
AAA: aval
bbB: bval
ccc: cval

EOT
@EXPECTED = ('Content-Length','Connection','Transfer-Encoding',
	'AAA','bbB','ccc');

$REQUEST = LW2::http_new_request();
$RESPONSE = LW2::http_new_response();

$REQUEST->{whisker}->{lowercase_incoming_headers} = 0;
$REQUEST->{whisker}->{normalize_incoming_headers} = 0;
$REQUEST->{whisker}->{ignore_duplicate_headers} = 0;
$REQUEST->{whisker}->{save_raw_headers} = 0;

$STREAM = buffer_stream($INITIAL_DATA);

$BAD=0;
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
if($result==1){
	$BAD++ if($enc ne 'identity');
	$BAD++ if($len != 5);
	$BAD++ if($conn ne 'keep-alive');
	$BAD++ if(!defined $RESPONSE->{'Content-Length'} 
		|| $RESPONSE->{'Content-Length'} != 5);
	$BAD++ if(!defined $RESPONSE->{'Connection'} 
		|| $RESPONSE->{'Connection'} ne 'Keep-alive');
	$BAD++ if(!defined $RESPONSE->{'Transfer-Encoding'} 
		|| $RESPONSE->{'Transfer-Encoding'} ne 'Identity');
	$BAD++ if(!defined $RESPONSE->{AAA} || $RESPONSE->{AAA} ne 'aval');	
	$BAD++ if(!defined $RESPONSE->{bbB} || $RESPONSE->{bbB} ne 'bval');	
	$BAD++ if(!defined $RESPONSE->{ccc} || $RESPONSE->{ccc} ne 'cval');	
	$BAD++ if(!defined $RESPONSE->{whisker});
	$BAD++ if(keys %$RESPONSE != 7);
	for($x=0;$x<~~@EXPECTED;$x++){
		$BAD++ if($RESPONSE->{whisker}->{header_order}->[$x] ne $EXPECTED[$x]);
	}
} else {
	$BAD++;
}
mok(!$BAD,'http_read_headers no.1');

$INITIAL_DATA = <<EOT;
AAA: aval
bbB: bval
etag: cval

EOT
$REQUEST->{whisker}->{lowercase_incoming_headers} = 1;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
@EXPECTED = ('aaa','bbb','etag');

$BAD=0;
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
if($result==1){
	$BAD++ if(defined $enc);
	$BAD++ if(defined $len);
	$BAD++ if(defined $conn);
	$BAD++ if(!defined $RESPONSE->{aaa} || $RESPONSE->{aaa} ne 'aval');	
	$BAD++ if(!defined $RESPONSE->{bbb} || $RESPONSE->{bbb} ne 'bval');	
	$BAD++ if(!defined $RESPONSE->{etag} || $RESPONSE->{etag} ne 'cval');	
	$BAD++ if(!defined $RESPONSE->{whisker});
	$BAD++ if(keys %$RESPONSE != 4);
	for($x=0;$x<~~@EXPECTED;$x++){
		$BAD++ if($RESPONSE->{whisker}->{header_order}->[$x] ne $EXPECTED[$x]);
	}
} else {
	$BAD++;
}
mok(!$BAD,'http_read_headers no.2');

$REQUEST->{whisker}->{lowercase_incoming_headers} = 0;
$REQUEST->{whisker}->{normalize_incoming_headers} = 1;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
@EXPECTED = ('Aaa','Bbb','ETag');

$BAD=0;
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
if($result==1){
	$BAD++ if(defined $enc);
	$BAD++ if(defined $len);
	$BAD++ if(defined $conn);
	$BAD++ if(!defined $RESPONSE->{Aaa} || $RESPONSE->{Aaa} ne 'aval');	
	$BAD++ if(!defined $RESPONSE->{Bbb} || $RESPONSE->{Bbb} ne 'bval');	
	$BAD++ if(!defined $RESPONSE->{ETag} || $RESPONSE->{ETag} ne 'cval');	
	$BAD++ if(!defined $RESPONSE->{whisker});
	$BAD++ if(keys %$RESPONSE != 4);
	for($x=0;$x<~~@EXPECTED;$x++){
		$BAD++ if($RESPONSE->{whisker}->{header_order}->[$x] ne $EXPECTED[$x]);
	}
} else {
	$BAD++;
}
mok(!$BAD,'http_read_headers no.3');

$INITIAL_DATA = <<EOT;
A: 1a
B: 1b
A: 2a
B: 2b

EOT
$REQUEST->{whisker}->{lowercase_incoming_headers} = 0;
$REQUEST->{whisker}->{normalize_incoming_headers} = 0;
$REQUEST->{whisker}->{ignore_duplicate_headers} = 0;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
@EXPECTED = ('A','B','A','B');

$BAD=0;
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
if($result==1){
	$BAD++ if(defined $enc);
	$BAD++ if(defined $len);
	$BAD++ if(defined $conn);
	$BAD++ if(!defined $RESPONSE->{A} || ref($RESPONSE->{A}) ne 'ARRAY' ||
		$RESPONSE->{A}->[0] ne '1a' || $RESPONSE->{A}->[1] ne '2a');
	$BAD++ if(!defined $RESPONSE->{B} || ref($RESPONSE->{B}) ne 'ARRAY' ||
		$RESPONSE->{B}->[0] ne '1b' || $RESPONSE->{B}->[1] ne '2b');
	$BAD++ if(!defined $RESPONSE->{whisker});
	$BAD++ if(keys %$RESPONSE != 3);
	for($x=0;$x<~~@EXPECTED;$x++){
		$BAD++ if($RESPONSE->{whisker}->{header_order}->[$x] ne $EXPECTED[$x]);
	}
} else {
	$BAD++;
}
mok(!$BAD,'http_read_headers no.4');


$REQUEST->{whisker}->{ignore_duplicate_headers} = 1;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
@EXPECTED = ('A','B','A','B');
$BAD=0;
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
if($result==1){
	$BAD++ if(defined $enc);
	$BAD++ if(defined $len);
	$BAD++ if(defined $conn);
	$BAD++ if(!defined $RESPONSE->{A} || ref($RESPONSE->{A}) ||
		$RESPONSE->{A} ne '2a');
	$BAD++ if(!defined $RESPONSE->{B} || ref($RESPONSE->{B}) ||
		$RESPONSE->{B} ne '2b');
	$BAD++ if(!defined $RESPONSE->{whisker});
	$BAD++ if(keys %$RESPONSE != 3);
	for($x=0;$x<~~@EXPECTED;$x++){
		$BAD++ if($RESPONSE->{whisker}->{header_order}->[$x] ne $EXPECTED[$x]);
	}
} else {
	$BAD++;
}
mok(!$BAD,'http_read_headers no.5');

$INITIAL_DATA = <<EOT;
Set-Cookie: A
set-cookie2: B
set-Cookie: C

EOT

$REQUEST->{whisker}->{lowercase_incoming_headers} = 0;
$REQUEST->{whisker}->{normalize_incoming_headers} = 1;
$REQUEST->{whisker}->{ignore_duplicate_headers} = 1;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
@EXPECTED = ('A','B','C');
$BAD=0;
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
if($result==1){
	$BAD++ if(defined $enc);
	$BAD++ if(defined $len);
	$BAD++ if(defined $conn);
	$BAD++ if(!defined $RESPONSE->{'Set-Cookie2'} || $RESPONSE->{'Set-Cookie2'} ne 'B');
	$BAD++ if(!defined $RESPONSE->{'Set-Cookie'} || ref($RESPONSE->{'Set-Cookie'}) ||
		$RESPONSE->{'Set-Cookie'} ne 'C');
	$BAD++ if(!defined $RESPONSE->{whisker});
	$BAD++ if(keys %$RESPONSE != 3);
	$BAD++ if(!defined $RESPONSE->{whisker}->{cookies});
	for($x=0;$x<~~@EXPECTED;$x++){
		$BAD++ if($RESPONSE->{whisker}->{cookies}->[$x] ne $EXPECTED[$x]);
	}
} else {
	$BAD++;
}
mok(!$BAD,'http_read_headers no.6');


$INITIAL_DATA = <<EOT;
A:1
B: 2
C:  3
D:	4

EOT

$REQUEST->{whisker}->{lowercase_incoming_headers} = 0;
$REQUEST->{whisker}->{normalize_incoming_headers} = 0;
$REQUEST->{whisker}->{ignore_duplicate_headers} = 0;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
$BAD=0;
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
if($result==1){
	$BAD++ if(!defined $RESPONSE->{whisker}->{abnormal_header_spacing} ||
		$RESPONSE->{whisker}->{abnormal_header_spacing} != 3);
} else {
	$BAD++;
}
mok(!$BAD,'http_read_headers no.7');


$INITIAL_DATA = "A: 1\nB: 2\nC: 3\n\n";
$REQUEST->{whisker}->{lowercase_incoming_headers} = 0;
$REQUEST->{whisker}->{normalize_incoming_headers} = 0;
$REQUEST->{whisker}->{ignore_duplicate_headers} = 0;
$REQUEST->{whisker}->{save_raw_headers} = 1;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
mok($RESPONSE->{whisker}->{raw_header_data} eq $INITIAL_DATA,
	'http_read_headers no.8');

$INITIAL_DATA = "A: 1\nB: 2\nC: 3\n\nmore\nstuff";
$REQUEST->{whisker}->{save_raw_headers} = 0;
$RESPONSE = LW2::http_new_response();
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
($result,$enc,$len,$conn) = LW2::http_read_headers($STREAM,$REQUEST,$RESPONSE);
mok($STREAM->{bufin} eq "more\nstuff", 'http_read_headers no.9');

$STREAM->{close}->();

####################################################
# http_construct_headers

$REQUEST = LW2::http_new_request();
foreach(keys %$REQUEST){
	delete $REQUEST->{$_} if($_ ne 'whisker'); }
$EOL = $REQUEST->{whisker}->{http_eol};
$REQUEST->{A}='a';
$EXPECTED = 'A: a'.$EOL;
$res = LW2::http_construct_headers($REQUEST);
mok($res eq $EXPECTED, 'http_construct_headers no.1');

$REQUEST->{A} = ['a', 'b'];
$EXPECTED = 'A: a'.$EOL.'A: b'.$EOL;
$res = LW2::http_construct_headers($REQUEST);
mok($res eq $EXPECTED, 'http_construct_headers no.2');

$REQUEST->{A} = 'a';
$REQUEST->{B} = 'b';
$REQUEST->{C} = 'c';
$REQUEST->{whisker}->{header_order} = ['A','C','B'];
$EXPECTED = 'A: a'.$EOL.'C: c'.$EOL.'B: b'.$EOL;
$res = LW2::http_construct_headers($REQUEST);
mok($res eq $EXPECTED, 'http_construct_headers no.3');

$REQUEST->{whisker}->{header_order} = ['C','B'];
$EXPECTED = 'C: c'.$EOL.'B: b'.$EOL.'A: a'.$EOL;
$res = LW2::http_construct_headers($REQUEST);
mok($res eq $EXPECTED, 'http_construct_headers no.4');

$REQUEST->{A} = ['1','3'];
$REQUEST->{B} = ['2','4'];
delete $REQUEST->{C};
$REQUEST->{whisker}->{header_order} = ['A','B','A','B'];
$EXPECTED = 'A: 1'.$EOL.'B: 2'.$EOL.'A: 3'.$EOL.'B: 4'.$EOL;
$res = LW2::http_construct_headers($REQUEST);
mok($res eq $EXPECTED, 'http_construct_headers no.5');

$REQUEST->{whisker}->{header_order} = ['A','B','A'];
$EXPECTED = 'A: 1'.$EOL.'B: 2'.$EOL.'A: 3'.$EOL.'B: 4'.$EOL;
$res = LW2::http_construct_headers($REQUEST);
mok($res eq $EXPECTED, 'http_construct_headers no.6');

####################################################################

$INITIAL_DATA = "Aaa\r\nBbb\r\n";
$STREAM = buffer_stream($INITIAL_DATA);

$res[0] = LW2::_http_getline($STREAM);
$res[1] = LW2::_http_getline($STREAM);
$res[2] = LW2::_http_getline($STREAM);

$BAD=0;
$BAD++ if(!defined $res[0] || $res[0] ne "Aaa\r\n" ||
	!defined $res[1] || $res[1] ne "Bbb\r\n" ||
	defined $res[2]);
mok(!$BAD, '_http_getline no.1');

$INITIAL_DATA = "Aaa";
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
$res[0] = LW2::_http_getline($STREAM);
mok(!defined $res[0], '_http_getline no.2');

$INITIAL_DATA = "Aaa\r";
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
$res[0] = LW2::_http_getline($STREAM);
mok(!defined $res[0], '_http_getline no.3');

$INITIAL_DATA = "Aaa\n";
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
$res[0] = LW2::_http_getline($STREAM);
mok(defined $res[0] && $res[0] eq "Aaa\n", '_http_getline no.4');

####################################################################

$INITIAL_DATA = "ABC";
$STREAM = buffer_stream($INITIAL_DATA);

$res[0] = LW2::_http_get($STREAM, 1);
$res[1] = LW2::_http_get($STREAM, 1);
$res[2] = LW2::_http_get($STREAM, 1);
$res[3] = LW2::_http_get($STREAM, 1);

$BAD=0;
$BAD++ if(!defined $res[0] || $res[0] ne "A" ||
	!defined $res[1] || $res[1] ne "B" ||
	!defined $res[2] || $res[2] ne "C" ||
	defined $res[3]);
mok(!$BAD, '_http_get no.1');

$INITIAL_DATA = "A";
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
$res = LW2::_http_get($STREAM, 2);
mok(!defined $res, '_http_get no.2');

####################################################################

$INITIAL_DATA = "ABCDE";
$STREAM = buffer_stream($INITIAL_DATA);

$res = LW2::_http_getall($STREAM, 3);
mok($res eq 'ABC', '_http_getall no.1');

$STREAM->{clearall}->();
$STREAM->{bufin} = 'AB';
$STREAM->{bufout} = 'CD';
$res = LW2::_http_getall($STREAM,4);
mok($res eq 'ABCD', '_http_getall no.2');

$BAD=0;
$INITIAL_DATA = 'ABCDE';
$STREAM->{clearall}->();
$STREAM->{write}->($INITIAL_DATA);
$res = LW2::_http_getall($STREAM, 10);
mok($res eq 'ABCDE', '_http_getall no.3');

####################################################################

$INITIAL_DATA = "Response";

# first we test the getall route
$REQUEST = LW2::http_new_request();
$RESPONSE = LW2::http_new_response();

$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST->{whisker}->{max_size} = 1024;
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'Response');
mok(!$BAD, 'http_read_body no.1');

$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST->{whisker}->{max_size}=8;
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'Response');
mok(!$BAD, 'http_read_body no.2');

$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST->{whisker}->{max_size}=3;
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'Res');
mok(!$BAD, 'http_read_body no.3');

$STREAM = buffer_stream($INITIAL_DATA);
delete $REQUEST->{whisker}->{max_size};
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'Response');
mok(!$BAD, 'http_read_body no.4');

# next the length-defined route
$REQUEST = LW2::http_new_request();

$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, 1024);
# read_body should return error/0, since we don't have enough data
mok($BAD, 'http_read_body no.5');

$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, 4);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'Resp');
mok(!$BAD, 'http_read_body no.6');

$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, 0);
$BAD++ if($RESPONSE->{whisker}->{data} ne '');
mok(!$BAD, 'http_read_body no.7');

$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST->{whisker}->{max_size}=3;
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, 8);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'Res');
mok(!$BAD, 'http_read_body no.8');

$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST->{whisker}->{max_size}=3;
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, undef, 1024);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'Res');
mok(!$BAD, 'http_read_body no.9');

# now for chunked encoding
$INITIAL_DATA = "3\r\nABC\r\n2\r\nDE\r\n0\r\n";
$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST = LW2::http_new_request();
$RESPONSE = LW2::http_new_response();
# defaults to save_raw_chunks=0
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'ABCDE');
mok(!$BAD, 'http_read_body no.10');

$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST->{whisker}->{max_size} = 4;
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'ABCD');
mok(!$BAD, 'http_read_body no.11');

$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST = LW2::http_new_request();
$REQUEST->{whisker}->{save_raw_chunks} = 1;
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne $INITIAL_DATA);
mok(!$BAD, 'http_read_body no.12');

$INITIAL_DATA = "1\r\nA\r\n00000001\r\nB\r\n1\r\nC\r\n0\r\n";
$STREAM = buffer_stream($INITIAL_DATA);
$REQUEST = LW2::http_new_request();
$RESPONSE = LW2::http_new_response();
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'ABC');
mok(!$BAD, 'http_read_body no.13');

$INITIAL_DATA = "1\nA\n1\nB\n1\nC\n0\n";
$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'ABC');
mok(!$BAD, 'http_read_body no.14');

$INITIAL_DATA = "1\nA\nA\nBCDEFGHIJK\n0\n";
$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'ABCDEFGHIJK');
mok(!$BAD, 'http_read_body no.15');

$INITIAL_DATA = "1\nA\nFFFFFFFFF\nAAA";
$STREAM = buffer_stream($INITIAL_DATA);
$BAD = LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
mok(!$BAD, 'http_read_body no.16');

$INITIAL_DATA = "5\nABC";
$RESPONSE = LW2::http_new_response();
$STREAM = buffer_stream($INITIAL_DATA);
$BAD = LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
mok(!$BAD, 'http_read_body no.17');

$INITIAL_DATA = "-3\nABC\n0\n";
$RESPONSE = LW2::http_new_response();
$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'ABC');
mok(!$BAD, 'http_read_body no.18');

$INITIAL_DATA = "2\nAB\n2\nCD\n0\nA: a\nB: b\n";
$RESPONSE = LW2::http_new_response();
$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne 'ABCD');
mok(!$BAD, 'http_read_body no.19');

$INITIAL_DATA = "2\nAB\n2\nCD\n0\nA: a\nB: b\n";
$REQUEST->{whisker}->{save_raw_chunks}=1;
$STREAM = buffer_stream($INITIAL_DATA);
$BAD = !LW2::http_read_body($STREAM, $REQUEST, $RESPONSE, 'chunked', undef);
$BAD++ if($RESPONSE->{whisker}->{data} ne $INITIAL_DATA);
mok(!$BAD, 'http_read_body no.20');

######################################################################
# start test server

$TS_PID = undef;
start_test_server();
if(!defined $TS_PID || $TS_PID==0){ # massive problem
	die("problem starting test server"); }

# setup so we kill the test server when we're done
END {
	kill 9, $TS_PID;
}

#####################################################################

LW2::http_reset();

$REQUEST = LW2::http_new_request();
$RESPONSE = LW2::http_new_response();
LW2::uri_split('http://localhost:8088/', $REQUEST);
$REQUEST->{whisker}->{normalize_incoming_headers}=0;
$REQUEST->{whisker}->{ignore_duplicate_headers}=0;
$REQUEST->{whisker}->{retry}=0;

$REQUEST->{whisker}->{uri} = '/TESTCASE/basic1/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{socket_state} != 0);
$BAD++ if($RESPONSE->{whisker}->{protocol} ne 'HTTP');
$BAD++ if($RESPONSE->{whisker}->{version} ne '1.0');
$BAD++ if($RESPONSE->{whisker}->{http_space1} ne ' ');
$BAD++ if($RESPONSE->{whisker}->{http_space2} ne ' ');
$BAD++ if($RESPONSE->{whisker}->{code} ne 200);
$BAD++ if($RESPONSE->{whisker}->{message} ne 'OK');
$BAD++ if($RESPONSE->{whisker}->{http_eol} ne "\x0a");
mok(!$BAD, 'live http testcase basic1');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/basic2/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{socket_state} != 0);
mok(!$BAD, 'live http testcase basic2');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/basic3/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{socket_state} != 1);
mok(!$BAD, 'live http testcase basic3');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/basic4/';
$REQUEST->{whisker}->{force_open} = 1;
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{socket_state} != 1);
mok(!$BAD, 'live http testcase basic4 force-open');
LW2::http_reset();
$REQUEST->{whisker}->{force_open} = 0;

$REQUEST->{whisker}->{force_close} = 0;
$REQUEST->{whisker}->{uri} = '/TESTCASE/basic3/';
$REQUEST->{whisker}->{force_close} = 1;
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{socket_state} != 0);
mok(!$BAD, 'live http testcase basic3 force-close');
LW2::http_reset();
$REQUEST->{whisker}->{force_close} = 0;

# lots of consecutive closed requests
$BAD=0;
$REQUEST->{whisker}->{uri} = '/TESTCASE/basic1/';
for(1..15){
	$BAD += validate_test_request($REQUEST, $RESPONSE);
	$BAD++ if($RESPONSE->{whisker}->{socket_state} != 0);
}
mok(!$BAD, 'live http testcase basic1 repetitive');
LW2::http_reset();

# lots of consecutive keep-alive requests
$BAD=0;
$REQUEST->{whisker}->{uri} = '/TESTCASE/basic3/';
for(1..15){
	$BAD += validate_test_request($REQUEST, $RESPONSE);
	$BAD++ if($RESPONSE->{whisker}->{socket_state} != 1);
	$BAD++ if($RESPONSE->{whisker}->{stats_syns} != 1);
}
mok(!$BAD, 'live http testcase basic3 repetitive');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/100cont1/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{'100_continue'} != 1);
mok(!$BAD, 'live http testcase 100cont1');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/100cont2/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{'100_continue'} != 3);
mok(!$BAD, 'live http testcase 100cont2');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/100cont3/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{'100_continue'} != 3);
mok(!$BAD, 'live http testcase 100cont3');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/spacer/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$BAD++ if($RESPONSE->{whisker}->{'http_space1'} ne "\t");
$BAD++ if($RESPONSE->{whisker}->{'http_space2'} ne "  ");
mok(!$BAD, 'live http testcase spacer');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/chunk1/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
mok(!$BAD, 'live http testcase chunk1');
LW2::http_reset();

$REQUEST->{whisker}->{uri} = '/TESTCASE/chunk1raw/';
$REQUEST->{whisker}->{save_raw_chunks} = 1;
$BAD = validate_test_request($REQUEST, $RESPONSE);
mok(!$BAD, 'live http testcase chunk1 raw');
LW2::http_reset();


######################################################################

# retries testing

# the retry test case will give us a keep-alive answer, and pretend
# to keep the connection alive but will really close it, which will
# cause an error when we try to attempt the next request

# first we do a request with retries disabled, to prove we get
# an error
$REQUEST->{whisker}->{retry} = 0;
$REQUEST->{whisker}->{uri} = '/TESTCASE/retry/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$REQUEST->{whisker}->{uri} = '/TESTCASE/basic4/';
$BAD++ if(! LW2::http_do_request($REQUEST, $RESPONSE));
mok(!$BAD, 'live http testcase retries no.1');
LW2::http_close($REQUEST);
LW2::http_reset();

# next we do the same thing with retries, which should succeed
$REQUEST->{whisker}->{retry} = 1;
$REQUEST->{whisker}->{uri} = '/TESTCASE/retry/';
$BAD = validate_test_request($REQUEST, $RESPONSE);
$REQUEST->{whisker}->{uri} = '/TESTCASE/basic4/';
$BAD += validate_test_request($REQUEST, $RESPONSE);
mok(!$BAD, 'live http testcase retries no.2');
LW2::http_close($REQUEST);
LW2::http_reset();


# HTTP stuff to test:
# - save_raw_headers -> raw_header_data
# - data_sock -> data_sock, data_stream
# - force_bodysnatch
# - hide_chunked_responses
# - save_raw_chunks
# - header_delete_on_success


######################################################################
######################################################################
# crawl stuff

$REQUEST = LW2::http_new_request();
$RESPONSE = LW2::http_new_response();
LW2::uri_split('http://localhost:8088/', $REQUEST);
$REQUEST->{whisker}->{normalize_incoming_headers}=0;
$REQUEST->{whisker}->{ignore_duplicate_headers}=0;
$REQUEST->{whisker}->{retry}=0;

$REQUEST->{whisker}->{uri} = '/TESTCASE/crawllinks/';
$BAD=0;
if(!LW2::http_do_request($REQUEST, $RESPONSE)){
	if($RESPONSE->{whisker}->{code} == 200){
		@urls = LW2::html_link_extractor(\$RESPONSE->{whisker}->{data});
		@urls = sort @urls;
		if(~~@urls != 39){
			$BAD++;
		} else {
			$x = 0;
			while($_ = shift @urls){
				$y = sprintf('%02d', $x);
				$BAD++ if($_ ne $y . '.html');
				$x++;
			}
		}
	} else {
		$BAD++;
	}
} else {
	$BAD++;
}
mok(!$BAD, 'html_link_extractor');
LW2::http_reset();

# reset crawl results
$SKIP_CRAWL=0;
$REQUEST->{whisker}->{uri} = '/CRAWLERRESET/';
$SKIP_CRAWL++ if(LW2::http_do_request($REQUEST,$RESPONSE));
if(!$SKIP_CRAWL){
	$SKIP_CRAWL++ if($RESPONSE->{whisker}->{code} != 200);
}

%TRACK = ();

if(!$SKIP_CRAWL){

	# basic crawl test
	$CRAWLER = LW2::crawl_new('http://localhost:8088/CRAWLERSTART/',
		35, $REQUEST, \%TRACK);
	$res = LW2::crawl($CRAWLER);

	# first we see what URLs we requested, according to the server
	$REQUEST->{whisker}->{uri} = '/CRAWLERRESULT/';
	$BAD = 0;
	if(!LW2::http_do_request($REQUEST,$RESPONSE)){
		$BAD++ if(LW2::md5($RESPONSE->{whisker}->{data}) ne
			'd8545f61b95130b94626ef24ba64a11a');
	} else {
		$BAD++;
	}
	# then we ensure we actually got the right number of URLs
	$BAD++ if($res != 16);
	mok(!$BAD,'crawler no.1');

	$REQUEST->{whisker}->{uri} = '/CRAWLERRESET/';
	LW2::http_do_request($REQUEST,$RESPONSE);

	# basic crawl test w/ HEAD requests
	$CRAWLER = LW2::crawl_new('http://localhost:8088/CRAWLERSTART/',
		35, $REQUEST, \%TRACK);
	$CRAWLER->{config}->{do_head} = 1;
	$res = LW2::crawl($CRAWLER);

	# first we see what URLs we requested, according to the server
	$REQUEST->{whisker}->{uri} = '/CRAWLERRESULT/';
	$BAD = 0;
	if(!LW2::http_do_request($REQUEST,$RESPONSE)){
		$BAD++ if(LW2::md5($RESPONSE->{whisker}->{data}) ne
			'5f57f08b9b194eb1ebd602acb0eb8da4');
	} else {
		$BAD++;
	}
	# then we ensure we actually got the right number of URLs
	$BAD++ if($res != 16);
	mok(!$BAD,'crawler no.2');

	# TODO: more crawl testing

} else {
	# we need to skip the tests so Test[::Simple] don't complain
	# note: keep the number below in sync with the number of tests
	# above
	for(1..1){
		mok(0, 'crawl placeholder for skipped tests');
	}
}


######################################################################

$REQUEST = LW2::http_new_request();
LW2::auth_set('basic', $REQUEST, 'Aladdin', 'open sesame', '');
$BAD=0;
$BAD++ if(!defined $REQUEST->{Authorization} ||
	$REQUEST->{Authorization} ne 'Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==');
mok(!$BAD,'auth_set no.1');
delete $REQUEST->{Authorization};

LW2::auth_set('proxy-basic', $REQUEST, 'Aladdin', 'open sesame', '');
$BAD=0;
$BAD++ if(!defined $REQUEST->{'Proxy-Authorization'} ||
	$REQUEST->{'Proxy-Authorization'} ne 'Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==');
mok(!$BAD,'auth_set no.2');
delete $REQUEST->{Authorization};

# TODO: auth_set('ntlm') && auth_set('proxy_ntlm')

$REQUEST->{Authorization} = 'foo';
$REQUEST->{'Proxy-Authorization'} = 'bar';
$REQUEST->{whisker}->{auth_callback} = '1';
$REQUEST->{whisker}->{auth_proxy_callback} = '1';
$REQUEST->{whisker}->{auth_data} = '1';
$REQUEST->{whisker}->{auth_proxy_data} = '1';
LW2::auth_unset($REQUEST);
$BAD=0;
$BAD++ if(defined $REQUEST->{Authorization});
$BAD++ if(defined $REQUEST->{'Proxy-Authorization'});
$BAD++ if(defined $REQUEST->{whisker}->{auth_callback});
$BAD++ if(defined $REQUEST->{whisker}->{auth_proxy_callback});
$BAD++ if(defined $REQUEST->{whisker}->{auth_data});
$BAD++ if(defined $REQUEST->{whisker}->{auth_proxy_data});
mok(!$BAD,'auth_unset');


$REQUEST = LW2::http_new_request();
$RESPONSE = LW2::http_new_response();
LW2::uri_split('http://localhost:8088/', $REQUEST);
$REQUEST->{whisker}->{normalize_incoming_headers}=0;
$REQUEST->{whisker}->{ignore_duplicate_headers}=0;
$REQUEST->{whisker}->{retry}=0;
$REQUEST->{whisker}->{uri} = '/AUTH/Basic+QWxhZGRpbjpvcGVuIHNlc2FtZQ==/';

@passwords = qw(A B C D AAA BBB CCC Word One Two open Sessame foo bar baz);
push @passwords, 'open sesame', 'more';
$res = LW2::auth_brute_force('basic', $REQUEST, 'Aladdin', \@passwords);
$BAD=0;
$BAD++ if(!defined $res || $res ne 'open sesame');
mok(!$BAD,'auth_brute_force');



######################################################################
######################################################################
# TODO:

# --auth
# auth_brute_force
#
# --cookie
# cookie_read
# cookie_write
# cookie_get_valid_names
#
# --dump
# dump
# dump_writefile
#
# --forms
# forms_read
# forms_write
#
# --html
# html_find_tags_rewrite
#
# --http
# http_do_request_timeout
#
# --multipart
# multipart_set
# multipart_get
# multipart_setfile
# multipart_getfile
# multipart_boundary
# multipart_write
# multipart_read
# multipart_read_data
# multipart_files_list
# multipart_params_list
#
# --ntlm
# ntlm_new
# ntlm_generate_responses
# ntlm_decode_challenge
# ntlm_header
# ntlm_client
#
# --simple
# get_page
# get_page_hash
# get_page_to_file
#
# --streams
# -...
#
# --time
# time_mktime
# time_gmtolocal
#
# --utils
# utils_recperm
# utils_port_open
# utils_save_page
# utils_bruteurl
# utils_request_fingerprint
# utils_flatten_lwhash
# utils_carp
# utils_croak

#########################################################################
#########################################################################

exit;

#####################

sub start_test_server {

	die("Couldn't locate testserver.pl!") if(!-e 'testserver.pl');

	$TS_PID = fork();
	die("Unable to fork for test server") if(!defined $TS_PID);

	$SIG{CHLD} = sub { 
		print STDERR "#! Test server unexpected exit!\n"; 
		exit; 
	    };

	if($TS_PID == 0){
		# child
		exec($^X,'testserver.pl','.');
		exit;
	}

	# parent
	select(undef,undef,undef,.5); # give time to startup
}

END {
	$SIG{CHLD} = 'IGNORE';
}

sub validate_test_request {
	my ($req, $resp) = @_;
	return 1 if( LW2::http_do_request($req,$resp) );
	return 2 if( $resp->{whisker}->{code} != 200 );
	my $h;
	foreach(@{$resp->{whisker}->{header_order}}){
		$h = $_ if(lc($_) eq 'md5sum');
	}
	return 3 if(!defined $h);
	my ($tmd5) = LW2::utils_find_key($resp, $h);
	$tmd5=~tr/a-f0-9//cd;
	my $md5 = LW2::md5( $resp->{whisker}->{data} );
	return 4 if($tmd5 ne lc($md5));
	return 0;
}


sub buffer_stream {
	my $initial_data = shift;
	my $REQUEST = LW2::http_new_request();
	$REQUEST->{whisker}->{buffer_stream}++;
	my $STREAM = LW2::stream_new($REQUEST);
	die("Buffer stream problem") if !defined $STREAM;
	$STREAM->{open}->();
	if(defined $initial_data){
		$STREAM->{write}->($initial_data);
	}
	return $STREAM;
}
