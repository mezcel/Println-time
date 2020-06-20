# demo

Some ( Rough and Hasty ) use case demo for the ```struct-fmt``` package.
* These demo are under developed. I was/am still learning the in's and outs of Golang. But they work and satisfy the package usecase.

## About

* Inputs a ```.json``` file and render it's contents onto a cli display using Golang with gocui.
* The primary query view mechanics are in place, the rest of the work will be in experimenting with the UI. Ultimately the UX will closely resemble [python-curses]( http://github.com/mezcel/python-curses ) and [printf-time]( http://github.com/mezcel/printf-time ).

## Objective Description

* This is a cli scripture rosary app wittten in Go.
* The rosary database is the same ```.json``` use in [python-curses]( http://github.com/mezcel/python-curses ).
* The ```struct``` ER schema is similar to the one used in [printf-time]( http://github.com/mezcel/printf-time ).

---

## Run

```sh
## Run
go run main.go

## Build
go build main.go -o "myApp.exe"
```

### App Scripts:

| script name | about |
| --- | --- |
|```structs.go```|Go Structs based on Json|
|```functions.go```|App Functions|
|```main-onefile.go```|Standalone Go file which does not import the ```struct-fmt``` package|
|```main.go```|Tui Rosary with the ```mezcel/struct-fmt``` package|
|```main-ui.go```|Gui Rosary with the ```mezcel/struct-fmt``` & the ```andlabs/ui``` package |

### UI Package

Install the ```andlabs/ui``` go package: [https://github.com/andlabs/ui.git](https://github.com/andlabs/ui)
```sh
## install the ui package dependency
go get github.com/andlabs/ui
```

