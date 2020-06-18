/* functions.go */

package myfunc

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func Cls() {
	// clear screen

	// insired from from: https://stackoverflow.com/a/22896706

	var clear map[string]func() //create a map for storing clear funcs

	clear = make(map[string]func()) //Initialize it

	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.

	if ok { //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func ReturnByteValue(jsonPath string) []byte {
	// Import and return json file

	jsonFile, err := os.Open(jsonPath)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func ReturnStartPosition(weekdayNo int) int {
	var positionNo int = 0
	//var weekdayNo int = int(time.Now().Weekday())

	switch weekdayNo {
	case 0: // Sunday
		positionNo = 237
		break
	case 1: // Monday
	case 5: // Saturday
		positionNo = 0
		break
	case 2: // Tuesday
	case 6: // Friday
		positionNo = 158
		break
	case 3: // Wednesday
		positionNo = 237
		break
	case 4: // Thursday
		positionNo = 79
		break
	}

	return positionNo
}
