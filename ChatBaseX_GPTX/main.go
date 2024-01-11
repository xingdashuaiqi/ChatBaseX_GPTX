package main

import (
	"ChatBaseX_GPPTX/config"
	"log"
)

func init() {
}

func main() {
	router := config.SetupConfig()
	err := router.Run("192.168.1.33:8090")
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}

}
