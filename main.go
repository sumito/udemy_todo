package main

import (
	"fmt"
	"udemy_todo/app/controllers"
	"udemy_todo/app/models"
)

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()

}
