#!/bin/bash

set -e  # Para parar em caso de erro

mkdir -p ./bin ./dist

# Nome do CLI
CLI_NAME="github-activity"

# Compilar para Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o ./bin/${CLI_NAME}

# Empacotar para .deb
echo "📦 Generating .deb package..."
nfpm pkg --config nfpm.yaml --packager deb --target ./dist

# Empacotar para .rpm
echo "📦 Generating .rpm package..."
nfpm pkg --config nfpm.yaml --packager rpm --target ./dist

# Empacotar para Arch (.pkg.tar.zst)
echo "📦 Generating Arch package..."
nfpm pkg --config nfpm.yaml --packager archlinux --target ./dist

echo ""

# Compilar para Windows
echo "Building for Windows..."
echo "📦 Generating Windows package..."
GOOS=windows GOARCH=amd64 go build -o ./bin/${CLI_NAME}.exe

echo "Build concluído!"
