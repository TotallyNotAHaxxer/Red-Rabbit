section .data
        string db "Current Branded CPU < XXXXXXXXXXXX >",0xA
section	.text
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
   mov  eax, 8
   mov  ebx, file_name
   mov  ecx, 0777        
   int  0x80             
   mov [fd_out], eax
   mov	edx,len          
   mov	ecx, msg         
   mov	ebx, [fd_out]    
   mov	eax,4            
   int	0x80             
   mov eax, 6
   mov ebx, [fd_out]
   mov eax, 4
   mov ebx, 1
   mov ecx, msg_done
   mov edx, len_done
   int  0x80
   mov eax, 5
   mov ebx, file_name
   mov ecx, 0           
   mov edx, 0777         
   int  0x80
   mov  [fd_in], eax
   mov eax, 3
   mov ebx, [fd_in]
   mov ecx, info
   mov edx, 26
   int 0x80
   mov eax, 6
   mov ebx, [fd_in]
   int  0x80    
   mov eax, 4
   mov ebx, 1
   mov ecx, info
   mov edx, 26
   int 0x80
   mov	eax,1             
   int	0x80              
section	.data
file_name db 'rec.txt'
msg db "Current Branded CPU < XXXXXXXXXXXX >",0xA
len equ  $-msg
msg_done db 'Wrote to rec.txt', 0xa
len_done equ $-msg_done
section .bss
fd_out resb 1
fd_in  resb 1
info resb  26