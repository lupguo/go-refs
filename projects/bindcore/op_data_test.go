package bindcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeOpData(t *testing.T) {
	accs := []*Account{
		{
			UID:     280001,
			UIDType: 0,
			RoleType: 31,
			HasAsset: true,
			UnionUID: "yye1-1adf",
		},
		{
			UID:      1442222212,
			UIDType:  2,
			UnionUID: "yye1-1adf",
		},
		{
			UID:       14499999,
			UIDType:   1012,
			AssetInfo: "{order:100}",
			HasAsset: true,
		},
	}
	opds, err := MakeOpData(accs, func() string {
		return "xxx"
	})
	assert.Nil(t, err)
	for _, opd := range opds {
		t.Logf("%+v", *opd)
	}
}
