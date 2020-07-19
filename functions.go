/* functions.go */

package structfmt

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

// Cls() will clear the cli/tui screen
func Cls() {
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

// ReturnByteValue(jsonPath string) will returns a byte[] of the json file input
func ReturnByteValue(jsonPath string) []byte {
	// Import and return json file

	jsonFile, err := os.Open(jsonPath)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened", jsonPath)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

// ReturnStartPosition(weekdayNo int) will return an int representing which day of the week it currently is.
// Inputs an int 0-6, where Sunday = 0, and 6 = Saturday
// Outputs and int equivalent to the primary key of the start of a mystery
func ReturnStartPosition(weekdayNo int) int {
	var positionNo int = 0
	//var weekdayNo int = int(time.Now().Weekday())

	switch weekdayNo {
	case 0: // Sunday
		positionNo = 237
		break
	case 1: // Monday
	case 5: // Friday
		positionNo = 0
		break
	case 2: // Tuesday
	case 6: // Saturday
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

// NextBead(accumulator int) will Return an int representing the next bead sequence position
func NextBead(accumulator int) int {
	// sequential navigation
	if accumulator < 315 {
		// forward progress
		accumulator++
	} else {
		// loop back to start
		accumulator = 0
	}

	return accumulator
}

// PreviousBead(accumulator int) will Return an int representing the next bead sequence position
func PreviousBead(accumulator int) int {
	// sequential navigation
	if accumulator > 0 {
		// forward progress
		accumulator--
	} else {
		// loop back to end
		accumulator = 315
	}

	return accumulator
}
