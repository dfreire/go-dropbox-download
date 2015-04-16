#!/bin/bash

set -e -x

echo "Creating release dir..."
mkdir -p release

echo "Creating darwin/386 binary..."
GOOS=darwin GOARCH=386 go build -o out/go-dropbox-download go-dropbox-download.go
cd out
tar cvzf ../release/go-dropbox-download-mac-32bit.tgz go-dropbox-download
cd ..

echo "Creating darwin/amd64 binary..."
GOOS=darwin GOARCH=amd64 go build -o out/go-dropbox-download go-dropbox-download.go
cd out
tar cvzf ../release/go-dropbox-download-mac-64bit.tgz go-dropbox-download
cd ..

echo "Creating linux/386 binary..."
GOOS=linux GOARCH=386 go build -o out/go-dropbox-download go-dropbox-download.go
cd out
tar cvzf ../release/go-dropbox-download-linux-32bit.tgz go-dropbox-download
cd ..

echo "Creating linux/amd64 binary..."
GOOS=linux GOARCH=amd64 go build -o out/go-dropbox-download go-dropbox-download.go
cd out
tar cvzf ../release/go-dropbox-download-linux-64bit.tgz go-dropbox-download
cd ..

echo "Creating windows/386 binary..."
GOOS=windows GOARCH=386 go build -o out/go-dropbox-download.exe go-dropbox-download.go
cd out
zip ../release/go-dropbox-download-windows-32bit.zip go-dropbox-download.exe
cd ..

echo "Creating windows/amd64 binary..."
GOOS=windows GOARCH=amd64 go build -o out/go-dropbox-download.exe go-dropbox-download.go
cd out
zip ../release/go-dropbox-download-windows-64bit.zip go-dropbox-download.exe
cd ..

rm -rf out
