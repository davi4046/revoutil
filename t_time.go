package revoutil

type Time struct {
	Numerator   uint8
	Denominator uint8
}

func (t Time) GetWholeNotesPerBar() float64 {
	return 1 / float64(t.Denominator) * float64(t.Numerator)
}
