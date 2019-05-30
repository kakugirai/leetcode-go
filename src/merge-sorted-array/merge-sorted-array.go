package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int
	var result []int
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	for i < m {
		result = append(result, nums1[i])
		i++
	}
	for j < n {
		result = append(result, nums2[j])
		j++
	}
	copy(nums1, result)
}

func main() {
	nums1 := []int{1, 3, 5, 0, 0, 0}
	nums2 := []int{2, 4, 6}
	merge(nums1, 3, nums2, 3)
}
