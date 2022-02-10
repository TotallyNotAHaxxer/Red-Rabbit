![alt text](https://camo.githubusercontent.com/d00f1df32e58970fe90ac8cf342057a8c888c118a672ff81580e047db53e601c/68747470733a2f2f7261772e6769746875622e636f6d2f6b726169682f7065726c2d726170746f722f6d61696e2f6578616d706c652e706e67)


```
##                                                        =--_
#                                         .-""""""-.     |* _)
#                                        /          \   /  /
#                                       /            \_/  /
#           _                          /|                /
#       _-'"/\                        / |    ____    _.-"            _
#    _-'   (  '-_            _       (   \  |\  /\  ||           .-'".".
#_.-'       '.   `'-._   .-'"/'.      "   | |/ /  | |/        _-"   (   '-_
#             '.      _-"   (   '-_       \ | /   \ |     _.-'       )     "-._
#           _.'   _.-'       )     "-._    ||\\   |\\  '"'        .-'
#         '               .-'          `'  || \\  ||))
#   __  _  ___  _ ____________ _____  ___ _|\ _|\_|\\/ _______________  ___   _
#                       c  c  " c C ""C  " ""  "" ""
#                   c       C
#              C        C
#                   C
#    C     c
#
#
#Powered 
#      By
#        /$$$$$$$                     /$$
#        | $$__  $$                   | $$
#        | $$  \ $$ /$$$$$$   /$$$$$$ | $$
#        | $$$$$$$//$$__  $$ /$$__  $$| $$
#        | $$____/| $$$$$$$$| $$  \__/| $$
#        | $$     | $$_____/| $$      | $$
#        | $$     |  $$$$$$$| $$      | $$
#        |__/      \_______/|__/      |__/
#
```

Infected is a perl script that autogenerates QR_Codes, with the option of parsing a list and getting the EXIF data of the QR code

# EXIF of the average generated bar code
```
┌─────────────╥───────────────────┬─────────────────────────┐
│ Data Number ║ Data              │After DATA EXIF          │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ColorType          │     RGB with Alpha      │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Interlace          │      Noninterlaced      │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║MIMEType           │        image/png        │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileInodeChangeDate│2022:01:10 15:09:22-05:00│
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Directory          │            .            │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileTypeExtension  │           png           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Compression        │     Deflate/Inflate     │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileType           │           PNG           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ImageSize          │         185x185         │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ImageHeight        │           185           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ExifToolVersion    │          12.16          │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Filter             │        Adaptive         │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileAccessDate     │2022:01:10 15:07:38-05:00│
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ImageWidth         │           185           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║BitDepth           │           16            │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileModifyDate     │2022:01:10 15:09:22-05:00│
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileName           │        main2.png        │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Megapixels         │          0.034          │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileSize           │       1372 bytes        │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FilePermissions    │        rw-r--r--        │
└─────────────╨───────────────────┴─────────────────────────┘
```

# Usage, Installs, and understanding 

```
Before i start this kinda was a personal project, when i first started my road to legacy programming i picked up perl ( specifcially perl 5 ) 
and well got used to it, so i made this small 6 line script which would automate generating QR codes but it really wasnt handy, 
it was raining today when i stumbled across the file named "Q4R.pl" which was stored in a very old USB, 
i opened it up and got alot of general ideas of what i could to so i turned it into a way better script 
with general purposes and usages

alot of scripts and most websites have you select a template and even URL 
1 by 1 by 1 constantly which can take a fuck load of time, this one allows you 
to take a list of URL's, throw it as a -f argument in the script amd generate them all in a matter 
of seconds, im not kidding. i added a sleep method for purpose so the user could read the output
before the parsing starts, but when thats done you could generate up to 30 QR codes and images a second, 
i parsed a list of 100 urls and it still generated in under 2 seconds, that was depending also on my connection 
given that each URL is tested before inputting or parsing to the generator

install => git clone https://github.com/ArkAngeL43/Infected.git ; cd infected ; chmod +x ./install.sh ; ./install.sh

SYNTAX:
      Perl main.pl -o <filename>.png <url to imbed>
      
      example of single usage 
      
      perl barcode.pl -o main2.png https://discord.gg/upxdNC3Xf3
      
      FILE PARSING 
      
      perl main.pl -f < path to filename >
      
      EXAMPLE
      
      perl main.pl -f urls.txt
```
<br>
<br>
# EXAMPLE OUTPUT

```
+==============================================================================================+
+                                                                                              +
+             __     __   __     ______   ______     ______     ______   ______     _____      +
+            /\ \   /\ "-.\ \   /\  ___\ /\  ___\   /\  ___\   /\__  _\ /\  ___\   /\  __-.    +
+            \ \ \  \ \ \-.  \  \ \  __\ \ \  __\   \ \ \____  \/_/\ \/ \ \  __\   \ \ \/\ \   + 
+             \ \_\  \ \_\\"\_\  \ \_\    \ \_____\  \ \_____\    \ \_\  \ \_____\  \ \____-   +
+              \/_/   \/_/ \/_/   \/_/     \/_____/   \/_____/     \/_/   \/_____/   \/____/   +
+                                                                                              +
+                                                                                              +
+                                                                                              +
+                                                                                              +
+                                                     ___._                                    +
+                                                   .'  <0>'-.._                               +
+                                                  /  /.--.____")                              +
+                                                 |   \   __.-'~                               +
+                                                 |  :  -'/                                    +
+                                                /:.  :.-'                                     +
+__________                                     | : '. |                                       +
+'--.____  '--------.______       _.----.-----./      :/                                       +
+        '--.__            `'----/       '-.      __ :/                                        +
+              '-.___           :           \   .'  )/                                         +
+                    '---._           _.-'   ] /  _/                                           +
+                         '-._      _/     _/ / _/                                             +
+                             \_ .-'____.-'__< |  \___                                         +
+                               <_______.\    \_\_---.7                                        +
+                              |   /'=r_.-'     _\\ =/                                         +
+                          .--'   /            ._/'>                                           +
+                        .'   _.-'                                                             +
+                       / .--'                                                                 +
+                      /,/                                                                     +
+                      |/`)                                                                    +
+                      'c=,                                                                    +
+==============================================================================================+

#########################################
 Testing URL ~~> https://discord.gg/upxdNC3Xf3
#########################################
[ INFO ]  URL turned with a 200 REQUEST during GET FRAME 


	[ INFO ] WARNING: FATAL: => main2.png SEEMS TO ALREADY EXIST BEFORE GENERATION...

	[ INFO ] WARNING: FATAL: => GENERATING NEW AND RANDOM FILENAME TO PREVENT EXIT CODE 1...

	[ INFO ] WARNING: NEW STRING NAME => ACCC6202.png
sleeping for 5 seconds....
[ INFO ]  Chcking if main2.png exists 

[ INFO ]  File main2.png exists 

┌─────────────╥───────────────────┬─────────────────────────┐
│ Data Number ║ Data              │After DATA EXIF          │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ColorType          │     RGB with Alpha      │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Interlace          │      Noninterlaced      │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║MIMEType           │        image/png        │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileInodeChangeDate│2022:01:10 15:09:22-05:00│
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Directory          │            .            │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileTypeExtension  │           png           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Compression        │     Deflate/Inflate     │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileType           │           PNG           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ImageSize          │         185x185         │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ImageHeight        │           185           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ExifToolVersion    │          12.16          │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Filter             │        Adaptive         │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileAccessDate     │2022:01:10 15:07:38-05:00│
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║ImageWidth         │           185           │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║BitDepth           │           16            │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileModifyDate     │2022:01:10 15:09:22-05:00│
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileName           │        main2.png        │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║Megapixels         │          0.034          │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FileSize           │       1372 bytes        │
├─────────────╫───────────────────┼─────────────────────────┤
│1            ║FilePermissions    │        rw-r--r--        │
└─────────────╨───────────────────┴─────────────────────────┘
[ WARN ] GENERAL: EXIF DATA COLLECTED: QR CODE GENERATED


```
