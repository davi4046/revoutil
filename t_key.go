package revoutil

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Key struct {
	RootPitchClass int
	Scale          Scale
}

func NewKey(pitchName string, scaleName string) *Key {
	pitch, ok := pitchClassMap[pitchName]
	if !ok {
		return nil
	}

	scale, ok := GetScaleByName(scaleName)
	if !ok {
		return nil
	}

	return &Key{
		RootPitchClass: pitch,
		Scale:          scale,
	}
}

func (k Key) GetPitchClassSet() []int {

	var pitchClassSet []int

	for _, v := range k.Scale.PitchClassSet {
		pitchClassSet = append(pitchClassSet, (v+k.RootPitchClass)%12)
	}

	return pitchClassSet
}

func (k Key) RelativeToAbsoluteChord(s string) string {
	parts := strings.Split(s, " ")

	degree, err := strconv.Atoi(parts[0])
	if err != nil {
		return s
	}

	pitchClassSet := k.GetPitchClassSet()

	degree %= len(pitchClassSet)

	parts[0] = fmt.Sprint(pitchClassSet[degree-1])

	if parts[1] == "_" {
		degreeTriads, ok := k.Scale.CommonTriads[degree]
		if ok {
			if len(degreeTriads) == 1 {
				parts[1] = maps.Keys(degreeTriads)[0]
			}
		}
	}

	return strings.Join(parts, " ")
}
