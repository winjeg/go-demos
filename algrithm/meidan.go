/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package algrithm

// 总体思想为把两个数组有序合并到新的数组， 并且直接计算结果
// 多用 O(M+N)的空间
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var finalArray []int
	n1Len, n2Len := len(nums1), len(nums2)
	if n1Len == 0 {
		finalArray = nums2
	}
	if n2Len == 0 {
		finalArray = nums1
	}

	if len(finalArray) == 0 {
		if nums1[0] > nums2[n2Len-1] {
			finalArray = make([]int, n1Len+n2Len, n1Len+n2Len)
			copy(finalArray, nums2)
			copy(finalArray[n2Len:], nums1)
		} else if nums1[n1Len-1 ] < nums2[0] {
			finalArray = make([]int, n1Len+n2Len, n1Len+n2Len)
			copy(finalArray, nums1)
			copy(finalArray[n1Len:], nums2)
		} else {
			finalArray = make([]int, 0, n1Len+n2Len)
			j, k := 0, 0
			for i := 0; i < n1Len+n2Len; i++ {
				if j < n1Len && k < n2Len && nums1[j] < nums2[k] {
					finalArray = append(finalArray, nums1[j])
					j++
				} else if j < n1Len && k < n2Len && nums1[j] > nums2[k] {
					finalArray = append(finalArray, nums2[k])
					k++
				} else if j < n1Len && k < n2Len && nums1[j] == nums2[k] {
					finalArray = append(finalArray, nums2[k], nums1[j])
					k++
					j++
				} else if j >= n1Len && k < n2Len {
					finalArray = append(finalArray, nums2[k:]...)
					break
				} else if j < n1Len && k >= n2Len {
					finalArray = append(finalArray, nums1[j:]...)
					break
				}
			}
		}
	}
	if (n1Len+n2Len)%2 == 0 {
		mid := (n1Len + n2Len) / 2
		return float64(finalArray[mid]+finalArray[mid-1]) / 2
	} else {
		return float64(finalArray[(n1Len+n2Len)/2])
	}
}

// 总体思想是， 走到中间的时候停掉， slow记录上一个值， quick记录当前值，
// 这么做的好处就是不用额外的空间， 时间复杂度也仅为O(M+N)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {

	i1, len1, i2, len2 := 0, len(nums1), 0, len(nums2)

	lenAll := len1 + len2

	if lenAll == 0 {
		return 0
	}

	// 类似快慢指针，记录中位数
	slow, quick := 0, 0

	for i := 0; i <= lenAll/2; i++ {
		if i1 == len1 || (i1 < len1 && i2 < len2 && nums1[i1] >= nums2[i2]) {
			slow = quick
			quick = nums2[i2]
			i2++
			continue
		}
		if i2 == len2 || (i1 < len1 && i2 < len2 && nums1[i1] <= nums2[i2]) {
			slow = quick
			quick = nums1[i1]
			i1++
		}
	}

	if lenAll%2 == 0 {
		return float64(slow+quick) / 2.0
	}

	return float64(quick)
}
