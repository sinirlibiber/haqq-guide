#!/bin/bash
echo -e "\033[0;31m"
echo "==================================================";
echo "     █▀▀▀ █░░█ █▀▄▀█ █░░█ █▀▀ █▀▀▄ █▀▀ █░░█       ";
echo "     █░▀█ █░░█ █░▀░█ █░░█ ▀▀█ █▀▀▄ █▀▀ █▄▄█ 		";
echo "     ▀▀▀▀ ░▀▀▀ ▀░░░▀ ░▀▀▀ ▀▀▀ ▀▀▀░ ▀▀▀ ▄▄▄█    	";
echo -e "\e[0m"
echo "=================================================="                                  

sleep 2


echo "=================================================="

echo -e "\e[1m\e[32m1. NODECORD GO YÜKLENİYOR... \e[0m" && sleep 1
wget https://golang.org/dl/go1.18.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz

# GO YAPILANDIRMA
cat <<EOF >> ~/.profile
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
EOF
source ~/.profile
go version

sleep 1

echo "=================================================="

echo -e "\e[1m\e[32m2. NODECORD REPOSU INDIRILIYOR ... \e[0m" && sleep 1
git clone https://github.com/sinirlibiber/HAQQREHBER/tree/main/Nodecord-main
cd ~/Nodecord
go install

echo "=================================================="

echo -e "\e[1m\e[32m KURULUM BAŞARIYLA TAMAMLANDI... \e[0m" && sleep 1
