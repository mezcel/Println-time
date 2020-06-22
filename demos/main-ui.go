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

type ReadingsText struct {
    DecadeName string
    MysteryName string
    ScriptureText string
    MessageText string
    PrayerText string
    Position int
}

var (
    fontButton *ui.FontButton
    alignment *ui.Combobox

    attrstr *ui.AttributedString
    btnNext *ui.Button
    btnPrevious *ui.Button
    btnClose *ui.Button

    lblDecadeName    *ui.Label
    lblMysteryName   *ui.Label
    lblScriptureText *ui.Label
    lblMessageText   *ui.Label
    lblPrayerText    *ui.Label

    textStructs ReadingsText

    // Global struct db variables
    RosaryBeads structfmt.RosaryBeads
    Beads structfmt.Beads
    Decades structfmt.Decades
    Mysterys structfmt.Mysterys
    Books structfmt.Books
    Scriptures structfmt.Scriptures
    Messages structfmt.Messages
    Prayers structfmt.Prayers
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
    textStructs.DecadeName = Decades.Decades[decadeIdx].DecadeName
    textStructs.MysteryName = Mysterys.Mysterys[mysteryIdx].MysteryName
    textStructs.ScriptureText = Scriptures.Scriptures[scriptureIdx].ScriptureText
    textStructs.MessageText = Messages.Messages[messageIdx].MesageText
    textStructs.PrayerText = Prayers.Prayers[prayerIdx].PrayerText
}

// PrintTui(<args>) will render string display strings in tui
func PrintTui() {
    // clear terminal screen
    structfmt.Cls()

    // View query on cli tui
    fmt.Println("Decade:\t\t"   + textStructs.DecadeName)
    fmt.Println("Mystery:\t"    + textStructs.MysteryName)
    fmt.Println("Message:\t"    + textStructs.MessageText)
    fmt.Println("Scripture:\t"  + textStructs.ScriptureText + "\n")
    fmt.Println("Prayer:\t\t"   + textStructs.PrayerText)

    fmt.Println("\n---\nPress the Ctrl+C to Exit")
}

/* *** GUI Configurations *** */

type areaHandler struct{}

// Draw Text Area
func (areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
    tmpString := "Decade:\n\t"      + textStructs.DecadeName +
                 "\n\nMystery:\n\t" + textStructs.MysteryName +
                 "\n\nMessage:\n\t" + textStructs.MessageText +
                 "\n\nScripture:\n\t" + textStructs.ScriptureText +
                 "\n\nPrayer:\n\t"  + textStructs.PrayerText

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

    // Define Forward Navigation Button
    btnNext = ui.NewButton("Next >>")
    btnNext.OnClicked(func(*ui.Button) {
        textStructs.Position = structfmt.NextBead(textStructs.Position)

        UpdateDisplayStrings()
        area.QueueRedrawAll()
        PrintTui()
    })

    // Define Backward Navigation Button
    btnPrevious = ui.NewButton("<< Back")
    btnPrevious.OnClicked(func(*ui.Button) {
        textStructs.Position = structfmt.PreviousBead(textStructs.Position)

        UpdateDisplayStrings()
        area.QueueRedrawAll()
        PrintTui()
    })

    // Place Navigation Buttons
    vbox.Append(btnPrevious, false)
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
    UpdateDisplayStrings()
    PrintTui()

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
    textStructs.Position = structfmt.ReturnStartPosition(WeekdayNo)

    ui.Main(setupUI)

    fmt.Println("done")
}
