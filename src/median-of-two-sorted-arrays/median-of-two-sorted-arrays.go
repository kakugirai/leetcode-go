package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	iMin, iMax, halfLen := 0, m, (m+n+1)/2

	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				if nums1[i] > nums2[j] {
					minRight = nums2[j]
				} else {
					minRight = nums1[i]
				}
			}

			return float64(maxLeft+minRight) / 2
		}
	}
	return 0.0
}

func main() {
	num1 := []int{1, 3, 8, 9, 15}
	num2 := []int{7, 11, 19, 21, 18, 25}
	fmt.Println(findMedianSortedArrays(num1, num2))
}
