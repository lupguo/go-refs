package tgorm

import (
	"fmt"
	"git.code.oa.com/rainbow/golang-sdk/confapi"
	"git.code.oa.com/rainbow/golang-sdk/config"
	"git.code.oa.com/rainbow/golang-sdk/log"
	"git.code.oa.com/rainbow/golang-sdk/types"
	"git.code.oa.com/rainbow/golang-sdk/watch"
	"testing"
)

func TestRainbow(t *testing.T) {
	// 一般全局一个实例就够用了
	rainbow, err := confapi.New(
		// types.ConnectStr("ip://9.56.2.161:8080"),
		// types.ConnectStr("cl5://65026305:65536"),
		types.ConnectStr("http://api.rainbow.oa.com:8080"),
		types.FileCachePath("/tmp/trpc_ketang_account_bind.config"),
		types.IsUsingLocalCache(true),
		types.IsUsingFileCache(true),
		types.LogLevel(log.LOG_LEVEL_DEBUG),
		types.LogName("/tmp/trpc_ketang_account_bind.log"),

		// 预拉取这个appid 和group
		types.AppID("df3218f5-bc8e-47b3-92a6-be8694be72cd"),
		types.Groups("trpc_ketang_account_bind"),
		types.EnvName("Test"),
	)
	if err != nil {
		t.Fatalf("[confapi.New]%s\n", err.Error())
	}

	// preload
	err = rainbow.PreLoad()
	if err != nil {
		t.Fatal(err)
	}

	getOpts := make([]types.AssignGetOption, 0)
	getOpts = append(getOpts, types.WithAppID("df3218f5-bc8e-47b3-92a6-be8694be72cd"))
	getOpts = append(getOpts, types.WithGroup("trpc_ketang_account_bind"))

	// get key
	key := "mysql_conf"
	val, err := rainbow.Get(key, getOpts...)
	if err != nil {
		fmt.Printf("[rainbow.Get]%s\n", err.Error())
		return
	}
	t.Log(val)

	// get group
	getOpts = append(getOpts, types.WithGroup("trpc_ketang_account_bind"))
	gval, err := rainbow.GetGroup(getOpts...)
	if err != nil {
		fmt.Printf("[rainbow.Get]%s\n", err.Error())
		return
	}
	t.Log(gval)

	// watchman group
	var watchman = watch.Watcher{
		GetOptions: types.GetOptions{
			AppID: "df3218f5-bc8e-47b3-92a6-be8694be72cd",
			Group: "trpc_ketang_account_bind",
		},
		CB: watchCallBack,
	}
	err = rainbow.AddWatcher(watchman)
	if err != nil {
		t.Fatal(err)
	}
}

// watchCallBack watch call back
func watchCallBack(oldVal watch.Result, newVal []*config.KeyValueItem) error {
	fmt.Printf("\n---------------------\n")
	fmt.Printf("old value:%+v\n", oldVal)
	fmt.Printf("new value:")
	for i := 0; i < len(newVal); i++ {
		fmt.Printf("%+v", *newVal[i])
	}
	fmt.Printf("\n---------------------\n")
	return nil
}