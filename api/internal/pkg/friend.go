package pkg

func GetMutualFriendList(list1, list2 []int) []int {
	out := []int{}
	bucket := map[int]bool{}
	for _, i := range list1 {
		for _, j := range list2 {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}
	return out
}

func RemoveDuplicateIDs(ids []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range ids {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func contain(item int, slice []int) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func FindDiff(slice1, slice2 []int) []int {
	var ret []int
	for _, item1 := range slice1 {
		contained := contain(item1, slice2)
		if !contained {
			ret = append(ret, item1)
		}
	}
	return ret
}

func ConcatIntSlices(slices ...[]int) []int {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	result := make([]int, totalLen)

	var i int
	for _, s := range slices {
		i += copy(result[i:], s)
	}
	return result
}
