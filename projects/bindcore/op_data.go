package bindcore

import (
	"github.com/pkg/errors"
)

const Update = "Update"
const Insert = "Insert"

type OpData struct {
	Op             string // 具体DB操作
	UID            uint64 // 账号UID
	UIDType        uint32 // 用户账号类型: 0,2,1001,1002,1005,1008,1012
	AssetUID       uint64
	AssetUIDType   uint32
	RoleType       uint32 // 是否机构账号（0:普通角色，1:机构角色）
	HasAsset       bool   // 有无资产账号（0:无资产，1:有资产）
	AssetInfo      string // 数据资产信息
	DefaultUID     uint64
	DefaultUIDType uint32
	UnionUID       string
}

// getOpHandle 获取DB操作类型
func getOpHandle(unionUID string) string {
	if unionUID != "" {
		return Update
	}
	return Insert
}

var ErrBindTypeCrash = errors.New("the type of account binding conflicts, account type cannot be consistent")

// MakeOpData 创建DB Data
func MakeOpData(accs []*Account, genUnionUIDFn func() string) ([]*OpData, error) {
	g := NewGroupAccounts(accs)
	if g.CrashCheck() {
		return nil, ErrBindTypeCrash
	}
	var data []*OpData
	for _, acc := range g {
		assUID, assUType := g.GetAssetAcc(acc.UID)
		dftUID, dftUType := g.GetDefaultAcc()
		data = append(data, &OpData{
			Op:             getOpHandle(acc.UnionUID),
			UID:            acc.UID,
			UIDType:        acc.UIDType,
			AssetUID:       assUID,
			AssetUIDType:   assUType,
			RoleType:       acc.RoleType,
			HasAsset:       acc.HasAsset,
			AssetInfo:      acc.AssetInfo,
			DefaultUID:     dftUID,
			DefaultUIDType: dftUType,
			UnionUID:       g.GetUnionUID(genUnionUIDFn),
		})
	}
	return data, nil
}
