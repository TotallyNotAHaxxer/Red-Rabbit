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

=item B<auth_brute_force>

Params: $auth_method, \%req, $user, \@passwords [, $domain, $fail_code ]

Return: $first_valid_password, undef if error/none found

Perform a HTTP authentication brute force against a server (host and URI 
defined in %req).  It will try every password in the password array for 
the given user.  The first password (in conjunction with the given user) 
that doesn't return HTTP 401 is returned (and the brute force is stopped 
at that point).  You should retry the request with the given password and
double-check that you got a useful HTTP return code that indicates
successful authentication (200, 302), and not something a bit more 
abnormal (407, 500, etc).  $domain is optional, and is only used for NTLM
auth.

Note: set up any proxy settings and proxy auth in %req before calling
this function.

You can brute-force proxy authentication by setting up the target proxy
as proxy_host and proxy_port in %req, using an arbitrary host and uri
(preferably one that is reachable upon successful proxy authorization),
and setting the $fail_code to 407.  The $auth_method passed to this
function should be a proxy-based one ('proxy-basic', 'proxy-ntlm', etc).

if your server returns something other than 401 upon auth failure, then
set $fail_code to whatever is returned (and it needs to be something
*different* than what is received on auth success, or this function
won't be able to tell the difference).

=cut

sub auth_brute_force {
    my ( $auth_method, $hrin, $user, $pwordref, $dom, $fail_code ) = @_;
    my ( $P, %hout );
    $fail_code ||= 401;

    return undef if ( !defined $auth_method || length($auth_method) == 0 );
    return undef if ( !defined $user        || length($user) == 0 );
    return undef if ( !( defined $hrin     && ref($hrin) ) );
    return undef if ( !( defined $pwordref && ref($pwordref) ) );

    map {
        ( $P = $_ ) =~ tr/\r\n//d;
        auth_set( $auth_method, $hrin, $user, $P, $dom );
        return undef if ( http_do_request( $hrin, \%hout ) );
        return $P if ( $hout{whisker}->{code} != $fail_code );
    } @$pwordref;

    return undef;
}

########################################################################

=item B<auth_unset>

Params: \%req

Return: nothing (modifies %req)

Modifes %req to disable all authentication (regular and proxy).

Note: it only removes the values set by auth_set().  Manually-defined
[Proxy-]Authorization headers will also be deleted (but you shouldn't 
be using the auth_* functions if you're manually handling your own auth...)

=cut

sub auth_unset {
    my $href = shift;
    return if ( !defined $href || !ref($href) );
    delete $$href{Authorization};
    delete $$href{'Proxy-Authorization'};
    delete $$href{whisker}->{auth_callback};
    delete $$href{whisker}->{auth_proxy_callback};
    delete $$href{whisker}->{auth_data};
    delete $$href{whisker}->{auth_proxy_data};
}

########################################################################

=item B<auth_set>

Params: $auth_method, \%req, $user, $password [, $domain]

Return: nothing (modifies %req)

Modifes %req to use the indicated authentication info.

Auth_method can be: 'basic', 'proxy-basic', 'ntlm', 'proxy-ntlm'.

Note: this function may not necessarily set any headers after being called.
Also, proxy-ntlm with SSL is not currently supported.

=cut

sub auth_set {
    my ( $method, $href, $user, $pass, $domain ) = ( lc(shift), @_ );

    return if ( !( defined $href && ref($href) ) );
    return if ( !defined $user || !defined $pass );

    if ( $method eq 'basic' ) {
        $$href{'Authorization'} =
          'Basic ' . encode_base64( $user . ':' . $pass, '' );
    }

    if ( $method eq 'proxy-basic' ) {
        $$href{'Proxy-Authorization'} =
          'Basic ' . encode_base64( $user . ':' . $pass, '' );
    }

    if ( $method eq 'ntlm' ) {
        http_close($href);
        $$href{whisker}->{auth_data} = ntlm_new( $user, $pass, $domain );
        $$href{whisker}->{auth_callback} = \&_ntlm_auth_callback;
    }

    if ( $method eq 'proxy-ntlm' ) {
        utils_croak('',"auth_set: proxy-ntlm auth w/ SSL not currently supported")
          if ( $href->{whisker}->{ssl} > 0 );
        http_close($href);
        $$href{whisker}->{auth_proxy_data} = ntlm_new( $user, $pass, $domain );
        $$href{whisker}->{auth_proxy_callback} = \&_ntlm_auth_proxy_callback;
    }

}

