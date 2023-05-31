package circmath

func MaxCircDist[T number](rhs T, lhs T, rmin T, rmax T) T {
	a := CircSub(rhs, lhs, rmin, rmax)
	b := CircSub(lhs, rhs, rmin, rmax)
	if a > b {
		return a
	} else {
		return -b
	}
}
