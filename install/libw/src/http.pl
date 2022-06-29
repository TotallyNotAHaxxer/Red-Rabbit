#BSD   Copyright (c) 2009, Jeff Forristal (wiretrip.net)
#BSD   All rights reserved.
#BSD
#BSD   Redistribution and use in source and binary forms, with or without 
#BSD   modification, are permitted provided that the following conditions 
#BSD   are met:
#BSD
#BSD   - Redistributions of source code must retain the above copyright 
#BSD   notice, this list of conditions and the following disclaimer.
#BSD
#BSD   - Redistributions in binary form must reproduce the above copyright 
#BSD   notice, this list of conditions and the following disclaimer in the 
#BSD   documentation and/or other materials provided with the distribution.
#BSD
#BSD   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS 
#BSD   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT 
#BSD   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS 
#BSD   FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE 
#BSD   COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, 
#BSD   INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, 
#BSD   BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; 
#BSD   LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER 
#BSD   CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT 
#BSD   LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN 
#BSD   ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE 
#BSD   POSSIBILITY OF SUCH DAMAGE.

##################################################################

# cluster global variables
%http_host_cache = ();

##################################################################

=item B<http_new_request>

Params: %parameters

Return: \%request_hash

This function basically 'objectifies' the creation of whisker
request hash objects.  You would call it like:

 $req = http_new_request( host=>'www.example.com', uri=>'/' )

where 'host' and 'uri' can be any number of {whisker} hash
control values (see http_init_request for default list).

=cut

sub http_new_request {
    my %X = @_;
    my ( $k, $v, %RET, %RES );

    http_init_request( \%RET );
    while ( ( $k, $v ) = each(%X) ) {
        $RET{whisker}->{$k} = $v;
    }
    $RES{whisker}          = {};
    $RES{whisker}->{MAGIC} = 31340;
    $RES{whisker}->{uri}   = '';
    return ( \%RET, \%RES ) if wantarray();
    return \%RET;
}

##################################################################

=item B<http_new_response>

Params: [none]

Return: \%response_hash

This function basically 'objectifies' the creation of whisker
response hash objects.  You would call it like:

	$resp = http_new_response()

=cut

sub http_new_response {
    my %RET;
    $RET{whisker}          = {};
    $RET{whisker}->{MAGIC} = 31340;
    $RET{whisker}->{uri}   = '';
    return \%RET;
}

##################################################################

=item B<http_init_request>

Params: \%request_hash_to_initialize

Return: Nothing (modifies input hash)

Sets default values to the input hash for use.  Sets the host to
'localhost', port 80, request URI '/', using HTTP 1.1 with GET
method.  The timeout is set to 10 seconds, no proxies are defined, and all
URI formatting is set to standard HTTP syntax.  It also sets the
Connection (Keep-Alive) and User-Agent headers.

NOTICE!!  It's important to use http_init_request before calling 
http_do_request, or http_do_request might puke.  Thus, a special magic 
value is placed in the hash to let http_do_request know that the hash has 
been properly initialized.  If you really must 'roll your own' and not use 
http_init_request before you call http_do_request, you will at least need 
to set the MAGIC value (amongst other things).

=cut

sub http_init_request {    # doesn't return anything
    my ($hin) = shift;

    return if ( !( defined $hin && ref($hin) ) );
    %$hin = ();            # clear control hash

    # control values
    $$hin{whisker} = {
        http_space1                   => ' ',
        http_space2                   => ' ',
        version                       => '1.1',
        method                        => 'GET',
        protocol                      => 'HTTP',
        port                          => 80,
        uri                           => '/',
        uri_prefix                    => '',
        uri_postfix                   => '',
        uri_param_sep                 => '?',
        host                          => 'localhost',
        timeout                       => 10,
        include_host_in_uri           => 0,
        ignore_duplicate_headers      => 1,
        normalize_incoming_headers    => 1,
        lowercase_incoming_headers    => 0,
        require_newline_after_headers => 0,
        invalid_protocol_return_value => 1,
        ssl                           => 0,
        ssl_save_info                 => 0,
        http_eol                      => "\x0d\x0a",
        force_close                   => 0,
        force_open                    => 0,
        retry                         => 1,
        trailing_slurp                => 0,
        force_bodysnatch              => 0,
        max_size                      => 0,
        MAGIC                         => 31339
    };

    # default header values
    $$hin{'Connection'} = 'Keep-Alive';
    $$hin{'User-Agent'} = "Mozilla (libwhisker/$LW2::VERSION)";
}

##################################################################

=item B<http_do_request>

Params: \%request, \%response [, \%configs]

Return: >=1 if error; 0 if no error (also modifies response hash)

*THE* core function of libwhisker.  http_do_request actually performs
the HTTP request, using the values submitted in %request, and placing result
values in %response.  This allows you to resubmit %request in subsequent 
requests (%response is automatically cleared upon execution).  You can 
submit 'runtime' config directives as %configs, which will be spliced into
$hin{whisker}->{} before anything else.  That means you can do:

LW2::http_do_request(\%req,\%resp,{'uri'=>'/cgi-bin/'});

This will set $req{whisker}->{'uri'}='/cgi-bin/' before execution, and
provides a simple shortcut (note: it does modify %req).

This function will also retry any requests that bomb out during the 
transaction (but not during the connecting phase).  This is controlled
by the {whisker}->{retry} value.  Also note that the returned error
message in hout is the *last* error received.  All retry errors are
put into {whisker}->{retry_errors}, which is an anonymous array.

Also note that all NTLM auth logic is implemented in http_do_request().
NTLM requires multiple requests in order to work correctly, and so this
function attempts to wrap that and make it all transparent, so that the
final end result is what's passed to the application.

This function will return 0 on success, 1 on HTTP protocol error, and 2
on non-recoverable network connection error (you can retry error 1, but
error 2 means that the server is totally unreachable and there's no
point in retrying).

=cut

sub http_do_request {
    my ( $hin, $hout ) = ( shift, shift );

    return 2 if ( !( defined $hin  && ref($hin) ) );
    return 2 if ( !( defined $hout && ref($hout) ) );

    # setup hash
    %$hout                     = ();
    $$hout{whisker}            = {};
    $$hout{whisker}->{'MAGIC'} = 31340;
    $$hout{whisker}->{uri}     = $$hin{whisker}->{uri};

    if (   !defined $$hin{whisker}
        || !defined $$hin{whisker}->{'MAGIC'}
        || $$hin{whisker}->{'MAGIC'} != 31339 )
    {
        $$hout{whisker}->{error} = 'Input hash not initialized';
        return 2;
    }

    if ( defined $_[0] ) {    # handle extra params
        my %hashref;
        if ( ref( $_[0] ) eq 'HASH' ) { %hashref = %{ $_[0] }; }
        else { %hashref = @_; }
        $$hin{whisker}->{$_} = $hashref{$_} foreach ( keys %hashref );
    }
    if ( defined $$hin{whisker}->{'anti_ids'} ) {    # handle anti_ids
        my %copy = %$hin;
        $copy{whisker} = {};
        %{ $copy{whisker} } = %{ $$hin{whisker} };
        encode_anti_ids( \%copy, $$hin{whisker}->{'anti_ids'} );
        $hin = \%copy;
    }

    # find/setup stream
    my $cache_key = stream_key($hin);
    my $stream;
    if ( !defined $http_host_cache{$cache_key} ) {
        $stream = stream_new($hin);
        $http_host_cache{$cache_key} = $stream;
    }
    else {
        $stream = $http_host_cache{$cache_key};
    }
    if ( !defined $stream ) {
        $$hout{whisker}->{error} = 'unable to allocate stream';
        return 2;
    }

    my $retry_count = $$hin{whisker}->{retry};
    my $puke_flag   = 0;
    my $ret         = 1;
    do {    # retries wrapper
        my ( $aret, $pass );

        if ( !$stream->{valid}->() ) {
            $stream->{clearall}->();
            if ( !$stream->{open}->($hin) ) {
                $$hout{whisker}->{error} =
                  'opening stream: ' . $stream->{error};
                $$hout{whisker}->{error} .=
                  '(reconnect problem after prior request)'
                  if ($puke_flag);
                return 2;
            }

            # freshly open stream/connection, handle auth
            if (   defined $$hin{whisker}->{proxy_host}
                && defined $$hin{whisker}->{auth_proxy_callback} )
            {
                $aret =
                  $$hin{whisker}->{auth_proxy_callback}
                  ->( $stream, $hin, $hout );
                return $aret if ( $aret != 0 );    # proxy auth error
            }
            if ( defined $$hin{whisker}->{auth_callback} ) {
                $aret =
                  $$hin{whisker}->{auth_callback}->( $stream, $hin, $hout );
                return 0     if ( $aret == 200 );    # auth not needed?
                return $aret if ( $aret != 0 );      # auth error
            }
        }

        _ssl_save_info( $hout, $stream )
          if ( $$hin{whisker}->{ssl} > 0
            && $$hin{whisker}->{ssl_save_info} > 0 );

        $ret = _http_do_request_ex( $stream, $hin, $hout );
        $puke_flag++
          if ( $ret == 1 && defined( $$hout{whisker}->{http_data_sent} ) );
        return $ret
          if ( $ret == 0 || $ret == 2 );    # success or fatal socket error
        $retry_count--;
    } while ( $retry_count >= 0 );

    # if we get here, we still had errors, but no more retries
    return $ret;

}

##################################################################

sub _http_do_request_ex {
    my ( $stream, $hin, $hout, $raw ) = @_;

    return 2 if ( !defined $stream );
    return 2 if ( !( defined $hin && ref($hin) ) );
    return 2 if ( !( defined $hout && ref($hout) ) );
    my $W = $hin->{whisker};

    # setup hash, if needed
    if ( !defined $$hout{whisker}->{MAGIC}
        || $$hout{whisker}->{MAGIC} != 31340 )
    {
        %$hout                     = ();
        $$hout{whisker}            = {};
        $$hout{whisker}->{'MAGIC'} = 31340;
        $$hout{whisker}->{uri}     = $$hin{whisker}->{uri};
    }

    ##### construct and send request
    $stream->{clear}->();

    if ( defined $raw && ref($raw) ) {
        $stream->{queue}->($$raw);

    }
    else {
        $stream->{queue}->( http_req2line($hin) );

        if ( $$W{version} ne '0.9' ) {
            $stream->{queue}->( http_construct_headers($hin) );
            $stream->{queue}->( $$W{raw_header_data} )
              if ( defined $$W{raw_header_data} );
            $stream->{queue}->( $$W{http_eol} );
            $stream->{queue}->( $$W{data} ) if ( defined $$W{data} );
        }    # http 0.9 support
    }

    # good time to fingerprint, if requested
    if ( defined $$W{request_fingerprint} ) {
        $$hout{whisker}->{request_fingerprint} =
          'md5:' . md5( $stream->{bufout} )
          if ( $$W{request_fingerprint} eq 'md5' );
        $$hout{whisker}->{request_fingerprint} =
          'md4:' . md4( $stream->{bufout} )
          if ( $$W{request_fingerprint} eq 'md4' );
    }

    # all data is wrangled...actually send it now
    if ( !$stream->{'write'}->() ) {
        $$hout{whisker}->{'error'} = 'sending request: ' . $stream->{error};
        $stream->{'close'}->();
        return 1;
    }

    # needed for SSL requests
    # NOTE: this is disabled because it's just a noop anyways
    # $stream->{writedone}->();

    $$hout{whisker}->{http_data_sent} = 1;
    $$hout{whisker}->{'lowercase_incoming_headers'} =
      $$W{'lowercase_incoming_headers'};

    ##### read and parse response
    my @H;
    if ( $$W{'version'} ne '0.9' ) {
        do {    # catch '100 Continue' responses
            my $resp = _http_getline($stream);

            if ( !defined $resp ) {
                $$hout{whisker}->{error} = 'error reading HTTP response';
                $$hout{whisker}->{data}  = $stream->{bufin};
                $stream->{'close'}->();
                return 1;
            }

            $$hout{whisker}->{'raw_header_data'} .= $resp
              if ( defined $$W{'save_raw_headers'} );

            if ( $resp !~
                /^([^\/]+)\/(\d\.\d)([ \t]+)(\d+)([ \t]*)(.*?)([\r\n]+)/ )
            {
                $$hout{whisker}->{'error'} = 'invalid HTTP response';
                $$hout{whisker}->{'data'}  = $resp;
                while ( defined( $_ = _http_getline($stream) ) ) {
                    $$hout{whisker}->{'data'} .= $_;
                }
                $stream->{'close'}->();
                return $$W{'invalid_protocol_return_value'} || 1;
            }

            $$hout{whisker}->{protocol}    = $1;
            $$hout{whisker}->{version}     = $2;
            $$hout{whisker}->{http_space1} = $3;
            $$hout{whisker}->{code}        = $4;
            $$hout{whisker}->{http_space2} = $5;
            $$hout{whisker}->{message}     = $6;
            $$hout{whisker}->{http_eol}    = $7;
            $$hout{whisker}->{'100_continue'}++ if ( $4 == 100 );

            @H = http_read_headers( $stream, $hin, $hout );
            if ( !$H[0] ) {
                $$hout{whisker}->{'error'} =
                  'Error in reading headers: ' . $H[1];
                $stream->{'close'}->();
                return 1;
            }

            if ( !defined $H[3] ) {    # connection
                my ($t) = utils_find_lowercase_key( $hin, 'connection' );
                $H[3] = $t || 'close';
            }

        } while ( $$hout{whisker}->{'code'} == 100 );

    }
    else {    # http ver 0.9, we need to fake it since headers are not sent
        $$hout{whisker}->{version}      = '0.9';
        $$hout{whisker}->{code}         = 200;
        $$hout{whisker}->{message} 	= '';
        $H[3]                           = 'close';
    }

    if ( $$hout{whisker}->{code}==404 && defined $$W{'shortcut_on_404'} ) {
        $stream->{'close'}->();
    }
    elsif ( defined $$W{data_sock} ) {
        $$hout{whisker}->{data_sock}   = $stream->{sock};
        $$hout{whisker}->{data_stream} = $stream;
    }
    else {
        if (
            $$W{'force_bodysnatch'}
            || (   $$W{'method'} ne 'HEAD'
                && $$hout{whisker}->{'code'} != 206
                && $$hout{whisker}->{'code'} != 102 )
          )
        {
            return 1
              if ( !http_read_body( $stream, $hin, $hout, $H[1], $H[2] ) );

            # {hide_chunked_responses} stuff follows
            if (   lc( $H[1] ) eq 'chunked'
                && defined $$hin{whisker}->{hide_chunked_responses}
                && $$hin{whisker}->{hide_chunked_responses} == 1
                && !defined $$hin{whisker}->{save_raw_chunks} )
            {
                $$hout{'Content-Length'} = length( $$hout{whisker}->{data} );
                utils_delete_lowercase_key( $hout, 'transfer-encoding' );
                my $new = [];
                my $cl  = 0;
                foreach ( @{ $$hout{whisker}->{header_order} } ) {
                    my $l = lc($_);
                    if ( $l eq 'content-length' ) {
                        $cl++;
                        next if ( $cl > 1 );
                    }
                    push @$new, $_ if ( $l ne 'transfer-encoding' );
                }
                push @$new, 'Content-Length' if ( $cl == 0 );
                $$hout{whisker}->{header_order} = $new;
            }
        }

        my ($ch) = LW2::utils_find_lowercase_key( $hin, 'connection' );
        my $cl = 0;
        $cl++
          if (
            (
                lc( $H[3] ) ne 'keep-alive' || ( defined $ch
                    && $ch =~ m/close/i )
            )
            && $$W{'force_open'} != 1
          );
        $cl++ if ( $$W{'force_close'} > 0 || $stream->{forceclose} > 0 );
        $cl++ if ( $$W{'ssl'} > 0 && $LW_SSL_KEEPALIVE == 0 );
        $stream->{'close'}->() if ($cl);
    }

    if ( defined $$W{'header_delete_on_success'}
        && ref( $$W{'header_delete_on_success'} ) )
    {
        foreach ( @{ $$W{'header_delete_on_success'} } ) {
            delete $hin->{$_} if ( exists $hin->{$_} );
        }
        delete $$W{header_delete_on_success};
    }

    $stream->{reqs}++;
    $$hout{whisker}->{'stats_reqs'}   = $stream->{reqs};
    $$hout{whisker}->{'stats_syns'}   = $stream->{syns};
    $$hout{whisker}->{'socket_state'} = $stream->{state};
    delete $$hout{whisker}->{'error'};    # no error
    return 0;

}

##################################################################

=item B<http_req2line>

Params: \%request, $uri_only_switch

Return: $request

req2line is used internally by http_do_request, as well as provides a
convienient way to turn a %request configuration into an actual HTTP request
line.  If $switch is set to 1, then the returned $request will be the URI
only ('/requested/page.html'), versus the entire HTTP request ('GET
/requested/page.html HTTP/1.0\n\n').  Also, if the 'full_request_override'
whisker config variable is set in %hin, then it will be returned instead
of the constructed URI.

=cut

sub http_req2line {
    my ( $S, $hin, $UO ) = ( '', @_ );
    $UO ||= 0;

    # notice: full_request_override can play havoc with proxy settings
    if ( defined $$hin{whisker}->{'full_request_override'} ) {
        return $$hin{whisker}->{'full_request_override'};

    }
    else {    # notice the components of a request--this is for flexibility
        if ( $UO != 1 ) {
            $S .= $$hin{whisker}->{'method'} . $$hin{whisker}->{'http_space1'};
            if ( $$hin{whisker}->{'include_host_in_uri'} > 0 ) {
                if ( $$hin{whisker}->{'ssl'} == 1 ) {
                    $S .= 'https://';
                }
                else {
                    $S .= 'http://';
                }

                if ( defined $$hin{whisker}->{'uri_user'} ) {
                    $S .= $$hin{whisker}->{'uri_user'};
                    if ( defined $$hin{whisker}->{'uri_password'} ) {
                        $S .= ':' . $$hin{whisker}->{'uri_password'};
                    }
                    $S .= '@';
                }

                $S .= $$hin{whisker}->{'host'} . ':' . $$hin{whisker}->{'port'};
            }
        }

        $S .=
            $$hin{whisker}->{'uri_prefix'}
          . $$hin{whisker}->{'uri'}
          . $$hin{whisker}->{'uri_postfix'};

        if ( defined $$hin{whisker}->{'parameters'}
            && $$hin{whisker}->{'parameters'} ne '' )
        {
            $S .=
                $$hin{whisker}->{'uri_param_sep'}
              . $$hin{whisker}->{'parameters'};
        }

        if ( $UO != 1 ) {
            if ( $$hin{whisker}->{'version'} ne '0.9' ) {
                $S .=
                    $$hin{whisker}->{'http_space2'}
                  . $$hin{whisker}->{'protocol'} . '/'
                  . $$hin{whisker}->{'version'};
            }
            $S .= $$hin{whisker}->{'http_eol'};
        }
    }
    return $S;
}

##################################################################

=item B<http_resp2line>

Params: \%response

Return: $response

http_resp2line provides a convienient way to turn a %response hash back 
into the original HTTP response line.

=cut

sub http_resp2line {
    my $hout = shift;
    my $out  = '';
    return undef if ( !defined $hout || !ref($hout) );
    return undef if ( $hout->{whisker}->{MAGIC} != 31340 );
    $out .= $$hout{whisker}->{protocol};
    $out .= '/';
    $out .= $$hout{whisker}->{version};
    $out .= $$hout{whisker}->{http_space1};
    $out .= $$hout{whisker}->{code};
    $out .= $$hout{whisker}->{http_space2};
    $out .= $$hout{whisker}->{message};
    $out .= $$hout{whisker}->{http_eol};
    return $out;
}

##################################################################

sub _http_getline {
    my $stream = shift;
    my ( $str, $t, $bc ) = ( '', 0, 0 );

    $t = index( $stream->{bufin}, "\n", 0 );
    while ( $t < 0 ) {
        return undef if !$stream->{read}->() || 
		length($stream->{bufin}) == $bc;
        $t = index( $stream->{bufin}, "\n", 0 );
    	$bc = length( $stream->{bufin} );
    }

    my $r = substr( $stream->{bufin}, 0, $t + 1 );
    $stream->{bufin} = substr( $stream->{bufin}, $t + 1 );

    #	substr($stream->{bufin},0,$t+1)='';
    return $r;
}

##################################################################

sub _http_get {    # read from socket w/ timeouts
    my ( $stream, $amount ) = @_;
    my ( $str, $t, $b )     = ( '', '', 0 );

    while ( $amount > length( $stream->{bufin} ) ) {
        return undef if !$stream->{read}->() ||
		length( $stream->{bufin} ) == $b;
	$b = length( $stream->{bufin} );
    }

    my $r = substr( $stream->{bufin}, 0, $amount );
    $stream->{bufin} = substr( $stream->{bufin}, $amount );

    #	substr($stream->{bufin},0,$amount)='';
    return $r;
}

##################################################################

sub _http_getall {
    my ( $tmp, $b, $stream, $max_size ) = ('', 0, @_);

    while ( $stream->{read}->() && length( $stream->{bufin} ) != $b) {
        last if ( $max_size && length( $stream->{bufin} ) >= $max_size );
        $b = length( $stream->{bufin} );	
    }
    ( $tmp, $stream->{bufin} ) = ( $stream->{bufin}, '' );
    $tmp = substr($tmp, 0, $max_size) if($max_size && 
    	length($tmp) > $max_size);
    return $tmp;
}

##################################################################

=item B<http_fixup_request>

Params: $hash_ref

Return: Nothing

This function takes a %hin hash reference and makes sure the proper 
headers exist (for example, it will add the Host: header, calculate the 
Content-Length: header for POST requests, etc).  For standard requests 
(i.e. you want the request to be HTTP RFC-compliant), you should call this 
function right before you call http_do_request.

=cut

sub http_fixup_request {
    my $hin = shift;

    return if ( !( defined $hin && ref($hin) ) );

    $$hin{whisker}->{uri} = '/' if ( $$hin{whisker}->{uri} eq '' );
    $$hin{whisker}->{http_space1}= ' ';
    $$hin{whisker}->{http_space2}= ' ';
    $$hin{whisker}->{protocol}= 'HTTP';
    $$hin{whisker}->{uri_param_sep}= '?';

    if ( $$hin{whisker}->{'version'} eq '1.1' ) {
        my ($host) = utils_find_lowercase_key($hin,'host');
        $$hin{'Host'} = $$hin{whisker}->{'host'} 
            if(!defined $host || $host eq '');
        $$hin{'Host'} .= ':' . $$hin{whisker}->{'port'}
          if ( index($$hin{'Host'},':') == -1 && 
          	( $$hin{whisker}->{port} != 80 || ( $$hin{whisker}->{ssl}==1 &&
              $$hin{whisker}->{port} != 443 ) ) );
        my ($conn) = utils_find_lowercase_key($hin,'connection');
        $$hin{'Connection'} = 'Keep-Alive' 
            if(!defined $conn || $conn eq '');

    } elsif( $$hin{whisker}->{'version'} eq '1.0' ){
        my ($conn) = utils_find_lowercase_key($hin,'connection');
        $$hin{'Connection'} = 'close' 
            if(!defined $conn || $conn eq '');
    }

    utils_delete_lowercase_key( $hin, 'content-length' );
    if ( $$hin{whisker}->{method} eq 'POST' || 
    		defined $$hin{whisker}->{data} ) {
	$$hin{whisker}->{data}||='';
        $$hin{'Content-Length'} = length( $$hin{whisker}->{'data'} );
        my ($v) = utils_find_lowercase_key( $hin, 'content-type' );
        if ( !defined $v || $v eq '' ) {
            $$hin{'Content-Type'} = 'application/x-www-form-urlencoded';
        }
    }

    #if(defined $$hin{whisker}->{'proxy_host'} && $$hin{whisker}->{ssl}==0){
    if ( defined $$hin{whisker}->{'proxy_host'} ) {
        $$hin{whisker}->{'include_host_in_uri'} = 1;
    }

}

##################################################################

=item B<http_reset>

Params: Nothing

Return: Nothing

The http_reset function will walk through the %http_host_cache, 
closing all open sockets and freeing SSL resources.  It also clears
out the host cache in case you need to rerun everything fresh.

Note: if you just want to close a single connection, and you have
a copy of the %request hash you used, you should use the http_close()
function instead.

=cut

sub http_reset {
    my $stream;

    foreach $stream ( keys %http_host_cache ) {
        $stream->{'close'}->() if(ref($stream));
        delete $http_host_cache{$stream};
    }
}

##################################################################

=item B<ssl_is_available>

Params: Nothing

Return: $boolean [, $lib_name, $version]

The ssl_is_available() function will inform you whether SSL requests
are allowed, which is dependant on whether the appropriate SSL
libraries are installed on the machine.  In scalar context, the
function will return 1 or 0.  In array context, the second element
will be the SSL library name that is currently being used by LW2,
and the third elment will be the SSL library version number.
Elements two and three (name and version) will be undefined if
called in array context and no SSL libraries are available.

=cut

sub ssl_is_available {
    return 0 if ( $LW_SSL_LIB == 0 );
    if ( $LW_SSL_LIB == 1 ) {
        return 1 if ( !wantarray() );
        return ( 1, "Net::SSLeay", $Net::SSLeay::VERSION );
    }
    elsif ( $LW_SSL_LIB == 2 ) {
        return 1 if ( !wantarray() );
        return ( 1, "Net::SSL", $Net::SSL::VERSION );
    }
    else {
        utils_carp('',"ssl_is_available: sanity check failed");
        return 0;
    }
}

##################################################################

sub _ssl_save_info {
    my ( $hr, $stream ) = @_;
    my $cert;

    if ( $stream->{streamtype} == 4 ) {
        my $SSL = $stream->{sock};
        $hr->{whisker}->{ssl_cipher} = Net::SSLeay::get_cipher($SSL);
        if ( $cert = Net::SSLeay::get_peer_certificate($SSL) ) {
            $hr->{whisker}->{ssl_cert_subject} =
              Net::SSLeay::X509_NAME_oneline(
                Net::SSLeay::X509_get_subject_name($cert) );
            $hr->{whisker}->{ssl_cert_issuer} =
              Net::SSLeay::X509_NAME_oneline(
                Net::SSLeay::X509_get_issuer_name($cert) );
        }
        return;
    }

    if ( $stream->{streamtype} == 5 ) {
        $hr->{whisker}->{ssl_cipher} = $stream->{sock}->get_cipher();
        if ( $cert = $stream->{sock}->get_peer_certificate() ) {
            $hr->{whisker}->{ssl_cert_subject} = $cert->subject_name();
            $hr->{whisker}->{ssl_cert_issuer}  = $cert->issuer_name();
        }
        return;
    }
}

##################################################################

=item B<http_read_headers>

Params: $stream, \%in, \%out

Return: $result_code, $encoding, $length, $connection

Read HTTP headers from the given stream, storing the results in %out.  On
success, $result_code will be 1 and $encoding, $length, and $connection
will hold the values of the Transfer-Encoding, Content-Length, and
Connection headers, respectively.  If any of those headers are not present,
then it will have an 'undef' value.  On an error, the $result_code will
be 0 and $encoding will contain an error message.

This function can be used to parse both request and response headers.

Note: if there are multiple Transfer-Encoding, Content-Length, or
Connection headers, then only the last header value is the one returned
by the function.

=cut

sub http_read_headers {
    my ( $stream, $in, $hout ) = @_;
    my $W = $in->{whisker};
    my ( $a, $b, $LC, $CL, $TE, $CO );

    # we use direct access into the stream buffers for quickest
    # parsing of the headers
    my $last;
    pos( $stream->{bufin} ) = 0;
    while (1) {
        $last = pos( $stream->{bufin} );
        if ( $stream->{bufin} !~ m/(.*?)[\r]{0,1}\n/g ) {
            if ( !$stream->{read}->() ) {
                last
                  if ( $$W{require_newline_after_headers} == 0
                    && length( $stream->{bufin} ) - 1 == $last );
                return ( 0, 'error reading in all headers' );
            }
            pos( $stream->{bufin} ) = $last;
            next;
        }
        last if ( $1 eq '' );

        # should we *not* puke on malformed header?
        return ( 0, 'malformed header' )
          if ( $1 !~ m/^([^:]+):([ \t]*)(.*)$/ );

        $$hout{whisker}->{'abnormal_header_spacing'}++ if ( $2 ne ' ' );

        $a  = $1;
        $b  = $3;
        $LC = lc($a);
        next if ( $LC eq 'whisker' );
        $TE = lc($b) if ( $LC eq 'transfer-encoding' );
        $CL = $b     if ( $LC eq 'content-length' );
        $CO = lc($b) if ( $LC eq 'connection' );
        push( @{ $$hout{whisker}->{cookies} }, $b )
          if ( $LC eq 'set-cookie' || $LC eq 'set-cookie2' );

        if ( $$W{'lowercase_incoming_headers'} > 0 ) {
            $a = $LC;
        }
        elsif ( $$W{'normalize_incoming_headers'} > 0 ) {
            $a = ucfirst($LC);
            $a = 'ETag' if ( $a eq 'Etag' );
            $a =~ s/(-[a-z])/uc($1)/eg;
        }

        push( @{ $$hout{whisker}->{header_order} }, $a );

        if ( defined $$hout{$a} && $$W{ignore_duplicate_headers} != 1 ) {
            $$hout{$a} = [ $$hout{$a} ] if ( !ref( $$hout{$a} ) );
            push( @{ $$hout{$a} }, $b );
        }
        else {
            $$hout{$a} = $b;
        }
    }

    my $found = pos( $stream->{bufin} );
    $$hout{whisker}->{'raw_header_data'} = substr( $stream->{bufin}, 0, $found )
      if ( defined $$W{'save_raw_headers'} );
    $stream->{bufin} = substr( $stream->{bufin}, $found );
    return ( 1, $TE, $CL, $CO );
}

##################################################################

=item B<http_read_body>

Params: $stream, \%in, \%out, $encoding, $length

Return: 1 on success, 0 on error (and sets $hout->{whisker}->{error})

Read the body from the given stream, placing it in $out->{whisker}->{data}.
Handles chunked encoding.  Can be used to read HTTP (POST) request or HTTP
response bodies.  $encoding parameter should be lowercase encoding type.

NOTE: $out->{whisker}->{data} is erased/cleared when this function is called,
leaving {data} to just contain this particular HTTP body.

=cut

sub http_read_body {
    my ( $temp, $stream, $hin, $hout, $enc, $len ) = ( '', @_ );
    my $max_size = $hin->{whisker}->{max_size} || 0;
    $$hout{whisker}->{data} = '';

    if ( defined $enc && lc($enc) eq 'chunked' ) {
        my $total = 0;
        my $x;
        my $saveraw = $$hin{whisker}->{save_raw_chunks} || 0;
        if ( !defined( $x = _http_getline($stream) ) ) {
            $$hout{whisker}->{'error'} = 'Error reading chunked data length';
            $stream->{'close'}->();
            return 0;
        }
        $a = $x;
        $a =~ tr/a-fA-F0-9//cd;
        if ( length($a) > 8 ) {
            $$hout{whisker}->{'error'} = 'Chunked size is too big: ' . $x;
            $stream->{'close'}->();
            return 0;
        }
        $len = hex($a);
        $len = $max_size if ( $max_size && $len > $max_size );

        $$hout{whisker}->{'data'} = $x if ($saveraw);

        while ( $len > 0 ) {    # chunked sucks
            if ( !defined( $temp = _http_get( $stream, $len ) ) ) {
                $$hout{whisker}->{'error'} = 'Error reading chunked data';
                $stream->{'close'}->();
                return 0;
            }
            $$hout{whisker}->{'data'} = $$hout{whisker}->{'data'} . $temp;
            $total += $len;
            if ( $max_size && $total >= $max_size ) {
                $stream->{'close'}->();
                return 1;
            }
            $temp = _http_getline($stream);
            $$hout{whisker}->{'data'} .= $temp if ( $saveraw && defined $temp );
            if ( defined $temp && $temp =~ /^[\r\n]*$/ ) {
                $temp = _http_getline($stream);
                $$hout{whisker}->{'data'} .= $temp
                  if ( $saveraw && defined $temp );
            }
            if ( !defined $temp ) {
                $$hout{whisker}->{'error'} = 'Error reading chunked data';
                $stream->{'close'}->();
                return 0;
            }
            $temp =~ tr/a-fA-F0-9//cd;
            if ( length($temp) > 8 ) {
                $$hout{whisker}->{'error'} =
                  'Chunked size is too big: ' . $temp;
                $stream->{'close'}->();
                return 0;
            }
            $len = hex($temp);
            $len = ( $max_size - $total )
              if ( $max_size && $len > ( $max_size - $total ) );
        }

        # read in trailer headers; currently doesn't account for max_size
        while ( defined( $_ = _http_getline($stream) ) ) {
            $$hout{whisker}->{'data'} .= $_ if ($saveraw);
            tr/\r\n//d;
            last if ( $_ eq '' );
        }

    }
    else {
        if ( defined $len ) {
            return 1 if ( $len <= 0 );
            $len = $max_size if ( $max_size && $len > $max_size );
            if (
                !defined(
                    $$hout{whisker}->{data} = _http_get( $stream, $len )
                )
              )
            {
                $stream->{'close'}->();

								# New LW2.5 feature: allow_short_reads will still return
								# success, even if all the data wasn't read.  This was
								# per request due to some 3Com switches sending out
								# the wrong content-length in HTTP response
								my $s = $$hin{whisker}->{allow_short_reads} || 0;
								if ( $s != 0 && length($stream->{'bufin'}) > 0 ) {
									# short read is requested, and there is some data, so
									# copy it over and return a non-error
									$$hout{whisker}->{'data'} = $stream->{'bufin'};
									return 1;
								}

                $$hout{whisker}->{'error'} =
                  'Error reading data: ' . $stream->{error};
                return 0;
            }
        }
        else {    # Yuck...read until server stops sending....
            $$hout{whisker}->{data} = _http_getall( $stream, $max_size );
            $stream->{'close'}->();
        }
        $$hout{whisker}->{'data'} ||= '';
    }
    return 1;
}

##################################################################

=item B<http_construct_headers>

Params: \%in

Return: $data

This function assembles the headers in the given hash into a data
string.

=cut

sub http_construct_headers {
    my $hin = shift;
    my ( %SENT, $output, $i );

    my $EOL = $hin->{whisker}->{http_eol} || "\x0d\x0a";
    if ( defined $hin->{whisker}->{header_order}
        && ref( $hin->{whisker}->{header_order} ) eq 'ARRAY' )
    {
        foreach ( @{ $hin->{whisker}->{header_order} } ) {
            next if ( $_ eq '' || $_ eq 'whisker' || !defined $hin->{$_} );
            if ( ref( $hin->{$_} ) ) {
                utils_croak("http_construct_headers: non-array header value reference")
                  if ( ref( $hin->{$_} ) ne 'ARRAY' );
                $SENT{$_} ||= 0;
                my $v = $$hin{$_}->[ $SENT{$_} ];
                $output .= "$_: $v$EOL";
            }
            else {
                $output .= "$_: $$hin{$_}$EOL";
            }
            $SENT{$_}++;
        }
    }

    foreach ( keys %$hin ) {
        next if ( $_ eq '' || $_ eq 'whisker' );
        if ( ref( $hin->{$_} ) ) {    # header with multiple values
	    utils_croak("http_construct_headers: non-array header value ref") 
	    	if ( ref( $hin->{$_} ) ne 'ARRAY' );
	    $SENT{$_} ||= 0;
	    for($i=$SENT{$_}; $i<~~@{ $hin->{$_} }; $i++) {
                $output .= "$_: " . $hin->{$_}->[$i] . $EOL;
            }
        }
        else {                       # normal header
            next if ( defined $SENT{$_} );
            $output .= "$_: $$hin{$_}$EOL";
        }
    }
    return $output;
}

##################################################################

=item B<http_close>

Params: \%request

Return: nothing

This function will close any open streams for the given request.

Note: in order for http_close() to find the right connection, all
original host/proxy/port parameters in %request must be the exact
same as when the original request was made.

=cut

sub http_close {
    my $hin       = shift;
    my $cache_key = stream_key($hin);
    return if ( !defined $http_host_cache{$cache_key} );
    my $stream = $http_host_cache{$cache_key};
    $stream->{'close'}->();
}

##################################################################

=item B<http_do_request_timeout>

Params: \%request, \%response, $timeout

Return: $result

This function is identical to http_do_request(), except that it
wraps the entire request in a timeout wrapper.  $timeout is the
number of seconds to allow for the entire request to be completed.

Note: this function uses alarm() and signals, and thus will only
work on Unix-ish platforms.  It should be safe to call on any
platform though.

=cut

sub http_do_request_timeout {
    my ( $req, $resp, $timeout ) = @_;
    $timeout ||= 30;

    my $result;
    eval {
        local $SIG{ALRM} = sub { die "timeout\n" };
        eval { alarm($timeout) };
        $result = LW2::http_do_request( $req, $resp );
        eval { alarm(0) };
    };
    if ($@) {
        $result                   = 1;
        $resp->{whisker}->{error} = 'Error with timeout wrapper';
        $resp->{whisker}->{error} = 'Total transaction timed out'
          if ( $@ =~ /timeout/ );
    }
    return $result;
}

