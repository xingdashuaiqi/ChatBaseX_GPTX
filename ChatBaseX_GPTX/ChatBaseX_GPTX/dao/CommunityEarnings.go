package dao

import (
	"ChatBaseX-GPTX0113/global"
	"ChatBaseX-GPTX0113/model"
	"fmt"
	"sort"
)

var MarketLevels = []model.Level{
	{1, 5000, 15},
	{2, 20000, 25},
	{3, 60000, 35},
	{4, 200000, 45},
	{5, 500000, 55},
	{6, 2000000, 60},
	{7, 5000000, 65},
}

func GetCommunityEarnings(userId int64) (model.TeamTree, error) {
	var distinctParentIDs []int64

	//if err := global.DBLink.Table("team_tree").
	//	Select("DISTINCT user_id").
	//	Where("parent_id == ?", 0).
	//	Pluck("parent_id", &distinctParentIDs).
	//	Error; err != nil {
	//	fmt.Println(" global.DBLink.Table GetCommunityEarnings err:", err.Error())
	//}

	//测试
	if err := global.DBLink.Table("team_tree").
		Select("DISTINCT user_id").
		Where("parent_id = ? and user_id = ?", 0, userId).
		Pluck("parent_id", &distinctParentIDs).
		Error; err != nil {
		fmt.Println(" global.DBLink.Table GetCommunityEarnings err:", err.Error())
	}

	//for _, i := range distinctParentIDs {
	//	var root model.TeamTree
	//	if err := global.DBLink.Where("user_id = ?", i).First(&root).Error; err != nil {
	//		return model.TeamTree{}, err
	//	}
	//	root.Achievements = TotalAmount(root.UserId)
	//	buildTreeRecursive(&root)
	//	return root, nil
	//
	//}
	var root model.TeamTree
	if err := global.DBLink.Where("user_id = ?", userId).First(&root).Error; err != nil {
		return model.TeamTree{}, err
	}
	root.Achievements = TotalAmount(root.UserId)
	err := buildTreeRecursive(&root)
	root.CalculateSubPerformance()
	calculateBigMarket(&root)
	SetMarketLevel(&root)

	fmt.Printf("%f\n", CalculateReturns(&root))
	fmt.Printf("%f\n", CommunityProfitCalc(CalculateReturns(&root), root.MarketLevel.Percentage))

	if err != nil {
		return model.TeamTree{}, err
	}
	return root, nil
}

// 递归构建树形结构的辅助函数
func buildTreeRecursive(node *model.TeamTree) error {
	var children []model.TeamTree
	if err := global.DBLink.Where("parent_id = ?", node.UserId).Find(&children).Error; err != nil {
		return err
	}
	for i := range children {
		children[i].Achievements = TotalAmount(children[i].UserId)
		if err := buildTreeRecursive(&children[i]); err != nil {
			return err
		}
	}

	node.Children = children
	return nil
}

// 获取用户的BigMarket
func calculateBigMarket(node *model.TeamTree) {
	if len(node.Children) == 0 {
		node.BigMarket = 0
	}

	var maxBigMarket []float64
	for i := range node.Children {
		calculateBigMarket(&node.Children[i])
		maxBigMarket = append(maxBigMarket, node.Children[i].Achievements+node.Children[i].SubPerformance)
	}

	node.BigMarket = MaxFloat64(maxBigMarket)

}

// 获取float64最大的值
func MaxFloat64(slice []float64) float64 {
	if len(slice) == 0 {
		return 0
	}

	// 使用 sort 包对切片进行排序
	sort.Float64s(slice)

	// 去除最大值
	return slice[len(slice)-1]
}

// 获取用户总消费
func TotalAmount(userId int64) float64 {
	var totalAmount float64
	row := global.DBLink.Table("user_sku_record").
		Select("COALESCE(SUM(amount), 0) as total_amount").
		Where("user_id = ? AND type = ?", userId, "AI_COLLECT").
		Row()

	if err := row.Scan(&totalAmount); err != nil {
		fmt.Println("TotalAmount row.Scan(&totalAmount) err:", err.Error())
	}

	return totalAmount
}

// 用于获取用户社区等级
func GetMarketLevel(amount float64) model.Level {
	if amount >= MarketLevels[0].Threshold && amount < MarketLevels[1].Threshold {
		return MarketLevels[0]
	}

	if amount >= MarketLevels[1].Threshold && amount < MarketLevels[2].Threshold {
		return MarketLevels[1]
	}

	if amount >= MarketLevels[2].Threshold && amount < MarketLevels[3].Threshold {
		return MarketLevels[2]
	}

	if amount >= MarketLevels[3].Threshold && amount < MarketLevels[4].Threshold {
		return MarketLevels[3]
	}

	if amount >= MarketLevels[4].Threshold && amount < MarketLevels[5].Threshold {
		return MarketLevels[4]
	}

	if amount >= MarketLevels[5].Threshold && amount < MarketLevels[6].Threshold {
		return MarketLevels[5]
	}
	if amount >= MarketLevels[6].Threshold {
		return MarketLevels[6]
	}

	// 如果金额超过最高阈值，则返回最后一个级别
	return model.Level{}
}

// 设置用户的社区等级
func SetMarketLevel(node *model.TeamTree) {
	for i := range node.Children {
		SetMarketLevel(&node.Children[i])
		node.Children[i].MarketLevel = GetMarketLevel(node.Children[i].SubPerformance - node.Children[i].BigMarket)
	}
	node.MarketLevel = GetMarketLevel(node.SubPerformance - node.BigMarket)
}

// 计算社区收益工具
func CommunityProfitCalc(amount float64, percentage float64) float64 {
	return amount * 0.01 / 0.6 * 0.4 * percentage / 100
}

func CalculateReturns(node *model.TeamTree) float64 {
	var amount float64
	for _, i := range node.Children {
		if i.MarketLevel.Name >= node.MarketLevel.Name {
			amount += i.Achievements
			return amount
		}
		if i.MarketLevel.Name < node.MarketLevel.Name {
			amount += i.SubPerformance - i.BigMarket
			amount += i.Achievements
			amount += CalculateReturns(&i)

		}
	}

	return amount
}
