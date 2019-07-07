#!/bin/bash

output=`$1 "*/15 0 1,15 * 1-5 /usr/bin/find"`
expected="minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
"
diff <(echo ${output}) <(echo ${expected})

if [[ $? != 0 ]]; then
    echo "expected did not equal output"
    exit 1
fi

output=`$1 "bad input"`

if [[ $? == 0 ]]; then
    echo "expected failure for bad input"
    exit 1
fi

echo "e2e success"