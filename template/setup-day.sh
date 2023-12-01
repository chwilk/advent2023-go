#!/bin/bash

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
DAY=${1:-25}
if ((DAY > 0)) && ((DAY < 26))
then
    dest="$SCRIPTDIR/day$DAY"
    mkdir -p $dest
    sed "s/XX/$DAY/" $SCRIPTDIR/template/day.template > $dest/day$DAY.go
    sed "s/XX/$DAY/" $SCRIPTDIR/template/test.template > $dest/day${DAY}_test.go
else
    echo -e "Usage: $0 DAY\n\twhere day is numeric between 01 and 25"
fi