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

#################################################################

=item B<uri_split>

Params: $uri_string [, \%request_hash]

Return: @uri_parts

Return an array of the following values, in order:  uri, protocol, host,
port, params, frag, user, password.  Values not defined are given an undef
value.  If a %request hash is passed in, then uri_split() will also set 
the appropriate values in the hash.

Note:  uri_split() will only set the %request hash if the protocol
is HTTP or HTTPS!

=cut

sub uri_split {
    my ( $uri, $work ) = ( shift, '', 0 );
    my ($hr) = shift;

    #       (uri,protocol,host,port,params,frag,user,pass)
    my @res = ( undef, undef, undef, 0, undef, undef, undef, undef );

    return undef if ( !defined $uri );

    # remove fragments
    ( $uri, $res[5] ) = split( '#', $uri, 2 ) if ( index( $uri, '#', 0 ) >= 0 );

    # get scheme and net_loc
    my $net_loc = undef;
    if ( $uri =~ s/^([-+.a-z0-9A-Z]+):// ) {
        $res[1] = lc($1);
        if ( substr( $uri, 0, 2 ) eq '//' ) {
            my $w = index( $uri, '/', 2 );
            if ( $w >= 0 ) {
                $net_loc = substr( $uri, 2, $w - 2 );
                $uri = substr( $uri, $w, length($uri) - $w );
            }
            else {
                ( $net_loc = $uri ) =~ tr#/##d;
                $uri = '/';
            }
        }
    }

    # parse net_loc info
    if ( defined $net_loc ) {
        if ( index( $net_loc, '@', 0 ) >= 0 ) {
            ( $res[6], $net_loc ) = split( /\@/, $net_loc, 2 );
            if ( index( $res[6], ':', 0 ) >= 0 ) {
                ( $res[6], $res[7] ) = split( ':', $res[6], 2 );
            }
        }
        $res[3] = $1 if ( $net_loc =~ s/:([0-9]+)$// );
        $res[2] = $net_loc;
    }

    # remove query info
    ( $uri, $res[4] ) = split( '\?', $uri, 2 )
      if ( index( $uri, '?', 0 ) >= 0 );

    # whatever is left over is the uri
    $res[0] = $uri;

    if ( $res[3] == 0 && defined $res[1] ) {
        $res[3] = 80  if ( $res[1] eq 'http' );
        $res[3] = 443 if ( $res[1] eq 'https' );
    }

    my $rel_uri = 0;
    $rel_uri++
      if ( $res[3] == 0
        && !defined $res[2]
        && !defined $res[1]
        && $res[0] ne '' );
    return @res if ( $res[3] == 0 && !$rel_uri );

    if ( defined $hr && ref($hr) ) {

        $$hr{whisker}->{uri} = $res[0] if ( defined $res[0] );
        if ( defined $res[4] ) {
            $$hr{whisker}->{parameters} = $res[4];
        }
        else { delete $$hr{whisker}->{parameters}; }

        return @res if ($rel_uri);

				if ( $res[1] eq 'https' ) {
	        $$hr{whisker}->{ssl} = 1;
	      } else { $$hr{whisker}->{ssl} = 0; }
        $$hr{whisker}->{host} = $res[2] if ( defined $res[2] );
        $$hr{whisker}->{port} = $res[3];

        if ( defined $res[6] ) {
            $$hr{whisker}->{uri_user} = $res[6];
        }
        else { delete $$hr{whisker}->{uri_user}; }
        if ( defined $res[7] ) {
            $$hr{whisker}->{uri_password} = $res[7];
        }
        else { delete $$hr{whisker}->{uri_password}; }
    }

    return @res;
}

#################################################################

=item B<uri_join>

Params: @vals

Return: $url

Takes the @vals array output from http_split_uri, and returns a single 
scalar/string with them joined again, in the form of:
protocol://user:pass@host:port/uri?params#frag

=cut

sub uri_join {
    my @V = @_;
    my $URL;

    $URL .= $V[1] . ':' if defined $V[1];
    if ( defined $V[2] ) {
        $URL .= '//';
        if ( defined $V[6] ) {
            $URL .= $V[6];
            $URL .= ':' . $V[7] if defined $V[7];
            $URL .= '@';
        }
        $URL .= $V[2];
    }

    if ( $V[3] > 0 ) {
        my $no = 0;
        $no++ if ( $V[3] == 80  && defined $V[1] && $V[1] eq 'http' );
        $no++ if ( $V[3] == 443 && defined $V[1] && $V[1] eq 'https' );
        $URL .= ':' . $V[3] if ( !$no );
    }

    $URL .= $V[0];
    $URL .= '?' . $V[4] if defined $V[4];
    $URL .= '#' . $V[5] if defined $V[5];
    return $URL;
}

#################################################################

=item B<uri_absolute>

Params: $uri, $base_uri [, $normalize_flag ]

Return: $absolute_uri

Double checks that the given $uri is in absolute form (that is,
"http://host/file"), and if not (it's in the form "/file"), then
it will append the given $base_uri to make it absolute.  This
provides a compatibility similar to that found in the URI
subpackage.

If $normalize_flag is set to 1, then the output will be passed
through uri_normalize before being returned.

=cut

sub uri_absolute {
    my ( $uri, $buri, $norm ) = @_;
    return undef if ( !defined $uri || !defined $buri );

    return $uri if ( $uri =~ m#^[-+.a-z0-9A-Z]+://# );

    if ( substr( $uri, 0, 1 ) eq '/' ) {
        if ( $buri =~ m#^[-+.a-z0-9A-Z]+://# ) {
            my @p = uri_split($buri);
            $buri = "$p[1]://$p[2]";
            $buri .= ":$p[3]" if ( ($p[1] eq 'http' && $p[3] != 80) ||
            	($p[1] eq 'https' && $p[3] != 443) );

            #			$buri.='/';
        }
        else {    # ah suck, base URI isn't absolute...
            return $uri;
        }
    }
    else {
        $buri =~ s/[?#].*$//;    # remove params and fragments
        $buri .= '/' if ( $buri =~ m#^[a-z]+://[^/]+$#i );
        $buri =~ s#/[^/]*$#/#;
    }

    return uri_normalize("$buri$uri")
      if ( defined $norm && $norm > 0 );
    return $buri . $uri;
}

#################################################################

=item B<uri_normalize>

Params: $uri [, $fix_windows_slashes ]

Return: $normalized_uri [ undef on error ]

Takes the given $uri and does any /./ and /../ dereferencing in
order to come up with the correct absolute URL.  If the $fix_
windows_slashes parameter is set to 1, all \ (back slashes) will
be converted to / (forward slashes).

Non-http/https URIs return an error.

=cut

sub uri_normalize {
    my ( $host, $uri, $win ) = ( '', @_ );

    $uri =~ tr#\\#/# if ( defined $win && $win > 0 );

    if ( $uri =~ s#^([-+.a-z0-9A-Z]+:)## ) {
        return undef if ( $1 ne 'http:' && $1 ne 'https:' );
        $host = $1;
        return undef unless ( $uri =~ s#^(//[^/]+)## );
        $host .= $1;
    }
    return "$host/" if ( $uri eq '' || $uri eq '/' );

    # fast path check
    return "$host$uri" if ( index( $uri, '/.' ) < 0 );

    my $extra = '';
    $extra = $1 if($uri =~ s/([?#].*)$//);    # remove params and fragments

    # parse order/steps as defined in RFC 1808
    1 while ( $uri =~ s#/\./#/# || $uri =~ s#//#/# );
    $uri =~ s#/\.$#/#;
    1 while ( $uri =~ s#[^/]+/\.\./## );
    1 while ( $uri =~ s#^/\.\./#/# );
    $uri =~ s#[^/]*/\.\.$##;
    $uri ||= '/';
    return $host . $uri . $extra;
}

#################################################################

=item B<uri_get_dir>

Params: $uri

Return: $uri_directory

Will take a URI and return the directory base of it, i.e. /rfp/page.php 
will return /rfp/.

=cut

sub uri_get_dir {
    my ( $w, $URL ) = ( 0, shift );

    return undef if ( !defined $URL );
    $URL = substr( $URL, 0, $w ) if ( ( $w = index( $URL, '#' ) ) >= 0 );
    $URL = substr( $URL, 0, $w ) if ( ( $w = index( $URL, '?' ) ) >= 0 );
    return $URL if ( substr( $URL, -1, 1 ) eq '/' );

    if ( ( $w = rindex( $URL, '/' ) ) >= 0 ) {
        $URL = substr( $URL, 0, $w + 1 );
    }
    else {
        $URL = '';
    }
    return $URL;
}

#################################################################

=item B<uri_strip_path_parameters>

Params: $uri [, \%param_hash]

Return: $stripped_uri

This function removes all URI path parameters of the form

 /blah1;foo=bar/blah2;baz

and returns the stripped URI ('/blah1/blah2').  If the optional
parameter hash reference is provided, the stripped parameters
are saved in the form of 'blah1'=>'foo=bar', 'blah2'=>'baz'.

Note: only the last value of a duplicate name is saved into the 
param_hash, if provided.  So a $uri of '/foo;A/foo;B/' will result 
in a single hash entry of 'foo'=>'B'.

=cut

sub uri_strip_path_parameters {
    my ( $uri, $hr ) = @_;
    my $s   = 0;
    $s++ if ( defined $hr && ref($hr) );

    my @p = split( /\//, $uri, -1 );
    map {
        if (s/;(.*)$//) { $$hr{$_} = $1 if ($s); }
    } @p;

		return join( '/', @p );
}

#################################################################

=item B<uri_parse_parameters>

Params: $parameter_string [, $decode, $multi_flag ]

Return: \%parameter_hash

This function takes a string in the form of:

 foo=1&bar=2&baz=3&foo=4

And parses it into a hash.  In the above example, the element 'foo'
has two values (1 and 4).  If $multi_flag is set to 1, then the
'foo' hash entry will hold an anonymous array of both values. 
Otherwise, the default is to just contain the last value (in this
case, '4').

If $decode is set to 1, then normal hex decoding is done on the
characters, where needed (both the name and value are decoded).

Note: if a URL parameter name appears without a value, then the
value will be set to undef.  E.g. for the string "foo=1&bar&baz=2",
the 'bar' hash element will have an undef value.

=cut

sub uri_parse_parameters {
    my ( $str, $decode, $multi ) = @_;
    my %P;
    if( $str !~ tr/=&// ){
    	$P{$str} = undef;
    	return \%P;
    }

    $multi  ||= 0;
    $decode ||= 0;
    foreach ( split( /&/, $str ) ) {
        my ( $name, $value ) = split( /=/, $_, 2 );
        if ($decode) {
            $name  = uri_unescape($name);
            $value = uri_unescape($value);
        }
        if ( defined $P{$name} && $multi ) {
            if ( ref( $P{$name} ) ) { push @{ $P{$name} }, $value; }
            else { $P{$name} = [ $P{$name}, $value ]; }
        }
        else {
            $P{$name} = $value;
        }
    }
    return \%P;
}

#################################################################

=item B<uri_escape>

Params: $data

Return: $encoded_data

This function encodes the given $data so it is safe to be used in URIs.

=cut

sub uri_escape {
    my $data = shift;
    return undef if ( !defined $data );
    $data =~ s/\%/\%25/g;
    $data =~ s/([+?&=#;@\\\/])/sprintf("%%%02x",ord($1))/eg;
    $data =~ tr/ /+/;
    $data =~ s/([^!-~])/sprintf("%%%02x",ord($1))/eg;
    return $data;
}

#################################################################

=item B<uri_unescape>

Params: $encoded_data

Return: $data

This function decodes the given $data out of URI format.

=cut

sub uri_unescape {
    my $data = shift;
    return undef if ( !defined $data );
    $data =~ tr/+/ /;
    $data =~ s/%([a-fA-F0-9][a-fA-F0-9])/pack("C",hex($1))/eg;
    return $data;
}

#################################################################

