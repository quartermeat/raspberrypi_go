#!/bin/bash

# Build for Debian aarch64
env GOOS=linux GOARCH=arm64 go build -o ./bin/myapp-arm64

# Replace with your actual remote host, user, and path
REMOTE_HOST="192.168.0.106"
REMOTE_USER="quartermeat"
REMOTE_PATH="./app/myapp"

# SFTP command to transfer the binary
sftp $REMOTE_USER@$REMOTE_HOST <<EOF
put ./bin/myapp-arm64 $REMOTE_PATH
exit
EOF