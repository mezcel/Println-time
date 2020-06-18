/* my-structs.go */

package struct-fmt

// ER Classes

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
