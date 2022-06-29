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

=item B<cookie_new_jar>

Params: none

Return: $jar

Create a new cookie jar, for use with the other functions.  Even though
the jar is technically just a hash, you should still use this function
in order to be future-compatible (should the jar format change).

=cut

sub cookie_new_jar {
    return {};
}

########################################################################

=item B<cookie_read>

Params: $jar, \%response [, \%request, $reject ]

Return: $num_of_cookies_read

Read in cookies from an %response hash, and put them in $jar.

Notice: cookie_read uses internal magic done by http_do_request
in order to read cookies regardless of 'Set-Cookie[2]' header
appearance.

If the optional %request hash is supplied, then it will be used to
calculate default host and path values, in case the cookie doesn't
specify them explicitly.  If $reject is set to 1, then the %request
hash values are used to calculate and reject cookies which are not
appropriate for the path and domains of the given request.

=cut

sub cookie_read {
    my ( $count, $jarref, $hrs, $hrq, $rej ) = ( 0, @_ );

    return 0 if ( !( defined $jarref && ref($jarref) ) );
    return 0 if ( !( defined $hrs   && ref($hrs) ) );
    return 0
      if (
        !(
            defined $$hrs{whisker}->{cookies}
            && ref( $$hrs{whisker}->{cookies} )
        )
      );

		my @opt;
		if(defined $hrq && ref($hrq)){
			push @opt, $hrq->{whisker}->{host};
			my $u = $hrq->{whisker}->{uri};
			$u=~s#/.*?$##;
			$u='/' if($u eq '');
			push @opt, $u, $rej;
		}

    foreach ( @{ $hrs->{whisker}->{cookies} } ) {
        cookie_parse( $jarref, $_ , @opt);
        $count++;
    }
    return $count;
}

########################################################################

=item B<cookie_parse>

Params: $jar, $cookie [, $default_domain, $default_path, $reject ]

Return: nothing

Parses the cookie into the various parts and then sets the appropriate 
values in the cookie $jar. If the cookie value is blank, it will delete 
it from the $jar.  See the 'docs/cookies.txt' document for a full
explanation of how Libwhisker parses cookies and what RFC aspects are 
supported.

The optional $default_domain value is taken literally.  Values with no 
leading dot (e.g. 'www.host.com') are considered to be strict hostnames 
and will only match the identical hostname.  Values with leading dots (e.g. 
'.host.com') are treated as sub-domain matches for a single domain level.
If the cookie does not indicate a domain, and a $default_domain is not
provided, then the cookie is considered to match all domains/hosts.

The optional $default_path is used when the cookie does not specify a path.
$default_path must be absolute (start with '/'), or it will be ignored.  If
the cookie does not specify a path, and $default_path is not provided, then
the default value '/' will be used.

Set $reject to 1 if you wish to reject cookies based upon the provided
$default_domain and $default_path.  Note that $default_domain and 
$default_path must be specified for $reject to actually do something 
meaningful.

=cut

sub cookie_parse {
    my ( $jarref, $header ) = (shift, shift);
		my ( $Dd, $Dp, $R ) = (shift, shift, shift||0);

    return if ( !( defined $jarref && ref($jarref) ) );
    return if ( !( defined $header && length($header) > 0 ) );

		my @C = ( undef, undef, undef, undef, 0 );

		$header =~ tr/\r\n//d;
		my ($f,%seen,$n,$t) = (1);
    while( length($header) ){
    	$header =~ s/^[ \t]+//;
    	last if(!($header =~ s/^([^ \t=;]+)//));
			# LW2.5 change: cookie name is no longer lower-cased
    	# my $an = lc($1);
    	my $an = $1;
			my $av = undef;
    	$header =~ s/^[ \t]+//;
    	if(substr($header,0,1) eq '='){
    		$header=~s/^=[ \t]*//;
    		if(substr($header,0,1) eq '"'){
    			my $p = index($header,'"',1);
    			last if($p == -1);
    			$av = substr($header,1,$p-1);
    			substr($header,0,$p+1)='';
    		} else {
					$av = $1 if($header =~ s/^([^ \t;,]*)//);
    		}
    	} else {
    		my $p = index($header,';');
    		substr($header,0,$p)='';
    	}
    	$header =~ s/^.*?;//;
			if($f){
				return if(!defined $av);
				($f,$n,$C[0])=(0,$an,$av);
			} else {
				$seen{$an}=$av if(!exists $seen{$an});
  		}
    }

		return if(!defined $n || $n eq '');
		my $del = 0;
		$del++ if($C[0] eq '');
		$del++ if(defined $seen{'max-age'} && $seen{'max-age'} eq '0');
		if($del){
        delete $$jarref{$n} if exists $$jarref{$n};			
        return;
		}

		if(defined $seen{domain} && $seen{domain} ne ''){
			$t = $seen{domain};
			$t='.'.$t if(substr($t,0,1) ne '.' && !_is_ip_address($t));
		} else {
			$t=$Dd;
		}
		$t=~s/\.+$// if(defined $t);
		$C[1]=$t;

		if(defined $seen{path}){
			$t = $seen{path};
		} else {
			$t=$Dp || '/';
		}
		$t=~s#/+$##;
		$t='/' if(substr($t,0,1) ne '/');
		$C[2]=$t;

		$C[4]=1 if(exists $seen{secure});

		return if($R && !_is_valid_cookie_match($C[1], $C[2], $Dd, $Dp));
    $$jarref{$n} = \@C;
}

########################################################################

sub _is_ip_address {
	my $n = shift;
	return 1 if($n=~/^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$/);
	return 0;
}

sub _is_valid_cookie_match {
	my ($cd, $cp, $td, $tp) = @_;
	return 0 if(index($tp,$cp)!=0);
	if(substr($cd,0,1) eq '.'){
		if( $td =~ /(.+)$cd$/ ){
			return 1 if(index($1,'.') == -1);
		}
		return 0;
	} else {
		return 0 if($cd ne $td);
	}
	return 1;
}

########################################################################

=item B<cookie_write>

Params: $jar, \%request, $override

Return: nothing

Goes through the given $jar and sets the Cookie header in %req pending the 
correct domain and path.  If $override is true, then the secure, domain and 
path restrictions of the cookies are ignored and all cookies are essentially
included.

Notice: cookie expiration is currently not implemented.  URL restriction
comparision is also case-insensitive.

=cut

sub cookie_write {
    my ( $jarref, $hin, $override ) = @_;
    my ( $name, $out ) = ( '', '' );

    return if ( !( defined $jarref && ref($jarref) ) );
    return if ( !( defined $hin    && ref($hin) ) );

    $override = $override || 0;
    $$hin{'whisker'}->{'ssl'} = $$hin{'whisker'}->{'ssl'} || 0;

    foreach $name ( keys %$jarref ) {
        next if ( $name eq '' );
        if($override){
            $out .= "$name=$$jarref{$name}->[0];";
            next;
        }
        next if ( $$hin{'whisker'}->{'ssl'} == 0 && $$jarref{$name}->[4] > 0 );
        if ( $$hin{'whisker'}->{'host'} =~ /$$jarref{$name}->[1]$/i
                && $$hin{'whisker'}->{'uri'} =~ /^$$jarref{$name}->[2])/ )
        {
            $out .= "$name=$$jarref{$name}->[0];";
        }
    }

    if ( $out ne '' ) { $$hin{'Cookie'} = $out; }

}

########################################################################

=item B<cookie_get>

Params: $jar, $name

Return: @elements

Fetch the named cookie from the $jar, and return the components.  The
returned items will be an array in the following order:

value, domain, path, expire, secure

value  = cookie value, should always be non-empty string
domain = domain root for cookie, can be undefined
path   = URL path for cookie, should always be a non-empty string
expire = undefined (depreciated, but exists for backwards-compatibility)
secure = whether or not the cookie is limited to HTTPs; value is 0 or 1

=cut

sub cookie_get {
    my ( $jarref, $name ) = @_;

    return undef if ( !( defined $jarref && ref($jarref) ) );

    if ( defined $$jarref{$name} ) {
        return @{ $$jarref{$name} };
    }

    return undef;
}

########################################################################

=item B<cookie_get_names>

Params: $jar

Return: @names

Fetch all the cookie names from the jar, which then let you cooke_get()
them individually.

=cut

sub cookie_get_names {
    my ( $jarref, $name ) = @_;

    return undef if ( !( defined $jarref && ref($jarref) ) );
    return keys %$jarref;
}

########################################################################

=item B<cookie_get_valid_names>

Params: $jar, $domain, $url, $ssl

Return: @names

Fetch all the cookie names from the jar which are valid for the given
$domain, $url, and $ssl values.  $domain should be string scalar of the
target host domain ('www.example.com', etc.).  $url should be the absolute 
URL for the page ('/index.html', '/cgi-bin/foo.cgi', etc.).  $ssl should be 
0 for non-secure cookies, or 1 for all (secure and normal) cookies.  The 
return value is an array of names compatible with cookie_get().

=cut

sub cookie_get_valid_names {
    my ( $jarref, $domain, $url, $ssl ) = @_;

    return () if ( !( defined $jarref && ref($jarref) ) );
		return () if ( !defined $domain || $domain eq '' );
		return () if ( !defined $url || $url eq '' );
		$ssl ||= 0;

		my (@r, $name);
    foreach $name ( keys %$jarref ) {
        next if ( $name eq '' );
        next if ( $$jarref{$name}->[4] > 0 && $ssl == 0 );
        if ( $domain =~ /$$jarref{$name}->[1]$/i
                && $url =~ /^$$jarref{$name}->[2])/i ) {
            push @r, $name;
        }
    }
    
    return @r;
}


########################################################################

=item B<cookie_set>

Params: $jar, $name, $value, $domain, $path, $expire, $secure

Return: nothing

Set the named cookie with the provided values into the %jar.  $name is 
required to be a non-empty string.  $value is required, and will delete
the named cookie from the $jar if it is an empty string.  $domain and
$path can be strings or undefined.  $expire is ignored (but exists
for backwards-compatibility).  $secure should be the numeric value of
0 or 1.

=cut

sub cookie_set {
    my ( $jarref, $name, $value, $domain, $path, $expire, $secure ) = @_;
    my @construct;

    return if ( !( defined $jarref && ref($jarref) ) );

    return if ( $name eq '' );
    if ( !defined $value || $value eq '' ) {
        delete $$jarref{$name};
        return;
    }
    $path   = $path   || '/';
    $secure = $secure || 0;

    @construct = ( $value, $domain, $path, undef, $secure );
    $$jarref{$name} = \@construct;
}

########################################################################

