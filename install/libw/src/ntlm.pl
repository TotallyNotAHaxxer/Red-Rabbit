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

########################################################################

=item B<ntlm_new>

Params: $username, $password [, $domain, $ntlm_only]

Return: $ntlm_object

Returns a reference to an array (otherwise known as the 'ntlm object')
which contains the various informations specific to a user/pass combo.
If $ntlm_only is set to 1, then only the NTLM hash (and not the LanMan
hash) will be generated.  This results in a speed boost, and is typically
fine for using against IIS servers.

The array contains the following items, in order:
username, password, domain, lmhash(password), ntlmhash(password)

=cut

sub ntlm_new {
    my ( $user, $pass, $domain, $flag ) = @_;
    $flag ||= 0;
    return undef if ( !defined $user );
    $pass   ||= '';
    $domain ||= '';
    my @a = ( "$user", "$pass", "$domain", undef, undef );
    my $t;

    if ( $flag == 0 ) {
        $t = substr( $pass, 0, 14 );
        $t =~ tr/a-z/A-Z/;
        $t .= "\0" x ( 14 - length($t) );
        $a[3] = des_E_P16($t);    # LanMan password hash
        $a[3] .= "\0" x ( 21 - length( $a[3] ) );
    }

    $t = md4( encode_unicode($pass) );
    $t =~ s/([a-z0-9]{2})/sprintf("%c",hex($1))/ieg;
    $t .= "\0" x ( 21 - length($t) );
    $a[4] = $t;                   # NTLM password hash

    &des_cache_reset();           # reset the keys hash
    return \@a;
}

########################################################################

sub ntlm_generate_responses {
    my ( $obj, $chal ) = @_;
    return ( undef, undef ) if ( !defined $obj || !defined $chal );
    return ( undef, undef ) if ( !ref($obj) );
    my $x = '';
    $x = des_E_P24( $obj->[3], $chal ) if ( defined $obj->[3] );
    return ( $x, des_E_P24( $obj->[4], $chal ) );
}

########################################################################

=item B<ntlm_decode_challenge>

Params: $challenge

Return: @challenge_parts

Splits the supplied challenge into the various parts.  The returned array
contains elements in the following order:

unicode_domain, ident, packet_type, domain_len, domain_maxlen,
domain_offset, flags, challenge_token, reserved, empty, raw_data

=cut

sub ntlm_decode_challenge {
    return undef if ( !defined $_[0] );
    my $chal = shift;
    my @res;

    @res = unpack( 'Z8VvvVVa8a8a8', substr( $chal, 0, 48 ) );
    push( @res, substr( $chal, 48 ) );
    unshift( @res, substr( $chal, $res[4], $res[2] ) );
    return @res;
}

########################################################################

sub ntlm_header {
    my ( $s, $h, $o ) = @_;
    my $l = length($s);
    return pack( 'vvV', 0, 0, $o - $h ) if ( $l == 0 );
    return pack( 'vvV', $l, $l, $o );
}

########################################################################

=item B<ntlm_client>

Params: $ntlm_obj [, $server_challenge]

Return: $response

ntlm_client() is responsible for generating the base64-encoded text you
include in the HTTP Authorization header.  If you call ntlm_client()
without a $server_challenge, the function will return the initial NTLM
request packet (message packet #1).  You send this to the server, and
take the server's response (message packet #2) and pass that as
$server_challenge, causing ntlm_client() to generate the final response
packet (message packet #3).

Note: $server_challenge is expected to be base64 encoded.

=cut

sub ntlm_client {
    my ( $obj, $p ) = @_;
    my $resp = "NTLMSSP\0";

    return undef if ( !defined $obj || !ref($obj) );

    if ( defined $p && $p ne '' ) {    # answer challenge
        $p =~ tr/ \t\r\n//d;
        $p = decode_base64($p);
        my @c  = ntlm_decode_challenge($p);
        my $uu = encode_unicode( $obj->[0] );    # username
        $resp .= pack( 'V', 3 );
        my ( $hl, $hn ) = ntlm_generate_responses( $obj, $c[7] );    # token
        return undef if ( !defined $hl || !defined $hn );
        my $o = 64;
        $resp .= ntlm_header( $hl, 64, $o );                         # LM hash
        $resp .= ntlm_header( $hn, 64, ( $o += length($hl) ) );      # NTLM hash
        $resp .= ntlm_header( $c[0], 64, ( $o += length($hn) ) );    # domain
        $resp .= ntlm_header( $uu, 64, ( $o += length( $c[0] ) ) );  # username
        $resp .= ntlm_header( $uu, 64, ( $o += length($uu) ) );    # workstation
        $resp .= ntlm_header( '', 64, ( $o += length($uu) ) );     # session
        $resp .= pack( 'V', $c[6] );
        $resp .= $hl . $hn . $c[0] . $uu . $uu;

    }
    else {    # initiate challenge
        $resp .= pack( 'VV', 1, 0x0000b207 );
        $resp .= ntlm_header( $obj->[0], 32, 32 );
        $resp .= ntlm_header( $obj->[2], 32, 32 + length( $obj->[0] ) );
        $resp .= $obj->[0] . $obj->[2];
    }

    return encode_base64( $resp, '' );
}

########################################################################

sub _ntlm_auth_callback {
    my ( $stream, $hi, $ho, $pflag ) = @_;
    my ( $ntlmobj, $header, $req_pre, $req_post, $aheader, $work, $ecode );
    my ($rheader);
    $pflag ||= 0;

    if ($pflag) {
        $ntlmobj                  = $$hi{whisker}->{auth_proxy_data};
        $header                   = 'Proxy-Authorization';
        $rheader                  = 'proxy-authenticate';
        $ecode                    = 407;
        $hi->{'Proxy-Connection'} = 'Keep-Alive';
    }
    else {
        $ntlmobj          = $$hi{whisker}->{auth_data};
        $header           = 'Authorization';
        $rheader          = 'www-authenticate';
        $ecode            = 401;
        $hi->{Connection} = 'Keep-Alive';
    }

    $ho->{whisker}->{error} = 'NTLM ' . $header;
    $hi->{$header} = 'NTLM ' . ntlm_client($ntlmobj);
    my $ret = _http_do_request_ex( $stream, $hi, $ho );
    return $ret if ($ret);
    return 200  if ( $$ho{whisker}->{code} == 200 );
    return 1    if ( $$ho{whisker}->{code} != $ecode );

    my $thead = utils_find_lowercase_key( $ho, $rheader );
    return 1 if ( !defined $thead );

    my ( $found, @auths );
    if ( ref($thead) ) { @auths = @$thead; }
    else { push @auths, $thead; }
    foreach (@auths) {
        $found = $1 if (m/^NTLM (.+)$/);
    }
    return 1 if ( !defined $found );

    $hi->{$header} = 'NTLM ' . ntlm_client( $ntlmobj, $found );
    push @{ $hi->{whisker}->{header_delete_on_success} }, $header;
    return 0;
}

sub _ntlm_auth_proxy_callback {
    return _ntlm_auth_callback( $_[0], $_[1], $_[2], 1 );
}

########################################################################

{    # start of DES local container #######################################
    my $generated = 0;
    my $perm1     = [
        57, 49, 41, 33, 25, 17, 9,  1,  58, 50, 42, 34, 26, 18,
        10, 2,  59, 51, 43, 35, 27, 19, 11, 3,  60, 52, 44, 36,
        63, 55, 47, 39, 31, 23, 15, 7,  62, 54, 46, 38, 30, 22,
        14, 6,  61, 53, 45, 37, 29, 21, 13, 5,  28, 20, 12, 4
    ];
    my $perm2 = [
        14, 17, 11, 24, 1,  5,  3,  28, 15, 6,  21, 10, 23, 19, 12, 4,
        26, 8,  16, 7,  27, 20, 13, 2,  41, 52, 31, 37, 47, 55, 30, 40,
        51, 45, 33, 48, 44, 49, 39, 56, 34, 53, 46, 42, 50, 36, 29, 32
    ];
    my $perm3 = [
        58, 50, 42, 34, 26, 18, 10, 2, 60, 52, 44, 36, 28, 20, 12, 4,
        62, 54, 46, 38, 30, 22, 14, 6, 64, 56, 48, 40, 32, 24, 16, 8,
        57, 49, 41, 33, 25, 17, 9,  1, 59, 51, 43, 35, 27, 19, 11, 3,
        61, 53, 45, 37, 29, 21, 13, 5, 63, 55, 47, 39, 31, 23, 15, 7
    ];
    my $perm4 = [
        32, 1,  2,  3,  4,  5,  4,  5,  6,  7,  8,  9,  8,  9,  10, 11,
        12, 13, 12, 13, 14, 15, 16, 17, 16, 17, 18, 19, 20, 21, 20, 21,
        22, 23, 24, 25, 24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32, 1
    ];
    my $perm5 = [
        16, 7, 20, 21, 29, 12, 28, 17, 1,  15, 23, 26, 5,  18, 31, 10,
        2,  8, 24, 14, 32, 27, 3,  9,  19, 13, 30, 6,  22, 11, 4,  25
    ];
    my $perm6 = [
        40, 8, 48, 16, 56, 24, 64, 32, 39, 7, 47, 15, 55, 23, 63, 31,
        38, 6, 46, 14, 54, 22, 62, 30, 37, 5, 45, 13, 53, 21, 61, 29,
        36, 4, 44, 12, 52, 20, 60, 28, 35, 3, 43, 11, 51, 19, 59, 27,
        34, 2, 42, 10, 50, 18, 58, 26, 33, 1, 41, 9,  49, 17, 57, 25
    ];
    my $sc = [ 1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1 ];

    sub des_E_P16 {
        my ($p14) = @_;
        my $sp8 = [ 0x4b, 0x47, 0x53, 0x21, 0x40, 0x23, 0x24, 0x25 ];
        my $p7 = substr( $p14, 0, 7 );
        my $p16 = des_smbhash( $sp8, $p7 );
        $p7 = substr( $p14, 7, 7 );
        $p16 .= des_smbhash( $sp8, $p7 );
        return $p16;
    }

    sub des_E_P24 {
        my ( $p21, $c8_str ) = @_;
        my @c8 = map { ord($_) } split( //, $c8_str );
        my $p24 = des_smbhash( \@c8, substr( $p21, 0, 7 ) );
        $p24 .= des_smbhash( \@c8, substr( $p21, 7,  7 ) );
        $p24 .= des_smbhash( \@c8, substr( $p21, 14, 7 ) );
    }

    sub des_permute {
        my ( $i, $out, $in, $p, $n ) = ( 0, @_ );
        foreach $i ( 0 .. ( $n - 1 ) ) {
            $out->[$i] = $in->[ $p->[$i] - 1 ];
        }
    }

    sub des_lshift {
        my ( $c, $d, $count ) = @_;
        my ( @outc, @outd, $i, $x );
        while ( $count-- ) {
            push @$c, shift @$c;
            push @$d, shift @$d;
        }
    }

    my %dohash_cache;    # cache for key data; saves some cycles
    my %key_cache;       # another cache for key data

    sub des_cache_reset {
        %dohash_cache = ();
        %key_cache    = ();
    }

    sub des_dohash {
        my ( $out, $in, $key ) = @_;
        my ( $i, $j, $k, @pk1, @c, @d, @cd, @ki, @pd1, @l, @r, @rl );

        # if(!defined $dohash_cache{$skey}){
        &des_permute( \@pk1, $key, $perm1, 56 );

        for ( $i = 0 ; $i < 28 ; $i++ ) {
            $c[$i] = $pk1[$i];
            $d[$i] = $pk1[ $i + 28 ];
        }
        for ( $i = 0 ; $i < 16 ; $i++ ) {
            my @array;
            &des_lshift( \@c, \@d, $sc->[$i] );
            @cd = ( @c, @d );
            &des_permute( \@array, \@cd, $perm2, 48 );
            $ki[$i] = \@array;

            #    $dohash_cache{$skey}->[$i]=\@array;
        }

        # } else {
        #	for($i=0;$i<16;$i++){
        #		$ki[$i]=$dohash_cache{$skey}->[$i];}
        # }

        des_dohash2( $in, \@l, \@r, \@ki );

        @rl = ( @r, @l );
        &des_permute( $out, \@rl, $perm6, 64 );
    }

    sub des_str_to_key {
        my ($str) = @_;
        my ( $i, @key, $out, @str );
        unshift( @str, ord($_) ) while ( $_ = chop($str) );
        $key[0] = $str[0] >> 1;
        $key[1] = ( ( $str[0] & 0x01 ) << 6 ) | ( $str[1] >> 2 );
        $key[2] = ( ( $str[1] & 0x03 ) << 5 ) | ( $str[2] >> 3 );
        $key[3] = ( ( $str[2] & 0x07 ) << 4 ) | ( $str[3] >> 4 );
        $key[4] = ( ( $str[3] & 0x0f ) << 3 ) | ( $str[4] >> 5 );
        $key[5] = ( ( $str[4] & 0x1f ) << 2 ) | ( $str[5] >> 6 );
        $key[6] = ( ( $str[5] & 0x3f ) << 1 ) | ( $str[6] >> 7 );
        $key[7] = $str[6] & 0x7f;
        foreach $i ( 0 .. 7 ) {
            $key[$i] = 0xff & ( $key[$i] << 1 );
        }
        @{ $key_cache{$str} } = @key;
        return \@key;
    }

    sub des_smbhash {
        my ( $in, $key ) = @_;
        my $key2;

        &des_generate if ( !$generated );
        if ( defined $key_cache{$key} ) {
            $key2 = $key_cache{$key};
        }
        else { $key2 = &des_str_to_key($key); }

        my ( $i, $div, $mod, @in, @outb, @inb, @keyb, @out );
        foreach $i ( 0 .. 63 ) {
            $div = int( $i / 8 );
            $mod = $i % 8;
            $inb[$i]  = ( $in->[$div] &   ( 1 << ( 7 - ($mod) ) ) ) ? 1 : 0;
            $keyb[$i] = ( $key2->[$div] & ( 1 << ( 7 - ($mod) ) ) ) ? 1 : 0;
            $outb[$i] = 0;
        }
        &des_dohash( \@outb, \@inb, \@keyb );
        foreach $i ( 0 .. 7 ) { $out[$i] = 0; }
        foreach $i ( 0 .. 63 ) {
            $out[ int( $i / 8 ) ] |= ( 1 << ( 7 - ( $i % 8 ) ) )
              if ( $outb[$i] );
        }
        my $out = pack( "C8", @out );

        return $out;
    }

    sub des_generate {    # really scary dragons here....this code is optimized
                          # for speed, and not readability
        my ( $i, $j );
        my $code = <<EOT;
{ my \$sbox = [[
[14,4,13,1,2,15,11,8,3,10,6,12,5,9,0,7],[0,15,7,4,14,2,13,1,10,6,12,11,9,5,3,8],
[4,1,14,8,13,6,2,11,15,12,9,7,3,10,5,0],[15,12,8,2,4,9,1,7,5,11,3,14,10,0,6,13]
],[
[15,1,8,14,6,11,3,4,9,7,2,13,12,0,5,10],[3,13,4,7,15,2,8,14,12,0,1,10,6,9,11,5],
[0,14,7,11,10,4,13,1,5,8,12,6,9,3,2,15],[13,8,10,1,3,15,4,2,11,6,7,12,0,5,14,9]
],[
[10,0,9,14,6,3,15,5,1,13,12,7,11,4,2,8],[13,7,0,9,3,4,6,10,2,8,5,14,12,11,15,1],
[13,6,4,9,8,15,3,0,11,1,2,12,5,10,14,7],[1,10,13,0,6,9,8,7,4,15,14,3,11,5,2,12]
],[
[7,13,14,3,0,6,9,10,1,2,8,5,11,12,4,15],[13,8,11,5,6,15,0,3,4,7,2,12,1,10,14,9],
[10,6,9,0,12,11,7,13,15,1,3,14,5,2,8,4],[3,15,0,6,10,1,13,8,9,4,5,11,12,7,2,14]
],[
[2,12,4,1,7,10,11,6,8,5,3,15,13,0,14,9],[14,11,2,12,4,7,13,1,5,0,15,10,3,9,8,6],
[4,2,1,11,10,13,7,8,15,9,12,5,6,3,0,14],[11,8,12,7,1,14,2,13,6,15,0,9,10,4,5,3]
],[
[12,1,10,15,9,2,6,8,0,13,3,4,14,7,5,11],[10,15,4,2,7,12,9,5,6,1,13,14,0,11,3,8],
[9,14,15,5,2,8,12,3,7,0,4,10,1,13,11,6],[4,3,2,12,9,5,15,10,11,14,1,7,6,0,8,13]
],[
[4,11,2,14,15,0,8,13,3,12,9,7,5,10,6,1],[13,0,11,7,4,9,1,10,14,3,5,12,2,15,8,6],
[1,4,11,13,12,3,7,14,10,15,6,8,0,5,9,2],[6,11,13,8,1,4,10,7,9,5,0,15,14,2,3,12]
],[
[13,2,8,4,6,15,11,1,10,9,3,14,5,0,12,7],[1,15,13,8,10,3,7,4,12,5,6,11,0,14,9,2],
[7,11,4,1,9,12,14,2,0,6,10,13,15,3,5,8],[2,1,14,7,4,10,8,13,15,12,9,0,3,5,6,11]
]];
EOT

        $code .=
          'sub des_dohash2 { my ($in,$l,$r,$ki)=@_; my (@p,$i,$j,$k,$m,$n);';
        for ( $i = 0 ; $i < 64 ; $i++ ) {
            $code .= "\$p[$i] = \$in->[" . ( $perm3->[$i] - 1 ) . "];\n";
        }
        for ( $i = 0 ; $i < 32 ; $i++ ) {
            $code .= "\$l->[$i]=\$p[$i]; \$r->[$i]=\$p[" . ( $i + 32 ) . "];\n";
        }
        $code .= 'for($i=0;$i<16;$i++){ local (@er,@erk,@b,@cb,@pcb,@r2);';
        for ( $i = 0 ; $i < 48 ; $i++ ) {
            $code .=
                "\$erk[$i]=\$r->["
              . ( $perm4->[$i] - 1 )
              . "]^(\$ki->[\$i]->[$i]);\n";
        }
        for ( $i = 0 ; $i < 8 ; $i++ ) {
            for ( $j = 0 ; $j < 6 ; $j++ ) {
                $code .= "\$b[$i][$j]=\$erk[" . ( $i * 6 + $j ) . "];\n";
            }
        }
        for ( $i = 0 ; $i < 8 ; $i++ ) {
            $code .= "\$m=(\$b[$i][0]<<1)|\$b[$i][5];\n";
            $code .=
"\$n=(\$b[$i][1]<<3)|(\$b[$i][2]<<2)|(\$b[$i][3]<<1)|\$b[$i][4];\n";
            for ( $j = 0 ; $j < 4 ; $j++ ) {
                $code .=
                    "\$b[$i][$j]=(\$sbox->[$i][\$m][\$n]&"
                  . ( 1 << ( 3 - $j ) )
                  . ")?1:0;\n";
            }
        }
        for ( $i = 0 ; $i < 8 ; $i++ ) {
            for ( $j = 0 ; $j < 4 ; $j++ ) {
                $code .= "\$cb[" . ( $i * 4 + $j ) . "]=\$b[$i][$j];\n";
            }
        }
        for ( $i = 0 ; $i < 32 ; $i++ ) {
            $code .= "\$pcb[$i]=\$cb[" . ( $perm5->[$i] - 1 ) . "];\n";
        }
        for ( $i = 0 ; $i < 32 ; $i++ ) {
            $code .= "\$r2[$i]=(\$l->[$i])^\$pcb[$i];\n";
        }
        for ( $i = 0 ; $i < 32 ; $i++ ) {
            $code .= "\$l->[$i]=\$r->[$i]; \$r->[$i]=\$r2[$i];\n";
        }
        $code .= '}}}';

        eval "$code";
        print "DEBUG: $code\n\n";
        $generated++;
    }

}    ##### end of DES container ################################################

