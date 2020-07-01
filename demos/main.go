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

func SetTUIKeyboard() {
    if runtime.GOOS == "linux" {
        fmt.Println("Hello from linux")

        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

    } else {

        fmt.Println( "Exiting because I did not make Windows CLI keyboard inputs." )
        os.Exit(3)
    }
}

func GetTUIKeyboardKeys( keyPress []byte, accumulator *int ) {
    os.Stdin.Read(keyPress)

    var keyString string = string(keyPress)

    switch keyString {
    case "q": // quit
        fmt.Println( "Quit" )
        exec.Command("reset").Run()
        os.Exit(3)

    case "h": // back
        //*accumulator--
        *accumulator = structfmt.PreviousBead(*accumulator)

    case "l": // next
        //*accumulator++
        *accumulator = structfmt.NextBead(*accumulator)

    default:
        fmt.Println( "\n---\nYou entered:\t", keyString, "\n Navigate using the following:\n  Key q:\tquit app\n  Key h:\tback to previous bead\n  Key l:\tprogress to the next bead" )
        fmt.Println("\n---\nPress the Enter to continue / Ctrl+C to Exit")
        fmt.Scanln()
    }
}

// Get column width of terminal display
func ReturnTermboxWidth() int {
    // Requires github.com/nsf/termbox-go import

    if err := termbox.Init(); err != nil {
        panic(err)
    }

    charWidth, _ := termbox.Size()
    termbox.Close()

    return charWidth
}

// Convert strin into an array of words
func CentenceArray(inStr string) []string {
    // requires strings import

    var centenceArr []string = strings.Split(inStr, " ")
    return centenceArr
}

// Perform custom word wrapping based on a defined char width
// Input a string and the number of chars to wrap long strings
func IndentedWrap(strOrig string, charWidth int) string {
    var newString string = ""
    var centenceArr []string = CentenceArray(strOrig)
    var centenceArrLen int = len(centenceArr)

    var charCount int = 0
    var i int = 0
    var wordLength int = 0

    for i = 0; i < centenceArrLen; i++ {

        if charCount < charWidth {
            newString += centenceArr[i] + " "
            wordLength = len(centenceArr[i]) + 2 // add the label offset
        } else {
            charCount = 0
            newString += "\n\t\t" + centenceArr[i] + " "
            wordLength = len(centenceArr[i]) + 7 // add the space and new line formating offset
        }

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
    SetTUIKeyboard()


    // Main loop
    for accumulator < 315 {
        // clear terminal screen
        structfmt.Cls()

        // Update position counter
        //accumulator = structfmt.NextBead(accumulator)

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

        // Column char width of a terminal
        var termWidth int = ReturnTermboxWidth() - 6 // misc width padding

        // View query on cli tui
        fmt.Println("Decade:\t\t" + decadeName)
        fmt.Println("Mystery:\t" + mysteryName)

        messageText = IndentedWrap(messageText, termWidth)
        fmt.Println("Message:\t" + messageText)

        scriptureText = IndentedWrap(scriptureText, termWidth)
        fmt.Println("Scripture:\t" + scriptureText + "\n")

        prayerText = IndentedWrap(prayerText, termWidth)
        fmt.Println("Prayer:\t\t" + prayerText)


        // Pause for read display
        // Keyboard navigation input
        GetTUIKeyboardKeys( keyPress, &accumulator )
    }
}
