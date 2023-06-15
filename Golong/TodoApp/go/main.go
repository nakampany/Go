package main

import (
	"fmt"
	"udemy/Golong/TodoApp/go/app/controllers"
	"udemy/Golong/TodoApp/go/app/models"
	"udemy/Golong/TodoApp/go/utils"
)

func init() {
	utils.LoggingSettings("./app.log")
}

// メイン関数
func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
