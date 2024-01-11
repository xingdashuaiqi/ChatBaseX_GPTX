package config

import (
	"ChatBaseX_GPPTX/config/setting"
	"ChatBaseX_GPPTX/global"
	"ChatBaseX_GPPTX/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	vp             *viper.Viper
	ContractConfig *setting.ContractConfig // 新增合约配置字段
}

// NewConfig 函数中初始化合约配置
func NewConfig() (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
    
	// 读取合约配置
	var contractConfig setting.ContractConfig
	if err := vp.UnmarshalKey("Contract", &contractConfig); err != nil {
		return nil, err
	}

	return &Config{vp, &contractConfig}, nil
}

func (config *Config) ReadSection(k string, v interface{}) error {
	err := config.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
func SetupConfig() *gin.Engine {
	conf, err := NewConfig()
	if err != nil {
		log.Panic("NewConfig error: ", err)
	}
    
	//读取数据库配置
	err = conf.ReadSection("Database", &global.DbConfig)
	if err != nil {
		log.Panic("ReadSection - Database error: ", err)
	}
    
	//读取合约配置
	err = conf.ReadSection("Contract", &global.ContractConfig)
	if err != nil {
		log.Panic("ReadSection - Contract error: ", err)
	}

	return router.SetupRouter()
}
