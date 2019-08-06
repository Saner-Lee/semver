package semver

func (v *version) compare(dst *version) int {
	diff := [LEN]int16{
		int16(v.major - dst.major),
		int16(v.minor - dst.minor),
		int16(v.patch - dst.patch),
	}

	for i := 0; i < LEN; i++ {
		if diff[i] < 0 {
			return -1
		}

		if diff[i] > 0 {
			return 1
		}
	}

	return 0
}

func (v *version) Gt(dst *version) bool {
	return v.compare(dst) > 0
}

func (v *version) Lt(dst *version) bool {
	return v.compare(dst) < 0
}

func (v *version) Eq(dst *version) bool {
	return v.compare(dst) == 0
}

func (v *version) Ge(dst *version) bool {
	return v.compare(dst) >= 0
}

func (v *version) Le(dst *version) bool {
	return v.compare(dst) <= 0
}

// (from,to)
func (v *version) InDoubleCloseRange(from, to *version) bool {
	return v.Gt(from) && v.Lt(to)
}

// [from,to]
func (v *version) InDoubleOpenRange(from, to *version) bool {
	return v.Ge(from) && v.Le(to)
}

// [from,to)
func (v *version) InLeftOpenRange(from, to *version) bool {
	return v.Ge(from) && v.Lt(to)
}

// (from,to]
func (v *version) InRightOpenRange(from, to *version) bool {
	return v.Gt(from) && v.Le(to)
}
