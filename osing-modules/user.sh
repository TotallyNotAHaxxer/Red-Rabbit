clear
ban() {
    go run banner.go
}
scanner() {
read -p $'\n\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Username:\e[0m ' username
check_insta=$(curl -s -H "Accept-Language: en" "https://www.instagram.com/$username" -L | grep -o 'The link you followed may be broken'; echo $?)
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Instagram: \e[0m"
if [[ $check_insta == *'1'* ]]; then
printf "\e[1;97m \e[0;97m https://www.instagram.com/%s\n" $username
printf "https://www.instagram.com/%s\n" $username > $username.txt
elif [[ $check_insta == *'0'* ]]; then
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Facebook: \e[0m"
check_face=$(curl -s "https://www.facebook.com/$username" -L -H "Accept-Language: en" | grep -o 'not found'; echo $?)
if [[ $check_face == *'1'* ]]; then
printf "\e[1;92m \e[0m https://www.facebook.com/%s\n" $username
printf "https://www.facebook.com/%s\n" $username >> $username.txt
elif [[ $check_face == *'0'* ]]; then
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Twitter: \e[0m"
check_twitter=$(curl -s "https://www.twitter.com/$username" -L -H "Accept-Language: en" | grep -o 'page doesn’t exist'; echo $?)
if [[ $check_twitter == *'1'* ]]; then
printf "\e[1;92m \e[0m https://www.twitter.com/%s\n" $username
printf "https://www.twitter.com/%s\n" $username >> $username.txt
elif [[ $check_twitter == *'0'* ]]; then
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ YouTube: \e[0m"
check_youtube=$(curl -s "https://www.youtube.com/$username" -L -H "Accept-Language: en" | grep -o 'Not Found'; echo $?)
if [[ $check_youtube == *'1'* ]]; then
printf "\e[1;92m \e[0m https://www.youtube.com/%s\n" $username
printf "https://www.youtube.com/%s\n" $username >> $username.txt
elif [[ $check_youtube == *'0'* ]]; then
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Blogger: \e[0m"
check=$(curl -s "https://$username.blogspot.com" -L -H "Accept-Language: en" -i | grep -o 'HTTP/2 404'; echo $?)
if [[ $check == *'1'* ]]; then
printf "\e[1;92m \e[0m https://%s.blogspot.com\n" $username
printf "https://%s.blogspot.com\n" $username >> $username.txt
elif [[ $check == *'0'* ]]; then
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ GooglePlus: \e[0m"
check=$(curl -s "https://plus.google.com/+$username/posts" -L -H "Accept-Language: en" -i | grep -o 'HTTP/2 404' ; echo $?)
if [[ $check == *'1'* ]]; then
printf "\e[1;92m \e[0m https://plus.google.com/+%s/posts\n" $username
printf "https://plus.google.com/+%s/posts\n" $username >> $username
elif [[ $check == *'0'* ]]; then
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Reddit: \e[0m"
check1=$(curl -s -i "https://www.reddit.com/user/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | head -n1 | grep -o 'HTTP/2 404' ; echo $?)
if [[ $check1 == *'0'* ]] ; then 
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
elif [[ $check1 == *'1'* ]]; then 
printf "\e[1;92m \e[0m https://www.reddit.com/user/%s\n" $username
printf "https://www.reddit.com/user/%s\n" $username >> $username.txt
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Wordpress: \e[0m"
check1=$(curl -s -i "https://$username.wordpress.com" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o 'Do you want to register' ; echo $?)
if [[ $check1 == *'0'* ]] ; then 
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
elif [[ $check1 == *'1'* ]]; then 
printf "\e[1;92m \e[0m https://%s.wordpress.com\n" $username
printf "https://%s.wordpress.com\n" $username >> $username.txt
fi
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Pinterest: \e[0m"
check1=$(curl -s -i "https://www.pinterest.com/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o '?show_error' ; echo $?)
if [[ $check1 == *'0'* ]] ; then 
printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"
elif [[ $check1 == *'1'* ]]; then 
printf "\e[1;92m \e[0m https://www.pinterest.com/%s\n" $username
printf "https://www.pinterest.com/%s\n" $username >> $username.txt
fi 
printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Github: \e[0m"

check1=$(curl -s -i "https://www.github.com/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.github.com/%s\n" $username

printf "https://www.github.com/%s\n" $username >> $username.txt

fi

 

## TUMBLR

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Tumblr: \e[0m"

check1=$(curl -s -i "https://$username.tumblr.com" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.tumblr.com\n" $username

printf "https://%s.tumblr.com\n" $username >> $username.txt

fi

 

## FLICKR

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Flickr: \e[0m"

check1=$(curl -s -i "https://www.flickr.com/people/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o 'Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.flickr.com/photos/%s\n" $username

printf "https://www.flickr.com/photos/%s\n" $username >> $username.txt

fi

 

## STEAM

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Steam: \e[0m"

check1=$(curl -s -i "https://steamcommunity.com/id/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o 'The specified profile could not be found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://steamcommunity.com/id/%s\n" $username

printf "https://steamcommunity.com/id/%s\n" $username >> $username.txt

fi

 

## VIMEO

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Vimeo: \e[0m"

check1=$(curl -s -i "https://vimeo.com/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://vimeo.com/%s\n" $username

printf "https://vimeo.com/%s\n" $username >> $username.txt

fi

 

 

## SoundCloud

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ SoundCloud: \e[0m"

check1=$(curl -s -i "https://soundcloud.com/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o '404 Not Found'; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://soundcloud.com/%s\n" $username

printf "https://soundcloud.com/%s\n" $username >> $username.txt

fi

 

## DISQUS

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Disqus: \e[0m"

check1=$(curl -s -i "https://disqus.com/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o '404 NOT FOUND' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://disqus.com/%s\n" $username

printf "https://disqus.com/%s\n" $username >> $username.txt

fi

 

## MEDIUM

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Medium: \e[0m"

check1=$(curl -s -i "https://medium.com/@$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://medium.com/@%s\n" $username

printf "https://medium.com/@%s\n" $username >> $username.txt

fi

 

## DEVIANTART

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ DeviantART: \e[0m"

check1=$(curl -s -i "https://$username.deviantart.com" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.deviantart.com\n" $username

printf "https://%s.deviantart.com\n" $username >> $username.txt

fi

 

## VK

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ VK: \e[0m"

check1=$(curl -s -i "https://vk.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://vk.com/%s\n" $username

printf "https://vk.com/%s\n" $username >> $username.txt

fi

 

## About.me

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ About.me: \e[0m"

check1=$(curl -s -i "https://about.me/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://about.me/%s\n" $username

printf "https://about.me/%s\n" $username >> $username.txt

fi

 

 

## Imgur

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Imgur: \e[0m"

check1=$(curl -s -i "https://imgur.com/user/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://imgur.com/user/%s\n" $username

printf "https://imgur.com/user/%s\n" $username >> $username.txt

fi

 

## FlipBoard

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Flipboard: \e[0m"

check1=$(curl -s -i "https://flipboard.com/@$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://flipboard.com/@%s\n" $username

printf "https://flipboard.com/@%s\n" $username >> $username.txt

fi

 

## SlideShare

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ SlideShare: \e[0m"

check1=$(curl -s -i "https://slideshare.net/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://slideshare.net/%s\n" $username

printf "https://slideshare.net/%s\n" $username >> $username.txt

fi

 

## Fotolog

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Fotolog: \e[0m"

check1=$(curl -s -i "https://fotolog.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://fotolog.com/%s\n" $username

printf "https://fotolog.com/%s\n" $username >> $username.txt

fi

 

 

## Spotify

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Spotify: \e[0m"

check1=$(curl -s -i "https://open.spotify.com/user/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://open.spotify.com/user/%s\n" $username

printf "https://open.spotify.com/user/%s\n" $username >> $username.txt

fi

 

## MixCloud

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ MixCloud: \e[0m"

check1=$(curl -s -i "https://www.mixcloud.com/$username" -H "Accept-Language: en" -L | grep -o 'error-message' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.mixcloud.com/%s\n" $username

printf "https://www.mixcloud.com/%s\n" $username >> $username.txt

fi

 

## Scribd

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Scribd: \e[0m"

check1=$(curl -s -i "https://www.scribd.com/$username" -H "Accept-Language: en" -L | grep -o 'show_404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.scribd.com/%s\n" $username

printf "https://www.scribd.com/%s\n" $username >> $username.txt

fi

 

## Badoo

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Badoo: \e[0m"

check1=$(curl -s -i "https://www.badoo.com/en/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.badoo.com/en/%s\n" $username

printf "https://www.badoo.com/en/%s\n" $username >> $username.txt

fi

 

# Patreon

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Patreon: \e[0m"

check1=$(curl -s -i "https://www.patreon.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.patreon.com/%s\n" $username

printf "https://www.patreon.com/%s\n" $username >> $username.txt

fi

 

## BitBucket

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ BitBucket: \e[0m"

check1=$(curl -s -i "https://bitbucket.org/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://bitbucket.org/%s\n" $username

printf "https://bitbucket.org/%s\n" $username >> $username.txt

fi

 

## DailyMotion

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ DailyMotion: \e[0m"

check1=$(curl -s -i "https://www.dailymotion.com/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.dailymotion.com/%s\n" $username

printf "https://www.dailymotion.com/%s\n" $username >> $username.txt

fi

 

## Etsy

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Etsy: \e[0m"

check1=$(curl -s -i "https://www.etsy.com/shop/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.etsy.com/shop/%s\n" $username

printf "https://www.etsy.com/shop/%s\n" $username >> $username.txt

fi

 

## CashMe

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ CashMe: \e[0m"

check1=$(curl -s -i "https://cash.me/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found'; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://cash.me/%s\n" $username

printf "https://cash.me/%s\n" $username >> $username.txt

fi

 

## Behance

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Behance: \e[0m"

check1=$(curl -s -i "https://www.behance.net/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found'; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.behance.net/%s\n" $username

printf "https://www.behance.net/%s\n" $username >> $username.txt

fi

 

## GoodReads

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ GoodReads: \e[0m"

check1=$(curl -s -i "https://www.goodreads.com/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.goodreads.com/%s\n" $username

printf "https://www.goodreads.com/%s\n" $username >> $username.txt

fi

 

## Instructables

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Instructables: \e[0m"

check1=$(curl -s -i "https://www.instructables.com/member/$username" -H "Accept-Language: en" -L | grep -o '404 NOT FOUND' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.instructables.com/member/%s\n" $username

printf "https://www.instructables.com/member/%s\n" $username >> $username.txt

fi

 

## KeyBase

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Keybase: \e[0m"

check1=$(curl -s -i "https://keybase.io/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://keybase.io/%s\n" $username

printf "https://keybase.io/%s\n" $username >> $username.txt

fi

 

## Kongregate

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Kongregate: \e[0m"

check1=$(curl -s -i "https://kongregate.com/accounts/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://kongregate.com/accounts/%s\n" $username

printf "https://kongregate.com/accounts/%s\n" $username >> $username.txt

fi

 

## Livejournal

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ LiveJournal: \e[0m"

check1=$(curl -s -i "https://$username.livejournal.com" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.livejournal.com\n" $username

printf "https://%s.livejournal.com\n" $username >> $username.txt

fi

 

## AngelList

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ AngelList: \e[0m"

check1=$(curl -s -i "https://angel.co/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://angel.co/%s\n" $username

printf "https://angel.co/%s\n" $username >> $username.txt

fi

 

## Last.fm

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ last.fm: \e[0m"

check1=$(curl -s -i "https://last.fm/user/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://last.fm/user/%s\n" $username

printf "https://last.fm/user/%s\n" $username >> $username.txt

fi

 

## Dribbble

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Dribbble: \e[0m"

check1=$(curl -s -i "https://dribbble.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://dribbble.com/%s\n" $username

printf "https://dribbble.com/%s\n" $username >> $username.txt

fi

 

## Codecademy

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Codecademy: \e[0m"

check1=$(curl -s -i "https://www.codecademy.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.codecademy.com/%s\n" $username

printf "https://www.codecademy.com/%s\n" $username >> $username.txt

fi

 

## Gravatar

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Gravatar: \e[0m"

check1=$(curl -s -i "https://en.gravatar.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://en.gravatar.com/%s\n" $username

printf "https://en.gravatar.com/%s\n" $username >> $username.txt

fi

 

## Pastebin

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Pastebin: \e[0m"

check1=$(curl -s -i "https://pastebin.com/u/$username" -H "Accept-Language: en" -L --user-agent '"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801"' | grep -o 'location: /index' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://pastebin.com/u/%s\n" $username

printf "https://pastebin.com/u/%s\n" $username >> $username.txt

fi

 

## Foursquare

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Foursquare: \e[0m"

check1=$(curl -s -i "https://foursquare.com/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://foursquare.com/%s\n" $username

printf "https://foursquare.com/%s\n" $username >> $username.txt

fi

 

## Roblox

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Roblox: \e[0m"

check1=$(curl -s -i "https://www.roblox.com/user.aspx?username=$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://foursquare.com/%s\n" $username

printf "https://foursquare.com/%s\n" $username >> $username.txt

fi

 

## Gumroad

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Gumroad: \e[0m"

check1=$(curl -s -i "https://www.gumroad.com/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.gumroad.com/%s\n" $username

printf "https://www.gumroad.com/%s\n" $username >> $username.txt

fi

 

## Newgrounds

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Newgrounds: \e[0m"

check1=$(curl -s -i "https://$username.newgrounds.com" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404 ' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.newgrounds.com\n" $username

printf "https://%s.newgrounds.com\n" $username >> $username.txt

fi

 

## Wattpad

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Wattpad: \e[0m"

check1=$(curl -s -i "https://www.wattpad.com/user/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404 ' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.wattpad.com/user/%s\n" $username

printf "https://www.wattpad.com/user/%s\n" $username >> $username.txt

fi

 

## Canva

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Canva: \e[0m"

check1=$(curl -s -i "https://www.canva.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404 ' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.canva.com/%s\n" $username

printf "https://www.canva.com/%s\n" $username >> $username.txt

fi

 

## CreativeMarket

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ CreativeMarket: \e[0m"

check1=$(curl -s -i "https://creativemarket.com/$username" -H "Accept-Language: en" -L | grep -o '404eef72' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://creativemarket.com/%s\n" $username

printf "https://creativemarket.com/%s\n" $username >> $username.txt

fi

 

## Trakt

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Trakt: \e[0m"

check1=$(curl -s -i "https://www.trakt.tv/users/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404 ' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.trakt.tv/users/%s\n" $username

printf "https://www.trakt.tv/users/%s\n" $username >> $username.txt

fi

 

## 500px

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ 500px: \e[0m"

check1=$(curl -s -i "https://500px.com/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://500px.com/%s\n" $username

printf "https://500px.com/%s\n" $username >> $username.txt

fi

 

## Buzzfeed

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Buzzfeed: \e[0m"

check1=$(curl -s -i "https://buzzfeed.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://buzzfeed.com/%s\n" $username

printf "https://buzzfeed.com/%s\n" $username >> $username.txt

fi

 

## TripAdvisor

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ TripAdvisor: \e[0m"

check1=$(curl -s -i "https://tripadvisor.com/members/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://tripadvisor.com/members/%s\n" $username

printf "https://tripadvisor.com/members/%s\n" $username >> $username.txt

fi

 

## HubPages

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ HubPages: \e[0m"

check1=$(curl -s -i "https://$username.hubpages.com" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.hubpages.com/\n" $username

printf "https://%s.hubpages.com/\n" $username >> $username.txt

fi

 

## Contently

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Contently: \e[0m"

check1=$(curl -s -i "https://$username.contently.com" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.contently.com\n" $username

printf "https://%s.contently.com\n" $username >> $username.txt

fi

 

## Houzz

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Houzz: \e[0m"

check1=$(curl -s -i "https://houzz.com/user/$username" -H "Accept-Language: en" -L | grep -o 'an error has occurred' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://houzz.com/user/%s\n" $username

printf "https://houzz.com/user/%s\n" $username >> $username.txt

fi

 

## blip.fm

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ blip.fm: \e[0m"

check1=$(curl -s -i "https://blip.fm/$username" -H "Accept-Language: en" -L | grep -o '404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://blip.fm/%s\n" $username

printf "https://blip.fm/%s\n" $username >> $username.txt

fi

 

## Wikipedia

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Wikipedia: \e[0m"

check1=$(curl -s -i "https://www.wikipedia.org/wiki/User:$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.wikipedia.org/wiki/User:%s\n" $username

printf "https://www.wikipedia.org/wiki/User:%s\n" $username >> $username.txt

fi

 

## HackerNews

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ HackerNews: \e[0m"

check1=$(curl -s -i "https://news.ycombinator.com/user?id=$username" -H "Accept-Language: en" -L | grep -o 'No such user' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://news.ycombinator.com/user?id=%s\n" $username

printf "https://news.ycombinator.com/user?id=%s\n" $username >> $username.txt

fi

 

## CodeMentor

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ CodeMentor: \e[0m"

check1=$(curl -s -i "https://www.codementor.io/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.codementor.io/%s\n" $username

printf "https://www.codementor.io/%s\n" $username >> $username.txt

fi

 

## ReverbNation

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ ReverbNation: \e[0m"

check1=$(curl -s -i "https://www.reverbnation.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.reverbnation.com/%s\n" $username

printf "https://www.reverbnation.com/%s\n" $username >> $username.txt

fi

 

## Designspiration 65

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Designspiration: \e[0m"

check1=$(curl -s -i "https://www.designspiration.net/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.designspiration.net/%s\n" $username

printf "https://www.designspiration.net/%s\n" $username >> $username.txt

fi

 

## Bandcamp

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Bandcamp: \e[0m"

check1=$(curl -s -i "https://www.bandcamp.com/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.bandcamp.com/%s\n" $username

printf "https://www.bandcamp.com/%s\n" $username >> $username.txt

fi

 

 

## ColourLovers

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ ColourLovers: \e[0m"

check1=$(curl -s -i "https://www.colourlovers.com/love/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.colourlovers.com/love/%s\n" $username

printf "https://www.colourlovers.com/love/%s\n" $username >> $username.txt

fi

 

 

## IFTTT

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ IFTTT: \e[0m"

check1=$(curl -s -i "https://www.ifttt.com/p/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.ifttt.com/p/%s\n" $username

printf "https://www.ifttt.com/p/%s\n" $username >> $username.txt

fi

 

## Ebay

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Ebay: \e[0m"

check1=$(curl -s -i "https://www.ebay.com/usr/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found\|eBay Profile - error' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.ebay.com/usr/%s\n" $username

printf "https://www.ebay.com/usr/%s\n" $username >> $username.txt

fi

 

## Slack

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Slack: \e[0m"

check1=$(curl -s -i "https://$username.slack.com" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.slack.com\n" $username

printf "https://%s.slack.com\n" $username >> $username.txt

fi

 

## OkCupid

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ OkCupid: \e[0m"

check1=$(curl -s -i "https://www.okcupid.com/profile/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.okcupid.com/profile/%s\n" $username

printf "https://www.okcupid.com/profile/%s\n" $username >> $username.txt

fi

 

## Trip

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Trip: \e[0m"

check1=$(curl -s -i "https://www.trip.skyscanner.com/user/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found\|HTTP/2 410' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.trip.skyscanner.com/user/%s\n" $username

printf "https://www.trip.skyscanner.com/user/%s\n" $username >> $username.txt

fi

 

## Ello

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Ello: \e[0m"

check1=$(curl -s -i "https://ello.co/$username" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://ello.co/%s\n" $username

printf "https://ello.co/%s\n" $username >> $username.txt

fi

 

## Tracky

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Tracky: \e[0m"

check1=$(curl -s -i "https://tracky.com/user/$username" -H "Accept-Language: en" -L | grep -o 'profile:username' ; echo $?)

 

if [[ $check1 == *'1'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'0'* ]]; then 

 

printf "\e[1;92m \e[0m https://tracky.com/~%s\n" $username

printf "https://tracky.com/~%s\n" $username >> $username.txt

fi

 

## Tripit

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Tripit: \e[0m"

check1=$(curl -s -i "https://www.tripit.com/people/$username#/profile/basic-info" -H "Accept-Language: en" -L | grep -o 'location: https://www.tripit.com/home' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://www.tripit.com/people/%s#/profile/basic-info\n" $username

printf "https://www.tripit.com/people/%s#/profile/basic-info\n" $username >> $username.txt

fi

 

## Basecamp

 

printf "\e[1;96m(⇀‸↼‶)⊃━☆ﾟ.*･｡ﾟ Basecamp: \e[0m"

check1=$(curl -s -i "https://$username.basecamphq.com/login" -H "Accept-Language: en" -L | grep -o 'HTTP/2 404\|404 Not Found' ; echo $?)

 

if [[ $check1 == *'0'* ]] ; then 

printf "\e[1;91m⚠️ NOT FOUND\e[0m\n"

elif [[ $check1 == *'1'* ]]; then 

 

printf "\e[1;92m \e[0m https://%s.basecamphq.com/login\n" $username

printf "https://%s.basecamphq.com/login\n" $username >> $username.txt

 

fi

}

ban

scanner