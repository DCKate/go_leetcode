// Largest Rectangle in Histogram
// need to improve !!!
package LRinH

import (
	"fmt"
)

type rectUnit struct {
	Data  int
	Count int
}

func checkPassArea(ntmp rectUnit, arr []rectUnit) (lr int, ret []rectUnit) {
	historyLargeCount := 0
	for _, kk := range arr {
		if ntmp.Data > kk.Data {
			kk.Count = kk.Count + ntmp.Count
			ret = append(ret, kk)
			// fmt.Println(ntmp.Data, " >= ", kk.Data, " ", ret)
			continue
		}
		end := kk.Data * kk.Count
		if lr <= end {
			lr = end
		}

		if kk.Count > historyLargeCount {
			historyLargeCount = kk.Count
		}
		// fmt.Println("finish >> ", kk.Data, " : ", kk.Count, " History :", historyLargeCount)
	}
	// only 0 need to be checked
	if ntmp.Data != 0 {
		ntmp.Count = ntmp.Count + historyLargeCount
		// fmt.Println("Add >> ", ntmp.Data, " : ", ntmp.Count)
		ret = append(ret, ntmp)
	}

	// fmt.Println("Final ", ret, " ", lr)
	return
}

func calculateRectArea(heights []int, count []int) int {
	larg := 0
	lr := 0
	var tmp []rectUnit
	tmp = append(tmp, rectUnit{
		Data:  heights[0],
		Count: count[0],
	})
	for ii, vv := range heights[1:] {
		now := rectUnit{
			Data:  vv,
			Count: count[ii+1],
		}
		lr, tmp = checkPassArea(now, tmp)
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

func prepareArray(heights []int) (result []int, count []int) {
	last := heights[0]
	ind := 0
	result = append(result, heights[0])
	count = append(count, 1)
	for _, vv := range heights[1:] {
		if vv == last {
			count[ind]++
			continue
		}
		last = vv
		ind = len(result)
		result = append(result, vv)
		count = append(count, 1)
	}
	// fmt.Println("after prepare ", result, " ", count)
	return
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	if len(heights) == 1 {
		return heights[0]
	}
	data, cout := prepareArray(heights)
	return calculateRectArea(data, cout)
}

func LargeTest() {
	fmt.Println(largestRectangleAreaV2([]int{2, 1, 2, 0, 3, 2, 2, 3}))
	fmt.Println(largestRectangleAreaV2([]int{2, 2, 2, 2, 2, 2, 2, 2}))
	fmt.Println(largestRectangleAreaV2([]int{0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(largestRectangleAreaV2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(largestRectangleAreaV2([]int{1, 2, 3, 4, 5}))
	fmt.Println(largestRectangleAreaV2([]int{5, 4, 3, 2, 1}))
	fmt.Println(largestRectangleAreaV2([]int{3, 2, 1}))
	fmt.Println(largestRectangleAreaV2([]int{3, 6, 5, 7, 4, 8, 1, 0}))
	fmt.Println(largestRectangleAreaV2([]int{1, 2, 3, 100, 5}))
	fmt.Println(largestRectangleAreaV2([]int{3, 5, 5, 2, 5, 5, 6, 6, 4, 4, 1, 1, 2, 5, 5, 6, 6, 4, 1, 3}))
}
