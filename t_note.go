package revoutil

type Note struct {
	// Midinote value from 0 to 255.
	Pitch int
	// Duration of the note in whole notes.
	Duration float64
	// Whether or not this note is a pause.
	IsPause bool

	Channel int
	Track   int
}
