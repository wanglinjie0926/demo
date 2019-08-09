#!/bin/bash
cd "$(dirname "$0")"
cd bin
chmod 777 ./client
chmod 777 ./start.sh
sh start.sh $1 $2 $3
