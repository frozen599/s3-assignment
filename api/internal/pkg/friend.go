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
