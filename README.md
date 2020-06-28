# struct-fmt

## About

* This is a Go package used to port data from a ```.json``` file into Go structs.
* The imported ```.json``` file  must have the same ER Schema used in: [python-curses]( http://github.com/mezcel/python-curses ), [printf-time]( http://github.com/mezcel/printf-time ), [jq-tput-terminal](https://github.com/mezcel/jq-tput-terminal), or [electron-container](https://github.com/mezcel/electron-container).
    * Demo Json: [demos/json](demos/json)

Run Use Case  Demos [README](demos)
```sh
## Install struct-fmt
go get github.com/mezcel/struct-fmt

## Navigate into the repo's demo directory
cd demos/

## launch scripts
go run main.go          ## Tui Rosary with mezcel/struct-fmt package
go run main-ui.go       ## Gui Rosary with mezcel/struct-fmt & andlabs/ui package
go run main-onefile.go  ## Tui Rosary without external packages
```

---

## Install

* Install Go: [download](https://golang.org/dl/)
* Install this Go package
```sh
## Install struct-fmt
go get github.com/mezcel/struct-fmt
```
* Background on Go workspaces: [link](#go-development-notes)
* The default Go repo location
    * **Win10:** > ```%USERPROFILE%\go\src\github.com\mezcel\struct-fmt```
    * **Debian:** > ```~/go/src/github.com/mezcel/struct-fmt```

* Update an existing package
    ```sh
    go get -u github.com/mezcel/struct-fmt
    ```

---

## Code Usage

Use this package in a Go script.
```go
// Import this package into a Go program
import "github.com/mezcel/struct-fmt"

// or
import structfmt "github.com/mezcel/struct-fmt"
```

## Use Case Demo
Demo Apps:
* Scripture Rosary Go App: ```./demos/main.go```
    * [README.md](demos/README.md)
* Run Demo
    ```sh
    ## install my packages
    go get github.com/mezcel/struct-fmt

    ## cd into demo directory
    cd $GOPATH/src/github.com/mezcel/struct-fmt/demos/

    ## install an additional package to easily get terminal width 
    go get github.com/nsf/termbox-go

    ## run script
    go run main.go
    ```

---

# Go Development Notes

## Setup a Local Go Workspace

### Linux (Debian)

Do this if the system did not auto configure for you.

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