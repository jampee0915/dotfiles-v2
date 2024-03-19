#!/bin/bash

DOT_DIR="$HOME/Downloads"
BINARY_FILE="dotfiles"

# バイナリファイルをダウンロード
curl -L -o ${DOT_DIR}/${BIN_FILE_NAME} https://github.com/jampee0915/dotfiles-v2/releases/download/v0.0.0/main

# 実行権限を付与
chmod +x ${DOT_DIR}/${BIN_FILE_NAME}

# 実行
${DOT_DIR}/${BIN_FILE_NAME}