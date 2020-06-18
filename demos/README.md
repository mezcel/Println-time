# demo

Use case demos for the struct-fmt package

## About

* Inputs a ```.json``` file and render it's contents onto a cli display using Golang with gocui.
* The primary query view mechanics are in place, the rest of the work will be in experimenting with the UI. Ultimately the UX will closely resemble [python-curses]( http://github.com/mezcel/python-curses ) and [printf-time]( http://github.com/mezcel/printf-time ).

## Objective Description

* This is a cli scripture rosary app wittten in Go.
* The rosary database is the same ```.json``` use in [python-curses]( http://github.com/mezcel/python-curses ).
* The ```struct``` ER schema is similar to the one used in [printf-time]( http://github.com/mezcel/printf-time ).

### App Scripts:

| script name | about |
| --- | --- |
|```structs.go```|Json Structs|
|```functions.go```|App Functions|


## Run

```go
// Run
go run main.go

// Build on Win10
go Build main.go -o "myApp.exe"
```