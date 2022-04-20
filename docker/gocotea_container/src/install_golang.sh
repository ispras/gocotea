#!/bin/bash

echo "Downloading Golang archive"
wget https://go.dev/dl/go1.18.linux-amd64.tar.gz -P /home/ubuntu

echo "Deleting previous Go installation"
rm -rf /usr/local/go

echo "Extracting go"
tar -C /usr/local -xzf /home/ubuntu/go1.18.linux-amd64.tar.gz

echo "Adding go to PATH"
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
