package bindcore

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapCount_Len(t *testing.T) {
	mc := MapCount{
		1:15,
		2:10,
		3:20,
		4:25,
	}
	sort.Sort(mc)
	assert.Equal(t, mc, MapCount{
		2:10,
		1:15,
		3:20,
		4:25,
	})
}
