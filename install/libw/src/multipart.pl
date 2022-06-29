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

=item B<multipart_set>

Params: \%multi_hash, $param_name, $param_value

Return: nothing

This function sets the named parameter to the given value within the
supplied multipart hash.

=cut

sub multipart_set {
    my ( $hr, $n, $v ) = @_;
    return if ( !ref($hr) );    # error check
    return undef if ( !defined $n || $n eq '' );
    $$hr{$n} = $v;
}

########################################################################

=item B<multipart_get>

Params: \%multi_hash, $param_name

Return: $param_value, undef on error

This function retrieves the named parameter to the given value within the
supplied multipart hash.  There is a special case where the named
parameter is actually a file--in which case the resulting value will be
"\0FILE".  In general, all special values will be prefixed with a NULL
character.  In order to get a file's info, use multipart_getfile().

=cut

sub multipart_get {
    my ( $hr, $n ) = @_;
    return undef if ( !ref($hr) );                 # error check
    return undef if ( !defined $n || $n eq '' );
    return $$hr{$n};
}

########################################################################

=item B<multipart_setfile>

Params: \%multi_hash, $param_name, $file_path [, $filename]

Return: undef on error, 1 on success

NOTE: this function does not actually add the contents of $file_path into
the %multi_hash; instead, multipart_write() inserts the content when
generating the final request.

=cut

sub multipart_setfile {
    my ( $hr, $n, $path ) = ( shift, shift, shift );
    my ($fname) = shift;

    return undef if ( !ref($hr) );                 # error check
    return undef if ( !defined $n || $n eq '' );
    return undef if ( !defined $path );
    return undef if ( !( -e $path && -f $path ) );

    if ( !defined $fname ) {
        $path =~ m/[\\\/]([^\\\/]+)$/;
        $fname = $1 || "whisker-file";
    }

    $$hr{$n} = "\0FILE";
    $$hr{"\0$n"} = [ $path, $fname ];
    return 1;
}

########################################################################

=item B<multipart_getfile>

Params: \%multi_hash, $file_param_name

Return: $path, $name ($path=undef on error)

multipart_getfile is used to retrieve information for a file
parameter contained in %multi_hash.  To use this you would most
likely do:

 ($path,$fname)=LW2::multipart_getfile(\%multi,"param_name");

=cut

sub multipart_getfile {
    my ( $hr, $n ) = @_;

    return undef if ( !ref($hr) );                                 # error check
    return undef if ( !defined $n || $n eq '' );
    return undef if ( !defined $$hr{$n} || $$hr{$n} ne "\0FILE" );

    return @{ $$hr{"\0$n"} };
}

########################################################################

=item B<multipart_boundary>

Params: \%multi_hash [, $new_boundary_name]

Return: $current_boundary_name

multipart_boundary is used to retrieve, and optionally set, the
multipart boundary used for the request.

NOTE: the function does no checking on the supplied boundary, so if 
you want things to work make sure it's a legit boundary.  Libwhisker
does *not* prefix it with any '---' characters.

=cut

sub multipart_boundary {
    my ( $hr, $new ) = @_;
    my $ret;

    return undef if ( !ref($hr) );    # error check

    if ( !defined $$hr{"\0BOUNDARY"} ) {

        # create boundary on the fly
        my $b  = uc( utils_randstr(20) );
        my $b2 = '-' x 32;
        $$hr{"\0BOUNDARY"} = "$b2$b";
    }

    $ret = $$hr{"\0BOUNDARY"};
    if ( defined $new ) {
        $$hr{"\0BOUNDARY"} = $new;
    }

    return $ret;
}

########################################################################

=item B<multipart_write>

Params: \%multi_hash, \%request

Return: 1 if successful, undef on error

multipart_write is used to parse and construct the multipart data
contained in %multi_hash, and place it ready to go in the given whisker
hash (%request) structure, to be sent to the server.

NOTE: file contents are read into the final %request, so it's possible for
the hash to get *very* large if you have (a) large file(s).

=cut

sub multipart_write {
    my ( $mp, $hr ) = @_;

    return undef if ( !ref($mp) );    # error check
    return undef if ( !ref($hr) );    # error check

    if ( !defined $$mp{"\0BOUNDARY"} ) {

        # create boundary on the fly
        my $b  = uc( utils_randstr(20) );
        my $b2 = '-' x 32;
        $$mp{"\0BOUNDARY"} = "$b2$b";
    }

    my $B   = $$mp{"\0BOUNDARY"};
    my $EOL = $$hr{whisker}->{http_eol} || "\x0d\x0a";

    my $keycount = 0;
    foreach ( keys %$mp ) {
        next if ( substr( $_, 0, 1 ) eq "\0" );
        $keycount++;
        if ( $$mp{$_} eq "\0FILE" ) {
            my ( $path, $name ) = multipart_getfile( $mp, $_ );
            next if ( !defined $path );
            $$hr{whisker}->{data} .= "$B$EOL";
            $$hr{whisker}->{data} .=
              "Content-Disposition: " . "form-data; name=\"$_\"; ";
            $$hr{whisker}->{data} .= "filename=\"$name\"$EOL";
            $$hr{whisker}->{data} .=
              "Content-Type: " . "application/octet-stream$EOL";
            $$hr{whisker}->{data} .= $EOL;
            next if ( !open( IN, "<$path" ) );
            binmode(IN);    # stupid Windows

            while (<IN>) {
                $$hr{whisker}->{data} .= $_;
            }
            close(IN);
            $$hr{whisker}->{data} .= $EOL;    # WARNING: is this right?
        }
        else {
            $$hr{whisker}->{data} .= "$B$EOL";
            $$hr{whisker}->{data} .=
              "Content-Disposition: " . "form-data; name=\"$_\"$EOL";
            $$hr{whisker}->{data} .= "$EOL$$mp{$_}$EOL";
        }
    }

    if ($keycount) {
        $$hr{whisker}->{data} .= "$B--$EOL";    # closing boundary
        $$hr{"Content-Length"} = length( $$hr{whisker}->{data} );
        $$hr{"Content-Type"}   = "multipart/form-data; boundary=$B";
        return 1;
    }
    else {

        # multipart hash didn't contain params to upload
        return undef;
    }
}

########################################################################

=item B<multipart_read>

Params: \%multi_hash, \%hout_response [, $filepath ]

Return: 1 if successful, undef on error

multipart_read will parse the data contents of the supplied
%hout_response hash, by passing the appropriate info to
multipart_read_data().  Please see multipart_read_data() for more
info on parameters and behaviour.

NOTE: this function will return an error if the given %hout_response
Content-Type is not set to "multipart/form-data".

=cut

sub multipart_read {
    my ( $mp, $hr, $fp ) = @_;

    return undef if ( !( defined $mp && ref($mp) ) );
    return undef if ( !( defined $hr && ref($hr) ) );

    my $ctype = utils_find_lowercase_key( $hr, 'content-type' );
    return undef if ( !defined $ctype );
    return undef if ( $ctype !~ m#^multipart/form-data#i );

    return multipart_read_data( $mp, \$$hr{'whisker'}->{'data'}, undef, $fp );

}

########################################################################

=item B<multipart_read_data>

Params: \%multi_hash, \$data, $boundary [, $filepath ]

Return: 1 if successful, undef on error

multipart_read_data parses the contents of the supplied data using 
the given boundary and puts the values in the supplied %multi_hash.  
Embedded files will *not* be saved unless a $filepath is given, which
should be a directory suitable for writing out temporary files.

NOTE: currently only application/octet-stream is the only supported
file encoding.  All other file encodings will not be parsed/saved.

=cut

sub multipart_read_data {
    my ( $mp, $dr, $bound, $fp ) = @_;

    return undef if ( !( defined $mp && ref($mp) ) );
    return undef if ( !( defined $dr && ref($dr) ) );

    # if $bound is undef, then we'll snag what looks to be
    # the first boundry from the data.
    if ( !defined $bound ) {
        if ( $$dr =~ /([-]{5,}[A-Z0-9]+)[\r\n]/i ) {
            $bound = $1;
        }
        else {

            # we didn't spot a typical boundary; error
            return undef;
        }
    }

    if ( defined $fp && !( -d $fp && -w $fp ) ) {
        $fp = undef;
    }

    my $line = utils_getline_crlf( $dr, 0 );
    return undef if ( !defined $line );
    return undef if ( index( $line, $bound ) != 0 );

    my $done = 0;
    while ( !$done ) {
        $done = _multipart_read_data_part( $mp, $dr, $bound, $fp );
    }

    return 1;
}

########################################################################

sub _multipart_read_data_part {
    my ( $mp, $dr, $bound, $fp ) = @_;

    my $dispinfo = utils_getline_crlf($dr);
    return 1 if ( !defined $dispinfo );
    return 1 if ( length($dispinfo) == 0 );
    my $lcdisp = lc($dispinfo);

    if ( index( $lcdisp, 'content-disposition: form-data;' ) != 0 ) {
        return 1;
    }    # bad disposition

    my ( $s, $e, $l );

    $s = index( $lcdisp, 'name="', 30 );
    $e = index( $lcdisp, '"',      $s + 6 );
    return 1 if ( $s == -1 || $e == -1 );
    my $NAME = substr( $dispinfo, $s + 6, $e - $s - 6 );

    $s = index( $lcdisp, 'filename="', $e );
    my $FILENAME = undef;
    if ( $s != -1 ) {
        $e = index( $lcdisp, '"', $s + 10 );
        return 1 if ( $e == -1 );    # puke; malformed filename
        $FILENAME = substr( $dispinfo, $s + 10, $e - $s - 10 );
        $s        = rindex( $FILENAME, '\\' );
        $e        = rindex( $FILENAME, '/' );
        $s = $e if ( $e > $s );
        $FILENAME = substr( $FILENAME, $s + 1, length($FILENAME) - $s );
    }

    my $CTYPE = utils_getline_crlf($dr);

    return 1 if ( !defined $CTYPE );
    $CTYPE = lc($CTYPE);

    if ( length($CTYPE) > 0 ) {
        $s = index( $CTYPE, 'content-type:' );
        return 1 if ( $s != 0 );    # bad ctype line
        $CTYPE = substr( $CTYPE, 13, length($CTYPE) - 13 );
        $CTYPE =~ tr/ \t//d;
        my $xx = utils_getline_crlf($dr);
        return 1 if ( !defined $xx );
        return 1 if ( length($xx) > 0 );
    }
    else {
        $CTYPE = 'application/octet-stream';
    }

    my $VALUE = '';
    while ( defined( $l = utils_getline_crlf($dr) ) ) {
        last if ( index( $l, $bound ) == 0 );
        $VALUE .= $l;
        $VALUE .= "\r\n";
    }

    substr( $VALUE, -2, 2 ) = '';

    if ( !defined $FILENAME ) {    # read in param
        $$mp{$NAME} = $VALUE;
        return 0;

    }
    else {                         # read in file
        $$mp{$NAME} = "\0FILE";
        return 0 if ( !defined $fp );

        # TODO: funky content types, like application/x-macbinary
        if ( $CTYPE ne 'application/octet-stream' ) {
            return 0;
        }

        my $rfn      = lc( utils_randstr(12) );
        my $fullpath = "$fp$rfn";

        $$mp{"\0$NAME"} = [ undef, $FILENAME ];
        return 0 if ( !open( OUT, ">$fullpath" ) );    # error opening file
        binmode(OUT);                                  # stupid Windows
        $$mp{"\0$NAME"} = [ $fullpath, $FILENAME ];
        print OUT $VALUE;
        close(OUT);

        return 0;

    }    # if !defined $FILENAME

    return 0;    # um, this should never be reached...
}

########################################################################

=item B<multipart_files_list>

Params: \%multi_hash

Return: @files

multipart_files_list returns an array of parameter names for all
the files that are contained in %multi_hash.

=cut

sub multipart_files_list {
    my ($mp) = shift;
    my @ret;

    return () if ( !( defined $mp && ref($mp) ) );
    while ( my ( $K, $V ) = each(%$mp) ) {
        push( @ret, $K ) if ( $V eq "\0FILE" );
    }
    return @ret;
}

########################################################################

=item B<multipart_params_list>

Params: \%multi_hash

Return: @params

multipart_files_list returns an array of parameter names for all
the regular parameters (non-file) that are contained in %multi_hash.

=cut

sub multipart_params_list {
    my ($mp) = shift;
    my @ret;

    return () if ( !( defined $mp && ref($mp) ) );
    while ( my ( $K, $V ) = each(%$mp) ) {
        push( @ret, $K ) if ( $V ne "\0FILE"
            && substr( $K, 0, 1 ) ne "\0" );
    }
    return @ret;
}

########################################################################

