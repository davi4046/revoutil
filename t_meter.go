package revoutil

type Meter struct {
	Numerator   uint8
	Denominator uint8
}

func (t Meter) GetWholeNotesPerBar() float64 {
	return 1 / float64(t.Denominator) * float64(t.Numerator)
}
