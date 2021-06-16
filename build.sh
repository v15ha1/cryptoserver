#!/bin/bash

cd "$(dirname "$0")/src"

EXEC_NAME=cryptoserver-app
EXEC_FULL_NAME=../run/bin/${EXEC_NAME}

rm -f ${EXEC_FULL_NAME}

echo "Running go build"
CGO_ENABLED=1 GOOS=linux go build -v -ldflags '-extldflags "-static" -extldflags "-lpthread"' -o ${EXEC_FULL_NAME} main.go

if test $? -eq 0
then
    echo "Build completed successfully."
    echo "Binary copied to : ${EXEC_FULL_NAME}"
    exit 0
else
    echo "Build failed."
    exit 1
fi
