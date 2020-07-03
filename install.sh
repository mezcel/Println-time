#!/bin/bash

## Install mezcel/struct-fmt
echo "Installing github.com/mezcel/struct-fmt ... "
go get github.com/mezcel/struct-fmt
echo "done."

## Demo dependencies

## Install andlabs/ui
## Used for Native Gui
echo "Installing github.com/andlabs/ui ... "
go get github.com/andlabs/ui
echo "done."

## Install nsf/termbox-go
## Plugin for reading cross platform terminal state/settings
echo "Installing github.com/nsf/termbox-go ... "
go get github.com/nsf/termbox-go
echo "done."
