#!/bin/bash

# build and install mrun binary

OUTPUT_LOCATION=/usr/bin/mrun

cd ./src/
sudo go build -o $OUTPUT_LOCATION main.go

alias mrun='$OUTPUT_LOCATION'
source ~/.bashrc
source ~/.zshrc
