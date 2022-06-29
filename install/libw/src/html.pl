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

################################################################

=item B<html_find_tags>

Params: \$data, \&callback_function [, $xml_flag, $funcref, \%tag_map]

Return: nothing

html_find_tags parses a piece of HTML and 'extracts' all found tags,
passing the info to the given callback function.  The callback function 
must accept two parameters: the current tag (as a scalar), and a hash ref 
of all the tag's elements. For example, the tag <a href="/file"> will
pass 'a' as the current tag, and a hash reference which contains
{'href'=>"/file"}.

The xml_flag, when set, causes the parser to do some extra processing
and checks to accomodate XML style tags such as <tag foo="bar"/>.

The optional %tagmap is a hash of lowercase tag names.  If a tagmap is
supplied, then the parser will only call the callback function if the
tag name exists in the tagmap.

The optional $funcref variable is passed straight to the callback
function, allowing you to pass flags or references to more complex
structures to your callback function.

=cut

{    # contained variables
    $DR  = undef;    # data reference
    $c   = 0;        # parser pointer
    $LEN = 0;

    sub html_find_tags {
        my ( $dataref, $callbackfunc, $xml, $fref, $tagmap ) = @_;

        return if ( !( defined $dataref      && ref($dataref) ) );
        return if ( !( defined $callbackfunc && ref($callbackfunc) ) );
        $xml ||= 0;

        my ( $INTAG, $CURTAG, $LCCURTAG, $ELEMENT, $VALUE, $cc ) = (0);
        my ( %TAG, $ret, $start, $tagstart, $tempstart, $x, $found );
        my $usetagmap = ( ( defined $tagmap && ref($tagmap) ) ? 1 : 0 );
        $CURTAG = $LCCURTAG = $ELEMENT = $VALUE = $cc = '';
        $DR     = $dataref;

        $LEN = length($$dataref);
        for ( $c = 0 ; $c < $LEN ; $c++ ) {

            $cc = substr( $$dataref, $c, 1 );
            next if ( !$INTAG && $cc ne '>' && $cc ne '<' );

            if ( $cc eq '<' ) {
                if ($INTAG) {

                    # we're already in a tag...
                    # we trick the parser into thinking we end cur tag
                    $cc = '>';
                    $c--;

                }
                elsif ($xml
                    && $LEN > ( $c + 9 )
                    && substr( $$dataref, $c + 1, 8 ) eq '![CDATA[' )
                {
                    $c += 9;
                    $tempstart = $c;
                    $found     = index( $$dataref, ']]>', $c );
                    $c         = $found + 2;
                    $c         = $LEN if ( $found < 0 );         # malformed XML
                         # what to do with CDATA?
                    next;

                }
                elsif ( $LEN > ( $c + 3 )
                    && substr( $$dataref, $c + 1, 3 ) eq '!--' )
                {
                    $tempstart = $c;
                    $c += 4;
                    $found = index( $$dataref, '-->', $c );
                    if ( $found < 0 ) {
                        $found = index( $$dataref, '>', $c );
                        $found = $LEN if ( $found < 0 );
                        $c = $found;
                    }
                    else {
                        $c = $found + 2;
                    }
                    if ( $usetagmap == 0 || defined $tagmap->{'!--'} ) {
                        my $dat = substr(
                            $$dataref,
                            $tempstart + 4,
                            $found - $tempstart - 4
                        );
                        &$callbackfunc( '!--', { '=' => $dat },
                            $dataref, $tempstart, $c - $tempstart + 1, $fref );
                    }
                    next;

                }
                elsif ( !$INTAG ) {
                    next if ( substr( $$dataref, $c + 1, 1 ) =~ tr/ \t\r\n// );
                    $c++;
                    $INTAG    = 1;
                    $tagstart = $c - 1;

                    $CURTAG = '';
                    while ( $c < $LEN
                        && ( $x = substr( $$dataref, $c, 1 ) ) !~
                        tr/ \t\r\n>=// )
                    {
                        $CURTAG .= $x;
                        $c++;
                    }

                    chop $CURTAG if ( $xml && substr( $CURTAG, -1, 1 ) eq '/' );
                    $c++ if ( defined $x && $x ne '>' );

                    $LCCURTAG = lc($CURTAG);
                    $INTAG = 0 if ( $LCCURTAG !~ tr/a-z0-9// );
                    next if ( $c >= $LEN );
                    $cc = substr( $$dataref, $c, 1 );
                }
            }

            if ( $cc eq '>' ) {
                next if ( !$INTAG );
                if ( $LCCURTAG eq 'script' && !$xml ) {
                    $tempstart = $c + 1;
                    pos($$dataref) = $c;
                    if ( $$dataref !~ m#(</script.*?>)#ig ) {

                        # what to do if closing script not found?
                        # right now, we'll just leave the tag alone;
                        # this won't affect the 'absorption' of the
                        # javascript code (and thus, affect parsing)
                    }
                    else {
                        $c = pos($$dataref) - 1;
                        my $l = length($1);
                        $TAG{'='} =
                          substr( $$dataref, $tempstart,
                            $c - $tempstart - $l + 1 );
                    }

                }
                elsif ( $LCCURTAG eq 'textarea' && !$xml ) {
                    $tempstart = $c + 1;
                    pos($$dataref) = $c;
                    if ( $$dataref !~ m#(</textarea.*?>)#ig ) {

                        # no closing textarea...
                    }
                    else {
                        $c = pos($$dataref) - 1;
                        my $l = length($1);
                        $TAG{'='} =
                          substr( $$dataref, $tempstart,
                            $c - $tempstart - $l + 1 );
                    }
                }

                $INTAG = 0;
                $TAG{'/'}++
                  if ( $xml && substr( $$dataref, $c - 1, 1 ) eq '/' );
                &$callbackfunc( $CURTAG, \%TAG, $dataref, $tagstart,
                    $c - $tagstart + 1, $fref )
                  if ( $usetagmap == 0 || defined $tagmap->{$LCCURTAG} );
                $CURTAG = $LCCURTAG = '';
                %TAG = ();
                next;
            }

            if ($INTAG) {
                $ELEMENT = '';
                $VALUE   = undef;

                # eat whitespace
                pos($$dataref) = $c;
                if ( $$dataref !~ m/[^ \t\r\n]/g ) {
                    $c = $LEN;
                    next;    # should we really abort?
                }
                $start = pos($$dataref) - 1;

                if ( $$dataref !~ m/[ \t\r\n<>=]/g ) {
                    $c = $LEN;
                    next;    # should we really abort?
                }
                $c = pos($$dataref) - 1;

                if ( $c > $start ) {
                    $ELEMENT = substr( $$dataref, $start, $c - $start );
                    chop $ELEMENT
                      if ( $xml && substr( $ELEMENT, -1, 1 ) eq '/' );
                }

                $cc = substr( $$dataref, $c, 1 );
                if ( $cc ne '>' ) {

                    # eat whitespace
                    if ( $cc =~ tr/ \t\r\n// ) {
                        $c++
                          while ( substr( $$dataref, $c, 1 ) =~ tr/ \t\r\n// );
                    }

                    if ( substr( $$dataref, $c, 1 ) eq '=' ) {
                        $c++;
                        $start = $c;
                        my $p = substr( $$dataref, $c, 1 );
                        if ( $p eq '"' || $p eq '\'' ) {
                            $c++;
                            $start++;
                            $c = index( $$dataref, $p, $c );
                            if ( $c < 0 ) { $c = $LEN; next; }    # Bad HTML
                            $VALUE = substr( $$dataref, $start, $c - $start );
                            $c++;
                            pos($$dataref) = $c;
                        }
                        else {
                            pos($$dataref) = $c;
                            if ( $$dataref !~ /[ \t\r\n>]/g ) {
                                $c = $LEN;
                            }
                            else {
                                $c     = pos($$dataref) - 1;
                                $VALUE =
                                  substr( $$dataref, $start, $c - $start );
                                chop $VALUE
                                  if ( $xml
                                    && substr( $$dataref, $c - 1, 2 ) eq '/>' );
                            }
                        }

                        if ( substr( $$dataref, $c, 1 ) =~ tr/ \t\r\n// ) {
                            if ( $$dataref !~ /[^ \t\r\n]/g ) {
                                $c = $LEN;
                                next;    # should we really abort?
                            }
                            $c = pos($$dataref) - 1;
                        }
                    }
                }    # if $c ne '>'
                $c--;
                $TAG{$ELEMENT} = $VALUE
                  if ( $ELEMENT ne '' || ( $xml && $ELEMENT ne '/' ) );
            }
        }

        # finish off any tags we had going
        if ($INTAG) {
            &$callbackfunc( $CURTAG, \%TAG, $dataref, $tagstart,
                $c - $tagstart + 1, $fref )
              if ( $usetagmap == 0 || defined $tagmap->{$LCCURTAG} );
        }

        $DR = undef;    # void dataref pointer
    }

################################################################

=item B<html_find_tags_rewrite>

Params: $position, $length, $replacement

Return: nothing

html_find_tags_rewrite() is used to 'rewrite' an HTML stream from
within an html_find_tags() callback function.  In general, you can
think of html_find_tags_rewrite working as:

substr(DATA, $position, $length) = $replacement

Where DATA is the current HTML string the html parser is using.
The reason you need to use this function and not substr() is
because a few internal parser pointers and counters need to be
adjusted to accomodate the changes.

If you want to remove a piece of the string, just set the
replacement to an empty string ('').  If you wish to insert a
string instead of overwrite, just set $length to 0; your string
will be inserted at the indicated $position.

=cut

    sub html_find_tags_rewrite {
        return if ( !defined $DR );
        my ( $pos, $len, $replace_str ) = @_;

        # replace the data
        substr( $$DR, $pos, $len ) = $replace_str;

        # adjust pointer and length
        my $l = ( length($replace_str) - $len );
        $c   += $l;
        $LEN += $l;
    }

################################################################

    sub _html_find_tags_adjust {
        my ( $p, $l ) = @_;
        $c   += $p;
        $LEN += $l;
    }
}    # end container

################################################################

=item B<html_link_extractor>

Params: \$html_data

Return: @urls

The html_link_extractor() function uses the internal crawl tests to
extract all the HTML links from the given HTML data stream.

Note: html_link_extractor() does not unique the returned array of
discovered links, nor does it attempt to remove javascript links
or make the links absolute.  It just extracts every raw link from
the HTML stream and returns it.  You'll have to do your own
post-processing.

=cut

sub html_link_extractor {
    my $data = shift;
    my $ptr;
    if ( ref($data) ) {
        $ptr = $data;
    }
    else {
        $ptr = \$data;
    }

    # emulate the crawl object parts we need
    my %OBJ = ( urls => [], forms => {} );
    $OBJ{response}                   = {};
    $OBJ{response}->{whisker}        = {};
    $OBJ{response}->{whisker}->{uri} = '';

    html_find_tags(
        $ptr,                           # data
        \&_crawl_extract_links_test,    # callback function
        0,                              # xml mode
        \%OBJ,                          # data object
        \%_crawl_linktags
    );                                  # tagmap

    return @{ $OBJ{urls} };
}

################################################################

