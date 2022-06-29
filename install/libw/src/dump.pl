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

=item B<dump>

Params: $name, \@array [, $name, \%hash, $name, \$scalar ]

Return: $code [ undef on error ]

The dump function will take the given $name and data reference, and
will create an ASCII perl code representation suitable for eval'ing
later to recreate the same structure.  $name is the name of the variable
that it will be saved as.  Example:

 $output = LW2::dump('request',\%request);

NOTE: dump() creates anonymous structures under the name given.  For
example, if you dump the hash %hin under the name 'hin', then when you
eval the dumped code you will need to use %$hin, since $hin is now a
*reference* to a hash.

=cut

sub dump {
    my %what = @_;
    my ( $final, $k, $v ) = ('');
    while ( ( $k, $v ) = each %what ) {
        return undef if ( ref($k) || !ref($v) );
        $final .= "\$$k = " . _dump( 1, $v, 1 );
        $final =~ s#,\n$##;
        $final .= ";\n";
    }
    return $final;
}

########################################################################

=item B<dump_writefile>

Params: $file, $name, \@array [, $name, \%hash, $name, \@scalar ]

Return: 0 if success; 1 if error

This calls dump() and saves the output to the specified $file.  

Note: LW does not checking on the validity of the file name, it's
creation, or anything of the sort.  Files are opened in overwrite
mode.

=cut

sub dump_writefile {
    my $file   = shift;
    my $output = &dump(@_);
    return 1 if ( !open( OUT, ">$file" ) || !defined $output );
    binmode(OUT);
    print OUT $output;
    close(OUT);
}

########################################################################

sub _dump {    # dereference and dump an element
    my ( $t,   $ref, $depth ) = @_;
    my ( $out, $k,   $v )     = ('');
    $depth ||= 1;

    # to protect against circular loops
    return 'undef' if ( $depth > 128 );

    if ( !defined $ref ) {
        return 'undef';
    }
    elsif ( ref($ref) eq 'HASH' ) {
        $out .= "{\n";
        while ( ( $k, $v ) = each %$ref ) {
#            next if ( $k eq '' );
            $out .= "\t" x $t;
            $out .= _dumpd($k) . ' => ';
            if ( ref($v) ) { $out .= _dump( $t + 1, $v, $depth + 1 ); }
            else { $out .= _dumpd($v); }
            $out .= ",\n" unless ( substr( $out, -2, 2 ) eq ",\n" );
        }
        $out =~ s#,\n$#\n#;
        $out .= "\t" x ( $t - 1 );
        $out .= "},\n";
    }
    elsif ( ref($ref) eq 'ARRAY' ) {
        $out .= "[";
        if ( ~~@$ref ) {
            $out .= "\n";
            foreach $v (@$ref) {
                $out .= "\t" x $t;
                if ( ref($v) ) { $out .= _dump( $t + 1, $v, $depth + 1 ); }
                else { $out .= _dumpd($v); }
                $out .= ",\n" unless ( substr( $out, -2, 2 ) eq ",\n" );
            }
            $out =~ s#,\n$#\n#;
            $out .= "\t" x ( $t - 1 );
        }
        $out .= "],\n";
    }
    elsif ( ref($ref) eq 'SCALAR' ) {
        $out .= _dumpd($$ref);
    }
    elsif ( ref($ref) eq 'REF' ) {
        $out .= _dump( $t, $$ref, $depth + 1 );
    }
    elsif ( ref($ref) ) {    # unknown/unsupported ref
        $out .= "undef";
    }
    else {                   # normal scalar
        $out .= _dumpd($ref);
    }
    return $out;
}

########################################################################

sub _dumpd {                 # escape a scalar string
    my $v = shift;
    return 'undef' if ( !defined $v );
    return "''"    if ( $v eq '' );
    return "$v"    if ( $v eq '0' || $v !~ tr/0-9//c && $v !~ m#^0+# );
    if ( $v !~ tr/ !-~//c ) {
        $v =~ s/(['\\])/\\$1/g;
        return "'$v'";
    }
    $v =~ s#\\#\\\\#g;
    $v =~ s#"#\\"#g;
    $v =~ s#\r#\\r#g;
    $v =~ s#\n#\\n#g;
    $v =~ s#\t#\\t#g;
    $v =~ s#\$#\\\$#g;
    $v =~ s#([^!-~ ])#sprintf('\\x%02x',ord($1))#eg;
    return "\"$v\"";
}

########################################################################

