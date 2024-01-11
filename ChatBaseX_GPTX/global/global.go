package global

import (
	"ChatBaseX_GPPTX/config/setting"

	"gorm.io/gorm"
)

var (
	DbConfig       *setting.DbConfig
	DBEngine       *gorm.DB
	ContractConfig *setting.ContractConfig // 新增合约配置
)
