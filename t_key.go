package revoutil

import (
	"fmt"
	"math"
	"strconv"

	"slices"

	"github.com/davi4046/revoutil/circmath"
)

type Key []int

func NewKey(pitch int, scale int) Key {
	binary := []rune(fmt.Sprintf("%12s", strconv.FormatInt(int64(scale), 2)))

	slices.Reverse(binary)

	var pitchClassSet []int

	for i, char := range binary {
		if char == '1' {
			pitchClassSet = append(pitchClassSet, i+pitch)
		}
	}

	return pitchClassSet
}

func (k Key) DegreeToMIDI(degree int) int {
	octave := int(math.Floor(float64(degree)/float64(len(k))) + 4)
	index := circmath.CircAdd(degree, 0, 0, len(k))
	return k[index] + 12*octave
}
