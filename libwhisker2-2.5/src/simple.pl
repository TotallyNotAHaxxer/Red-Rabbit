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

=item B<get_page>

Params: $url [, \%request]

Return: $code, $data ($code will be set to undef on error, $data will
			contain error message)

This function will fetch the page at the given URL, and return the HTTP response code
and page contents.  Use this in the form of:
($code,$html)=LW2::get_page("http://host.com/page.html")

The optional %request will be used if supplied.  This allows you to set
headers and other parameters.

=cut

sub get_page {
    my ( $URL, $hr ) = ( shift, shift );
    return ( undef, 'No URL supplied' ) if ( length($URL) == 0 );

    my ( %req, %resp );
    my $rptr;

    if ( defined $hr && ref($hr) ) {
        $rptr = $hr;
    }
    else {
        $rptr = \%req;
        http_init_request( \%req );
    }

    my @u = uri_split( $URL, $rptr );
    return ( undef, 'Non-HTTP URL supplied' )
      if ( $u[1] ne 'http' && $u[1] ne 'https' );
    http_fixup_request($rptr);

    if ( http_do_request( $rptr, \%resp ) ) {
        return ( undef, $resp{'whisker'}->{'error'} );
    }
    return ( $resp{'whisker'}->{'code'}, $resp{'whisker'}->{'data'} );
}

########################################################################

=item B<get_page_hash>

Params: $url [, \%request]

Return: $hash_ref (undef on no URL)

This function will fetch the page at the given URL, and return the whisker
HTTP response hash.  The return code of the function is set to
$hash_ref->{whisker}->{get_page_hash}, and uses the http_do_request()
return values.

Note: undef is returned if no URL is supplied

=cut

sub get_page_hash {
    my ( $URL, $hr ) = ( shift, shift );
    return undef if ( length($URL) == 0 );

    my ( %req, %resp );
    my $rptr;

    if ( defined $hr && ref($hr) ) {
        $rptr = $hr;
    }
    else {
        $rptr = \%req;
        http_init_request( \%req );
    }

    my @u = uri_split( $URL, $rptr );    # this is newer >=1.1 syntax
    return undef if ( $u[1] ne 'http' && $u[1] ne 'https' );
    http_fixup_request($rptr);

    my $r = http_do_request( $rptr, \%resp );
    $resp{whisker}->{get_page_hash} = $r;
    return \%resp;
}

########################################################################

=item B<get_page_to_file>

Params: $url, $filepath [, \%request]

Return: $code ($code will be set to undef on error)

This function will fetch the page at the given URL, place the resulting HTML
in the file specified, and return the HTTP response code.  The optional
%request hash sets the default parameters to be used in the request.

NOTE: libwhisker does not do any file checking; libwhisker will open the
supplied filepath for writing, overwriting any previously-existing files.
Libwhisker does not differentiate between a bad request, and a bad file
open.  If you're having troubles making this function work, make sure
that your $filepath is legal and valid, and that you have appropriate
write permissions to create/overwrite that file.

=cut

sub get_page_to_file {
    my ( $URL, $filepath, $hr ) = @_;

    return undef if ( length($URL) == 0 );
    return undef if ( length($filepath) == 0 );

    my ( %req, %resp );
    my $rptr;

    if ( defined $hr && ref($hr) ) {
        $rptr = $hr;
    }
    else {
        $rptr = \%req;
        http_init_request( \%req );
    }

    my @u = uri_split( $URL, $rptr );    # this is newer >=1.1 syntax
    return undef if ( $u[1] ne 'http' && $u[1] ne 'https' );
    http_fixup_request($rptr);
    return undef if ( http_do_request( $rptr, \%resp ) );

    open( OUT, ">$filepath" ) || return undef;
    binmode(OUT);                        # stupid Windows
    print OUT $resp{'whisker'}->{'data'};
    close(OUT);

    return $resp{'whisker'}->{'code'};
}

