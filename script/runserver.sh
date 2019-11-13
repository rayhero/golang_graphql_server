#!/bin/sh
srcPath="src/nier"
pkgFile="main.go"
app="gqlserver"
src="$srcPath/$app/$pkgFile"

printf "\nStart running: $app\n"
# Set all ENV vars for the server to run
time go run $src
# export $(grep -v '^#' .env | xargs) && time go run $src
# This should unset all the ENV vars, just in case.
# unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped running: $app\n\n"