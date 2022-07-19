package textx

import (
	"encoding/base64"
	"testing"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"x-learn/protocol/testpb3/person/user"
)

func TestEqualUser(t *testing.T) {
	// userA
	ua := &user.UserA{
		Name:   "Luping",
		Age:    18,
		Desc:   []byte(`good`),
		Weight: 110,
		Male:   true,
		Url:    []string{`https://a.com`, `https://b.com`},
		Data: &user.Data{
			Id:   130,
			Name: "Terry",
		},
	}
	// pb=>json marshal
	jd, err := protojson.Marshal(ua)
	if err != nil {
		t.Error(err)
	}
	t.Logf("jd=>%s", jd)

	// pb=>bytes
	bd, err := proto.Marshal(ua)
	if err != nil {
		t.Error(err)
	}
	userABase64 := base64.StdEncoding.EncodeToString(bd)
	t.Logf("userABase64=>%s", userABase64)

	// unmarshal to userB
	var ub user.UserB
	bdUserB, err := base64.StdEncoding.DecodeString(userABase64)
	err = proto.Unmarshal(bdUserB, &ub)
	if err != nil {
		t.Error(err)
	}
	t.Logf("userB=>%+v", ub)
}
