package tool

// return 0 if total is 0
func Percent(n, total float64) float64 {
	if total == 0 {
		return 0
	}
	return n / total * 100
}
