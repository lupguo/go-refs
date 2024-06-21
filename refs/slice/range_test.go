package slice

import (
	"testing"

	"x-learn/advanced/klog/log"
)

type user struct {
	id   uint32
	name string
}

func TestRangeForSlice(t *testing.T) {
	users := []*user{{1, "clark"}, {2, "terry"}}
	userMap := make(map[uint32]*user)

	for _, u := range users {
		userMap[u.id] = u
		t.Logf("u=>%v,type u=>%T &u=%p", u, u, &u)
	}

	t.Logf("user map: %+v", userMap)

	for k, u := range userMap {
		log.Infof("k=>%v, u=>%v", k, u)
	}
}

func TestRangeForSliceV2(t *testing.T) {
	users := []user{{1, "clark"}, {2, "terry"}}
	userMap := make(map[uint32]*user)

	for _, u := range users {
		userMap[u.id] = &u
		t.Logf("u=>%v,type u=>%T &u=%p", u, u, &u)
	}

	t.Logf("user map: %+v", userMap)

	for k, u := range userMap {
		log.Infof("k=>%v, u=>%v", k, u)
	}
}
