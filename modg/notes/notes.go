/*

 ______           __        ______         __     __     __ __          ___ ___ ______
|   __ \.-----.--|  |______|   __ \.---.-.|  |--.|  |--.|__|  |_ ______|   |   |    __|
|      <|  -__|  _  |______|      <|  _  ||  _  ||  _  ||  |   _|______|   |   |  __  |
|___|__||_____|_____|      |___|__||___._||_____||_____||__|____|       \_____/|______|
							Red Rabbit Version 6

This project was rewritten in perl5, Go, and Ruby

You may ask why the rewrite?:
				Simply because the code in version5 was not to my standards, and i also had alot more scripts to impliment i never did
				so i will not make this a big deal even though it is, and rather just update the repo, simply i decided to rewrite the main
				file of red rabbit in go, since my knowlege stays mainly in go and perl, while i do have my ruby skill i would rather rewrite
				alot of this in go, and given that go is easier to use for some things as well as perl without the use of third party libraries



This main file will have alot of the base function in rr6 built into this main file rather than running a million other files, a main reason i chose
to rewrite this in go is tbh most of the scripts used in RR5 were just automated go scripts using system commands in the ruby file, so why not just
switch red rabbit to go and then inline alot of the main functions RR5 had, sure this still will be automated alot with the os system however i now
have way more control over the output and direction as to where these scripts will go, especially since RR6 has its own perl and go modules to work
out of. Alot of people told me version5 was enough but to me it wasnt, i am not keeping this project for the people im keeping it for myself, sure
this will be released however this does not mean it was rewritten for people to use it more it was to show myself where i am at with my skill and knowlege of
programming and hacking, as well as my ideas. and as far as i am concerned as long as the team i have is still standing, my skill keeps advancing, and my brains
grow i will keep updating this framework with more and more commands. I did this alone, and im not backing down from it


pods hold props


note this project is fucking massive, its a large project and to someone who is not as experienced it will take time to learn especially with usage of flags mixed will a console
and its perl wrapper / code lib, this project runs off of JSON, Go, Perl, and Ruby libraries all which are run and taken in as commands by this main file, main.go SHOULD NOT BE MESSED
WITH AT ALL, all other modules are pre installed in the /usr/bin/perl5 standard directory for perl5 libs, go libs will be stored locally here since they ly inline in this file
and ruby modules are stored in the moduled directory so the r6.rb and r6.pl filenames can import them from that specific path, it is suggested you do not mess with the following files

r6.pl
r6.rb
main.go
OFSL.pm
TABLE.pm
all of the .pod
all of the .1.in
all of the .txt and .json files
all of the .data files
all of the .html and .css files

basically anything you dont want to touch due to its sensativity, and how this all ties to the main file, if you do not know what you are doing please leave now and just
simply use this tool and run it as is, overall unless you were a recruited dev apart of the scare security team then you should not touch this code

this project is protected by the BSD-3-clause license

License: BSD-3-clause
 Redistribution and use in source and binary forms, with or without modification,
 are permitted provided that the following conditions are met:
 .
     * Redistributions of source code must retain the above copyright notice,
     this list of conditions and the following disclaimer.
 .
     * Redistributions in binary form must reproduce the above copyright notice,
     this list of conditions and the following disclaimer in the documentation
     and/or other materials provided with the distribution.
 .
     * Neither the name of RE43P3R, or Scare_Security nor the names of its contributors
     may be used to endorse or promote products derived from this software
     without specific prior written permission.
 .
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
 ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
 ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


 please read the main.pod or notes.pod files for more infromation revolving reliability for the perl modules OFSL and TABLE

/*

-=- KNOWN BUG (ARP MODULE) -=-

panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x740fc7]

goroutine 12 [running]:
main.ARPW(0x77359400, 0xc00007ebc0, 0x0)
	/home/reaper/Desktop/RR6/main.go:2857 +0x1c7
main.scan(0xc00007ebc0)
	/home/reaper/Desktop/RR6/main.go:2691 +0x9ad
main.arpmain.func1({0x7, 0x5dc, {0xc000022620, 0x9}, {0xc0001f01dc, 0x6, 0xa24}, 0x3})
	/home/reaper/Desktop/RR6/main.go:2899 +0x11d
created by main.arpmain
	/home/reaper/Desktop/RR6/main.go:2897 +0x1ff
exit status 2


*/
package notes
