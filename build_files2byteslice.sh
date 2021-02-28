#!/bin/bash

# build files2byteslice
# go get github.com/hajimehoshi/file2byteslice
# cd ~/go/src/github.com/hajimehoshi/file2byteslice
git clone https://github.com/hajimehoshi/file2byteslice.git
cd file2byteslice
go get
go build -o file2byteslice ./cmd/file2byteslice/main.go

# copy to bin dir
# i guess we are assuming ~/go/bin is in PATH
cp file2byteslice ~/go/bin/

# cleanup
cd ..
rm -rf file2byteslice
