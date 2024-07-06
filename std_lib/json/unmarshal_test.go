package json

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalArray(t *testing.T) {
	str := `[144115351647524390,287853405]`
	var uids []uint64
	err := json.Unmarshal([]byte(str), &uids)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", uids)
}

func TestMarshalArray(t *testing.T) {
	uids := []uint64{1, 2, 3}
	b, err := json.Marshal(uids)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}

type Tenant struct {
	TenantId string `json:"tenant_id"` // 租户ID
	UserName string `json:"user_name"` // 租户名称
	Password string `json:"password"`  // 租户秘钥
}

type User struct {
	ID int `json:"id"`
}

// unmarshal_test.go:61: 0x14000019540, isNil(zere)=false
// unmarshal_test.go:68: unmarshal zero1 ok
// unmarshal_test.go:73: 0x140000195b8, isNil(zere)=false
// unmarshal_test.go:77: unmarshal zero2 ok
// unmarshal_test.go:82: 0x14000010098, isNil(zere)=false
// unmarshal_test.go:86: unmarshal zero3 ok
// unmarshal_test.go:91: 0x0, isNil(zere)=true
// unmarshal_test.go:93: unmarshal zero4 got err,json: Unmarshal(nil *json.User)
// unmarshal_test.go:100: 0x14000019620, isNil(zere)=false
// unmarshal_test.go:104: unmarshal zero5 ok
func TestNilUnmarshal(t *testing.T) {
	type User struct {
		ID int `json:"id"`
	}

	var zero1 User
	t.Logf("%p, isNil(zere)=%t", &zero1, &zero1 == nil)

	// 1. ok
	userBytes := []byte(`{"id":100}`)
	if err := json.Unmarshal(userBytes, &zero1); err != nil {
		t.Errorf("unmarshal zero1 got err,%s", err)
	} else {
		t.Logf("unmarshal zero1 ok")
	}

	// 2. ok
	zero2 := User{}
	t.Logf("%p, isNil(zere)=%t", &zero2, &zero2 == nil)
	if err := json.Unmarshal(userBytes, &zero2); err != nil {
		t.Errorf("unmarshal zero2 got err,%s", err)
	} else {
		t.Logf("unmarshal zero2 ok")
	}

	// 3. ok
	var zero3 *User
	t.Logf("%p, isNil(zere)=%t", &zero3, &zero3 == nil)
	if err := json.Unmarshal(userBytes, &zero3); err != nil {
		t.Errorf("unmarshal zero3 got err,%s", err)
	} else {
		t.Logf("unmarshal zero3 ok")
	}

	// 4. unmarshal got err,json: Unmarshal(nil *json.User)
	var zero4 *User
	t.Logf("%p, isNil(zere)=%t", zero4, zero4 == nil)
	if err := json.Unmarshal(userBytes, zero4); err != nil {
		t.Errorf("unmarshal zero4 got err,%s", err)
	} else {
		t.Logf("unmarshal zero4 ok")
	}

	// 5. ok
	zero5 := &User{}
	t.Logf("%p, isNil(zere)=%t", zero5, zero5 == nil)
	if err := json.Unmarshal(userBytes, zero5); err != nil {
		t.Errorf("unmarshal zero5 got err,%s", err)
	} else {
		t.Logf("unmarshal zero5 ok")
	}

}

func TestUnmashal22(t *testing.T) {
	su := []*User{
		{1},
		{2},
	}

	v, err := json.Marshal(su)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", v)
	// unmarshal

	// pointer
	var sur []*User
	err = json.Unmarshal(v, &sur)
	if err != nil {
		t.Error(err)
	}
	for k, v := range sur {
		t.Logf("sur k=%d,v= %+v", k, v)
	}

	var u2 User
	err = json.Unmarshal([]byte(`{"id":33}`), &u2)
	if err != nil {
		t.Error(err)
	}
	t.Logf("u2=>%v", u2)

}
