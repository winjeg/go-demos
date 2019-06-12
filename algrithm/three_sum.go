package algrithm


func threeSum(nums []int) [][]int {
	tmpMap := make(map[int]int, 10)
	for i := range nums {
		tmpMap[nums[i]] = i
	}
	result := make([][]int, 0, 10)
	for j := range nums {
		for k := range nums {
			sum := nums[j] + nums[k]
			v, ok := tmpMap[-sum]
			if ok && v != j && v != k && k > j && v > k {
				ts := []int{nums[j], nums[k], -sum}
				if !contains(result, ts) {
					result = append(result, ts)
				}
			}
		}
	}
	return result
}

func contains(c [][]int, e []int) bool {
	for i := range c {
		if duplicate(c[i], e) {
			return true
		}
	}
	return false
}

func duplicate(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 || len(s2) == 0 || s1 == nil || s2 == nil {
		return false
	}
	r := make(map[int]bool, 5)
	for i := range s1 {
		r[s1[i]] = true
	}
	if allZero(s1) && allZero(s2) {
		return true
	}

	if (allZero(s1) && !allZero(s2) && hasZero(s2)) || (allZero(s2) && !allZero(s1) && hasZero(s1)) {
		return false
	}

	equals := true
	for j := range s2 {
		if _, ok := r[s2[j]]; !ok {
			return false
		}
	}
	return equals
}

func hasZero(s []int) bool {
	hasZero := false
	for _, v := range s {
		if v == 0 {
			hasZero = true
			break
		}
	}
	return hasZero
}

func allZero(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
