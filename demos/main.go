/* file     : main.go
dependencies: go get github.com/mezcel/struct-fmt */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	structfmt "github.com/mezcel/struct-fmt"
	"github.com/nsf/termbox-go"
)

func SetTtyKeyInput() {
	if runtime.GOOS == "linux" {
		fmt.Println("Hello from linux")

		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	}
}

func GetTUIKeyboardKeys(keyPress []byte, accumulator *int) {
	os.Stdin.Read(keyPress)

	var keyString string = string(keyPress)

	switch keyString {
	case "q": // quit
		fmt.Println("Quit")
		exec.Command("reset").Run()
		os.Exit(3)

	case "h": // back
		//*accumulator--
		*accumulator = structfmt.PreviousBead(*accumulator)

	case "l": // next
		//*accumulator++
		*accumulator = structfmt.NextBead(*accumulator)
	}

}

// Get column width of terminal display
func ReturnTermboxWidth() int {
	// Requires github.com/nsf/termbox-go import

	if err := termbox.Init(); err != nil {
		panic(err)
	}

	readingsWidth, _ := termbox.Size()
	termbox.Close()

	return readingsWidth
}

// Convert strin into an array of words
func CentenceArray(inStr string) []string {
	// requires strings import

	var centenceArr []string = strings.Split(inStr, " ")
	return centenceArr
}

// Perform custom word wrapping based on a defined char width
// Input a string and the number of chars to wrap long strings
func IndentedWrap(strOrig string, readingsWidth int) string {
	var newString string = ""
	var centenceArr []string = CentenceArray(strOrig)
	var centenceArrLen int = len(centenceArr)

	var charCount int = 0
	var i int = 0
	var wordLength int = 0

	for i = 0; i < centenceArrLen; i++ {

		if charCount < readingsWidth {
			newString += centenceArr[i] + " "
		} else {
			charCount = 0
			newString += "\n\t\t" + centenceArr[i] + " "
		}

		// word with trailing space
		wordLength = len(centenceArr[i]) + 1
		charCount += wordLength
	}

	return newString
}

func main() {

	// Open a jsonFile
	var jsonPath string = "json/rosaryJSON-nab.json"
	var byteValue []byte = structfmt.ReturnByteValue(jsonPath)

	// Make a struct DB from a json file

	// Declare local struct variables
	var rosaryBeads structfmt.RosaryBeads
	var beads structfmt.Beads
	var decades structfmt.Decades
	var mysterys structfmt.Mysterys
	var books structfmt.Books
	var scriptures structfmt.Scriptures
	var messages structfmt.Messages
	var prayers structfmt.Prayers

	// Port json data into struct data
	json.Unmarshal(byteValue, &rosaryBeads)
	json.Unmarshal(byteValue, &beads)
	json.Unmarshal(byteValue, &decades)
	json.Unmarshal(byteValue, &mysterys)
	json.Unmarshal(byteValue, &books)
	json.Unmarshal(byteValue, &scriptures)
	json.Unmarshal(byteValue, &messages)
	json.Unmarshal(byteValue, &prayers)

	// Main Loop

	// Flag which day of the week it is
	var weekdayNo int = int(time.Now().Weekday())

	// Initial starting position based on which day of the week it is
	var accumulator int = structfmt.ReturnStartPosition(weekdayNo)

	var keyPress []byte = make([]byte, 1)
	SetTtyKeyInput()

	// Main loop
	for accumulator <= 315 {
		// clear terminal screen
		structfmt.Cls()

		// Query FKs
		var decadeIdx int = rosaryBeads.RosaryBeads[accumulator].DecadeIndex
		var mysteryIdx int = rosaryBeads.RosaryBeads[accumulator].MysteryIndex
		var prayerIdx int = rosaryBeads.RosaryBeads[accumulator].PrayerIndex
		var scriptureIdx int = rosaryBeads.RosaryBeads[accumulator].ScriptureIndex
		var messageIdx int = rosaryBeads.RosaryBeads[accumulator].MessageIndex

		// Query attribute strings
		var decadeName string = decades.Decades[decadeIdx].DecadeName
		var mysteryName string = mysterys.Mysterys[mysteryIdx].MysteryName
		var scriptureText string = scriptures.Scriptures[scriptureIdx].ScriptureText
		var messageText string = messages.Messages[messageIdx].MesageText
		var prayerText string = prayers.Prayers[prayerIdx].PrayerText

		// Set the carrage return length 
		// based on terminal width and tab space from the query lables.
		var readingsWidth int = ReturnTermboxWidth() - 21

		// View query on cli tui
		decadeName = IndentedWrap(decadeName, readingsWidth)
		fmt.Println("Decade:\t\t" + decadeName)

		mysteryName = IndentedWrap(mysteryName, readingsWidth)
		fmt.Println("Mystery:\t" + mysteryName)

		messageText = IndentedWrap(messageText, readingsWidth)
		fmt.Println("Message:\t" + messageText)

		scriptureText = IndentedWrap(scriptureText, readingsWidth)
		fmt.Println("Scripture:\t" + scriptureText + "\n")

		prayerText = IndentedWrap(prayerText, readingsWidth)
		fmt.Println("Prayer:\t\t" + prayerText)

		// Pause for read display
		// Keyboard navigation input
		fmt.Printf("-------------\t-------------------------------------\nControl Keys:\t( h = back, l = next, q = quit ) ")

		// decorative prompt workarround for win10 terminals
		if runtime.GOOS == "windows" {
			fmt.Printf("?: ")
		}

		GetTUIKeyboardKeys(keyPress, &accumulator)
	}
}
