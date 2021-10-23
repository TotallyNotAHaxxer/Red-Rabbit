use strict;
use warnings;
use feature 'say';
use HTTP::Tiny;
use Term::ANSIColor;


sub Second_Test(){
    my $Client = HTTP::Tiny->new();
    my @urls = (
        "https://doxbin.org/",
    );

    for my $url (@urls) {
        my $response = $Client->get($url);
        say $url, "Response ===> ", $response->{status};
        if($url, $response->{status} == '200'){
        say "okay connection";      
        }
        elsif($url, $response->{status} == '307'){
            say "Temporary Redirection testing again.....";
            test_again()
        }
        else{
            say "im not sure about this";
        }
    }
}

sub clear(){
    system("clear")
}


sub main(){
    my $Client = HTTP::Tiny->new();
    my @urls = (
        'http://www.yahoo.com',
        'https://www.google.com',
        "http://nosuchsiteexists.com",
        "https://instagram.com",
        "https://www.facebook.com/",
        "https://www.twitter.com/",
        "https://doxbin.org/",
        "https://www.etsy.com/shop/",
        "https://cash.me/",
        "https://www.behance.net/",
        "https://www.goodreads.com/",
        "https://www.instructables.com/",
        "https://google.com/",
        "https://keybase.io/",
        "https://kongregate.com/accounts/",
        "https://www.livejournal.com",
        "https://angel.co/",
        "https://last.fm/user/user123",
        "https://dribbble.com/",
        "https://www.codecademy.com/",
        "https://en.gravatar.com/",
        "https://pastebin.com/u/",
        "https://foursquare.com/",
        "https://foursquare.com/",
        "https://www.gumroad.com/",
        "https://.newgrounds.com",
        "https://www.wattpad.com/user/user123",
        "https://www.canva.com/",
        "https://creativemarket.com/",
        "https://www.trakt.tv",
        "https://500px.com/",
        "https://buzzfeed.com/",
    );
    print colored("TESTING THE FOLLOWING URLS\n@urls\n", "blink");
    sleep(6);
    for my $url (@urls) {
        my $response = $Client->get($url);
        print color("red"),"#########################################\n", color("reset");
        print color("red"),"# Testing URL ~~> $url\n", color("reset");
        print color("red"),"#########################################\n", color("reset");
        say $url, " Response ===> ", $response->{status};
        if($url, $response->{status} == '200'){
        say "okay connection";      
        }
        elsif($url, $response->{status} == '307'){
            say "Temporary Redirection testing again.....";
        }
        elsif($url, $response->{status} == '599'){
            say $url;
            say "Might be having issues connecting";
        }
        elsif($url, $response->{status} == '404'){
            say $url;
            say "CAN NOT CONNECT TO SERVER WARNING!!! THIS WILL AFFECT THE TOOLS ACCURACY";
        }
        else{
            say "This Error Code Doesnt Seem To Be In Index Range";
            say "If this is an error please note this can affect the";
            say "Tools Accuracy!!!!!!!";
        }
    }
}

clear();
main();
sleep(5);
clear();
say "Returning to Red-Rabbit";
sleep(3);
system(" cd .. ; sudo ruby main.rb")
