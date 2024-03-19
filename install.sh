#!/bin/bash

DOT_DIR="$HOME/Downloads"

# バイナリファイルをダウンロード
curl -L -o ${DOT_DIR}/dotfiles-script https://github.com/jampee0915/dotfiles-v2/releases/download/v0.0.0/main

# 実行権限を付与
chmod +x ${DOT_DIR}/dotfiles-script

# 実行
${DOT_DIR}/dotfiles-script