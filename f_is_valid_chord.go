package revoutil

import (
	"regexp"
)

func IsValidChord(s string) bool {
	r := regexp.MustCompile(`^([A-G](#|b)|\d+) (maj|min|aug|dim|sus2|sus4|dom|_) (3|5|7|9|11|13)(( (\+|-|#|b)(3|5|7|9|11|13))?)*$`)
	return r.MatchString(s)
}
