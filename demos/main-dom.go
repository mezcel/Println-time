/*
 * file         : main-dom.go
 * dependencies : go get github.com/mezcel/struct-fmt
                  go get github.com/nsf/termbox-go
 * about:       : Web browser dom web page hosed by go server
 * git:         : https://github.com/mezcel/struct-fmt/blob/master/demos/main-dom.go
 * */

package main

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

type PageVariables struct {
	Date string
	Time string

	DecadeName    string
	MysteryName   string
	ScriptureText string
	MessageText   string
	PrayerText    string
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
}

/* *** Html Configurations *** */

func IndexPage(w http.ResponseWriter, r *http.Request) {

	// Increment position counter
	textStructs.Position = textStructs.Position + 1

	// Update global struct vars
	UpdateDisplayStrings()

	now := time.Now()              // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),

		DecadeName:    textStructs.DecadeName,
		MysteryName:   textStructs.MysteryName,
		ScriptureText: textStructs.ScriptureText,
		MessageText:   textStructs.MessageText,
		PrayerText:    textStructs.PrayerText,
	}

	t, err := template.ParseFiles("index.html") //parse the html file index.html
	if err != nil {                             // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, HomePageVars) //execute the template and pass it the PageVariables struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}

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
	textStructs.Position = structfmt.ReturnStartPosition(WeekdayNo) - 1

	http.HandleFunc("/", IndexPage)
	//http.HandleFunc("/Navigate", Next)

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("done")
}