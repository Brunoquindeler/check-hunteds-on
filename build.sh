#!/usr/bin/zsh

# Linux
build_linux() {
    rm -f ./builds/linux/check-hunteds-linux
    go build -o ./builds/linux/check-hunteds-linux .
    cp -n hunteds.yml ./builds/linux/
}


# Windows
build_windows() {
    rm -f ./builds/windows/check-hunteds.exe
    go generate
    env GOOS=windows GOARCH=amd64 go build -o ./builds/windows/check-hunteds.exe .
    cp -n hunteds.yml ./builds/windows/
}

rm ./resource.syso

build_linux

build_windows