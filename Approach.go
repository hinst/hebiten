package hebiten

func ApproachFloat64(x, delta, target float64) float64 {
	if x < target {
		x = x + delta
		if target < x {
			x = target
		}
	} else {
		x = x - delta
		if x < target {
			x = target
		}
	}
	return x
}
