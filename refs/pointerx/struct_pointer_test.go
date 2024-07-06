package pointerx

import (
	"testing"
)

type student struct {
	Id   int
	Name string
	fav  *favorite // 爱好
}

type favorite struct {
	Name     string // 喜好
	Type     int    // 1: sport 2: music
	Duration int    // 时间
}

func TestStuUpd(t *testing.T) {
	fav := &favorite{
		Name:     "踢球",
		Type:     1,
		Duration: 2,
	}
	xiaoLi := &student{
		Id:   100,
		Name: "XiaoLi",
		fav:  fav,
	}
	t.Logf("init stu[%v] fav: %v", xiaoLi, fav)
	t.Logf("init stu: %v", xiaoLi.fav)

	// 更新fav内容
	fav.Name = "音乐"
	fav.Type = 2

	t.Logf("after update stu[%v] fav: %v", xiaoLi, fav)
	t.Logf("after update stu: %v", xiaoLi.fav)
}
