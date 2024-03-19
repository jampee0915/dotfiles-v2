#!/bin/bash

DOT_DIR="$HOME/Downloads"

curl -L -o ${DOT_DIR}/dotfiles-script https://github.com/jampee0915/dotfiles-v2/releases/download/v0.0.0/main
chmod +x ${DOT_DIR}/dotfiles-script

cd ${DOT_DIR}
./dotfiles-script
