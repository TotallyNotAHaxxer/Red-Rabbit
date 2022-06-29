#
# This is a compatiblity 'bridge' which will translate the
# libwhisker 2.x API into libwhisker 1.x format.  This should
# only be used to support legacy programs which refuse to port
# to LW2, but should be using LW2 over LW[1] because of bug fixes.
#

package LW;

require 'LW2.pm';

$LW::VERSION = '1.10';
$LW::BRIDGE = '2.0';

#
# NOTE: The following two lines depend on external files; remove/comment
#       out if you need single-file portability
#
use strict;
use vars qw(%available $LW_HAS_SOCKET $LW_HAS_SSL $TIMEOUT
	$LW_SSL_LIB $LW_NONBLOCK_CONNECT $FUNC %_deprec
	%crawl_server_tags %crawl_referrers %_remap_to
	%crawl_offsites %crawl_cookies %crawl_forms %_remap_from
	%crawl_linktags %crawl_config
	);

#### GLOBAL VARIABLE STUFF ####

%available		= ();
$LW_HAS_SOCKET		= (defined $Socket::VERSION)?1:0;
$LW_HAS_SSL		= ($LW2::LW_SSL_LIB>0)?1:0;
$LW_SSL_LIB		= $LW2::LW_SSL_LIB;
$LW_NONBLOCK_CONNECT	= $LW2::LW_NONBLOCK_CONNECT;

%crawl_server_tags=();
%crawl_referrers=();
%crawl_offsites=();
%crawl_cookies=();
%crawl_forms=();
%crawl_linktags = %LW2::_crawl_linktags;
%crawl_config = %LW2::_crawl_config;

$TIMEOUT=10;	# doesn't do anything

#### BRIDGED FUNCTIONS ####

# antiids.pl = DONE
sub anti_ids { 
	warn("Anti-IDS: session splicing is not supported")
		if($_[1]=~/9/);
	my $hr = $_[0];
	_remap($hr);
	LW2::encode_anti_ids(@_);
	_remap_from($hr);
}

# auth.pl = DONE
sub auth_set_header		{ goto &LW2::auth_set; }
sub do_auth			{ goto &LW2::auth_set; }
sub auth_brute_force {
	_remap($_[1]);
	goto &LW2::auth_brute_force;
}

# bruteurl.pl = DONE
sub bruteurl { 
	_remap($_[0]);
	goto &LW2::utils_bruteurl;
}

# cookie.pl = DONE
sub cookie_read			{ goto &LW2::cookie_read; }
sub cookie_write		{ goto &LW2::cookie_write; }
sub cookie_parse		{ goto &LW2::cookie_parse; }
sub cookie_get			{ goto &LW2::cookie_get; }
sub cookie_set			{ goto &LW2::cookie_set; }

# crawl.pl = DONE
sub crawl_get_config {
	my $key=shift;
	return $crawl_config{$key};
}
sub crawl_set_config {
	return if(!defined $_[0]);
	my %opts=@_;
	while( my($k,$v)=each %opts){
		$crawl_config{lc($k)}=$v; }
}
sub crawl { # crawl changed *a lot*, so we have a lot of fixing to do...
	my ($START, $DEPTH, $TRACK, $HREQ)=@_;
	_remap($HREQ);
	my $CRAWL = LW2::crawl_new($START,$DEPTH,$HREQ,$TRACK);
	$crawl_config{ref_hin}=$CRAWL->{request};
	$crawl_config{ref_hout}=$CRAWL->{response};
	$crawl_config{ref_jar}=$CRAWL->{jar};
	$crawl_config{ref_links}=$CRAWL->{urls};
	my @p = LW2::uri_split($START);
	$crawl_config{host}=$p[2];
	$crawl_config{port}=$p[3];
	$crawl_config{start}=$p[0];
	%{$CRAWL->{config}} = %crawl_config;
	%crawl_server_tags=();		%crawl_referrers=();
	%crawl_offsites=();		%crawl_cookies=();
	%crawl_forms=();
	my $res = $CRAWL->{crawl}->();
	return if(!defined $res);
	%crawl_server_tags = %{$CRAWL->{server_tags}};
	%crawl_referrers = %{$CRAWL->{referrers}};
	%crawl_offsites = %{$CRAWL->{offsites}};
	%crawl_cookies = %{$CRAWL->{cookies}};
	%crawl_forms = %{$CRAWL->{forms}};
}


# dump.pl = DONE
sub dumper { 
	my $res = &LW2::dump(@_);
	$res = 'ERROR' if(!defined $res);
	return $res;
}
sub dumper_writefile		{ goto &LW2::dump_writefile; }

# easy.pl = DONE
sub upload_file { die("<upload_file is not implemented>"); }
sub get_page { 
	_remap($_[1]);
	goto &LW2::get_page;
}
sub get_page_hash {
	_remap($_[1]);
	my $res = LW2::get_page_hash(@_);
	return _remap_from($res);
}
sub get_page_to_file {
	_remap($_[2]);
	goto &LW2::get_page_to_file;
}
sub download_file {
	_remap($_[2]);
	goto &LW2::get_page_to_file;
}

# encode.pl = DONE
sub encode_base64		{ goto &LW2::encode_base64; }
sub encode_base64_perl		{ goto &LW2::encode_base64; }
sub decode_base64		{ goto &LW2::decode_base64; }
sub decode_base64_perl		{ goto &LW2::decode_base64; }
sub encode_str2uri		{ goto &LW2::encode_uri_hex; }
sub encode_str2ruri		{ goto &LW2::encode_uri_randomhex; }
sub encode_unicode		{ goto &LW2::encode_unicode; }

# forms.pl = DONE
sub forms_read { 
	warn("LW1.x forms support was broken; LW2 is fixed, but not compatible");
	goto &LW2::forms_read; 
}
sub forms_write {
	warn("LW1.x forms support was broken; LW2 is fixed, but not compatible");
	goto &LW2::forms_write;
}

# html.pl = DONE
{ $FUNC = '';
sub html_find_tags {
	my ($dr,$func)=@_;	
	$FUNC = $func;
	LW2::html_find_tags($dr,\&_html_callback_wrapper);
	$FUNC = '';
}

sub _html_callback_wrapper {
	return if($FUNC eq '');
	my $res = &$FUNC(@_);
	LW2::_html_find_tags_adjust($res,0) if(defined $res && $res > 0);
}}

# http.pl = DONE
sub http_reset { goto &LW2::http_reset; }
sub http_init_request {
	my $href = shift;
	LW2::http_init_request($href);
	$href->{whisker}->{version}='1.0'; # default for LW1.x
	_remap_from($href);
	$href->{Connection}='close'; # default for LW1.x
}
sub http_do_request {
	my ($req,$resp,%conf)=@_;
	my ($k,$v);
	while(($k,$v)=each(%conf)){
		$req->{whisker}->{$k}=$v; }
	_remap($req);
	my $res = LW2::http_do_request($req,$resp);
	_remap_from($resp);
	return $res;
}
sub http_fixup_request {
	my $req=shift;
	_remap($req);
	LW2::http_fixup_request($req);
	_remap_from($req);	
}

# mdx.pl = DONE
sub md5				{ goto &LW2::md5; }
sub md5_perl			{ goto &LW2::md5; }
sub md4				{ goto &LW2::md4; }
sub md4_perl			{ goto &LW2::md4; }

# multipart.pl = DONE
sub multipart_set		{ goto &LW2::mutipart_set; }
sub multipart_get		{ goto &LW2::mutipart_get; }
sub multipart_setfile		{ goto &LW2::mutipart_setfile; }
sub multipart_getfile		{ goto &LW2::mutipart_getfile; }
sub multipart_boundary		{ goto &LW2::mutipart_boundary; }
sub multipart_write		{ goto &LW2::mutipart_write; }
sub multipart_read		{ goto &LW2::mutipart_read; }
sub multipart_read_data		{ goto &LW2::mutipart_read_data; }
sub multipart_files_list	{ goto &LW2::mutipart_files_list; }
sub multipart_params_list	{ goto &LW2::mutipart_params_list; }

# ntlm.pl = DONE
sub ntlm_new			{ goto &LW2::ntlm_new; }
sub ntlm_client			{ goto &LW2::ntlm_client; }

# utils.pl = DONE
sub utils_recperm		{ goto &LW2::utils_recperm; }
sub utils_array_shuffle		{ goto &LW2::utils_array_shuffle; }
sub utils_randstr		{ goto &LW2::utils_randstr; }
sub utils_get_dir		{ goto &LW2::uri_get_dir; }
sub utils_port_open		{ goto &LW2::utils_port_open; }
sub utils_getline		{ goto &LW2::utils_getline; }
sub utils_getline_crlf		{ goto &LW2::utils_getline_crlf; }
sub utils_absolute_uri		{ goto &LW2::uri_absolute; }
sub utils_normalize_uri		{ goto &LW2::uri_normalize; }
sub utils_save_page		{ goto &LW2::utils_save_page; }
sub utils_getopts		{ goto &LW2::utils_getopts; }
sub utils_unidecode_uri		{ goto &LW2::decode_unicode; }
sub utils_text_wrapper		{ goto &LW2::utils_text_wrapper; }
sub utils_lowercase_headers	{ goto &LW2::utils_lowercase_keys; }
sub utils_lowercase_hashkeys	{ goto &LW2::utils_lowercase_keys; }
sub utils_find_lowercase_key	{ goto &LW2::utils_find_lowercase_key; }
sub utils_join_uri		{ goto &LW2::uri_join; }
sub utils_split_uri {
	my $hr = $_[1];
	my @res = &LW2::uri_split(@_);
	_remap_from($hr);
	return @res;
}

#### COMPATIBILITY SUPPORT FUNCTIONS ####

%_remap_to = (
	'req_spacer'	=> 'http_space1',
	'req_spacer2'	=> 'http_space2',
	'http_ver'	=> 'version',
	'http_protocol'	=> 'protocol',
	'uri_param'	=> 'parameters',
	'sockstate'	=> 'socket_state',
	'recv_header_order' => 'header_order',
	'http_resp_message' => 'message'
);

%_remap_from = ();
while(my($k,$v)=each(%_remap_to)){
	$_remap_from{$v}=$k; }

%_deprec = (
	'method_postfix' => 1,
	'http_req_trailer' => 1,
	'queue_md5' => 1,
	'retry_errors' => 1,
	'ids_session_splice' => 1
);

sub _remap_from {
	_remap($_[0],1);
}

sub _remap {
	my $hr = shift;
	return undef if(!defined $hr || !ref($hr));
	my $from = shift||0;
	my $MAP = \%_remap_to;
	$MAP = \%_remap_from 
		if($from || $hr->{whisker}->{MAGIC} eq '31340');
	my @k = keys %{ $hr->{whisker} };
	foreach(@k){
		$hr->{whisker}->{http_resp} = $hr->{whisker}->{code}
			if($_ eq 'code');
		warn("whisker option '$_' will be ignored")
			if(exists $_deprec{$_});
		next if(!defined $MAP->{$_});
		$hr->{whisker}->{ $MAP->{$_} } = $hr->{whisker}->{$_};
	}
}
1;

