#!/bin/bash
nohup ./client -teamID=$1 -ip="$2" -port=$3 >../battle.out 2>&1 &
