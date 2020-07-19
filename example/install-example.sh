#!/bin/bash

## Install mezcel/struct-fmt
echo "Installing github.com/mezcel/struct-fmt ... "
go get github.com/mezcel/struct-fmt
echo "done."

## Demo dependencies

## Install andlabs/ui
## Native Window Manager windowed client
## This is a library that aims to provide simple GUI software development in Go
echo "Installing github.com/andlabs/ui ... "
go get github.com/andlabs/ui
echo "done."

## Install nsf/termbox-go
## Plugin for reading cross platform terminal state/settings
## Termbox is a library that provides a minimalistic API which allows the programmer to write text-based user interfaces.
echo "Installing github.com/nsf/termbox-go ... "
go get github.com/nsf/termbox-go
echo "done."

## Install PuerkitoBio/goquery
## Plugin optimised for web scraping
echo "Installing github.com/PuerkitoBio/goquery ... "
go get github.com/PuerkitoBio/goquery
echo "done."