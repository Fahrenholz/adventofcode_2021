#!/usr/bin/env bash

current=$1
days=(zero test two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen twenty twentyone twentytwo twentythree twentyfour twentyfive)

mkdir "day${days[current]}"
cp main_template.txt "day${days[current]}/main.go"
touch "day${days[current]}/inputs.txt"


