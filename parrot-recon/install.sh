#!/bin/bash

working_dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)

echo "[+] Installing Tools Required For Parrot-Recon"
sudo apt install go nmap nikto amass gobuster dirbuster sslyze sublist3r wpscan
go get -u -v github.com/lukasikic/subzy
go install github.com/lukasikic/subzy@latest
go get -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei

echo "[+] Script Done!"
