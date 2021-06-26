package json

import (
	"encoding/json"
	"testing"
)

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

func TestNilUnmarshal(t *testing.T) {
	var zero *User
	var u = &User{}
	zero = u
	_ = u
	t.Logf("%p, isNil(zere)=%t", zero, zero == nil)

	err := json.Unmarshal([]byte("{\"id\": 100}"), zero)
	if err != nil {
		t.Errorf("unmarshal got err,%s", err)
	}

	t.Logf("zero=%+v", zero)
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
