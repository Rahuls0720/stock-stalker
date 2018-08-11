#!/bin/bash

cd ..

# Compile and run with arg 'fb'
go build -o main src/*.go
./main fb

# Clean up
rm main
