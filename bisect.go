package bisect

// Finds the zero of the given function between low and high. Assumes function is
// continuous between low and high, and that there is a single zero in the interval
func Bisect(function func(int) int, low, mid, high int) {
	if low == high {
		return high
	}

	val := function(mid)

	if val < 0 {
		return Bisect(function, low, (low+mid)/2, mid-1)
	} else if val > 0 {
		return Bisect(function, mid+1, (mid+high)/2, high)
	} else {
		return mid
	}
}
