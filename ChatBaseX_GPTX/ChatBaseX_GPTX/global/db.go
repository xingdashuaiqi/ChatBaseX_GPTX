package global

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

var (
	DBLink *gorm.DB
)

// 连接到数据库
func SetupDBLink() error {
	var err error
	DBLink, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ai?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		// 全局禁用表名复数
		DBLink.SingularTable(true)
		//打开sql日志
		DBLink.LogMode(true)
		return nil
	} else {
		return err
	}
}
