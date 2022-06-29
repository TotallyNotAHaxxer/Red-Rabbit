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

=item B<time_mktime>

Params: $seconds, $minutes, $hours, $day_of_month, $month, $year_minus_1900

Return: $seconds [ -1 on error ]

Performs a general mktime calculation with the given time components.
Note that the input parameter values are expected to be in the format
output by localtime/gmtime.  Namely, $seconds is 0-60 (yes, there can
be a leap second value of 60 occasionally), $minutes is 0-59, $hours
is 0-23, $days is 1-31, $month is 0-11, and $year is 70-127.  This
function is limited in that it will not process dates prior to 1970 or
after 2037 (that way 32-bit time_t overflow calculations aren't required).

Additional parameters passed to the function are ignored, so it is
safe to use the full localtime/gmtime output, such as:

	$seconds = LW2::time_mktime( localtime( time ) );

Note: this function does not adjust for time zone, daylight savings
time, etc.  You must do that yourself.

=cut

sub time_mktime {
	my ($sec,$min,$hour,$day,$mon,$yr)=@_;
	my @md=(0,31,59,90,120,151,181,212,243,273,304,334);
	foreach(@_[0..5]){
		return -1 if !defined $_ || $_<0; }
	return -1 if($sec>60 || $min>59 || $hour>23 || $day>31 || $mon>11
		|| $yr>137 || $yr<70);
	$yr += 1900;
	my $res = ($yr-1970)*365+$md[$mon];
	$res += int(($yr-1969)/4) + int(($yr-1601)/400);
	$res -= int(($yr-1901)/100);
	$res = ($res+$day-1)*24;
	$res = ($res+$hour)*60;
	$res = ($res+$min)*60;
	return $res+$sec;
}


=item B<time_gmtolocal>

Params: $seconds_gmt

Return: $seconds_local_timezone

Takes a seconds value in UTC/GMT time and adjusts it to reflect the current
timezone.  This function is slightly expensive; it takes the gmtime() and
localtime() representations of the current time, calculates the delta 
difference by turning them back into seconds via time_mktime, and then 
applies this delta difference to $seconds_gmt.

Note that if you give this function a time and subtract the return value from
the original time, you will get the delta value.  At that point, you can just
apply the delta directly and skip calling this function, which is a massive
performance boost.  However, this will cause problems if you have a long
running program which crosses daylight savings time boundaries, as the DST
adjustment will not be accounted for unless you recalculate the new delta.

=cut

sub time_gmtolocal {
	my $t = shift;
	my $now = time;
	my $utc = time_mktime(gmtime($now));
	my $me = time_mktime(localtime($now));
	return $t - ($utc - $me);
}

