package bindcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAccounts(t *testing.T) {
	accs := []*Account{
		{
			UID:      280001,
			UIDType:  0,
			RoleType: 31,
			HasAsset: true,
		},
		{
			UID:     1442222212,
			UIDType: 2,
		},
		{
			UID:       14499999,
			UIDType:   1012,
			AssetInfo: "{order:100}",
			HasAsset:  true,
		},
	}

	qquid := uint64(280001)     // 有资产
	wxuid := uint64(1442222212) // 无资产
	phuid := uint64(14499999)   // 有资产

	g := NewGroupAccounts(accs)
	assert.Equal(t, g.GetAssetUIDs(), []uint64{qquid, phuid})

	assert.Equal(t, g.GetUnionUID(nil), "xxx")
	assert.Equal(t, g.GetPhoneAccUID(), phuid)
	assert.Equal(t, g.GetWxAccUID(), wxuid)
	assert.Equal(t, g.GetQqAccUID(), qquid)

	acc, uType := g.GetAssetAcc(qquid)
	assert.Equal(t, acc, qquid)
	assert.Equal(t, uType, uint32(0))

	defaultAcc, assUType := g.GetDefaultAcc()
	assert.Equal(t, defaultAcc, qquid)
	assert.Equal(t, assUType, uint32(0))
}

func TestCrashCheck(t *testing.T) {
	accs1 := []*Account{
		{
			UID:      280001,
			UIDType:  0,
			RoleType: 31,
			HasAsset: true,
		},
		{
			UID:     1442222212,
			UIDType: 2,
		},
		{
			UID:       14499999,
			UIDType:   1012,
			AssetInfo: "{order:100}",
			HasAsset:  true,
		},
	}
	accs2 := []*Account{
		{
			UID:      280001,
			UIDType:  0,
			RoleType: 31,
			HasAsset: true,
		},
		{
			UID:     1442222212,
			UIDType: 0,
		},
	}

	g1 := NewGroupAccounts(accs1)
	g2 := NewGroupAccounts(accs2)
	assert.Equal(t, g1.CrashCheck(), false)
	assert.Equal(t, g2.CrashCheck(), true)

}
