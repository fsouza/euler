#!/bin/sh

if [ $1 -lt 10 ];
then
    dir_name="problem_00$1"
else
    if [ $1 -lt 100 ];
    then
        dir_name="problem_0$1"
    else
        dir_name="problem_$1"
    fi
fi

cd $dir_name && 6g $1.go && 6l -o main.out $1.6 && ./main.out && make clean
