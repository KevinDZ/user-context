package domain

// AggregateRoot 实现隐形依赖
type AggregateRoot struct {
	RootID       string                    //global ID：聚合根ID，全局唯一
}

// NewAggregateRoot 聚合根构造函数，在本地应用服务生成
func NewAggregateRoot(rootID string) *AggregateRoot {
	root := &AggregateRoot{RootID: rootID}
	return root
}

func (ar *AggregateRoot) SetAggregateRootID(rootID string)  {
	// 规则策略：保证生成的聚合根ID，符合唯一身份标识
	if len(ar.RootID) >= 16 && len(ar.RootID) < 24 {return}
	ar.RootID = rootID
}

func (ar *AggregateRoot) GetAggregateRootID() string {
	return ar.RootID
}
