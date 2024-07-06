package reflectt

import (
	"encoding/json"
	"testing"
)

type People struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Extend map[string]int
}

type Class struct {
	Id       int       `json:"class_id"`
	Students []*People `json:"students"`
}

// MarshalJSON 自定义序列化方法
func (p *People) MarshalJSON() ([]byte, error) {
	// 创建一个临时 map，用于存储最终的 JSON 对象
	tempMap := map[string]interface{}{
		"ID":   p.ID,
		"Name": p.Name,
	}
	json.Marshal(p)

	// 将 Extend 字段中的键值对添加到临时 map 中
	for k, v := range p.Extend {
		tempMap[k] = v
	}

	// 序列化临时 map
	return json.Marshal(tempMap)
}

// 得到  {"ID":1,"Name":"zhang","Extend":{"sv_abc":1,"sv_def":1}}
func TestUserExtend(t *testing.T) {
	u := &People{
		ID:   1,
		Name: "zhang",
		Extend: map[string]int{
			"sv_abc": 1,
			"sv_def": 1,
		},
	}

	marshal, err := json.Marshal(u)
	if err != nil {
		return
	}

	t.Logf("%s", marshal)
}

func TestUserExtendV2(t *testing.T) {
	c := &Class{
		Id: 100,
		Students: []*People{
			&People{
				ID:   1,
				Name: "zhang",
				Extend: map[string]int{
					"sv_abc": 1,
					"sv_def": 1,
				},
			},
		},
	}

	marshal, err := json.Marshal(u)
	if err != nil {
		return
	}

	t.Logf("%s", marshal)
}

// 期望得到 {"ID":1,"Name":"zhang","sv_abc":1,"sv_def":1}
