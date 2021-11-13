# keeb

Tiny keybinding utility written in Go. Built for [The Key](https://stackoverflow.blog/2021/03/31/the-key-copy-paste/)

> Note: I don't know why anyone would want to do this, but this is primarily for posterity so I don't forget this entire setup.

## Setup
To use this program you either need an input device that can emit F17, F18 and F19 or use `The Key` with the [firmware](./resources/f17-18-19-mod-shift.hex) I built.

More details at [Drop](https://drop.com/talk/93641/how-to-configure-stack-overflow-the-key-macropad) on how to flash the firmware.

## Installing
`keeb` is a tiny go program that can be compiled with

```shell
go get
go build cmd/keeb.go # Outputs a binary at ./keeb
```

Alternatively you can install it globally via

```shell
go install # Installs to $GOPATH/bin
```

## Running

```shell
keeb # Runs the program if you installed it globally
```

Ideally you'd be running this a daemon on system start.

That's it! `keeb` would respond to the binded keys you can pretty much do anything you want with it.