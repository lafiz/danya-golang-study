#!/bin/bash
declare -A args
args=(["clean"]=true ["verbose"]=false)

build() {
    eval "go build -a -buildmode=exe -trimpath ${1} -o bin/app.exe src/main.go"
}

while [[ $# -gt 0 ]]; do
    case "$1" in
    -v | --verbose)
        args[verbose]=true
        shift
        ;;
    -c | --clean)
        args[clean]=true
        shift
        ;;
    *)
        args[verbose]=false
        args[clean]=false
        ;;
    esac
done

if [[ "${args[clean]}" = true ]]; then
    echo "Cleaning build cache"
    go clean -i -r -cache -testcache -modcache -fuzzcache
    echo "Old compiled binaries removed"
fi

echo "Compiling..."

verbose_options=""

if [[ "${args[verbose]}" = true ]]; then
    verbose_options="-v -x"
fi

build "$verbose_options"

echo "Compiled"
