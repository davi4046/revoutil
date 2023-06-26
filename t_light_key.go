package revoutil

import (
	"fmt"
	"math"
	"strconv"

	"github.com/davi4046/revoutil/circmath"
)

type LightKey []int

func NewLightKey(pitch int, scale int) LightKey {
	binary := fmt.Sprintf("%12s", strconv.FormatInt(int64(scale), 2))

	var pitchClassSet []int

	for i, char := range binary {
		if char == '1' {
			pitchClassSet = append(pitchClassSet, (i+pitch)%12)
		}
	}

	return pitchClassSet
}

func (k LightKey) DegreeToMIDI(degree int) int {
	octave := int(math.Floor(float64(degree)/float64(len(k))) + 4)
	index := circmath.CircAdd(degree, 0, 0, len(k))
	return k[index] + 12*octave
}
