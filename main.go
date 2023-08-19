package main

import (
	"fmt"

	"github.com/durotimicodes/xanda_task_R3_D3/server"
)

func main() {

	err := server.StartServer()
	if err != nil {
		fmt.Println("error starting server in main", err)
		return
	}

}
