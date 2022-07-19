package json

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

type Person struct {
	Name  string
	Fav   []string
	Class map[string]string
}

type Student struct {
	Xuehao string
	Person Person
}

type PointerStudent struct {
	Xuehao string
	Person *Person
}

func TestMarshalDiffPointer(t *testing.T) {
	type Person struct {
		Name  string
		Fav   []string
		Class map[string]string
	}
	stu1 := &Person{Name: "clark"}
	stu2 := Person{Name: "clark"}
	ms1, _ := json.Marshal(stu1)
	ms2, _ := json.Marshal(stu2)
	t.Logf("ms1=>%s", ms1)
	t.Logf("ms1=>%s", ms2)
}

func TestMarshalStructs(t *testing.T) {
	st1 := &Student{
		Xuehao: "200510",
		Person: Person{
			Name: "Clark",
			Fav:  []string{"游泳", "羽毛球", "电动"},
			Class: map[string]string{
				"xiaoxue":  "汉山小学",
				"chuzhong": "唐田中学",
				"gaozhong": "第九中学",
			},
		},
	}
	if sb1, err := json.Marshal(st1); err != nil {
		t.Error(err)
	} else {
		t.Logf("%s\n", sb1)
	}
}

func TestUnMarshal(t *testing.T) {
	tests := []struct {
		name string
		json string
	}{
		{"t1", `{"Xuehao":"200510","Person":{"Name":"Clark"}}`},
		{"t2", `{"Xuehao":"200510","Person":{"Name":"Clark","Fav":["游泳","羽毛球","电动"],"Class":{"chuzhong":"唐田中学","gaozhong":"第九中学","xiaoxue":"汉山小学"}}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &Student{}
			if err := json.Unmarshal([]byte(tt.json), st); err != nil {
				t.Error(err)
			} else {
				t.Logf("%#v", st)
			}
		})
	}
}

func TestUnMarashalPointer(t *testing.T) {
	tests := []struct {
		name string
		json string
	}{
		{"t1", `{"Xuehao":"200510","Person":{"Name":"Clark"}}`},
		{"t2", `{"Xuehao":"200510","Person":{"Name":"Clark","Fav":["游泳","羽毛球","电动"],"Class":{"chuzhong":"唐田中学","gaozhong":"第九中学","xiaoxue":"汉山小学"}}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &PointerStudent{}
			if err := json.Unmarshal([]byte(tt.json), st); err != nil {
				t.Error(err)
			} else {
				t.Logf("%#v\n%#v\n%#v\n%#v\n", st, st.Person, st.Person.Class, st.Person.Fav)
			}
		})
	}
}

func TestMarshalX(t *testing.T) {
	rds := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "clark",
	})
	ctx := context.Background()
	wset1 := rds.Set(ctx, "tr:slice1", []int{1, 2, 3}, 600*time.Second)
	s, err2 := wset1.Result()
	t.Logf("tr:slice12 write => result=%s, err=%v", s, err2)
	wset2 := rds.Set(ctx, "tr:slice2", &[]int{1, 2, 3}, 600*time.Second)
	s, err2 = wset2.Result()
	t.Logf("tr:slice12 write => result=%s, err=%v", s, err2)
	//

	cmd1 := rds.Get(ctx, "tr:slice1")
	result, err := cmd1.Result()
	t.Logf("tr:slice12 => result=%s, err=%v", result, err)
	cmd2 := rds.Get(ctx, "tr:slice12")
	result, err = cmd2.Result()
	t.Logf("tr:slice12 => result=%s, err=%v", result, err)
}

func TestJsonEncodeX(t *testing.T) {
	tests := []struct {
		name      string
		val       interface{}
		container interface{}
	}{
		// {"slice-marshal-container-struct", []int{1, 2}, []int{}},
		{"slice-marshal-container-pointer", []int{1, 2}, &[]int{}},
		// {"slice-marshal-pointer-struct", &[]int{1,2}, []int{}},
		{"slice-marshal-pointer-pointer", &[]int{1, 2}, &[]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := json.Marshal(tt.val)
			if err != nil {
				t.Errorf("can not marshal val %#v", tt.val)
			}
			t.Logf("marshal val: %s", m)
			// unmarshal
			err = json.Unmarshal(m, tt.container)
			if err != nil {
				t.Errorf("can not unmarshal val %#v", tt.val)
			}
			t.Logf("unmarshal interface: %#v", tt.container)
			rtks := reflect.TypeOf(tt.container).Kind().String()
			rvks := reflect.ValueOf(tt.container).Kind().String()
			rt := reflect.TypeOf(tt.container)
			rve := reflect.ValueOf(tt.container).Elem()
			rvet := reflect.ValueOf(tt.container).Elem().Type()
			t.Logf("assert interface type: %s, value type:%s, %v, elem value:%v, elem type: %v", rtks, rvks, rt, rve, rvet)
		})
	}

}

func TestUnmarshalJson(t *testing.T) {
	testStrs := []string{
		`[]`,
		`[1]`,
		`[1,2]`,
		`[
1,
2,
3]`,
	}
	for _, str := range testStrs {
		var list []uint64
		err := json.Unmarshal([]byte(str), &list)
		if err != nil {
			t.Errorf("unmarsh %v got err %v", str, err)
			continue
		}
		t.Logf("unmarshal succ got list %v", list)
	}
}
