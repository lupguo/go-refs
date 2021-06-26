package bindcore

type MapCount map[uint64]int

func (m MapCount) Len() int {
	return len(m)
}

func (m MapCount) Less(i, j int) bool {
	return m[uint64(i)] < m[uint64(j)]
}

func (m MapCount) Swap(i, j int) {
	ii, jj := uint64(j), uint64(j)
	m[ii], m[jj] = m[jj], m[ii]
}

