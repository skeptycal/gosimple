#!/usr/bin/env zsh

 check_last () {
        local result=$?
        (( result == 0 )) || { echo "$msg"; return 1; }
        local _def_msg=${default_message:="command failure (pid: $$"}
        local msg="$@"
        msg=${msg:=${_def_msg}}
        # TODO: log the error if logger is active
    }

set -e
echo "" >| coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go build -race -v .
    check_last "failed to compile"
    go test -race -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
