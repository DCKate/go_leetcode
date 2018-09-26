package LRinH

func largestRectangleAreaV2(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	if len(heights) == 1 {
		return heights[0]
	}
	large := 0
	lessFromLeft := make([]int, len(heights))
	lessFromRight := make([]int, len(heights))
	lessFromRight[len(heights)-1] = len(heights)
	lessFromLeft[0] = -1
	for ii := 1; ii < len(heights); ii++ {
		p := ii - 1
		for p >= 0 && heights[ii] <= heights[p] {
			// fmt.Println(p, " vv ", vv, " ", heights[p])
			p = lessFromLeft[p]

		}
		lessFromLeft[ii] = p
	}

	for ii := len(heights) - 2; ii >= 0; ii-- {
		p := ii + 1
		for p < len(heights) && heights[ii] <= heights[p] {
			// fmt.Println(p, " vv ", heights[ii], " ", heights[p])
			p = lessFromRight[p]

		}
		lessFromRight[ii] = p
	}
	// fmt.Println("left ", lessFromLeft, " right ", lessFromRight, " value ", heights)
	for ii, vv := range heights {
		width := lessFromRight[ii] - lessFromLeft[ii] - 1
		area := width * vv
		if large < area {
			large = area
		}
	}
	return large
}
