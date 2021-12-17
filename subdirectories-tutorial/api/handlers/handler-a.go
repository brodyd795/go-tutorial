package handlers

import (
	"fmt"
	"subdirectories/api/services"

	u "subdirectories/utils"
)

func HandlerA() {
	fmt.Println("Handler A")
	services.ServiceA()
	u.UtilX()
}
