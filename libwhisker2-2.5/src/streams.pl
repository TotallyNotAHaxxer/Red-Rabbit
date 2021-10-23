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

@_stream_FUNCS = (
    [ 'open', 'close', 'read', 'write', 'writedone', 'valid' ],    # stream_NULL
    [ 'socket', 'all', 'socket', 'socket', 'noop', 'socket' ]
    ,    # stream_SOCKTCP   1
    [ 'socket', 'all', 'socket', 'socket', 'noop', 'never' ]
    ,    # stream_SOCKUDP   2
    [ 'file', 'all', 'socket', 'file', 'noop', 'never' ],   # stream_FILE      3
    [ 'ssl', 'all', 'ssl', 'ssl', 'noop', 'netssleay' ],    # stream_NETSSLEAY 4
    [ 'ssl', 'all', 'ssl', 'ssl', 'noop', 'never' ],        # stream_NETSSL    5
    [ 'buffer', 'buffer', 'buffer', 'buffer', 'noop',
        'never' ]                                           # stream_BUFFER    6
);

sub stream_key {
    my ( $key, $type, $wh ) = ( '', 1, shift );

    if ( defined $wh->{whisker}->{UDP} && $wh->{whisker}->{UDP} > 0 ) {
        $type = 2;
        $key  = 'udp:';
    }

    if ( $wh->{whisker}->{ssl} > 0 ) {
        $type = 4 if ( $LW_SSL_LIB == 1 );
        $type = 5 if ( $LW_SSL_LIB == 2 );
        $key = 'ssl:';
    }

    if ( defined $wh->{whisker}->{file_stream} ) {
        $type = 3;
        $key  = 'file=' . $wh->{whisker}->{file_stream} . ':';
    }

    if ( defined $wh->{whisker}->{buffer_stream} ) {
        $type = 6;
        $key  = 'buffer:';
    }

    my ( $x, $h, $p ) = (0);
    if ( defined $wh->{whisker}->{proxy_host} ) {
        $h = $wh->{whisker}->{proxy_host};
        $p = $wh->{whisker}->{proxy_port} || 80;
        $x++;
        $key .= 'proxy:';
        if ( $type == 5 ) {
            $x                = 0;
            $ENV{HTTPS_PROXY} = "$h:$p";
            $h                = $wh->{whisker}->{host};
            $p                = $wh->{whisker}->{port};
        }
    }
    else {
        $h = $wh->{whisker}->{host};
        $p = $wh->{whisker}->{port};
    }

    $key .= $h . ':' . $p;
    if ( defined $wh->{whisker}->{stream_num} ) {
        $key .= '/' . $wh->{whisker}->{stream_num};
    }

    return $key if ( !wantarray() );
    return ( $type, $h, $p, $x, $key );
}

sub stream_setsock {
    my $fd = shift;
    my $wh = http_new_request( host => 'localhost', port => 80, ssl => 0 );
    my $xr = stream_new($wh);
    return undef if ( $xr->{streamtype} != 1 );
    $xr->{sock}  = $fd;
    $xr->{state} = 1;
    $xr->{eof}   = 0;
    $xr->{clearall}->();
    return $xr;
}

{
    $SYMCOUNT = 0;

    sub stream_new {
        my ( $c, $rh ) = ( 0, shift );
        my $sock = _stream_newsock();
        my %x;
        %x = (
            bufin      => '',
            bufout     => '',
            error      => '',
            streamtype => 0,
            eof        => 0,
            ctx        => undef,
            sock       => $sock,
            state      => 0,
            syns       => 0,
            reqs       => 0,
            timeout    => $rh->{whisker}->{timeout} || 10,
            nonblock   => 0,
            forceclose => 0
        );

        ( $x{streamtype}, $x{chost}, $x{cport}, $x{proxy}, $x{key} ) =
          stream_key($rh);
        return undef if ( $x{streamtype} == 0 );
        return undef
          if (
            $LW_SSL_LIB == 0
            && (   $x{streamtype} == 4
                || $x{streamtype} == 5 )
          );
        return undef
          if ( $x{streamtype} != 3
            && $x{streamtype} != 6
            && !defined $Socket::VERSION );

        $x{nonblock} = $LW_NONBLOCK_CONNECT if ( $x{streamtype} == 1 );
        $x{forceclose} = 1 if ( $x{streamtype} == 5 );

        $x{slurp} = $rh->{whisker}->{trailing_slurp} || 0;

        my @N = @{ $_stream_FUNCS[ $x{streamtype} ] };
        for ( $c = 0 ; $c < 6 ; $c++ ) {
            my $n = $_stream_FUNCS[0]->[$c];
            my $e =
              '$x{"' . $n . '"}=sub{&_stream_' . $N[$c] . "_$n" . '(\%x,@_)}';
            eval "$e";
        }
        $x{queue} = sub { $x{bufout} .= shift };
        $x{clearall} = sub { $x{bufin} = $x{bufout} = '' };
        $x{clear} = sub { $x{bufout} = '' };
        return bless \%x, 'LW2::stream';
    }

    sub _stream_newsock {    # same as Symbol::gensym
        my $pkg  = "LW2::";
        my $name = "_STREAM_" . $SYMCOUNT++;
        delete $$pkg{$name};
        return \*{ $pkg . $name };
    }
}

sub _stream_all_close {
    my $xr = shift;
    $xr->{state} = 0;
    if ( $xr->{streamtype} == 4 ) {
        eval { $xr->{sock}->shutdown() };
        eval { close( $xr->{origsock} ) };

        #		eval { Net::SSLeay::free($xr->{sock}) };
    }
    else {
        eval { close( $xr->{sock} ) };
    }
}

sub _stream_never_valid {
    return 0;
}

sub __bad_netssleay_error {
    my $err = Net::SSLeay::ERR_get_error;
    return 0
      if ( $err == Net::SSLeay::ERROR_NONE
        || $err == Net::SSLeay::ERROR_WANT_READ
        || $err == Net::SSLeay::ERROR_WANT_WRITE );
    return 1;
}

sub _stream_netssleay_valid {
    my $xr = shift;
    return 0 if ( $LW_SSL_KEEPALIVE == 0 || $xr->{state} == 0 );
    return 0 if ( &Net::SSLeay::OPENSSL_VERSION_NUMBER < 0x0090601f );

    my $lo = Net::SSLeay::pending( $xr->{sock} );
    if ( $lo > 0 ) {    # leftover data to slurp
        if ( !$xr->{slurp} ) {
            return 0 if ( !_stream_ssl_read($xr) );
        }
        else {

            # todo
            #$xr->{slurped}.=$x."\0";
        }
    }
    return 0 if ( __bad_netssleay_error() );

    my ( $r, $e, $vin ) = ( undef, undef, '' );
    my $fno = fileno( $xr->{origsock} );
    vec( $vin, $fno, 1 ) = 1;
    if ( select( ( $r = $vin ), undef, ( $e = $vin ), .0001 ) ) {
        return 0 if ( vec( $e, $fno, 1 ) );
        if ( vec( $r, $fno, 1 ) ) {    # waiting data, let's peek
            my $temp = Net::SSLeay::peek( $xr->{sock}, 1 );
            return 0 if ( __bad_netssleay_error() );
            return 0 if ( $temp <= 0 );
        }
    }

    return 1;
}

sub _stream_socket_valid {
    my $xr = shift;
    return 0 if ( $xr->{state} == 0 );
    my ( $o, $vin ) = ( undef, '' );
    vec( $vin, fileno( $xr->{sock} ), 1 ) = 1;
    if ( select( ( $o = $vin ), undef, undef, .0001 ) ) {
        my ( $hold, $res );
        do {
            $res = sysread( $xr->{sock}, $hold, 4096 );
            return _stream_err( $xr, 1, 'is_valid sysread failed' )
              if ( !defined $res );    # error
            return 0 if ( $res == 0 ); # EOF
            if ( !$xr->{slurp} ) {
                $xr->{bufin} .= $hold;
            }
            else {
                $xr->{slurped} .= $hold . "\0";
            }
        } while ( $res && select( ( $o = $vin ), undef, undef, .0001 ) );
    }
    return 1;
}

sub _stream_socket_read {
    my $xr = shift;
    return 0 if ( $xr->{state} == 0 );
    my ( $vin, $t ) = ( '', '' );
    vec( $vin, fileno( $xr->{sock} ), 1 ) = 1;
    return 0 if ( !select( $vin, undef, undef, $xr->{timeout} ) );
    my $res = sysread( $xr->{sock}, $t, 4096 );
    return _stream_err( $xr, 1, 'sysread failed' ) if ( !defined $res );
    if ( $res == 0 ) {
        $xr->{eof} = 1;
        return 0;
    }
    $xr->{bufin} .= $t;
    $xr->{eof} = 0;
    return 1;
}

sub _stream_ssl_read {
    my ( $xr, $t ) = ( shift, '' );
    return 0 if ( $xr->{state} == 0 );
    if ( $xr->{streamtype} == 4 ) {
        local $SIG{ALRM} = sub { die "lw_timeout\n" };
        local $SIG{PIPE} = sub { die "lw_pipe\n" };
        eval {
            eval { alarm( $xr->{timeout} ) };

            #			sleep(1) while(!Net::SSLeay::pending($xr->{sock}));
            $t = Net::SSLeay::read( $xr->{sock} );
            eval { alarm(0) };
        };
        return 0 if ( $@ || __bad_netssleay_error() || !defined $t || $t eq '' );
    }
    elsif ( $xr->{streamtype} == 5 ) {
        return 0 if ( !$xr->{sock}->read( $t, 4096 ) );
    }
    $xr->{bufin} .= $t;
    return 1;
}

sub _stream_noop_writedone { }

sub _stream_ssl_writedone {
    my $xr = shift;
    if ( $xr->{streamtype} == 4 ) {    # Net::SSLeay
        shutdown $xr->{origsock}, 1;
    }
    else {                             # Net::SSL
                                       #shutdown $xr->{sock}, 1;
    }
}

sub _stream_socket_write {
    my ( $xr, $data, $v, $wrote ) = ( shift, shift, '', 0 );
    return 0 if ( $xr->{state} == 0 );
    $xr->{bufout} .= $data if ( defined $data );
    my $len = length( $xr->{bufout} );
    return 1 if ( $len == 0 );
    vec( $v, fileno( $xr->{sock} ), 1 ) = 1;
    return _stream_err( $xr, 1, 'stream write test failed' )
      if ( !select( undef, $v, undef, .0001 ) );
    my $piperr = 0;
    local $SIG{PIPE} = sub { $piperr++ };

    #	$wrote=syswrite($xr->{sock},$xr->{bufout},$len);
    #	return _stream_err($xr,1,'syswrite failed')
    #		if(!defined $wrote || $piperr);
    #	$xr->{error} = 'could not send entire queue' && return 0
    #		if($wrote!=$len);
    #	$xr->{bufout}='';
    #	return 1;

    do {
        $wrote = syswrite( $xr->{sock}, $xr->{bufout}, $len );
        if ( defined $wrote ) {
            substr( $xr->{bufout}, 0, $wrote ) = '';
        }
        else {
            if ( $! != EWOULDBLOCK ) {
                $piperr++;
            }
            else {
                vec( $v, fileno( $xr->{sock} ), 1 ) = 1;
                $piperr++ if ( !select( undef, $v, undef, $xr->{timeout} ) );
            }
        }
        return _stream_err( $xr, 1, 'syswrite failed' ) if ($piperr);
        $len = length( $xr->{bufout} );
    } while ( $len > 0 );
    return 1;
}

sub _stream_ssl_write {
    my ( $xr, $data, $wrote, $err ) = ( shift, shift, 0, '' );
    return 0 if ( $xr->{state} == 0 );
    $xr->{bufout} .= $data if ( defined $data );
    my $len = length( $xr->{bufout} );
    return 1 if ( $len == 0 );
    if ( $xr->{streamtype} == 4 ) {
        ( $wrote, $err ) =
          Net::SSLeay::ssl_write_all( $xr->{sock}, \$xr->{bufout} );
        if ( __bad_netssleay_error() || !$wrote ) {
            $xr->{error} = "SSL error: $err";
            return 0;
        }
        if ( $wrote != $len ) {
            $xr->{error} = 'could not send entire queue';
            return 0;
        }
    }
    elsif ( $xr->{streamtype} == 5 ) {
        $xr->{sock}->print( $xr->{bufout} );

        # bummer, no error checking?
    }
    $xr->{bufout} = '';
    return 1;
}

sub _stream_socket_alloc {
    my ( $xr, $wh ) = @_;

    if ( $xr->{streamtype} == 2 ) {
        return _stream_err( $xr, 0, 'socket problems (UDP)' )
          if (
            !socket(
                $xr->{sock}, PF_INET,
                SOCK_DGRAM, getprotobyname('udp') || 0
            )
          );
    }
    else {
        return _stream_err( $xr, 0, 'socket() problems' )
          if (
            !socket(
                $xr->{sock}, PF_INET,
                SOCK_STREAM, getprotobyname('tcp') || 0
            )
          );
    }

    if ( defined $wh->{whisker}->{bind_socket} ) {
        my $p = $wh->{whisker}->{bind_port} || '*';
        $p =~ tr/0-9*//cd;
        return _stream_err( $xr, 0, 'Bad bind_port value' )
          if ( $p eq '' );
        my $a = INADDR_ANY;
        $a = inet_aton( $wh->{whisker}->{bind_addr} )
          if ( defined $wh->{whisker}->{bind_addr} );
        return _stream_err( $xr, 0, 'Bad bind_addr value' )
          if ( !defined $a );
        if ( $p =~ tr/*// ) {
            for ( $p = 14011 ; $p < 65535 ; $p++ ) {
                if ( !bind( $xr->{sock}, sockaddr_in( $p, $a ) ) ) {
                    return _stream_err( $xr, 0, 'bind() on socket failed' )
                      if ( $! ne 'Address already in use' );
                }
                else {
                    last;
                }
            }
            return _stream_err( $xr, 0, 'bind() cannot find open socket' )
              if ( $p >= 65535 );
        }
        else {
            return _stream_err( $xr, 0, 'bind() on socket failed' )
              if ( !bind( $xr->{sock}, sockaddr_in( $p, $a ) ) );
        }
    }

    if ( !defined $xr->{iaton} ) {
        $xr->{iaton} = inet_aton( $xr->{chost} );
        return _stream_err( $xr, 0, 'can\'t resolve hostname' )
          if ( !defined $xr->{iaton} );
    }
    $xr->{socket_alloc}++;
    return 1;
}

sub _stream_socket_nonblock {
    my ( $fl, $xr, $nonblock ) = ( 0, @_ );

    if ( $^O =~ /Win32/ ) {
        $fl = 1 if ($nonblock);

        # 0x8004667e = FIONBIO in Winsock2.h
        if ( !ioctl( $xr->{sock}, 0x8004667e, \$fl ) ) {
            return 0;
        }
    }
    else {
        if ( !( $fl = fcntl( $xr->{sock}, F_GETFL, 0 ) ) ) {
            return 0;
        }
        $fl |= O_NONBLOCK if ($nonblock);
        $fl &= ~O_NONBLOCK if ( !$nonblock );
        if ( !( fcntl( $xr->{sock}, F_SETFL, $fl ) ) ) {
            return 0;
        }

    }
    return 1;
}

sub _stream_socket_open {
    my ( $vin, $xr, $wh ) = ( '', @_ );
    return 0 if ( !defined $wh );

    $xr->{'close'}->() if ( $xr->{state} > 0 );
    return 0 if ( !_stream_socket_alloc( $xr, $wh ) );
    $xr->{timeout} = $wh->{whisker}->{timeout} || 10;

    if ( $xr->{nonblock} ) {
        if ( !_stream_socket_nonblock( $xr, 1 ) ) {
            $xr->{nonblock} = 0;
            $LW_NONBLOCK_CONNECT = 0;
        }
        else {
            my $R =
              connect( $xr->{sock}, sockaddr_in( $xr->{cport}, $xr->{iaton} ) );
            if ( !$R ) {
                return _stream_err( $xr, 1, 'can\'t connect (connect error)' )
                  if ( $! != EINPROGRESS && $! != EWOULDBLOCK );
                vec( $vin, fileno( $xr->{sock} ), 1 ) = 1;
                return _stream_err( $xr, 1, 'can\'t connect (timeout)' )
                  if ( !select( undef, $vin, $vin, $xr->{timeout} )
                    || !getpeername( $xr->{sock} ) );
            }

            # leave in nonblock for normal TCP
            #			if($xr->{streamtype} != 1 && !_stream_socket_nonblock($xr,0)){
            #				$LW_NONBLOCK_CONNECT=0;
            #				return _stream_err($xr,1,'setting sock to block');
            #			}
        }
    }

    if ( !$xr->{nonblock} ) {
        eval {
            local $SIG{ALRM} = sub { die "timeout\n" };
            eval { alarm( $xr->{timeout} ) };
            if (
                !connect(
                    $xr->{sock}, sockaddr_in( $xr->{cport}, $xr->{iaton} )
                )
              )
            {
                eval { alarm(0) };
                die "connect failed\n";
            }
            eval { alarm(0) };
        };
        return _stream_err( $xr, 0,
            'can\'t connect (' . substr( $@, 0, index( $@, "\n" ) ) . ')' )
          if ($@);
    }

    binmode( $xr->{sock} );
    my $S = select( $xr->{sock} );
    $|++;
    select($S);
    $xr->{state} = 1;
    $xr->{syns}++;
    return 1;
}

sub _stream_ssl_open {
    my ( $xr, $wh ) = @_;
    return 0         if ( !defined $wh );
    $xr->{close}->() if ( $xr->{state} > 0 );
    my $W = $wh->{whisker};

    if ( $xr->{streamtype} == 5 ) {

        # these have to always be set, to overwrite any previous
        # set values (using ENV is a crappy way to do this)
        $ENV{HTTPS_KEY_FILE}  = $W->{ssl_rsacertfile} || '';
        $ENV{HTTPS_CERT_FILE} = $W->{ssl_certfile}    || '';
	eval {
            $xr->{sock}           = Net::SSL->new(
                PeerAddr => $xr->{chost},
                PeerPort => $xr->{cport},
                Timeout  => $xr->{timeout}
            );
	};
        return _stream_err( $xr, 0, 'can\'t connect: ' . $@ ) 
		if ($@ || !defined $xr->{sock});
        $xr->{sock}->autoflush(1);
        $xr->{state} = 1;

        # Net::SSL doesn't use stream_socket_open, so fake syns
        $xr->{syns}++;
        return 1;
    }

    return 0 if ( $xr->{streamtype} != 4 );

    # otherwise, we're stream_NETSSLEAY

    if ( !defined $xr->{ctx} ) {
        return _stream_err( $xr, 0, 'ssl ctx create' )
          if ( !( $xr->{ctx} = Net::SSLeay::CTX_new() ) );
        Net::SSLeay::CTX_set_options( $xr->{ctx}, &Net::SSLeay::OP_ALL );
        if ( defined $W->{ssl_rsacertfile} ) {
            if (
                !(
                    Net::SSLeay::CTX_use_RSAPrivateKey_file(
                        $xr->{ctx}, $W->{ssl_rsacertfile},
                        &Net::SSLeay::FILETYPE_PEM
                    )
                )
              )
            {
                return _stream_err( $xr, 0, 'ssl ctx rsacert' );
            }
        }
        if ( defined $W->{ssl_certfile} ) {
            if (
                !(
                    Net::SSLeay::CTX_use_certificate_file(
                        $xr->{ctx}, $W->{ssl_certfile},
                        &Net::SSLeay::FILETYPE_PEM
                    )
                )
              )
            {
                return _stream_err( $xr, 0, 'ssl ctx cert' );
            }
        }
    }

		# just to be safe, catch any errors that didn't get returned
		return _stream_err($xr, 0, 'ssl setup error' )
			if( __bad_netssleay_error() );

    return _stream_err( $xr, 0, 'ssl create new' )
      if ( !( $xr->{sslobj} = Net::SSLeay::new( $xr->{ctx} ) ) );
    if ( defined $W->{ssl_ciphers} ) {
        if (
            !(
                Net::SSLeay::set_cipher_list(
                    $xr->{sslobj}, $W->{ssl_ciphers}
                )
            )
          )
        {
            return _stream_err( $xr, 0, 'ssl set ciphers' );
        }
    }

    # now we use a normal socket to connect
    return 0 if ( !_stream_socket_open( $xr, $wh ) );
    $xr->{state} = 1;

    if ( $xr->{proxy} ) {
        my $C = 'CONNECT ' . $W->{host} . ':' . $W->{port} . " HTTP/1.0\r\n";
        $C .= 'Proxy-Authorization: ' . $wh->{'Proxy-Authorization'} . "\r\n"
          if ( defined $wh->{'Proxy-Authorization'} );
        $C .= "\r\n";

        my $r = syswrite( $xr->{sock}, $C, length($C) );
        return _stream_err( $xr, 1, 'sending proxy connect string' )
          if ( !defined $r || $r != length($C) );

        # now we need to read proxy response and parse it
        do {
            return _stream_err( $xr, 1, 'ssl proxy request failed' )
              if ( !_stream_socket_read($xr) );
          } while ( index( $xr->{bufin}, "\n\n" ) == -1
            && index( $xr->{bufin}, "\r\n\r\n" ) == -1 );
        return _stream_err( $xr, 1, 'proxy couldn\'t make connection' )
          if ( $xr->{bufin} !~ /^HTTP\/1.[0-9]+\W+200/ );

        #$xr->{bufin}='';
        $xr->{clearall}->();
    }

    Net::SSLeay::set_fd( $xr->{sslobj}, fileno( $xr->{sock} ) );
    Net::SSLeay::set_session( $xr->{sslobj}, $xr->{sslsession} )
      if ( defined $xr->{sslsession} );
    return _stream_err( $xr, 1, 'ssl connect failed' )
      if ( !( Net::SSLeay::connect( $xr->{sslobj} ) ) ||
      	__bad_netssleay_error() );

    # my $x = Net::SSLeay::ctrl( $xr->{sslobj}, 6, 0, '' );
    $xr->{sslsession} = Net::SSLeay::get_session( $xr->{sslobj} )
      if ( defined $W->{ssl_resume} && $W->{ssl_resume} > 0 );

    # little trickery to abstract/normalize stuff
    $xr->{origsock} = $xr->{sock};
    $xr->{sock}     = $xr->{sslobj};
    return 1;
}

sub _stream_file_open {
    my ( $xr, $wh ) = @_;
    $xr->{close}->() if ( $xr->{state} > 0 );
    my $file = $wh->{whisker}->{file_stream};
    return _stream_err( $xr, 0, 'invalid file' )
      if ( !-e $file || !-f $file );
    return _stream_err( $xr, 0, 'file open failure' )
      if ( !sysopen( $xr->{sock}, $file, 'r' ) );
    binmode($xr->{sock}); # Stupid Windows
    $xr->{state} = 1;
}

sub _stream_file_write {
    my $xr = shift;
    $xr->{bufout} = '';
    return 1;
}

sub _stream_buffer_open {
    my ( $xr, $wh ) = @_;
    $xr->{close}->() if ( $xr->{state} > 0 );
    $xr->{state} = 1;
}

sub _stream_buffer_close {
    my $xr = shift;
    $xr->{state} = 0;
    $xr->{bufout} = $xr->{bufin} = '';
}

sub _stream_buffer_read {
    my $xr = shift;
    return 0 if ( $xr->{state} == 0 );
    if ( length( $xr->{bufout} ) > 0 ) {
        $xr->{bufin} .= $xr->{bufout};
        $xr->{bufout} = '';
    }
    if ( length( $xr->{bufin} ) == 0 ) {
        $xr->{eof} = 1;
        return 0;
    }
    $xr->{eof} = 0;
    return 1;
}

sub _stream_buffer_write {
    my ( $xr, $data ) = ( shift, shift );
    return 0 if ( $xr->{state} == 0 );
    $xr->{bufout} .= $data if ( defined $data );
    my $len = length( $xr->{bufout} );
    return 1 if ( $len == 0 );
    $xr->{bufin} .= $xr->{bufout};
    $xr->{bufout} = '';
    return 1;
}

sub _stream_err {
    my ( $xr, $close, $error ) = @_;
    $xr->{error} = $error;
    $xr->{error} .= ": $!" if ( defined $! && $! ne '' );
    $xr->{'close'}->() if ($close);
    $xr->{state} = 0;
    return 0;
}

