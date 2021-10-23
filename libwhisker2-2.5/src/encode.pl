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

{    # package variables
    my $MIMEBASE64_TRYLOADING = 1;

########################################################################

=item B<encode_base64>

Params: $data [, $eol]

Return: $b64_encoded_data

This function does Base64 encoding.  If the binary MIME::Base64 module
is available, it will use that; otherwise, it falls back to an internal
perl version.  The perl version carries the following copyright:

 Copyright 1995-1999 Gisle Aas <gisle@aas.no>

NOTE: the $eol parameter will be inserted every 76 characters.  This is
used to format the data for output on a 80 character wide terminal.

=cut

    sub encode_base64 {
        if ($MIMEBASE64_TRYLOADING) {
            eval "require MIME::Base64";
            $MIMEBASE64_TRYLOADING = 0;
        }
        goto &MIME::Base64::encode_base64 if ($MIME::Base64::VERSION);
        my $res = "";
        my $eol = $_[1];
        $eol = "\n" unless defined $eol;
        pos( $_[0] ) = 0;
        while ( $_[0] =~ /(.{1,45})/gs ) {
            $res .= substr( pack( 'u', $1 ), 1 );
            chop($res);
        }
        $res =~ tr|` -_|AA-Za-z0-9+/|;
        my $padding = ( 3 - length( $_[0] ) % 3 ) % 3;
        $res =~ s/.{$padding}$/'=' x $padding/e if $padding;
        if ( length $eol ) {
            $res =~ s/(.{1,76})/$1$eol/g;
        }
        $res;
    }

########################################################################

=item B<decode_base64>

Params: $data

Return: $b64_decoded_data

A perl implementation of base64 decoding.  The perl code for this function
was actually taken from an older MIME::Base64 perl module, and bears the 
following copyright:

Copyright 1995-1999 Gisle Aas <gisle@aas.no>

=cut

    sub decode_base64 {
        if ($MIMEBASE64_TRYLOADING) {
            eval "require MIME::Base64";
            $MIMEBASE64_TRYLOADING = 0;
        }
        goto &MIME::Base64::decode_base64 if ($MIME::Base64::VERSION);
        my $str = shift;
        my $res = "";
        $str =~ tr|A-Za-z0-9+=/||cd;
        $str =~ s/=+$//;                # remove padding
        $str =~ tr|A-Za-z0-9+/| -_|;    # convert to uuencoded format
        while ( $str =~ /(.{1,60})/gs ) {
            my $len = chr( 32 + length($1) * 3 / 4 );    # compute length byte
            $res .= unpack( "u", $len . $1 );            # uudecode
        }
        $res;
    }

########################################################################

}    # end package variables

########################################################################

=item B<encode_uri_hex>

Params: $data

Return: $result

This function encodes every character (except the / character) with normal 
URL hex encoding.

=cut

sub encode_uri_hex {    # normal hex encoding
    my $str = shift;
    $str =~ s/([^\/])/sprintf("%%%02x",ord($1))/ge;
    return $str;
}

#########################################################################

=item B<encode_uri_randomhex>

Params: $data

Return: $result

This function randomly encodes characters (except the / character) with 
normal URL hex encoding.

=cut

sub encode_uri_randomhex {    # random normal hex encoding
    my @T = split( //, shift );
    my $s;
    foreach (@T) {
        if (m#[;=:&@\?]#) {
            $s .= $_;
            next;
        }
        if ( ( rand() * 2 ) % 2 == 1 ) { $s .= sprintf( "%%%02x", ord($_) ); }
        else { $s .= $_; }
    }
    return $s;
}

#########################################################################

=item B<encode_uri_randomcase>

Params: $data

Return: $result

This function randomly changes the case of characters in the string.

=cut

sub encode_uri_randomcase {
    my ( $x, $uri ) = ( '', shift );
    return $uri if ( $uri !~ tr/a-zA-Z// );    # fast-path
    my @T = split( //, $uri );
    for ( $x = 0 ; $x < ( scalar @T ) ; $x++ ) {
        if ( ( rand() * 2 ) % 2 == 1 ) {
            $T[$x] =~ tr/A-Za-z/a-zA-Z/;
        }
    }
    return join( '', @T );
}

#########################################################################

=item B<encode_unicode>

Params: $data

Return: $result

This function converts a normal string into Windows unicode format
(non-overlong or anything fancy).

=cut

sub encode_unicode {
    my ( $c, $r ) = ( '', '' );
    foreach $c ( split( //, shift ) ) {
        $r .= pack( "v", ord($c) );
    }
    return $r;
}

#########################################################################

=item B<decode_unicode>

Params: $unicode_string

Return: $decoded_string

This function attempts to decode a unicode (UTF-8) string by
converting it into a single-byte-character string.  Overlong 
characters are converted to their standard characters in place; 
non-overlong (aka multi-byte) characters are substituted with the 
0xff; invalid encoding characters are left as-is.

Note: this function is useful for dealing with the various unicode
exploits/vulnerabilities found in web servers; it is *not* good for
doing actual UTF-8 parsing, since characters over a single byte are
basically dropped/replaced with a placeholder.

=cut

sub decode_unicode {
    my $str = $_[0];
    return $str if ( $str !~ tr/!-~//c );    # fastpath
    my ( $lead, $count, $idx );
    my $out = '';
    my $len = length($str);
    my ( $ptr, $no, $nu ) = ( 0, 0, 0 );

    while ( $ptr < $len ) {
        my $c = substr( $str, $ptr, 1 );
        if ( ord($c) >= 0xc0 && ord($c) <= 0xfd ) {
            $count = 0;
            $c     = ord($c) << 1;
            while ( ( $c & 0x80 ) == 0x80 ) {
                $c <<= 1;
                last if ( $count++ == 4 );
            }
            $c = ( $c & 0xff );
            for ( $idx = 1 ; $idx < $count ; $idx++ ) {
                my $o = ord( substr( $str, $ptr + $idx, 1 ) );
                $no = 1 if ( $o != 0x80 );
                $nu = 1 if ( $o < 0x80 || $o > 0xbf );
            }
            my $o = ord( substr( $str, $ptr + $idx, 1 ) );
            $nu = 1 if ( $o < 0x80 || $o > 0xbf );
            if ($nu) {
                $out .= substr( $str, $ptr++, 1 );
            }
            else {
                if ($no) {
                    $out .= "\xff";    # generic replacement char
                }
                else {
                    my $prior =
                      ord( substr( $str, $ptr + $count - 1, 1 ) ) << 6;
                    $out .= pack( "C",
                        (( ord( substr( $str, $ptr + $count, 1 ) ) & 0x7f ) +
                          $prior ) & 255 );
                }
                $ptr += $count + 1;
            }
            $no = $nu = 0;
        }
        else {
            $out .= $c;
            $ptr++;
        }
    }
    return $out;
}

########################################################################

=item B<encode_anti_ids>

Params: \%request, $modes

Return: nothing

encode_anti_ids computes the proper anti-ids encoding/tricks 
specified by $modes, and sets up %hin in order to use those tricks.  
Valid modes are (the mode numbers are the same as those found in whisker 
1.4):

=over 4

=item 1 Encode some of the characters via normal URL encoding

=item 2 Insert directory self-references (/./)

=item 3 Premature URL ending (make it appear the request line is done)

=item 4 Prepend a long random string in the form of "/string/../URL"

=item 5 Add a fake URL parameter

=item 6 Use a tab instead of a space as a request spacer

=item 7 Change the case of the URL (works against Windows and Novell)

=item 8 Change normal seperators ('/') to Windows version ('\')

=item 9 Session splicing [NOTE: not currently available]

=item A Use a carriage return (0x0d) as a request spacer

=item B Use binary value 0x0b as a request spacer

=back

You can set multiple modes by setting the string to contain all the modes
desired; i.e. $modes="146" will use modes 1, 4, and 6.

=cut

sub encode_anti_ids {
    my ( $rhin, $modes ) = ( shift, shift );
    my ( @T, $x, $c, $s, $y );
    my $ENCODED = 0;
    my $W       = $$rhin{'whisker'};

    return if ( !( defined $rhin && ref($rhin) ) );

    # in case they didn't do it already
    $$rhin{'whisker'}->{'uri_orig'} = $$rhin{'whisker'}->{'uri'};

    # note: order is important!

    # mode 9 - session splicing
    #if($modes=~/9/){
    #	$$rhin{'whisker'}->{'ids_session_splice'}=1;
    #}

    # mode 4 - prepend long random string
    if ( $modes =~ /4/ ) {
        $s = '';
        if ( $$W{'uri'} =~ m#^/# ) {
            $y = &utils_randstr;
            $s .= $y while ( length($s) < 512 );
            $$W{'uri'} = "/$s/.." . $$W{'uri'};
        }
    }

    # mode 7  - (windows) random case sensitivity
    if ( $modes =~ /7/ ) {
        $$W{'uri'} = encode_uri_randomcase( $$W{'uri'} );
    }

    # mode 2 - directory self-reference (/./)
    if ( $modes =~ /2/ ) {
        $$W{'uri'} =~ s#/#/./#g;
    }

    # mode 8 - windows directory separator (\)
    if ( $modes =~ /8/ ) {
        $$W{'uri'} =~ s#/#\\#g;
        $$W{'uri'} =~ s#^\\#/#;
        $$W{'uri'} =~ s#^([a-zA-Z0-9_]+):\\#$1://#;
        $$W{'uri'} =~ s#\\$#/#;
    }

    # mode 1 - random URI (non-UTF8) encoding
    if ( $modes =~ /1/ ) {
        if ( $ENCODED == 0 ) {
            $$W{'uri'} = encode_uri_randomhex( $$W{'uri'} );
            $ENCODED = 1;
        }
    }

    # mode 5 - fake parameter
    if ( $modes =~ /5/ ) {
        ( $s, $y ) = ( &utils_randstr, &utils_randstr );
        $$W{'uri'} = "/$s.html%3F$y=/../$$W{'uri'}";
    }

    # mode 3 - premature URL ending
    if ( $modes =~ /3/ ) {
        $s = &utils_randstr;
        $$W{'uri'} = "/%20HTTP/1.1%0d%0aAccept%3a%20$s/../..$$W{'uri'}";
    }

    # mode 6 - TAB as request spacer
    if ( $modes =~ /6/ ) {
        $$W{'http_space1'} = "\t";
    }

    # mode A - CR as request spacer
    if ( $modes =~ /A/i ) {
        $$W{'http_space1'} = $$W{'http_space2'} = "\x0d";
    }

    # mode B - 0x0b as request spacer
    if ( $modes =~ /B/i ) {
        $$W{'http_space1'} = $$W{'http_space2'} = "\x0b";
    }

}

