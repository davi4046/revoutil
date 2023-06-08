package revoutil

type GeneratorInput struct {
	Index int
	Key   Key
	Time  Time
	Tempo int
}

type Key struct {
	PitchName string
	ScaleName string
}

type Time struct {
	Numerator   int
	Denominator int
}
