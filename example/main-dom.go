/*
 * file         : main-dom.go
 * dependencies : go get github.com/mezcel/struct-fmt
 * about:       : Web browser dom web page hosed by go server
 * git:         : https://github.com/mezcel/struct-fmt/blob/master/demos/main-dom.go
 * */

package example

import (
	"encoding/json"
	"fmt"
	"time"

	structfmt "github.com/mezcel/struct-fmt"

	"html/template"
	"log"
	"net/http"
)

// Global Vars used for UI text display
type ReadingsText struct {
	DecadeName    string
	MysteryName   string
	ScriptureText string
	MessageText   string
	PrayerText    string

	Position         int
	LoopBody         int
	SmallbeadPercent int
	MysteryPercent   int
}

var (
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

	// Intro and Outro Progress
	if textStructs.LoopBody == 0 {
		if textStructs.MysteryPercent == 50 {
			textStructs.SmallbeadPercent = 0
		} else {
			textStructs.SmallbeadPercent = (textStructs.SmallbeadPercent * 10) / 7
		}
	}
}

/* *** Html Configurations *** */

// Initial Page Load
func IndexPage(w http.ResponseWriter, r *http.Request) {

	// Update global struct vars
	UpdateDisplayStrings()

	// Update Html Page strings with struct vars
	UpdateIndexPageVars(w)

}

// Update index.html with the latest server variables
func UpdateIndexPageVars(w http.ResponseWriter) {

	t, err := template.ParseFiles("html/index.html") //parse the html file index.html
	if err != nil {                                  // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, textStructs) //execute the template
	if err != nil {                 // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

// The next button actions
func Next(w http.ResponseWriter, r *http.Request) {

	// Increment position counter
	textStructs.Position = structfmt.NextBead(textStructs.Position)

	// Update global struct vars
	UpdateDisplayStrings()

	// Update Html Page strings with struct vars
	UpdateIndexPageVars(w)

}

// The back button actions
func Back(w http.ResponseWriter, r *http.Request) {

	// Increment position counter
	textStructs.Position = structfmt.PreviousBead(textStructs.Position)

	// Update global struct vars
	UpdateDisplayStrings()

	// Update Html Page strings with struct vars
	UpdateIndexPageVars(w)

}

// Reset back button actions
func Reset(w http.ResponseWriter, r *http.Request) {

	// Flag which day of the week it is
	var WeekdayNo int = int(time.Now().Weekday())

	// Initial starting position based on which day of the week it is
	textStructs.Position = structfmt.ReturnStartPosition(WeekdayNo)

	// Update global struct vars
	UpdateDisplayStrings()

	// Update Html Page strings with struct vars
	UpdateIndexPageVars(w)

}

// Favicon external resource
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/img/favicon.ico")
}

// Style external resource
func CssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/css/w3.css")
}

func CssHandler2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/css/w3-colors-ios.css")
}

// Script external resource
func JsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/js/myScript.js")
}

/* *** Main *** */
func main() {
	fmt.Println("\nStarted main() ...")
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

	// Initial page load
	http.HandleFunc("/", IndexPage)

	// favicon
	http.HandleFunc("/img/favicon.ico", FaviconHandler)

	// css
	http.HandleFunc("/css/w3.css", CssHandler)
	http.HandleFunc("/css/w3-colors-ios.css", CssHandler2)

	// js
	http.HandleFunc("/js/myScript.js", JsHandler)

	// Navigation Page Refresh
	http.HandleFunc("/Next", Next)
	http.HandleFunc("/Back", Back)
	http.HandleFunc("/Reset", Reset)

	// Server at port
	fmt.Println("\n\t- Go Server is running the main-dom.go app ...")
	fmt.Println("\t- Open a web browser and navigate to \"localhost:8080\".")
	fmt.Println("\n( From within this prompt,\n\tpress Ctrl-C to terminate server hosting. )\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
