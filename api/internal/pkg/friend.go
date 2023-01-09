package pkg

import (
	"github.com/frozen599/s3-assignment/api/internal/models"
)

func GetMutualFriendList(list1, list2 []models.Relationship) []int {
	var ret []int
	for _, item1 := range list1 {
		for _, item2 := range list2 {
			if item1.RelationshipType == item2.RelationshipType &&
				(item2.UserID2 == item2.UserID2) {
				ret = append(ret, item1.UserID2)
			}
		}
	}
	return ret
}
