/*
 * file         : main-ui.go
 * dependencies : go get github.com/mezcel/struct-fmt
 *                go get github.com/andlabs/ui
 * about:       : Based on andlabs' drawtext.go example script
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

var (
    fontButton *ui.FontButton
    alignment *ui.Combobox

    attrstr *ui.AttributedString
    btnNext *ui.Button
    btnClose *ui.Button

    lblDecadeName    *ui.Label
    lblMysteryName   *ui.Label
    lblScriptureText *ui.Label
    lblMessageText   *ui.Label
    lblPrayerText    *ui.Label
)

// Global string vars used for display readings
var DecadeName string = "DecadeName"
var MysteryName string = "MysteryName"
var ScriptureText string = "ScriptureText"
var MessageText string = "MessageText"
var PrayerText string = "PrayerText"
var Accumulator int = 0

// Global struct db variables
var RosaryBeads structfmt.RosaryBeads
var Beads structfmt.Beads
var Decades structfmt.Decades
var Mysterys structfmt.Mysterys
var Books structfmt.Books
var Scriptures structfmt.Scriptures
var Messages structfmt.Messages
var Prayers structfmt.Prayers

/* *** Text Readings Vaues *** */

// updateGlobals(<args>) will update global string vars
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

// printTui(<args>) will render string display strings in tui
func printTui(DecadeName string, MysteryName string, MessageText string, ScriptureText string, PrayerText string) {
    // clear terminal screen
    structfmt.Cls()

    // View query on cli tui
    fmt.Println("Decade:\t\t" + DecadeName)
    fmt.Println("Mystery:\t" + MysteryName)
    fmt.Println("Message:\t" + MessageText)
    fmt.Println("Scripture:\t" + ScriptureText + "\n")

    fmt.Println("Prayer:\t\t" + PrayerText)

    fmt.Println("\n---\nPress the Ctrl+C to Exit")
}

// updateUI(Accumulator int) will update global string vars and render string display strings in tui
func updateUI(Accumulator int) {
    updateGlobals(RosaryBeads, Beads, Decades, Mysterys, Books, Scriptures, Messages, Prayers, Accumulator)
    printTui(DecadeName, MysteryName, MessageText, ScriptureText, PrayerText)
}

/* *** GUI Configurations *** */

type areaHandler struct{}

// Draw Text Area
func (areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
    tmpString := "Decade:\n\t" + DecadeName + "\n\nMystery:\n\t" + MysteryName + "\n\nMessage:\n\t" + MessageText + "\n\nScripture:\n\t" + ScriptureText + "\n\nPrayer:\n\t" + PrayerText

    attrstr = ui.NewAttributedString(tmpString)

    tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
        String:     attrstr,
        DefaultFont:    fontButton.Font(),
        Width:      p.AreaWidth,
        Align:      ui.DrawTextAlign(alignment.Selected()),
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
    // reject all keys
    return false
}

// setupUI() is the Native GUI layout design
func setupUI() {

    // Define Main Window
    mainwin := ui.NewWindow("Go Rosary GUI - andlabs/ui", 800, 500, true)
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
    mainwin.SetChild(hbox)

    vbox := ui.NewVerticalBox()
    vbox.SetPadded(true)
    hbox.Append(vbox, false)

    // Place Text Attributes Label
    vbox.Append(ui.NewLabel("Text Attributes:"), false)

    area := ui.NewArea(areaHandler{})

    // Define Font Button
    fontButton = ui.NewFontButton()
    fontButton.OnChanged(func(*ui.FontButton) {
        area.QueueRedrawAll()
    })

    // Place Font Button
    vbox.Append(fontButton, false)

    // Text Area Form

    form := ui.NewForm()
    form.SetPadded(true)
    vbox.Append(form, false)

    alignment = ui.NewCombobox()
    // note that the items match with the values of the uiDrawTextAlign values
    alignment.Append("Left")
    alignment.Append("Center")
    alignment.Append("Right")
    alignment.SetSelected(0)        // start with left alignment
    alignment.OnSelected(func(*ui.Combobox) {
        area.QueueRedrawAll()
    })

    hbox.Append(area, true)
    form.Append("Alignment", alignment, false)

    // Place Separator
    vbox.Append(ui.NewHorizontalSeparator(), false)

    // Place Bead Navigation Label
    vbox.Append(ui.NewLabel("Bead Navigation:"), false)

    // Define Navigation Button
    btnNext = ui.NewButton("Next Bead >>")
    btnNext.OnClicked(func(*ui.Button) {
        Accumulator = structfmt.NextBead(Accumulator)
        updateUI(Accumulator)
        area.QueueRedrawAll()
    })

    // Place Navigation Button
    vbox.Append(btnNext, false)

    // Place Separator
    vbox.Append(ui.NewHorizontalSeparator(), false)

    // Place Exit Label
    vbox.Append(ui.NewLabel("Quit:"), false)

    // Define Close Button
    btnClose = ui.NewButton("Close GUI Window")
    btnClose.OnClicked(func(*ui.Button) {
        ui.Quit()
    })

    // Place Close Button
    vbox.Append(btnClose, false)

    // Update position counter and global text variables
    updateGlobals(RosaryBeads, Beads, Decades, Mysterys, Books, Scriptures, Messages, Prayers, Accumulator)
    printTui(DecadeName, MysteryName, MessageText, ScriptureText, PrayerText)

    mainwin.Show()
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
    Accumulator = structfmt.ReturnStartPosition(WeekdayNo)

    ui.Main(setupUI)
}
