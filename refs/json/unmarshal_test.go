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

func TestUnmarshal(t *testing.T) {
	teant := new(Tenant)
	s := `{\"tenant_id\":\"tn-c65190017f394836863e83ca429b500c\",\"user_name\":\"contactsAdmin\",\"password\":\"Demo@123\"}`
	err := json.Unmarshal([]byte(s), teant)
	if err != nil {
		t.Logf("err = %#v", err)
	}
	t.Logf("teant=%s", teant)
}

type User struct {
	ID int `json:"id"`
}

//    unmarshal_test.go:34: 0x1400001d540, isNil(zere)=false
//    unmarshal_test.go:41: unmarshal zero1 ok
//    unmarshal_test.go:46: 0x1400001d5b8, isNil(zere)=false
//    unmarshal_test.go:50: unmarshal zero2 ok
//    unmarshal_test.go:55: 0x14000010098, isNil(zere)=false
//    unmarshal_test.go:59: unmarshal zero3 ok
//    unmarshal_test.go:64: 0x0, isNil(zere)=true
//    unmarshal_test.go:66: unmarshal zero4 got err,json: Unmarshal(nil *json.User)
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

	// 3. unmarshal got err,json: Unmarshal(nil *json.User)
	var zero4 *User
	t.Logf("%p, isNil(zere)=%t", zero4, zero4 == nil)
	if err := json.Unmarshal(userBytes, zero4); err != nil {
		t.Errorf("unmarshal zero4 got err,%s", err)
	} else {
		t.Logf("unmarshal zero4 ok")
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
