/* my-funcs.go */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func cls() {
	// clear screen in Win10

	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func returnByteValue(jsonPath string) []byte {
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

func returnStartPosition(weekdayNo int) int {
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
