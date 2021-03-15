package tailor

import "regexp"

type point struct {
	x, y float32
}

type extremes struct {
	min, max point
}

func min(a, b float32) float32 {
	if a > b {
		return b
	}
	return a
}

func max(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func getNumbers(s string) []string {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return re.FindAllString(s, -1)
}
