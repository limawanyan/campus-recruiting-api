#!/bin/bash
printf "\nRegenerating swagger doc\n\n"
go install github.com/swaggo/swag/cmd/swag@v1.7.9
time swag init
printf "\nDone.\n\n"