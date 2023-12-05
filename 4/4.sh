#!/bin/bash

# part 1
sum=0
while IFS= read -r line
do
    ((sum+=$(echo "$line" | cut -d ':' -f2 | rg --pcre2 -o '\b(\d+)\b(?=.*\b\1\b)' | wc -l | awk '{ result = 0; for (i = 1; i <= $0; i++) { if (result == 0) result = 1; else result *= 2 } print result }')))
done < "../input/4"
echo "$sum"
