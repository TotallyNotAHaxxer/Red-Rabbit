#/usr/bin/perl
#
# OG code            => ArkAngeL43, Perl5, 
# Most used function => push, for formatting a list of values
# List               => -e filename, optsgeotp
# 
#
# Used a list of features to determine what the file is, and the data inside of the file 
#
# List from the offical perl5 documentation 
#
#List __START__:
#
            #-A
            #Script start time minus file last access time, in days.
            #
            #-B
            #Is it a binary file?
            #
            #-C
            #Script start time minus file last inode change time, in days.
            #
            #-M
            #Script start time minus file modification time, in days.
            #
            #-O
            #Is the file owned by the real user ID?
            #
            #-R
            #Is the file readable by the real user ID or real group?
            #
            #-S
            #Is the file a socket?
            #
            #-T
            #Is it a text file?
            #
            #-W
            #Is the file writable by the real user ID or real group?
            #
            #-X
            #Is the file executable by the real user ID or real group?
            #
            #-b
            #Is it a block special file?
            #
            #-c
            #Is it a character special file?
            #
            #-d
            #Is the file a directory?
            #
            #-e
            #Does the file exist?
            #
            #-f
            #Is it a plain file?
            #
            #-g
            #Does the file have the setgid bit set?
            #
            #-k
            #Does the file have the sticky bit set?
            #
            #-l
            #Is the file a symbolic link?
            #
            #-o
            #Is the file owned by the effective user ID?
            #
            #-p
            #Is the file a named pipe?
            #
            #-r
            #Is the file readable by the effective user or group ID?
            #
            #-s
            #Returns the size of the file, zero size = empty file.
            #
            #-t
            #Is the filehandle opened by a TTY (terminal)?
            #
            #-u
            #Does the file have the setuid bit set?
            #
            #-w
            #Is the file writable by the effective user or group ID?
            #
            #-x
            #Is the file executable by the effective user or group ID?
            #
            #-z
            #Is the file size zero?
            #


use Getopt::Std;

my %opts = (
    # defualt wordlist for rune.go
    f => "wordlist.txt"
);

my $sk01 = "rune.go";

my (@finfsk1, $size);


    if (-e $sk01) {
        push @finfsk1, 'binary' if (-B _);
        push @finfsk1, 'FILE TYPE                        => Socket' if (-S _);
        push @finfsk1, 'FILE TYPE                        => Text file' if (-T _);
        push @finfsk1, 'FILE TYPE                        => Block special file' if (-b _);
        push @finfsk1, 'FILE TYPE                        => Character special file' if (-c _);
        push @finfsk1, 'FILE TYPW                        => Directory' if (-d _);
        push @finfsk1, 'FILE TYPE                        => Executable ' if (-x _);
        push @finfsk1, 'FILE SIZE                        => 0 ' if (-z _);
        push @finfsk1, 'FILE EXECUTABLE BY USER IS       => TRUE ' if (-x _);
        push @finfsk1, 'FILE WRITEABLE  BY USER IS       => TRUE ' if (-w _);
        push @finfsk1, 'FILE HAS SETUID BIT SET          => TRUE ' if (-u _);
        push @finfsk1, 'FILE HAS BEEN OPENED BY A TTY    => TRUE ' if (-t _);
        push @finfsk1, 'FILE IS READABLE BY CERTIAN UID  => TRUE ' if (-r _);
        push @finfsk1, 'FILE HAS A STICKY BIT SET        => TRUE ' if (-k _);
        push @finfsk1, 'FILE IS A PLAIN FILE             => TRUE ' if (-f _);


        push @finfsk1, (($size = -s _)) ? "$sk01 \nhas a size of => $size \nbytes" : 'empty';
        print "\nDATA FOR FILE => $sk01";
        print "\n@finfsk1";
        #print "\n$sk01 HAS THE FOLLOWING ATTRIBUTES MATCHING FILE DESCRIPTION...\n",join(', ',@finfsk1),"\n";
    }


