package revokey

import (
	"embed"
	"encoding/json"
	"log"
)

type Scale struct {
	Name           string
	Decimal        int
	BinarySequence string

	CommonNames map[string][]string
	// The number of pitches in the scale.
	Cardinality int
	// The tones in this scale, expressed as numbers from 0 to 11.
	PitchClassSet []int
	// A code assigned by theorist Allen Forte, for this pitch class set and all of its transpositional (rotation) and inversional (reflection) transformations.
	ForteNumber string
	// Some scales have rotational symmetry, sometimes known as "limited transposition". If there are any rotational symmetries, these are the intervals of periodicity.
	RotationalSymmetry []int
	// If a scale has an axis of reflective symmetry, then it can transform into itself by inversion. It also implies that the scale has Ridge Tones. Notably an axis of reflection can occur directly on a tone or half way between two tones.
	ReflectionAxes []float64
	// A palindromic scale has the same pattern of intervals both ascending and descending.
	Palindromicity bool
	// A chiral scale can not be transformed into its inverse by rotation. If a scale is chiral, then it has an enantiomorph.
	IsChiral bool
	// A hemitone is two tones separated by a semitone interval. Hemitonia describes how many such hemitones exist.
	Hemitonia int
	// A cohemitone is an instance of two adjacent hemitones. Cohemitonia describes how many such cohemitones exist.
	Cohemitonia int
	// An imperfection is a tone which does not have a perfect fifth above it in the scale. This value is the quantity of imperfections in this scale.
	Imperfections int
	// Describes if this scale is in prime form, using the Starr/Rahn algorithm.
	IsPrime    bool
	Generation struct {
		// Indicates if the scale can be constructed using a generator, and an origin.
		CanConstruct bool
		Generator    int
		Origin       int
	}
	// A deep scale is one where the interval vector has 6 different digits, an indicator of maximum hierarchization.
	IsDeepScale bool
	// Defines the scale as the sequence of intervals between one tone and the next.
	IntervalStructure []int
	// Describes the intervallic content of the scale, read from left to right as the number of occurences of each interval size from semitone, up to six semitones.
	IntervalVector []int
	// First described by Michael Buchler (2001), this is a vector showing the prominence of intervals relative to the maximum and minimum possible for the scale's cardinality. A saturation of 0 means the interval is present minimally, a saturation of 1 means it is the maximum possible.
	ProportionalSaturationVector []float64
	// Describes the specific interval sizes that exist for each generic interval size. Each generic <g> has a spectrum {n,...}. The Spectrum Width is the difference between the highest and lowest values in each spectrum.
	DistributionSpectra map[int][]int
	// Determined by the Distribution Spectra; this is the sum of all spectrum widths divided by the scale cardinality.
	SpectraVariation float64
	// A scale is maximally even if the tones are optimally spaced apart from each other.
	IsMaximallyEven bool
	// A scale is a maximal area set if a polygon described by vertices dodecimetrically placed around a circle produces the maximal interior area for scales of the same cardinality. All maximally even sets have maximal area, but not all maximal area sets are maximally even.
	IsMaximalAreaSet bool
	// Area of the polygon described by vertices placed for each tone of the scale dodecimetrically around a unit circle, ie a circle with radius of 1.
	InteriorArea float64
	// Perimeter of the polygon described by vertices placed for each tone of the scale dodecimetrically around a unit circle.
	PolygonPerimeter float64
	// A scale has Myhill Property if the Distribution Spectra have exactly two specific intervals for every generic interval.
	HasMyhillProperty bool
	// A scale is balanced if the distribution of its tones would satisfy the "centrifuge problem", ie are placed such that it would balance on its centre point.
	IsBalanced bool
	// Ridge Tones are those that appear in all transpositions of a scale upon the members of that scale. Ridge Tones correspond directly with axes of reflective symmetry.
	RidgeTones []int
	// Also known as Rothenberg Propriety, named after its inventor. Propriety describes whether every specific interval is uniquely mapped to a generic interval. A scale is either "Proper", "Strictly Proper", or "Improper".
	Propriety string
	// Defined by Norman Carey (2002), the heteromorphic profile is an ordered triple of (c, a, d) where c is the number of contradictions, a is the number of ambiguities, and d is the number of differences. When c is zero, the scale is Proper. When a is also zero, the scale is Strictly Proper.
	HetermorphicProfile struct {
		Contradictions int
		Ambiguities    int
		Differences    int
	}
	// The Coherence Quotient is a score between 0 and 1, indicating the proportion of coherence failures (ambiguity or contradiction) in the scale, against the maximum possible for a cardinality. A high coherence quotient indicates a less complex scale, whereas a quotient of 0 indicates a maximally complex scale.
	CoherenceQuotient float64
	// The Sameness Quotient is a score between 0 and 1, indicating the proportion of differences in the heteromorphic profile, against the maximum possible for a cardinality. A higher quotient indicates a less complex scale, whereas a quotient of 0 indicates a scale with maximum complexity.
	SamenessQuotient float64
	// These are the common triads (major, minor, augmented and diminished) that you can create from members of this scale.
	CommonTriads map[int]map[string]struct {
		PitchClasses        [3]int
		Degree              int
		Eccentricity        int
		ClosenessCentrality float64
	}
	// Also known as Bi-Triadic Hexatonics (a term coined by mDecks), and related to Generic Modality Compression (a method for guitar by Mick Goodrick and Tim Miller), these are two common triads that when combined use all the tones in this scale.
	TriadPolychords []map[int][]string
	// Modes are the rotational transformation of this scale.
	Modes []int
	// The prime form of this scale.
	Prime               int
	ComplementaryFamily string
	// The inverse of a scale is a reflection using the root as its axis.
	Inverse int
	// Based on the work of Niels Verosky, hierarchizability is the measure of repeated patterns with "place-finding" remainder bits, applied recursively to the binary representation of a scale. For a full explanation, read Niels' paper, Hierarchizability as a Predictor of Scale Candidacy. The variable k is the maximum number of remainders allowed at each level of recursion, for them to count as an increment of hierarchizability. A high hierarchizability score is a good indicator of scale candidacy, ie a measure of usefulness for producing pleasing music. There is a strong correlation between scales with maximal hierarchizability and scales that are in popular use in a variety of world musical traditions.
	Hierarchizability []struct {
		K                 int
		Hierarchizability int
		BreakdownPattern  string
	}
	// Only scales that are chiral will have an enantiomorph.
	Enantiomorph    int
	Transformations []struct {
		Abbreviation string
		Operation    struct {
			A int
			B int
		}
		Result int
	}
	// These are other scales that are similar to this one, created by adding a tone, removing a tone, or moving one note up or down a semitone.
	NearbyScales []int
}

//go:embed scales.json
var f embed.FS

func GetScaleByName(name string) (Scale, bool) {

	file, err := f.Open("scales.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if _, err := decoder.Token(); err != nil {
		log.Fatalln(err)
	}

	for decoder.More() {
		var scale Scale

		if err := decoder.Decode(&scale); err != nil {
			log.Fatalln(err)
		}

		if scale.Name == name {
			return scale, true
		}
	}

	return Scale{}, false
}

func GetScaleByDecimal(decimal int) (Scale, bool) {

	file, err := f.Open("scales.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if _, err := decoder.Token(); err != nil {
		log.Fatalln(err)
	}

	for decoder.More() {
		var scale Scale

		if err := decoder.Decode(&scale); err != nil {
			log.Fatalln(err)
		}

		if scale.Decimal == decimal {
			return scale, true
		}
	}

	return Scale{}, false
}
