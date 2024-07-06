package json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BindOneKeyAccReq struct {
	// TinyId         uint64 `protobuf:"varint,1,opt,name=tiny_id,json=tinyId,proto3" json:"tiny_id"`                        // 资产账号TinyID，默认从协议头中获取(客户端不用传递)
	EncryptedPhone string `protobuf:"bytes,2,opt,name=encrypted_phone,json=encryptedPhone,proto3" json:"encrypted_phone"` //加密手机号
	MsgId          string `protobuf:"bytes,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id"`                            // 消息id
	Token          string `protobuf:"bytes,4,opt,name=token,proto3" json:"token"`                                         // 登陆票据
}

func TestMarshalPB(t *testing.T) {
	k := &BindOneKeyAccReq{
		EncryptedPhone: "cc963c614de08b40853457303546308c36d7651a3f3ed3d02ad0db",
		MsgId:          "2102241216583816696086528",
		Token:          "5b117ce3be150c996fdd6c79e95b4e6c",
	}
	marshal, err := json.Marshal(k)
	assert.Nil(t, err)
	t.Logf("%s", marshal)
}
