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

=item B<utils_recperm>

Params: $uri, $depth, \@dir_parts, \@valid, \&func, \%track, \%arrays, \&cfunc

Return: nothing

This is a special function which is used to recursively-permutate through
a given directory listing.  This is really only used by whisker, in order
to traverse down directories, testing them as it goes.  See whisker 2.0 for
exact usage examples.

=cut

# '/', 0, \@dir.split, \@valid, \&func, \%track, \%arrays, \&cfunc
sub utils_recperm {
    my ( $d, $p, $pp, $pn, $r, $fr, $dr, $ar, $cr ) = ( '', shift, shift, @_ );
    $p =~ s#/+#/#g;
    if ( $pp >= @$pn ) {
        push @$r, $p if &$cr( $$dr{$p} );
    }
    else {
        my $c = $$pn[$pp];
        if ( $c !~ /^\@/ ) {
            utils_recperm( $p . $c . '/', $pp + 1, @_ )
              if ( &$fr( $p . $c . '/' ) );
        }
        else {
            $c =~ tr/\@//d;
            if ( defined $$ar{$c} ) {
                foreach $d ( @{ $$ar{$c} } ) {
                    if ( &$fr( $p . $d . '/' ) ) {
                        utils_recperm( $p . $d . '/', $pp + 1, @_ );
                    }
                }
            }
        }
    }
}

#################################################################

=item B<utils_array_shuffle>

Params: \@array

Return: nothing

This function will randomize the order of the elements in the given array.

=cut

sub utils_array_shuffle {    # fisher yates shuffle....w00p!
    my $array = shift;
    my $i;
    for ( $i = @$array ; --$i ; ) {
        my $j = int rand( $i + 1 );
        next if $i == $j;
        @$array[ $i, $j ] = @$array[ $j, $i ];
    }
}    # end array_shuffle, from Perl Cookbook (rock!)

#################################################################

=item B<utils_randstr>

Params: [ $size, $chars ]

Return: $random_string

This function generates a random string between 10 and 20 characters
long, or of $size if specified.  If $chars is specified, then the
random function picks characters from the supplied string.  For example,
to have a random string of 10 characters, composed of only the characters
'abcdef', then you would run:

 utils_randstr(10,'abcdef');

The default character string is alphanumeric.

=cut

sub utils_randstr {
    my $str;
    my $drift = shift || ( ( rand() * 10 ) % 10 ) + 10;

    # 'a'..'z' doesn't seem to work on string assignment :(
    my $CHARS = shift
      || 'abcdefghijklmnopqrstuvwxyz'
      . 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
      . '0123456789';

    my $L = length($CHARS);
    for ( 1 .. $drift ) {
        $str .= substr( $CHARS, ( ( rand() * $L ) % $L ), 1 );
    }
    return $str;
}

#################################################################

=item B<utils_port_open>

Params: $host, $port

Return: $result

Quick function to attempt to make a connection to the given host and
port.  If a connection was successfully made, function will return true
(1).  Otherwise it returns false (0).

Note: this uses standard TCP connections, thus is not recommended for use
in port-scanning type applications.  Extremely slow.

=cut

sub utils_port_open {    # this should be platform-safe
    my ( $target, $port ) = @_;

    return 0 if ( !defined $target || !defined $port );
    return 0 if ( !defined $Socket::VERSION );

    if ( !( socket( S, PF_INET, SOCK_STREAM, 0 ) ) ) { return 0; }
    if ( connect( S, sockaddr_in( $port, inet_aton($target) ) ) ) {
        close(S);
        return 1;
    }
    else { return 0; }
}

#################################################################

=item B<utils_lowercase_keys>

Params: \%hash

Return: $number_changed

Will lowercase all the header names (but not values) of the given hash.

=cut

sub utils_lowercase_keys {
    my $href = shift;

    return if ( !( defined $href && ref($href) ) );

    my $count = 0;
    while ( my ( $key, $val ) = each %$href ) {
        if ( $key =~ tr/A-Z// ) {
            $count++;
            delete $$href{$key};
            $$href{ lc($key) } = $val;
        }
    }
    return $count;
}

#################################################################

=item B<utils_find_lowercase_key>

Params: \%hash, $key

Return: $value, undef on error or not exist

Searches the given hash for the $key (regardless of case), and
returns the value. If the return value is placed into an array, the
will dereference any multi-value references and return an array of
all values.

WARNING!  In scalar context, $value can either be a single-value
scalar or an array reference for multiple scalar values.  That means
you either need to check the return value and act appropriately, or
use an array context (even if you only want a single value).  This is
very important, even if you know there are no multi-value hash keys.
This function may still return an array of multiple values even if
all hash keys are single value, since lowercasing the keys could result
in multiple keys matching.  For example, a hash with the values
{ 'Foo'=>'a', 'fOo'=>'b' } technically has two keys with the lowercase
name 'foo', and so this function will either return an array or array
reference with both 'a' and 'b'.

=cut

sub utils_find_lowercase_key {
    return utils_find_key( $_[0], $_[1], 1 );
}

#################################################################

=item B<utils_find_key>

Params: \%hash, $key

Return: $value, undef on error or not exist

Searches the given hash for the $key (case-sensitive), and
returns the value. If the return value is placed into an array, the
will dereference any multi-value references and return an array of
all values.

=cut

sub utils_find_key {
    my ( $href, $key, $dolower ) = ( shift, shift, shift || 0 );

    return undef if ( !( defined $href && ref($href) ) );
    return undef if ( !defined $key );

    if ($dolower) {
        $key = lc($key);
        my ( $k, $v );
				my @match;
        while ( ( $k, $v ) = each %$href ) {
            if ( lc($k) eq $key ) {
                if( ref($v) ) {
                    push @match, @$v;
                } else {
                    push @match, $v;
                }
            }
        }
        return @match if wantarray();
        return \@match if( ~~@match > 1 );
        return $match[0];
    }
    else {
        return @{ $href->{$key} } if ( ref( $href->{$key} ) && wantarray() );
        return $href->{$key};
    }
    return undef;
}

#################################################################

=item B<utils_delete_lowercase_key>

Params: \%hash, $key

Return: $number_found

Searches the given hash for the $key (regardless of case), and
deletes the key out of the hash if found.  The function returns
the number of keys found and deleted (since multiple keys can
exist under the names 'Key', 'key', 'keY', 'KEY', etc.).

=cut

sub utils_delete_lowercase_key {
    my ( $href, $key ) = ( shift, lc(shift) );

    return undef if ( !( defined $href && ref($href) ) );
    return undef if ( !defined $key );

    my $deleted = 0;
    foreach ( keys %$href ) {
        if ( lc($_) eq $key ) {
            delete $href->{$_};
            $deleted++;
        }
    }
    return $deleted;
}

#################################################################

=item B<utils_getline>

Params: \$data [, $resetpos ]

Return: $line (undef if no more data)

Fetches the next \n terminated line from the given data.  Use
the optional $resetpos to reset the internal position pointer.
Does *NOT* return trialing \n.

=cut

{
    my $POS = 0;

    sub utils_getline {
        my ( $dr, $rp ) = @_;

        return undef if ( !( defined $dr && ref($dr) ) );
        $POS = $rp if ( defined $rp );

        my $where = index( $$dr, "\x0a", $POS );
        return undef if ( $where == -1 );

        my $str = substr( $$dr, $POS, $where - $POS );
        $POS = $where + 1;

        return $str;
    }
}

#################################################################

=item B<utils_getline_crlf>

Params: \$data [, $resetpos ]

Return: $line (undef if no more data)

Fetches the next \r\n terminated line from the given data.  Use
the optional $resetpos to reset the internal position pointer.
Does *NOT* return trialing \r\n.

=cut

{
    my $POS = 0;

    sub utils_getline_crlf {
        my ( $dr, $rp ) = @_;

        return undef if ( !( defined $dr && ref($dr) ) );
        $POS = $rp if ( defined $rp );

        my $tpos = $POS;
        while (1) {
            my $where = index( $$dr, "\x0a", $tpos );
            return undef if ( $where == -1 );

            if ( substr( $$dr, $where - 1, 1 ) eq "\x0d" ) {
                my $str = substr( $$dr, $POS, $where - $POS - 1 );
                $POS = $where + 1;
                return $str;
            }
            else {
                $tpos = $where + 1;
            }
        }
    }
}

#################################################################

=item B<utils_save_page>

Params: $file, \%response

Return: 0 on success, 1 on error

Saves the data portion of the given whisker %response hash to the
indicated file.  Can technically save the data portion of a
%request hash too.  A file is not written if there is no data.

Note: LW does not do any special file checking; files are opened
in overwrite mode.

=cut

sub utils_save_page {
    my ( $file, $hr ) = @_;
    return 1 if ( !ref($hr) || ref($file) );
    return 0
      if ( !defined $$hr{'whisker'}
        || !defined $$hr{'whisker'}->{'data'} );
    open( OUT, ">$file" ) || return 1;
    binmode(OUT); # Stupid Windows
    print OUT $$hr{'whisker'}->{'data'};
    close(OUT);
    return 0;
}

#################################################################

=item B<utils_getopts>

Params: $opt_str, \%opt_results

Return: 0 on success, 1 on error

This function is a general implementation of GetOpts::Std.  It will
parse @ARGV, looking for the options specified in $opt_str, and will
put the results in %opt_results.  Behavior/parameter values are
similar to GetOpts::Std's getopts().

Note: this function does *not* support long options (--option),
option grouping (-opq), or options with immediate values (-ovalue).
If an option is indicated as having a value, it will take the next
argument regardless.

=cut

sub utils_getopts {
    my ( $str, $ref ) = @_;
    my ( %O, $l );
    my @left;

    return 1 if ( $str =~ tr/-:a-zA-Z0-9//c );

    while ( $str =~ m/([a-z0-9]:{0,1})/ig ) {
        $l = $1;
        if ( $l =~ tr/://d ) {
            $O{$l} = 1;
        }
        else { $O{$l} = 0; }
    }

    while ( $l = shift(@ARGV) ) {
        push( @left, $l ) && next if ( substr( $l, 0, 1 ) ne '-' );
        push( @left, $l ) && next if ( $l eq '-' );
        substr( $l, 0, 1 ) = '';
        if ( length($l) != 1 ) {
            %$ref = ();
            return 1;
        }
        if ( $O{$l} == 1 ) {
            my $x = shift(@ARGV);
            $$ref{$l} = $x;
        }
        else { $$ref{$l} = 1; }
    }

    @ARGV = @left;
    return 0;
}

#################################################################

=item B<utils_text_wrapper>

Params: $long_text_string [, $crlf, $width ]

Return: $formatted_test_string

This is a simple function used to format a long line of text for
display on a typical limited-character screen, such as a unix
shell console.

$crlf defaults to "\n", and $width defaults to 76.

=cut

sub utils_text_wrapper {
    my ( $out, $w, $str, $crlf, $width ) = ( '', 0, @_ );
    $crlf  ||= "\n";
    $width ||= 76;
    $str .= $crlf if ( $str !~ /$crlf$/ );
    return $str if ( length($str) <= $width );
    while ( length($str) > $width ) {
        my $w1 = rindex( $str, ' ',  $width );
        my $w2 = rindex( $str, "\t", $width );
        if ( $w1 > $w2 ) { $w = $w1; }
        else { $w = $w2; }
        if ( $w == -1 ) {
            $w = $width;
        }
        else { substr( $str, $w, 1 ) = ''; }
        $out .= substr( $str, 0, $w, '' );
        $out .= $crlf;
    }
    return $out . $str;
}

#################################################################

=item B<utils_bruteurl>

Params: \%req, $pre, $post, \@values_in, \@values_out

Return: Nothing (adds to @out)
        
Bruteurl will perform a brute force against the host/server specified in
%req.  However, it will make one request per entry in @in, taking the
value and setting $hin{'whisker'}->{'uri'}= $pre.value.$post.  Any URI
responding with an HTTP 200 or 403 response is pushed into @out.  An
example of this would be to brute force usernames, putting a list of
common usernames in @in, setting $pre='/~' and $post='/'.

=cut

sub utils_bruteurl {
    my ( $hin, $upre, $upost, $arin, $arout ) = @_;
    my ( $U, %hout );

    return if ( !( defined $hin   && ref($hin) ) );
    return if ( !( defined $arin  && ref($arin) ) );
    return if ( !( defined $arout && ref($arout) ) );
    return if ( !defined $upre  || length($upre) == 0 );
    return if ( !defined $upost || length($upost) == 0 );

    http_fixup_request($hin);

    map {
        ( $U = $_ ) =~ tr/\r\n//d;
        next if ( $U eq '' );
        if (
            !http_do_request( $hin, \%hout, { 'uri' => $upre . $U . $upost } ) )
        {
            if (   $hout{'whisker'}->{'code'} == 200
                || $hout{'whisker'}->{'code'} == 403 )
            {
                push( @{$arout}, $U );
            }
        }
    } @$arin;
}

#################################################################

=item B<utils_join_tag>

Params: $tag_name, \%attributes

Return: $tag_string [undef on error]
        
This function takes the $tag_name (like 'A') and a hash full of
attributes (like {href=>'http://foo/'}) and returns the 
constructed HTML tag string (<A href="http://foo">).

=cut

sub utils_join_tag {
    my ( $name, $href ) = @_;
    return undef if ( !defined $name || $name eq '' );
    return undef if ( !defined $href || !ref($href) );
    my ( $out, $k, $v ) = ( "<$name", '', '' );
    while ( ( $k, $v ) = each %$href ) {
        next if ( $k eq '' );
        $out .= " $k";
        $out .= "=\"$v\"" if ( defined $v );
    }
    $out .= '>';
    return $out;
}

#################################################################

=item B<utils_request_clone>

Params: \%from_request, \%to_request

Return: 1 on success, 0 on error

This function takes the connection/request-specific values from the
given from_request hash, and copies them to the to_request hash.

=cut

sub utils_request_clone {
    my ( $from, $to ) = @_;
    return 0 if ( !defined $from || !ref($from) );
    return 0 if ( !defined $to   || !ref($to) );
    return 0 if ( !defined $from->{whisker}->{MAGIC} );

    %$to = ();

    # copy headers
    my ( $k, $v );
    while ( ( $k, $v ) = each(%$from) ) {
        next if ( $k eq 'whisker' );
        if ( ref($v) ) {
            @{ $to->{$k} } = @$v;
        }
        else {
            $to->{$k} = $v;
        }
    }

    # copy whisker control values
    $to->{whisker} = {};
    while ( ( $k, $v ) = each( %{ $from->{whisker} } ) ) {
        if ( ref($v) ) {
            @{ $to->{whisker}->{$k} } = @$v;
        }
        else {
            $to->{whisker}->{$k} = $v;
        }
    }

    return 1;
}

#################################################################

=item B<utils_request_fingerprint>

Params: \%request [, $hash ]

Return: $fingerprint [undef on error]
        
This function constructs a 'fingerprint' of the given request by
using a cryptographic hashing function on the constructed original
HTTP request.

Note: $hash can be 'md5' (default) or 'md4'.

=cut

sub utils_request_fingerprint {
    my ( $href, $hash ) = @_;
    $hash ||= 'md5';
    return undef if ( !defined $href || !ref($href) );
    return undef if ( !defined $href->{whisker}->{MAGIC} );

    my $data = '';
    if ( $href->{whisker}->{MAGIC} == 31339 ) {    # LW2 request
        $data = http_req2line($href);
        if ( $href->{whisker}->{version} ne '0.9' ) {
            $data .= http_construct_headers($href);
            $data .= $href->{whisker}->{raw_header_data}
              if ( defined $href->{whisker}->{raw_header_data} );
            $data .= $href->{whisker}->{http_eol};
            $data .= $href->{whisker}->{data}
              if ( defined $href->{whisker}->{data} );
        }                                          # http 0.9 support

        return 'md5:' . md5($data) if ( $hash eq 'md5' );
        return 'md4:' . md4($data) if ( $hash eq 'md4' );
    }

    return undef;
}

#################################################################

=item B<utils_flatten_lwhash>

Params: \%lwhash

Return: $flat_version [undef on error]
        
This function takes a %request or %response libwhisker hash, and
creates an approximate flat data string of the original request/
response (i.e. before it was parsed into components and placed
into the libwhisker hash).

=cut

sub utils_flatten_lwhash {
    my $hr = shift;
    return undef if ( !defined $hr || !ref($hr) );
    my $out;

    if ( $hr->{whisker}->{MAGIC} == 31339 ) {
        $out = http_req2line($hr);
    }
    elsif ( $hr->{whisker}->{MAGIC} == 31340 ) {
        $out = http_resp2line($hr);
    }
    else {
        return undef;
    }

    $out .= http_construct_headers($hr);
    $out .= $hr->{whisker}->{http_eol} || "\x0d\x0a";
    if ( defined $hr->{whisker}->{data}
        && length( $hr->{whisker}->{data} ) > 0 )
    {
        $out .= $hr->{whisker}->{data};
    }

    return $out;
}

#################################################################

sub _utils_carp_common {
	my ($x,$pack,$m) = (0, shift || '',join('',@_) || '(Unknown error)');
	my @s = caller($x++);
	@s=caller($x++) while(defined $s[0] && ($s[0] eq 'LW2' || $s[0] eq $pack));
	return $m if !defined $s[0];
	return "$m at $s[1] line $s[2]\n";
}

=item B<utils_carp>

Params: [ $package_name ]

Return: nothing
        
This function acts like Carp's carp function.  It warn's with the file and 
line number of user's code which causes a problem.  It traces up the call 
stack and reports the first function that is not in the LW2 or optional 
$package_name package package.

=cut

sub utils_carp {
	warn _utils_carp_common(@_);
}

=item B<utils_croak>

Params: [ $package_name ]

Return: nothing
        
This function acts like Carp's croak function.  It die's with the file and 
line number of user's code which causes a problem.  It traces up the call 
stack and reports the first function that is not in the LW2 or optional 
$package_name package package.

=cut

sub utils_croak {
	die _utils_carp_common(@_);
}
