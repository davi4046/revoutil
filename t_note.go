package revoutil

type Note struct {
	// Midinote value from 0 to 255.
	Value int
	// Duration of the note in whole notes.
	Duration float64
	// Whether or not this note is a pause.

	Channel int
	Track   int

	IsPause bool
}
