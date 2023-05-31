package circmath

func CircSub[T number](a T, b T, rmin T, rmax T) T {
	result := a - b
	r := rmax - rmin
	for result >= rmax {
		result -= r
	}
	for result < rmin {
		result += r
	}
	return result
}
