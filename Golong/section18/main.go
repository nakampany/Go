package main

import (
	"fmt"
	"log"

	"udemy/Golong/section18/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
	log.Println("test")
}
