@ECHO OFF

REM Install mezcel/struct-fmt
REM My app package
ECHO Installing github.com/mezcel/struct-fmt ...
go get -u github.com/mezcel/struct-fmt
ECHO done.

REM Install andlabs/ui
REM This is a library that aims to provide simple GUI software development in Go
ECHO Installing github.com/andlabs/ui ...
go get -u github.com/andlabs/ui
ECHO done.

REM Install nsf/termbox-go
REM Termbox is a library that provides a minimalistic API which allows the programmer to write text-based user interfaces.
ECHO Installing github.com/nsf/termbox-go ... 
go get github.com/nsf/termbox-go
ECHO done.

REM Install PuerkitoBio/goquery
REM Plugin optimised for web scraping
ECHO Installing github.com/PuerkitoBio/goquery ... 
go get github.com/PuerkitoBio/goquery
ECHO done.

