# Shroot-finder
```
 
███████╗ ██████╗██╗  ██╗██████╗ ██╗   ██╗████████╗███████╗    ███████╗██╗███╗   ██╗██████╗ ███████╗██████╗ 
██╔════╝██╔════╝██║  ██║██╔══██╗██║   ██║╚══██╔══╝██╔════╝    ██╔════╝██║████╗  ██║██╔══██╗██╔════╝██╔══██╗
███████╗██║     ███████║██████╔╝██║   ██║   ██║   █████╗█████╗█████╗  ██║██╔██╗ ██║██║  ██║█████╗  ██████╔╝
╚════██║██║     ██╔══██║██╔══██╗██║   ██║   ██║   ██╔══╝╚════╝██╔══╝  ██║██║╚██╗██║██║  ██║██╔══╝  ██╔══██╗
███████║╚██████╗██║  ██║██║  ██║╚██████╔╝   ██║   ███████╗    ██║     ██║██║ ╚████║██████╔╝███████╗██║  ██║
╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚══════╝    ╚═╝     ╚═╝╚═╝  ╚═══╝╚═════╝ ╚══════╝╚═╝  ╚═╝
                                                                                                           



```


This is a SQL Finder, Admin panel finder that goes basses off requests. Made from 100% golang


what can it do?

Schrute-Finder is a simple tool to gather information on a URL like finding adming panels and attempting to sql inject a header and test if the header respond ok indicating it may be injectable

```
how to 


--------------------------------------------------------------------
[*] Value usage
go run main.go |-s|-a|-w true   -t <target> 
--------------------------- Basic value usages-------------
-> go run main.go -s true -a true -w true -t www.example.com | this will get server information, enable admin and SQL testing, then take the HTML and save it 
-> Enable a host -> -a|-s|-w -> true -> -s true same with the others 
Usage :> go run main.go -t https://www.parrot-pentest.com/
  -a string
    	admin panel finding to enable this type true or | go run main.go -a true
  -f string
    	Eliminate time waiting process to prevent False positives -> this could return false results
  -h	Print usage instructions
  -s string
    	SQL vulnerability parsing to enable this type true or | go run main.go -s true 
  -t string
    	target URL
  -v	Print version
  -w string
    	Write HTML data to a file | to enable this use -w true


```

install the f*king go-lang first 

```
go get github.com/briandowns/spinner
```
