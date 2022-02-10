// simple test for the binary scanner, represents stack overflow and buffer overflow 
// vulnerabilities
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]){
char buffer[100];
strcpy(buffer,  argv[1]);
return 0;
}