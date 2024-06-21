package maps

import (
	"fmt"
	"testing"

	"x-learn/advance/klog/log"
)

type cmdHandler func() error

func h1() error {
	fmt.Println("h1")
	return nil
}

func h2() error {
	fmt.Println("h2")
	return nil
}

func TestMapIf(t *testing.T) {
	map1 := make(map[string]int)
	mapVal, mapSet := map1["key1"]
	t.Logf("mapVal=%+v, mapSet=%+v", mapVal, mapSet)

	// 插入几个元素
	map1["key1"] = 10
	map1["key2"] = 20
	map1["key3"] = 30

	// 查看k, mapSet
	mapVal, mapSet = map1["key2"]
	t.Logf("mapVal=%+v, mapSet=%+v", mapVal, mapSet)
	mapVal, mapSet = map1["key4"]
	t.Logf("mapVal=%+v, mapSet=%+v", mapVal, mapSet)

	// if v, ok = map[x]type; v为map设定的值，未设定为零值；ok为元素是否设定
	if v, ok := map1["key4"]; !ok {
		t.Logf("map1[key4] not set, map value=>%v, ok=>%t", v, ok)
	} else {
		t.Logf("map1[key4] value => %v", v)
	}

	// for range map出来的k是map的索引，v为map元素值
	for k, v := range map1 {
		t.Logf("k=>%+v, v=>%+v", k, v)
	}

}

func TestUndefined(t *testing.T) {
	cmdMap := make(map[string]cmdHandler)
	cmdMap["h1"] = h1
	cmdMap["h2"] = h2
}

var mmap map[int]*struct{ ID int }

func TestValExist(t *testing.T) {
	if mmap[1] == nil {
		t.Log("equal nil")

		// panic
		// mmap[1] = &struct{ ID int }{ID: 100}

		// must make map
		mmap = make(map[int]*struct{ ID int })
		mmap[1] = &struct{ ID int }{ID: 100}
	}
	t.Logf("here..., %+v", mmap)
}

func TestDeleteMap(t *testing.T) {
	m := make(map[int]int)
	delete(m, 1)
	delete(m, 2)

}

func TestOKMap(t *testing.T) {
	var m map[int]int
	log.Infof("%+v => %p", m, m)
	log.Infof("%+v => %p", m, &m)
	log.Infof("m==nil? %v", m == nil)
	log.Infof("&m==nil? %v", &m == nil)
	if _, ok := m[100]; !ok {
		log.Error("not ok")
	}
	log.Error("here!")
}

func TestNilMapGet(t *testing.T) {
	var m map[int]int
	if v, ok := m[100]; ok {
		t.Logf("v=%v, ok=%v", v, ok)
	} else {
		t.Log("m[100] not exist")
	}
}
