#!/bin/bash

cd ..

# Compile and run with arg 'fb', flag -t/-T for tracking
go build -o main src/*.go
./main -t fb 150 250

# Clean up
rm main
