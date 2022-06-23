#!/bin/bash

shellExit()
{
if [ $1 -eq 1 ]; then
    printf "\nfailed!!!\n\n"
    exit 1
fi
}

printf "\nRestart server port:8081 \n\n"

lsof -i:8081 | grep LISTEN | awk '{print $2}' | xargs kill -s SIGINT && go run ./main.go -env fat
shellExit $?

printf "\nDone.\n\n"
