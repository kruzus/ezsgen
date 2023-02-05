#!/bin/sh

// @TODO: finish this build automation, check folder if it exists, etc...

result=${PWD##*/}    

   

printf 'Building project: %s to bin folder...\n' "${PWD##*/}" 

go build -ldflags "-w" -o bin/app

printf 'Project built at: %s\n' "bin/app"
 