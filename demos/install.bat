@ECHO OFF

REM Install mezcel/struct-fmt
ECHO Installing github.com/mezcel/struct-fmt ...
go get -u github.com/mezcel/struct-fmt
ECHO done.

REM Install andlabs/ui
ECHO Installing github.com/andlabs/ui ...
go get -u github.com/andlabs/ui
ECHO done.

REM Install nsf/termbox-go
ECHO Installing github.com/nsf/termbox-go ... 
go get github.com/nsf/termbox-go
ECHO done.
