package utils

import (
	"github.com/frozen599/s3-assignment/api/internal/models"
)

func GetMutualFriendList(list1, list2 []models.Relationship) []models.Relationship {
	var ret []models.Relationship
	for _, item1 := range list1 {
		for _, item2 := range list2 {
			if item1.ID == item2.ID {
				ret = append(ret, item1)
			}
		}
	}
	return ret
}
