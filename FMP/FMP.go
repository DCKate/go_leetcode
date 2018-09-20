package FMP

func FindFirstMissingPositive(in []int) int {
	ref := make([]int, len(in))
	for _, vv := range in {
		if vv >= 1 && vv <= len(in) {
			ref[vv-1] = vv
		}
	}
	for ii, vv := range ref {
		if vv == 0 {
			return ii + 1
		}
	}
	return len(in) + 1
}
