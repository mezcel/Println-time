/** file: main2.go */

package source

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	// Open a jsonFile
	var jsonPath string = "json/rosaryJSON-nab.json"
	var byteValue []byte = ReturnByteValue(jsonPath)

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

	// Main Loop

	var weekdayNo int = int(time.Now().Weekday())
	var accumulator int = ReturnStartPosition(weekdayNo)

	for accumulator < 315 {
		Cls() // clear terminal screen

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

		fmt.Println("Decade:\t\t" + decadeName)
		fmt.Println("Mystery:\t" + mysteryName)
		fmt.Println("Message:\t" + messageText)
		fmt.Println("Scripture:\t" + scriptureText + "\n")
		fmt.Println("Prayer:\t\t" + prayerText)

		fmt.Println("\n---\nPress the Enter to continue / Ctrl+C to Exit")
		fmt.Scanln()
	}
}
