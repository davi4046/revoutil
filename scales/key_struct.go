package scales

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
			// FIND EN MÅDE AT OPBEVARE LANGE BESKEDER
			// fmt.Printf("INFO: The quality of the chord '%s' could not be inferred during conversion from scale-relative to non-scale-relative chord as the quality of the triad built on scale degree %d in the scale '%s' is ambiguous.\n", s, degree, k.Scale.Name)
		}
	}

	return strings.Join(parts, " ")
}
