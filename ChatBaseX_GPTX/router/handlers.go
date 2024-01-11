package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 进仓
type DepositRequestDeposit struct {
	UserId      string `json:"UserId" binding:"required"`
	Amount1     string `json:"amount1" binding:"required"`
	Amount2     string `json:"amount2" binding:"required"`
	Amount3     string `json:"amount3" binding:"required"`
	Amount1Rate string `json:"amount1Rate" binding:"required"`
	Amount2Rate string `json:"amount2Rate" binding:"required"`
	Amount3Rate string `json:"amount3Rate" binding:"required"`
}

// 奖励
type DepositRequestAixReward struct {
	UserId int `json:"UserId" binding:"required"`
	Amount int `json:"Amount" binding:"required"`
}

// AIX账户地址
type AIXrewardMapping struct {
	Addresses []string `json:"addresses"`
	Amount    []string `json:"amount"`
}

// DepositHandler 处理用户进仓请求的函数
func DepositHandler(c *gin.Context) {
	// 从请求中解析 DepositRequest 结构体
	var depositRequest DepositRequestDeposit
	if err := c.ShouldBindJSON(&depositRequest); err != nil {
		c.JSON(400, gin.H{"error": "无法解析请求数据", "details": err.Error()})
		return
	}
    
	// 使用 DepositRequest 中的 Amount 和 UserID 参数
	UserId, _ := strconv.ParseFloat(depositRequest.UserId, 64)
	Amount1, _ := strconv.ParseFloat(depositRequest.Amount1, 64)
	Amount2, _ := strconv.ParseFloat(depositRequest.Amount2, 64)
	Amount3, _ := strconv.ParseFloat(depositRequest.Amount3, 64)
	Amount1Rate, _ := strconv.ParseFloat(depositRequest.Amount1Rate, 64)
	Amount2Rate, _ := strconv.ParseFloat(depositRequest.Amount2Rate, 64)
	Amount3Rate, _ := strconv.ParseFloat(depositRequest.Amount3Rate, 64)
  
	//日化后总额
	DateSpaceOne := Amount1 * Amount1Rate
	DateSpaceTwo := Amount2 * Amount2Rate
	DateSpaceThree := Amount3 * Amount3Rate
	totalreward := DateSpaceOne + DateSpaceTwo + DateSpaceThree
	AixReward := AIXrewardHandler(totalreward)
    
	// 构造 JSON 响应
	response := gin.H{
		"userId":      UserId,
		"totalreward": totalreward,
		"AixReward":   AixReward,
	}
    
	// 打印结果
	log.Printf("Usertotalreward: %v\n", response)
	// 打印结果
	fmt.Printf("UserId: %d\n", response["userId"])
	fmt.Printf("TotalReward: %v\n", response["totalreward"])
	// 将结果以 JSON 形式返回给客户端
	c.JSON(http.StatusOK, response)
}
    
// DepositHandler 处理用户进仓请求的函数
func AIXrewardHandler(totalreward float64) float64 {

	//日化后总额
	AixReward := totalreward * 0.2
	return AixReward

}
