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
