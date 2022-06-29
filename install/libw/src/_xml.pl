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

{    # static variable container:
    $XML              = undef;          # current document hash ptr
    @tag_stack        = ("/\t/\t0");    # current tag stack order
    $XML_IGNORING_XSL = 0;

    sub _xml_callback {
        my ( $TAG, $hr, $dr, $start, $len ) = @_;

        # tags we don't care about
        return if ( $TAG =~ /^(!--|\?xml|!doctype)/ );

        # ignore all XSL
        $XML_IGNORING_XSL++ if ( $TAG eq 'xsl:stylesheet' );
        if ( $TAG eq '/xsl:stylesheet' ) {
            $XML_IGNORING_XSL = 0;
            return;
        }
        return if ($XML_IGNORING_XSL);

        my $ns = undef;
        $ns = $1 if ( $TAG =~ s/^(.+?):// );

        if ( substr( $TAG, 0, 1 ) ne '/' ) {
            my ( $n, $cur, $end ) = split( /\t/, $tag_stack[-1] );
            my $ab = _xml_join_name( $TAG, $cur );
            my $flag = 0;
            if ( defined( $$hr{'/'} ) ) {
                $flag++;
                delete $$hr{'/'};
            }

            # indicate previous element has children
            $tag_stack[-1] = "$n\t$cur\t0";

            my $rname = _xml_add_tag( $XML, $ab, undef, %$hr, $ns );
            my $e = $start + $len;
            push( @tag_stack, "$TAG\t$rname\t$e" ) unless $flag;

            # for some reason, we need to reset the array ptr for
            # push and pop to work correctly; might have something to
            # do with it being a static variable
            foreach (@tag_stack) { }
        }
        else {
            my ( $n, $path, $end ) = split( /\t/, pop(@tag_stack) );
            if ( "/$n" ne $TAG ) {

                # WHOA, bad XML?  We are now going to puke.
                %$XML = ();
                $XML->{'ERROR'} =
                  "Closing tag <$TAG> not expected" . " (expecting </$n>)";

                # short circuit parsing to the end
                _html_find_tags_adjust( 0, 0 );
            }

            if ( $end > 0 ) {    # set value
                my $val = substr( $$dr, $end, $start - $end );
                _xml_set_value( $XML, $path, _xml_deentify($val) );
            }
        }
        return;
    }

################################################################

    sub _xml_deentify {          # /* INTERNAL */
        my $val = shift;
        $val =~ s/&lt;/</g;
        $val =~ s/&gt;/>/g;
        $val =~ s/&quot;/"/g;
        $val =~ s/&#([0-9]{1,3});/chr($1)/eg;
        $val =~ s/&amp;/&/g;                    # must be last
        return $val;
    }

################################################################

=item B<xml_read_data>

Params: \$data

Return: $XML_object, undef on error

This function takes a reference to a scalar containing XML
markup and parses it into the %XML hash format used by the
other XML functions.

If a parsing error is encountered, the returned hash only has 
one element, 'ERROR', which contains an error message.

=cut

    sub xml_read_data {
        my $p = shift;
        return undef if ( !defined $p );
        $p         = \$p if ( !ref($p) );
        $XML       = {};                    # new anonymous hash
        @tag_stack = ("/\t/\t0");           # reset the tag stack
        html_find_tags( $p, \&_xml_callback, 1 );
        return $XML;                        # return new document
    }

################################################################

=item B<xml_read_file>

Params: $filename

Return: $XML_object, undef on error

xml_read_file opens the given $filename and attempts to
parse the XML data found within.

If a parsing error is encountered, the returned hash only has 
one element, 'ERROR', which contains an error message.

=cut

    sub xml_read_file {
        my $filename = shift;
        return undef if ( !defined $filename || $filename eq '' );
        return undef if ( !-e $filename || !-f $filename );

        my $data = '';
        open( IN, "<$filename" ) or return undef;
	binmode(IN); # Stupid Windows
        $data .= $_ while (<IN>);
        close(IN);

        return xml_read_data( \$data );
    }

}    # end of static variable container

################################################################

sub _xml_add_tag {
    my ( $hr, $name, $value, %e, $ns ) = @_;
    my $realname = $name;

    if ( exists $$hr{$name} ) {

        # already have a tag named that; make array
        my @t = _xml_get_4arr( $hr, $name );
        $t[0]++;
        $$hr{$name} = [ $t[0], $t[1], $t[2], $t[3] ];
        $realname .= '[' . $t[0] . ']';
    }

    # there's three different storage formats, in order
    # to reduce the amount of anonymous structures in
    # the main hash
    if ( ( scalar keys %e ) > 0 || defined $ns ) {   # we need a full anon array
        my $p = undef;
        if ( ( scalar keys %e ) > 0 ) {
            $p = [];
            foreach ( keys %e ) {                    # need to add each param
                $_ =~ s/^.+?://;                     # remove namespace
                next if ( $_ eq '' );
                push @$p, $_;
                $$hr{ _xml_join_name( "\@$_", $realname ) } = $e{$_};
            }
        }
        $$hr{$realname} = [ 0, $p, $value, $ns ];
    }
    elsif ( defined $value ) {                       # simple value
        $$hr{$realname} = $value;
    }
    else {    # empty tag with no value or elements
        $$hr{$realname} = undef;
    }

    return $realname;
}

################################################################

sub _xml_get_4arr {
    my ( $hr, $name ) = @_;
    return undef if ( !exists $$hr{$name} );
    return ( 0, undef, undef, undef ) if ( !defined $$hr{$name} );
    return ( 0, undef, $$hr{$name}, undef ) if ( !ref( $$hr{$name} ) );
    return @{ $$hr{$name} };
}

################################################################

sub _xml_check {
    my ( $hr, $name, $root ) = @_;
    return undef if ( !defined $hr || !ref($hr) );
    return undef if ( !defined $name || $name eq '' );
    $name = _xml_join_name( $name, $root ) if ( defined $root );
    return undef if ( !exists( $$hr{$name} ) );
    return $name;
}

################################################################

sub _xml_set_value {
    my ( $hr, $name, $value, $root ) = @_;
    return undef unless ( $name = _xml_check( $hr, $name, $root ) );
    if ( ref( $$hr{$name} ) ) {
        $$hr{$name}->[2] = $value;
    }
    else {
        $$hr{$name} = $value;
    }
    return 1;
}

################################################################

=item B<xml_get_multi>

Params: $XML_obj, $name [, $root] 

Return: @elements, undef on error

xml_get_multi() returns an array of absolute element pathnames for
the given named element, which can then be iterated over to access
child elements/parameters of each element.

=cut

sub xml_get_multi {
    my ( $count, $c, $hr, $name, $root ) = ( 0, 1, @_ );
    return undef unless ( $name = _xml_check( $hr, $name, $root ) );
    my @ret = ($name);
    return @ret if ( !defined $$hr{$name} || !ref( $$hr{$name} ) );
    $count = $$hr{$name}->[0];
    return @ret if ( $count == 0 );
    for ( $c = 1 ; $c <= $count ; $c++ ) {
        push @ret, $name . "[$c]";
    }
    return @ret;
}

################################################################

=item B<xml_is_multi>

Params: $XML_obj, $name [, $root] 

Return: 1 if multi, 0 if not, undef on error

This function checks to see if the named element has multiple elements
present in the given $XML_object.

=cut

sub xml_is_multi {
    my ( $hr, $name, $root ) = @_;
    return undef unless ( $name = _xml_check( $hr, $name, $root ) );
    return 0 if ( !ref( $$hr{$name} ) );
    return $$hr{$name}->[0];
}

################################################################

sub _xml_join_name {
    my ( $name, $root ) = @_;
    return undef if ( !defined $name );
    $root = '/' if ( !defined $root || $root eq '' );
    my $join = "$root/$name";
    $join =~ s#/{2,}#/#g;
    chop $join if ( substr( $join, -1, 1 ) eq '/' );
    return $join;
}

################################################################

=item B<xml_get_element_value>

Params: $XML_obj, $name [, $root] 

Return: $value, undef on error/not found

Fetches the value of the element of name $name.

=cut

sub xml_get_element_value {
    my ( $hr, $name, $root ) = @_;
    return undef unless ( $name = _xml_check( $hr, $name, $root ) );
    my @x = _xml_get_4arr( $hr, $name );
    return $x[2];
}

sub xml_gev {    # typing laziness shortcut
    goto &xml_get_element_value;
}

################################################################

=item B<xml_get_element_namespace>

Params: $XML_obj, $name [, $root] 

Return: $value, undef on error/not found/unknown

Returns the namespace of an element, if it was present in the XML
document.  If not known, then undef is returned.

=cut

sub xml_get_element_namespace {
    my ( $hr, $name, $root ) = @_;
    return undef unless ( $name = _xml_check( $hr, $name, $root ) );
    my @x = _xml_get_4arr( $hr, $name );
    return $x[3];
}

sub xml_gens {    # typing laziness shortcut
    goto &xml_get_element_namespace;
}

################################################################

=item B<xml_get_element_parameters>

Params: $XML_obj, $name [, $root] 

Return: @parameter_names, undef on error/not found

Returns a list of parameter names set for the named element.

=cut

sub xml_get_element_parameters {
    my ( $hr, $name, $root ) = @_;
    return undef unless ( $name = _xml_check( $hr, $name, $root ) );
    return undef if ( !ref( $$hr{$name} ) );
    my @t = _xml_get_4arr( $hr, $name );
    return undef if ( !defined( $t[1] ) );
    return @{ $t[1] };
}

sub xml_gep {    # typing laziness shortcut
    goto &xml_get_element_parameters;
}

################################################################

=item B<xml_if_exist>

Params: $XML_obj, $name [, $root] 

Return: 1 if found, undef if not found/error

Check to see if the named element exists.

=cut

sub xml_if_exist {
    my ( $hr, $name, $root ) = @_;
    return undef unless ( $name = _xml_check( $hr, $name, $root ) );
    return 1;
}

################################################################

