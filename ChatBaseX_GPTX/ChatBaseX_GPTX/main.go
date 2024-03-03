package main

import (
	"ChatBaseX-GPTX0113/global"
	"ChatBaseX-GPTX0113/router"
	"fmt"
	"log"
)

func init() {
	err := global.SetupDBLink()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("初始化数据库失败 err:%v", err)
	}
}

func main() {
	r := router.Router()
	//"localhost:9090"
	r.Run(":9090")
}
