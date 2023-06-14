package revoutil

type Note struct {
	// Midinote value from 0 to 255.
	Pitch int
	// Start point of the note in whole notes.
	Start float64
	// End point of the note in whole notes.
	End float64
	// Start point of the note in bars.
	StartBar float64
	// End point of the note in bars.
	EndBar float64

	Channel uint
	Track   uint
}

func (n Note) GetDuration() float64 {
	return n.End - n.Start
}
