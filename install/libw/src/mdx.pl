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

{    # start md5 packaged varbs
    my ( @S, @T, @M );
    my $code           = '';
    my $MD5_TRYLOADING = 1;

=item B<md5>

Params: $data

Return: $hex_md5_string

This function takes a data scalar, and composes a MD5 hash of it, and 
returns it in a hex ascii string.  It will use the fastest MD5 function
available.

=cut

    sub md5 {
        return undef if ( !defined $_[0] );    # oops, forgot the data
        if ($MD5_TRYLOADING) {
            $MD5_TRYLOADING = 0;
            eval "require MD5";
        }
        return MD5->hexhash( $_[0] ) if ($MD5::VERSION);
        my $DATA = _md5_pad( $_[0] );
        &_md5_init() if ( !defined $M[0] );
        return _md5_perl_generated( \$DATA );
    }

########################################################################

    sub _md5_init {
        return if ( defined $S[0] );
        my $i;
        for ( $i = 1 ; $i <= 64 ; $i++ ) {
            $T[ $i - 1 ] = int( ( 2**32 ) * abs( sin($i) ) );
        }
        my @t = ( 7, 12, 17, 22, 5, 9, 14, 20, 4, 11, 16, 23, 6, 10, 15, 21 );
        for ( $i = 0 ; $i < 64 ; $i++ ) {
            $S[$i] = $t[ ( int( $i / 16 ) * 4 ) + ( $i % 4 ) ];
        }
        @M = (
            0, 1, 2,  3,  4,  5,  6,  7,  8,  9,  10, 11, 12, 13, 14, 15,
            1, 6, 11, 0,  5,  10, 15, 4,  9,  14, 3,  8,  13, 2,  7,  12,
            5, 8, 11, 14, 1,  4,  7,  10, 13, 0,  3,  6,  9,  12, 15, 2,
            0, 7, 14, 5,  12, 3,  10, 1,  8,  15, 6,  13, 4,  11, 2,  9
        );
        &_md5_generate();

        # check to see if it works correctly
        my $TEST = _md5_pad('foobar');
        if ( _md5_perl_generated( \$TEST ) ne
            '3858f62230ac3c915f300c664312c63f' )
        {
            utils_carp('md5: MD5 self-test not successful.');
        }
    }

########################################################################

    # This function is from Digest::Perl::MD5, and bears the following
    # copyrights:
    #
    # Copyright 2000 Christian Lackas, Imperia Software Solutions
    # Copyright 1998-1999 Gisle Aas.
    # Copyright 1995-1996 Neil Winton.
    # Copyright 1991-1992 RSA Data Security, Inc.
    #

    sub _md5_pad {
        my $l = length( my $msg = shift() . chr(128) );
        $msg .= "\0" x ( ( $l % 64 <= 56 ? 56 : 120 ) - $l % 64 );
        $l = ( $l - 1 ) * 8;
        $msg .= pack 'VV', $l & 0xffffffff, ( $l >> 16 >> 16 );
        return $msg;
    }

########################################################################

    sub _md5_generate {
        my $N = 'abcddabccdabbcda';
        my ( $i, $M ) = ( 0, '' );
        $M = '&0xffffffff' if ( ( 1 << 16 ) << 16 );    # mask for 64bit systems

        $code = <<EOT;
        sub _md5_perl_generated {
	BEGIN { \$^H |= 1; }; # use integer
        my (\$A,\$B,\$C,\$D)=(0x67452301,0xefcdab89,0x98badcfe,0x10325476);
        my (\$a,\$b,\$c,\$d,\$t,\$i);
        my \$dr=shift;
        my \$l=length(\$\$dr);
        for my \$L (0 .. ((\$l/64)-1) ) {
                my \@D = unpack('V16', substr(\$\$dr, \$L*64,64));
                (\$a,\$b,\$c,\$d)=(\$A,\$B,\$C,\$D);
EOT

        for ( $i = 0 ; $i < 16 ; $i++ ) {
            my ( $a, $b, $c, $d ) =
              split( '', substr( $N, ( $i % 4 ) * 4, 4 ) );
            $code .=
              "\$t=((\$$d^(\$$b\&(\$$c^\$$d)))+\$$a+\$D[$M[$i]]+$T[$i])$M;\n";
            $code .=
"\$$a=(((\$t<<$S[$i])|((\$t>>(32-$S[$i]))&((1<<$S[$i])-1)))+\$$b)$M;\n";
        }
        for ( ; $i < 32 ; $i++ ) {
            my ( $a, $b, $c, $d ) =
              split( '', substr( $N, ( $i % 4 ) * 4, 4 ) );
            $code .=
              "\$t=((\$$c^(\$$d\&(\$$b^\$$c)))+\$$a+\$D[$M[$i]]+$T[$i])$M;\n";
            $code .=
"\$$a=(((\$t<<$S[$i])|((\$t>>(32-$S[$i]))&((1<<$S[$i])-1)))+\$$b)$M;\n";
        }
        for ( ; $i < 48 ; $i++ ) {
            my ( $a, $b, $c, $d ) =
              split( '', substr( $N, ( $i % 4 ) * 4, 4 ) );
            $code .= "\$t=((\$$b^\$$c^\$$d)+\$$a+\$D[$M[$i]]+$T[$i])$M;\n";
            $code .=
"\$$a=(((\$t<<$S[$i])|((\$t>>(32-$S[$i]))&((1<<$S[$i])-1)))+\$$b)$M;\n";
        }
        for ( ; $i < 64 ; $i++ ) {
            my ( $a, $b, $c, $d ) =
              split( '', substr( $N, ( $i % 4 ) * 4, 4 ) );
            $code .= "\$t=((\$$c^(\$$b|(~\$$d)))+\$$a+\$D[$M[$i]]+$T[$i])$M;\n";
            $code .=
"\$$a=(((\$t<<$S[$i])|((\$t>>(32-$S[$i]))&((1<<$S[$i])-1)))+\$$b)$M;\n";
        }

        $code .= <<EOT;
                \$A=\$A+\$a\&0xffffffff; \$B=\$B+\$b\&0xffffffff;
                \$C=\$C+\$c\&0xffffffff; \$D=\$D+\$d\&0xffffffff;
        } # for
	return unpack('H*', pack('V4',\$A,\$B,\$C,\$D)); }
EOT
        eval "$code";
    }

}    # md5 package container

########################################################################

{    # start md4 packaged varbs
    my ( @S, @T, @M );
    my $code = '';

=item B<md4>

Params: $data

Return: $hex_md4_string

This function takes a data scalar, and composes a MD4 hash of it, and 
returns it in a hex ascii string.  It will use the fastest MD4 function
available.

=cut

    sub md4 {
        return undef if ( !defined $_[0] );    # oops, forgot the data
        my $DATA = _md5_pad( $_[0] );
        &_md4_init() if ( !defined $M[0] );
        return _md4_perl_generated( \$DATA );
    }

########################################################################

    sub _md4_init {
        return if ( defined $S[0] );
        my $i;
        my @t = ( 3, 7, 11, 19, 3, 5, 9, 13, 3, 9, 11, 15 );
        for ( $i = 0 ; $i < 48 ; $i++ ) {
            $S[$i] = $t[ ( int( $i / 16 ) * 4 ) + ( $i % 4 ) ];
        }
        @M = (
            0, 1, 2, 3,  4, 5,  6, 7,  8, 9, 10, 11, 12, 13, 14, 15,
            0, 4, 8, 12, 1, 5,  9, 13, 2, 6, 10, 14, 3,  7,  11, 15,
            0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5,  13, 3,  11, 7,  15
        );

        my $N = 'abcddabccdabbcda';
        my $M = '';
        $M = '&0xffffffff' if ( ( 1 << 16 ) << 16 );    # mask for 64bit systems

        $code = <<EOT;
        sub _md4_perl_generated {
	BEGIN { \$^H |= 1; }; # use integer
        my (\$A,\$B,\$C,\$D)=(0x67452301,0xefcdab89,0x98badcfe,0x10325476);
        my (\$a,\$b,\$c,\$d,\$t,\$i);
        my \$dr=shift;
        my \$l=length(\$\$dr);
        for my \$L (0 .. ((\$l/64)-1) ) {
                my \@D = unpack('V16', substr(\$\$dr, \$L*64,64));
                (\$a,\$b,\$c,\$d)=(\$A,\$B,\$C,\$D);
EOT

        for ( $i = 0 ; $i < 16 ; $i++ ) {
            my ( $a, $b, $c, $d ) =
              split( '', substr( $N, ( $i % 4 ) * 4, 4 ) );
            $code .= "\$t=((\$$d^(\$$b\&(\$$c^\$$d)))+\$$a+\$D[$M[$i]])$M;\n";
            $code .=
"\$$a=(((\$t<<$S[$i])|((\$t>>(32-$S[$i]))&((1<<$S[$i])-1))))$M;\n";
        }
        for ( ; $i < 32 ; $i++ ) {
            my ( $a, $b, $c, $d ) =
              split( '', substr( $N, ( $i % 4 ) * 4, 4 ) );
            $code .=
"\$t=(( (\$$b&\$$c)|(\$$b&\$$d)|(\$$c&\$$d) )+\$$a+\$D[$M[$i]]+0x5a827999)$M;\n";
            $code .=
"\$$a=(((\$t<<$S[$i])|((\$t>>(32-$S[$i]))&((1<<$S[$i])-1))))$M;\n";
        }
        for ( ; $i < 48 ; $i++ ) {
            my ( $a, $b, $c, $d ) =
              split( '', substr( $N, ( $i % 4 ) * 4, 4 ) );
            $code .=
              "\$t=(( \$$b^\$$c^\$$d )+\$$a+\$D[$M[$i]]+0x6ed9eba1)$M;\n";
            $code .=
"\$$a=(((\$t<<$S[$i])|((\$t>>(32-$S[$i]))&((1<<$S[$i])-1))))$M;\n";
        }

        $code .= <<EOT;
                \$A=\$A+\$a\&0xffffffff; \$B=\$B+\$b\&0xffffffff;
                \$C=\$C+\$c\&0xffffffff; \$D=\$D+\$d\&0xffffffff;
        } # for
	return unpack('H*', pack('V4',\$A,\$B,\$C,\$D)); }
EOT
        eval "$code";

        my $TEST = _md5_pad('foobar');
        if ( _md4_perl_generated( \$TEST ) ne
            '547aefd231dcbaac398625718336f143' )
        {
            utils_carp('md4: MD4 self-test not successful.');
        }
    }

}    # md4 package container

