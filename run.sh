#!/bin/sh

dir_name=`printf "problem_%03d" $1`
cd $dir_name && 6g $1.go && 6l $1.6 && ./6.out && make clean
