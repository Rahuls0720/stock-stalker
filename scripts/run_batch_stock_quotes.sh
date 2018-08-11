#!/bin/bash

cd ..

# Compile and run with arg 'fb'
go build -o main src/*.go
./main fb tsla aapl

# Clean up
rm main
