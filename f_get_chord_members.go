package revoutil

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func GetChordMembers(s string) ([]int, error) {
	parts := strings.Split(s, " ")

	var rootPitchClass int

	rootPitchClass, ok := PitchClassMap[parts[0]]
	if !ok {
		i, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid root")
		}
		rootPitchClass = i % 12
	}

	var quality string

	if parts[1] == "_" {
		switch parts[2] {
		case "3", "5":
			quality = "maj"
		case "7", "9", "11", "13":
			quality = "dom"
		default:
			return nil, fmt.Errorf("invalid extension")
		}
	} else {
		quality = parts[1]
	}

	pitchClassSet, ok := qualityMap[quality]
	if !ok {
		return nil, fmt.Errorf("invalid quality")
	}

	for i := range pitchClassSet {
		pitchClassSet[i] += rootPitchClass
	}

	var members []int

	switch parts[2] {
	case "3":
		members = pitchClassSet[:2]
	case "5":
		members = pitchClassSet[:3]
	case "7":
		members = pitchClassSet[:4]
	case "9":
		members = pitchClassSet[:5]
	case "11":
		members = pitchClassSet[:6]
	case "13":
		members = pitchClassSet[:7]
	default:
		return nil, fmt.Errorf("invalid extension")
	}

	// Perform operations
	for _, s := range parts[3:] {

		targetIndex := slices.Index(extensions, string(s[1:])) + 1

		if targetIndex == 0 {
			// Target does not correspond to an extension
			return nil, fmt.Errorf("invalid operation")
		}

		targetPitchClass := pitchClassSet[targetIndex]

		switch string(s[0]) {
		case "+":
			if i := slices.Index(members, targetPitchClass); i == -1 {
				members = append(members, targetPitchClass)
			}
		case "-":
			if i := slices.Index(members, targetPitchClass); i != -1 {
				members = append(members[:i], members[i+1:]...)
			}
		case "#":
			if i := slices.Index(members, targetPitchClass); i != -1 {
				members[i] += 1
			}
		case "b":
			if i := slices.Index(members, targetPitchClass); i != -1 {
				members[i] -= 1
			}
		default:
			return nil, fmt.Errorf("invalid operation")
		}
	}

	return members, nil
}
