// Median of Two Sorted Arrays
package MofTSA

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merg []int
	p1 := 0
	p2 := 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] <= nums2[p2] {
			merg = append(merg, nums1[p1])
			p1++
			continue
		}
		merg = append(merg, nums2[p2])
		p2++
	}
	if p1 < len(nums1) {
		merg = append(merg, nums1[p1:]...)
	}
	if p2 < len(nums2) {
		merg = append(merg, nums2[p2:]...)
	}

	fmt.Println(merg)
	if len(merg)%2 == 1 {
		ind := (len(merg) - 1) / 2
		return float64(merg[ind])
	}
	right := len(merg) / 2
	return float64(merg[right]+merg[right-1]) / 2
}

func MedianSortedArraysTest() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 1, 1, 1, 1, 1}, []int{}))
}
