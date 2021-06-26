package bindcore

type Account struct {
	UID        uint64 // 账号UID
	UIDType    uint32 // 用户账号类型: 0,2,1001,1002,1005,1008,1012
	RoleType   uint32 // 是否机构账号（0:普通角色，1:机构角色）
	HasAsset   bool   // 有无资产账号（0:无资产，1:有资产）
	AssetInfo  string // 数据资产信息
	DefaultUID uint64
	UnionUID   string
}

type GroupHandler interface {
	// -- 冲突检测
	CrashCheck() bool

	// -- 关键值
	GetAssetAcc(uid uint64) (assUID uint64, assUType uint32)
	GetDefaultAcc() (dftUID uint64, dftUType uint32)
	GetUnionUID(newGenUnion func() string) string

	// -- 资产优先级
	GetAgentAccUID() uint64
	GetAssetUIDs() []uint64
	GetPhoneAccUID() uint64
	GetWxAccUID() uint64
	GetQqAccUID() uint64
}

// Group 组账号
type Group map[uint64]*Account

// NewGroupAccounts 新创建组
func NewGroupAccounts(accs []*Account) Group {
	g := make(Group)
	for _, acc := range accs {
		g[acc.UID] = acc
	}
	return g
}

// CrashCheck 检测是否冲突
func (g Group) CrashCheck() bool {
	ct := make(map[uint32]int)
	for _, acc := range g {
		ct[acc.UIDType]++
	}
	for _, c := range ct {
		if c > 1 {
			return true
		}
	}
	return false
}

// GetAssetAcc 获取uid的应设定的资产账号
func (g Group) GetAssetAcc(uid uint64) (assUID uint64, assUType uint32) {
	acc := g[uid]
	// 账号已是资产账号，则资产账号就是其自身
	if acc.HasAsset || acc.RoleType > 0 {
		return acc.UID, acc.UIDType
	}

	// 空账号，按优先级选择资产账号
	aguid := g.GetAgentAccUID() // 机构
	list := g.GetAssetUIDs()    // 资产账号列表
	phuid := g.GetPhoneAccUID() // 手机号
	wxuid := g.GetWxAccUID()    // 微信UID
	qquid := g.GetQqAccUID()    // QQ

	mscore := make(MapCount)
	for _, acc := range g {
		// 机构优先
		if aguid > 0 && acc.UID == aguid {
			mscore[acc.UID] += 10
		}
		// 资产优先
		if len(list) > 0 && inlist(acc.UID, list) {
			mscore[acc.UID] += 5
		}
		// 手机号优先
		if phuid > 0 && acc.UID == phuid {
			mscore[acc.UID] += 3
		}
		// WX号优先
		if wxuid > 0 && acc.UID == wxuid {
			mscore[acc.UID] += 2
		}
		// QQ号优先
		if qquid > 0 && acc.UID == qquid {
			mscore[acc.UID]++
		}
	}
	// 统计得分
	var muid uint64
	var max int
	for uid, score := range mscore {
		if score > max {
			max = score
			muid = uid
		}
	}
	return g[muid].UID, g[muid].UIDType
}

// inlist 是否在列表中
func inlist(uid uint64, list []uint64) bool {
	for _, u := range list {
		if u == uid {
			return true
		}
	}
	return false
}

// GetDefaultAcc 获取改组账号中的默认资产账号
func (g Group) GetDefaultAcc() (assUID uint64, assUType uint32) {
	// 已设置了默认账号
	for _, acc := range g {
		if acc.DefaultUID > 0 {
			return acc.UID, acc.UIDType
		}
	}

	// 为设置默认资产账号，尝试按按优先级选择默认资产账号
	aguid := g.GetAgentAccUID() // 机构
	list := g.GetAssetUIDs()    // 资产账号列表
	phuid := g.GetPhoneAccUID() // 手机号
	wxuid := g.GetWxAccUID()    // 微信UID
	qquid := g.GetQqAccUID()    // QQ

	mscore := make(MapCount)
	for _, acc := range g {
		// 机构优先
		if aguid > 0 && aguid == acc.UID {
			mscore[acc.UID] += 10
		}
		// 资产优先
		if len(list) > 0 && inlist(acc.UID, list) {
			mscore[acc.UID] += 5
		}
		// 手机号优先
		if phuid > 0 {
			mscore[acc.UID] += 3
		}
		// WX号优先
		if wxuid > 0 {
			mscore[acc.UID] += 2
		}
		// QQ号优先
		if qquid > 0 {
			mscore[acc.UID]++
		}
	}
	// 统计得分
	var muid uint64
	var max int
	for uid, score := range mscore {
		if score > max {
			max = score
			muid = uid
		}
	}
	return g[muid].UID, g[muid].UIDType
}

// GetUnionUID 获取联合账号的UID
func (g Group) GetUnionUID(genUnionUID func() string) string {
	for _, acc := range g {
		if acc.UnionUID != "" {
			return acc.UnionUID
		}
	}
	return genUnionUID()
}

// GetAgentAccUID 获取机构账号
func (g Group) GetAgentAccUID() uint64 {
	for _, acc := range g {
		if acc.RoleType > 0 {
			return acc.UID
		}
	}
	return 0
}

// GetAssetUIDs 获取组账号中的资产账号
func (g Group) GetAssetUIDs() []uint64 {
	var assAccs []uint64
	for _, acc := range g {
		if acc.HasAsset {
			assAccs = append(assAccs, acc.UID)
		}
	}
	return assAccs
}

// GetPhoneAccUID 获取手机账号
func (g Group) GetPhoneAccUID() uint64 {
	for _, acc := range g {
		if acc.UIDType == 1012 {
			return acc.UID
		}
	}
	return 0
}

// GetWxAccUID 获取微信账号
func (g Group) GetWxAccUID() uint64 {
	for _, acc := range g {
		if acc.UIDType == 2 {
			return acc.UID
		}
	}
	return 0
}

// GetQqAccUID 获取QQ账号
func (g Group) GetQqAccUID() uint64 {
	for _, acc := range g {
		if acc.UIDType == 0 {
			return acc.UID
		}
	}
	return 0
}
