package model

import "time"

type UserAmountRecord struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int64     `gorm:"column:user_id"`
	Status      string    `gorm:"column:status"`
	Type        int       `gorm:"column:type"`
	Amount      float64   `gorm:"column:amount"`
	ExtInfo     string    `gorm:"column:ext_info"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
	IsDeleted   int       `gorm:"column:is_deleted"`
	SkuRecordID int       `gorm:"column:sku_record_id"`
}

func (UserAmountRecord) TableName() string {
	return "user_amount_record"
}
