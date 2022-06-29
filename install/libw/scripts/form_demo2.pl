#!perl
#
# form_demo2.pl
#
# This is a libwhisker sample script which shows how to use the forms
# functions, and parse the forms structure, to rewrite form elements.
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


$FORMS = LW2::forms_read( \$HTML );

foreach $key (keys %{ $FORMS->[0] }){

	# skip main <form> definition
	next if($key eq "\0");

	# basic data
	$type = lc($FORMS->[0]->{$key}->[0]->[0]);
	$value = lc($FORMS->[0]->{$key}->[0]->[1]);
	$attribs = lc($FORMS->[0]->{$key}->[0]->[2]);

	# turn <select> into <textarea> full of all options
	if($type eq 'select'){
		@values=();
		foreach ( @{ $FORMS->[0]->{$key} } ){
			next if($_->[0] ne 'option');
			push @values, $_->[1];
		}
		$value=join("\n",@values);
		$FORMS->[0]->{$key} = [[ 'textarea', $value, [] ]];
		next;
	}

	# turn <input radio> into <textarea> full of all options
	if($type eq 'input-radio'){
		@values=();
		foreach ( @{ $FORMS->[0]->{$key} } ){
			next if(lc($_->[0]) ne 'input-radio');
			push @values, $_->[1];
		}
		$value=join("\n",@values);
		$FORMS->[0]->{$key} = [[ 'textarea', $value, [] ]];
		next;
	}


	# <textareas> don't need changing
	if($type eq 'textarea'){
		$FORMS->[0]->{$key} = [[ 'textarea', $value, [] ]];
		next;
	}

	# turn <input hidden> into <input text>
	if($type eq 'input-hidden'){
		$FORMS->[0]->{$key} = [[ 'input-text', $value, [] ]];
		next;
	}

	# turn <input checkbox> into <input text> with 1 or 0 value
	if($type eq 'input-checkbox'){
		if(attributes_lookup($attribs,'checked')){
			$value = '1';
		} else {
			$value = '0';
		}
		$FORMS->[0]->{$key} = [[ 'input-text', $value, [] ]];
		next;
	}

	# fallthrough, delete attribs
	$FORMS->[0]->{$key} = [[ $type, $value, [] ]];
}

# print LW2::dump('form',$FORMS->[0] );
print LW2::forms_write( $FORMS->[0] );

########################################################

sub attributes_lookup {
	my ( $attribs_ref, $name ) = @_;

	foreach (@$attribs_ref){
		return $_ if($_ eq $name);
		return $_ if($_ =~ m/^$name=/);
	}
	return undef;
}
