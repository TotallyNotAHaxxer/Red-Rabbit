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

#####################################################

# cluster global variables
%_crawl_config = (
    'save_cookies'         => 0,
    'reuse_cookies'        => 1,
    'save_offsites'        => 0,
    'save_non_http'        => 0,
    'follow_moves'         => 1,
    'url_limit'            => 1000,
    'use_params'           => 0,
    'params_double_record' => 0,
    'skip_ext'             => {
        gif => 1,
        jpg => 1,
        png => 1,
        gz  => 1,
        swf => 1,
        pdf => 1,
        zip => 1,
        wav => 1,
        mp3 => 1,
        asf => 1,
        tgz => 1
    },
    'save_skipped'    => 0,
    'save_referrers'  => 0,
    'use_referrers'   => 1,
    'do_head'         => 0,
    'callback'        => 0,
    'netloc_bug'      => 1,
    'normalize_uri'   => 1,
    'source_callback' => 0
);

%_crawl_linktags = (
    'a'          => 'href',
    'applet'     => [qw(codebase archive code)],
    'area'       => 'href',
    'base'       => 'href',
    'bgsound'    => 'src',
    'blockquote' => 'cite',
    'body'       => 'background',
    'del'        => 'cite',
    'embed'      => [qw(src pluginspage)],
    'form'       => 'action',
    'frame'      => [qw(src longdesc)],
    'iframe'     => [qw(src longdesc)],
    'ilayer'     => 'background',
    'img'        => [qw(src lowsrc longdesc usemap)],
    'input'      => [qw(src usemap)],
    'ins'        => 'cite',
    'isindex'    => 'action',
    'head'       => 'profile',
    'layer'      => [qw(background src)],
    'link'       => 'href',

    #	 'meta'    => 'http-equiv',
    'object' => [qw(codebase data archive usemap)],
    'q'      => 'cite',
    'script' => 'src',
    'table'  => 'background',
    'td'     => 'background',
    'th'     => 'background',
    'xmp'    => 'href',
);

#####################################################

=item B<crawl_new>

Params: $START, $MAX_DEPTH, \%request_hash [, \%tracking_hash ]

Return: $crawl_object

The crawl_new() functions initializes a crawl object (hash) to the default
values, and then returns it for later use by crawl().  $START is the starting
URL (in the form of 'http://www.host.com/url'), and MAX_DEPTH is the maximum
number of levels to crawl (the START URL counts as 1, so a value of 2 will
crawl the START URL and all URLs found on that page).  The request_hash
is a standard initialized request hash to be used for requests; you should
set any authentication information or headers in this hash in order for
the crawler to use them.  The optional tracking_hash lets you supply a
hash for use in tracking URL results (otherwise crawl_new() will allocate
a new anon hash).

=cut

sub crawl_new {
    my ( $start, $depth, $reqref, $trackref ) = @_;
    my %X;

    return undef if ( !defined $start  || !defined $depth );
    return undef if ( !defined $reqref || !ref($reqref) );
    $trackref = {} if ( !defined $trackref || !ref($trackref) );

    $X{track}   = $trackref;
    $X{request} = $reqref;
    $X{depth}   = $depth || 2;
    $X{start}   = $start;
    $X{magic}   = 7340;

    $X{reset} = sub {
        $X{errors}      = [];    # all errors encountered
        $X{urls}        = [];    # temp; used to hold all URLs on page
        $X{server_tags} = {};    # all server tags found
        $X{referrers}   = {};    # who refers to what URLs
        $X{offsites}    = {};    # all URLs that point offsite
        $X{response}    = {};    # temp; the response hash
        $X{non_http}    = {};    # all non_http URLs found
        $X{cookies}     = {};    # all cookies found
        $X{forms}       = {};    # all forms found
        $X{jar}         = {};    # temp; cookie jar
        $X{url_queue}   = [];    # temp; URLs to still fetch

        $X{config} = {};
        %{ $X{config} } = %_crawl_config;

        %{ $X{track} } = ();
        $X{parsed_page_count} = 0;
    };

    $X{crawl} = sub { crawl( \%X, @_ ) };
    $X{reset}->();

    return \%X;
}

#####################################################

=item B<crawl>

Params: $crawl_object [, $START, $MAX_DEPTH ]

Return: $count [ undef on error ] 

The heart of the crawl package.  Will perform an HTTP crawl on the
specified HOST, starting at START URI, proceeding up to MAX_DEPTH. 

Crawl_object needs to be the variable returned by crawl_new().  You can
also indirectly call crawl() via the crawl_object itself:

	$crawl_object->{crawl}->($START,$MAX_DEPTH)

Returns the number of URLs actually crawled (not including those skipped).

=cut

{    # START OF CRAWL CONTAINER

    sub crawl {
        my ( $C, $START, $MAX_DEPTH ) = @_;
        return undef if ( !defined $C || !ref($C) || $C->{magic} != 7340 );

        # shortcuts, to reduce dereferences and typing
        my $CONFIG = $C->{config};
        my $TRACK  = $C->{track};
        my $URLS   = $C->{urls};
        my $RESP   = $C->{response};
        my $REQ    = $C->{request};
        my $Q      = $C->{url_queue};

        $START ||= $C->{start};
        $C->{depth} = $MAX_DEPTH || $C->{depth};

        my ( $COUNT, $T, @ST ) = ( 0, '' );

        # ST[] = [ 0.HOST, 1.PORT, 2.URL, 3.DEPTH, 4.CWD, 5.REF ]

        my @v = uri_split($START);

        my $error = undef;
        $error = 'Start protocol not http or https'
          if ( $v[1] ne 'http' && $v[1] ne 'https' );
        $error = 'Bad start host' if ( !defined $v[2] || $v[2] eq '' );
        push( @{ $C->{errors} }, $error ) && return undef if ( defined $error );

        @ST = ( $v[2], $v[3], $v[0], 1, '', '' );

        $REQ->{whisker}->{ssl}  = 1 if ( $v[1] eq 'https' );
        $REQ->{whisker}->{host} = $ST[0];
        $REQ->{whisker}->{port} = $ST[1];
        $REQ->{whisker}->{lowercase_incoming_headers} = 1;
        $REQ->{whisker}->{ignore_duplicate_headers}   = 0;
        delete $REQ->{whisker}->{parameters};
        http_fixup_request($REQ);

        push @$Q, \@ST;

        while (@$Q) {
            @ST = @{ shift @$Q };

            next if ( defined $TRACK->{ $ST[2] } && $TRACK->{ $ST[2] } ne '?' );
            if ( $ST[3] > $C->{depth} ) {
                $TRACK->{ $ST[2] } = '?' if ( $CONFIG->{save_skipped} > 0 );
                next;
            }

            $ST[4] = uri_get_dir( $ST[2] );
            $REQ->{whisker}->{uri} = $ST[2];
            if ( $ST[5] ne '' && $CONFIG->{use_referrers} > 0 ) {
                $REQ->{Referrer} = $ST[5];
            }

            my $result = _crawl_do_request( $REQ, $RESP, $C );
            if ( $result == 1 || $result == 2 ) {
                push @{ $C->{errors} }, "$ST[2]: $RESP->{whisker}->{error}";
                next;
            }

            $COUNT++;
            $TRACK->{ $ST[2] } = $RESP->{whisker}->{code}
              if ( $result == 0 || $result == 4 );
            $TRACK->{ $ST[2] } = '?'
              if ( ( $result == 3 || $result == 5 )
                && $CONFIG->{save_skipped} > 0 );

            if ( defined $RESP->{server} && !ref( $RESP->{server} ) ) {
                $C->{server_tags}->{ $RESP->{server} }++;
            }

            if ( defined $RESP->{'set-cookie'} ) {
                if ( $CONFIG->{save_cookies} > 0 ) {
                    if ( ref( $RESP->{'set-cookie'} ) ) {
                        $C->{cookies}->{$_}++
                          foreach ( @{ $RESP->{'set-cookie'} } );
                    }
                    else {
                        $C->{cookies}->{ $RESP->{'set-cookie'} }++;
                    }
                }
                cookie_read( $C->{jar}, $RESP )
                  if ( $CONFIG->{reuse_cookies} > 0 );
            }

            next if ( $result == 4 || $result == 5 );
            next if ( scalar @$Q > $CONFIG->{url_limit} );

            if ( $result == 0 ) {    # page should be parsed
                if ( $CONFIG->{source_callback} != 0
                    && ref( $CONFIG->{source_callback} ) eq 'CODE' )
                {
                    &{ $CONFIG->{source_callback} }($C);
                }

                html_find_tags( \$RESP->{whisker}->{data},
                    \&_crawl_extract_links_test, 0, $C, \%_crawl_linktags );
                $C->{parsed_page_count}++;
            }

            push @$URLS, $RESP->{location} if ( $result == 3 );

            foreach $T (@$URLS) {
                $T =~ tr/\0\r\n//d;
                next if ( length($T) == 0 );
                next if ( $T =~ /^#/i );       # fragment

                push @{ $C->{referrers}->{$T} }, $ST[2]
                  if ( $CONFIG->{save_referrers} > 0 );

                if (   $T =~ /^([a-zA-Z0-9]*):/
                    && lc($1) ne 'http'
                    && lc($1) ne 'https' )
                {
                    push @{ $C->{non_http}->{$T} }, $ST[2]
                      if ( $CONFIG->{save_non_http} > 0 );
                    next;
                }

                if ( substr( $T, 0, 2 ) eq '//' && $CONFIG->{netloc_bug} > 0 ) {
                    if ( $REQ->{whisker}->{ssl} > 0 ) { $T = 'https:' . $T; }
                    else { $T = 'http:' . $T; }
                }
                if ( $CONFIG->{callback} != 0 ) {
                    next if &{ $CONFIG->{callback} }( $T, $C );
                }

                $T = uri_absolute( $T, $ST[4], $CONFIG->{normalize_uri} );

                # (uri,protocol,host,port,params,frag,user,pass)
                @v = uri_split($T);

                # make sure URL is on same host and port
                if (   ( defined $v[2] && $v[2] ne $ST[0] )
                    || ( $v[3] > 0 && $v[3] != $ST[1] ) )
                {
                    $C->{offsites}->{ uri_join(@v) }++
                      if ( $CONFIG->{save_offsites} > 0 );
                    next;
                }

                if ( $v[0] =~ /\.([a-z0-9]+)$/i ) {
                    if ( defined $CONFIG->{skip_ext}->{ lc($1) } ) {
                        $TRACK->{ $v[0] } = '?'
                          if ( $CONFIG->{save_skipped} > 0 );
                        next;
                    }
                }

                if ( defined $v[4] && $CONFIG->{use_params} > 0 ) {
                    $TRACK->{ $v[0] } = '?'
                      if ( $CONFIG->{params_double_record} > 0
                        && !defined $TRACK->{ $v[0] } );
                    $v[0] = $v[0] . '?' . $v[4];
                }

                next
                  if ( defined $TRACK->{ $v[0] } )
                  ;    # we've processed this already

                # ST[] = [ 0.HOST, 1.PORT, 2.URL, 3.DEPTH, 4.CWD, 5.REF ]
                push @$Q, [ $ST[0], $ST[1], $v[0], $ST[3] + 1, '', $ST[2] ];
            }    # foreach

            @$URLS = ();    # reset for next round
        }    # while

        return $COUNT;
    }    # end sub crawl

#####################################################

    sub _crawl_extract_links_test {
        my ( $TAG, $hr, $dr, $start, $len, $OBJ ) = ( lc(shift), @_ );

        return undef if ( !scalar %$hr );    # fastpath quickie

        # we know this is defined, due to our tagmap
        my $t = $_crawl_linktags{$TAG};

        while ( my ( $key, $val ) = each %$hr ) {    # normalize element values
            $$hr{ lc($key) } = $val;
        }

        # all of this just to catch meta refresh URLs
        if (   $TAG eq 'meta'
            && defined $$hr{'http-equiv'}
            && $$hr{'http-equiv'} eq 'refresh'
            && defined $$hr{'content'}
            && $$hr{'content'} =~ m/url=(.+)/i )
        {
            push( @{ $OBJ->{urls} }, $1 );

        }
        elsif ( ref($t) ) {
            foreach (@$t) {
                push( @{ $OBJ->{urls} }, $$hr{$_} ) if ( defined $$hr{$_} );
            }
        }
        else {
            push( @{ $OBJ->{urls} }, $$hr{$t} ) if ( defined $$hr{$t} );
        }

        if ( $TAG eq 'form' && defined $$hr{action} ) {
            my $u = $OBJ->{response}->{whisker}->{uri};
            $OBJ->{forms}->{ uri_absolute( $$hr{action}, $u, 1 ) }++;
        }

        return undef;
    }

################################################################

    sub _crawl_do_request_ex {
        my ( $hrin, $hrout, $OBJ ) = @_;
        my $ret;

        $ret = http_do_request( $hrin, $hrout );

        return ( 2, $ret )
          if ( $ret == 2 );     # if there was connection error, do not continue
        if   ( $ret == 0 ) {    # successful request

            # WARNING: what if *all* HEAD respones are 302'd on purpose, but
            #          all GETs are normal?
            if (   $$hrout{whisker}->{code} < 308
                && $$hrout{whisker}->{code} > 300 )
            {
                if ( $OBJ->{config}->{follow_moves} > 0 ) {
                    return ( 3, $ret )
                      if ( defined $$hrout{location}
                        && !ref( $$hrout{location} ) );
                }
                return ( 5, $ret );    # not avail
            }

            if ( $$hrout{whisker}->{code} == 200 ) {

                # no content-type is treated as text/htm
                if ( defined $$hrout{'content-type'}
                    && $$hrout{'content-type'} !~ /^text\/htm/i )
                {
                    return ( 4, $ret );
                }
            }
        }
        return ( -1, $ret );    # fallthrough
    }

################################################################

    sub _crawl_do_request {
        my ( $hrin, $hrout, $OBJ ) = @_;
        my ( $cret, $lwret );

        if ( $OBJ->{config}->{do_head} && $$hrin{whisker}->{method} ne 'HEAD' )
        {
            my $save = $$hrin{whisker}->{method};
            $$hrin{whisker}->{method} = 'HEAD';
            ( $cret, $lwret ) = _crawl_do_request_ex( $hrin, $hrout, $OBJ );
            $$hrin{whisker}->{method} = $save;

            return $cret if ( $cret > 0 );

            if ( $lwret == 0 ) {    # successful request
                if ( $$hrout{whisker}->{code} == 501 ) {    # HEAD not allowed
                    $OBJ->{config}->{do_head} = 0;    # no more HEAD requests
                }
            }

            # request errors are essentially redone via GET, below
        }

        ( $cret, $lwret ) = _crawl_do_request_ex( $hrin, $hrout, $OBJ );
        return $lwret if ( $cret < 0 );
        return $cret;
    }

}    # CRAWL_CONTAINER

################################################################

