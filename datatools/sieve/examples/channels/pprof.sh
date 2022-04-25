#!/bin/zsh

i=0
rm -rf ./heap* 2>&1 >/dev/null
while true; do
    echo "taking pprof snapshot to heap.${i}.pprof"
    curl http://localhost:8080/debug/pprof/heap >| "heap.${i}.pprof" 2>/dev/null
    i=$(( i+1 ))
    sleep 3
done;