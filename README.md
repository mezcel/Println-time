# struct-fmt

## About

* This is a Go package which inputs a formatted scripture rosary json and parses them into structs.
* The json must have the same ER Schema found in: [python-curses]( http://github.com/mezcel/python-curses ) or [printf-time]( http://github.com/mezcel/printf-time ).

## Install

Install Go: [golang.org](https://golang.org/dl/)

```go
// Install struct-fmt
go get github.com/mezcel/struct-fmt

/*  // Uninstall
    go clean -i -x github.com/mezcel/struct-fmt */
```

## Usage

```go
import "github.com/mezcel/struct-fmt"
```

## Use Case Demo
Demo App:
    ```./demos/main.go```

---

# Go Development Notes

## Setup a Local Go Workspace

### Linux (Debian)

```sh
## Go workspace directory structure
mkdir -p $HOME/go{bin,src}

## Set local environment variables (obsolete, but still good practice)

## Run Go apps anywhere on system
## also make aliases in .bashrc or .zshrc, .profile, ect.

echo -e "\n## Go Workspace Environment Variable Alias\n" >> ~/.profile
echo 'export GOPATH=$HOME/go' >> ~/.profile

## Make Go tools available on system
echo 'export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin' >> ~/.profile
```

### Win10
```ps1
## Go workspace directory structure
mkdir ~/go/bin, ~/go/src

## Chocolately should have made $GOPATH upon Go installation
## Add the workspace's ```bin``` subdirectory to path
setx PATH "$($env:path);$GOPATH\bin"
```

## Make personal Go packages

```sh
## Create future projects in the following directory structure
$GOPATH/src/github.com/<user-name>/<project-name>
```

#### Import a Github published Go app into another Go project
```go
go get github.com/<user-name>/<project-name>
```

#### Make binaries
```sh
cd $GOPATH/src/github.com/<user-name>/<project-name>

go install
go build
```