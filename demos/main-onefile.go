/** file: main-onefile.go
run: go run myApp.go
build: go build myApp.go -o "myApp.exe" */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Structs

type RosaryBead struct {
	RosaryBeadID     int `json:"rosaryBeadID"`
	BeadIndex        int `json:"beadIndex"`
	DecadeIndex      int `json:"decadeIndex"`
	MysteryIndex     int `json:"mysteryIndex"`
	PrayerIndex      int `json:"prayerIndex"`
	ScriptureIndex   int `json:"scriptureIndex"`
	MessageIndex     int `json:"messageIndex"`
	LoopBody         int `json:"loopBody"`
	SmallbeadPercent int `json:"smallbeadPercent"`
	MysteryPercent   int `json:"mysteryPercent"`
}

type Bead struct {
	BeadID   int    `json:"beadID"`
	BeadType string `json:"beadType"`
}

type Decade struct {
	DecadeID       int    `json:"beadID"`
	MysteryIndex   int    `json:"mysteryIndex"`
	DecadeNo       int    `json:"decadeNo"`
	DecadeName     string `json:"decadeName"`
	DecadeInfo     string `json:"decadeInfo"`
	InfoRefference string `json:"infoRefference"`
}

type Mystery struct {
	MysteryID   int    `json:"mysteryID"`
	MysteryNo   int    `json:"mysteryNo"`
	MysteryName string `json:"mysteryName"`
}

type Book struct {
	BookID   int    `json:"bookID"`
	BookName string `json:"bookName"`
}

type Scripture struct {
	ScriptureID   int    `json:"scriptureID"`
	BookIndex     int    `json:"bookIndex"`
	ChapterIndex  int    `json:"chapterIndex"`
	VerseIndex    int    `json:"verseIndex"`
	ScriptureText string `json:"scriptureText"`
}

type Message struct {
	MessageID  int    `json:"messageID"`
	MesageText string `json:"mesageText"`
}

type Prayer struct {
	PrayerID   int    `json:"prayerID"`
	PrayerName string `json:"prayerName"`
	PrayerText string `json:"prayerText"`
}

// ER DB

type RosaryBeads struct {
	RosaryBeads []RosaryBead `json:"rosaryBead"`
}

type Beads struct {
	Beads []Bead `json:"bead"`
}

type Decades struct {
	Decades []Decade `json:"decade"`
}

type Mysterys struct {
	Mysterys []Mystery `json:"mystery"`
}

type Books struct {
	Books []Book `json:"book"`
}

type Scriptures struct {
	Scriptures []Scripture `json:"scripture"`
}

type Messages struct {
	Messages []Message `json:"message"`
}

type Prayers struct {
	Prayers []Prayer `json:"prayer"`
}

func cls() {
	// clear screen

	// insired from: https://stackoverflow.com/a/22896706

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

// Functions / Methods

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

// Main

func main() {

	// Open a jsonFile
	var jsonPath string = "json/rosaryJSON-nab.json"
	var byteValue []byte = returnByteValue(jsonPath)

	// Make a struct DB from a json file

	var rosaryBeads RosaryBeads
	var beads Beads
	var decades Decades
	var mysterys Mysterys
	var books Books
	var scriptures Scriptures
	var messages Messages
	var prayers Prayers

	json.Unmarshal(byteValue, &rosaryBeads)
	json.Unmarshal(byteValue, &beads)
	json.Unmarshal(byteValue, &decades)
	json.Unmarshal(byteValue, &mysterys)
	json.Unmarshal(byteValue, &books)
	json.Unmarshal(byteValue, &scriptures)
	json.Unmarshal(byteValue, &messages)
	json.Unmarshal(byteValue, &prayers)

	var weekdayNo int = int(time.Now().Weekday())
	var accumulator int = returnStartPosition(weekdayNo)

	// Main Loop
	for accumulator < 315 {
		cls() // clear terminal screen

		// position progress step increment
		accumulator++

		var decadeIdx int = rosaryBeads.RosaryBeads[accumulator].DecadeIndex
		var mysteryIdx int = rosaryBeads.RosaryBeads[accumulator].MysteryIndex
		var prayerIdx int = rosaryBeads.RosaryBeads[accumulator].PrayerIndex
		var scriptureIdx int = rosaryBeads.RosaryBeads[accumulator].ScriptureIndex
		var messageIdx int = rosaryBeads.RosaryBeads[accumulator].MessageIndex

		var decadeName string = decades.Decades[decadeIdx].DecadeName
		var mysteryName string = mysterys.Mysterys[mysteryIdx].MysteryName
		var scriptureText string = scriptures.Scriptures[scriptureIdx].ScriptureText
		var messageText string = messages.Messages[messageIdx].MesageText
		var prayerText string = prayers.Prayers[prayerIdx].PrayerText

		fmt.Println("Decade:    " + decadeName)
		fmt.Println("Mystery:   " + mysteryName)
		fmt.Println("Message:   " + messageText)
		fmt.Println("Scripture:\n\t   " + scriptureText)
		fmt.Println("\nPrayer:\n\t   " + prayerText)

		fmt.Println("\n---\nPress the Enter to continue / Ctrl+C to Exit")
		fmt.Scanln()
	}
}
