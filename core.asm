section .data 
        string db "Using Current CPU =>  < XXXXXXXXXXXX >",0xA
section .text
global _start
_start:
       xor eax,eax
       cpuid 
       mov edi,string
       mov [edi+25],ebx 
       mov [edi+29],edx
       mov [edi+33],ecx
       mov eax,4
       mov ebx,1
       mov ecx,string
       mov edx,41
       int 0x80
       mov eax,1
       xor ebx,ebx
       int 0x80 