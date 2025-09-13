#!/bin/sh -eu
set -e

EXECUTABLE="GOBOL-LSP"
RED="\033[31m"
GREEN="\033[32m"
YELLOW="\033[33m"
BLUE="\033[34m"
RESET="\033[0m"

if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go and try again.${RESET}"
    exit 1
fi

echo -e "${BLUE}Cloning 'https://github.com/jtomaspm/GOBOL-LSP'${RESET}"
git clone https://github.com/jtomaspm/GOBOL-LSP $EXECUTABLE

echo -e "${BLUE}Building executable...${RESET}"
cd $EXECUTABLE
go build -o $EXECUTABLE ./cmd/lsp/main.go

echo -e "${BLUE}Installing in '$HOME/.local/bin/$EXECUTABLE'${RESET}"
mkdir -p "$HOME/.local/bin"
if [ -f "$HOME/.local/bin/$EXECUTABLE" ]; then
    echo -e "${YELLOW}Removed old version...${RESET}"
    rm "$HOME/.local/bin/$EXECUTABLE"
fi
mv $EXECUTABLE "$HOME/.local/bin/$EXECUTABLE"
chmod +x "$HOME/.local/bin/$EXECUTABLE"

echo -e "${BLUE}Removing source code...${RESET}"
cd ..
rm -rf $EXECUTABLE

echo -e "${GREEN}✅ Installed GOBOL-LSP successfully!${RESET}"