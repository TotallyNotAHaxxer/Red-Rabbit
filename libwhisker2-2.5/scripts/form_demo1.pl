#!perl
#
# form_demo1.pl
#
# This is an example libwhisker script which shows how to parse the
# forms structure.  This program describes a form.
#

use LW2;

$HTML =<<EOT;

<form method="POST" action="/test.cgi">
<input type="text" name="first-input" value="one" onchange="foo();">
<input type="text" name="second-input" value="two" onchange="bar();">
<input type="text" value="three">
<input type="checkbox" name="first-check" value="mycheck">
<input type="radio" name="the-radio" value="1">
<input type="radio" name="the-radio" value="2" selected>
<input type="radio" name="the-radio" value="3">
<textarea name="areatext">my text</textarea>
<select name="myselect">
<option>1</option>
<option>2</option>
<option selected>3</option>
</select>
<input type="submit">
</form>

EOT


# read the form(s)
$FOUND = LW2::forms_read( \$HTML );

# $FOUND is an array of found forms; let's just look at the first one
$FIRST_FORM = $FOUND->[0];

# first we'll deal with the <form> data
print "Form name is: ", $FIRST_FORM->{"\0"}->[0], "\n";
print "Form method is: ", $FIRST_FORM->{"\0"}->[1], "\n";
print "Form action is: ", $FIRST_FORM->{"\0"}->[2], "\n";
print "\n";

# now we just enumerate the keys and handle them
foreach $key (keys %$FIRST_FORM){

	# is this an unnamed key?
	$unnamed = 0;
	if($key =~ m/^unknown[0-9]+$/){
		$unnamed++; }

	# check the element type
	$type = $FIRST_FORM->{$key}->[0]->[0];

	# handle multi-element selects
	if($type eq 'select'){

		# storage for the selected value
		$selected_value=undef;

		# possible values
		@possible_values = ();

		# loop through each element
		foreach ( @{ $FIRST_FORM->{$key} }){

			# we only care about options at this point
			next if($_->[0] ne 'option');

			# the first option will be the selected value if
			# an actual flagged as 'selected' value is not found
			$selected_value = $_->[1] if(!defined $selected_value);

			# check to see if this value is selected
			if(attributes_lookup($_->[2],'selected')){
				$selected_value = $_->[1]; }

			# save this value as a possible value
			push @possible_values, $_->[1];
		}

		if($unnamed){
			print "[$key]\tis an unnamed select with value ",
				$selected_value;
		} else {
			print "[$key]\tis a select with value ",
				$selected_value;
		}

		if(attributes_lookup($attribs,'onchange')){
			print " (w/ javascript onchange())";
		}
		print "\n";

		print "\t\t- Possible values are: ",
			join(',', @possible_values), "\n";

		# next element
		next;
	}

	# handle radio boxes
	if(lc($type) eq 'input-radio'){

		# storage for the selected value
		$selected_value=undef;

		# possible values
		@possible_values = ();

		# loop through each element
		foreach $element ( @{ $FIRST_FORM->{$key} }){
			
			# another way to deference things
			$type = $element->[0];
			$value = $element->[1];
			$attribs = $element->[2];

			# we only care about input-radios
			next if(lc($type) ne 'input-radio');

			# check to see if this value is selected
			if(attributes_lookup($attribs,'selected')){
				$selected_value = $value; }

			# save this value as a possible value
			push @possible_values, $value;
		}

		if($unnamed){
			print "[$key]\tis an unnamed radio input ";
		} else {
			print "[$key]\tis a radio input ";
		}
		
		if(defined $selected_value){
			print "with value $selected_value";
		} else {
			print "with no default value";
		}

		if(attributes_lookup($attribs,'onchange')){
			print " (w/ javascript onchange())";
		}
		print "\n";

		print "\t\t- Possible values are: ",
			join(',', @possible_values), "\n";

		next;
	}


	# handle the other simple types
	foreach $element ( @{ $FIRST_FORM->{$key} } ){

		$type = $element->[0];
		$value = $element->[1];
		$attribs = $element->[2];

		if( $type eq 'textarea'){
			print "[$key]\tis textarea with ",
				length( $value ),
				" chars in default value\n";

		} elsif ($type =~ m/^input-(.+)/){
			$input_type = lc($1);

			if($unnamed){
				print "[$key]\tis an unnamed $input_type input";
			} else {
				print "[$key]\tis a $input_type input";
			}

			if($input_type eq 'text' || $input_type eq 'hidden'){
				print " with value '$value'";

			} elsif($input_type eq 'submit'){

			} elsif($input_type eq 'radio'){
				if(attributes_lookup($attribs,'selected')){
					print " with actual value '$value'";
				} else {
					print " with possible value '$value'";
				}

			} elsif($input_type eq 'checkbox'){
				if(attributes_lookup($attribs,'checked')){
					print " which is checked";
				} else {
					print " which is not checked";
				}
			}
		
			if(attributes_lookup($attribs,'onchange')){
				print " (w/ javascript onchange())";
			}
			print "\n";
		}
	}


}

sub attributes_lookup {
	my ( $attribs_ref, $name ) = @_;

	foreach (@$attribs_ref){
		return $_ if($_ eq $name);
		return $_ if($_ =~ m/^$name=/);
	}
	return undef;
}
