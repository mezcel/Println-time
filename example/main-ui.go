/*
 * file         : main-ui.go
 * dependencies : go get github.com/mezcel/struct-fmt
 *                go get github.com/andlabs/ui
 * about:       : This is a demo example using an OS's native desktop environment UI client.
 * git:         : https://github.com/mezcel/struct-fmt/blob/master/example/main-ui.go
 * */
package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"

	"encoding/json"
	"fmt"
	"time"

	structfmt "github.com/mezcel/struct-fmt"
)

// Globa Vars used for UI text display
type ReadingsText struct {
	DecadeName       string
	MysteryName      string
	ScriptureText    string
	MessageText      string
	PrayerText       string
	Position         int
	LoopBody         int
	SmallbeadPercent int
	MysteryPercent   int
}

var (
	// UI Variables
	fontButton *ui.FontButton
	alignment  *ui.Combobox

	attrstr     *ui.AttributedString
	btnNext     *ui.Button
	btnPrevious *ui.Button
	btnClose    *ui.Button

	lblDecadeName    *ui.Label
	lblMysteryName   *ui.Label
	lblScriptureText *ui.Label
	lblMessageText   *ui.Label
	lblPrayerText    *ui.Label

	// Display string Struct
	textStructs ReadingsText

	// Global struct db variables
	RosaryBeads structfmt.RosaryBeads
	Beads       structfmt.Beads
	Decades     structfmt.Decades
	Mysterys    structfmt.Mysterys
	Books       structfmt.Books
	Scriptures  structfmt.Scriptures
	Messages    structfmt.Messages
	Prayers     structfmt.Prayers
)

/* *** Text Readings Values *** */

// UpdateDisplayStrings() will update global string vars
func UpdateDisplayStrings() {

	var idx int = textStructs.Position

	// Query FKs
	var decadeIdx int = RosaryBeads.RosaryBeads[idx].DecadeIndex
	var mysteryIdx int = RosaryBeads.RosaryBeads[idx].MysteryIndex
	var prayerIdx int = RosaryBeads.RosaryBeads[idx].PrayerIndex
	var scriptureIdx int = RosaryBeads.RosaryBeads[idx].ScriptureIndex
	var messageIdx int = RosaryBeads.RosaryBeads[idx].MessageIndex

	// Query attribute strings

	textStructs.LoopBody = RosaryBeads.RosaryBeads[idx].LoopBody
	textStructs.SmallbeadPercent = RosaryBeads.RosaryBeads[idx].SmallbeadPercent
	textStructs.MysteryPercent = RosaryBeads.RosaryBeads[idx].MysteryPercent

	textStructs.DecadeName = Decades.Decades[decadeIdx].DecadeName

	textStructs.MysteryName = Mysterys.Mysterys[mysteryIdx].MysteryName
	textStructs.ScriptureText = Scriptures.Scriptures[scriptureIdx].ScriptureText
	textStructs.MessageText = Messages.Messages[messageIdx].MesageText
	textStructs.PrayerText = Prayers.Prayers[prayerIdx].PrayerText
}

/* *** GUI Configurations *** */

// Append decorated string
func AppendWithAttributes(what string, attrs ...ui.Attribute) {
	start := len(attrstr.String())
	end := start + len(what)
	attrstr.AppendUnattributed(what)
	for _, a := range attrs {
		attrstr.SetAttribute(a, start, end)
	}
}

type areaHandler struct{}

// Draw Text Area
func (areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {

	// Formatted display string
	attrstr = ui.NewAttributedString("")
	AppendWithAttributes("Decade:\n\t", ui.TextWeightBold)
	attrstr.AppendUnattributed(textStructs.DecadeName)
	AppendWithAttributes("\n\nMystery:\n\t", ui.TextWeightBold)
	attrstr.AppendUnattributed(textStructs.MysteryName)
	AppendWithAttributes("\n\nMessage:\n\t", ui.TextWeightBold)
	attrstr.AppendUnattributed(textStructs.MessageText)
	AppendWithAttributes("\n\nScripture:\n\t", ui.TextWeightBold)
	attrstr.AppendUnattributed(textStructs.ScriptureText)
	AppendWithAttributes("\n\nPrayer:\n\t", ui.TextWeightBold)
	attrstr.AppendUnattributed(textStructs.PrayerText)

	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String:      attrstr,
		DefaultFont: fontButton.Font(),
		Width:       p.AreaWidth,
		Align:       ui.DrawTextAlign(alignment.Selected()),
	})
	defer tl.Free()
	p.Context.Text(tl, 0, 0)
}

func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// do nothing
}

func (areaHandler) MouseCrossed(a *ui.Area, left bool) {
	// do nothing
}

func (areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {

	switch ke.Key {
	case 113: // q
		ui.Quit()
	default:
		fmt.Println("key press:", ke.Key, ", key Up:", ke.Up)
	}

	return false // Documentation recommends a false return
}

// Segment Progressbar Value
func ProgressSegmentInteger(inputInt int, loopBody int) int {
	var intReturn int = 0

	if loopBody == 1 {
		// decade
		intReturn = inputInt * 10
	} else {
		// intro
		intReturn = (inputInt * 100) / 7
	}

	if intReturn > 100 {
		intReturn = 100
	}

	return intReturn
}

// Total Progressbar Value
func ProgressTotalInteger(inputInt int) int {
	var intReturn int = 0
	intReturn = inputInt * 2

	return intReturn
}

// Update the progress bar display
func UpdateProgressBar(pbarDecade *ui.ProgressBar, pbarMystery *ui.ProgressBar) {

	decadeProgress := ProgressSegmentInteger(textStructs.SmallbeadPercent, textStructs.LoopBody)
	mysteryProgress := ProgressTotalInteger(textStructs.MysteryPercent)

	pbarDecade.SetValue(decadeProgress)
	pbarMystery.SetValue(mysteryProgress)
}

// Forward Navigation
func NextClick(area *ui.Area) {
	textStructs.Position = structfmt.NextBead(textStructs.Position)
	UpdateDisplayStrings()
	area.QueueRedrawAll()
}

// Backward Navigation
func PreviousClick(area *ui.Area) {
	textStructs.Position = structfmt.PreviousBead(textStructs.Position)
	UpdateDisplayStrings()
	area.QueueRedrawAll()
}

// SetupUI() is the Native GUI layout design
func SetupUI() {

	// Define Main Window
	mainwin := ui.NewWindow("Golang Rosary GUI", 400, 600, true)
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

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	hboxFont := ui.NewHorizontalBox()
	hboxFont.SetPadded(true)

	hboxNav := ui.NewHorizontalBox()
	hboxNav.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	mainwin.SetChild(vbox)

	area := ui.NewArea(areaHandler{})

	// Text Area Form
	form := ui.NewForm()
	form.SetPadded(true)

	// Text Alignment Combobox
	alignment = ui.NewCombobox()
	alignment.Append("Left")
	alignment.Append("Center")
	alignment.Append("Right")
	alignment.SetSelected(0) // start with left alignment
	alignment.OnSelected(func(*ui.Combobox) {
		area.QueueRedrawAll()
	})

	form.Append("Alignment", alignment, false)

	// Define Font Button
	fontButton = ui.NewFontButton()
	fontButton.OnChanged(func(*ui.FontButton) {
		area.QueueRedrawAll()
	})

	hboxFont.Append(fontButton, false)
	hboxFont.Append(form, false)
	vbox.Append(hboxFont, false)
	vbox.Append(ui.NewHorizontalSeparator(), false)
	vbox.Append(area, true)

	// Progress Bar
	pbarDecade := ui.NewProgressBar()
	pbarMystery := ui.NewProgressBar()

	// Place Separator
	vbox.Append(ui.NewHorizontalSeparator(), false)

	// Place Bead Navigation Label
	vbox.Append(ui.NewLabel("Bead Navigation:"), false)

	// Define Forward Navigation Button
	btnNext = ui.NewButton(" Next >> ")
	btnNext.OnClicked(func(*ui.Button) {
		NextClick(area)
		UpdateProgressBar(pbarDecade, pbarMystery)
	})

	// Define Backward Navigation Button
	btnPrevious = ui.NewButton(" << Back ")
	btnPrevious.OnClicked(func(*ui.Button) {
		PreviousClick(area)
		UpdateProgressBar(pbarDecade, pbarMystery)
	})

	// Place Navigation Buttons
	hboxNav.Append(btnPrevious, false)
	hboxNav.Append(btnNext, false)
	vbox.Append(hboxNav, false)

	// Place Separator
	vbox.Append(ui.NewHorizontalSeparator(), false)

	// Progress Bars
	vbox.Append(pbarDecade, false)
	vbox.Append(pbarMystery, false)

	// Place Separator
	vbox.Append(ui.NewHorizontalSeparator(), false)

	// Define Close Button
	btnClose = ui.NewButton("Close Window")
	btnClose.OnClicked(func(*ui.Button) {
		ui.Quit()
	})

	// Place Close Button
	vbox.Append(btnClose, false)

	// Update position counter and global text variables
	UpdateDisplayStrings()

	mainwin.Show()
	fmt.Println("\tmain-ui.go is running a native desktop manager window...")
}

/* *** Main *** */
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
	textStructs.Position = structfmt.ReturnStartPosition(WeekdayNo)

	ui.Main(SetupUI)

	fmt.Println("Done.\n\tApp is terminated.")
}
