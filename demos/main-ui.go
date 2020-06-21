/* file		: main-ui.go
dependencies: go get github.com/mezcel/struct-fmt
			  go get github.com/andlabs/ui */

package main

import (
	"strings"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/buger/goterm"

	"encoding/json"
	"fmt"
	"time"

	structfmt "github.com/mezcel/struct-fmt"
)

var (
	btnNext *ui.Button

	lblDecadeName    *ui.Label
	lblMysteryName   *ui.Label
	lblScriptureText *ui.Label
	lblMessageText   *ui.Label
	lblPrayerText    *ui.Label
)

// Query attribute strings
var DecadeName string = "na"
var MysteryName string = "na"
var ScriptureText string = "na"
var MessageText string = "na"
var PrayerText string = "na"
var Accumulator int = 0

// Make a struct DB from a json file

// Declare local struct variables
var RosaryBeads structfmt.RosaryBeads
var Beads structfmt.Beads
var Decades structfmt.Decades
var Mysterys structfmt.Mysterys
var Books structfmt.Books
var Scriptures structfmt.Scriptures
var Messages structfmt.Messages
var Prayers structfmt.Prayers

func WrapCentence(terminalWidth int, inSentence string) string {
	var strOutput string = ""

	wordsArr := strings.Fields(inSentence)
	wordCount := len(wordsArr)
	counter := 0

	for i := 0; i < wordCount; i++ {
		counter = counter + len(wordsArr[i]) + 1

		if counter <= terminalWidth {
			strOutput = strOutput + wordsArr[i] + " "
		} else {
			strOutput = strOutput + wordsArr[i] + "\n\t\t"
			counter = 0
		}
	}

	return strOutput
}

func updateGlobals(RosaryBeads structfmt.RosaryBeads, Beads structfmt.Beads, Decades structfmt.Decades, Mysterys structfmt.Mysterys, Books structfmt.Books, Scriptures structfmt.Scriptures, Messages structfmt.Messages, Prayers structfmt.Prayers, Accumulator int) {

	// Query FKs
	var decadeIdx int = RosaryBeads.RosaryBeads[Accumulator].DecadeIndex
	var mysteryIdx int = RosaryBeads.RosaryBeads[Accumulator].MysteryIndex
	var prayerIdx int = RosaryBeads.RosaryBeads[Accumulator].PrayerIndex
	var scriptureIdx int = RosaryBeads.RosaryBeads[Accumulator].ScriptureIndex
	var messageIdx int = RosaryBeads.RosaryBeads[Accumulator].MessageIndex

	// Query attribute strings
	DecadeName = Decades.Decades[decadeIdx].DecadeName
	MysteryName = Mysterys.Mysterys[mysteryIdx].MysteryName
	ScriptureText = Scriptures.Scriptures[scriptureIdx].ScriptureText
	MessageText = Messages.Messages[messageIdx].MesageText
	PrayerText = Prayers.Prayers[prayerIdx].PrayerText
}

func printTui(DecadeName string, MysteryName string, MessageText string, ScriptureText string, PrayerText string) {
	// clear terminal screen
	structfmt.Cls()

	// View query on cli tui
	fmt.Println("Decade:\t\t" + DecadeName)
	fmt.Println("Mystery:\t" + MysteryName)
	fmt.Println("Message:\t" + MessageText)
	fmt.Println("Scripture:\t" + ScriptureText + "\n")

	terminalWidth := goterm.Width()
	PrayerText = WrapCentence(terminalWidth, PrayerText)

	fmt.Println("Prayer:\t\t" + PrayerText)

	fmt.Println("\n---\nPress the Ctrl+C to Exit")
}

func setLabels(DecadeName string, MysteryName string, MessageText string, ScriptureText string, PrayerText string) {
	lblDecadeName.SetText(DecadeName)
	lblMysteryName.SetText(MysteryName)
	lblScriptureText.SetText(MessageText)
	lblMessageText.SetText(ScriptureText)
	lblPrayerText.SetText(PrayerText)
}

func updateUI(Accumulator int) {
	updateGlobals(RosaryBeads, Beads, Decades, Mysterys, Books, Scriptures, Messages, Prayers, Accumulator)
	printTui(DecadeName, MysteryName, MessageText, ScriptureText, PrayerText)
	setLabels(DecadeName, MysteryName, MessageText, ScriptureText, PrayerText)
}

func setupUI() {
	// Variables are global

	mainwin := ui.NewWindow("Go Rosary GUI - andlabs/ui", 640, 480, true)
	mainwin.SetMargined(true)
	mainwin.OnClosing(func(*ui.Window) bool {
		mainwin.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	mainwin.SetChild(vbox)

	btnNext = ui.NewButton("Next Bead")
	btnNext.OnClicked(func(*ui.Button) {
		Accumulator = structfmt.NextBead(Accumulator)
		updateUI(Accumulator)
	})
	vbox.Append(btnNext, false)

	// Update position counter
	updateGlobals(RosaryBeads, Beads, Decades, Mysterys, Books, Scriptures, Messages, Prayers, Accumulator)
	printTui(DecadeName, MysteryName, MessageText, ScriptureText, PrayerText)

	lblDecadeName = ui.NewLabel(DecadeName)
	lblMysteryName = ui.NewLabel(MysteryName)
	lblScriptureText = ui.NewLabel(MessageText)
	lblMessageText = ui.NewLabel(ScriptureText)
	lblPrayerText = ui.NewLabel(PrayerText)

	vbox.Append(lblDecadeName, false)
	vbox.Append(lblMysteryName, false)
	vbox.Append(lblScriptureText, false)
	vbox.Append(lblMessageText, false)
	vbox.Append(lblPrayerText, false)

	mainwin.Show()
}

func main() {

	// Open a jsonFile
	var jsonPath string = "json/rosaryJSON-nab.json"
	var byteValue []byte = structfmt.ReturnByteValue(jsonPath)
	// Port json data into struct data
	json.Unmarshal(byteValue, &RosaryBeads)
	json.Unmarshal(byteValue, &Beads)
	json.Unmarshal(byteValue, &Decades)
	json.Unmarshal(byteValue, &Mysterys)
	json.Unmarshal(byteValue, &Books)
	json.Unmarshal(byteValue, &Scriptures)
	json.Unmarshal(byteValue, &Messages)
	json.Unmarshal(byteValue, &Prayers)

	// Flag which day of the week it is
	var WeekdayNo int = int(time.Now().Weekday())

	// Initial starting position based on which day of the week it is
	Accumulator = structfmt.ReturnStartPosition(WeekdayNo)

	////////////////

	ui.Main(setupUI)
}
