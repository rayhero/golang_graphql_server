#!/bin/sh
srcPath="src/nier"
pkgFile="main.go"
outputPath="build"
app="gqlserver"
output="$outputPath/$app"
src="$srcPath/$app/$pkgFile"

printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"