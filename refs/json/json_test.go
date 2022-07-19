package json

import (
	"encoding/json"
	"testing"
	"time"
)


func TestJsonMarshal(t *testing.T) {
	m := map[string]interface{}{
		"Aaa":           "Aaa",
		"aaa":           "aaa",
		"Num1":          1,
		"Float1":        2.00,
		"Bool1":         true,
		"Bool2":         false,
		"LONG_SETTING":  "LONG_SETTING",
		"Long_settings": "Long_settings",
		"small_prefix":  "small_prefix",
	}
	d1, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	t.Logf("%s", d1)

	d2, err := json.MarshalIndent(m, "|", "x")
	if err != nil {
		panic(err)
	}
	t.Logf("d2 %s", d2)
}

type RetMessage struct {
	Code   int `json:"code"`
	Errmsg string
	Data   RetData
}

type RetData struct {
	Name      string
	Age       uint
	LotsBooks []string
}

type kv map[string]interface{}

func TestJsonUnMarshal(t *testing.T) {
	// marshal
	succ := kv{
		"code":   0,
		"errmsg": "success",
		"data": kv{
			"name":       "clark",
			"age":        33,
			"lots_books": []string{"计算机网络", "数据机构", "算法导论"},
		},
		"DFloat": 2.00,
	}
	jd1, err := json.Marshal(succ)
	if err != nil {
		panic(err)
	}
	// unmarshal
	jd := new(RetMessage)
	err = json.Unmarshal(jd1, &jd)
	if err != nil {
		panic(err)
	}
	t.Logf("json data: %#v", jd)
	t.Logf("kv age:%T %[1]v", jd.Data.Age)
	t.Logf("kv age:%T %[1]v", jd.Data.Name)

	for i, book := range jd.Data.LotsBooks {
		t.Logf("%d, %s", i, book)
	}
}

type Fav struct {
	title    string
	add_time int64
}

type FavInfos []Fav

func TestJsonUnMarshalMaps(t *testing.T) {
	// marshal
	succ := kv{
		"code":   0,
		"errmsg": "success",
		"data": kv{
			"name":       "clark",
			"age":        33,
			"lots_books": []string{"计算机网络", "数据机构", "算法导论"},
			"fav_infos": []Fav{
				{"read book", 332211},
				{"play game", 6666},
			},
		},
		"DFloat": 2.00,
	}
	jd1, err := json.Marshal(succ)
	if err != nil {
		panic(err)
	}
	// unmarshal
	jd := make(kv)
	err = json.Unmarshal(jd1, &jd)
	if err != nil {
		panic(err)
	}
	t.Logf("json data: %#v", jd)
	t.Logf("code: %T, %[1]v", jd["code"])
	t.Logf("code: %T, %[1]v", jd["code"].(float64))
	t.Logf("code: %T, %[1]v", jd["errmsg"])
	t.Logf("code: %T, %[1]v", jd["data"])
	t.Logf("code2: %T, %[1]v", jd["data"].(kv))

}

func TestJsonFruit(t *testing.T) {
	type Fruit struct {
		Name     string `json:"Name"`
		PriceTag string `json:"PriceTag"`
	}

	type FruitBasket struct {
		Name    string
		Fruit   map[string]Fruit
		Id      int64 `json:"ref"` // 声明对应的json key
		Created time.Time
	}
	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit" : {
	    "1": {
		"Name": "Apple",
		"PriceTag": "$1"
	    },
	    "2": {
		"Name": "Pear",
		"PriceTag": "$1.5"
	    }
        },
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%#v", basket)
	for _, item := range basket.Fruit {
		t.Log(item.Name, item.PriceTag)
	}
}

func TestMarshal2(t *testing.T) {
	// apollo获取智校请求地址
	reqdata := map[string]string{
		"appid":        "app.Appid",
		"secret":       "app.SecretKey",
		"code":         "app.AuthCode",
		"redirect_url": "app.RedirectUrl",
	}
	jdata, err := json.Marshal(reqdata)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", jdata)
	t.Logf("%s", jdata)

	// unmarshal to mapping
	k := make(map[string]string)
	if err = json.Unmarshal(jdata, &k); err != nil {
		t.Errorf("unmarshal fail: %s", err)
	} else {
		t.Logf("%#v", k)
	}

	// unmarshal to struct
	type stt struct {
		Appid       string
		Secret      string
		Code        string `json:"code"`
		RedirectUrl string `json:"redirect_url"`
	}
	p := new(stt)
	if err := json.Unmarshal(jdata, p); err != nil {
		t.Errorf("unmarshal to struct fail: %s", err)
	} else {
		t.Logf("%#v", p)
	}
}

type ConfigDB struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parse_time"`
	Loc       string `json:"loc"`
	LogMode   bool   `json:"log_mode"`
}

func TestJSON(t *testing.T) {

	master := &ConfigDB{
		Host:      "127.0.0.1",
		Port:      "3306",
		User:      "root",
		Password:  "Secrect123.",
		Database:  "ONLINE_EDU_TEST",
		Charset:   "utf8mb4",
		ParseTime: true,
		Loc:       "Local",
		LogMode:   true,
	}

	if b, err := json.Marshal(master); err != nil {
		t.Logf("error %s", err)
	} else {
		t.Logf("%s", b)
	}
}

func TestUnMarshalJSON(t *testing.T) {
	var cfgDefaul = `{
		"host":"127.0.0.1",
		"port":"3306",
		"user":"root",
		"password":"Secrect123.",
		"database":"ONLINE_EDU_TEST",
		"charset":"utf8mb4",
		"parse_time":true,
		"loc":"Local",
		"log_mode":true
	}
	`
	k := new(ConfigDB)
	if err := json.Unmarshal([]byte(cfgDefaul), k); err != nil {
		t.Logf("error %s", err)
	} else {
		t.Logf("%#v", k)
	}
}

func TestJsonMarshalX(t *testing.T) {
	var list = []int{1,2,3}
	marshal, _ := json.Marshal(list)
	t.Logf("list => %s", string(marshal))
}