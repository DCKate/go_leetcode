// Largest Rectangle in Histogram
// need to improved !!!
package LRinH

import (
	"fmt"
)

type rectUnit struct {
	Data  int
	Count int
}

func checkPassArea(h int, arr []rectUnit) (lr int, ret []rectUnit) {
	tmp := rectUnit{
		Data:  h,
		Count: 1,
	}
	for _, kk := range arr {
		if h == kk.Data {
			tmp.Count = kk.Count + 1
			continue
		} else if h > kk.Data {
			kk.Count++
			ret = append(ret, kk)
			// fmt.Println(h, " >= ", kk.Data, " ", ret)
			tmp.Count = 1
			continue
		}

		end := kk.Data * kk.Count
		if lr <= end {
			lr = end
		}
		if tmp.Count < kk.Count+1 {
			tmp.Count = kk.Count + 1
		}
	}
	ret = append(ret, tmp)
	// fmt.Println("Final ", ret)
	return
}

func largestRectangleArea(heights []int) int {
	larg := 0
	lr := 0
	var tmp []rectUnit

	if len(heights) == 0 {
		return 0
	}

	if len(heights) == 1 {
		return heights[0]
	}

	tmp = append(tmp, rectUnit{
		Data:  heights[0],
		Count: 1,
	})
	for _, vv := range heights[1:] {
		lr, tmp = checkPassArea(vv, tmp)
		if lr > larg {
			larg = lr
		}
	}

	for _, kk := range tmp {
		end := kk.Data * kk.Count
		if larg <= end {
			larg = end
		}
	}
	return larg
}

func LargeTest() {
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 2, 2, 2, 2}))
	// fmt.Println(largestRectangleArea([]int{1, 2, 3, 4, 5}))
	// fmt.Println(largestRectangleArea([]int{5, 4, 3, 2, 1}))
	// fmt.Println(largestRectangleArea([]int{3, 2, 1}))
	fmt.Println(largestRectangleArea([]int{3, 6, 5, 7, 4, 8, 1, 0}))
	// fmt.Println(largestRectangleArea([]int{1, 2, 3, 100, 5}))
}
