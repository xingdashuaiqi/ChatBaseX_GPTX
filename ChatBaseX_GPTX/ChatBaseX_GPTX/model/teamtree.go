package model

type TeamTree struct {
	Id             int64      `gorm:"column:id",json:"id"`
	UserId         int64      `gorm:"column:user_id",json:"user_id"`
	ParentId       int64      `gorm:"column:parent_id",json:"parent_id"`
	CreateTime     string     `gorm:"column:create_time",json:"create_time"`
	Achievements   float64    `gorm:"-"` //个人业绩
	SubPerformance float64    `gorm:"-"` //伞下业绩
	BigMarket      float64    `gorm:"-"` //大市场
	MarketLevel    Level      `gorm:"-"` //市场级别
	Children       []TeamTree `gorm:"-"`
}

func (TeamTree) TableName() string {
	return "team_tree"
}

func (node *TeamTree) CalculateSubPerformance() float64 {
	if len(node.Children) == 0 {
		// 如果节点没有子节点，则其 SubPerformance 等于 Achievements
		return 0
	}

	// 递归计算子节点的 SubPerformance
	subperformance := 0.0
	for i := range node.Children {
		subperformance += node.Children[i].CalculateSubPerformance() + node.Children[i].Achievements

	}

	// 更新当前节点的 SubPerformance
	node.SubPerformance = subperformance

	return subperformance
}
