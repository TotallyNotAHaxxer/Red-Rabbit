use strict;
use warnings;
use feature 'say';
use POSIX;
use Image::ExifTool ':Public';
use Data::HexDump;
use String::CRC32;

use GD;


no  warnings 'redefine';


my $injection_method = shift;
my $filename         = shift;
my $payload          = shift;
my $chunk            = shift;
my $pixelheight      = shift; 
my $pixelwidth       = shift;

my @pngtil = (
    (pack 'I>', length $payload),
    "\x50\x55\x6e\x4b",
    $payload,
    (pack 'I>', crc32('IEND' . $payload)),
    "\x00IEND",
);

my @build_hex_JPG = (
    "\xff\xd8",
    "\xff\xdb",
    pack('S>', 67),
    "\x00" . "\x01" x 64,
    "\xff\xc2",
    "\x00\x0b",
    "\x08\x00\x01\x00\x01\x01\x01\x11\x00",
    "\xff\xc4",
    "\x00\x14",
    "\x00\x01\x00\x00\x00\x00\x00\x00\x00"."\x00\x00\x00\x00\x00\x00\x00\x00\x03",
    "\xff\xda",
    "\x00\x08",
    "\x01\x01\x00\x00\x00\x01\x3f",
    "\xff\xd9",
);

# hex dumping 

sub hexdump {
    my $hd = Data::HexDump->new();
    $hd->file($filename);
    print while $_ = $hd->dump;
}


#
# JPG injection
#

sub create_jpg {
    say "<RR6> Stego Module: Generating output file";
    sysopen my $fh, $filename, O_CREAT|O_WRONLY;
    for my $hexc (@build_hex_JPG) {
        syswrite $fh, $hexc
    }
    close $fh;
    say "<RR6> Stego Module: File saved to: $filename";
}

sub inject_payload_com {
    say "<RR6> Stego Module: Injecting payload into COMMENT";
    my $exifTool = Image::ExifTool->new;
    $exifTool->SetNewValue('Comment', $payload)or die "<RR6> Stego Module: Failed to Set New Value\n";
    $exifTool->WriteInfo($filename)or die "<RR6> Stego Module:  Fail to WriteInfo\n";
    say "<RR6> Stego Module: Payload was injected successfully\n";
}

sub inject_payload_dqt {
    say "<RR6> Stego Module: Injecting payload into DQT table";
    my $payload_len = length $payload;
    sysopen my $fh, $filename, O_WRONLY;
    sysseek    $fh, (7 + (64 - $payload_len)), SEEK_SET;
    syswrite   $fh, $payload;
    close      $fh;
    say "<RR6> Stego Module: Payload was injected succesfully\n";
}


sub main_injection_JPG {
    &create_jpg if ! -f $filename;
    if (uc $chunk eq 'COM') {
        &create_jpg if ! -f $filename;
        &inject_payload_com;
    }
    elsif (uc $chunk eq 'DQT') {
        die "The payload size must not exceed 64 bytes!\n" if length($payload) > 64;
        &create_jpg; 
        &inject_payload_dqt;
    }
    else {
        die "section to inject argument must be COM or DQT!\n";
    }
    hexdump();
    say "<RR6> Stego Module: Payload injection into $filename was injected without fail"
}

#
# WEBP Injection
#

sub create_webp {
    say "<RR6> Stego Create: Generating output file";
    my $minimal_webp = "\x52\x49\x46\x46\x1a\x00\x00\x00"."\x57\x45\x42\x50\x56\x50\x38\x4c"."\x0d\x00\x00\x00\x2f\x00\x00\x00"."\x10\x07\x10\x11\x11\x88\x88\xfe"."\x07\x00";
    sysopen my $fh, $filename, O_CREAT|O_WRONLY;
    syswrite   $fh, $minimal_webp;
    close      $fh;
    say "<RR6> Stego Create: File saved to: $filename \n";
}

sub inject_payload_webp {
    say "<RR6> Stego Module: Injecting payload into $filename";
    sysopen my $fh, $filename, O_WRONLY;
    sysseek    $fh, 4, SEEK_SET;
    syswrite   $fh, "\x2f\x2a";
    sysseek    $fh, 15, SEEK_SET;
    syswrite   $fh, "\x4c\xff\xff\xff";
    sysseek    $fh, 0, SEEK_END;
    syswrite   $fh, "\x2a\x2f\x3d\x31\x3b";
    syswrite   $fh, $payload;
    syswrite   $fh, "\x3b";
    close      $fh;
    say "<RR6> Stego Module: Payload was injected successfully\n";
}

sub main_webp_injection {
    if (-f $filename) {
        say "<RR6> Stego Module: File exists, injecting $filename";
    } 
    else {
        say "<RR6> Stego Module: $filename does not exist. Creating a new blank WEBP file named $filename";
        &create_webp;
    }
    &inject_payload_webp;
}


#
# GIF format injection 
#

# needed this if statement here
sub create_gif {
    say "<RR6> Stego Module: Generating output file";
    if ($pixelheight eq "" or $pixelwidth eq "") {
        say "<RR6> Stego Module ERR -> WARNING: You did not specify a pixel width or height, using standard 476px 280px";
        my $img = GD::Image->new(1200,800,1,);
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->gif;
        close      $fh;
        say "<RR6> Stego Module: File saved as $filename \n";
    } 
    else {
        say "<RR6> Stego Module: Using current pixel height and width | $pixelheight - $pixelwidth | ";
        my $img = GD::Image->new(
            $pixelwidth,
            $pixelheight,
        1,
        );
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->gif;
        close      $fh;
        say "<RR6> Stego Module: File saved as -> $filename \n";
    }
}

sub inject_payload_GIF {
    say "<RR6> Stego Module: Injecting payload into $filename";
    sysopen my $fh, $filename, O_WRONLY;
    sysseek    $fh, 6, SEEK_SET;
    syswrite   $fh, "\x2f\x2a";
    sysseek    $fh, 0, SEEK_END;
    syswrite   $fh, "\x2a\x2f\x3d\x31\x3b";
    syswrite   $fh, $payload;
    syswrite   $fh, "\x3b";
    close      $fh;
    say "<RR6> Stego Module: Payload was injected successfully\n";
}



sub main_inject_GIF {
    if (-f $filename) {
        say "<RR6> Stego Module: File exists, no need to create a new one"
    } 
    else {
        say "<RR6> Stego Module: File does NOT exist, creating new GIF";
        &create_gif;
    }
    say "<RR6> Stego Module: Payload injected -> $payload";
    &inject_payload_GIF;
    &hexdump;
}

#
# BMP image injection
#

sub create_BMP {
    say "<RR6> Stego Module: [BMP] Generating output file";
    my $bmp_minimal = "\x42\x4d\x1e\x00\x00\x00\x00\x00\x00\x00\x1a\x00"."\x00\x00\x0c\x00\x00\x00\x01\x00\x01\x00\x01\x00"."\x18\x00\x00\x00\xff\x00";
    sysopen my $fh, $filename, O_CREAT|O_WRONLY;
    syswrite   $fh, $bmp_minimal;
    close      $fh;
    say "<RR6> Stego Module - File saved as: $filename\n";
}

sub inject_payload_BMP {
    say "<RR6> Stego Module: Injecting payload into $filename";
    sysopen my $fh, $filename, O_RDWR;
    sysseek    $fh, 2, SEEK_SET;
    syswrite   $fh, "\x2f\x2a";
    sysseek    $fh, 0, SEEK_END;
    syswrite   $fh, "\x2a\x2f\x3d\x31\x3b";
    syswrite   $fh, $payload;
    syswrite   $fh, "\x3b";
    close      $fh;
    say "<RR6> Stego Module: Payload [ $payload ] was injected successfully\n";
}


sub main_inject_BMP {
    say "<RR6> Stego Module: Injecting $payload into $filename";
    if (-f $filename) {
        say "<RR6> Stego Module: $filename exists, no reason to make a new BMP file";
    } 
    else {
        say "<RR6> Stego Module: $filename Does NOT exist, creating BMP file";
        &create_BMP;
    }
    &inject_payload_BMP;
    &hexdump;
}


#
# PNG IMAGE INJECTION
#

sub systell {
    sysseek $_[0], 0, SEEK_CUR
}

sub rewind {
    sysseek $_[0], systell($_[0]) - $_[1], SEEK_CUR
}

sub create_PNG {
    say "<RR6> Stego Module: [PNG] Generating new file | reason -> $filename does not exist ";
    if ($pixelheight eq "" or $pixelwidth eq "") {
        say "<RR6> Stego Module: [PNG] Pixel width or height was not defined using defualt";
        my $img = GD::Image->new(860, 881, 1,);
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->png;
        close      $fh;
        say "<RR6> Stego Module: [PNG] File saved as: $filename\n";
    }
    else {
        say "<RR6> Stego Module: [PNG] Pixel width or height was defined | $pixelwidth - $pixelheight | ";
        my $img = GD::Image->new($pixelwidth, $pixelheight, 1,);
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->png;
        close      $fh;
        say "<RR6> Stego Module: [PNG] File saved as -> $filename\n";
    }
}

sub inject_payload_PNG {
    say "<RR6> Stego Module: [PNG] Injecting payload < $payload > into < $filename > \n";
    sysopen our $fh, $filename, O_RDWR;
    {
        my $format;
        sysseek     $fh, 1, SEEK_SET;
        sysread     $fh, $format, 3;
        die "<RR6> Stego Module: [PNG] USER ERROR -> WARNING -> FATAL : $filename is not a PNG file and does not have the values to be one" if $format ne "PNG";
    }
    sysseek     $fh, 8, SEEK_SET;
    sub read_chunks {
        *read_next_chunk = \&read_chunks;
        my ($chunk_size, $chunk_type, $content, $crc);
        sysread $fh, $chunk_size, 4;
        sysread $fh, $chunk_type, 4;
        $chunk_size = unpack('I>', $chunk_size);
        say "[+] Chunk size: $chunk_size";
        say "[+] Chunk type: $chunk_type";
        return if $chunk_type eq 'IEND';
        sysread $fh, $content, $chunk_size;
        sysread $fh, $crc, 4;
        say '[+] CRC: ' . unpack('H8', $crc);
        &read_next_chunk;
    }
    &read_chunks;
    rewind   $fh, 8;
    say "\n<RR6> Stego Module: [PNG] Inject payload to the new chunk: 'pUnk'";
    for my $edit (@pngtil) {
        syswrite $fh, $edit;
    }
    close    $fh;
    say "<RR6> Stego Module: [PNG] Payload was injected successfully\n";
}



sub main_inject_PNG {
    if (-f $filename) {
        say "<RR6> Stego Module: [PNG] Filename exists, no need to create a new one";
    }
    else {
        say "<RR6> Stego Module: [PNG] Filename does NOT exist, creating a new one";
        &create_PNG;
    }
    &inject_payload_PNG;
}

# run main function
sub main {
    if ($injection_method eq "png") {
        &main_inject_PNG;
    } 
    if ($injection_method eq "bmp") {
        &main_inject_BMP;
    }
    if ($injection_method eq "webp") {
        &main_webp_injection
    }
    if ($injection_method eq "gif") {
        &main_inject_GIF
    }
    if ($injection_method eq "jpg") {
        &main_injection_JPG;
    }
}

main;