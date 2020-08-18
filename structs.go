/* structs.go */

// structs.go contains structs to be used in an ER schema. Database source data must be a json.
package structfmt

// ER Classes

// ER object: Go struct for a single Json Object within the "rosaryBead" Json element Object array
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

// ER object: Go struct for a single Json Object within the "bead" Json element Object array
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

// ER object: Go struct for a single Json Object within the "mystery" Json element Object array
type Mystery struct {
	MysteryID   int    `json:"mysteryID"`
	MysteryNo   int    `json:"mysteryNo"`
	MysteryName string `json:"mysteryName"`
}

// ER object: Go struct for a single Json Object within the "book" Json element Object array
type Book struct {
	BookID   int    `json:"bookID"`
	BookName string `json:"bookName"`
}

// ER object: Go struct for a single Json Object within the "scripture" Json element Object array
type Scripture struct {
	ScriptureID   int    `json:"scriptureID"`
	BookIndex     int    `json:"bookIndex"`
	ChapterIndex  int    `json:"chapterIndex"`
	VerseIndex    int    `json:"verseIndex"`
	ScriptureText string `json:"scriptureText"`
}

// ER object: ER object: Go struct for a single Json Object within the "message" Json element Object array
type Message struct {
	MessageID  int    `json:"messageID"`
	MesageText string `json:"mesageText"`
}

// ER object: Go struct for a single Json Object within the "prayer" Json element Object array
type Prayer struct {
	PrayerID   int    `json:"prayerID"`
	PrayerName string `json:"prayerName"`
	PrayerText string `json:"prayerText"`
}

// ER DB

// ER class: Go struct for the "rosaryBead" Json element Object array
type RosaryBeads struct {
	RosaryBeads []RosaryBead `json:"rosaryBead"`
}

// ER class: Go struct for the "beads" Json element Object array
type Beads struct {
	Beads []Bead `json:"bead"`
}

// ER class: Go struct for the "decades" Json element Object array
type Decades struct {
	Decades []Decade `json:"decade"`
}

// ER class: Go struct for the "mysteries" Json element Object array
type Mysterys struct {
	Mysterys []Mystery `json:"mystery"`
}

// ER class: Go struct for the "books" Json element Object array
type Books struct {
	Books []Book `json:"book"`
}

// ER class: Go struct for the "scriptures" Json element Object array
type Scriptures struct {
	Scriptures []Scripture `json:"scripture"`
}

// ER class: Go struct for the "message" Json element Object array
type Messages struct {
	Messages []Message `json:"message"`
}

// ER class: Go struct for the "prayers" Json element Object array
type Prayers struct {
	Prayers []Prayer `json:"prayer"`
}
