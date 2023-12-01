#!/bin/bash

function first_last() {
    line=$1

    first=""
    last=""
    for ((i = 0; i < ${#line}; i++)); do
        char="${line:i:1}"
        [[ $char =~ [1-9] ]] && { [ -z "$first" ] && first="$char"; last="$char"; }
    done

    [ -n "$first" ] && [ -n "$last" ] && ((SUM += "$first$last"))
}

while IFS= read -r line
do
    first_last "$line"
done < "../input/1"

echo "$SUM"

SUM=0
while IFS= read -r line
do
    line=$(sed -e 's/one/o1e/g' -e 's/two/t2o/g' -e 's/three/t3e/g' -e 's/four/f4r/g' -e 's/five/f5e/g' -e 's/six/s6x/g' -e 's/seven/s7n/g' -e 's/eight/e8t/g' -e 's/nine/n9e/g' <<< "$line")
    first_last "$line"
done < "../input/1"

echo "$SUM"

