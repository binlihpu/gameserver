#!/bin/bash

tabtoy -mode=v3 \
-index=index.xlsx \
-go_out=../../gamedata/data_gen.go \
-json_out=data_gen.json \
-package=gamedata

if [ $? -ne 0 ] ; then
	read -rsp $'Errors occurred...\n' ; 
	exit 1 
fi