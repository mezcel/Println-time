# Println-time

## About

* Inputs a ```.json``` file and render it's contents onto a cli display using Golang with gocui.
* The primary query view mechanics are in place, the rest of the work will be in experimenting with the UI. Ultimately the UX will closely resemble [python-curses]( http://github.com/mezcel/python-curses ) and [printf-time]( http://github.com/mezcel/printf-time ).
* The "curses", the display formatting, and the navigational control options are not designed or implemented yet.

## Objective Description

* This is a cli scripture rosary app wittten in Go.
* The rosary database is the same ```json``` use in [python-curses]( http://github.com/mezcel/python-curses ). The ```struct``` ER schema is similar to the one used in [printf-time]( http://github.com/mezcel/printf-time ).
* This app was/is an app used to help me learn Go.

## Install Dependency

Install Go: [golang.org](https://golang.org/dl/)

```go
// install the gocui package dependency
go get github.com/jroimartin/gocui
```
## Run

```go
// Run
go run my-structs.go my-funcs.go main.go

// Build as an Exe
go build (my-structs.go my-funcs.go main.go) -o "myApp.exe"
```

## Dependencies

### App Scripts:

| script name | about |
| --- | --- |
|```my-structs.go```|Json Structs|
|```my-funcs.go```|App Functions|
|```main.go```|Main()|

### Additional Go Packages

| package | install | about |
| --- | --- | --- |
|[gocui.git](https://github.com/jroimartin/gocui)| ```go get github.com/jroimartin/gocui``` | Go package aimed at creating Console User Interfaces. It is tput-like/curses-like |
| [godoc.org](https://godoc.org/github.com/jroimartin/gocui) |```go doc github.com/jroimartin/gocui``` | Documentation

---

## Notes and Thoughts

Input a ```.json`` file and render it's contents onto a cli display using Golang with gocui.