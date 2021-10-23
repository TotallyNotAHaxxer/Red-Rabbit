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

=item B<FORMS FUNCTIONS>

The goal is to parse the variable, human-readable HTML into concrete
structures useable by your program.  The forms functions does do a good job
at making these structures, but I will admit: they are not exactly simple,
and thus not a cinch to work with.  But then again, representing something
as complex as a HTML form is not a simple thing either.  I think the
results are acceptable for what's trying to be done.  Anyways...

Forms are stored in perl hashes, with elements in the following format:

 $form{'element_name'}=@([ 'type', 'value', @params ])

Thus every element in the hash is an array of anonymous arrays.  The first
array value contains the element type (which is 'select', 'textarea',
'button', or an 'input' value of the form 'input-text', 'input-hidden',
'input-radio', etc).

The second value is the value, if applicable (it could be undef if no
value was specified).  Note that select elements will always have an undef
value--the actual values are in the subsequent options elements.

The third value, if defined, is an anonymous array of additional tag
parameters found in the element (like 'onchange="blah"', 'size="20"',
'maxlength="40"', 'selected', etc).

The array does contain one special element, which is stored in the hash
under a NULL character ("\0") key.  This element is of the format:

 $form{"\0"}=['name', 'method', 'action', @parameters];

The element is an anonymous array that contains strings of the form's
name, method, and action (values can be undef), and a @parameters array
similar to that found in normal elements (above).

Accessing individual values stored in the form hash becomes a test of your
perl referencing skills.  Hint: to access the 'value' of the third element
named 'choices', you would need to do:

 $form{'choices'}->[2]->[1];

The '[2]' is the third element (normal array starts with 0), and the
actual value is '[1]' (the type is '[0]', and the parameter array is
'[2]').

=cut

################################################################

# Cluster global variables
%_forms_ELEMENTS = (
    'form'     => 1,
    'input'    => 1,
    'textarea' => 1,
    'button'   => 1,
    'select'   => 1,
    'option'   => 1,
    '/select'  => 1
);

################################################################

=item B<forms_read>

Params: \$html_data

Return: \@found_forms

This function parses the given $html_data into libwhisker form hashes.  
It returns a reference to an array of hash references to the found 
forms.

=cut

sub forms_read {
    my $dr = shift;
    return undef if ( !ref($dr) || length($$dr) == 0 );

    my $A = [ {}, [] ];

    html_find_tags( $dr, \&_forms_parse_callback, 0, $A, \%_forms_ELEMENTS );

    if ( scalar %{ $A->[0] } ) {
        push( @{ $A->[1] }, $A->[0] );
    }

    return $A->[1];
}

################################################################

=item B<forms_write>

Params: \%form_hash

Return: $html_of_form [undef on error]

This function will take the given %form hash and compose a generic HTML
representation of it, formatted with tabs and newlines in order to make it
neat and tidy for printing.

Note: this function does *not* escape any special characters that were
embedded in the element values.

=cut

sub forms_write {
    my $hr = shift;
    return undef if ( !ref($hr) || !( scalar %$hr ) );
    return undef if ( !defined $$hr{"\0"} );

    my $t = '<form name="' . $$hr{"\0"}->[0] . '" method="';
    $t .= $$hr{"\0"}->[1] . '" action="' . $$hr{"\0"}->[2] . '"';
    if ( defined $$hr{"\0"}->[3] ) {
        $t .= ' ' . join( ' ', @{ $$hr{"\0"}->[3] } );
    }
    $t .= ">\n";

    my ( $name, $ar );
    while ( ( $name, $ar ) = each(%$hr) ) {
        next if ( $name eq "\0" );
        next if ( $name eq '' && $ar->[0]->[0] eq '' );
        foreach $a (@$ar) {
            my $P = '';
            $P = ' ' . join( ' ', @{ $$a[2] } ) if ( defined $$a[2] );
            $t .= "\t";

            if ( $$a[0] eq 'textarea' ) {
                $t .= "<textarea name=\"$name\"$P>$$a[1]";
                $t .= "</textarea>\n";

            }
            elsif ( $$a[0] =~ m/^input-(.+)$/ ) {
                $t .= "<input type=\"$1\" name=\"$name\" ";
                $t .= "value=\"$$a[1]\"$P>\n";

            }
            elsif ( $$a[0] eq 'option' ) {
                $t .= "\t<option value=\"$$a[1]\"$P>$$a[1]\n";

            }
            elsif ( $$a[0] eq 'select' ) {
                $t .= "<select name=\"$name\"$P>\n";

            }
            elsif ( $$a[0] eq '/select' ) {
                $t .= "</select$P>\n";

            }
            else {    # button
                $t .= "<button name=\"$name\" value=\"$$a[1]\">\n";
            }
        }
    }

    $t .= "</form>\n";
    return $t;
}

################################################################

{    # these are 'private' static variables for &_forms_parse_html
    my $CURRENT_SELECT = undef;
    my $UNKNOWNS       = 0;

    sub _forms_parse_callback {
        my ( $TAG, $hr, $dr, $start, $len, $ar ) = ( lc(shift), @_ );
        my ( $saveparam, $parr, $key ) = ( 0, undef, '' );

        my $_forms_CURRENT = $ar->[0];
        my $_forms_FOUND   = $ar->[1];

        if ( scalar %$hr ) {
            while ( my ( $key, $val ) = each %$hr ) {
                if ( $key =~ tr/A-Z// ) {
                    delete $$hr{$key};
                    if ( defined $val ) { $$hr{ lc($key) } = $val; }
                    else { $$hr{ lc($key) } = undef; }
                }
            }
        }

        if ( $TAG eq 'form' ) {
            if ( scalar %$_forms_CURRENT ) {    # save last form
                push( @$_forms_FOUND, $_forms_CURRENT );
                $ar->[0] = {};
                $_forms_CURRENT = $ar->[0];
            }

            $_forms_CURRENT->{"\0"} =
              [ $$hr{name}, $$hr{method}, $$hr{action}, [] ];
            delete $$hr{'name'};
            delete $$hr{'method'};
            delete $$hr{'action'};
            $key      = "\0";
            $UNKNOWNS = 0;

        }
        elsif ( $TAG eq 'input' ) {
            $$hr{type}  = 'text'                  if ( !defined $$hr{type} );
            $$hr{name}  = 'unknown' . $UNKNOWNS++ if ( !defined $$hr{name} );
            $$hr{value} = undef                   if ( !defined $$hr{value} );
            $key        = $$hr{name};

            push @{ $_forms_CURRENT->{$key} },
              [ 'input-' . $$hr{type}, $$hr{value}, [] ];
            delete $$hr{'name'};
            delete $$hr{'type'};
            delete $$hr{'value'};

        }
        elsif ( $TAG eq 'select' ) {
            $$hr{name} = 'unknown' . $UNKNOWNS++ if ( !defined $$hr{name} );
            $key = $$hr{name};
            push @{ $_forms_CURRENT->{$key} }, [ 'select', undef, [] ];
            $CURRENT_SELECT = $key;
            delete $$hr{name};

        }
        elsif ( $TAG eq '/select' ) {
            push @{ $_forms_CURRENT->{$CURRENT_SELECT} },
              [ '/select', undef, [] ];
            $CURRENT_SELECT = undef;
            return undef;

        }
        elsif ( $TAG eq 'option' ) {
            return undef if ( !defined $CURRENT_SELECT );
            if ( !defined $$hr{value} ) {
                my $stop = index( $$dr, '<', $start + $len );
                return undef if ( $stop == -1 );    # MAJOR PUKE
                $$hr{value} =
                  substr( $$dr, $start + $len, ( $stop - $start - $len ) );
                $$hr{value} =~ tr/\r\n//d;
            }
            push @{ $_forms_CURRENT->{$CURRENT_SELECT} },
              [ 'option', $$hr{value}, [] ];
            delete $$hr{value};

        }
        elsif ( $TAG eq 'textarea' ) {
            my $stop = $start + $len;
            $$hr{value} = $$hr{'='};
            delete $$hr{'='};
            $$hr{name} = 'unknown' . $UNKNOWNS++ if ( !defined $$hr{name} );
            $key = $$hr{name};
            push @{ $_forms_CURRENT->{$key} }, [ 'textarea', $$hr{value}, [] ];
            delete $$hr{'name'};
            delete $$hr{'value'};

        }
        else {    # button
            $$hr{name}  = 'unknown' . $UNKNOWNS++ if ( !defined $$hr{name} );
            $$hr{value} = undef                   if ( !defined $$hr{value} );
            $key        = $$hr{name};
            push @{ $_forms_CURRENT->{$key} }, [ 'button', $$hr{value}, [] ];
            delete $$hr{'name'};
            delete $$hr{'value'};
        }

        if ( scalar %$hr ) {
            if ( $TAG eq 'form' ) { $parr = $_forms_CURRENT->{$key}->[3]; }
            else {
                $parr = $_forms_CURRENT->{$key}->[-1];
                $parr = $parr->[2];
            }

            my ( $k, $v );
            while ( ( $k, $v ) = each(%$hr) ) {
                if ( defined $v ) { push @$parr, "$k=\"$v\""; }
                else { push @$parr, $k; }
            }
        }

        return undef;
    }
}

