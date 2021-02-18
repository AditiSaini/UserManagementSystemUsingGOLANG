#!/bin/bash

echo Creating binaries
cd ../servers/TCPServer
go build
cd ../HTTPServer
go build
echo All binaries created